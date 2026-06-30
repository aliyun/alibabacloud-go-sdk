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
	client.EndpointMap = map[string]*string{
		"cn-hangzhou": dara.String("ga.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("ga"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds IP entries to an access control policy group and allows or restricts the forwarding of access requests to listeners for these IP entries by using Settings such as whitelists or blacklists, enabling precise control over client requests. You can call the AddEntriesToAcl operation to add IP entries to an access control policy group.
//
// Description:
//
// - **AddEntriesToAcl*	- is an asynchronous operation. After a request is sent, the system returns a request ID, but the IP entries are not yet added. The addition node continues in the background. You can call [GetAcl](https://help.aliyun.com/document_detail/258292.html) or [ListAcls](https://help.aliyun.com/document_detail/258291.html) to query the status of the access control policy group:
//
//   - If the access control policy group is in the **configuring*	- state, the IP entries are being added. In this state, you can only perform query operations and cannot perform other operations.
//
//   - If the access control policy group is in the **active*	- state, the IP entries are added.
//
// - You cannot concurrently call **AddEntriesToAcl*	- to add IP entries to an access control policy group within the same Global Accelerator (GA) instance.
//
// @param request - AddEntriesToAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEntriesToAclResponse
func (client *Client) AddEntriesToAclWithOptions(request *AddEntriesToAclRequest, runtime *dara.RuntimeOptions) (_result *AddEntriesToAclResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds IP entries to an access control policy group and allows or restricts the forwarding of access requests to listeners for these IP entries by using Settings such as whitelists or blacklists, enabling precise control over client requests. You can call the AddEntriesToAcl operation to add IP entries to an access control policy group.
//
// Description:
//
// - **AddEntriesToAcl*	- is an asynchronous operation. After a request is sent, the system returns a request ID, but the IP entries are not yet added. The addition node continues in the background. You can call [GetAcl](https://help.aliyun.com/document_detail/258292.html) or [ListAcls](https://help.aliyun.com/document_detail/258291.html) to query the status of the access control policy group:
//
//   - If the access control policy group is in the **configuring*	- state, the IP entries are being added. In this state, you can only perform query operations and cannot perform other operations.
//
//   - If the access control policy group is in the **active*	- state, the IP entries are added.
//
// - You cannot concurrently call **AddEntriesToAcl*	- to add IP entries to an access control policy group within the same Global Accelerator (GA) instance.
//
// @param request - AddEntriesToAclRequest
//
// @return AddEntriesToAclResponse
func (client *Client) AddEntriesToAcl(request *AddEntriesToAclRequest) (_result *AddEntriesToAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddEntriesToAclResponse{}
	_body, _err := client.AddEntriesToAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the AssociateAclsWithListener operation to associate access control policy groups with a listener.
//
// Description:
//
// - **AssociateAclsWithListener*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the access control policy group is not yet associated with the listener. The association node continues to run in the background. You can invoke [DescribeListener](https://help.aliyun.com/document_detail/153254.html) to query the listener status:
//
//   - If the listener is in the **updating*	- state, the access control policy group is being associated with the listener. In this state, you can only execute query operations and cannot execute other operations.
//
//   - If the listener is in the **active*	- state, the access control policy group is associated with the listener.
//
// - You cannot concurrently associate access control policy groups with listeners within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - AssociateAclsWithListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateAclsWithListenerResponse
func (client *Client) AssociateAclsWithListenerWithOptions(request *AssociateAclsWithListenerRequest, runtime *dara.RuntimeOptions) (_result *AssociateAclsWithListenerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the AssociateAclsWithListener operation to associate access control policy groups with a listener.
//
// Description:
//
// - **AssociateAclsWithListener*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the access control policy group is not yet associated with the listener. The association node continues to run in the background. You can invoke [DescribeListener](https://help.aliyun.com/document_detail/153254.html) to query the listener status:
//
//   - If the listener is in the **updating*	- state, the access control policy group is being associated with the listener. In this state, you can only execute query operations and cannot execute other operations.
//
//   - If the listener is in the **active*	- state, the access control policy group is associated with the listener.
//
// - You cannot concurrently associate access control policy groups with listeners within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - AssociateAclsWithListenerRequest
//
// @return AssociateAclsWithListenerResponse
func (client *Client) AssociateAclsWithListener(request *AssociateAclsWithListenerRequest) (_result *AssociateAclsWithListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateAclsWithListenerResponse{}
	_body, _err := client.AssociateAclsWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds multiple certificates to an HTTPS listener of a Global Accelerator instance. Combined with virtual endpoint groups and forwarding rules, this enables accelerated access to multiple HTTPS domain names. You can call the AssociateAdditionalCertificatesWithListener operation to bind additional certificates to an HTTPS listener.
//
// Description:
//
// - Only HTTPS protocol listeners support attaching extension certificates.
//
// - The **AssociateAdditionalCertificatesWithListener*	- operation is asynchronous. After you send a request, the system returns a request ID, but the attachment between the HTTPS listener and the extension certificates is not yet complete because the association node is still running in the background. You can invoke [DescribeListener](https://help.aliyun.com/document_detail/153254.html) to query the listener status:
//
//   - If the listener is in the **updating*	- state, the HTTPS listener and extension certificates are being attached. In this state, you can only execute query operations.
//
//   - If the listener is in the **active*	- state, the HTTPS listener and extension certificates are attached.
//
// - The **AssociateAdditionalCertificatesWithListener*	- operation does not support concurrent requests to attach extension certificates to HTTPS listeners within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - AssociateAdditionalCertificatesWithListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateAdditionalCertificatesWithListenerResponse
func (client *Client) AssociateAdditionalCertificatesWithListenerWithOptions(request *AssociateAdditionalCertificatesWithListenerRequest, runtime *dara.RuntimeOptions) (_result *AssociateAdditionalCertificatesWithListenerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds multiple certificates to an HTTPS listener of a Global Accelerator instance. Combined with virtual endpoint groups and forwarding rules, this enables accelerated access to multiple HTTPS domain names. You can call the AssociateAdditionalCertificatesWithListener operation to bind additional certificates to an HTTPS listener.
//
// Description:
//
// - Only HTTPS protocol listeners support attaching extension certificates.
//
// - The **AssociateAdditionalCertificatesWithListener*	- operation is asynchronous. After you send a request, the system returns a request ID, but the attachment between the HTTPS listener and the extension certificates is not yet complete because the association node is still running in the background. You can invoke [DescribeListener](https://help.aliyun.com/document_detail/153254.html) to query the listener status:
//
//   - If the listener is in the **updating*	- state, the HTTPS listener and extension certificates are being attached. In this state, you can only execute query operations.
//
//   - If the listener is in the **active*	- state, the HTTPS listener and extension certificates are attached.
//
// - The **AssociateAdditionalCertificatesWithListener*	- operation does not support concurrent requests to attach extension certificates to HTTPS listeners within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - AssociateAdditionalCertificatesWithListenerRequest
//
// @return AssociateAdditionalCertificatesWithListenerResponse
func (client *Client) AssociateAdditionalCertificatesWithListener(request *AssociateAdditionalCertificatesWithListenerRequest) (_result *AssociateAdditionalCertificatesWithListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.AssociateAdditionalCertificatesWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Alibaba Cloud Global Accelerator (GA) Integration with Cloud Products
//
// @param request - AssociateResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateResourcesResponse
func (client *Client) AssociateResourcesWithOptions(request *AssociateResourcesRequest, runtime *dara.RuntimeOptions) (_result *AssociateResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Alibaba Cloud Global Accelerator (GA) Integration with Cloud Products
//
// @param request - AssociateResourcesRequest
//
// @return AssociateResourcesResponse
func (client *Client) AssociateResources(request *AssociateResourcesRequest) (_result *AssociateResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateResourcesResponse{}
	_body, _err := client.AssociateResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If you want to protect your Global Accelerator (GA) applications from large-scale DDoS attacks and ensure service stability and availability, you can call the AttachDdosToAccelerator operation to associate an Anti-DDoS Pro or Anti-DDoS Premium instance with a GA instance.
//
// Description:
//
// Note the following when you call this operation:
//
// - AttachDdosToAccelerator is an asynchronous operation. After you send a request, the system returns a request ID, but the Anti-DDoS Pro or Anti-DDoS Premium instance is not yet associated with the Global Accelerator (GA) instance. The associate task continues to run in the background. You can call [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) or [ListAccelerators](https://help.aliyun.com/document_detail/153236.html) to query the status of the GA instance:
//
//   - If the GA instance is in the **configuring*	- state, the Anti-DDoS Pro or Anti-DDoS Premium instance is being associated with the GA instance. In this state, you can only perform query operations.
//
//   - If the GA instance is in the **active*	- state, the Anti-DDoS Pro or Anti-DDoS Premium instance is associated with the GA instance.
//
// - The AttachDdosToAccelerator operation does not support concurrent requests to associate Anti-DDoS Pro or Anti-DDoS Premium instances with the same GA instance.
//
// @param request - AttachDdosToAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachDdosToAcceleratorResponse
func (client *Client) AttachDdosToAcceleratorWithOptions(request *AttachDdosToAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *AttachDdosToAcceleratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If you want to protect your Global Accelerator (GA) applications from large-scale DDoS attacks and ensure service stability and availability, you can call the AttachDdosToAccelerator operation to associate an Anti-DDoS Pro or Anti-DDoS Premium instance with a GA instance.
//
// Description:
//
// Note the following when you call this operation:
//
// - AttachDdosToAccelerator is an asynchronous operation. After you send a request, the system returns a request ID, but the Anti-DDoS Pro or Anti-DDoS Premium instance is not yet associated with the Global Accelerator (GA) instance. The associate task continues to run in the background. You can call [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) or [ListAccelerators](https://help.aliyun.com/document_detail/153236.html) to query the status of the GA instance:
//
//   - If the GA instance is in the **configuring*	- state, the Anti-DDoS Pro or Anti-DDoS Premium instance is being associated with the GA instance. In this state, you can only perform query operations.
//
//   - If the GA instance is in the **active*	- state, the Anti-DDoS Pro or Anti-DDoS Premium instance is associated with the GA instance.
//
// - The AttachDdosToAccelerator operation does not support concurrent requests to associate Anti-DDoS Pro or Anti-DDoS Premium instances with the same GA instance.
//
// @param request - AttachDdosToAcceleratorRequest
//
// @return AttachDdosToAcceleratorResponse
func (client *Client) AttachDdosToAccelerator(request *AttachDdosToAcceleratorRequest) (_result *AttachDdosToAcceleratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachDdosToAcceleratorResponse{}
	_body, _err := client.AttachDdosToAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a Simple Log Service (SLS) Logstore with an endpoint group.
//
// Description:
//
// - **AttachLogStoreToEndpointGroup*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the association between the SLS Logstore and the endpoint group is not yet complete. The association node continues to run in the background. You can invoke [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) to query the status of the endpoint group:
//
//   - If the endpoint group is in the **updating*	- state, the SLS Logstore is being associated with the endpoint group. In this state, you can only execute query operations.
//
//   - If the endpoint group is in the **active*	- state, the SLS Logstore is associated with the endpoint group.
//
// - **AttachLogStoreToEndpointGroup*	- does not support concurrent association of SLS Logstores with endpoint groups within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - AttachLogStoreToEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachLogStoreToEndpointGroupResponse
func (client *Client) AttachLogStoreToEndpointGroupWithOptions(request *AttachLogStoreToEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *AttachLogStoreToEndpointGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a Simple Log Service (SLS) Logstore with an endpoint group.
//
// Description:
//
// - **AttachLogStoreToEndpointGroup*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the association between the SLS Logstore and the endpoint group is not yet complete. The association node continues to run in the background. You can invoke [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) to query the status of the endpoint group:
//
//   - If the endpoint group is in the **updating*	- state, the SLS Logstore is being associated with the endpoint group. In this state, you can only execute query operations.
//
//   - If the endpoint group is in the **active*	- state, the SLS Logstore is associated with the endpoint group.
//
// - **AttachLogStoreToEndpointGroup*	- does not support concurrent association of SLS Logstores with endpoint groups within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - AttachLogStoreToEndpointGroupRequest
//
// @return AttachLogStoreToEndpointGroupResponse
func (client *Client) AttachLogStoreToEndpointGroup(request *AttachLogStoreToEndpointGroupRequest) (_result *AttachLogStoreToEndpointGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachLogStoreToEndpointGroupResponse{}
	_body, _err := client.AttachLogStoreToEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the BandwidthPackageAddAccelerator operation to attach a bandwidth plan to an Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// - **BandwidthPackageAddAccelerator*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the bandwidth plan is not yet attached to the Alibaba Cloud Global Accelerator (GA) instance. The attachment node continues to run in the background. You can invoke [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) to query the status of the bandwidth plan:
//
//   - If the bandwidth plan is in the **binding*	- state, the bandwidth plan is being attached to the Alibaba Cloud Global Accelerator (GA) instance. In this state, you can only execute query operations.
//
//   - If the bandwidth plan is in the **active*	- state, the bandwidth plan is attached to the Alibaba Cloud Global Accelerator (GA) instance.
//
// - The **BandwidthPackageAddAccelerator*	- operation does not support concurrent requests to attach bandwidth plans to the same Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - BandwidthPackageAddAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BandwidthPackageAddAcceleratorResponse
func (client *Client) BandwidthPackageAddAcceleratorWithOptions(request *BandwidthPackageAddAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *BandwidthPackageAddAcceleratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the BandwidthPackageAddAccelerator operation to attach a bandwidth plan to an Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// - **BandwidthPackageAddAccelerator*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the bandwidth plan is not yet attached to the Alibaba Cloud Global Accelerator (GA) instance. The attachment node continues to run in the background. You can invoke [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) to query the status of the bandwidth plan:
//
//   - If the bandwidth plan is in the **binding*	- state, the bandwidth plan is being attached to the Alibaba Cloud Global Accelerator (GA) instance. In this state, you can only execute query operations.
//
//   - If the bandwidth plan is in the **active*	- state, the bandwidth plan is attached to the Alibaba Cloud Global Accelerator (GA) instance.
//
// - The **BandwidthPackageAddAccelerator*	- operation does not support concurrent requests to attach bandwidth plans to the same Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - BandwidthPackageAddAcceleratorRequest
//
// @return BandwidthPackageAddAcceleratorResponse
func (client *Client) BandwidthPackageAddAccelerator(request *BandwidthPackageAddAcceleratorRequest) (_result *BandwidthPackageAddAcceleratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BandwidthPackageAddAcceleratorResponse{}
	_body, _err := client.BandwidthPackageAddAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the BandwidthPackageRemoveAccelerator operation to disassociate a bandwidth plan from an Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// - Before you invoke the **BandwidthPackageRemoveAccelerator*	- operation, make sure that no acceleration regions or endpoint groups exist under the Alibaba Cloud Global Accelerator (GA) instance.
//
//   - To delete an acceleration region, see [DeleteIpSet](https://help.aliyun.com/document_detail/2253276.html) or [DeleteIpSets](https://help.aliyun.com/document_detail/2253278.html).
//
//   - To delete an endpoint group, see [DeleteEndpointGroup](https://help.aliyun.com/document_detail/2253305.html), [DeleteEndpointGroups](https://help.aliyun.com/document_detail/2253311.html), or [DeleteCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/2303183.html).
//
// - The **BandwidthPackageRemoveAccelerator*	- operation is asynchronous. After you send a request, the system returns a request ID, but the disassociation has not yet completed. The disassociation node continues to run in the background. You can invoke [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) to query the status of the bandwidth plan:
//
//   - If the bandwidth plan is in the **unbinding*	- state, the bandwidth plan is being disassociated from the GA instance. In this state, you can only execute query operations.
//
//   - If the bandwidth plan is in the **active*	- state, the bandwidth plan is disassociated from the GA instance.
//
// - The **BandwidthPackageRemoveAccelerator*	- operation does not support concurrent disassociation of bandwidth plans from the same GA instance.
//
// @param request - BandwidthPackageRemoveAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BandwidthPackageRemoveAcceleratorResponse
func (client *Client) BandwidthPackageRemoveAcceleratorWithOptions(request *BandwidthPackageRemoveAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *BandwidthPackageRemoveAcceleratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the BandwidthPackageRemoveAccelerator operation to disassociate a bandwidth plan from an Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// - Before you invoke the **BandwidthPackageRemoveAccelerator*	- operation, make sure that no acceleration regions or endpoint groups exist under the Alibaba Cloud Global Accelerator (GA) instance.
//
//   - To delete an acceleration region, see [DeleteIpSet](https://help.aliyun.com/document_detail/2253276.html) or [DeleteIpSets](https://help.aliyun.com/document_detail/2253278.html).
//
//   - To delete an endpoint group, see [DeleteEndpointGroup](https://help.aliyun.com/document_detail/2253305.html), [DeleteEndpointGroups](https://help.aliyun.com/document_detail/2253311.html), or [DeleteCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/2303183.html).
//
// - The **BandwidthPackageRemoveAccelerator*	- operation is asynchronous. After you send a request, the system returns a request ID, but the disassociation has not yet completed. The disassociation node continues to run in the background. You can invoke [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) to query the status of the bandwidth plan:
//
//   - If the bandwidth plan is in the **unbinding*	- state, the bandwidth plan is being disassociated from the GA instance. In this state, you can only execute query operations.
//
//   - If the bandwidth plan is in the **active*	- state, the bandwidth plan is disassociated from the GA instance.
//
// - The **BandwidthPackageRemoveAccelerator*	- operation does not support concurrent disassociation of bandwidth plans from the same GA instance.
//
// @param request - BandwidthPackageRemoveAcceleratorRequest
//
// @return BandwidthPackageRemoveAcceleratorResponse
func (client *Client) BandwidthPackageRemoveAccelerator(request *BandwidthPackageRemoveAcceleratorRequest) (_result *BandwidthPackageRemoveAcceleratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BandwidthPackageRemoveAcceleratorResponse{}
	_body, _err := client.BandwidthPackageRemoveAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the resource group to which a Global Accelerator resource belongs by calling the ChangeResourceGroup operation.
//
// Description:
//
// The **ChangeResourceGroup*	- operation does not support concurrent modifications to the resource group of Global Accelerator resources within the same Alibaba Cloud Global Accelerator (GA) instance.
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the resource group to which a Global Accelerator resource belongs by calling the ChangeResourceGroup operation.
//
// Description:
//
// The **ChangeResourceGroup*	- operation does not support concurrent modifications to the resource group of Global Accelerator resources within the same Alibaba Cloud Global Accelerator (GA) instance.
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
func (client *Client) ConfigEndpointProbeWithOptions(request *ConfigEndpointProbeRequest, runtime *dara.RuntimeOptions) (_result *ConfigEndpointProbeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ConfigEndpointProbeResponse
func (client *Client) ConfigEndpointProbe(request *ConfigEndpointProbeRequest) (_result *ConfigEndpointProbeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigEndpointProbeResponse{}
	_body, _err := client.ConfigEndpointProbeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Global Accelerator (GA) is a network acceleration service that provides coverage across the globe. It uses Alibaba Cloud\\"s high-quality Border Gateway Protocol (BGP) bandwidth and global transmission network to provide low-latency access from nearby locations. This reduces the impact of network issues, such as latency, jitter, and packet loss, on your service quality. GA provides a high-availability and high-performance network acceleration service for users worldwide. You can call the CreateAccelerator operation to create a Global Accelerator instance.
//
// Description:
//
// The **CreateAccelerator*	- operation is asynchronous. After you send a request, the system returns a Global Accelerator instance ID, but the instance is still being created in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the status of the Global Accelerator instance:
//
// - If a Global Accelerator instance is in the **init*	- state, the instance is being created. You can only perform query operations on the instance.
//
// - If a Global Accelerator instance is in the **active*	- state, the instance is created.
//
// @param request - CreateAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAcceleratorResponse
func (client *Client) CreateAcceleratorWithOptions(request *CreateAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *CreateAcceleratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Global Accelerator (GA) is a network acceleration service that provides coverage across the globe. It uses Alibaba Cloud\\"s high-quality Border Gateway Protocol (BGP) bandwidth and global transmission network to provide low-latency access from nearby locations. This reduces the impact of network issues, such as latency, jitter, and packet loss, on your service quality. GA provides a high-availability and high-performance network acceleration service for users worldwide. You can call the CreateAccelerator operation to create a Global Accelerator instance.
//
// Description:
//
// The **CreateAccelerator*	- operation is asynchronous. After you send a request, the system returns a Global Accelerator instance ID, but the instance is still being created in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the status of the Global Accelerator instance:
//
// - If a Global Accelerator instance is in the **init*	- state, the instance is being created. You can only perform query operations on the instance.
//
// - If a Global Accelerator instance is in the **active*	- state, the instance is created.
//
// @param request - CreateAcceleratorRequest
//
// @return CreateAcceleratorResponse
func (client *Client) CreateAccelerator(request *CreateAcceleratorRequest) (_result *CreateAcceleratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAcceleratorResponse{}
	_body, _err := client.CreateAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the CreateAcl operation to create an access control policy group.
//
// Description:
//
// *CreateAcl*	- is an asynchronous operation. After you invoke the operation, the system returns an access control policy group ID but the access control policy group is not yet created. The creation node continues to run in the background. You can invoke [GetAcl](https://help.aliyun.com/document_detail/258292.html) or [ListAcls](https://help.aliyun.com/document_detail/258291.html) to query the status of the access control policy group:
//
// - If the access control policy group is in the **init*	- state, the access control policy group is being created. In this state, you can only execute query operations and cannot execute other operations.
//
// - If the access control policy group is in the **active*	- state, the access control policy group is created.
//
// @param request - CreateAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAclResponse
func (client *Client) CreateAclWithOptions(request *CreateAclRequest, runtime *dara.RuntimeOptions) (_result *CreateAclResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the CreateAcl operation to create an access control policy group.
//
// Description:
//
// *CreateAcl*	- is an asynchronous operation. After you invoke the operation, the system returns an access control policy group ID but the access control policy group is not yet created. The creation node continues to run in the background. You can invoke [GetAcl](https://help.aliyun.com/document_detail/258292.html) or [ListAcls](https://help.aliyun.com/document_detail/258291.html) to query the status of the access control policy group:
//
// - If the access control policy group is in the **init*	- state, the access control policy group is being created. In this state, you can only execute query operations and cannot execute other operations.
//
// - If the access control policy group is in the **active*	- state, the access control policy group is created.
//
// @param request - CreateAclRequest
//
// @return CreateAclResponse
func (client *Client) CreateAcl(request *CreateAclRequest) (_result *CreateAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAclResponse{}
	_body, _err := client.CreateAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an origin probing task by calling the CreateApplicationMonitor operation.
//
// Description:
//
// You can call the **CreateApplicationMonitor*	- operation to create an origin probing task. This task monitors the end-to-end network quality from the probing point through Global Accelerator (GA) to the origin server in real time, helping you quickly locate network faults and perform targeted network optimization.
//
// Before you begin:
//
// - Only subscription Alibaba Cloud Global Accelerator (GA) instances of Medium Ⅰ or higher specifications support origin probing tasks.
//
// - Origin probing tasks cannot be created for UDP protocol listeners.
//
// - The service port of the monitoring address must be within the listener port range.
//
// - The **CreateApplicationMonitor*	- operation is asynchronous. After you invoke this operation, the system returns a node ID for the origin probing task, but the node is not yet created. The node creation continues in the background. You can invoke [DescribeApplicationMonitor](https://help.aliyun.com/document_detail/408463.html) or [ListApplicationMonitor](https://help.aliyun.com/document_detail/408462.html) to query the status of the origin probing task:
//
//   - If the origin probing task is in the **init*	- state, the task is being created. In this state, you can only perform query operations.
//
//   - If the origin probing task is in the **active*	- state, the task is created.
//
// - The **CreateApplicationMonitor*	- operation does not support concurrent creation of origin probing nodes within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - CreateApplicationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationMonitorResponse
func (client *Client) CreateApplicationMonitorWithOptions(request *CreateApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationMonitorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an origin probing task by calling the CreateApplicationMonitor operation.
//
// Description:
//
// You can call the **CreateApplicationMonitor*	- operation to create an origin probing task. This task monitors the end-to-end network quality from the probing point through Global Accelerator (GA) to the origin server in real time, helping you quickly locate network faults and perform targeted network optimization.
//
// Before you begin:
//
// - Only subscription Alibaba Cloud Global Accelerator (GA) instances of Medium Ⅰ or higher specifications support origin probing tasks.
//
// - Origin probing tasks cannot be created for UDP protocol listeners.
//
// - The service port of the monitoring address must be within the listener port range.
//
// - The **CreateApplicationMonitor*	- operation is asynchronous. After you invoke this operation, the system returns a node ID for the origin probing task, but the node is not yet created. The node creation continues in the background. You can invoke [DescribeApplicationMonitor](https://help.aliyun.com/document_detail/408463.html) or [ListApplicationMonitor](https://help.aliyun.com/document_detail/408462.html) to query the status of the origin probing task:
//
//   - If the origin probing task is in the **init*	- state, the task is being created. In this state, you can only perform query operations.
//
//   - If the origin probing task is in the **active*	- state, the task is created.
//
// - The **CreateApplicationMonitor*	- operation does not support concurrent creation of origin probing nodes within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - CreateApplicationMonitorRequest
//
// @return CreateApplicationMonitorResponse
func (client *Client) CreateApplicationMonitor(request *CreateApplicationMonitorRequest) (_result *CreateApplicationMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApplicationMonitorResponse{}
	_body, _err := client.CreateApplicationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a bandwidth plan.
//
// Description:
//
// You must create a basic bandwidth plan to use Global Accelerator (GA) for network acceleration. A basic bandwidth plan supports the following bandwidth types:
//
// - **Basic bandwidth**: The acceleration area and the area where the endpoint is deployed are in the Chinese mainland. The accelerated service is deployed on Alibaba Cloud.
//
// - **Enhanced bandwidth**: The acceleration area and the area where the endpoint is deployed are in the Chinese mainland. This bandwidth type can accelerate services on both Alibaba Cloud and public networks outside Alibaba Cloud.
//
// - **Advanced bandwidth**: The acceleration area and the area where the endpoint is deployed are outside the Chinese mainland. This bandwidth type can accelerate services on both Alibaba Cloud and public networks outside Alibaba Cloud. To accelerate access for users in the Chinese mainland, you can select China (Hong Kong) as the acceleration area.
//
// Note the following when you call this operation:
//
// - The **CreateBandwidthPackage*	- operation is asynchronous. After you send a request, the system returns a bandwidth plan ID, but the bandwidth plan is not created immediately. The system creates the bandwidth plan in the background. You can call the [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) operation to query the status of the bandwidth plan:
//
//   - If a bandwidth plan is in the **init*	- state, the bandwidth plan is being created. In this state, you can only query the bandwidth plan and cannot perform other operations.
//
//   - If a bandwidth plan is in the **active*	- state, the bandwidth plan is created.
//
// - The **CreateBandwidthPackage*	- operation does not support concurrent requests to create bandwidth plans for the same Global Accelerator instance.
//
// @param request - CreateBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBandwidthPackageResponse
func (client *Client) CreateBandwidthPackageWithOptions(request *CreateBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *CreateBandwidthPackageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You must create a basic bandwidth plan to use Global Accelerator (GA) for network acceleration. A basic bandwidth plan supports the following bandwidth types:
//
// - **Basic bandwidth**: The acceleration area and the area where the endpoint is deployed are in the Chinese mainland. The accelerated service is deployed on Alibaba Cloud.
//
// - **Enhanced bandwidth**: The acceleration area and the area where the endpoint is deployed are in the Chinese mainland. This bandwidth type can accelerate services on both Alibaba Cloud and public networks outside Alibaba Cloud.
//
// - **Advanced bandwidth**: The acceleration area and the area where the endpoint is deployed are outside the Chinese mainland. This bandwidth type can accelerate services on both Alibaba Cloud and public networks outside Alibaba Cloud. To accelerate access for users in the Chinese mainland, you can select China (Hong Kong) as the acceleration area.
//
// Note the following when you call this operation:
//
// - The **CreateBandwidthPackage*	- operation is asynchronous. After you send a request, the system returns a bandwidth plan ID, but the bandwidth plan is not created immediately. The system creates the bandwidth plan in the background. You can call the [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) operation to query the status of the bandwidth plan:
//
//   - If a bandwidth plan is in the **init*	- state, the bandwidth plan is being created. In this state, you can only query the bandwidth plan and cannot perform other operations.
//
//   - If a bandwidth plan is in the **active*	- state, the bandwidth plan is created.
//
// - The **CreateBandwidthPackage*	- operation does not support concurrent requests to create bandwidth plans for the same Global Accelerator instance.
//
// @param request - CreateBandwidthPackageRequest
//
// @return CreateBandwidthPackageResponse
func (client *Client) CreateBandwidthPackage(request *CreateBandwidthPackageRequest) (_result *CreateBandwidthPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBandwidthPackageResponse{}
	_body, _err := client.CreateBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateBasicAccelerateIpWithOptions(request *CreateBasicAccelerateIpRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicAccelerateIpResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateBasicAccelerateIpResponse
func (client *Client) CreateBasicAccelerateIp(request *CreateBasicAccelerateIpRequest) (_result *CreateBasicAccelerateIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBasicAccelerateIpResponse{}
	_body, _err := client.CreateBasicAccelerateIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateBasicAccelerateIpEndpointRelationWithOptions(request *CreateBasicAccelerateIpEndpointRelationRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicAccelerateIpEndpointRelationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateBasicAccelerateIpEndpointRelationResponse
func (client *Client) CreateBasicAccelerateIpEndpointRelation(request *CreateBasicAccelerateIpEndpointRelationRequest) (_result *CreateBasicAccelerateIpEndpointRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBasicAccelerateIpEndpointRelationResponse{}
	_body, _err := client.CreateBasicAccelerateIpEndpointRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the CreateBasicAccelerateIpEndpointRelations operation to batch attach accelerated IP addresses to endpoints for a basic Global Accelerator instance.
//
// Description:
//
// - **CreateBasicAccelerateIpEndpointRelations*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the task of batch attaching accelerated IP addresses to endpoints is still in progress. You can call [GetBasicAccelerateIp](https://help.aliyun.com/document_detail/466794.html) or [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) to query the status of accelerated IP addresses and endpoints respectively to confirm whether the attachments are created:
//
//   - If an accelerated IP address or endpoint is in the **binding*	- state, the attachment is being created. In this state, you can only perform query operations.
//
//   - If all accelerated IP addresses and endpoints are in the **bound*	- state, and the attachment status returned by [ListBasicAccelerateIpEndpointRelations](https://help.aliyun.com/document_detail/466803.html) is **active**, the batch task of attaching accelerated IP addresses to endpoints is complete.
//
// - **CreateBasicAccelerateIpEndpointRelations*	- does not support concurrent batch attaching of accelerated IP addresses to endpoints within the same basic Global Accelerator instance.
//
// @param request - CreateBasicAccelerateIpEndpointRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicAccelerateIpEndpointRelationsResponse
func (client *Client) CreateBasicAccelerateIpEndpointRelationsWithOptions(request *CreateBasicAccelerateIpEndpointRelationsRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicAccelerateIpEndpointRelationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the CreateBasicAccelerateIpEndpointRelations operation to batch attach accelerated IP addresses to endpoints for a basic Global Accelerator instance.
//
// Description:
//
// - **CreateBasicAccelerateIpEndpointRelations*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the task of batch attaching accelerated IP addresses to endpoints is still in progress. You can call [GetBasicAccelerateIp](https://help.aliyun.com/document_detail/466794.html) or [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) to query the status of accelerated IP addresses and endpoints respectively to confirm whether the attachments are created:
//
//   - If an accelerated IP address or endpoint is in the **binding*	- state, the attachment is being created. In this state, you can only perform query operations.
//
//   - If all accelerated IP addresses and endpoints are in the **bound*	- state, and the attachment status returned by [ListBasicAccelerateIpEndpointRelations](https://help.aliyun.com/document_detail/466803.html) is **active**, the batch task of attaching accelerated IP addresses to endpoints is complete.
//
// - **CreateBasicAccelerateIpEndpointRelations*	- does not support concurrent batch attaching of accelerated IP addresses to endpoints within the same basic Global Accelerator instance.
//
// @param request - CreateBasicAccelerateIpEndpointRelationsRequest
//
// @return CreateBasicAccelerateIpEndpointRelationsResponse
func (client *Client) CreateBasicAccelerateIpEndpointRelations(request *CreateBasicAccelerateIpEndpointRelationsRequest) (_result *CreateBasicAccelerateIpEndpointRelationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBasicAccelerateIpEndpointRelationsResponse{}
	_body, _err := client.CreateBasicAccelerateIpEndpointRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Basic Alibaba Cloud Global Accelerator (GA) instances leverage Alibaba Cloud\\"s premium global the Internet bandwidth and high-quality transmission network to provide users with point-to-point acceleration. Basic Alibaba Cloud Global Accelerator (GA) instances are primarily used for Layer 3 (IP protocol) network acceleration. You can invoke the CreateBasicAccelerator operation to create a basic Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// *CreateBasicAccelerator*	- is an asynchronous operation. After you invoke this operation, the system returns a basic Alibaba Cloud Global Accelerator (GA) instance ID, but the instance is not yet created. The creation node continues to execute in the background. You can invoke [GetBasicAccelerator](https://help.aliyun.com/document_detail/353188.html) or [ListBasicAccelerators](https://help.aliyun.com/document_detail/353189.html) to query the status of the basic GA instance:
//
// - If the basic GA instance is in the **init*	- state, the instance is being created. In this state, you can only perform query operations.
//
// - If the basic GA instance is in the **active*	- state, the instance is created.
//
// @param request - CreateBasicAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicAcceleratorResponse
func (client *Client) CreateBasicAcceleratorWithOptions(request *CreateBasicAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicAcceleratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Basic Alibaba Cloud Global Accelerator (GA) instances leverage Alibaba Cloud\\"s premium global the Internet bandwidth and high-quality transmission network to provide users with point-to-point acceleration. Basic Alibaba Cloud Global Accelerator (GA) instances are primarily used for Layer 3 (IP protocol) network acceleration. You can invoke the CreateBasicAccelerator operation to create a basic Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// *CreateBasicAccelerator*	- is an asynchronous operation. After you invoke this operation, the system returns a basic Alibaba Cloud Global Accelerator (GA) instance ID, but the instance is not yet created. The creation node continues to execute in the background. You can invoke [GetBasicAccelerator](https://help.aliyun.com/document_detail/353188.html) or [ListBasicAccelerators](https://help.aliyun.com/document_detail/353189.html) to query the status of the basic GA instance:
//
// - If the basic GA instance is in the **init*	- state, the instance is being created. In this state, you can only perform query operations.
//
// - If the basic GA instance is in the **active*	- state, the instance is created.
//
// @param request - CreateBasicAcceleratorRequest
//
// @return CreateBasicAcceleratorResponse
func (client *Client) CreateBasicAccelerator(request *CreateBasicAcceleratorRequest) (_result *CreateBasicAcceleratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBasicAcceleratorResponse{}
	_body, _err := client.CreateBasicAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the CreateBasicEndpoint operation to create an endpoint for a basic Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// - **CreateBasicEndpoint*	- is an asynchronous operation. After you invoke this operation, the system returns an endpoint ID for the basic Alibaba Cloud Global Accelerator (GA) instance, but the endpoint is not yet created. The creation task continues to execute in the background. You can invoke [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) to query the endpoint status:
//
//   - When the endpoint is in the **init*	- state, the endpoint is being created. In this state, you can only execute query operations.
//
//   - When the endpoint is in the **active*	- state, the endpoint is created.
//
// - **CreateBasicEndpoint*	- does not support concurrent endpoint creation within the same basic Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - CreateBasicEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicEndpointResponse
func (client *Client) CreateBasicEndpointWithOptions(request *CreateBasicEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the CreateBasicEndpoint operation to create an endpoint for a basic Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// - **CreateBasicEndpoint*	- is an asynchronous operation. After you invoke this operation, the system returns an endpoint ID for the basic Alibaba Cloud Global Accelerator (GA) instance, but the endpoint is not yet created. The creation task continues to execute in the background. You can invoke [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) to query the endpoint status:
//
//   - When the endpoint is in the **init*	- state, the endpoint is being created. In this state, you can only execute query operations.
//
//   - When the endpoint is in the **active*	- state, the endpoint is created.
//
// - **CreateBasicEndpoint*	- does not support concurrent endpoint creation within the same basic Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - CreateBasicEndpointRequest
//
// @return CreateBasicEndpointResponse
func (client *Client) CreateBasicEndpoint(request *CreateBasicEndpointRequest) (_result *CreateBasicEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBasicEndpointResponse{}
	_body, _err := client.CreateBasicEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the CreateBasicEndpointGroup operation to create an endpoint group for a basic Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// - **CreateBasicEndpointGroup*	- is an asynchronous operation. After you invoke this operation, the system returns an endpoint group ID before the endpoint group is created. The endpoint group is being created in the background. You can invoke [GetBasicEndpointGroup](https://help.aliyun.com/document_detail/362984.html) to query the status of the endpoint group:
//
//   - If the endpoint group is in the **init*	- state, the endpoint group is being created. In this state, you can only perform query operations.
//
//   - If the endpoint group is in the **active*	- state, the endpoint group is created.
//
// - **CreateBasicEndpointGroup*	- does not support concurrent requests to create an endpoint group for the same basic Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - CreateBasicEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicEndpointGroupResponse
func (client *Client) CreateBasicEndpointGroupWithOptions(request *CreateBasicEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicEndpointGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the CreateBasicEndpointGroup operation to create an endpoint group for a basic Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// - **CreateBasicEndpointGroup*	- is an asynchronous operation. After you invoke this operation, the system returns an endpoint group ID before the endpoint group is created. The endpoint group is being created in the background. You can invoke [GetBasicEndpointGroup](https://help.aliyun.com/document_detail/362984.html) to query the status of the endpoint group:
//
//   - If the endpoint group is in the **init*	- state, the endpoint group is being created. In this state, you can only perform query operations.
//
//   - If the endpoint group is in the **active*	- state, the endpoint group is created.
//
// - **CreateBasicEndpointGroup*	- does not support concurrent requests to create an endpoint group for the same basic Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - CreateBasicEndpointGroupRequest
//
// @return CreateBasicEndpointGroupResponse
func (client *Client) CreateBasicEndpointGroup(request *CreateBasicEndpointGroupRequest) (_result *CreateBasicEndpointGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBasicEndpointGroupResponse{}
	_body, _err := client.CreateBasicEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateBasicEndpointsWithOptions(request *CreateBasicEndpointsRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicEndpointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateBasicEndpointsResponse
func (client *Client) CreateBasicEndpoints(request *CreateBasicEndpointsRequest) (_result *CreateBasicEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBasicEndpointsResponse{}
	_body, _err := client.CreateBasicEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the CreateBasicIpSet operation to create an acceleration region for a basic Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// When you invoke this operation, take note of the following items:
//
// - A basic Alibaba Cloud Global Accelerator (GA) instance supports only one acceleration region and supports only the IPv4 protocol.
//
// - **CreateBasicIpSet*	- is an asynchronous operation. After a request is sent, the system returns an acceleration region instance ID but the acceleration region is not yet created. The creation node continues to run in the background. You can invoke [GetBasicIpSet](https://help.aliyun.com/document_detail/362987.html) to query the status of the acceleration region:
//
//   - If the acceleration region is in the **init*	- state, the acceleration region is being created. In this state, you can only execute query operations.
//
//   - If the acceleration region is in the **active*	- state, the acceleration region is created.
//
// - The **CreateBasicIpSet*	- operation does not support concurrent creation of acceleration regions within the same basic Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - CreateBasicIpSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicIpSetResponse
func (client *Client) CreateBasicIpSetWithOptions(request *CreateBasicIpSetRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicIpSetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the CreateBasicIpSet operation to create an acceleration region for a basic Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// When you invoke this operation, take note of the following items:
//
// - A basic Alibaba Cloud Global Accelerator (GA) instance supports only one acceleration region and supports only the IPv4 protocol.
//
// - **CreateBasicIpSet*	- is an asynchronous operation. After a request is sent, the system returns an acceleration region instance ID but the acceleration region is not yet created. The creation node continues to run in the background. You can invoke [GetBasicIpSet](https://help.aliyun.com/document_detail/362987.html) to query the status of the acceleration region:
//
//   - If the acceleration region is in the **init*	- state, the acceleration region is being created. In this state, you can only execute query operations.
//
//   - If the acceleration region is in the **active*	- state, the acceleration region is created.
//
// - The **CreateBasicIpSet*	- operation does not support concurrent creation of acceleration regions within the same basic Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - CreateBasicIpSetRequest
//
// @return CreateBasicIpSetResponse
func (client *Client) CreateBasicIpSet(request *CreateBasicIpSetRequest) (_result *CreateBasicIpSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBasicIpSetResponse{}
	_body, _err := client.CreateBasicIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the CreateCustomRoutingEndpointGroupDestinations operation to create mapping configurations for an endpoint group that is associated with a custom route listener.
//
// Description:
//
// An Alibaba Cloud Global Accelerator (GA) instance can generate a port mapping table based on the configured listener port range, the mapping configurations (protocols and port ranges) of the destination endpoint group, and the IP address information of the endpoints (vSwitches). This enables deterministic routing of traffic to specific IP addresses and ports within the vSwitches.
//
// This operation creates mapping configurations for an endpoint group that is associated with a custom route listener. When you invoke this operation, take note of the following items:
//
// - **CreateCustomRoutingEndpointGroupDestinations*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the mapping configurations for the endpoint group are not yet created. The creation node continues to run in the background. You can invoke the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) operation to query the status of the endpoint group and confirm whether the mapping configurations are created:
//
//   - If the endpoint group is in the **updating*	- state, the mapping configurations are being created. In this state, you can only execute query operations.
//
//   - If the endpoint group is in the **active*	- state, the mapping configurations are created.
//
// - **CreateCustomRoutingEndpointGroupDestinations*	- does not support concurrent creation of mapping configurations for endpoint groups associated with custom route listeners within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// ### Before you begin
//
// Before you create mapping configurations for an endpoint group associated with a custom route listener, make sure that you have completed the following operations:
//
// - A standard Alibaba Cloud Global Accelerator (GA) instance is created. For more information, see [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html).
//
// - A bandwidth plan is attached to the standard Global Accelerator instance. For more information, see [BandwidthPackageAddAccelerator](https://help.aliyun.com/document_detail/153239.html).
//
// - You have completed the deployment of the required applications as backend services to accept forwarded requests from Global Accelerator. Custom route listeners support only vSwitches as backend service types.
//
// - You have obtained the permissions to use custom route listeners and created a custom route listener. The custom route listener type is in invitational preview. To use this feature, contact your account manager. To create a custom route listener, see [CreateListener](https://help.aliyun.com/document_detail/153253.html).
//
// - You have created an endpoint group for the custom route listener. For more information, see [CreateCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/449363.html).
//
// @param request - CreateCustomRoutingEndpointGroupDestinationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomRoutingEndpointGroupDestinationsResponse
func (client *Client) CreateCustomRoutingEndpointGroupDestinationsWithOptions(request *CreateCustomRoutingEndpointGroupDestinationsRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomRoutingEndpointGroupDestinationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the CreateCustomRoutingEndpointGroupDestinations operation to create mapping configurations for an endpoint group that is associated with a custom route listener.
//
// Description:
//
// An Alibaba Cloud Global Accelerator (GA) instance can generate a port mapping table based on the configured listener port range, the mapping configurations (protocols and port ranges) of the destination endpoint group, and the IP address information of the endpoints (vSwitches). This enables deterministic routing of traffic to specific IP addresses and ports within the vSwitches.
//
// This operation creates mapping configurations for an endpoint group that is associated with a custom route listener. When you invoke this operation, take note of the following items:
//
// - **CreateCustomRoutingEndpointGroupDestinations*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the mapping configurations for the endpoint group are not yet created. The creation node continues to run in the background. You can invoke the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) operation to query the status of the endpoint group and confirm whether the mapping configurations are created:
//
//   - If the endpoint group is in the **updating*	- state, the mapping configurations are being created. In this state, you can only execute query operations.
//
//   - If the endpoint group is in the **active*	- state, the mapping configurations are created.
//
// - **CreateCustomRoutingEndpointGroupDestinations*	- does not support concurrent creation of mapping configurations for endpoint groups associated with custom route listeners within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// ### Before you begin
//
// Before you create mapping configurations for an endpoint group associated with a custom route listener, make sure that you have completed the following operations:
//
// - A standard Alibaba Cloud Global Accelerator (GA) instance is created. For more information, see [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html).
//
// - A bandwidth plan is attached to the standard Global Accelerator instance. For more information, see [BandwidthPackageAddAccelerator](https://help.aliyun.com/document_detail/153239.html).
//
// - You have completed the deployment of the required applications as backend services to accept forwarded requests from Global Accelerator. Custom route listeners support only vSwitches as backend service types.
//
// - You have obtained the permissions to use custom route listeners and created a custom route listener. The custom route listener type is in invitational preview. To use this feature, contact your account manager. To create a custom route listener, see [CreateListener](https://help.aliyun.com/document_detail/153253.html).
//
// - You have created an endpoint group for the custom route listener. For more information, see [CreateCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/449363.html).
//
// @param request - CreateCustomRoutingEndpointGroupDestinationsRequest
//
// @return CreateCustomRoutingEndpointGroupDestinationsResponse
func (client *Client) CreateCustomRoutingEndpointGroupDestinations(request *CreateCustomRoutingEndpointGroupDestinationsRequest) (_result *CreateCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.CreateCustomRoutingEndpointGroupDestinationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the CreateCustomRoutingEndpointGroups operation to create endpoint groups for a custom routing type listener in batches.
//
// Description:
//
// Global Accelerator allocates traffic to endpoints within endpoint groups based on the forwarding method defined by the listener routing type.
//
// - After you configure an intelligent routing listener, the Alibaba Cloud Global Accelerator (GA) instance automatically selects the nearest healthy endpoint group for traffic forwarding based on latency factors (primarily depending on geographic location and network link conditions), and ultimately delivers client network access requests to healthy endpoints.
//
// - After you configure a custom routing type listener, the Alibaba Cloud Global Accelerator (GA) instance generates a port mapping table based on the configured listener port range, destination endpoint group protocol and port range, and IP address information of the endpoints (vSwitches), to deterministically route traffic to specific IP addresses and ports within vSwitches.
//
// This operation creates endpoint groups for a custom routing type listener. To create endpoint groups for an intelligent routing listener, invoke [CreateEndpointGroup](https://help.aliyun.com/document_detail/153259.html).
//
// When you invoke this operation, take note of the following items:
//
// - **CreateCustomRoutingEndpointGroups*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the endpoint groups for the custom routing type listener are not yet created. The creation task continues to execute in the background. You can invoke [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) or [ListCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/449374.html) to query the status of the endpoint groups:
//
//   - If an endpoint group is in the **init*	- state, the endpoint groups are being created in batches. In this state, you can only execute query operations.
//
//   - When all endpoint groups are in the **active*	- state, the batch creation is complete.
//
// - **CreateCustomRoutingEndpointGroups*	- does not support concurrent creation of endpoint groups for custom routing type listeners within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// ### Before you begin
//
// Before you create endpoint groups for a custom routing type listener, make sure that you have completed the following operations:
//
// - A standard Global Accelerator instance is created. For more information, see [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html).
//
// - A bandwidth plan is attached to the standard Alibaba Cloud Global Accelerator (GA) instance. For more information, see [BandwidthPackageAddAccelerator](https://help.aliyun.com/document_detail/153239.html).
//
// - You have deployed the relevant applications as backend services for Global Accelerator to accept forwarded requests. Custom routing type listeners support only vSwitches as the backend service type.
//
// - You have applied for permissions to use custom routing type listeners and created a custom routing type listener. The custom routing type for listeners is in invitational preview. To use this feature, contact your account manager. To create a custom routing type listener, see [CreateListener](https://help.aliyun.com/document_detail/153253.html).
//
// @param request - CreateCustomRoutingEndpointGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomRoutingEndpointGroupsResponse
func (client *Client) CreateCustomRoutingEndpointGroupsWithOptions(request *CreateCustomRoutingEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomRoutingEndpointGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the CreateCustomRoutingEndpointGroups operation to create endpoint groups for a custom routing type listener in batches.
//
// Description:
//
// Global Accelerator allocates traffic to endpoints within endpoint groups based on the forwarding method defined by the listener routing type.
//
// - After you configure an intelligent routing listener, the Alibaba Cloud Global Accelerator (GA) instance automatically selects the nearest healthy endpoint group for traffic forwarding based on latency factors (primarily depending on geographic location and network link conditions), and ultimately delivers client network access requests to healthy endpoints.
//
// - After you configure a custom routing type listener, the Alibaba Cloud Global Accelerator (GA) instance generates a port mapping table based on the configured listener port range, destination endpoint group protocol and port range, and IP address information of the endpoints (vSwitches), to deterministically route traffic to specific IP addresses and ports within vSwitches.
//
// This operation creates endpoint groups for a custom routing type listener. To create endpoint groups for an intelligent routing listener, invoke [CreateEndpointGroup](https://help.aliyun.com/document_detail/153259.html).
//
// When you invoke this operation, take note of the following items:
//
// - **CreateCustomRoutingEndpointGroups*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the endpoint groups for the custom routing type listener are not yet created. The creation task continues to execute in the background. You can invoke [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) or [ListCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/449374.html) to query the status of the endpoint groups:
//
//   - If an endpoint group is in the **init*	- state, the endpoint groups are being created in batches. In this state, you can only execute query operations.
//
//   - When all endpoint groups are in the **active*	- state, the batch creation is complete.
//
// - **CreateCustomRoutingEndpointGroups*	- does not support concurrent creation of endpoint groups for custom routing type listeners within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// ### Before you begin
//
// Before you create endpoint groups for a custom routing type listener, make sure that you have completed the following operations:
//
// - A standard Global Accelerator instance is created. For more information, see [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html).
//
// - A bandwidth plan is attached to the standard Alibaba Cloud Global Accelerator (GA) instance. For more information, see [BandwidthPackageAddAccelerator](https://help.aliyun.com/document_detail/153239.html).
//
// - You have deployed the relevant applications as backend services for Global Accelerator to accept forwarded requests. Custom routing type listeners support only vSwitches as the backend service type.
//
// - You have applied for permissions to use custom routing type listeners and created a custom routing type listener. The custom routing type for listeners is in invitational preview. To use this feature, contact your account manager. To create a custom routing type listener, see [CreateListener](https://help.aliyun.com/document_detail/153253.html).
//
// @param request - CreateCustomRoutingEndpointGroupsRequest
//
// @return CreateCustomRoutingEndpointGroupsResponse
func (client *Client) CreateCustomRoutingEndpointGroups(request *CreateCustomRoutingEndpointGroupsRequest) (_result *CreateCustomRoutingEndpointGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomRoutingEndpointGroupsResponse{}
	_body, _err := client.CreateCustomRoutingEndpointGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the CreateCustomRoutingEndpointTrafficPolicies operation to create an endpoint traffic policy (custom route type listener).
//
// Description:
//
// This operation takes effect only when the traffic policy of the backend service for the endpoint is set to allow traffic to specified destinations that can accept access traffic. You can invoke [DescribeCustomRoutingEndpoint](https://help.aliyun.com/document_detail/449386.html) to query the traffic policy of the backend service for a specified endpoint. This operation takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom*	- (specifying destinations that can accept access traffic).
//
// Before you invoke this operation, take note of the following items:
//
// - **CreateCustomRoutingEndpointTrafficPolicies*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the endpoint traffic policies for the custom route type listener are not yet created. The creation task continues to run in the background. You can invoke [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) to query the status of the endpoint group to confirm whether the traffic policies are created.
//
//   - If the endpoint group is in the **updating*	- state, the traffic policies are being created. In this state, you can only execute query operations.
//
//   - If the endpoint group is in the **active*	- state, the traffic policies are created.
//
// - The **CreateCustomRoutingEndpointTrafficPolicies*	- operation does not support concurrent creation of endpoint traffic policies within the same Global Accelerator instance.
//
// ### Before you begin
//
// Before you create an endpoint traffic policy, make sure that you have completed the following operations:
//
// - A standard Global Accelerator instance is created. For more information, see [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html).
//
// - If the billing method of the standard Global Accelerator instance is **pay-by-bandwidth**, a basic bandwidth plan is attached to the standard Global Accelerator instance. For more information, see [BandwidthPackageAddAccelerator](https://help.aliyun.com/document_detail/153239.html).
//
// - You have deployed the required applications as backend services to accept forwarded requests from Global Accelerator. The backend service type for custom route type listeners supports only vSwitches.
//
// - You have obtained the permissions to use custom route type listeners and created a custom route type listener. The custom route type for listeners is in invitational preview. To use this feature, contact your account manager. To create a custom route type listener, see [CreateListener](https://help.aliyun.com/document_detail/153253.html).
//
// - An endpoint group is created for the custom route type listener. For more information, see [CreateCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/449363.html).
//
// - An endpoint is created for the custom route type listener. For more information, see [CreateCustomRoutingEndpoints](https://help.aliyun.com/document_detail/449382.html).
//
// @param request - CreateCustomRoutingEndpointTrafficPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomRoutingEndpointTrafficPoliciesResponse
func (client *Client) CreateCustomRoutingEndpointTrafficPoliciesWithOptions(request *CreateCustomRoutingEndpointTrafficPoliciesRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the CreateCustomRoutingEndpointTrafficPolicies operation to create an endpoint traffic policy (custom route type listener).
//
// Description:
//
// This operation takes effect only when the traffic policy of the backend service for the endpoint is set to allow traffic to specified destinations that can accept access traffic. You can invoke [DescribeCustomRoutingEndpoint](https://help.aliyun.com/document_detail/449386.html) to query the traffic policy of the backend service for a specified endpoint. This operation takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom*	- (specifying destinations that can accept access traffic).
//
// Before you invoke this operation, take note of the following items:
//
// - **CreateCustomRoutingEndpointTrafficPolicies*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the endpoint traffic policies for the custom route type listener are not yet created. The creation task continues to run in the background. You can invoke [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) to query the status of the endpoint group to confirm whether the traffic policies are created.
//
//   - If the endpoint group is in the **updating*	- state, the traffic policies are being created. In this state, you can only execute query operations.
//
//   - If the endpoint group is in the **active*	- state, the traffic policies are created.
//
// - The **CreateCustomRoutingEndpointTrafficPolicies*	- operation does not support concurrent creation of endpoint traffic policies within the same Global Accelerator instance.
//
// ### Before you begin
//
// Before you create an endpoint traffic policy, make sure that you have completed the following operations:
//
// - A standard Global Accelerator instance is created. For more information, see [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html).
//
// - If the billing method of the standard Global Accelerator instance is **pay-by-bandwidth**, a basic bandwidth plan is attached to the standard Global Accelerator instance. For more information, see [BandwidthPackageAddAccelerator](https://help.aliyun.com/document_detail/153239.html).
//
// - You have deployed the required applications as backend services to accept forwarded requests from Global Accelerator. The backend service type for custom route type listeners supports only vSwitches.
//
// - You have obtained the permissions to use custom route type listeners and created a custom route type listener. The custom route type for listeners is in invitational preview. To use this feature, contact your account manager. To create a custom route type listener, see [CreateListener](https://help.aliyun.com/document_detail/153253.html).
//
// - An endpoint group is created for the custom route type listener. For more information, see [CreateCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/449363.html).
//
// - An endpoint is created for the custom route type listener. For more information, see [CreateCustomRoutingEndpoints](https://help.aliyun.com/document_detail/449382.html).
//
// @param request - CreateCustomRoutingEndpointTrafficPoliciesRequest
//
// @return CreateCustomRoutingEndpointTrafficPoliciesResponse
func (client *Client) CreateCustomRoutingEndpointTrafficPolicies(request *CreateCustomRoutingEndpointTrafficPoliciesRequest) (_result *CreateCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.CreateCustomRoutingEndpointTrafficPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the CreateCustomRoutingEndpoints operation to create endpoints for a custom route type listener.
//
// Description:
//
// After you configure a custom route type listener, the Alibaba Cloud Global Accelerator (GA) instance generates a port mapping table based on the configured listener port range, the protocol and port range of the destination endpoint group, and the IP address information of the endpoints (vSwitches). This way, traffic is deterministically routed to specific IP addresses and ports in the vSwitches.
//
// This operation creates endpoints for a custom route type listener. When you invoke this operation, take note of the following items:
//
// - **CreateCustomRoutingEndpoints*	- is an asynchronous operation. After a request is sent, the system returns a request ID, but the endpoints are not yet created. The creation node continues to run in the background. You can invoke the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) operation to query the status of the endpoint group to confirm whether the endpoints are created:
//
//   - If the endpoint group is in the **updating*	- state, the endpoints are being created. In this state, you can only execute query operations.
//
//   - If the endpoint group is in the **active*	- state, the endpoints are created.
//
// - The **CreateCustomRoutingEndpoints*	- operation does not support concurrent requests to create endpoints for custom route listeners within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// ### Before you begin
//
// Before you create endpoints for a custom route type listener, make sure that the following operations are complete:
//
// - A standard Global Accelerator instance is created. For more information, see [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html).
//
// - A bandwidth plan is attached to the standard Alibaba Cloud Global Accelerator (GA) instance. For more information, see [BandwidthPackageAddAccelerator](https://help.aliyun.com/document_detail/153239.html).
//
// - Applications are deployed as backend services of Global Accelerator to accept forwarded requests. Custom route type listeners support only vSwitches as the backend service type.
//
// - You have obtained the permissions to use custom route type listeners and created a custom route type listener. The custom route type for listeners is in invitational preview. To use this feature, contact your account manager. To create a custom route type listener, see [CreateListener](https://help.aliyun.com/document_detail/153253.html).
//
// - An endpoint group for the custom route type listener is created. For more information, see [CreateCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/449363.html).
//
// @param request - CreateCustomRoutingEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomRoutingEndpointsResponse
func (client *Client) CreateCustomRoutingEndpointsWithOptions(request *CreateCustomRoutingEndpointsRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomRoutingEndpointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the CreateCustomRoutingEndpoints operation to create endpoints for a custom route type listener.
//
// Description:
//
// After you configure a custom route type listener, the Alibaba Cloud Global Accelerator (GA) instance generates a port mapping table based on the configured listener port range, the protocol and port range of the destination endpoint group, and the IP address information of the endpoints (vSwitches). This way, traffic is deterministically routed to specific IP addresses and ports in the vSwitches.
//
// This operation creates endpoints for a custom route type listener. When you invoke this operation, take note of the following items:
//
// - **CreateCustomRoutingEndpoints*	- is an asynchronous operation. After a request is sent, the system returns a request ID, but the endpoints are not yet created. The creation node continues to run in the background. You can invoke the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) operation to query the status of the endpoint group to confirm whether the endpoints are created:
//
//   - If the endpoint group is in the **updating*	- state, the endpoints are being created. In this state, you can only execute query operations.
//
//   - If the endpoint group is in the **active*	- state, the endpoints are created.
//
// - The **CreateCustomRoutingEndpoints*	- operation does not support concurrent requests to create endpoints for custom route listeners within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// ### Before you begin
//
// Before you create endpoints for a custom route type listener, make sure that the following operations are complete:
//
// - A standard Global Accelerator instance is created. For more information, see [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html).
//
// - A bandwidth plan is attached to the standard Alibaba Cloud Global Accelerator (GA) instance. For more information, see [BandwidthPackageAddAccelerator](https://help.aliyun.com/document_detail/153239.html).
//
// - Applications are deployed as backend services of Global Accelerator to accept forwarded requests. Custom route type listeners support only vSwitches as the backend service type.
//
// - You have obtained the permissions to use custom route type listeners and created a custom route type listener. The custom route type for listeners is in invitational preview. To use this feature, contact your account manager. To create a custom route type listener, see [CreateListener](https://help.aliyun.com/document_detail/153253.html).
//
// - An endpoint group for the custom route type listener is created. For more information, see [CreateCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/449363.html).
//
// @param request - CreateCustomRoutingEndpointsRequest
//
// @return CreateCustomRoutingEndpointsResponse
func (client *Client) CreateCustomRoutingEndpoints(request *CreateCustomRoutingEndpointsRequest) (_result *CreateCustomRoutingEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomRoutingEndpointsResponse{}
	_body, _err := client.CreateCustomRoutingEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an accelerated domain name and associates it with one or more GA instances.
//
// Description:
//
// After you associate an accelerated domain name that has obtained an ICP number with a Global Accelerator (GA) instance, you do not need to complete filing for the accelerated domain name or its subdomains on Alibaba Cloud.
//
// This operation adds an accelerated domain name and associates it with GA instances. Take note of the following items when calling this operation:
//
// - If your accelerated domain name is hosted in the Chinese mainland, you must obtain an ICP number for the domain name.
//
// - The same accelerated domain name cannot be repeatedly associated with the same GA instance.
//
// - You cannot repeatedly call the **CreateDomain*	- operation by using the same Alibaba Cloud account within a specific period of time.
//
// @param request - CreateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an accelerated domain name and associates it with one or more GA instances.
//
// Description:
//
// After you associate an accelerated domain name that has obtained an ICP number with a Global Accelerator (GA) instance, you do not need to complete filing for the accelerated domain name or its subdomains on Alibaba Cloud.
//
// This operation adds an accelerated domain name and associates it with GA instances. Take note of the following items when calling this operation:
//
// - If your accelerated domain name is hosted in the Chinese mainland, you must obtain an ICP number for the domain name.
//
// - The same accelerated domain name cannot be repeatedly associated with the same GA instance.
//
// - You cannot repeatedly call the **CreateDomain*	- operation by using the same Alibaba Cloud account within a specific period of time.
//
// @param request - CreateDomainRequest
//
// @return CreateDomainResponse
func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an endpoint group.
//
// Description:
//
// - Before you create a virtual endpoint group for a Layer 4 listener, you must first create a default endpoint group.
//
// - **CreateEndpointGroup*	- is an asynchronous operation. After you send a request, the system returns an endpoint group ID and begins creating the endpoint group in the background. You can call [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) to query the status of the endpoint group:
//
//   - If the endpoint group is in the **init*	- state, it is being created. In this state, you can only perform query operations.
//
//   - If the endpoint group is in the **active*	- state, it has been created.
//
// - You cannot make concurrent calls to the **CreateEndpointGroup*	- operation for the same Global Accelerator instance.
//
// @param request - CreateEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEndpointGroupResponse
func (client *Client) CreateEndpointGroupWithOptions(request *CreateEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateEndpointGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - Before you create a virtual endpoint group for a Layer 4 listener, you must first create a default endpoint group.
//
// - **CreateEndpointGroup*	- is an asynchronous operation. After you send a request, the system returns an endpoint group ID and begins creating the endpoint group in the background. You can call [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) to query the status of the endpoint group:
//
//   - If the endpoint group is in the **init*	- state, it is being created. In this state, you can only perform query operations.
//
//   - If the endpoint group is in the **active*	- state, it has been created.
//
// - You cannot make concurrent calls to the **CreateEndpointGroup*	- operation for the same Global Accelerator instance.
//
// @param request - CreateEndpointGroupRequest
//
// @return CreateEndpointGroupResponse
func (client *Client) CreateEndpointGroup(request *CreateEndpointGroupRequest) (_result *CreateEndpointGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEndpointGroupResponse{}
	_body, _err := client.CreateEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates endpoint groups in batches.
//
// Description:
//
// - Creates endpoint groups in batches. Default and virtual endpoint groups cannot be created in a single call.
//
// - This API does not support creating virtual endpoint groups for Layer-4 listeners. To create a virtual endpoint group for a Layer-4 listener, call [CreateEndpointGroup](https://help.aliyun.com/document_detail/2302394.html).
//
// - **CreateEndpointGroups*	- is an asynchronous API. It returns a request ID and creates the endpoint groups in the background. You can call [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) or [ListEndpointGroups](https://help.aliyun.com/document_detail/153261.html) to query the status of an endpoint group:
//
//   - If an endpoint group is in the **init*	- state, it is initializing. You can only query the endpoint group in this state.
//
//   - The batch creation is complete when all endpoint groups are in the **active*	- state.
//
// - You cannot make concurrent calls to **CreateEndpointGroups*	- for the same Global Accelerator instance.
//
// @param request - CreateEndpointGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEndpointGroupsResponse
func (client *Client) CreateEndpointGroupsWithOptions(request *CreateEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *CreateEndpointGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates endpoint groups in batches.
//
// Description:
//
// - Creates endpoint groups in batches. Default and virtual endpoint groups cannot be created in a single call.
//
// - This API does not support creating virtual endpoint groups for Layer-4 listeners. To create a virtual endpoint group for a Layer-4 listener, call [CreateEndpointGroup](https://help.aliyun.com/document_detail/2302394.html).
//
// - **CreateEndpointGroups*	- is an asynchronous API. It returns a request ID and creates the endpoint groups in the background. You can call [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) or [ListEndpointGroups](https://help.aliyun.com/document_detail/153261.html) to query the status of an endpoint group:
//
//   - If an endpoint group is in the **init*	- state, it is initializing. You can only query the endpoint group in this state.
//
//   - The batch creation is complete when all endpoint groups are in the **active*	- state.
//
// - You cannot make concurrent calls to **CreateEndpointGroups*	- for the same Global Accelerator instance.
//
// @param request - CreateEndpointGroupsRequest
//
// @return CreateEndpointGroupsResponse
func (client *Client) CreateEndpointGroups(request *CreateEndpointGroupsRequest) (_result *CreateEndpointGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEndpointGroupsResponse{}
	_body, _err := client.CreateEndpointGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If you need to distribute traffic based on request attributes such as the domain name, path, HTTP headers, and cookies, you can create custom forwarding rules for a listener. The listener evaluates incoming requests against these rules and performs different forwarding actions. To create forwarding rules, call the `CreateForwardingRules` API.
//
// Description:
//
// Before you call this API, you should understand how forwarding rules work and their matching conditions. For more information, see [Forwarding rules](https://help.aliyun.com/document_detail/204224.html).
//
// When you call this API, note the following:
//
// - The **CreateForwardingRules*	- API is asynchronous. After the call is made, the system returns a forwarding rule ID, but the rule is still being created. You can call [ListForwardingRules](https://help.aliyun.com/document_detail/205817.html) to query the status of the forwarding rule:
//
//   - If a forwarding rule is in the **configuring*	- status, it is still being created, and you can only perform query operations.
//
//   - If a forwarding rule is in the **active*	- status, the rule has been created.
//
// - You cannot use the **CreateForwardingRules*	- API to create forwarding rules concurrently for the same Global Accelerator instance.
//
// @param request - CreateForwardingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateForwardingRulesResponse
func (client *Client) CreateForwardingRulesWithOptions(request *CreateForwardingRulesRequest, runtime *dara.RuntimeOptions) (_result *CreateForwardingRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If you need to distribute traffic based on request attributes such as the domain name, path, HTTP headers, and cookies, you can create custom forwarding rules for a listener. The listener evaluates incoming requests against these rules and performs different forwarding actions. To create forwarding rules, call the `CreateForwardingRules` API.
//
// Description:
//
// Before you call this API, you should understand how forwarding rules work and their matching conditions. For more information, see [Forwarding rules](https://help.aliyun.com/document_detail/204224.html).
//
// When you call this API, note the following:
//
// - The **CreateForwardingRules*	- API is asynchronous. After the call is made, the system returns a forwarding rule ID, but the rule is still being created. You can call [ListForwardingRules](https://help.aliyun.com/document_detail/205817.html) to query the status of the forwarding rule:
//
//   - If a forwarding rule is in the **configuring*	- status, it is still being created, and you can only perform query operations.
//
//   - If a forwarding rule is in the **active*	- status, the rule has been created.
//
// - You cannot use the **CreateForwardingRules*	- API to create forwarding rules concurrently for the same Global Accelerator instance.
//
// @param request - CreateForwardingRulesRequest
//
// @return CreateForwardingRulesResponse
func (client *Client) CreateForwardingRules(request *CreateForwardingRulesRequest) (_result *CreateForwardingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateForwardingRulesResponse{}
	_body, _err := client.CreateForwardingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateIpSetsWithOptions(request *CreateIpSetsRequest, runtime *dara.RuntimeOptions) (_result *CreateIpSetsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateIpSetsResponse
func (client *Client) CreateIpSets(request *CreateIpSetsRequest) (_result *CreateIpSetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIpSetsResponse{}
	_body, _err := client.CreateIpSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a listener for your GA instance.
//
// Description:
//
// Note the following when you call this operation:
//
// - **CreateListener*	- is an asynchronous operation. After you send a request, the system returns a listener ID but the listener is still being created in the background. You can call [DescribeListener](https://help.aliyun.com/document_detail/153254.html) to check the listener\\"s status:
//
//   - An **init*	- status indicates that the listener is being created. In this state, you can only perform query operations.
//
//   - An **active*	- status indicates that the listener is ready.
//
// - You cannot concurrently create multiple listeners for the same Global Accelerator instance by using the **CreateListener*	- operation.
//
// @param request - CreateListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateListenerResponse
func (client *Client) CreateListenerWithOptions(request *CreateListenerRequest, runtime *dara.RuntimeOptions) (_result *CreateListenerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a listener for your GA instance.
//
// Description:
//
// Note the following when you call this operation:
//
// - **CreateListener*	- is an asynchronous operation. After you send a request, the system returns a listener ID but the listener is still being created in the background. You can call [DescribeListener](https://help.aliyun.com/document_detail/153254.html) to check the listener\\"s status:
//
//   - An **init*	- status indicates that the listener is being created. In this state, you can only perform query operations.
//
//   - An **active*	- status indicates that the listener is ready.
//
// - You cannot concurrently create multiple listeners for the same Global Accelerator instance by using the **CreateListener*	- operation.
//
// @param request - CreateListenerRequest
//
// @return CreateListenerResponse
func (client *Client) CreateListener(request *CreateListenerRequest) (_result *CreateListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateListenerResponse{}
	_body, _err := client.CreateListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateSpareIpsWithOptions(request *CreateSpareIpsRequest, runtime *dara.RuntimeOptions) (_result *CreateSpareIpsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateSpareIpsResponse
func (client *Client) CreateSpareIps(request *CreateSpareIpsRequest) (_result *CreateSpareIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSpareIpsResponse{}
	_body, _err := client.CreateSpareIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Global Accelerator (GA) instance.
//
// Description:
//
// - You cannot delete subscription GA instances.
//
// - **DeleteAccelerator*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the status of the task.
//
//   - If the GA instance is in the **deleting*	- state, the GA instance is being deleted. In this case, you can perform only query operations.
//
//   - If the GA instance cannot be queried, the GA instance is deleted.
//
// @param request - DeleteAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAcceleratorResponse
func (client *Client) DeleteAcceleratorWithOptions(request *DeleteAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *DeleteAcceleratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - You cannot delete subscription GA instances.
//
// - **DeleteAccelerator*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the status of the task.
//
//   - If the GA instance is in the **deleting*	- state, the GA instance is being deleted. In this case, you can perform only query operations.
//
//   - If the GA instance cannot be queried, the GA instance is deleted.
//
// @param request - DeleteAcceleratorRequest
//
// @return DeleteAcceleratorResponse
func (client *Client) DeleteAccelerator(request *DeleteAcceleratorRequest) (_result *DeleteAcceleratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAcceleratorResponse{}
	_body, _err := client.DeleteAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the DeleteAcl operation to delete an access control policy group.
//
// Description:
//
// *DeleteAcl*	- is an asynchronous operation. After a request is sent, the system returns a request ID, but the access control policy group is not immediately deleted. The deletion node continues to run in the background. You can invoke [GetAcl](https://help.aliyun.com/document_detail/258292.html) to query the status of the access control policy group:
//
// - If the access control policy group is in the **deleting*	- state, the access control policy group is being deleted. In this state, you can only execute query operations and cannot execute other operations.
//
// - If the access control policy group cannot be found, the access control policy group is deleted.
//
// @param request - DeleteAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAclResponse
func (client *Client) DeleteAclWithOptions(request *DeleteAclRequest, runtime *dara.RuntimeOptions) (_result *DeleteAclResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the DeleteAcl operation to delete an access control policy group.
//
// Description:
//
// *DeleteAcl*	- is an asynchronous operation. After a request is sent, the system returns a request ID, but the access control policy group is not immediately deleted. The deletion node continues to run in the background. You can invoke [GetAcl](https://help.aliyun.com/document_detail/258292.html) to query the status of the access control policy group:
//
// - If the access control policy group is in the **deleting*	- state, the access control policy group is being deleted. In this state, you can only execute query operations and cannot execute other operations.
//
// - If the access control policy group cannot be found, the access control policy group is deleted.
//
// @param request - DeleteAclRequest
//
// @return DeleteAclResponse
func (client *Client) DeleteAcl(request *DeleteAclRequest) (_result *DeleteAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAclResponse{}
	_body, _err := client.DeleteAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DeleteApplicationMonitor operation to delete an origin probing task.
//
// Description:
//
// - The **DeleteApplicationMonitor*	- operation is asynchronous. After you send a request, the system returns a request ID, but the origin probing node is not yet deleted. The deletion node continues to run in the background. You can invoke [ListApplicationMonitor](https://help.aliyun.com/document_detail/408462.html) to query the status of the origin probing node:
//
//   - If the origin probing node is in the **deleting*	- state, the node is being deleted. In this state, you can only execute query operations.
//
//   - If the origin probing node cannot be found, the node is deleted.
//
// - The **DeleteApplicationMonitor*	- operation does not support concurrent deletion of origin probing nodes within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - DeleteApplicationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationMonitorResponse
func (client *Client) DeleteApplicationMonitorWithOptions(request *DeleteApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationMonitorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DeleteApplicationMonitor operation to delete an origin probing task.
//
// Description:
//
// - The **DeleteApplicationMonitor*	- operation is asynchronous. After you send a request, the system returns a request ID, but the origin probing node is not yet deleted. The deletion node continues to run in the background. You can invoke [ListApplicationMonitor](https://help.aliyun.com/document_detail/408462.html) to query the status of the origin probing node:
//
//   - If the origin probing node is in the **deleting*	- state, the node is being deleted. In this state, you can only execute query operations.
//
//   - If the origin probing node cannot be found, the node is deleted.
//
// - The **DeleteApplicationMonitor*	- operation does not support concurrent deletion of origin probing nodes within the same Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - DeleteApplicationMonitorRequest
//
// @return DeleteApplicationMonitorResponse
func (client *Client) DeleteApplicationMonitor(request *DeleteApplicationMonitorRequest) (_result *DeleteApplicationMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApplicationMonitorResponse{}
	_body, _err := client.DeleteApplicationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteBandwidthPackageWithOptions(request *DeleteBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *DeleteBandwidthPackageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteBandwidthPackageResponse
func (client *Client) DeleteBandwidthPackage(request *DeleteBandwidthPackageRequest) (_result *DeleteBandwidthPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBandwidthPackageResponse{}
	_body, _err := client.DeleteBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteBasicAccelerateIpWithOptions(request *DeleteBasicAccelerateIpRequest, runtime *dara.RuntimeOptions) (_result *DeleteBasicAccelerateIpResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteBasicAccelerateIpResponse
func (client *Client) DeleteBasicAccelerateIp(request *DeleteBasicAccelerateIpRequest) (_result *DeleteBasicAccelerateIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBasicAccelerateIpResponse{}
	_body, _err := client.DeleteBasicAccelerateIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the attach relationship between an accelerated IP address and an endpoint of a basic Global Accelerator (GA) instance.
//
// Description:
//
// - **DeleteBasicAccelerateIpEndpointRelation*	- is an asynchronous operation. After a request is sent, the system returns a request ID, but the attach relationship between the accelerated IP address and the endpoint of the basic Global Accelerator (GA) instance is not immediately removed. The deletion task continues to run in the background. You can call the following operations to check whether the attach relationship between the accelerated IP address and the endpoint is deleted:
//
//   - Call [GetBasicAccelerateIp](https://help.aliyun.com/document_detail/466794.html) or [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) to query the status of the accelerated IP address and the endpoint respectively. If the status of the accelerated IP address and the endpoint is **unbinding**, the attach relationship is being deleted. In this state, you can only perform query operations and cannot perform other operations.
//
//   - Call [ListBasicAccelerateIpEndpointRelations](https://help.aliyun.com/document_detail/466803.html) to query the attach status between the accelerated IP address and the endpoint. If no attach information is returned, the attach relationship between the accelerated IP address and the endpoint is deleted.
//
// - **DeleteBasicAccelerateIpEndpointRelation*	- does not support concurrent deletion of attach relationships between accelerated IP addresses and endpoints within the same basic GA instance.
//
// @param request - DeleteBasicAccelerateIpEndpointRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBasicAccelerateIpEndpointRelationResponse
func (client *Client) DeleteBasicAccelerateIpEndpointRelationWithOptions(request *DeleteBasicAccelerateIpEndpointRelationRequest, runtime *dara.RuntimeOptions) (_result *DeleteBasicAccelerateIpEndpointRelationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the attach relationship between an accelerated IP address and an endpoint of a basic Global Accelerator (GA) instance.
//
// Description:
//
// - **DeleteBasicAccelerateIpEndpointRelation*	- is an asynchronous operation. After a request is sent, the system returns a request ID, but the attach relationship between the accelerated IP address and the endpoint of the basic Global Accelerator (GA) instance is not immediately removed. The deletion task continues to run in the background. You can call the following operations to check whether the attach relationship between the accelerated IP address and the endpoint is deleted:
//
//   - Call [GetBasicAccelerateIp](https://help.aliyun.com/document_detail/466794.html) or [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) to query the status of the accelerated IP address and the endpoint respectively. If the status of the accelerated IP address and the endpoint is **unbinding**, the attach relationship is being deleted. In this state, you can only perform query operations and cannot perform other operations.
//
//   - Call [ListBasicAccelerateIpEndpointRelations](https://help.aliyun.com/document_detail/466803.html) to query the attach status between the accelerated IP address and the endpoint. If no attach information is returned, the attach relationship between the accelerated IP address and the endpoint is deleted.
//
// - **DeleteBasicAccelerateIpEndpointRelation*	- does not support concurrent deletion of attach relationships between accelerated IP addresses and endpoints within the same basic GA instance.
//
// @param request - DeleteBasicAccelerateIpEndpointRelationRequest
//
// @return DeleteBasicAccelerateIpEndpointRelationResponse
func (client *Client) DeleteBasicAccelerateIpEndpointRelation(request *DeleteBasicAccelerateIpEndpointRelationRequest) (_result *DeleteBasicAccelerateIpEndpointRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBasicAccelerateIpEndpointRelationResponse{}
	_body, _err := client.DeleteBasicAccelerateIpEndpointRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the DeleteBasicAccelerator operation to delete a specified basic Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// - Upfront (subscription) basic Alibaba Cloud Global Accelerator (GA) instances cannot be deleted. You can unsubscribe on the <props="china">[Unsubscribe](https://usercenter2.aliyun.com/refund/refund)<props="intl">[Unsubscribe](https://usercenter2-intl.aliyun.com/refund/refund) page. Before you unsubscribe, make sure that the basic Alibaba Cloud Global Accelerator (GA) instance has no acceleration area or endpoint group configurations and is not attached to a bandwidth plan.
//
//   - To delete an acceleration area, refer to [DeleteBasicIpSet](https://help.aliyun.com/document_detail/2253388.html).
//
//   - To delete an endpoint group, refer to [DeleteBasicEndpointGroup](https://help.aliyun.com/document_detail/2253399.html).
//
//   - To disassociate a bandwidth plan from a basic Alibaba Cloud Global Accelerator (GA) instance, refer to [BandwidthPackageRemoveAccelerator](https://help.aliyun.com/document_detail/153240.html).
//
// - Before you invoke this operation to delete a pay-as-you-go basic Alibaba Cloud Global Accelerator (GA) instance, make sure that data migration is complete and that the acceleration area and endpoint group configurations under the instance are deleted.
//
//   - To delete an acceleration area, refer to [DeleteBasicIpSet](https://help.aliyun.com/document_detail/2253388.html).
//
//   - To delete an endpoint group, refer to [DeleteBasicEndpointGroup](https://help.aliyun.com/document_detail/2253399.html).
//
// - **DeleteBasicAccelerator*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the node in the background. You can invoke [GetBasicAccelerator](https://help.aliyun.com/document_detail/353188.html) to query the status of the basic Alibaba Cloud Global Accelerator (GA) instance:
//
//   - If the instance is in the **deleting*	- state, the instance is being deleted. In this state, you can only execute query operations.
//
//   - If the instance cannot be found, the instance is deleted.
//
// @param request - DeleteBasicAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBasicAcceleratorResponse
func (client *Client) DeleteBasicAcceleratorWithOptions(request *DeleteBasicAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *DeleteBasicAcceleratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the DeleteBasicAccelerator operation to delete a specified basic Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// - Upfront (subscription) basic Alibaba Cloud Global Accelerator (GA) instances cannot be deleted. You can unsubscribe on the <props="china">[Unsubscribe](https://usercenter2.aliyun.com/refund/refund)<props="intl">[Unsubscribe](https://usercenter2-intl.aliyun.com/refund/refund) page. Before you unsubscribe, make sure that the basic Alibaba Cloud Global Accelerator (GA) instance has no acceleration area or endpoint group configurations and is not attached to a bandwidth plan.
//
//   - To delete an acceleration area, refer to [DeleteBasicIpSet](https://help.aliyun.com/document_detail/2253388.html).
//
//   - To delete an endpoint group, refer to [DeleteBasicEndpointGroup](https://help.aliyun.com/document_detail/2253399.html).
//
//   - To disassociate a bandwidth plan from a basic Alibaba Cloud Global Accelerator (GA) instance, refer to [BandwidthPackageRemoveAccelerator](https://help.aliyun.com/document_detail/153240.html).
//
// - Before you invoke this operation to delete a pay-as-you-go basic Alibaba Cloud Global Accelerator (GA) instance, make sure that data migration is complete and that the acceleration area and endpoint group configurations under the instance are deleted.
//
//   - To delete an acceleration area, refer to [DeleteBasicIpSet](https://help.aliyun.com/document_detail/2253388.html).
//
//   - To delete an endpoint group, refer to [DeleteBasicEndpointGroup](https://help.aliyun.com/document_detail/2253399.html).
//
// - **DeleteBasicAccelerator*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the node in the background. You can invoke [GetBasicAccelerator](https://help.aliyun.com/document_detail/353188.html) to query the status of the basic Alibaba Cloud Global Accelerator (GA) instance:
//
//   - If the instance is in the **deleting*	- state, the instance is being deleted. In this state, you can only execute query operations.
//
//   - If the instance cannot be found, the instance is deleted.
//
// @param request - DeleteBasicAcceleratorRequest
//
// @return DeleteBasicAcceleratorResponse
func (client *Client) DeleteBasicAccelerator(request *DeleteBasicAcceleratorRequest) (_result *DeleteBasicAcceleratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBasicAcceleratorResponse{}
	_body, _err := client.DeleteBasicAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the DeleteBasicEndpoint operation to delete an endpoint of a basic Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// - **DeleteBasicEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID, but the endpoint is not yet deleted and the deletion node continues in the background. You can invoke [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) to query the status of the endpoint:
//
//   - If the endpoint is in the **deleting*	- state, the endpoint is being deleted. In this state, you can only execute query operations.
//
//   - If the endpoint cannot be found, the endpoint is deleted.
//
// - **DeleteBasicEndpoint*	- does not support concurrent deletion of endpoints within the same basic Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - DeleteBasicEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBasicEndpointResponse
func (client *Client) DeleteBasicEndpointWithOptions(request *DeleteBasicEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteBasicEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the DeleteBasicEndpoint operation to delete an endpoint of a basic Alibaba Cloud Global Accelerator (GA) instance.
//
// Description:
//
// - **DeleteBasicEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID, but the endpoint is not yet deleted and the deletion node continues in the background. You can invoke [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) to query the status of the endpoint:
//
//   - If the endpoint is in the **deleting*	- state, the endpoint is being deleted. In this state, you can only execute query operations.
//
//   - If the endpoint cannot be found, the endpoint is deleted.
//
// - **DeleteBasicEndpoint*	- does not support concurrent deletion of endpoints within the same basic Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - DeleteBasicEndpointRequest
//
// @return DeleteBasicEndpointResponse
func (client *Client) DeleteBasicEndpoint(request *DeleteBasicEndpointRequest) (_result *DeleteBasicEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBasicEndpointResponse{}
	_body, _err := client.DeleteBasicEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteBasicEndpointGroupWithOptions(request *DeleteBasicEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteBasicEndpointGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteBasicEndpointGroupResponse
func (client *Client) DeleteBasicEndpointGroup(request *DeleteBasicEndpointGroupRequest) (_result *DeleteBasicEndpointGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBasicEndpointGroupResponse{}
	_body, _err := client.DeleteBasicEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteBasicIpSetWithOptions(request *DeleteBasicIpSetRequest, runtime *dara.RuntimeOptions) (_result *DeleteBasicIpSetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteBasicIpSetResponse
func (client *Client) DeleteBasicIpSet(request *DeleteBasicIpSetRequest) (_result *DeleteBasicIpSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBasicIpSetResponse{}
	_body, _err := client.DeleteBasicIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteCustomRoutingEndpointGroupDestinationsWithOptions(request *DeleteCustomRoutingEndpointGroupDestinationsRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomRoutingEndpointGroupDestinationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteCustomRoutingEndpointGroupDestinationsResponse
func (client *Client) DeleteCustomRoutingEndpointGroupDestinations(request *DeleteCustomRoutingEndpointGroupDestinationsRequest) (_result *DeleteCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.DeleteCustomRoutingEndpointGroupDestinationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteCustomRoutingEndpointGroupsWithOptions(request *DeleteCustomRoutingEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomRoutingEndpointGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteCustomRoutingEndpointGroupsResponse
func (client *Client) DeleteCustomRoutingEndpointGroups(request *DeleteCustomRoutingEndpointGroupsRequest) (_result *DeleteCustomRoutingEndpointGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomRoutingEndpointGroupsResponse{}
	_body, _err := client.DeleteCustomRoutingEndpointGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteCustomRoutingEndpointTrafficPoliciesWithOptions(request *DeleteCustomRoutingEndpointTrafficPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteCustomRoutingEndpointTrafficPoliciesResponse
func (client *Client) DeleteCustomRoutingEndpointTrafficPolicies(request *DeleteCustomRoutingEndpointTrafficPoliciesRequest) (_result *DeleteCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.DeleteCustomRoutingEndpointTrafficPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteCustomRoutingEndpointsWithOptions(request *DeleteCustomRoutingEndpointsRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomRoutingEndpointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteCustomRoutingEndpointsResponse
func (client *Client) DeleteCustomRoutingEndpoints(request *DeleteCustomRoutingEndpointsRequest) (_result *DeleteCustomRoutingEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomRoutingEndpointsResponse{}
	_body, _err := client.DeleteCustomRoutingEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDomainAcceleratorRelationWithOptions(request *DeleteDomainAcceleratorRelationRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainAcceleratorRelationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteDomainAcceleratorRelationResponse
func (client *Client) DeleteDomainAcceleratorRelation(request *DeleteDomainAcceleratorRelationRequest) (_result *DeleteDomainAcceleratorRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainAcceleratorRelationResponse{}
	_body, _err := client.DeleteDomainAcceleratorRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteEndpointGroupWithOptions(request *DeleteEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteEndpointGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteEndpointGroupResponse
func (client *Client) DeleteEndpointGroup(request *DeleteEndpointGroupRequest) (_result *DeleteEndpointGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEndpointGroupResponse{}
	_body, _err := client.DeleteEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteEndpointGroupsWithOptions(request *DeleteEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *DeleteEndpointGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteEndpointGroupsResponse
func (client *Client) DeleteEndpointGroups(request *DeleteEndpointGroupsRequest) (_result *DeleteEndpointGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEndpointGroupsResponse{}
	_body, _err := client.DeleteEndpointGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteForwardingRulesWithOptions(request *DeleteForwardingRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteForwardingRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteForwardingRulesResponse
func (client *Client) DeleteForwardingRules(request *DeleteForwardingRulesRequest) (_result *DeleteForwardingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteForwardingRulesResponse{}
	_body, _err := client.DeleteForwardingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteIpSetWithOptions(request *DeleteIpSetRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpSetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteIpSetResponse
func (client *Client) DeleteIpSet(request *DeleteIpSetRequest) (_result *DeleteIpSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIpSetResponse{}
	_body, _err := client.DeleteIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteIpSetsWithOptions(request *DeleteIpSetsRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpSetsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteIpSetsResponse
func (client *Client) DeleteIpSets(request *DeleteIpSetsRequest) (_result *DeleteIpSetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIpSetsResponse{}
	_body, _err := client.DeleteIpSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteListenerWithOptions(request *DeleteListenerRequest, runtime *dara.RuntimeOptions) (_result *DeleteListenerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteListenerResponse
func (client *Client) DeleteListener(request *DeleteListenerRequest) (_result *DeleteListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteListenerResponse{}
	_body, _err := client.DeleteListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteSpareIpsWithOptions(request *DeleteSpareIpsRequest, runtime *dara.RuntimeOptions) (_result *DeleteSpareIpsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteSpareIpsResponse
func (client *Client) DeleteSpareIps(request *DeleteSpareIpsRequest) (_result *DeleteSpareIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSpareIpsResponse{}
	_body, _err := client.DeleteSpareIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeAccelerator operation to query information about a specified Global Accelerator instance.
//
// @param request - DescribeAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAcceleratorResponse
func (client *Client) DescribeAcceleratorWithOptions(request *DescribeAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *DescribeAcceleratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeAccelerator operation to query information about a specified Global Accelerator instance.
//
// @param request - DescribeAcceleratorRequest
//
// @return DescribeAcceleratorResponse
func (client *Client) DescribeAccelerator(request *DescribeAcceleratorRequest) (_result *DescribeAcceleratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAcceleratorResponse{}
	_body, _err := client.DescribeAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeAcceleratorAutoRenewAttributeWithOptions(request *DescribeAcceleratorAutoRenewAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAcceleratorAutoRenewAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeAcceleratorAutoRenewAttributeResponse
func (client *Client) DescribeAcceleratorAutoRenewAttribute(request *DescribeAcceleratorAutoRenewAttributeRequest) (_result *DescribeAcceleratorAutoRenewAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAcceleratorAutoRenewAttributeResponse{}
	_body, _err := client.DescribeAcceleratorAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeAcceleratorServiceStatusWithOptions(request *DescribeAcceleratorServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAcceleratorServiceStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeAcceleratorServiceStatusResponse
func (client *Client) DescribeAcceleratorServiceStatus(request *DescribeAcceleratorServiceStatusRequest) (_result *DescribeAcceleratorServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAcceleratorServiceStatusResponse{}
	_body, _err := client.DescribeAcceleratorServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeApplicationMonitorWithOptions(request *DescribeApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeApplicationMonitorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeApplicationMonitorResponse
func (client *Client) DescribeApplicationMonitor(request *DescribeApplicationMonitorRequest) (_result *DescribeApplicationMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApplicationMonitorResponse{}
	_body, _err := client.DescribeApplicationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a bandwidth plan by calling the DescribeBandwidthPackage operation.
//
// @param request - DescribeBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBandwidthPackageResponse
func (client *Client) DescribeBandwidthPackageWithOptions(request *DescribeBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *DescribeBandwidthPackageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a bandwidth plan by calling the DescribeBandwidthPackage operation.
//
// @param request - DescribeBandwidthPackageRequest
//
// @return DescribeBandwidthPackageResponse
func (client *Client) DescribeBandwidthPackage(request *DescribeBandwidthPackageRequest) (_result *DescribeBandwidthPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBandwidthPackageResponse{}
	_body, _err := client.DescribeBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeBandwidthPackageAutoRenewAttributeWithOptions(request *DescribeBandwidthPackageAutoRenewAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeBandwidthPackageAutoRenewAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeBandwidthPackageAutoRenewAttributeResponse
func (client *Client) DescribeBandwidthPackageAutoRenewAttribute(request *DescribeBandwidthPackageAutoRenewAttributeRequest) (_result *DescribeBandwidthPackageAutoRenewAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBandwidthPackageAutoRenewAttributeResponse{}
	_body, _err := client.DescribeBandwidthPackageAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the pricing and specification details of Global Accelerator commodity options available for purchase.
//
// @param request - DescribeCommodityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCommodityResponse
func (client *Client) DescribeCommodityWithOptions(request *DescribeCommodityRequest, runtime *dara.RuntimeOptions) (_result *DescribeCommodityResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the pricing and specification details of Global Accelerator commodity options available for purchase.
//
// @param request - DescribeCommodityRequest
//
// @return DescribeCommodityResponse
func (client *Client) DescribeCommodity(request *DescribeCommodityRequest) (_result *DescribeCommodityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCommodityResponse{}
	_body, _err := client.DescribeCommodityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCommodityPriceWithOptions(request *DescribeCommodityPriceRequest, runtime *dara.RuntimeOptions) (_result *DescribeCommodityPriceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCommodityPriceResponse
func (client *Client) DescribeCommodityPrice(request *DescribeCommodityPriceRequest) (_result *DescribeCommodityPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCommodityPriceResponse{}
	_body, _err := client.DescribeCommodityPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCustomRoutingEndPointTrafficPolicyWithOptions(request *DescribeCustomRoutingEndPointTrafficPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomRoutingEndPointTrafficPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCustomRoutingEndPointTrafficPolicyResponse
func (client *Client) DescribeCustomRoutingEndPointTrafficPolicy(request *DescribeCustomRoutingEndPointTrafficPolicyRequest) (_result *DescribeCustomRoutingEndPointTrafficPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomRoutingEndPointTrafficPolicyResponse{}
	_body, _err := client.DescribeCustomRoutingEndPointTrafficPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCustomRoutingEndpointWithOptions(request *DescribeCustomRoutingEndpointRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomRoutingEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCustomRoutingEndpointResponse
func (client *Client) DescribeCustomRoutingEndpoint(request *DescribeCustomRoutingEndpointRequest) (_result *DescribeCustomRoutingEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomRoutingEndpointResponse{}
	_body, _err := client.DescribeCustomRoutingEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCustomRoutingEndpointGroupWithOptions(request *DescribeCustomRoutingEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomRoutingEndpointGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCustomRoutingEndpointGroupResponse
func (client *Client) DescribeCustomRoutingEndpointGroup(request *DescribeCustomRoutingEndpointGroupRequest) (_result *DescribeCustomRoutingEndpointGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomRoutingEndpointGroupResponse{}
	_body, _err := client.DescribeCustomRoutingEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCustomRoutingEndpointGroupDestinationsWithOptions(request *DescribeCustomRoutingEndpointGroupDestinationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomRoutingEndpointGroupDestinationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCustomRoutingEndpointGroupDestinationsResponse
func (client *Client) DescribeCustomRoutingEndpointGroupDestinations(request *DescribeCustomRoutingEndpointGroupDestinationsRequest) (_result *DescribeCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.DescribeCustomRoutingEndpointGroupDestinationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a specified endpoint group.
//
// @param request - DescribeEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndpointGroupResponse
func (client *Client) DescribeEndpointGroupWithOptions(request *DescribeEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeEndpointGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a specified endpoint group.
//
// @param request - DescribeEndpointGroupRequest
//
// @return DescribeEndpointGroupResponse
func (client *Client) DescribeEndpointGroup(request *DescribeEndpointGroupRequest) (_result *DescribeEndpointGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEndpointGroupResponse{}
	_body, _err := client.DescribeEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns the configuration of a specified acceleration region, including its accelerated IP addresses.
//
// @param request - DescribeIpSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpSetResponse
func (client *Client) DescribeIpSetWithOptions(request *DescribeIpSetRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpSetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the configuration of a specified acceleration region, including its accelerated IP addresses.
//
// @param request - DescribeIpSetRequest
//
// @return DescribeIpSetResponse
func (client *Client) DescribeIpSet(request *DescribeIpSetRequest) (_result *DescribeIpSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIpSetResponse{}
	_body, _err := client.DescribeIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the configurations of a specific listener.
//
// Description:
//
// This operation queries the configuration of a specified listener, such as its routing type, status, creation timestamp, and port information.
//
// @param request - DescribeListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeListenerResponse
func (client *Client) DescribeListenerWithOptions(request *DescribeListenerRequest, runtime *dara.RuntimeOptions) (_result *DescribeListenerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the configurations of a specific listener.
//
// Description:
//
// This operation queries the configuration of a specified listener, such as its routing type, status, creation timestamp, and port information.
//
// @param request - DescribeListenerRequest
//
// @return DescribeListenerResponse
func (client *Client) DescribeListener(request *DescribeListenerRequest) (_result *DescribeListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeListenerResponse{}
	_body, _err := client.DescribeListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeLogStoreOfEndpointGroupWithOptions(request *DescribeLogStoreOfEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogStoreOfEndpointGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeLogStoreOfEndpointGroupResponse
func (client *Client) DescribeLogStoreOfEndpointGroup(request *DescribeLogStoreOfEndpointGroupRequest) (_result *DescribeLogStoreOfEndpointGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogStoreOfEndpointGroupResponse{}
	_body, _err := client.DescribeLogStoreOfEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Call the DetachDdosFromAccelerator operation to detach an Anti-DDoS Pro or Anti-DDoS Premium instance from a Global Accelerator instance.
//
// Description:
//
// - **DetachDdosFromAccelerator*	- is an asynchronous operation. After you send a request, the system returns a request ID and performs the operation in the background. The Anti-DDoS Pro or Anti-DDoS Premium instance is not immediately detached. You can call [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) or [ListAccelerators](https://help.aliyun.com/document_detail/153236.html) to query the state of the Global Accelerator instance:
//
//   - If the Global Accelerator instance is in the **configuring*	- state, the Anti-DDoS Pro or Anti-DDoS Premium instance is being detached. In this state, you can only perform query operations.
//
//   - If the Global Accelerator instance is in the **active*	- state, the Anti-DDoS Pro or Anti-DDoS Premium instance is detached.
//
// - The **DetachDdosFromAccelerator*	- operation does not support concurrent requests to detach Anti-DDoS instances from the same Global Accelerator instance.
//
// @param request - DetachDdosFromAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachDdosFromAcceleratorResponse
func (client *Client) DetachDdosFromAcceleratorWithOptions(request *DetachDdosFromAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *DetachDdosFromAcceleratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DetachDdosFromAccelerator operation to detach an Anti-DDoS Pro or Anti-DDoS Premium instance from a Global Accelerator instance.
//
// Description:
//
// - **DetachDdosFromAccelerator*	- is an asynchronous operation. After you send a request, the system returns a request ID and performs the operation in the background. The Anti-DDoS Pro or Anti-DDoS Premium instance is not immediately detached. You can call [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) or [ListAccelerators](https://help.aliyun.com/document_detail/153236.html) to query the state of the Global Accelerator instance:
//
//   - If the Global Accelerator instance is in the **configuring*	- state, the Anti-DDoS Pro or Anti-DDoS Premium instance is being detached. In this state, you can only perform query operations.
//
//   - If the Global Accelerator instance is in the **active*	- state, the Anti-DDoS Pro or Anti-DDoS Premium instance is detached.
//
// - The **DetachDdosFromAccelerator*	- operation does not support concurrent requests to detach Anti-DDoS instances from the same Global Accelerator instance.
//
// @param request - DetachDdosFromAcceleratorRequest
//
// @return DetachDdosFromAcceleratorResponse
func (client *Client) DetachDdosFromAccelerator(request *DetachDdosFromAcceleratorRequest) (_result *DetachDdosFromAcceleratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachDdosFromAcceleratorResponse{}
	_body, _err := client.DetachDdosFromAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DetachLogStoreFromEndpointGroupWithOptions(request *DetachLogStoreFromEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *DetachLogStoreFromEndpointGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DetachLogStoreFromEndpointGroupResponse
func (client *Client) DetachLogStoreFromEndpointGroup(request *DetachLogStoreFromEndpointGroupRequest) (_result *DetachLogStoreFromEndpointGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachLogStoreFromEndpointGroupResponse{}
	_body, _err := client.DetachLogStoreFromEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DetectApplicationMonitorWithOptions(request *DetectApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *DetectApplicationMonitorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DetectApplicationMonitorResponse
func (client *Client) DetectApplicationMonitor(request *DetectApplicationMonitorRequest) (_result *DetectApplicationMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectApplicationMonitorResponse{}
	_body, _err := client.DetectApplicationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DisableApplicationMonitorWithOptions(request *DisableApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationMonitorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DisableApplicationMonitorResponse
func (client *Client) DisableApplicationMonitor(request *DisableApplicationMonitorRequest) (_result *DisableApplicationMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableApplicationMonitorResponse{}
	_body, _err := client.DisableApplicationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Detach integrated cloud product from GA
//
// @param request - DisassociateResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateResourcesResponse
func (client *Client) DisassociateResourcesWithOptions(request *DisassociateResourcesRequest, runtime *dara.RuntimeOptions) (_result *DisassociateResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Detach integrated cloud product from GA
//
// @param request - DisassociateResourcesRequest
//
// @return DisassociateResourcesResponse
func (client *Client) DisassociateResources(request *DisassociateResourcesRequest) (_result *DisassociateResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisassociateResourcesResponse{}
	_body, _err := client.DisassociateResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DissociateAclsFromListenerWithOptions(request *DissociateAclsFromListenerRequest, runtime *dara.RuntimeOptions) (_result *DissociateAclsFromListenerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DissociateAclsFromListenerResponse
func (client *Client) DissociateAclsFromListener(request *DissociateAclsFromListenerRequest) (_result *DissociateAclsFromListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DissociateAclsFromListenerResponse{}
	_body, _err := client.DissociateAclsFromListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DissociateAdditionalCertificatesFromListenerWithOptions(request *DissociateAdditionalCertificatesFromListenerRequest, runtime *dara.RuntimeOptions) (_result *DissociateAdditionalCertificatesFromListenerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DissociateAdditionalCertificatesFromListenerResponse
func (client *Client) DissociateAdditionalCertificatesFromListener(request *DissociateAdditionalCertificatesFromListenerRequest) (_result *DissociateAdditionalCertificatesFromListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DissociateAdditionalCertificatesFromListenerResponse{}
	_body, _err := client.DissociateAdditionalCertificatesFromListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) EnableApplicationMonitorWithOptions(request *EnableApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationMonitorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return EnableApplicationMonitorResponse
func (client *Client) EnableApplicationMonitor(request *EnableApplicationMonitorRequest) (_result *EnableApplicationMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableApplicationMonitorResponse{}
	_body, _err := client.EnableApplicationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAclWithOptions(request *GetAclRequest, runtime *dara.RuntimeOptions) (_result *GetAclResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAclResponse
func (client *Client) GetAcl(request *GetAclRequest) (_result *GetAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAclResponse{}
	_body, _err := client.GetAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetBasicAccelerateIpWithOptions(request *GetBasicAccelerateIpRequest, runtime *dara.RuntimeOptions) (_result *GetBasicAccelerateIpResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetBasicAccelerateIpResponse
func (client *Client) GetBasicAccelerateIp(request *GetBasicAccelerateIpRequest) (_result *GetBasicAccelerateIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBasicAccelerateIpResponse{}
	_body, _err := client.GetBasicAccelerateIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attachment information of an accelerated IP address or endpoint of a basic Global Accelerator (GA) instance.
//
// @param request - GetBasicAccelerateIpEndpointRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBasicAccelerateIpEndpointRelationResponse
func (client *Client) GetBasicAccelerateIpEndpointRelationWithOptions(request *GetBasicAccelerateIpEndpointRelationRequest, runtime *dara.RuntimeOptions) (_result *GetBasicAccelerateIpEndpointRelationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attachment information of an accelerated IP address or endpoint of a basic Global Accelerator (GA) instance.
//
// @param request - GetBasicAccelerateIpEndpointRelationRequest
//
// @return GetBasicAccelerateIpEndpointRelationResponse
func (client *Client) GetBasicAccelerateIpEndpointRelation(request *GetBasicAccelerateIpEndpointRelationRequest) (_result *GetBasicAccelerateIpEndpointRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBasicAccelerateIpEndpointRelationResponse{}
	_body, _err := client.GetBasicAccelerateIpEndpointRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the GetBasicAccelerateIpIdleCount operation to query the number of idle accelerated IP addresses of a basic Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - GetBasicAccelerateIpIdleCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBasicAccelerateIpIdleCountResponse
func (client *Client) GetBasicAccelerateIpIdleCountWithOptions(request *GetBasicAccelerateIpIdleCountRequest, runtime *dara.RuntimeOptions) (_result *GetBasicAccelerateIpIdleCountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the GetBasicAccelerateIpIdleCount operation to query the number of idle accelerated IP addresses of a basic Alibaba Cloud Global Accelerator (GA) instance.
//
// @param request - GetBasicAccelerateIpIdleCountRequest
//
// @return GetBasicAccelerateIpIdleCountResponse
func (client *Client) GetBasicAccelerateIpIdleCount(request *GetBasicAccelerateIpIdleCountRequest) (_result *GetBasicAccelerateIpIdleCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBasicAccelerateIpIdleCountResponse{}
	_body, _err := client.GetBasicAccelerateIpIdleCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetBasicAcceleratorWithOptions(request *GetBasicAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *GetBasicAcceleratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetBasicAcceleratorResponse
func (client *Client) GetBasicAccelerator(request *GetBasicAcceleratorRequest) (_result *GetBasicAcceleratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBasicAcceleratorResponse{}
	_body, _err := client.GetBasicAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetBasicEndpointWithOptions(request *GetBasicEndpointRequest, runtime *dara.RuntimeOptions) (_result *GetBasicEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetBasicEndpointResponse
func (client *Client) GetBasicEndpoint(request *GetBasicEndpointRequest) (_result *GetBasicEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBasicEndpointResponse{}
	_body, _err := client.GetBasicEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetBasicEndpointGroupWithOptions(request *GetBasicEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *GetBasicEndpointGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetBasicEndpointGroupResponse
func (client *Client) GetBasicEndpointGroup(request *GetBasicEndpointGroupRequest) (_result *GetBasicEndpointGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBasicEndpointGroupResponse{}
	_body, _err := client.GetBasicEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetBasicIpSetWithOptions(request *GetBasicIpSetRequest, runtime *dara.RuntimeOptions) (_result *GetBasicIpSetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetBasicIpSetResponse
func (client *Client) GetBasicIpSet(request *GetBasicIpSetRequest) (_result *GetBasicIpSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBasicIpSetResponse{}
	_body, _err := client.GetBasicIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Obtain the cloud products associated with a Global Accelerator (GA) instance
//
// @param request - GetGlobalAcceleratorResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGlobalAcceleratorResourcesResponse
func (client *Client) GetGlobalAcceleratorResourcesWithOptions(request *GetGlobalAcceleratorResourcesRequest, runtime *dara.RuntimeOptions) (_result *GetGlobalAcceleratorResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtain the cloud products associated with a Global Accelerator (GA) instance
//
// @param request - GetGlobalAcceleratorResourcesRequest
//
// @return GetGlobalAcceleratorResourcesResponse
func (client *Client) GetGlobalAcceleratorResources(request *GetGlobalAcceleratorResourcesRequest) (_result *GetGlobalAcceleratorResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGlobalAcceleratorResourcesResponse{}
	_body, _err := client.GetGlobalAcceleratorResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetHealthStatusWithOptions(request *GetHealthStatusRequest, runtime *dara.RuntimeOptions) (_result *GetHealthStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetHealthStatusResponse
func (client *Client) GetHealthStatus(request *GetHealthStatusRequest) (_result *GetHealthStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHealthStatusResponse{}
	_body, _err := client.GetHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the GetInvalidDomainCount operation to retrieve the total number of invalid domain names.
//
// @param request - GetInvalidDomainCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInvalidDomainCountResponse
func (client *Client) GetInvalidDomainCountWithOptions(request *GetInvalidDomainCountRequest, runtime *dara.RuntimeOptions) (_result *GetInvalidDomainCountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the GetInvalidDomainCount operation to retrieve the total number of invalid domain names.
//
// @param request - GetInvalidDomainCountRequest
//
// @return GetInvalidDomainCountResponse
func (client *Client) GetInvalidDomainCount(request *GetInvalidDomainCountRequest) (_result *GetInvalidDomainCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInvalidDomainCountResponse{}
	_body, _err := client.GetInvalidDomainCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the maximum bandwidth of an acceleration area.
//
// Description:
//
// 本接口用于查询带宽计费方式为**按带宽**的标准型全球加速实例各加速地域的带宽峰值限额，即全球加速实例所绑定基础带宽包的带宽值。
//
// @param request - GetIpsetsBandwidthLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIpsetsBandwidthLimitResponse
func (client *Client) GetIpsetsBandwidthLimitWithOptions(request *GetIpsetsBandwidthLimitRequest, runtime *dara.RuntimeOptions) (_result *GetIpsetsBandwidthLimitResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Description:
//
// 本接口用于查询带宽计费方式为**按带宽**的标准型全球加速实例各加速地域的带宽峰值限额，即全球加速实例所绑定基础带宽包的带宽值。
//
// @param request - GetIpsetsBandwidthLimitRequest
//
// @return GetIpsetsBandwidthLimitResponse
func (client *Client) GetIpsetsBandwidthLimit(request *GetIpsetsBandwidthLimitRequest) (_result *GetIpsetsBandwidthLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIpsetsBandwidthLimitResponse{}
	_body, _err := client.GetIpsetsBandwidthLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a CNAME spare IP address.
//
// @param request - GetSpareIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpareIpResponse
func (client *Client) GetSpareIpWithOptions(request *GetSpareIpRequest, runtime *dara.RuntimeOptions) (_result *GetSpareIpResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a CNAME spare IP address.
//
// @param request - GetSpareIpRequest
//
// @return GetSpareIpResponse
func (client *Client) GetSpareIp(request *GetSpareIpRequest) (_result *GetSpareIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSpareIpResponse{}
	_body, _err := client.GetSpareIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAccelerateAreasWithOptions(request *ListAccelerateAreasRequest, runtime *dara.RuntimeOptions) (_result *ListAccelerateAreasResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAccelerateAreasResponse
func (client *Client) ListAccelerateAreas(request *ListAccelerateAreasRequest) (_result *ListAccelerateAreasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccelerateAreasResponse{}
	_body, _err := client.ListAccelerateAreasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Global Accelerator instances.
//
// @param request - ListAcceleratorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAcceleratorsResponse
func (client *Client) ListAcceleratorsWithOptions(request *ListAcceleratorsRequest, runtime *dara.RuntimeOptions) (_result *ListAcceleratorsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Global Accelerator instances.
//
// @param request - ListAcceleratorsRequest
//
// @return ListAcceleratorsResponse
func (client *Client) ListAccelerators(request *ListAcceleratorsRequest) (_result *ListAcceleratorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAcceleratorsResponse{}
	_body, _err := client.ListAcceleratorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAclsWithOptions(request *ListAclsRequest, runtime *dara.RuntimeOptions) (_result *ListAclsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAclsResponse
func (client *Client) ListAcls(request *ListAclsRequest) (_result *ListAclsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAclsResponse{}
	_body, _err := client.ListAclsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListApplicationMonitorWithOptions(request *ListApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationMonitorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListApplicationMonitorResponse
func (client *Client) ListApplicationMonitor(request *ListApplicationMonitorRequest) (_result *ListApplicationMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationMonitorResponse{}
	_body, _err := client.ListApplicationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListApplicationMonitorDetectResultWithOptions(request *ListApplicationMonitorDetectResultRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationMonitorDetectResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListApplicationMonitorDetectResultResponse
func (client *Client) ListApplicationMonitorDetectResult(request *ListApplicationMonitorDetectResultRequest) (_result *ListApplicationMonitorDetectResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationMonitorDetectResultResponse{}
	_body, _err := client.ListApplicationMonitorDetectResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ListAvailableAccelerateAreas operation to query available acceleration areas.
//
// @param request - ListAvailableAccelerateAreasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableAccelerateAreasResponse
func (client *Client) ListAvailableAccelerateAreasWithOptions(request *ListAvailableAccelerateAreasRequest, runtime *dara.RuntimeOptions) (_result *ListAvailableAccelerateAreasResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ListAvailableAccelerateAreas operation to query available acceleration areas.
//
// @param request - ListAvailableAccelerateAreasRequest
//
// @return ListAvailableAccelerateAreasResponse
func (client *Client) ListAvailableAccelerateAreas(request *ListAvailableAccelerateAreasRequest) (_result *ListAvailableAccelerateAreasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAvailableAccelerateAreasResponse{}
	_body, _err := client.ListAvailableAccelerateAreasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAvailableBusiRegionsWithOptions(request *ListAvailableBusiRegionsRequest, runtime *dara.RuntimeOptions) (_result *ListAvailableBusiRegionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAvailableBusiRegionsResponse
func (client *Client) ListAvailableBusiRegions(request *ListAvailableBusiRegionsRequest) (_result *ListAvailableBusiRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAvailableBusiRegionsResponse{}
	_body, _err := client.ListAvailableBusiRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListBandwidthPackagesWithOptions(request *ListBandwidthPackagesRequest, runtime *dara.RuntimeOptions) (_result *ListBandwidthPackagesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListBandwidthPackagesResponse
func (client *Client) ListBandwidthPackages(request *ListBandwidthPackagesRequest) (_result *ListBandwidthPackagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBandwidthPackagesResponse{}
	_body, _err := client.ListBandwidthPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of bandwidth plans.
//
// Description:
//
// This operation is deprecated. Use [ListBandwidthPackages](https://help.aliyun.com/document_detail/2253239.html) instead.
//
// @param request - ListBandwidthackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBandwidthackagesResponse
func (client *Client) ListBandwidthackagesWithOptions(request *ListBandwidthackagesRequest, runtime *dara.RuntimeOptions) (_result *ListBandwidthackagesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of bandwidth plans.
//
// Description:
//
// This operation is deprecated. Use [ListBandwidthPackages](https://help.aliyun.com/document_detail/2253239.html) instead.
//
// @param request - ListBandwidthackagesRequest
//
// @return ListBandwidthackagesResponse
func (client *Client) ListBandwidthackages(request *ListBandwidthackagesRequest) (_result *ListBandwidthackagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBandwidthackagesResponse{}
	_body, _err := client.ListBandwidthackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListBasicAccelerateIpEndpointRelationsWithOptions(request *ListBasicAccelerateIpEndpointRelationsRequest, runtime *dara.RuntimeOptions) (_result *ListBasicAccelerateIpEndpointRelationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListBasicAccelerateIpEndpointRelationsResponse
func (client *Client) ListBasicAccelerateIpEndpointRelations(request *ListBasicAccelerateIpEndpointRelationsRequest) (_result *ListBasicAccelerateIpEndpointRelationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBasicAccelerateIpEndpointRelationsResponse{}
	_body, _err := client.ListBasicAccelerateIpEndpointRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListBasicAccelerateIpsWithOptions(request *ListBasicAccelerateIpsRequest, runtime *dara.RuntimeOptions) (_result *ListBasicAccelerateIpsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListBasicAccelerateIpsResponse
func (client *Client) ListBasicAccelerateIps(request *ListBasicAccelerateIpsRequest) (_result *ListBasicAccelerateIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBasicAccelerateIpsResponse{}
	_body, _err := client.ListBasicAccelerateIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListBasicAcceleratorsWithOptions(request *ListBasicAcceleratorsRequest, runtime *dara.RuntimeOptions) (_result *ListBasicAcceleratorsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListBasicAcceleratorsResponse
func (client *Client) ListBasicAccelerators(request *ListBasicAcceleratorsRequest) (_result *ListBasicAcceleratorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBasicAcceleratorsResponse{}
	_body, _err := client.ListBasicAcceleratorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListBasicEndpointsWithOptions(request *ListBasicEndpointsRequest, runtime *dara.RuntimeOptions) (_result *ListBasicEndpointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListBasicEndpointsResponse
func (client *Client) ListBasicEndpoints(request *ListBasicEndpointsRequest) (_result *ListBasicEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBasicEndpointsResponse{}
	_body, _err := client.ListBasicEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListBusiRegionsWithOptions(request *ListBusiRegionsRequest, runtime *dara.RuntimeOptions) (_result *ListBusiRegionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListBusiRegionsResponse
func (client *Client) ListBusiRegions(request *ListBusiRegionsRequest) (_result *ListBusiRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBusiRegionsResponse{}
	_body, _err := client.ListBusiRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available acceleration areas and regions by calling the ListCommonAreas operation.
//
// Description:
//
// This operation is used to query the available acceleration areas and regions for the intelligent recommendation and free trial on the Global Accelerator wizard page. You can filter results based on specified conditions.
//
// @param request - ListCommonAreasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCommonAreasResponse
func (client *Client) ListCommonAreasWithOptions(request *ListCommonAreasRequest, runtime *dara.RuntimeOptions) (_result *ListCommonAreasResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available acceleration areas and regions by calling the ListCommonAreas operation.
//
// Description:
//
// This operation is used to query the available acceleration areas and regions for the intelligent recommendation and free trial on the Global Accelerator wizard page. You can filter results based on specified conditions.
//
// @param request - ListCommonAreasRequest
//
// @return ListCommonAreasResponse
func (client *Client) ListCommonAreas(request *ListCommonAreasRequest) (_result *ListCommonAreasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCommonAreasResponse{}
	_body, _err := client.ListCommonAreasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the ListCustomRoutingEndpointGroupDestinations operation to query the destination configurations of an endpoint group for a custom route listener.
//
// @param request - ListCustomRoutingEndpointGroupDestinationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomRoutingEndpointGroupDestinationsResponse
func (client *Client) ListCustomRoutingEndpointGroupDestinationsWithOptions(request *ListCustomRoutingEndpointGroupDestinationsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomRoutingEndpointGroupDestinationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the ListCustomRoutingEndpointGroupDestinations operation to query the destination configurations of an endpoint group for a custom route listener.
//
// @param request - ListCustomRoutingEndpointGroupDestinationsRequest
//
// @return ListCustomRoutingEndpointGroupDestinationsResponse
func (client *Client) ListCustomRoutingEndpointGroupDestinations(request *ListCustomRoutingEndpointGroupDestinationsRequest) (_result *ListCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.ListCustomRoutingEndpointGroupDestinationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListCustomRoutingEndpointGroupsWithOptions(request *ListCustomRoutingEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomRoutingEndpointGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListCustomRoutingEndpointGroupsResponse
func (client *Client) ListCustomRoutingEndpointGroups(request *ListCustomRoutingEndpointGroupsRequest) (_result *ListCustomRoutingEndpointGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomRoutingEndpointGroupsResponse{}
	_body, _err := client.ListCustomRoutingEndpointGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the ListCustomRoutingEndpointTrafficPolicies operation to query the list of endpoint traffic policies for a custom routing type listener.
//
// @param request - ListCustomRoutingEndpointTrafficPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomRoutingEndpointTrafficPoliciesResponse
func (client *Client) ListCustomRoutingEndpointTrafficPoliciesWithOptions(request *ListCustomRoutingEndpointTrafficPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the ListCustomRoutingEndpointTrafficPolicies operation to query the list of endpoint traffic policies for a custom routing type listener.
//
// @param request - ListCustomRoutingEndpointTrafficPoliciesRequest
//
// @return ListCustomRoutingEndpointTrafficPoliciesResponse
func (client *Client) ListCustomRoutingEndpointTrafficPolicies(request *ListCustomRoutingEndpointTrafficPoliciesRequest) (_result *ListCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.ListCustomRoutingEndpointTrafficPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListCustomRoutingEndpointsWithOptions(request *ListCustomRoutingEndpointsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomRoutingEndpointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListCustomRoutingEndpointsResponse
func (client *Client) ListCustomRoutingEndpoints(request *ListCustomRoutingEndpointsRequest) (_result *ListCustomRoutingEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomRoutingEndpointsResponse{}
	_body, _err := client.ListCustomRoutingEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the port mapping table of a custom routing listener.
//
// Description:
//
// After you configure a custom routing listener and an endpoint group, the Global Accelerator (GA) instance generates a port mapping table. This table is based on the listener port range, the protocols and port ranges of the destination endpoint group, and the IP addresses of the endpoints (vSwitches). A custom routing listener uses this port mapping table to deterministically route traffic to specific IP addresses and ports in a vSwitch. This operation queries the generated port mapping table.
//
// @param request - ListCustomRoutingPortMappingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomRoutingPortMappingsResponse
func (client *Client) ListCustomRoutingPortMappingsWithOptions(request *ListCustomRoutingPortMappingsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomRoutingPortMappingsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// After you configure a custom routing listener and an endpoint group, the Global Accelerator (GA) instance generates a port mapping table. This table is based on the listener port range, the protocols and port ranges of the destination endpoint group, and the IP addresses of the endpoints (vSwitches). A custom routing listener uses this port mapping table to deterministically route traffic to specific IP addresses and ports in a vSwitch. This operation queries the generated port mapping table.
//
// @param request - ListCustomRoutingPortMappingsRequest
//
// @return ListCustomRoutingPortMappingsResponse
func (client *Client) ListCustomRoutingPortMappings(request *ListCustomRoutingPortMappingsRequest) (_result *ListCustomRoutingPortMappingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomRoutingPortMappingsResponse{}
	_body, _err := client.ListCustomRoutingPortMappingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the ListCustomRoutingPortMappingsByDestination operation to query the port mappings of a specified backend instance for a custom route listener.
//
// @param request - ListCustomRoutingPortMappingsByDestinationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomRoutingPortMappingsByDestinationResponse
func (client *Client) ListCustomRoutingPortMappingsByDestinationWithOptions(request *ListCustomRoutingPortMappingsByDestinationRequest, runtime *dara.RuntimeOptions) (_result *ListCustomRoutingPortMappingsByDestinationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the ListCustomRoutingPortMappingsByDestination operation to query the port mappings of a specified backend instance for a custom route listener.
//
// @param request - ListCustomRoutingPortMappingsByDestinationRequest
//
// @return ListCustomRoutingPortMappingsByDestinationResponse
func (client *Client) ListCustomRoutingPortMappingsByDestination(request *ListCustomRoutingPortMappingsByDestinationRequest) (_result *ListCustomRoutingPortMappingsByDestinationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomRoutingPortMappingsByDestinationResponse{}
	_body, _err := client.ListCustomRoutingPortMappingsByDestinationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDomainsWithOptions(request *ListDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDomainsResponse
func (client *Client) ListDomains(request *ListDomainsRequest) (_result *ListDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDomainsResponse{}
	_body, _err := client.ListDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListEndpointGroupIpAddressCidrBlocksWithOptions(request *ListEndpointGroupIpAddressCidrBlocksRequest, runtime *dara.RuntimeOptions) (_result *ListEndpointGroupIpAddressCidrBlocksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListEndpointGroupIpAddressCidrBlocksResponse
func (client *Client) ListEndpointGroupIpAddressCidrBlocks(request *ListEndpointGroupIpAddressCidrBlocksRequest) (_result *ListEndpointGroupIpAddressCidrBlocksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEndpointGroupIpAddressCidrBlocksResponse{}
	_body, _err := client.ListEndpointGroupIpAddressCidrBlocksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists endpoint groups.
//
// @param request - ListEndpointGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEndpointGroupsResponse
func (client *Client) ListEndpointGroupsWithOptions(request *ListEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListEndpointGroupsResponse, _err error) {
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

	if !dara.IsNil(request.EndpointGroupRegion) {
		query["EndpointGroupRegion"] = request.EndpointGroupRegion
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists endpoint groups.
//
// @param request - ListEndpointGroupsRequest
//
// @return ListEndpointGroupsResponse
func (client *Client) ListEndpointGroups(request *ListEndpointGroupsRequest) (_result *ListEndpointGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEndpointGroupsResponse{}
	_body, _err := client.ListEndpointGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries forwarding rules.
//
// Description:
//
// > Queries only custom forwarding rules. The default forwarding rule is not included in the results.
//
// @param request - ListForwardingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListForwardingRulesResponse
func (client *Client) ListForwardingRulesWithOptions(request *ListForwardingRulesRequest, runtime *dara.RuntimeOptions) (_result *ListForwardingRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// > Queries only custom forwarding rules. The default forwarding rule is not included in the results.
//
// @param request - ListForwardingRulesRequest
//
// @return ListForwardingRulesResponse
func (client *Client) ListForwardingRules(request *ListForwardingRulesRequest) (_result *ListForwardingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListForwardingRulesResponse{}
	_body, _err := client.ListForwardingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListIpSetsWithOptions(request *ListIpSetsRequest, runtime *dara.RuntimeOptions) (_result *ListIpSetsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListIpSetsResponse
func (client *Client) ListIpSets(request *ListIpSetsRequest) (_result *ListIpSetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIpSetsResponse{}
	_body, _err := client.ListIpSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListIspTypesWithOptions(request *ListIspTypesRequest, runtime *dara.RuntimeOptions) (_result *ListIspTypesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListIspTypesResponse
func (client *Client) ListIspTypes(request *ListIspTypesRequest) (_result *ListIspTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIspTypesResponse{}
	_body, _err := client.ListIspTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListListenerCertificatesWithOptions(request *ListListenerCertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListListenerCertificatesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListListenerCertificatesResponse
func (client *Client) ListListenerCertificates(request *ListListenerCertificatesRequest) (_result *ListListenerCertificatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListListenerCertificatesResponse{}
	_body, _err := client.ListListenerCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of listeners.
//
// Description:
//
// This operation retrieves the listeners of a Global Accelerator instance. The response includes the routing type, status, creation timestamp, and port details for each listener.
//
// @param request - ListListenersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListListenersResponse
func (client *Client) ListListenersWithOptions(request *ListListenersRequest, runtime *dara.RuntimeOptions) (_result *ListListenersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of listeners.
//
// Description:
//
// This operation retrieves the listeners of a Global Accelerator instance. The response includes the routing type, status, creation timestamp, and port details for each listener.
//
// @param request - ListListenersRequest
//
// @return ListListenersResponse
func (client *Client) ListListeners(request *ListListenersRequest) (_result *ListListenersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListListenersResponse{}
	_body, _err := client.ListListenersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListSpareIpsWithOptions(request *ListSpareIpsRequest, runtime *dara.RuntimeOptions) (_result *ListSpareIpsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListSpareIpsResponse
func (client *Client) ListSpareIps(request *ListSpareIpsRequest) (_result *ListSpareIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSpareIpsResponse{}
	_body, _err := client.ListSpareIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListSystemSecurityPoliciesWithOptions(request *ListSystemSecurityPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListSystemSecurityPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListSystemSecurityPoliciesResponse
func (client *Client) ListSystemSecurityPolicies(request *ListSystemSecurityPoliciesRequest) (_result *ListSystemSecurityPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSystemSecurityPoliciesResponse{}
	_body, _err := client.ListSystemSecurityPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Activates the pay-as-you-go Global Accelerator (GA) service. If you want to use pay-as-you-go GA instances, you must activate the pay-as-you-go GA service first.
//
// @param request - OpenAcceleratorServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenAcceleratorServiceResponse
func (client *Client) OpenAcceleratorServiceWithOptions(request *OpenAcceleratorServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenAcceleratorServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OpenAcceleratorServiceResponse
func (client *Client) OpenAcceleratorService(request *OpenAcceleratorServiceRequest) (_result *OpenAcceleratorServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenAcceleratorServiceResponse{}
	_body, _err := client.OpenAcceleratorServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the approval status of cross-border permissions for an Alibaba Cloud account (main account).
//
// @param request - QueryCrossBorderApprovalStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCrossBorderApprovalStatusResponse
func (client *Client) QueryCrossBorderApprovalStatusWithOptions(request *QueryCrossBorderApprovalStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryCrossBorderApprovalStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the approval status of cross-border permissions for an Alibaba Cloud account (main account).
//
// @param request - QueryCrossBorderApprovalStatusRequest
//
// @return QueryCrossBorderApprovalStatusResponse
func (client *Client) QueryCrossBorderApprovalStatus(request *QueryCrossBorderApprovalStatusRequest) (_result *QueryCrossBorderApprovalStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCrossBorderApprovalStatusResponse{}
	_body, _err := client.QueryCrossBorderApprovalStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveEntriesFromAclWithOptions(request *RemoveEntriesFromAclRequest, runtime *dara.RuntimeOptions) (_result *RemoveEntriesFromAclResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RemoveEntriesFromAclResponse
func (client *Client) RemoveEntriesFromAcl(request *RemoveEntriesFromAclRequest) (_result *RemoveEntriesFromAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveEntriesFromAclResponse{}
	_body, _err := client.RemoveEntriesFromAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReplaceBandwidthPackageWithOptions(request *ReplaceBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *ReplaceBandwidthPackageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReplaceBandwidthPackageResponse
func (client *Client) ReplaceBandwidthPackage(request *ReplaceBandwidthPackageRequest) (_result *ReplaceBandwidthPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReplaceBandwidthPackageResponse{}
	_body, _err := client.ReplaceBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Removes tags from Global Accelerator (GA) resources.
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Use the UpdateAccelerator operation to modify a Global Accelerator instance.
//
// Description:
//
// - **UpdateAccelerator*	- is an asynchronous operation. After you send a request, the system returns a request ID and performs the modification in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the state of a Global Accelerator (GA) instance:
//
//   - If the GA instance is in the **configuring*	- state, the instance is being modified. In this state, you can only perform query operations.
//
//   - If the GA instance is in the **active*	- state, the modification is complete.
//
// - You cannot call the **UpdateAccelerator*	- operation concurrently on the same GA instance.
//
// @param request - UpdateAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAcceleratorResponse
func (client *Client) UpdateAcceleratorWithOptions(request *UpdateAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *UpdateAcceleratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use the UpdateAccelerator operation to modify a Global Accelerator instance.
//
// Description:
//
// - **UpdateAccelerator*	- is an asynchronous operation. After you send a request, the system returns a request ID and performs the modification in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the state of a Global Accelerator (GA) instance:
//
//   - If the GA instance is in the **configuring*	- state, the instance is being modified. In this state, you can only perform query operations.
//
//   - If the GA instance is in the **active*	- state, the modification is complete.
//
// - You cannot call the **UpdateAccelerator*	- operation concurrently on the same GA instance.
//
// @param request - UpdateAcceleratorRequest
//
// @return UpdateAcceleratorResponse
func (client *Client) UpdateAccelerator(request *UpdateAcceleratorRequest) (_result *UpdateAcceleratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAcceleratorResponse{}
	_body, _err := client.UpdateAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the UpdateAcceleratorAutoRenewAttribute operation to modify the auto-renewal property of a Global Accelerator instance.
//
// Description:
//
// The **UpdateAcceleratorAutoRenewAttribute*	- operation does not support concurrent modifications to the auto-renewal property of the same Global Accelerator instance.
//
// @param request - UpdateAcceleratorAutoRenewAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAcceleratorAutoRenewAttributeResponse
func (client *Client) UpdateAcceleratorAutoRenewAttributeWithOptions(request *UpdateAcceleratorAutoRenewAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateAcceleratorAutoRenewAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the UpdateAcceleratorAutoRenewAttribute operation to modify the auto-renewal property of a Global Accelerator instance.
//
// Description:
//
// The **UpdateAcceleratorAutoRenewAttribute*	- operation does not support concurrent modifications to the auto-renewal property of the same Global Accelerator instance.
//
// @param request - UpdateAcceleratorAutoRenewAttributeRequest
//
// @return UpdateAcceleratorAutoRenewAttributeResponse
func (client *Client) UpdateAcceleratorAutoRenewAttribute(request *UpdateAcceleratorAutoRenewAttributeRequest) (_result *UpdateAcceleratorAutoRenewAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAcceleratorAutoRenewAttributeResponse{}
	_body, _err := client.UpdateAcceleratorAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateAcceleratorConfirmWithOptions(request *UpdateAcceleratorConfirmRequest, runtime *dara.RuntimeOptions) (_result *UpdateAcceleratorConfirmResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateAcceleratorConfirmResponse
func (client *Client) UpdateAcceleratorConfirm(request *UpdateAcceleratorConfirmRequest) (_result *UpdateAcceleratorConfirmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAcceleratorConfirmResponse{}
	_body, _err := client.UpdateAcceleratorConfirmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the type of transmission network for a Global Accelerator (GA) instance.
//
// Description:
//
// Changes the type of transmission network for a **standard*	- GA instance whose bandwidth metering method is **pay-by-data-transfer**. Before you call this operation, make sure that the following requirements are met:
//
// - Cloud Data Transfer (CDT) is activated. When you call the [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html) operation and set **BandwidthBillingType*	- to **CDT*	- to create a **standard*	- GA instance whose bandwidth metering method is **pay-by-data-transfer**, CDT is automatically activated. The data transfer fees are managed by CDT.
//
// - If you want to set **CrossBorderMode*	- to **private**, which specifies cross-border Express Connect circuit as the type of transmission network, make sure that real-name verification is complete for your enterprise account. For more information, see [Real-name verification](https://help.aliyun.com/document_detail/52595.html).
//
// @param request - UpdateAcceleratorCrossBorderModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAcceleratorCrossBorderModeResponse
func (client *Client) UpdateAcceleratorCrossBorderModeWithOptions(request *UpdateAcceleratorCrossBorderModeRequest, runtime *dara.RuntimeOptions) (_result *UpdateAcceleratorCrossBorderModeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Changes the type of transmission network for a **standard*	- GA instance whose bandwidth metering method is **pay-by-data-transfer**. Before you call this operation, make sure that the following requirements are met:
//
// - Cloud Data Transfer (CDT) is activated. When you call the [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html) operation and set **BandwidthBillingType*	- to **CDT*	- to create a **standard*	- GA instance whose bandwidth metering method is **pay-by-data-transfer**, CDT is automatically activated. The data transfer fees are managed by CDT.
//
// - If you want to set **CrossBorderMode*	- to **private**, which specifies cross-border Express Connect circuit as the type of transmission network, make sure that real-name verification is complete for your enterprise account. For more information, see [Real-name verification](https://help.aliyun.com/document_detail/52595.html).
//
// @param request - UpdateAcceleratorCrossBorderModeRequest
//
// @return UpdateAcceleratorCrossBorderModeResponse
func (client *Client) UpdateAcceleratorCrossBorderMode(request *UpdateAcceleratorCrossBorderModeRequest) (_result *UpdateAcceleratorCrossBorderModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAcceleratorCrossBorderModeResponse{}
	_body, _err := client.UpdateAcceleratorCrossBorderModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables cross-border data transmission for a Global Accelerator (GA) instance.
//
// Description:
//
// Enables or disables cross-border data transmission for basic or standard GA instances that use Cloud Data Transfer (CDT) to bill data transfers.
//
// @param request - UpdateAcceleratorCrossBorderStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAcceleratorCrossBorderStatusResponse
func (client *Client) UpdateAcceleratorCrossBorderStatusWithOptions(request *UpdateAcceleratorCrossBorderStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateAcceleratorCrossBorderStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Enables or disables cross-border data transmission for basic or standard GA instances that use Cloud Data Transfer (CDT) to bill data transfers.
//
// @param request - UpdateAcceleratorCrossBorderStatusRequest
//
// @return UpdateAcceleratorCrossBorderStatusResponse
func (client *Client) UpdateAcceleratorCrossBorderStatus(request *UpdateAcceleratorCrossBorderStatusRequest) (_result *UpdateAcceleratorCrossBorderStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAcceleratorCrossBorderStatusResponse{}
	_body, _err := client.UpdateAcceleratorCrossBorderStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateAclAttributeWithOptions(request *UpdateAclAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateAclAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateAclAttributeResponse
func (client *Client) UpdateAclAttribute(request *UpdateAclAttributeRequest) (_result *UpdateAclAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAclAttributeResponse{}
	_body, _err := client.UpdateAclAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateAdditionalCertificateWithListenerWithOptions(request *UpdateAdditionalCertificateWithListenerRequest, runtime *dara.RuntimeOptions) (_result *UpdateAdditionalCertificateWithListenerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateAdditionalCertificateWithListenerResponse
func (client *Client) UpdateAdditionalCertificateWithListener(request *UpdateAdditionalCertificateWithListenerRequest) (_result *UpdateAdditionalCertificateWithListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAdditionalCertificateWithListenerResponse{}
	_body, _err := client.UpdateAdditionalCertificateWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of an origin probing task by calling the UpdateApplicationMonitor operation.
//
// Description:
//
// *UpdateApplicationMonitor*	- is an asynchronous operation. After you call this operation, the system returns a request ID, but the origin probing task is not yet modified. The modification task continues to run in the background. You can call [DescribeApplicationMonitor](https://help.aliyun.com/document_detail/408463.html) or [ListApplicationMonitor](https://help.aliyun.com/document_detail/408462.html) to check whether the origin probing task configuration has been modified:
//
// - If the modified parameter values have not changed, the origin probing task is still being modified. In this case, you can only perform query operations and cannot perform other operations.
//
// - If the modified parameter values have changed, the origin probing task has been modified.
//
// @param request - UpdateApplicationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationMonitorResponse
func (client *Client) UpdateApplicationMonitorWithOptions(request *UpdateApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationMonitorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of an origin probing task by calling the UpdateApplicationMonitor operation.
//
// Description:
//
// *UpdateApplicationMonitor*	- is an asynchronous operation. After you call this operation, the system returns a request ID, but the origin probing task is not yet modified. The modification task continues to run in the background. You can call [DescribeApplicationMonitor](https://help.aliyun.com/document_detail/408463.html) or [ListApplicationMonitor](https://help.aliyun.com/document_detail/408462.html) to check whether the origin probing task configuration has been modified:
//
// - If the modified parameter values have not changed, the origin probing task is still being modified. In this case, you can only perform query operations and cannot perform other operations.
//
// - If the modified parameter values have changed, the origin probing task has been modified.
//
// @param request - UpdateApplicationMonitorRequest
//
// @return UpdateApplicationMonitorResponse
func (client *Client) UpdateApplicationMonitor(request *UpdateApplicationMonitorRequest) (_result *UpdateApplicationMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApplicationMonitorResponse{}
	_body, _err := client.UpdateApplicationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateBandwidthPackagaAutoRenewAttributeWithOptions(request *UpdateBandwidthPackagaAutoRenewAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateBandwidthPackagaAutoRenewAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateBandwidthPackagaAutoRenewAttributeResponse
func (client *Client) UpdateBandwidthPackagaAutoRenewAttribute(request *UpdateBandwidthPackagaAutoRenewAttributeRequest) (_result *UpdateBandwidthPackagaAutoRenewAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBandwidthPackagaAutoRenewAttributeResponse{}
	_body, _err := client.UpdateBandwidthPackagaAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the UpdateBandwidthPackage operation to modify the configuration of a bandwidth plan.
//
// Description:
//
// Before you call this operation, note the following:
//
// - If you do not change the bandwidth value, **UpdateBandwidthPackage*	- is a synchronous operation. The modification takes effect immediately.
//
// - If you change the bandwidth value of a bandwidth plan that is not associated with a Global Accelerator (GA) instance, **UpdateBandwidthPackage*	- is an asynchronous operation. The system returns a request ID, but the modification is not complete. The system performs the task in the background. You can call the [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) operation to query whether the configuration of the bandwidth plan is modified:
//
//   - If the values of the parameters that you want to modify remain unchanged, the bandwidth plan is being modified. In this case, you can only perform query operations.
//
//   - If the values of the parameters that you want to modify have changed, the modification is complete.
//
// - If you change the bandwidth value of a bandwidth plan that is associated with a GA instance, **UpdateBandwidthPackage*	- is an asynchronous operation. The system returns a request ID, but the modification is not complete. The system performs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation and query the status of the GA instance to check whether the bandwidth plan is modified:
//
//   - If the GA instance is in the **configuring*	- state, the bandwidth plan that is associated with the GA instance is being modified. In this case, you can only perform query operations.
//
//   - If the GA instance is in the **active*	- state, the bandwidth plan that is associated with the GA instance has been modified.
//
// - You cannot repeatedly call the **UpdateBandwidthPackage*	- operation to modify the configuration of the same bandwidth plan.
//
// @param request - UpdateBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBandwidthPackageResponse
func (client *Client) UpdateBandwidthPackageWithOptions(request *UpdateBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *UpdateBandwidthPackageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the UpdateBandwidthPackage operation to modify the configuration of a bandwidth plan.
//
// Description:
//
// Before you call this operation, note the following:
//
// - If you do not change the bandwidth value, **UpdateBandwidthPackage*	- is a synchronous operation. The modification takes effect immediately.
//
// - If you change the bandwidth value of a bandwidth plan that is not associated with a Global Accelerator (GA) instance, **UpdateBandwidthPackage*	- is an asynchronous operation. The system returns a request ID, but the modification is not complete. The system performs the task in the background. You can call the [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) operation to query whether the configuration of the bandwidth plan is modified:
//
//   - If the values of the parameters that you want to modify remain unchanged, the bandwidth plan is being modified. In this case, you can only perform query operations.
//
//   - If the values of the parameters that you want to modify have changed, the modification is complete.
//
// - If you change the bandwidth value of a bandwidth plan that is associated with a GA instance, **UpdateBandwidthPackage*	- is an asynchronous operation. The system returns a request ID, but the modification is not complete. The system performs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation and query the status of the GA instance to check whether the bandwidth plan is modified:
//
//   - If the GA instance is in the **configuring*	- state, the bandwidth plan that is associated with the GA instance is being modified. In this case, you can only perform query operations.
//
//   - If the GA instance is in the **active*	- state, the bandwidth plan that is associated with the GA instance has been modified.
//
// - You cannot repeatedly call the **UpdateBandwidthPackage*	- operation to modify the configuration of the same bandwidth plan.
//
// @param request - UpdateBandwidthPackageRequest
//
// @return UpdateBandwidthPackageResponse
func (client *Client) UpdateBandwidthPackage(request *UpdateBandwidthPackageRequest) (_result *UpdateBandwidthPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBandwidthPackageResponse{}
	_body, _err := client.UpdateBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateBasicAcceleratorWithOptions(request *UpdateBasicAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *UpdateBasicAcceleratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateBasicAcceleratorResponse
func (client *Client) UpdateBasicAccelerator(request *UpdateBasicAcceleratorRequest) (_result *UpdateBasicAcceleratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBasicAcceleratorResponse{}
	_body, _err := client.UpdateBasicAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateBasicEndpointWithOptions(request *UpdateBasicEndpointRequest, runtime *dara.RuntimeOptions) (_result *UpdateBasicEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateBasicEndpointResponse
func (client *Client) UpdateBasicEndpoint(request *UpdateBasicEndpointRequest) (_result *UpdateBasicEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBasicEndpointResponse{}
	_body, _err := client.UpdateBasicEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of an endpoint group that is associated with a basic Global Accelerator (GA) instance.
//
// Description:
//
// - **UpdateBasicEndpointGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. The system modifies the configurations of an endpoint group that is associated with a basic GA instance by deleting the endpoint group and creating a new endpoint group. You can call the [GetBasicAccelerator](https://help.aliyun.com/document_detail/353188.html) operation to query the status of the task.
//
//   - If the basic GA instance is in the **configuring*	- state, the configurations of the endpoint group are being modified. In this case, you can perform only query operations.
//
//   - If the basic GA instance is in the **active*	- state, the configurations of the endpoint group are modified.
//
// - The **UpdateBasicEndpointGroup*	- operation cannot be repeatedly called for the same basic GA instance within a specific period of time.
//
// @param request - UpdateBasicEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBasicEndpointGroupResponse
func (client *Client) UpdateBasicEndpointGroupWithOptions(request *UpdateBasicEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateBasicEndpointGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of an endpoint group that is associated with a basic Global Accelerator (GA) instance.
//
// Description:
//
// - **UpdateBasicEndpointGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. The system modifies the configurations of an endpoint group that is associated with a basic GA instance by deleting the endpoint group and creating a new endpoint group. You can call the [GetBasicAccelerator](https://help.aliyun.com/document_detail/353188.html) operation to query the status of the task.
//
//   - If the basic GA instance is in the **configuring*	- state, the configurations of the endpoint group are being modified. In this case, you can perform only query operations.
//
//   - If the basic GA instance is in the **active*	- state, the configurations of the endpoint group are modified.
//
// - The **UpdateBasicEndpointGroup*	- operation cannot be repeatedly called for the same basic GA instance within a specific period of time.
//
// @param request - UpdateBasicEndpointGroupRequest
//
// @return UpdateBasicEndpointGroupResponse
func (client *Client) UpdateBasicEndpointGroup(request *UpdateBasicEndpointGroupRequest) (_result *UpdateBasicEndpointGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBasicEndpointGroupResponse{}
	_body, _err := client.UpdateBasicEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateBasicIpSetWithOptions(request *UpdateBasicIpSetRequest, runtime *dara.RuntimeOptions) (_result *UpdateBasicIpSetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateBasicIpSetResponse
func (client *Client) UpdateBasicIpSet(request *UpdateBasicIpSetRequest) (_result *UpdateBasicIpSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBasicIpSetResponse{}
	_body, _err := client.UpdateBasicIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the UpdateCustomRoutingEndpointGroupAttribute operation to modify the name and description of an endpoint group associated with a custom routing listener.
//
// @param request - UpdateCustomRoutingEndpointGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomRoutingEndpointGroupAttributeResponse
func (client *Client) UpdateCustomRoutingEndpointGroupAttributeWithOptions(request *UpdateCustomRoutingEndpointGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomRoutingEndpointGroupAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the UpdateCustomRoutingEndpointGroupAttribute operation to modify the name and description of an endpoint group associated with a custom routing listener.
//
// @param request - UpdateCustomRoutingEndpointGroupAttributeRequest
//
// @return UpdateCustomRoutingEndpointGroupAttributeResponse
func (client *Client) UpdateCustomRoutingEndpointGroupAttribute(request *UpdateCustomRoutingEndpointGroupAttributeRequest) (_result *UpdateCustomRoutingEndpointGroupAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomRoutingEndpointGroupAttributeResponse{}
	_body, _err := client.UpdateCustomRoutingEndpointGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateCustomRoutingEndpointGroupDestinationsWithOptions(request *UpdateCustomRoutingEndpointGroupDestinationsRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomRoutingEndpointGroupDestinationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateCustomRoutingEndpointGroupDestinationsResponse
func (client *Client) UpdateCustomRoutingEndpointGroupDestinations(request *UpdateCustomRoutingEndpointGroupDestinationsRequest) (_result *UpdateCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.UpdateCustomRoutingEndpointGroupDestinationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateCustomRoutingEndpointTrafficPoliciesWithOptions(request *UpdateCustomRoutingEndpointTrafficPoliciesRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateCustomRoutingEndpointTrafficPoliciesResponse
func (client *Client) UpdateCustomRoutingEndpointTrafficPolicies(request *UpdateCustomRoutingEndpointTrafficPoliciesRequest) (_result *UpdateCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.UpdateCustomRoutingEndpointTrafficPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateCustomRoutingEndpointsWithOptions(request *UpdateCustomRoutingEndpointsRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomRoutingEndpointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateCustomRoutingEndpointsResponse
func (client *Client) UpdateCustomRoutingEndpoints(request *UpdateCustomRoutingEndpointsRequest) (_result *UpdateCustomRoutingEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomRoutingEndpointsResponse{}
	_body, _err := client.UpdateCustomRoutingEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the UpdateDomain operation to update a domain name.
//
// Description:
//
// This operation is used to update an accelerated domain name. If the new accelerated domain name is deployed in the Chinese mainland, the domain name must have obtained an ICP filing.
//
// The **UpdateDomain*	- operation does not support concurrent updates of accelerated domain names within the same account.
//
// @param request - UpdateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomainWithOptions(request *UpdateDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the UpdateDomain operation to update a domain name.
//
// Description:
//
// This operation is used to update an accelerated domain name. If the new accelerated domain name is deployed in the Chinese mainland, the domain name must have obtained an ICP filing.
//
// The **UpdateDomain*	- operation does not support concurrent updates of accelerated domain names within the same account.
//
// @param request - UpdateDomainRequest
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomain(request *UpdateDomainRequest) (_result *UpdateDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDomainResponse{}
	_body, _err := client.UpdateDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries and updates the ICP filing status of an accelerated domain name.
//
// Description:
//
// This operation queries the latest ICP filing status of an accelerated domain name and updates the status accordingly.
//
// The **UpdateDomainState*	- operation holds an exclusive lock on the GA instance. While the operation is in progress, you cannot call the same operation with the same Alibaba Cloud account.
//
// @param request - UpdateDomainStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainStateResponse
func (client *Client) UpdateDomainStateWithOptions(request *UpdateDomainStateRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainStateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries and updates the ICP filing status of an accelerated domain name.
//
// Description:
//
// This operation queries the latest ICP filing status of an accelerated domain name and updates the status accordingly.
//
// The **UpdateDomainState*	- operation holds an exclusive lock on the GA instance. While the operation is in progress, you cannot call the same operation with the same Alibaba Cloud account.
//
// @param request - UpdateDomainStateRequest
//
// @return UpdateDomainStateResponse
func (client *Client) UpdateDomainState(request *UpdateDomainStateRequest) (_result *UpdateDomainStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDomainStateResponse{}
	_body, _err := client.UpdateDomainStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of an endpoint group.
//
// Description:
//
// - The **UpdateEndpointGroup*	- API is an asynchronous API. After you call this API, the system returns a request ID and starts the update in the background; the configuration is not modified immediately. Call [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) to check the status of the endpoint group:
//
//   - If an endpoint group is in the **updating*	- status, its configuration is being modified, and you can only perform queries.
//
//   - If an endpoint group is in the **active*	- status, the update is complete.
//
// - The **UpdateEndpointGroup*	- API does not support concurrent updates to endpoint groups in the same Global Accelerator (GA) instance.
//
// @param request - UpdateEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEndpointGroupResponse
func (client *Client) UpdateEndpointGroupWithOptions(request *UpdateEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateEndpointGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of an endpoint group.
//
// Description:
//
// - The **UpdateEndpointGroup*	- API is an asynchronous API. After you call this API, the system returns a request ID and starts the update in the background; the configuration is not modified immediately. Call [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) to check the status of the endpoint group:
//
//   - If an endpoint group is in the **updating*	- status, its configuration is being modified, and you can only perform queries.
//
//   - If an endpoint group is in the **active*	- status, the update is complete.
//
// - The **UpdateEndpointGroup*	- API does not support concurrent updates to endpoint groups in the same Global Accelerator (GA) instance.
//
// @param request - UpdateEndpointGroupRequest
//
// @return UpdateEndpointGroupResponse
func (client *Client) UpdateEndpointGroup(request *UpdateEndpointGroupRequest) (_result *UpdateEndpointGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEndpointGroupResponse{}
	_body, _err := client.UpdateEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateEndpointGroupAttributeWithOptions(request *UpdateEndpointGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateEndpointGroupAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateEndpointGroupAttributeResponse
func (client *Client) UpdateEndpointGroupAttribute(request *UpdateEndpointGroupAttributeRequest) (_result *UpdateEndpointGroupAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEndpointGroupAttributeResponse{}
	_body, _err := client.UpdateEndpointGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies endpoint groups for a listener in a batch.
//
// Description:
//
// ### Usage notes
//
// - **UpdateEndpointGroups*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation continues to run in the background. You can call the or [](t2323644.xdita#)operation to query the state of an endpoint group.
//
//   - If an endpoint group is in the **updating*	- state, its configuration is being modified. In this state, you can only perform query operations.
//
//   - If an endpoint group is in the **active*	- state, its configuration has been modified.
//
// - You cannot concurrently call the **UpdateEndpointGroups*	- operation to modify the configurations of endpoint groups that belong to the same Global Accelerator (GA) instance.
//
// @param request - UpdateEndpointGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEndpointGroupsResponse
func (client *Client) UpdateEndpointGroupsWithOptions(request *UpdateEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *UpdateEndpointGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies endpoint groups for a listener in a batch.
//
// Description:
//
// ### Usage notes
//
// - **UpdateEndpointGroups*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation continues to run in the background. You can call the or [](t2323644.xdita#)operation to query the state of an endpoint group.
//
//   - If an endpoint group is in the **updating*	- state, its configuration is being modified. In this state, you can only perform query operations.
//
//   - If an endpoint group is in the **active*	- state, its configuration has been modified.
//
// - You cannot concurrently call the **UpdateEndpointGroups*	- operation to modify the configurations of endpoint groups that belong to the same Global Accelerator (GA) instance.
//
// @param request - UpdateEndpointGroupsRequest
//
// @return UpdateEndpointGroupsResponse
func (client *Client) UpdateEndpointGroups(request *UpdateEndpointGroupsRequest) (_result *UpdateEndpointGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEndpointGroupsResponse{}
	_body, _err := client.UpdateEndpointGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// To update forwarding rules, call the UpdateForwardingRules API.
//
// Description:
//
// - **UpdateForwardingRules*	- is an asynchronous API. A call to this API returns a request ID and runs the update in the background. You can call [ListForwardingRules](https://help.aliyun.com/document_detail/205817.html) to query the status of the forwarding rule:
//
//   - A status of **configuring*	- indicates that the forwarding rule is being updated. During this process, you can only perform query operations.
//
//   - A status of **active*	- indicates that the update is complete.
//
// - You cannot use **UpdateForwardingRules*	- to concurrently update forwarding rules within the same Global Accelerator instance.
//
// @param request - UpdateForwardingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateForwardingRulesResponse
func (client *Client) UpdateForwardingRulesWithOptions(request *UpdateForwardingRulesRequest, runtime *dara.RuntimeOptions) (_result *UpdateForwardingRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// To update forwarding rules, call the UpdateForwardingRules API.
//
// Description:
//
// - **UpdateForwardingRules*	- is an asynchronous API. A call to this API returns a request ID and runs the update in the background. You can call [ListForwardingRules](https://help.aliyun.com/document_detail/205817.html) to query the status of the forwarding rule:
//
//   - A status of **configuring*	- indicates that the forwarding rule is being updated. During this process, you can only perform query operations.
//
//   - A status of **active*	- indicates that the update is complete.
//
// - You cannot use **UpdateForwardingRules*	- to concurrently update forwarding rules within the same Global Accelerator instance.
//
// @param request - UpdateForwardingRulesRequest
//
// @return UpdateForwardingRulesResponse
func (client *Client) UpdateForwardingRules(request *UpdateForwardingRulesRequest) (_result *UpdateForwardingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateForwardingRulesResponse{}
	_body, _err := client.UpdateForwardingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateIpSetWithOptions(request *UpdateIpSetRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpSetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateIpSetResponse
func (client *Client) UpdateIpSet(request *UpdateIpSetRequest) (_result *UpdateIpSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIpSetResponse{}
	_body, _err := client.UpdateIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateIpSetsWithOptions(request *UpdateIpSetsRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpSetsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateIpSetsResponse
func (client *Client) UpdateIpSets(request *UpdateIpSetsRequest) (_result *UpdateIpSetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIpSetsResponse{}
	_body, _err := client.UpdateIpSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of a listener for a Global Accelerator (GA) instance.
//
// Description:
//
// Modifies the protocol, ports, and other configurations of a listener to meet your business requirements.
//
// When you call this operation, take note of the following items:
//
// - **UpdateListener*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeListener](https://help.aliyun.com/document_detail/153254.html) operation to query the status of a listener.
//
//   - If the listener is in the **updating*	- state, it indicates that its configurations are being modified. In this case, you can perform only query operations.
//
//   - If the listener is in the **active*	- state, it indicates that its configurations are modified.
//
// - The **UpdateListener*	- operation cannot be repeatedly called to modify listener configurations for the same GA instance within a specific period of time.
//
// @param request - UpdateListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateListenerResponse
func (client *Client) UpdateListenerWithOptions(request *UpdateListenerRequest, runtime *dara.RuntimeOptions) (_result *UpdateListenerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of a listener for a Global Accelerator (GA) instance.
//
// Description:
//
// Modifies the protocol, ports, and other configurations of a listener to meet your business requirements.
//
// When you call this operation, take note of the following items:
//
// - **UpdateListener*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeListener](https://help.aliyun.com/document_detail/153254.html) operation to query the status of a listener.
//
//   - If the listener is in the **updating*	- state, it indicates that its configurations are being modified. In this case, you can perform only query operations.
//
//   - If the listener is in the **active*	- state, it indicates that its configurations are modified.
//
// - The **UpdateListener*	- operation cannot be repeatedly called to modify listener configurations for the same GA instance within a specific period of time.
//
// @param request - UpdateListenerRequest
//
// @return UpdateListenerResponse
func (client *Client) UpdateListener(request *UpdateListenerRequest) (_result *UpdateListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateListenerResponse{}
	_body, _err := client.UpdateListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Simple Log Service log configuration
//
// @param request - UpdateLogStoreConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogStoreConfigResponse
func (client *Client) UpdateLogStoreConfigWithOptions(request *UpdateLogStoreConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLogStoreConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Simple Log Service log configuration
//
// @param request - UpdateLogStoreConfigRequest
//
// @return UpdateLogStoreConfigResponse
func (client *Client) UpdateLogStoreConfig(request *UpdateLogStoreConfigRequest) (_result *UpdateLogStoreConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLogStoreConfigResponse{}
	_body, _err := client.UpdateLogStoreConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateServiceManagedControlWithOptions(request *UpdateServiceManagedControlRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceManagedControlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateServiceManagedControlResponse
func (client *Client) UpdateServiceManagedControl(request *UpdateServiceManagedControlRequest) (_result *UpdateServiceManagedControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateServiceManagedControlResponse{}
	_body, _err := client.UpdateServiceManagedControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
