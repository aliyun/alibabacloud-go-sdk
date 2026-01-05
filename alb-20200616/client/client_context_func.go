// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds IP entries to an access control list (ACL).
//
// Description:
//
//	  Each ACL can contain IP addresses or CIDR blocks. Take note of the following limits on ACLs:
//
//	    	- The maximum number of IP entries that can be added to an ACL with each Alibaba Cloud account at a time: 20
//
//	    	- The maximum number of IP entries that each ACL can contain: 1,000
//
//		- **AddEntriesToAcl*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAclEntries](https://help.aliyun.com/document_detail/213616.html) operation to query the status of the task.
//
//	    	- If the ACL is in the **Adding*	- state, the IP entries are being added.
//
//	    	- If the ACL is in the **Available*	- state, the IP entries are added.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddEntriesToAcl"),
		Version:     dara.String("2020-06-16"),
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
// Adds backend servers to a server group.
//
// Description:
//
// *AddServersToServerGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background.
//
// 1.  You can call the [ListServerGroups](https://help.aliyun.com/document_detail/213627.html) operation to query the status of a server group.
//
//   - If a server group is in the **Configuring*	- state, it indicates that the server group is being modified.
//
//   - If a server group is in the **Available*	- state, it indicates that the server group is running.
//
// 2.  You can call the [ListServerGroupServers](https://help.aliyun.com/document_detail/213628.html) operation to query the status of a backend server.
//
//   - If a backend server is in the **Adding*	- state, it indicates that the backend server is being added to a server group.
//
//   - If a backend server is in the **Available*	- state, it indicates that the server is running.
//
// @param request - AddServersToServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddServersToServerGroupResponse
func (client *Client) AddServersToServerGroupWithContext(ctx context.Context, request *AddServersToServerGroupRequest, runtime *dara.RuntimeOptions) (_result *AddServersToServerGroupResponse, _err error) {
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

	if !dara.IsNil(request.ServerGroupId) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	if !dara.IsNil(request.Servers) {
		query["Servers"] = request.Servers
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddServersToServerGroup"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddServersToServerGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies a health check template to a server group.
//
// @param request - ApplyHealthCheckTemplateToServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyHealthCheckTemplateToServerGroupResponse
func (client *Client) ApplyHealthCheckTemplateToServerGroupWithContext(ctx context.Context, request *ApplyHealthCheckTemplateToServerGroupRequest, runtime *dara.RuntimeOptions) (_result *ApplyHealthCheckTemplateToServerGroupResponse, _err error) {
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

	if !dara.IsNil(request.HealthCheckTemplateId) {
		query["HealthCheckTemplateId"] = request.HealthCheckTemplateId
	}

	if !dara.IsNil(request.ServerGroupId) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyHealthCheckTemplateToServerGroup"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyHealthCheckTemplateToServerGroupResponse{}
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
// *DeleteDhcpOptionsSet*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAclRelations](https://help.aliyun.com/document_detail/213618.html) operation to query the status of the task.
//
//   - If an ACL is in the **Associating*	- state, the ACL is being associated with a listener.
//
//   - If an ACL is in the **Associated*	- state, the ACL is associated with a listener.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateAclsWithListener"),
		Version:     dara.String("2020-06-16"),
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
// Associates additional certificates with a listener.
//
// Description:
//
// *AssociateAdditionalCertificatesWithListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetListenerAttribute](https://help.aliyun.com/document_detail/2254865.html) operation to query the status of the task:
//
//   - If the HTTPS or QUIC listener is in the **Associating*	- state, the additional certificates are being associated.
//
//   - If the HTTPS or QUIC listener is in the **Associated*	- state, the additional certificates are associated.
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
	if !dara.IsNil(request.Certificates) {
		query["Certificates"] = request.Certificates
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateAdditionalCertificatesWithListener"),
		Version:     dara.String("2020-06-16"),
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
// Associates an EIP bandwidth plan with an Application Load Balancer (ALB) instance.
//
// Description:
//
// *AttachCommonBandwidthPackageToLoadBalancer*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/214362.html) to query the status of the task.
//
//   - If the ALB instance is in the **Configuring*	- state, the EIP bandwidth plan is being associated with the ALB instance.
//
//   - If the ALB instance is in the **Active*	- state, the EIP bandwidth plan is associated with the ALB instance.
//
// @param request - AttachCommonBandwidthPackageToLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachCommonBandwidthPackageToLoadBalancerResponse
func (client *Client) AttachCommonBandwidthPackageToLoadBalancerWithContext(ctx context.Context, request *AttachCommonBandwidthPackageToLoadBalancerRequest, runtime *dara.RuntimeOptions) (_result *AttachCommonBandwidthPackageToLoadBalancerResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachCommonBandwidthPackageToLoadBalancer"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachCommonBandwidthPackageToLoadBalancerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds the elastic IP address (EIP) and virtual IP address (VIP) of a zone to a DNS record.
//
// Description:
//
// This operation is supported only by Application Load Balancer (ALB) instances that use static IP addresses. Before you call this operation, you must call the StartShiftLoadBalancerZones operation to remove the zone from the ALB instance.
//
// @param request - CancelShiftLoadBalancerZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelShiftLoadBalancerZonesResponse
func (client *Client) CancelShiftLoadBalancerZonesWithContext(ctx context.Context, request *CancelShiftLoadBalancerZonesRequest, runtime *dara.RuntimeOptions) (_result *CancelShiftLoadBalancerZonesResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.ZoneMappings) {
		query["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelShiftLoadBalancerZones"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelShiftLoadBalancerZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates AScript rules.
//
// Description:
//
// ### [](#)Prerequisites
//
//   - A standard or WAF-enabled Application Load Balancer (ALB) instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/214358.html).
//
// ### [](#)Usage notes
//
// **CreateAScripts*	- an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAScripts](https://help.aliyun.com/document_detail/472574.html) operation to query the status of a script.
//
//   - If the script is in the **Creating*	- state, the script is being created.
//
//   - If the script is in the **Available**, the script is created.
//
// @param request - CreateAScriptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAScriptsResponse
func (client *Client) CreateAScriptsWithContext(ctx context.Context, request *CreateAScriptsRequest, runtime *dara.RuntimeOptions) (_result *CreateAScriptsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AScripts) {
		query["AScripts"] = request.AScripts
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAScripts"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAScriptsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL) in a region.
//
// Description:
//
// ## Usage notes
//
// The **CreateAcl*	- operation is asynchronous. After you send a request, the system returns a request ID. However, the operation is still being performed in the system background. You can call the [ListAcls](https://help.aliyun.com/document_detail/213617.html) operation to query the status of an ACL:
//
//   - If an ACL is in the **Creating*	- state, the ACL is being created.
//
//   - If an ACL is in the **Available*	- state, the ACL is created.
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
	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
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
		Version:     dara.String("2020-06-16"),
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
// Creates a health check template in a region.
//
// @param request - CreateHealthCheckTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHealthCheckTemplateResponse
func (client *Client) CreateHealthCheckTemplateWithContext(ctx context.Context, request *CreateHealthCheckTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateHealthCheckTemplateResponse, _err error) {
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

	if !dara.IsNil(request.HealthCheckCodes) {
		query["HealthCheckCodes"] = request.HealthCheckCodes
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckHost) {
		query["HealthCheckHost"] = request.HealthCheckHost
	}

	if !dara.IsNil(request.HealthCheckHttpVersion) {
		query["HealthCheckHttpVersion"] = request.HealthCheckHttpVersion
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckMethod) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !dara.IsNil(request.HealthCheckPath) {
		query["HealthCheckPath"] = request.HealthCheckPath
	}

	if !dara.IsNil(request.HealthCheckProtocol) {
		query["HealthCheckProtocol"] = request.HealthCheckProtocol
	}

	if !dara.IsNil(request.HealthCheckTemplateName) {
		query["HealthCheckTemplateName"] = request.HealthCheckTemplateName
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHealthCheckTemplate"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHealthCheckTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a listener.
//
// Description:
//
// ## Usage notes
//
// **CreateListener*	- is an asynchronous operation. After you call this operation, the system returns a request ID. However, the operation is still being performed in the background. You can call the [GetListenerAttribute](https://help.aliyun.com/document_detail/214353.html) operation to query the status of the HTTP, HTTPS, or QUIC listener.
//
//   - If the HTTP, HTTPS, or QUIC listener is in the **Provisioning*	- state, it indicates that the listener is being created.
//
//   - If the HTTP, HTTPS, or QUIC listener is in the **Running*	- state, it indicates that the listener has been created successfully.
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
	if !dara.IsNil(request.CaCertificates) {
		query["CaCertificates"] = request.CaCertificates
	}

	if !dara.IsNil(request.CaEnabled) {
		query["CaEnabled"] = request.CaEnabled
	}

	if !dara.IsNil(request.Certificates) {
		query["Certificates"] = request.Certificates
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DefaultActions) {
		query["DefaultActions"] = request.DefaultActions
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.GzipEnabled) {
		query["GzipEnabled"] = request.GzipEnabled
	}

	if !dara.IsNil(request.Http2Enabled) {
		query["Http2Enabled"] = request.Http2Enabled
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerDescription) {
		query["ListenerDescription"] = request.ListenerDescription
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.QuicConfig) {
		query["QuicConfig"] = request.QuicConfig
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.SecurityPolicyId) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.XForwardedForConfig) {
		query["XForwardedForConfig"] = request.XForwardedForConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateListener"),
		Version:     dara.String("2020-06-16"),
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
// Creates an Application Load Balancer (ALB) instance in a region.
//
// Description:
//
// *CreateLoadBalancer*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/214362.html) operation to query the status of an ALB instance.
//
//   - If an ALB instance is in the **Provisioning*	- state, it indicates that the ALB instance is being created.
//
//   - If an ALB instance is in the **Active*	- state, it indicates that the ALB instance is created.
//
// @param request - CreateLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerResponse
func (client *Client) CreateLoadBalancerWithContext(ctx context.Context, request *CreateLoadBalancerRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadBalancerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressAllocatedMode) {
		query["AddressAllocatedMode"] = request.AddressAllocatedMode
	}

	if !dara.IsNil(request.AddressIpVersion) {
		query["AddressIpVersion"] = request.AddressIpVersion
	}

	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeletionProtectionEnabled) {
		query["DeletionProtectionEnabled"] = request.DeletionProtectionEnabled
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.LoadBalancerBillingConfig) {
		query["LoadBalancerBillingConfig"] = request.LoadBalancerBillingConfig
	}

	if !dara.IsNil(request.LoadBalancerEdition) {
		query["LoadBalancerEdition"] = request.LoadBalancerEdition
	}

	if !dara.IsNil(request.LoadBalancerName) {
		query["LoadBalancerName"] = request.LoadBalancerName
	}

	if !dara.IsNil(request.ModificationProtectionConfig) {
		query["ModificationProtectionConfig"] = request.ModificationProtectionConfig
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneMappings) {
		query["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancer"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a forwarding rule for a listener.
//
// Description:
//
// Take note of the following limits:
//
//   - When you configure the **Redirect*	- action, you can use the default value only for the **HttpCode*	- parameter. Do not use the default values for the other parameters.
//
//   - If you specify the **Rewrite*	- action together with other actions in a forwarding rule, make sure that the **ForwardGroup*	- action is specified.
//
//   - **CreateRule*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListRules](https://help.aliyun.com/document_detail/214379.html) operation to query the status of a forwarding rule.
//
//   - If a forwarding rule is in the **Provisioning*	- state, the forwarding rule is being created.
//
//   - If a forwarding rule is in the **Available*	- state, the forwarding rule is created.
//
//   - You can set **RuleConditions*	- and **RuleActions*	- to add conditions and actions to a forwarding rule. The limits on conditions and actions are:
//
//   - Limits on conditions: 5 for a basic Application Load Balancer (ALB) instance, 10 for a standard ALB instance, and 10 for a WAF-enabled ALB instance.
//
//   - Limits on actions: 3 for a basic ALB instance, 5 for a standard ALB instance, and 5 for a WAF-enabled ALB instance.
//
// @param request - CreateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRuleResponse
func (client *Client) CreateRuleWithContext(ctx context.Context, request *CreateRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
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

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RuleActions) {
		query["RuleActions"] = request.RuleActions
	}

	if !dara.IsNil(request.RuleConditions) {
		query["RuleConditions"] = request.RuleConditions
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRule"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multiple forwarding rules at a time.
//
// Description:
//
// When you call this operation, take note of the following limits:
//
//   - When you configure the **Redirect*	- action, do not use the default values for parameters other than **HttpCode**.
//
//   - If you specify multiple actions in a forward rule, you must specify the **ForwardGroup*	- parameter along with the **Rewrite*	- parameter.
//
//   - **CreateRules*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListRules](https://help.aliyun.com/document_detail/214379.html) operation to query the status of the forwarding rules.
//
//   - If the forwarding rules are in the **Provisioning*	- state, the forwarding rules are being created.
//
//   - If the forwarding rules are in the **Available*	- state, the forwarding rules are created.
//
//   - You can set **RuleConditions*	- and **RuleActions*	- to add conditions and actions to a forwarding rule. Take note of the following limits on the number of conditions and the number of actions in each forwarding rule:
//
//   - Conditions: 5 for each basic ALB instance, 10 for each standard ALB instance, and 10 for each WAF-enabled ALB instance.
//
//   - Actions: 3 for each basic ALB instance, 5 for each standard ALB instance, and 5 for each WAF-enabled ALB instance.
//
// @param request - CreateRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRulesResponse
func (client *Client) CreateRulesWithContext(ctx context.Context, request *CreateRulesRequest, runtime *dara.RuntimeOptions) (_result *CreateRulesResponse, _err error) {
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

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Rules) {
		bodyFlat["Rules"] = request.Rules
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRules"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom security policy in a region.
//
// @param request - CreateSecurityPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityPolicyResponse
func (client *Client) CreateSecurityPolicyWithContext(ctx context.Context, request *CreateSecurityPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ciphers) {
		query["Ciphers"] = request.Ciphers
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityPolicyName) {
		query["SecurityPolicyName"] = request.SecurityPolicyName
	}

	if !dara.IsNil(request.TLSVersions) {
		query["TLSVersions"] = request.TLSVersions
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecurityPolicy"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecurityPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a server group in a region.
//
// Description:
//
// *CreateServerGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [ListServerGroups](https://help.aliyun.com/document_detail/213627.html) to query the status of a server group.
//
//   - If a server group is in the **Creating*	- state, it indicates that the server group is being created.
//
//   - If a server group is in the **Available*	- state, it indicates that the server group is created.
//
// @param request - CreateServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServerGroupResponse
func (client *Client) CreateServerGroupWithContext(ctx context.Context, request *CreateServerGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateServerGroupResponse, _err error) {
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

	if !dara.IsNil(request.ConnectionDrainConfig) {
		query["ConnectionDrainConfig"] = request.ConnectionDrainConfig
	}

	if !dara.IsNil(request.CrossZoneEnabled) {
		query["CrossZoneEnabled"] = request.CrossZoneEnabled
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.HealthCheckConfig) {
		query["HealthCheckConfig"] = request.HealthCheckConfig
	}

	if !dara.IsNil(request.Ipv6Enabled) {
		query["Ipv6Enabled"] = request.Ipv6Enabled
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.ServerGroupName) {
		query["ServerGroupName"] = request.ServerGroupName
	}

	if !dara.IsNil(request.ServerGroupType) {
		query["ServerGroupType"] = request.ServerGroupType
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.SlowStartConfig) {
		query["SlowStartConfig"] = request.SlowStartConfig
	}

	if !dara.IsNil(request.StickySessionConfig) {
		query["StickySessionConfig"] = request.StickySessionConfig
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UchConfig) {
		query["UchConfig"] = request.UchConfig
	}

	if !dara.IsNil(request.UpstreamKeepaliveEnabled) {
		query["UpstreamKeepaliveEnabled"] = request.UpstreamKeepaliveEnabled
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServerGroup"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServerGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes AScript rules.
//
// Description:
//
// *DeleteAScripts*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAScripts](https://help.aliyun.com/document_detail/472574.html) operation to query the status of the task:
//
//   - If an AScript rule is in the **Deleting*	- state, the AScript rule is being deleted.
//
//   - If an AScript rule cannot be found, the AScript rule is deleted.
//
// @param request - DeleteAScriptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAScriptsResponse
func (client *Client) DeleteAScriptsWithContext(ctx context.Context, request *DeleteAScriptsRequest, runtime *dara.RuntimeOptions) (_result *DeleteAScriptsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AScriptIds) {
		query["AScriptIds"] = request.AScriptIds
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAScripts"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAScriptsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access control list (ACL).
//
// Description:
//
// *DeleteAcl*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAcls](https://help.aliyun.com/document_detail/213617.html) operation to query the status of the task.
//
//   - If the ACL is in the **Deleting*	- state, the ACL is being deleted.
//
//   - If the ACL cannot be found, the ACL is deleted.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAcl"),
		Version:     dara.String("2020-06-16"),
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
// Deletes health check templates.
//
// @param request - DeleteHealthCheckTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHealthCheckTemplatesResponse
func (client *Client) DeleteHealthCheckTemplatesWithContext(ctx context.Context, request *DeleteHealthCheckTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DeleteHealthCheckTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.HealthCheckTemplateIds) {
		query["HealthCheckTemplateIds"] = request.HealthCheckTemplateIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHealthCheckTemplates"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHealthCheckTemplatesResponse{}
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
// *DeleteListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [GetListenerAttribute](https://help.aliyun.com/document_detail/2254865.html) to query the status of the task.
//
//   - If the listener is in the **Deleting*	- state, the listener is being deleted.
//
//   - If the listener cannot be found, the listener is deleted.
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
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteListener"),
		Version:     dara.String("2020-06-16"),
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
// Deletes an Application Load Balancer (ALB) instance.
//
// Description:
//
// *DeleteLoadBalancer*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/214362.html) to query the status of the task.
//
//   - If an ALB instance is in the **Deleting*	- state, the ALB instance is being deleted.
//
//   - If an ALB instance cannot be found, the ALB instance is deleted.
//
// @param request - DeleteLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLoadBalancerResponse
func (client *Client) DeleteLoadBalancerWithContext(ctx context.Context, request *DeleteLoadBalancerRequest, runtime *dara.RuntimeOptions) (_result *DeleteLoadBalancerResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLoadBalancer"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a forwarding rule.
//
// Description:
//
// *DeleteRule*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListRules](https://help.aliyun.com/document_detail/214379.html) operation to query the status of a forwarding rule:
//
//   - If the forwarding rule is in the **Deleting*	- state, the forwarding rule is being deleted.
//
//   - If the forwarding rule cannot be found, the forwarding rule is deleted.
//
// @param request - DeleteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleResponse
func (client *Client) DeleteRuleWithContext(ctx context.Context, request *DeleteRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
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

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRule"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more forwarding rules from a listener at a time.
//
// Description:
//
// *DeleteRules*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListRules](https://help.aliyun.com/document_detail/214379.html) operation to query the status of forwarding rules.
//
//   - If the forwarding rules are in the **Deleting*	- state, the forwarding rules are being deleted.
//
//   - If the forwarding rules cannot be found, the forwarding rules are deleted.
//
// @param request - DeleteRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRulesResponse
func (client *Client) DeleteRulesWithContext(ctx context.Context, request *DeleteRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteRulesResponse, _err error) {
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

	if !dara.IsNil(request.RuleIds) {
		query["RuleIds"] = request.RuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRules"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom security policy.
//
// @param request - DeleteSecurityPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityPolicyResponse
func (client *Client) DeleteSecurityPolicyWithContext(ctx context.Context, request *DeleteSecurityPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityPolicyResponse, _err error) {
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

	if !dara.IsNil(request.SecurityPolicyId) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecurityPolicy"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a server group.
//
// Description:
//
// *DeleteServerGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListServerGroups](https://help.aliyun.com/document_detail/213627.html) operation to query the status of the task.
//
//   - If a server group is in the **Deleting*	- state, it indicates that the server group is being deleted.
//
//   - If a specified server group cannot be found, it indicates that the server group has been deleted.
//
// @param request - DeleteServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServerGroupResponse
func (client *Client) DeleteServerGroupWithContext(ctx context.Context, request *DeleteServerGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteServerGroupResponse, _err error) {
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

	if !dara.IsNil(request.ServerGroupId) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServerGroup"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServerGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资源预留
//
// @param request - DescribeCapacityReservationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCapacityReservationResponse
func (client *Client) DescribeCapacityReservationWithContext(ctx context.Context, request *DescribeCapacityReservationRequest, runtime *dara.RuntimeOptions) (_result *DescribeCapacityReservationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCapacityReservation"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCapacityReservationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available regions.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2020-06-16"),
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
// Queries zones in a region.
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
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2020-06-16"),
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
// Disassociates an elastic IP address (EIP) bandwidth plan from an Application Load Balancer (ALB) instance.
//
// Description:
//
// *DetachCommonBandwidthPackageFromLoadBalancer*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/214359.html) operation to query the status of the task.
//
//   - If an ALB instance is in the **Configuring*	- state, the EIP bandwidth plan is being disassociated from the ALB instance.
//
//   - If an ALB instance is in the **Active*	- state, the EIP bandwidth plan is disassociated from the ALB instance.
//
// @param request - DetachCommonBandwidthPackageFromLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachCommonBandwidthPackageFromLoadBalancerResponse
func (client *Client) DetachCommonBandwidthPackageFromLoadBalancerWithContext(ctx context.Context, request *DetachCommonBandwidthPackageFromLoadBalancerRequest, runtime *dara.RuntimeOptions) (_result *DetachCommonBandwidthPackageFromLoadBalancerResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachCommonBandwidthPackageFromLoadBalancer"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachCommonBandwidthPackageFromLoadBalancerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables deletion protection for an Application Load Balancer (ALB) instance.
//
// @param request - DisableDeletionProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDeletionProtectionResponse
func (client *Client) DisableDeletionProtectionWithContext(ctx context.Context, request *DisableDeletionProtectionRequest, runtime *dara.RuntimeOptions) (_result *DisableDeletionProtectionResponse, _err error) {
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

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableDeletionProtection"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDeletionProtectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the access log feature for a Server Load Balancer (SLB) instance.
//
// @param request - DisableLoadBalancerAccessLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableLoadBalancerAccessLogResponse
func (client *Client) DisableLoadBalancerAccessLogWithContext(ctx context.Context, request *DisableLoadBalancerAccessLogRequest, runtime *dara.RuntimeOptions) (_result *DisableLoadBalancerAccessLogResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableLoadBalancerAccessLog"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableLoadBalancerAccessLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the type of the IPv6 address that is used by a dual-stack Application Load Balancer (ALB) instance from public to private.
//
// Description:
//
// ### Prerequisites
//
// An ALB instance is created and IPv4/IPv6 dual stack is enabled for the instance. You can call the [CreateLoadBalancer](https://help.aliyun.com/document_detail/214358.html) operation and set **AddressIpVersion*	- to **DualStack*	- to create a dual-stack ALB instance.
//
// > If you set **AddressIpVersion*	- to **DualStack**:
//
//   - If you set **AddressType*	- to **Internet**, the ALB instance uses a public IPv4 IP address and a private IPv6 address.
//
//   - If you set **AddressType*	- to **Intranet**, the ALB instance uses a private IPv4 IP address and a private IPv6 address.
//
// ### Description
//
//   - After the DisableLoadBalancerIpv6Internet operation is called, the value of **Ipv6AddressType*	- is changed to **Intranet*	- and the type of the IPv6 address of the ALB instance is changed from public to private. If you upgrade the instance or the instance scales elastic network interfaces (ENIs) along with workloads, private IPv6 addresses are automatically enabled for the instance and the new ENIs. You can call the [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/214362.html) operation to query the value of **Ipv6AddressType**.
//
//   - **DisableLoadBalancerIpv6Internet*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/214362.html) operation to query the status of the task.
//
//   - If the ALB instance is in the **Configuring*	- state, the network type of the IPv6 address that is used by the ALB instance is being changed.
//
//   - If the ALB instance is in the **Active*	- state, the network type of the IPv6 address that is used by the ALB instance is changed.
//
// @param request - DisableLoadBalancerIpv6InternetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableLoadBalancerIpv6InternetResponse
func (client *Client) DisableLoadBalancerIpv6InternetWithContext(ctx context.Context, request *DisableLoadBalancerIpv6InternetRequest, runtime *dara.RuntimeOptions) (_result *DisableLoadBalancerIpv6InternetResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableLoadBalancerIpv6Internet"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableLoadBalancerIpv6InternetResponse{}
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
// *DissociateAclsFromListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [ListAclRelations](https://help.aliyun.com/document_detail/213618.html) to query the status of an ACL.
//
//   - If an ACL is in the **Dissociating*	- state, the ACL is being disassociated from the listener.
//
//   - If an ACL is in the **Dissociated*	- state, the ACL is disassociated from the listener.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateAclsFromListener"),
		Version:     dara.String("2020-06-16"),
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
// Disassociates additional certificates from a listener.
//
// Description:
//
// *DissociateAdditionalCertificatesFromListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListListenerCertificates](https://help.aliyun.com/document_detail/214354.html) operation to query the status of the task. - If an additional certificate is in the **Dissociating*	- state, the additional certificate is being disassociated. - If an additional certificate is in the **Dissociated*	- state, the additional certificate is disassociated.
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
	if !dara.IsNil(request.Certificates) {
		query["Certificates"] = request.Certificates
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateAdditionalCertificatesFromListener"),
		Version:     dara.String("2020-06-16"),
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
// Enables deletion protection for a resource.
//
// @param request - EnableDeletionProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDeletionProtectionResponse
func (client *Client) EnableDeletionProtectionWithContext(ctx context.Context, request *EnableDeletionProtectionRequest, runtime *dara.RuntimeOptions) (_result *EnableDeletionProtectionResponse, _err error) {
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

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableDeletionProtection"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDeletionProtectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the access log feature for an Application Load Balancer (ALB) instance.
//
// @param request - EnableLoadBalancerAccessLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableLoadBalancerAccessLogResponse
func (client *Client) EnableLoadBalancerAccessLogWithContext(ctx context.Context, request *EnableLoadBalancerAccessLogRequest, runtime *dara.RuntimeOptions) (_result *EnableLoadBalancerAccessLogResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LogProject) {
		query["LogProject"] = request.LogProject
	}

	if !dara.IsNil(request.LogStore) {
		query["LogStore"] = request.LogStore
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableLoadBalancerAccessLog"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableLoadBalancerAccessLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the type of the IPv6 address that is used by a dual-stack Application Load Balancer (ALB) instance from private to public.
//
// Description:
//
// ### Prerequisites
//
// An ALB instance is created and IPv4/IPv6 dual stack is enabled for the instance. You can call the [CreateLoadBalancer](https://help.aliyun.com/document_detail/214358.html) operation and set **AddressIpVersion*	- to **DualStack*	- to create a dual-stack ALB instance.
//
// > If you set **AddressIpVersion*	- to **DualStack**:
//
//   - If you set **AddressType*	- to **Internet**, the ALB instance uses a public IPv4 IP address and a private IPv6 address.
//
//   - If you set **AddressType*	- to **Intranet**, the ALB instance uses a private IPv4 IP address and a private IPv6 address.
//
// ### Description
//
//   - After the EnableLoadBalancerIpv6Internet operation is called, the value of **Ipv6AddressType*	- is changed to **Internet*	- and the type of the IPv6 address of the ALB instance is changed from private to public. If you upgrade the instance or the instance scales elastic network interfaces (ENIs) along with workloads, public IPv6 addresses are automatically enabled for the instance and the new ENIs. You can call the [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/214362.html) operation to query the value of **Ipv6AddressType**.
//
//   - **EnableLoadBalancerIpv6Internet*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/214362.html) operation to query the status of the task.
//
//   - If the ALB instance is in the **Configuring*	- state, the network type of the IPv6 address that is used by the ALB instance is being changed.
//
//   - If the ALB instance is in the **Active*	- state, the network type of the IPv6 address that is used by the ALB instance is changed.
//
// @param request - EnableLoadBalancerIpv6InternetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableLoadBalancerIpv6InternetResponse
func (client *Client) EnableLoadBalancerIpv6InternetWithContext(ctx context.Context, request *EnableLoadBalancerIpv6InternetRequest, runtime *dara.RuntimeOptions) (_result *EnableLoadBalancerIpv6InternetResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableLoadBalancerIpv6Internet"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableLoadBalancerIpv6InternetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a health check template.
//
// @param request - GetHealthCheckTemplateAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHealthCheckTemplateAttributeResponse
func (client *Client) GetHealthCheckTemplateAttributeWithContext(ctx context.Context, request *GetHealthCheckTemplateAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetHealthCheckTemplateAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HealthCheckTemplateId) {
		query["HealthCheckTemplateId"] = request.HealthCheckTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHealthCheckTemplateAttribute"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHealthCheckTemplateAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a listener.
//
// @param request - GetListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetListenerAttributeResponse
func (client *Client) GetListenerAttributeWithContext(ctx context.Context, request *GetListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetListenerAttributeResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetListenerAttribute"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetListenerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the health check status of a listener and its forwarding rules.
//
// @param request - GetListenerHealthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetListenerHealthStatusResponse
func (client *Client) GetListenerHealthStatusWithContext(ctx context.Context, request *GetListenerHealthStatusRequest, runtime *dara.RuntimeOptions) (_result *GetListenerHealthStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludeRule) {
		query["IncludeRule"] = request.IncludeRule
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetListenerHealthStatus"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetListenerHealthStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an Application Load Balancer (ALB) instance.
//
// @param request - GetLoadBalancerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoadBalancerAttributeResponse
func (client *Client) GetLoadBalancerAttributeWithContext(ctx context.Context, request *GetLoadBalancerAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetLoadBalancerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLoadBalancerAttribute"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLoadBalancerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries AScript rules.
//
// @param request - ListAScriptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAScriptsResponse
func (client *Client) ListAScriptsWithContext(ctx context.Context, request *ListAScriptsRequest, runtime *dara.RuntimeOptions) (_result *ListAScriptsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AScriptIds) {
		query["AScriptIds"] = request.AScriptIds
	}

	if !dara.IsNil(request.AScriptNames) {
		query["AScriptNames"] = request.AScriptNames
	}

	if !dara.IsNil(request.ListenerIds) {
		query["ListenerIds"] = request.ListenerIds
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
		Action:      dara.String("ListAScripts"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAScriptsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the entries of an access control list (ACL).
//
// @param request - ListAclEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAclEntriesResponse
func (client *Client) ListAclEntriesWithContext(ctx context.Context, request *ListAclEntriesRequest, runtime *dara.RuntimeOptions) (_result *ListAclEntriesResponse, _err error) {
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
		Action:      dara.String("ListAclEntries"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAclEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the listeners that are associated with access control lists (ACLs).
//
// @param request - ListAclRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAclRelationsResponse
func (client *Client) ListAclRelationsWithContext(ctx context.Context, request *ListAclRelationsRequest, runtime *dara.RuntimeOptions) (_result *ListAclRelationsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAclRelations"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAclRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access control lists (ACLs) in a region.
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

	if !dara.IsNil(request.AclNames) {
		query["AclNames"] = request.AclNames
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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
		Version:     dara.String("2020-06-16"),
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
// Queries asynchronous tasks in a region.
//
// @param request - ListAsynJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsynJobsResponse
func (client *Client) ListAsynJobsWithContext(ctx context.Context, request *ListAsynJobsRequest, runtime *dara.RuntimeOptions) (_result *ListAsynJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAsynJobs"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAsynJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries health check templates in a region.
//
// @param request - ListHealthCheckTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHealthCheckTemplatesResponse
func (client *Client) ListHealthCheckTemplatesWithContext(ctx context.Context, request *ListHealthCheckTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListHealthCheckTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HealthCheckTemplateIds) {
		query["HealthCheckTemplateIds"] = request.HealthCheckTemplateIds
	}

	if !dara.IsNil(request.HealthCheckTemplateNames) {
		query["HealthCheckTemplateNames"] = request.HealthCheckTemplateNames
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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
		Action:      dara.String("ListHealthCheckTemplates"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHealthCheckTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificates that are associated with a listener, including additional certificates and the default certificate.
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
	if !dara.IsNil(request.CertificateIds) {
		query["CertificateIds"] = request.CertificateIds
	}

	if !dara.IsNil(request.CertificateType) {
		query["CertificateType"] = request.CertificateType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListListenerCertificates"),
		Version:     dara.String("2020-06-16"),
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
// Queries the listeners in a region.
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
	if !dara.IsNil(request.ListenerIds) {
		query["ListenerIds"] = request.ListenerIds
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerIds) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListListeners"),
		Version:     dara.String("2020-06-16"),
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
// Queries the configurations of instances.
//
// @param request - ListLoadBalancersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLoadBalancersResponse
func (client *Client) ListLoadBalancersWithContext(ctx context.Context, request *ListLoadBalancersRequest, runtime *dara.RuntimeOptions) (_result *ListLoadBalancersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressIpVersion) {
		query["AddressIpVersion"] = request.AddressIpVersion
	}

	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.DNSName) {
		query["DNSName"] = request.DNSName
	}

	if !dara.IsNil(request.Ipv6AddressType) {
		query["Ipv6AddressType"] = request.Ipv6AddressType
	}

	if !dara.IsNil(request.LoadBalancerBussinessStatus) {
		query["LoadBalancerBussinessStatus"] = request.LoadBalancerBussinessStatus
	}

	if !dara.IsNil(request.LoadBalancerIds) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !dara.IsNil(request.LoadBalancerNames) {
		query["LoadBalancerNames"] = request.LoadBalancerNames
	}

	if !dara.IsNil(request.LoadBalancerStatus) {
		query["LoadBalancerStatus"] = request.LoadBalancerStatus
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcIds) {
		query["VpcIds"] = request.VpcIds
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLoadBalancers"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLoadBalancersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the forwarding rules in a region.
//
// @param request - ListRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesResponse
func (client *Client) ListRulesWithContext(ctx context.Context, request *ListRulesRequest, runtime *dara.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.ListenerIds) {
		query["ListenerIds"] = request.ListenerIds
	}

	if !dara.IsNil(request.LoadBalancerIds) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RuleIds) {
		query["RuleIds"] = request.RuleIds
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRules"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom security policies in a region.
//
// @param request - ListSecurityPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecurityPoliciesResponse
func (client *Client) ListSecurityPoliciesWithContext(ctx context.Context, request *ListSecurityPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListSecurityPoliciesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityPolicyIds) {
		query["SecurityPolicyIds"] = request.SecurityPolicyIds
	}

	if !dara.IsNil(request.SecurityPolicyNames) {
		query["SecurityPolicyNames"] = request.SecurityPolicyNames
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecurityPolicies"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecurityPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the listeners that are associated with security policies.
//
// @param request - ListSecurityPolicyRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecurityPolicyRelationsResponse
func (client *Client) ListSecurityPolicyRelationsWithContext(ctx context.Context, request *ListSecurityPolicyRelationsRequest, runtime *dara.RuntimeOptions) (_result *ListSecurityPolicyRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityPolicyIds) {
		query["SecurityPolicyIds"] = request.SecurityPolicyIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecurityPolicyRelations"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecurityPolicyRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries servers in a server group.
//
// @param request - ListServerGroupServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServerGroupServersResponse
func (client *Client) ListServerGroupServersWithContext(ctx context.Context, request *ListServerGroupServersRequest, runtime *dara.RuntimeOptions) (_result *ListServerGroupServersResponse, _err error) {
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

	if !dara.IsNil(request.ServerGroupId) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	if !dara.IsNil(request.ServerIds) {
		query["ServerIds"] = request.ServerIds
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServerGroupServers"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServerGroupServersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries server groups.
//
// @param request - ListServerGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServerGroupsResponse
func (client *Client) ListServerGroupsWithContext(ctx context.Context, request *ListServerGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListServerGroupsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServerGroupIds) {
		query["ServerGroupIds"] = request.ServerGroupIds
	}

	if !dara.IsNil(request.ServerGroupNames) {
		query["ServerGroupNames"] = request.ServerGroupNames
	}

	if !dara.IsNil(request.ServerGroupType) {
		query["ServerGroupType"] = request.ServerGroupType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServerGroups"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServerGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tag keys.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithContext(ctx context.Context, request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

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
		Action:      dara.String("ListTagKeys"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of resources.
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
		Version:     dara.String("2020-06-16"),
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
// Queries tag values.
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithContext(ctx context.Context, request *ListTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
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
		Action:      dara.String("ListTagValues"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an Application Load Balancer (ALB) instance to a security group.
//
// Description:
//
//	  By default, security groups are unavailable. To use security groups, contact your account manager.
//
//		- Make sure that a security group is created. For more information about how to create security groups, see [CreateSecurityGroup](https://help.aliyun.com/document_detail/2679843.html).
//
//		- Each ALB instance can be added to at most four security groups.
//
//		- To query the security groups of an ALB instance, call the [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/2254835.html) operation.
//
//		- GetLoadBalancerAttribute is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAsynJobs](https://help.aliyun.com/document_detail/2254893.html) operation to query the status of the task.
//
//	    	- If the task is in the Succeeded state, the ALB instance is added to the security group.
//
//	    	- If the task is in the Processing state, the ALB instance is being added to the security group. In this case, you can query the task but cannot perform other operations.
//
// @param request - LoadBalancerJoinSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoadBalancerJoinSecurityGroupResponse
func (client *Client) LoadBalancerJoinSecurityGroupWithContext(ctx context.Context, request *LoadBalancerJoinSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *LoadBalancerJoinSecurityGroupResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.SecurityGroupIds) {
		query["SecurityGroupIds"] = request.SecurityGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LoadBalancerJoinSecurityGroup"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LoadBalancerJoinSecurityGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an Application Load Balancer (ALB) instance from a security group.
//
// Description:
//
//	LoadBalancerLeaveSecurityGroup is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAsynJobs](https://help.aliyun.com/document_detail/2254893.html) operation to query the status of the task.
//
//	  	- If the task is in the Succeeded state, the ALB instance is removed from the security group.
//
//	  	- If the task is in the Processing state, the ALB instance is being removed from the security group. In this case, you can query the task but cannot perform other operations.
//
// @param request - LoadBalancerLeaveSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoadBalancerLeaveSecurityGroupResponse
func (client *Client) LoadBalancerLeaveSecurityGroupWithContext(ctx context.Context, request *LoadBalancerLeaveSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *LoadBalancerLeaveSecurityGroupResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.SecurityGroupIds) {
		query["SecurityGroupIds"] = request.SecurityGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LoadBalancerLeaveSecurityGroup"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LoadBalancerLeaveSecurityGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改资源预留
//
// @param request - ModifyCapacityReservationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCapacityReservationResponse
func (client *Client) ModifyCapacityReservationWithContext(ctx context.Context, request *ModifyCapacityReservationRequest, runtime *dara.RuntimeOptions) (_result *ModifyCapacityReservationResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.MinimumLoadBalancerCapacity) {
		query["MinimumLoadBalancerCapacity"] = request.MinimumLoadBalancerCapacity
	}

	if !dara.IsNil(request.ResetCapacityReservation) {
		query["ResetCapacityReservation"] = request.ResetCapacityReservation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCapacityReservation"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCapacityReservationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a resource to another resource group.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithContext(ctx context.Context, request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
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
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes entries from an access control list (ACL).
//
// Description:
//
// *RemoveEntriesFromAcl*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAclEntries](https://help.aliyun.com/document_detail/213616.html) operation to query the status of the task.
//
//   - If an ACL is in the **Removing*	- state, the entries are being removed.
//
//   - If an ACL cannot be found, the entries are removed.
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
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Entries) {
		query["Entries"] = request.Entries
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveEntriesFromAcl"),
		Version:     dara.String("2020-06-16"),
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
// Removes backend servers from a server group.
//
// Description:
//
// *RemoveServersFromServerGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background.
//
// 1.  You can call [ListServerGroups](https://help.aliyun.com/document_detail/2254862.html) to query the status of a server group.
//
//   - If the server group is in the **Configuring*	- state, the server group is being modified.
//
//   - If the server group is in the **Available*	- state, the server group is running.
//
// 2.  You can call [ListServerGroupServers](https://help.aliyun.com/document_detail/2254863.html) to query the status of a backend server.
//
//   - If the backend server is in the **Removing*	- state, the backend server is being removed from the server group.
//
//   - If the backend server cannot be found, the backend server is no longer in the server group.
//
// @param request - RemoveServersFromServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveServersFromServerGroupResponse
func (client *Client) RemoveServersFromServerGroupWithContext(ctx context.Context, request *RemoveServersFromServerGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveServersFromServerGroupResponse, _err error) {
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

	if !dara.IsNil(request.ServerGroupId) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	if !dara.IsNil(request.Servers) {
		query["Servers"] = request.Servers
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveServersFromServerGroup"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveServersFromServerGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces backend servers in a server group.
//
// Description:
//
// *ReplaceServersInServerGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background.
//
// 1.  You can call the [ListServerGroups](https://help.aliyun.com/document_detail/213627.html) operation to query the status of a server group.
//
//   - If a server group is in the **Configuring*	- state, it indicates that the server group is being modified.
//
//   - If a server group is in the **Available*	- state, it indicates that the server group is running.
//
// 2.  You can call the [ListServerGroupServers](https://help.aliyun.com/document_detail/213628.html) operation to query the status of a backend server.
//
//   - If a backend server is in the **Replacing*	- state, it indicates that the server is being removed from the server group and a new server is added to the server group.
//
//   - If a backend server is in the \\*\\*Available\\*\\	- state, it indicates that the server is running.
//
// @param request - ReplaceServersInServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceServersInServerGroupResponse
func (client *Client) ReplaceServersInServerGroupWithContext(ctx context.Context, request *ReplaceServersInServerGroupRequest, runtime *dara.RuntimeOptions) (_result *ReplaceServersInServerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddedServers) {
		query["AddedServers"] = request.AddedServers
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RemovedServers) {
		query["RemovedServers"] = request.RemovedServers
	}

	if !dara.IsNil(request.ServerGroupId) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceServersInServerGroup"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceServersInServerGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a listener.
//
// Description:
//
// *StartListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [GetListenerAttribute](https://help.aliyun.com/document_detail/2254865.html) to query the status of the task.
//
//   - If a listener is in the **Configuring*	- state, the listener is being enabled.
//
//   - If a listener is in the **Running*	- state, the listener is enabled.
//
// @param request - StartListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartListenerResponse
func (client *Client) StartListenerWithContext(ctx context.Context, request *StartListenerRequest, runtime *dara.RuntimeOptions) (_result *StartListenerResponse, _err error) {
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

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartListener"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an elastic IP address (EIP) or a virtual IP address (VIP) of a zone from a DNS record.
//
// Description:
//
// This operation is supported by Application Load Balancer (ALB) instances that use static IP addresses. The zone cannot be removed if the ALB instance has only one available zone.
//
// @param request - StartShiftLoadBalancerZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartShiftLoadBalancerZonesResponse
func (client *Client) StartShiftLoadBalancerZonesWithContext(ctx context.Context, request *StartShiftLoadBalancerZonesRequest, runtime *dara.RuntimeOptions) (_result *StartShiftLoadBalancerZonesResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.ZoneMappings) {
		query["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartShiftLoadBalancerZones"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartShiftLoadBalancerZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a listener.
//
// Description:
//
// *StopListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetListenerAttribute](https://help.aliyun.com/document_detail/2254865.html) operation to query the status of the task:
//
//   - If a listener is in the **Configuring*	- state, the listener is being disabled.
//
//   - If a listener is in the **Stopped*	- state, the listener is disabled.
//
// @param request - StopListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopListenerResponse
func (client *Client) StopListenerWithContext(ctx context.Context, request *StopListenerRequest, runtime *dara.RuntimeOptions) (_result *StopListenerResponse, _err error) {
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

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopListener"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to resources.
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
		Version:     dara.String("2020-06-16"),
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
// Removes tags from resources.
//
// @param request - UnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithContext(ctx context.Context, request *UnTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnTagResources"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates AScript rules.
//
// Description:
//
// *UpdateAScripts*	- is an an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAScripts](https://help.aliyun.com/document_detail/472574.html) operation to query the status of an AScript rule.
//
//   - If the rule is in the **Configuring*	- state, the rule is being updated.
//
//   - If the rule is in the **Available*	- state, the rule is updated.
//
// @param request - UpdateAScriptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAScriptsResponse
func (client *Client) UpdateAScriptsWithContext(ctx context.Context, request *UpdateAScriptsRequest, runtime *dara.RuntimeOptions) (_result *UpdateAScriptsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AScripts) {
		query["AScripts"] = request.AScripts
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAScripts"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAScriptsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attributes of an access control list (ACL), such as the name.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAclAttribute"),
		Version:     dara.String("2020-06-16"),
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
// Modifies the attributes, such as the name and protocol, of a health check template.
//
// @param request - UpdateHealthCheckTemplateAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHealthCheckTemplateAttributeResponse
func (client *Client) UpdateHealthCheckTemplateAttributeWithContext(ctx context.Context, request *UpdateHealthCheckTemplateAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateHealthCheckTemplateAttributeResponse, _err error) {
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

	if !dara.IsNil(request.HealthCheckCodes) {
		query["HealthCheckCodes"] = request.HealthCheckCodes
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckHost) {
		query["HealthCheckHost"] = request.HealthCheckHost
	}

	if !dara.IsNil(request.HealthCheckHttpVersion) {
		query["HealthCheckHttpVersion"] = request.HealthCheckHttpVersion
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckMethod) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !dara.IsNil(request.HealthCheckPath) {
		query["HealthCheckPath"] = request.HealthCheckPath
	}

	if !dara.IsNil(request.HealthCheckProtocol) {
		query["HealthCheckProtocol"] = request.HealthCheckProtocol
	}

	if !dara.IsNil(request.HealthCheckTemplateId) {
		query["HealthCheckTemplateId"] = request.HealthCheckTemplateId
	}

	if !dara.IsNil(request.HealthCheckTemplateName) {
		query["HealthCheckTemplateName"] = request.HealthCheckTemplateName
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHealthCheckTemplateAttribute"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHealthCheckTemplateAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attributes of a listener, such as the name and the default action.
//
// Description:
//
// *UpdateListenerAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetListenerAttribute](https://help.aliyun.com/document_detail/2254865.html) operation to query the status of the task.
//
//   - If a listener is in the **Configuring*	- state, the configuration of the listener is being modified.
//
//   - If a listener is in the **Running*	- state, the configuration of the listener is modified.
//
// @param request - UpdateListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateListenerAttributeResponse
func (client *Client) UpdateListenerAttributeWithContext(ctx context.Context, request *UpdateListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaCertificates) {
		query["CaCertificates"] = request.CaCertificates
	}

	if !dara.IsNil(request.CaEnabled) {
		query["CaEnabled"] = request.CaEnabled
	}

	if !dara.IsNil(request.Certificates) {
		query["Certificates"] = request.Certificates
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DefaultActions) {
		query["DefaultActions"] = request.DefaultActions
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.GzipEnabled) {
		query["GzipEnabled"] = request.GzipEnabled
	}

	if !dara.IsNil(request.Http2Enabled) {
		query["Http2Enabled"] = request.Http2Enabled
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerDescription) {
		query["ListenerDescription"] = request.ListenerDescription
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.QuicConfig) {
		query["QuicConfig"] = request.QuicConfig
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
		Action:      dara.String("UpdateListenerAttribute"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateListenerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the log configuration of a listener, such as the access log configuration.
//
// Description:
//
// *UpdateListenerLogConfig*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [GetListenerAttribute](https://help.aliyun.com/document_detail/2254865.html) to query the status of the task:
//
//   - If a listener is in the **Configuring*	- state, the log configuration of the listener is being modified.
//
//   - If a listener is in the **Running*	- state, the log configuration of the listener is modified.
//
// > You can update the log configuration of a listener only after you enable the access log feature.
//
// @param request - UpdateListenerLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateListenerLogConfigResponse
func (client *Client) UpdateListenerLogConfigWithContext(ctx context.Context, request *UpdateListenerLogConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateListenerLogConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessLogRecordCustomizedHeadersEnabled) {
		query["AccessLogRecordCustomizedHeadersEnabled"] = request.AccessLogRecordCustomizedHeadersEnabled
	}

	if !dara.IsNil(request.AccessLogTracingConfig) {
		query["AccessLogTracingConfig"] = request.AccessLogTracingConfig
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateListenerLogConfig"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateListenerLogConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the network type of an Application Load Balancer (ALB) instance.
//
// Description:
//
// ## Prerequisites
//
//   - An ALB instance is created. For more information about how to create an ALB instance, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/214358.html).
//
//   - If you want to change the network type from internal-facing to Internet-facing, you must first create an elastic IP address (EIP). For more information, see [AllocateEipAddress](https://help.aliyun.com/document_detail/120192.html).
//
// ## Usage notes
//
// **UpdateLoadBalancerAddressTypeConfig*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/214362.html) operation to query the status of the task.
//
//   - If an ALB instance is in the **Configuring*	- state, the network type is being changed.
//
//   - If an ALB instance is in the **Active*	- state, the network type has been changed.
//
// @param request - UpdateLoadBalancerAddressTypeConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoadBalancerAddressTypeConfigResponse
func (client *Client) UpdateLoadBalancerAddressTypeConfigWithContext(ctx context.Context, request *UpdateLoadBalancerAddressTypeConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLoadBalancerAddressTypeConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.ZoneMappings) {
		query["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLoadBalancerAddressTypeConfig"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLoadBalancerAddressTypeConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an Application Load Balancer (ALB) instance, such as the name and the configuration read-only mode.
//
// Description:
//
// *UpdateLoadBalancerAttribute*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/214362.html) to query the status of the task.
//
//   - If the ALB instance is in the **Configuring*	- state, the ALB instance is being modified.
//
//   - If the ALB instance is in the **Active*	- state, the ALB instance is modified.
//
// @param request - UpdateLoadBalancerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoadBalancerAttributeResponse
func (client *Client) UpdateLoadBalancerAttributeWithContext(ctx context.Context, request *UpdateLoadBalancerAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateLoadBalancerAttributeResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LoadBalancerName) {
		query["LoadBalancerName"] = request.LoadBalancerName
	}

	if !dara.IsNil(request.ModificationProtectionConfig) {
		query["ModificationProtectionConfig"] = request.ModificationProtectionConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLoadBalancerAttribute"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLoadBalancerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the edition of an Application Load Balancer (ALB) instance.
//
// Description:
//
//	  You can only upgrade a basic ALB instance to a standard ALB instance or a WAF-enabled ALB instance. You cannot downgrade a standard ALB instance or a WAF-enabled ALB instance to a basic ALB instance. For more information, see [Upgrade an ALB instance](https://help.aliyun.com/document_detail/214654.html).
//
//		- **UpdateLoadBalancerEdition*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/214362.html) operation to query the status of an ALB instance.
//
//	    	- If the ALB instance is in the **Configuring*	- state, the edition of the ALB instance is being modified.
//
//	    	- If the ALB instance is in the **Active*	- state, the edition of the ALB instance is modified.
//
// @param request - UpdateLoadBalancerEditionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoadBalancerEditionResponse
func (client *Client) UpdateLoadBalancerEditionWithContext(ctx context.Context, request *UpdateLoadBalancerEditionRequest, runtime *dara.RuntimeOptions) (_result *UpdateLoadBalancerEditionResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerEdition) {
		query["LoadBalancerEdition"] = request.LoadBalancerEdition
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLoadBalancerEdition"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLoadBalancerEditionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the zones of an Application Load Balancer (ALB) instance.
//
// Description:
//
// *UpdateLoadBalancerZones*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [GetLoadBalancerAttribute](https://help.aliyun.com/document_detail/214362.html) to query the status of the task.
//
//   - If an ALB instance is in the **Configuring*	- state, the zones are being modified.
//
//   - If an ALB instance is in the **Active*	- state, the zones are modified.
//
// > You may be charged after you call UpdateLoadBalancerZones.
//
// @param request - UpdateLoadBalancerZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoadBalancerZonesResponse
func (client *Client) UpdateLoadBalancerZonesWithContext(ctx context.Context, request *UpdateLoadBalancerZonesRequest, runtime *dara.RuntimeOptions) (_result *UpdateLoadBalancerZonesResponse, _err error) {
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.ZoneMappings) {
		query["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLoadBalancerZones"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLoadBalancerZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a forwarding rule, such as the match condition, action, and name.
//
// Description:
//
//	  **UpdateRuleAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListRules](https://help.aliyun.com/document_detail/214379.html) operation to query the status of a forwarding rule:
//
//	    	- If a forwarding rule is in the **Configuring*	- state, the forwarding rule is being updated.
//
//	    	- If a forwarding rule is in the **Available*	- state, the forwarding rule is updated.
//
//		- You can set **RuleConditions*	- and **RuleActions*	- to add conditions and actions to a forwarding rule. Take note of the following limits on the number of conditions and the number of actions in each forwarding rule:
//
//	    	- Number of conditions: You can specify at most 5 for a basic Application Load Balancer (ALB) instance, at most 10 for a standard ALB instance, and at most 10 for a WAF-enabled ALB instance.
//
//	    	- Number of actions: You can specify at most 3 for a basic ALB instance, at most 5 for a standard ALB instance, and at most 5 for a WAF-enabled ALB instance.
//
// @param request - UpdateRuleAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleAttributeResponse
func (client *Client) UpdateRuleAttributeWithContext(ctx context.Context, request *UpdateRuleAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleAttributeResponse, _err error) {
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

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RuleActions) {
		query["RuleActions"] = request.RuleActions
	}

	if !dara.IsNil(request.RuleConditions) {
		query["RuleConditions"] = request.RuleConditions
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
		Action:      dara.String("UpdateRuleAttribute"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRuleAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of forwarding rules.
//
// Description:
//
// *UpdateRulesAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListRules](https://help.aliyun.com/document_detail/214379.html) operation to query the status of the task.
//
//   - If a forwarding rule is in the **Configuring*	- state, the forwarding rule is being updated.
//
//   - If a forwarding rule is in the **Available*	- state, the forwarding rule is updated.
//
//   - You can set **RuleConditions*	- and **RuleActions*	- to add conditions and actions to a forwarding rule. Take note of the following limits on the maximum number of conditions and the maximum number of actions in each forwarding rule:
//
//   - Limits on conditions: 5 for a basic Application Load Balancer (ALB) instance, 10 for a standard ALB instance, and 10 for a WAF-enabled ALB instance.
//
//   - Limits on actions: 3 for a basic ALB instance, 5 for a standard ALB instance, and 5 for a WAF-enabled ALB instance.
//
// @param request - UpdateRulesAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRulesAttributeResponse
func (client *Client) UpdateRulesAttributeWithContext(ctx context.Context, request *UpdateRulesAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateRulesAttributeResponse, _err error) {
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

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Rules) {
		bodyFlat["Rules"] = request.Rules
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRulesAttribute"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRulesAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attributes of a security policy, such as the TLS protocol version and the supported cipher suites.
//
// Description:
//
// ##
//
// **UpdateSecurityPolicyAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [ListSecurityPolicies](https://help.aliyun.com/document_detail/213609.html) to query the status of the task.
//
//   - If a security policy is in the **Configuring*	- state, the security policy is being updated.
//
//   - If a security policy is in the **Available*	- state, the security policy is updated.
//
// @param request - UpdateSecurityPolicyAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityPolicyAttributeResponse
func (client *Client) UpdateSecurityPolicyAttributeWithContext(ctx context.Context, request *UpdateSecurityPolicyAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityPolicyAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ciphers) {
		query["Ciphers"] = request.Ciphers
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.SecurityPolicyId) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !dara.IsNil(request.SecurityPolicyName) {
		query["SecurityPolicyName"] = request.SecurityPolicyName
	}

	if !dara.IsNil(request.TLSVersions) {
		query["TLSVersions"] = request.TLSVersions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecurityPolicyAttribute"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityPolicyAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a server group, such as health checks, session persistence, server group names, routing algorithms, and protocols.
//
// Description:
//
// ## Description
//
// **UpdateServerGroupAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListServerGroups](https://help.aliyun.com/document_detail/213627.html) operation to query the status of a server group:
//
//   - If a server group is in the **Configuring*	- state, the configuration of the server group is being modified.
//
//   - If a server group is in the **Available*	- state, the configuration of the server group is modified.
//
// @param request - UpdateServerGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServerGroupAttributeResponse
func (client *Client) UpdateServerGroupAttributeWithContext(ctx context.Context, request *UpdateServerGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateServerGroupAttributeResponse, _err error) {
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

	if !dara.IsNil(request.ConnectionDrainConfig) {
		query["ConnectionDrainConfig"] = request.ConnectionDrainConfig
	}

	if !dara.IsNil(request.CrossZoneEnabled) {
		query["CrossZoneEnabled"] = request.CrossZoneEnabled
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.HealthCheckConfig) {
		query["HealthCheckConfig"] = request.HealthCheckConfig
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.ServerGroupId) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	if !dara.IsNil(request.ServerGroupName) {
		query["ServerGroupName"] = request.ServerGroupName
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.SlowStartConfig) {
		query["SlowStartConfig"] = request.SlowStartConfig
	}

	if !dara.IsNil(request.StickySessionConfig) {
		query["StickySessionConfig"] = request.StickySessionConfig
	}

	if !dara.IsNil(request.UchConfig) {
		query["UchConfig"] = request.UchConfig
	}

	if !dara.IsNil(request.UpstreamKeepaliveEnabled) {
		query["UpstreamKeepaliveEnabled"] = request.UpstreamKeepaliveEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServerGroupAttribute"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServerGroupAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations, such as the backend server weight and description, of a server group.
//
// Description:
//
// *UpdateServerGroupServersAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background.
//
// 1.  You can call the [ListServerGroups](https://help.aliyun.com/document_detail/213627.html) operation to query the status of a server group.
//
//   - If a server group is in the **Configuring*	- state, it indicates that the server group is being modified.
//
//   - If a server group is in the **Available*	- state, it indicates that the server group is running.
//
// 2.  You can call the [ListServerGroupServers](https://help.aliyun.com/document_detail/213628.html) operation to query the status of a backend server.
//
//   - If a backend server is in the **Configuring*	- state, it indicates that the backend server is being modified.
//
//   - If a backend server is in the **Available*	- state, it indicates that the backend server is running.
//
// @param request - UpdateServerGroupServersAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServerGroupServersAttributeResponse
func (client *Client) UpdateServerGroupServersAttributeWithContext(ctx context.Context, request *UpdateServerGroupServersAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateServerGroupServersAttributeResponse, _err error) {
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

	if !dara.IsNil(request.ServerGroupId) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	if !dara.IsNil(request.Servers) {
		query["Servers"] = request.Servers
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServerGroupServersAttribute"),
		Version:     dara.String("2020-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServerGroupServersAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
