// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Binds a custom domain name to a private gateway.
//
// @param tmpReq - AttachGatewayDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachGatewayDomainResponse
func (client *Client) AttachGatewayDomainWithContext(ctx context.Context, ClusterId *string, GatewayId *string, tmpReq *AttachGatewayDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AttachGatewayDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AttachGatewayDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomDomain) {
		request.CustomDomainShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomDomain, dara.String("CustomDomain"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomDomainShrink) {
		query["CustomDomain"] = request.CustomDomainShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachGatewayDomain"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/domain/attach"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachGatewayDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clones a service.
//
// @param tmpReq - CloneServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneServiceResponse
func (client *Client) CloneServiceWithContext(ctx context.Context, ClusterId *string, ServiceName *string, tmpReq *CloneServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CloneServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Labels) {
		request.LabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Labels, dara.String("Labels"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelsShrink) {
		query["Labels"] = request.LabelsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Commits the Worker0 container in the custom container service and deploys the container as a new image.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommitServiceResponse
func (client *Client) CommitServiceWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CommitServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CommitService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/commit"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CommitServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL) for a private gateway. The IP CIDR blocks added to the ACL can access the private gateway.
//
// @param tmpReq - CreateAclPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAclPolicyResponse
func (client *Client) CreateAclPolicyWithContext(ctx context.Context, ClusterId *string, GatewayId *string, tmpReq *CreateAclPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAclPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAclPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AclPolicyList) {
		request.AclPolicyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AclPolicyList, dara.String("AclPolicyList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AclPolicyListShrink) {
		query["AclPolicyList"] = request.AclPolicyListShrink
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAclPolicy"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/acl_policy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAclPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application service to obtain the inference capabilities of large models.
//
// @param request - CreateAppServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppServiceResponse
func (client *Client) CreateAppServiceWithContext(ctx context.Context, request *CreateAppServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAppServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QuotaId) {
		query["QuotaId"] = request.QuotaId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.AppVersion) {
		body["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.Replicas) {
		body["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.ServiceName) {
		body["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceSpec) {
		body["ServiceSpec"] = request.ServiceSpec
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/app_services"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a stress testing task.
//
// @param request - CreateBenchmarkTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBenchmarkTaskResponse
func (client *Client) CreateBenchmarkTaskWithContext(ctx context.Context, request *CreateBenchmarkTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateBenchmarkTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBenchmarkTask"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/benchmark-tasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBenchmarkTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建故障注入任务
//
// @param request - CreateFaultInjectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFaultInjectionResponse
func (client *Client) CreateFaultInjectionWithContext(ctx context.Context, ClusterId *string, ServiceName *string, InstanceName *string, request *CreateFaultInjectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFaultInjectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FaultArgs) {
		body["FaultArgs"] = request.FaultArgs
	}

	if !dara.IsNil(request.FaultType) {
		body["FaultType"] = request.FaultType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFaultInjection"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/instances/" + dara.PercentEncode(dara.StringValue(InstanceName)) + "/faults"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFaultInjectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a gateway.
//
// @param request - CreateGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayResponse
func (client *Client) CreateGatewayWithContext(ctx context.Context, request *CreateGatewayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenewal) {
		body["AutoRenewal"] = request.AutoRenewal
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.EnableInternet) {
		body["EnableInternet"] = request.EnableInternet
	}

	if !dara.IsNil(request.EnableIntranet) {
		body["EnableIntranet"] = request.EnableIntranet
	}

	if !dara.IsNil(request.GatewayType) {
		body["GatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Replicas) {
		body["Replicas"] = request.Replicas
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGateway"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an internal endpoint of a private gateway.
//
// @param request - CreateGatewayIntranetLinkedVpcRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayIntranetLinkedVpcResponse
func (client *Client) CreateGatewayIntranetLinkedVpcWithContext(ctx context.Context, ClusterId *string, GatewayId *string, request *CreateGatewayIntranetLinkedVpcRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGatewayIntranetLinkedVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.EnableAuthoritativeDns) {
		query["EnableAuthoritativeDns"] = request.EnableAuthoritativeDns
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGatewayIntranetLinkedVpc"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/intranet_endpoint_linked_vpc"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGatewayIntranetLinkedVpcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a VPC peering connection on an internal endpoint of a gateway.
//
// @param tmpReq - CreateGatewayIntranetLinkedVpcPeerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayIntranetLinkedVpcPeerResponse
func (client *Client) CreateGatewayIntranetLinkedVpcPeerWithContext(ctx context.Context, ClusterId *string, GatewayId *string, tmpReq *CreateGatewayIntranetLinkedVpcPeerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGatewayIntranetLinkedVpcPeerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateGatewayIntranetLinkedVpcPeerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PeerVpcs) {
		request.PeerVpcsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PeerVpcs, dara.String("PeerVpcs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PeerVpcsShrink) {
		query["PeerVpcs"] = request.PeerVpcsShrink
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGatewayIntranetLinkedVpcPeer"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/intranet_endpoint_linked_vpc_peer"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGatewayIntranetLinkedVpcPeerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource group.
//
// Description:
//
// *Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/144261.html) of Elastic Algorithm Service (EAS).
//
// @param request - CreateResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithContext(ctx context.Context, request *CreateResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenewal) {
		body["AutoRenewal"] = request.AutoRenewal
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.EcsInstanceCount) {
		body["EcsInstanceCount"] = request.EcsInstanceCount
	}

	if !dara.IsNil(request.EcsInstanceType) {
		body["EcsInstanceType"] = request.EcsInstanceType
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.ResourceName) {
		body["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SelfManagedResourceOptions) {
		body["SelfManagedResourceOptions"] = request.SelfManagedResourceOptions
	}

	if !dara.IsNil(request.SystemDiskSize) {
		body["SystemDiskSize"] = request.SystemDiskSize
	}

	if !dara.IsNil(request.Zone) {
		body["Zone"] = request.Zone
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResource"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates instances in a dedicated resource group.
//
// @param request - CreateResourceInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceInstancesResponse
func (client *Client) CreateResourceInstancesWithContext(ctx context.Context, ClusterId *string, ResourceId *string, request *CreateResourceInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenewal) {
		body["AutoRenewal"] = request.AutoRenewal
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.EcsInstanceCount) {
		body["EcsInstanceCount"] = request.EcsInstanceCount
	}

	if !dara.IsNil(request.EcsInstanceType) {
		body["EcsInstanceType"] = request.EcsInstanceType
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.SystemDiskSize) {
		body["SystemDiskSize"] = request.SystemDiskSize
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	if !dara.IsNil(request.Zone) {
		body["Zone"] = request.Zone
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceInstances"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/instances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the LogShipper feature of Log Service for a resource group.
//
// @param request - CreateResourceLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceLogResponse
func (client *Client) CreateResourceLogWithContext(ctx context.Context, ClusterId *string, ResourceId *string, request *CreateResourceLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LogStore) {
		body["LogStore"] = request.LogStore
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceLog"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/log"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a model service in Elastic Algorithm Service (EAS).
//
// Description:
//
// *Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/144261.html) of Elastic Algorithm Service (EAS).
//
// @param tmpReq - CreateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceResponse
func (client *Client) CreateServiceWithContext(ctx context.Context, tmpReq *CreateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Labels) {
		request.LabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Labels, dara.String("Labels"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Develop) {
		query["Develop"] = request.Develop
	}

	if !dara.IsNil(request.LabelsShrink) {
		query["Labels"] = request.LabelsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the Autoscaler feature and creates an Autoscaler controller for a service.
//
// @param request - CreateServiceAutoScalerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceAutoScalerResponse
func (client *Client) CreateServiceAutoScalerWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *CreateServiceAutoScalerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceAutoScalerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Behavior) {
		body["behavior"] = request.Behavior
	}

	if !dara.IsNil(request.Max) {
		body["max"] = request.Max
	}

	if !dara.IsNil(request.Min) {
		body["min"] = request.Min
	}

	if !dara.IsNil(request.ScaleStrategies) {
		body["scaleStrategies"] = request.ScaleStrategies
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceAutoScaler"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/autoscaler"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceAutoScalerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the Cron Horizontal Pod Autoscaler (CronHPA) feature for a service.
//
// @param request - CreateServiceCronScalerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceCronScalerResponse
func (client *Client) CreateServiceCronScalerWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *CreateServiceCronScalerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceCronScalerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeDates) {
		body["ExcludeDates"] = request.ExcludeDates
	}

	if !dara.IsNil(request.ScaleJobs) {
		body["ScaleJobs"] = request.ScaleJobs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceCronScaler"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/cronscaler"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceCronScalerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the traffic mirroring feature for a service. After the feature is enabled, requests received by the service can be mirrored to another service.
//
// @param request - CreateServiceMirrorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceMirrorResponse
func (client *Client) CreateServiceMirrorWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *CreateServiceMirrorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceMirrorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Ratio) {
		body["Ratio"] = request.Ratio
	}

	if !dara.IsNil(request.Target) {
		body["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceMirror"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/mirror"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceMirrorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual resource group.
//
// @param request - CreateVirtualResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirtualResourceResponse
func (client *Client) CreateVirtualResourceWithContext(ctx context.Context, request *CreateVirtualResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVirtualResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DisableSpotProtectionPeriod) {
		body["DisableSpotProtectionPeriod"] = request.DisableSpotProtectionPeriod
	}

	if !dara.IsNil(request.Resources) {
		body["Resources"] = request.Resources
	}

	if !dara.IsNil(request.VirtualResourceName) {
		body["VirtualResourceName"] = request.VirtualResourceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVirtualResource"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/virtualresources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVirtualResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access control list (ACL) for a private gateway. The IP CIDR block that is deleted from the ACL cannot access the private gateway.
//
// @param tmpReq - DeleteAclPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAclPolicyResponse
func (client *Client) DeleteAclPolicyWithContext(ctx context.Context, ClusterId *string, GatewayId *string, tmpReq *DeleteAclPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAclPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAclPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AclPolicyList) {
		request.AclPolicyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AclPolicyList, dara.String("AclPolicyList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AclPolicyListShrink) {
		query["AclPolicyList"] = request.AclPolicyListShrink
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAclPolicy"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/acl_policy"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAclPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a stress testing task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBenchmarkTaskResponse
func (client *Client) DeleteBenchmarkTaskWithContext(ctx context.Context, ClusterId *string, TaskName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteBenchmarkTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBenchmarkTask"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/benchmark-tasks/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(TaskName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBenchmarkTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除故障注入任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaultInjectionResponse
func (client *Client) DeleteFaultInjectionWithContext(ctx context.Context, ClusterId *string, ServiceName *string, InstanceName *string, FaultType *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFaultInjectionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaultInjection"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/instances/" + dara.PercentEncode(dara.StringValue(InstanceName)) + "/faults/" + dara.PercentEncode(dara.StringValue(FaultType))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaultInjectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a private gateway.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGatewayWithContext(ctx context.Context, ClusterId *string, GatewayId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGateway"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网关内网访问端点
//
// @param request - DeleteGatewayIntranetLinkedVpcRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayIntranetLinkedVpcResponse
func (client *Client) DeleteGatewayIntranetLinkedVpcWithContext(ctx context.Context, ClusterId *string, GatewayId *string, request *DeleteGatewayIntranetLinkedVpcRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGatewayIntranetLinkedVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayIntranetLinkedVpc"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/intranet_endpoint_linked_vpc"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayIntranetLinkedVpcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a VPC peering connection from an internal endpoint of a gateway.
//
// @param tmpReq - DeleteGatewayIntranetLinkedVpcPeerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayIntranetLinkedVpcPeerResponse
func (client *Client) DeleteGatewayIntranetLinkedVpcPeerWithContext(ctx context.Context, ClusterId *string, GatewayId *string, tmpReq *DeleteGatewayIntranetLinkedVpcPeerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGatewayIntranetLinkedVpcPeerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteGatewayIntranetLinkedVpcPeerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PeerVpcs) {
		request.PeerVpcsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PeerVpcs, dara.String("PeerVpcs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PeerVpcsShrink) {
		query["PeerVpcs"] = request.PeerVpcsShrink
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayIntranetLinkedVpcPeer"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/intranet_endpoint_linked_vpc_peer"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayIntranetLinkedVpcPeerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource group that contains no resources or instances.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithContext(ctx context.Context, ClusterId *string, ResourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResource"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the virtual private cloud (VPC) direct connection feature for a dedicated resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceDLinkResponse
func (client *Client) DeleteResourceDLinkWithContext(ctx context.Context, ClusterId *string, ResourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceDLinkResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceDLink"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/dlink"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceDLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the tags of an instance in a resource group.
//
// @param tmpReq - DeleteResourceInstanceLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceInstanceLabelResponse
func (client *Client) DeleteResourceInstanceLabelWithContext(ctx context.Context, ClusterId *string, ResourceId *string, tmpReq *DeleteResourceInstanceLabelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceInstanceLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteResourceInstanceLabelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.Keys) {
		request.KeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keys, dara.String("Keys"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.LabelKeys) {
		request.LabelKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelKeys, dara.String("LabelKeys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AllInstances) {
		query["AllInstances"] = request.AllInstances
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.KeysShrink) {
		query["Keys"] = request.KeysShrink
	}

	if !dara.IsNil(request.LabelKeysShrink) {
		query["LabelKeys"] = request.LabelKeysShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceInstanceLabel"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/label"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceInstanceLabelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes instances in a dedicated resource group. You can delete only pay-as-you-go instances as a regular user.
//
// @param request - DeleteResourceInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceInstancesResponse
func (client *Client) DeleteResourceInstancesWithContext(ctx context.Context, ClusterId *string, ResourceId *string, request *DeleteResourceInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllFailed) {
		query["AllFailed"] = request.AllFailed
	}

	if !dara.IsNil(request.InstanceList) {
		query["InstanceList"] = request.InstanceList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceInstances"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/instances"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the LogShipper feature of Log Service for a dedicated resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceLogResponse
func (client *Client) DeleteResourceLogWithContext(ctx context.Context, ClusterId *string, ResourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceLogResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceLog"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/log"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceResponse
func (client *Client) DeleteServiceWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the existing Autoscaler controller and disables the Autoscaler feature for a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceAutoScalerResponse
func (client *Client) DeleteServiceAutoScalerWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceAutoScalerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceAutoScaler"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/autoscaler"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceAutoScalerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the Cronscaler feature for a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceCronScalerResponse
func (client *Client) DeleteServiceCronScalerWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceCronScalerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceCronScaler"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/cronscaler"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceCronScalerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts the instances of a service.
//
// @param request - DeleteServiceInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceInstancesResponse
func (client *Client) DeleteServiceInstancesWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *DeleteServiceInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Container) {
		query["Container"] = request.Container
	}

	if !dara.IsNil(request.InstanceList) {
		query["InstanceList"] = request.InstanceList
	}

	if !dara.IsNil(request.IsReplica) {
		query["IsReplica"] = request.IsReplica
	}

	if !dara.IsNil(request.SoftRestart) {
		query["SoftRestart"] = request.SoftRestart
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceInstances"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/instances"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes existing service tags.
//
// @param tmpReq - DeleteServiceLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceLabelResponse
func (client *Client) DeleteServiceLabelWithContext(ctx context.Context, ClusterId *string, ServiceName *string, tmpReq *DeleteServiceLabelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteServiceLabelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Keys) {
		request.KeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keys, dara.String("Keys"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.LabelKeys) {
		request.LabelKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelKeys, dara.String("LabelKeys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.KeysShrink) {
		query["Keys"] = request.KeysShrink
	}

	if !dara.IsNil(request.LabelKeysShrink) {
		query["LabelKeys"] = request.LabelKeysShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceLabel"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/label"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceLabelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the traffic mirroring feature for a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceMirrorResponse
func (client *Client) DeleteServiceMirrorWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceMirrorResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceMirror"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/mirror"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceMirrorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual resource group that contains no resources or instances.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirtualResourceResponse
func (client *Client) DeleteVirtualResourceWithContext(ctx context.Context, ClusterId *string, VirtualResourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteVirtualResourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVirtualResource"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/virtualresources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(VirtualResourceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVirtualResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about the configurations of a stress testing task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBenchmarkTaskResponse
func (client *Client) DescribeBenchmarkTaskWithContext(ctx context.Context, ClusterId *string, TaskName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeBenchmarkTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBenchmarkTask"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/benchmark-tasks/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(TaskName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBenchmarkTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the report of a stress testing task.
//
// @param request - DescribeBenchmarkTaskReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBenchmarkTaskReportResponse
func (client *Client) DescribeBenchmarkTaskReportWithContext(ctx context.Context, ClusterId *string, TaskName *string, request *DescribeBenchmarkTaskReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeBenchmarkTaskReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportType) {
		query["ReportType"] = request.ReportType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBenchmarkTaskReport"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/benchmark-tasks/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(TaskName)) + "/report"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBenchmarkTaskReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a private gateway.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGatewayResponse
func (client *Client) DescribeGatewayWithContext(ctx context.Context, ClusterId *string, GatewayId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeGatewayResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGateway"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a service group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupResponse
func (client *Client) DescribeGroupWithContext(ctx context.Context, ClusterId *string, GroupName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroup"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/groups/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GroupName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of endpoints of service groups.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupEndpointsResponse
func (client *Client) DescribeGroupEndpointsWithContext(ctx context.Context, ClusterId *string, GroupName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeGroupEndpointsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupEndpoints"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/groups/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GroupName)) + "/endpoints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instance types for an available instance in a shared resource group.
//
// @param tmpReq - DescribeMachineSpecRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMachineSpecResponse
func (client *Client) DescribeMachineSpecWithContext(ctx context.Context, tmpReq *DescribeMachineSpecRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeMachineSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeMachineSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceTypes) {
		request.InstanceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceTypes, dara.String("InstanceTypes"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.InstanceTypesShrink) {
		query["InstanceTypes"] = request.InstanceTypesShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMachineSpec"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/public/instance_types"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMachineSpecResponse{}
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/regions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Queries the information about a resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceResponse
func (client *Client) DescribeResourceWithContext(ctx context.Context, ClusterId *string, ResourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeResourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResource"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries detailed configurations about a virtual private cloud (VPC) direct connection of a dedicated resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceDLinkResponse
func (client *Client) DescribeResourceDLinkWithContext(ctx context.Context, ClusterId *string, ResourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeResourceDLinkResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceDLink"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/dlink"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceDLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about the LogShipper configurations of Log Service for a dedicated resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceLogResponse
func (client *Client) DescribeResourceLogWithContext(ctx context.Context, ClusterId *string, ResourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeResourceLogResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceLog"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/log"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceResponse
func (client *Client) DescribeServiceWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the Autoscaler configurations of a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceAutoScalerResponse
func (client *Client) DescribeServiceAutoScalerWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeServiceAutoScalerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceAutoScaler"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/autoscaler"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceAutoScalerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Cron Horizontal Pod Autoscaler (CronHPA) configurations of a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceCronScalerResponse
func (client *Client) DescribeServiceCronScalerWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeServiceCronScalerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceCronScaler"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/cronscaler"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceCronScalerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the diagnostics details of a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceDiagnosisResponse
func (client *Client) DescribeServiceDiagnosisWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeServiceDiagnosisResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceDiagnosis"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/diagnosis"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceDiagnosisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of service endpoints.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceEndpointsResponse
func (client *Client) DescribeServiceEndpointsWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeServiceEndpointsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceEndpoints"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/endpoints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about recent service deployment events.
//
// @param request - DescribeServiceEventRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceEventResponse
func (client *Client) DescribeServiceEventWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *DescribeServiceEventRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeServiceEventResponse, _err error) {
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

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceEvent"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the diagnostics details of an instance that runs Elastic Algorithm Service (EAS).
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceInstanceDiagnosisResponse
func (client *Client) DescribeServiceInstanceDiagnosisWithContext(ctx context.Context, ClusterId *string, ServiceName *string, InstanceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeServiceInstanceDiagnosisResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceInstanceDiagnosis"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/instances/" + dara.PercentEncode(dara.StringValue(InstanceName)) + "/diagnosis"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceInstanceDiagnosisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the logs of a service.
//
// @param request - DescribeServiceLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceLogResponse
func (client *Client) DescribeServiceLogWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *DescribeServiceLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeServiceLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContainerName) {
		query["ContainerName"] = request.ContainerName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Previous) {
		query["Previous"] = request.Previous
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceLog"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about the traffic mirroring settings of a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceMirrorResponse
func (client *Client) DescribeServiceMirrorWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeServiceMirrorResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceMirror"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/mirror"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceMirrorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the logon-free URL of the service.
//
// @param request - DescribeServiceSignedUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceSignedUrlResponse
func (client *Client) DescribeServiceSignedUrlWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *DescribeServiceSignedUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeServiceSignedUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Expire) {
		query["Expire"] = request.Expire
	}

	if !dara.IsNil(request.Internal) {
		query["Internal"] = request.Internal
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceSignedUrl"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/signed_url"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceSignedUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical prices of preemptible instances. For more information about preemptible instances, see Create and use preemptible instances.
//
// @param request - DescribeSpotDiscountHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSpotDiscountHistoryResponse
func (client *Client) DescribeSpotDiscountHistoryWithContext(ctx context.Context, request *DescribeSpotDiscountHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeSpotDiscountHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.IsProtect) {
		query["IsProtect"] = request.IsProtect
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSpotDiscountHistory"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/public/spot_discount"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSpotDiscountHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Views the details of a virtual resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVirtualResourceResponse
func (client *Client) DescribeVirtualResourceWithContext(ctx context.Context, ClusterId *string, VirtualResourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeVirtualResourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVirtualResource"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/virtualresources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(VirtualResourceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVirtualResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a custom domain name from a private gateway.
//
// @param tmpReq - DetachGatewayDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachGatewayDomainResponse
func (client *Client) DetachGatewayDomainWithContext(ctx context.Context, ClusterId *string, GatewayId *string, tmpReq *DetachGatewayDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DetachGatewayDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetachGatewayDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomDomain) {
		request.CustomDomainShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomDomain, dara.String("CustomDomain"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomDomainShrink) {
		query["CustomDomain"] = request.CustomDomainShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachGatewayDomain"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/domain/detach"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachGatewayDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches a container service to development mode or exits development mode.
//
// @param request - DevelopServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DevelopServiceResponse
func (client *Client) DevelopServiceWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *DevelopServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DevelopServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Exit) {
		query["Exit"] = request.Exit
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DevelopService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/develop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DevelopServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries access control lists (ACLs) created for a private gateway.
//
// @param request - ListAclPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAclPolicyResponse
func (client *Client) ListAclPolicyWithContext(ctx context.Context, ClusterId *string, GatewayId *string, request *ListAclPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAclPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAclPolicy"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/acl_policy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAclPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of stress testing tasks that are created by the current user.
//
// @param request - ListBenchmarkTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBenchmarkTaskResponse
func (client *Client) ListBenchmarkTaskWithContext(ctx context.Context, request *ListBenchmarkTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListBenchmarkTaskResponse, _err error) {
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

	if !dara.IsNil(request.ModelId) {
		query["ModelId"] = request.ModelId
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

	if !dara.IsNil(request.RequestMethod) {
		query["RequestMethod"] = request.RequestMethod
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBenchmarkTask"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/benchmark-tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBenchmarkTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of private gateways.
//
// @param request - ListGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayResponse
func (client *Client) ListGatewayWithContext(ctx context.Context, request *ListGatewayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGatewayResponse, _err error) {
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

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayName) {
		query["GatewayName"] = request.GatewayName
	}

	if !dara.IsNil(request.GatewayType) {
		query["GatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.InternetEnabled) {
		query["InternetEnabled"] = request.InternetEnabled
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

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGateway"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of custom domain names of a private gateway.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayDomainsResponse
func (client *Client) ListGatewayDomainsWithContext(ctx context.Context, ClusterId *string, GatewayId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGatewayDomainsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayDomains"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/domains"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of the internal endpoints of a private gateway.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayIntranetLinkedVpcResponse
func (client *Client) ListGatewayIntranetLinkedVpcWithContext(ctx context.Context, ClusterId *string, GatewayId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGatewayIntranetLinkedVpcResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayIntranetLinkedVpc"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/intranet_endpoint_linked_vpc"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayIntranetLinkedVpcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of all VPC peering connections on internal endpoint of a gateway.
//
// @param request - ListGatewayIntranetLinkedVpcPeerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayIntranetLinkedVpcPeerResponse
func (client *Client) ListGatewayIntranetLinkedVpcPeerWithContext(ctx context.Context, ClusterId *string, GatewayId *string, request *ListGatewayIntranetLinkedVpcPeerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGatewayIntranetLinkedVpcPeerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayIntranetLinkedVpcPeer"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/intranet_endpoint_linked_vpc_peer"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayIntranetLinkedVpcPeerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the zones supported by a gateway within an intranet.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayIntranetSupportedZoneResponse
func (client *Client) ListGatewayIntranetSupportedZoneWithContext(ctx context.Context, GatewayId *string, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGatewayIntranetSupportedZoneResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayIntranetSupportedZone"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId)) + "/intranet_supported_zone"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayIntranetSupportedZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries created service groups.
//
// @param request - ListGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithContext(ctx context.Context, request *ListGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
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

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.TrafficMode) {
		query["TrafficMode"] = request.TrafficMode
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroups"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/groups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workers in a resource group.
//
// @param request - ListResourceInstanceWorkerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceInstanceWorkerResponse
func (client *Client) ListResourceInstanceWorkerWithContext(ctx context.Context, ClusterId *string, ResourceId *string, InstanceName *string, request *ListResourceInstanceWorkerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceInstanceWorkerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Ready) {
		query["Ready"] = request.Ready
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkerName) {
		query["WorkerName"] = request.WorkerName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceInstanceWorker"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/instance/" + dara.PercentEncode(dara.StringValue(InstanceName)) + "/workers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceInstanceWorkerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instances in a dedicated resource group.
//
// @param tmpReq - ListResourceInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceInstancesResponse
func (client *Client) ListResourceInstancesWithContext(ctx context.Context, ClusterId *string, ResourceId *string, tmpReq *ListResourceInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListResourceInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Label) {
		request.LabelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Label, dara.String("Label"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceIP) {
		query["InstanceIP"] = request.InstanceIP
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceStatus) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !dara.IsNil(request.LabelShrink) {
		query["Label"] = request.LabelShrink
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

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceInstances"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListResourceServices is deprecated
//
// Summary:
//
// Queries a list of services that are deployed in the dedicated resource group.
//
// @param request - ListResourceServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceServicesResponse
func (client *Client) ListResourceServicesWithContext(ctx context.Context, ClusterId *string, ResourceId *string, request *ListResourceServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceServicesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceServices"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/services"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of dedicated resource groups for the current user.
//
// @param request - ListResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithContext(ctx context.Context, request *ListResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceStatus) {
		query["ResourceStatus"] = request.ResourceStatus
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResources"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the containers of a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceContainersResponse
func (client *Client) ListServiceContainersWithContext(ctx context.Context, ClusterId *string, ServiceName *string, InstanceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServiceContainersResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceContainers"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/instances/" + dara.PercentEncode(dara.StringValue(InstanceName)) + "/containers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceContainersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取故障注入信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceFaultInjectionInfoResponse
func (client *Client) ListServiceInstanceFaultInjectionInfoWithContext(ctx context.Context, ClusterId *string, ServiceName *string, InstanceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceFaultInjectionInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceInstanceFaultInjectionInfo"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/instances/" + dara.PercentEncode(dara.StringValue(InstanceName)) + "/faults"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceInstanceFaultInjectionInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instances of a service.
//
// @param request - ListServiceInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstancesResponse
func (client *Client) ListServiceInstancesWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *ListServiceInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServiceInstancesResponse, _err error) {
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

	if !dara.IsNil(request.HostIP) {
		query["HostIP"] = request.HostIP
	}

	if !dara.IsNil(request.InstanceIP) {
		query["InstanceIP"] = request.InstanceIP
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceStatus) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.IsSpot) {
		query["IsSpot"] = request.IsSpot
	}

	if !dara.IsNil(request.ListReplica) {
		query["ListReplica"] = request.ListReplica
	}

	if !dara.IsNil(request.MemberType) {
		query["MemberType"] = request.MemberType
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

	if !dara.IsNil(request.ReplicaName) {
		query["ReplicaName"] = request.ReplicaName
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceInstances"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the historical versions of a service.
//
// @param request - ListServiceVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceVersionsResponse
func (client *Client) ListServiceVersionsWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *ListServiceVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServiceVersionsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceVersions"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists services.
//
// @param tmpReq - ListServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithContext(ctx context.Context, tmpReq *ListServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListServicesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Label) {
		request.LabelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Label, dara.String("Label"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoscalerEnabled) {
		query["AutoscalerEnabled"] = request.AutoscalerEnabled
	}

	if !dara.IsNil(request.CronscalerEnabled) {
		query["CronscalerEnabled"] = request.CronscalerEnabled
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.Gateway) {
		query["Gateway"] = request.Gateway
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.IncludeNoWorkspace) {
		query["IncludeNoWorkspace"] = request.IncludeNoWorkspace
	}

	if !dara.IsNil(request.LabelShrink) {
		query["Label"] = request.LabelShrink
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

	if !dara.IsNil(request.ParentServiceUid) {
		query["ParentServiceUid"] = request.ParentServiceUid
	}

	if !dara.IsNil(request.QuotaId) {
		query["QuotaId"] = request.QuotaId
	}

	if !dara.IsNil(request.ResourceAliasName) {
		query["ResourceAliasName"] = request.ResourceAliasName
	}

	if !dara.IsNil(request.ResourceBurstable) {
		query["ResourceBurstable"] = request.ResourceBurstable
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceStatus) {
		query["ServiceStatus"] = request.ServiceStatus
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	if !dara.IsNil(request.ServiceUid) {
		query["ServiceUid"] = request.ServiceUid
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.TrafficState) {
		query["TrafficState"] = request.TrafficState
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServices"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tenant plug-ins.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantAddonsResponse
func (client *Client) ListTenantAddonsWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTenantAddonsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTenantAddons"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/tenantaddons"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTenantAddonsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of virtual resource groups for the current user.
//
// @param request - ListVirtualResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirtualResourceResponse
func (client *Client) ListVirtualResourceWithContext(ctx context.Context, request *ListVirtualResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVirtualResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.VirtualResourceId) {
		query["VirtualResourceId"] = request.VirtualResourceId
	}

	if !dara.IsNil(request.VirtualResourceName) {
		query["VirtualResourceName"] = request.VirtualResourceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVirtualResource"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/virtualresources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVirtualResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates resource group instances.
//
// @param request - MigrateResourceInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateResourceInstanceResponse
func (client *Client) MigrateResourceInstanceWithContext(ctx context.Context, ClusterId *string, ResourceId *string, request *MigrateResourceInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MigrateResourceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestResourceId) {
		body["DestResourceId"] = request.DestResourceId
	}

	if !dara.IsNil(request.InstanceIds) {
		body["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.MigrateToHybrid) {
		body["MigrateToHybrid"] = request.MigrateToHybrid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateResourceInstance"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/instances/migrate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateResourceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets tenant configurations.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReinstallTenantAddonResponse
func (client *Client) ReinstallTenantAddonWithContext(ctx context.Context, ClusterId *string, TenantAddonName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReinstallTenantAddonResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReinstallTenantAddon"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/tenantaddons/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(TenantAddonName)) + "/reinstall"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReinstallTenantAddonResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switch the traffic state or weight of the service.
//
// @param request - ReleaseServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseServiceResponse
func (client *Client) ReleaseServiceWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *ReleaseServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReleaseServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TrafficState) {
		body["TrafficState"] = request.TrafficState
	}

	if !dara.IsNil(request.Weight) {
		body["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/release"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartServiceResponse
func (client *Client) RestartServiceWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/restart"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a stress testing task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartBenchmarkTaskResponse
func (client *Client) StartBenchmarkTaskWithContext(ctx context.Context, ClusterId *string, TaskName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartBenchmarkTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartBenchmarkTask"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/benchmark-tasks/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(TaskName)) + "/start"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartBenchmarkTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartServiceResponse
func (client *Client) StartServiceWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/start"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a stress testing task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopBenchmarkTaskResponse
func (client *Client) StopBenchmarkTaskWithContext(ctx context.Context, ClusterId *string, TaskName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopBenchmarkTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopBenchmarkTask"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/benchmark-tasks/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(TaskName)) + "/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopBenchmarkTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a running service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopServiceResponse
func (client *Client) StopServiceWithContext(ctx context.Context, ClusterId *string, ServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an application service.
//
// @param request - UpdateAppServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppServiceResponse
func (client *Client) UpdateAppServiceWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *UpdateAppServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAppServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QuotaId) {
		query["QuotaId"] = request.QuotaId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.AppVersion) {
		body["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.Replicas) {
		body["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.ServiceSpec) {
		body["ServiceSpec"] = request.ServiceSpec
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/app_services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a stress testing task.
//
// @param request - UpdateBenchmarkTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBenchmarkTaskResponse
func (client *Client) UpdateBenchmarkTaskWithContext(ctx context.Context, ClusterId *string, TaskName *string, request *UpdateBenchmarkTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateBenchmarkTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBenchmarkTask"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/benchmark-tasks/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(TaskName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBenchmarkTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a private gateway.
//
// @param request - UpdateGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayResponse
func (client *Client) UpdateGatewayWithContext(ctx context.Context, GatewayId *string, ClusterId *string, request *UpdateGatewayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnableInternet) {
		body["EnableInternet"] = request.EnableInternet
	}

	if !dara.IsNil(request.EnableIntranet) {
		body["EnableIntranet"] = request.EnableIntranet
	}

	if !dara.IsNil(request.EnableSSLRedirection) {
		body["EnableSSLRedirection"] = request.EnableSSLRedirection
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.IsDefault) {
		body["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Replicas) {
		body["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.VSwitchIds) {
		body["VSwitchIds"] = request.VSwitchIds
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGateway"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/gateways/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GatewayId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the specific fields of a service group.
//
// @param request - UpdateGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithContext(ctx context.Context, ClusterId *string, GroupName *string, request *UpdateGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TrafficMode) {
		body["TrafficMode"] = request.TrafficMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGroup"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/groups/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(GroupName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a dedicated resource group. Only the name of a dedicated resource group can be updated.
//
// @param request - UpdateResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithContext(ctx context.Context, ClusterId *string, ResourceId *string, request *UpdateResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceName) {
		body["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.SelfManagedResourceOptions) {
		body["SelfManagedResourceOptions"] = request.SelfManagedResourceOptions
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResource"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of a virtual private cloud (VPC) direct connection for a dedicated resource group.
//
// @param request - UpdateResourceDLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceDLinkResponse
func (client *Client) UpdateResourceDLinkWithContext(ctx context.Context, ClusterId *string, ResourceId *string, request *UpdateResourceDLinkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceDLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCIDRs) {
		body["DestinationCIDRs"] = request.DestinationCIDRs
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.VSwitchId) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VSwitchIdList) {
		body["VSwitchIdList"] = request.VSwitchIdList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceDLink"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/dlink"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceDLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the service scheduling status of an instance in a dedicated resource group.
//
// @param request - UpdateResourceInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceInstanceResponse
func (client *Client) UpdateResourceInstanceWithContext(ctx context.Context, ClusterId *string, ResourceId *string, InstanceId *string, request *UpdateResourceInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["Action"] = request.Action
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceInstance"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/instances/" + dara.PercentEncode(dara.StringValue(InstanceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the tag of an instance in a resource group.
//
// @param tmpReq - UpdateResourceInstanceLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceInstanceLabelResponse
func (client *Client) UpdateResourceInstanceLabelWithContext(ctx context.Context, ClusterId *string, ResourceId *string, tmpReq *UpdateResourceInstanceLabelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceInstanceLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateResourceInstanceLabelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AllInstances) {
		query["AllInstances"] = request.AllInstances
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceInstanceLabel"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/label"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceInstanceLabelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a model or processor of a service. If only the metadata.instance field is updated, manual scaling can be performed.
//
// @param request - UpdateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceResponse
func (client *Client) UpdateServiceWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *UpdateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberToUpdate) {
		query["MemberToUpdate"] = request.MemberToUpdate
	}

	if !dara.IsNil(request.UpdateType) {
		query["UpdateType"] = request.UpdateType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateService"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the Autoscaler configurations of a service.
//
// @param request - UpdateServiceAutoScalerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceAutoScalerResponse
func (client *Client) UpdateServiceAutoScalerWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *UpdateServiceAutoScalerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceAutoScalerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Behavior) {
		body["behavior"] = request.Behavior
	}

	if !dara.IsNil(request.Max) {
		body["max"] = request.Max
	}

	if !dara.IsNil(request.Min) {
		body["min"] = request.Min
	}

	if !dara.IsNil(request.ScaleStrategies) {
		body["scaleStrategies"] = request.ScaleStrategies
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceAutoScaler"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/autoscaler"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceAutoScalerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the Cron Horizontal Pod Autoscaler (CronHPA) settings of a service.
//
// @param request - UpdateServiceCronScalerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceCronScalerResponse
func (client *Client) UpdateServiceCronScalerWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *UpdateServiceCronScalerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceCronScalerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeDates) {
		body["ExcludeDates"] = request.ExcludeDates
	}

	if !dara.IsNil(request.ScaleJobs) {
		body["ScaleJobs"] = request.ScaleJobs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceCronScaler"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/cronscaler"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceCronScalerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates attributes of service instances. Only isolation can be performed for service instances.
//
// @param request - UpdateServiceInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceInstanceResponse
func (client *Client) UpdateServiceInstanceWithContext(ctx context.Context, ClusterId *string, ServiceName *string, InstanceName *string, request *UpdateServiceInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsReplica) {
		query["IsReplica"] = request.IsReplica
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Detach) {
		body["Detach"] = request.Detach
	}

	if !dara.IsNil(request.Hibernate) {
		body["Hibernate"] = request.Hibernate
	}

	if !dara.IsNil(request.Isolate) {
		body["Isolate"] = request.Isolate
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceInstance"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/instances/" + dara.PercentEncode(dara.StringValue(InstanceName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds service tags or updates existing service tags.
//
// @param request - UpdateServiceLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceLabelResponse
func (client *Client) UpdateServiceLabelWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *UpdateServiceLabelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceLabelResponse, _err error) {
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
		Action:      dara.String("UpdateServiceLabel"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/label"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceLabelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the traffic mirroring configurations of a service.
//
// @param request - UpdateServiceMirrorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceMirrorResponse
func (client *Client) UpdateServiceMirrorWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *UpdateServiceMirrorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceMirrorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Ratio) {
		body["Ratio"] = request.Ratio
	}

	if !dara.IsNil(request.Target) {
		body["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceMirror"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/mirror"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceMirrorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the safety lock of a service to minimize misoperations on the service.
//
// @param request - UpdateServiceSafetyLockRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceSafetyLockResponse
func (client *Client) UpdateServiceSafetyLockWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *UpdateServiceSafetyLockRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceSafetyLockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lock) {
		body["Lock"] = request.Lock
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceSafetyLock"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/lock"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceSafetyLockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the version of a service or rolls back the service to a specific version.
//
// @param request - UpdateServiceVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceVersionResponse
func (client *Client) UpdateServiceVersionWithContext(ctx context.Context, ClusterId *string, ServiceName *string, request *UpdateServiceVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		body["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceVersion"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/services/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(ServiceName)) + "/version"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a virtual resource group.
//
// @param request - UpdateVirtualResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVirtualResourceResponse
func (client *Client) UpdateVirtualResourceWithContext(ctx context.Context, ClusterId *string, VirtualResourceId *string, request *UpdateVirtualResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateVirtualResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DisableSpotProtectionPeriod) {
		body["DisableSpotProtectionPeriod"] = request.DisableSpotProtectionPeriod
	}

	if !dara.IsNil(request.Resources) {
		body["Resources"] = request.Resources
	}

	if !dara.IsNil(request.VirtualResourceName) {
		body["VirtualResourceName"] = request.VirtualResourceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVirtualResource"),
		Version:     dara.String("2021-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/virtualresources/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/" + dara.PercentEncode(dara.StringValue(VirtualResourceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVirtualResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
