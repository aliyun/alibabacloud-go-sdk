// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates or updates an Elastic Container Instance (ECI) scaling configuration by using a configuration file.
//
// Description:
//
// When you manage ECI scaling configurations by using YAML configuration files, take note of the following items:
//
// - If you specify the ID of an ECI scaling configuration, the system updates the ECI scaling configuration based on the YAML configuration file.
//
// - If you do not specify the ID of an ECI scaling configuration, the system creates an ECI scaling configuration based on the YAML configuration file.
//
// @param request - ApplyEciScalingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyEciScalingConfigurationResponse
func (client *Client) ApplyEciScalingConfigurationWithContext(ctx context.Context, request *ApplyEciScalingConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ApplyEciScalingConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.Format) {
		query["Format"] = request.Format
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScalingConfigurationId) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyEciScalingConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyEciScalingConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and manages a scaling group by using a configuration file.
//
// Description:
//
// The ApplyScalingGroup operation supports the Kubernetes Deployment resource definition format to quickly create Elastic Container Instance (ECI) scaling groups. It also supports ECI extended annotations for Kubernetes YAML. For more information, see the "Supported annotations" section.
//
// Mapping between YAML configurations and scaling groups: A scaling group name is mapped by the namespace, kind, and name tuple in the YAML definition. In the same region, the same YAML configuration can only be mapped to the same scaling group. Example: If you use a Deployment YAML configuration with the name nginx in the default namespace, it is mapped to a scaling group named k8s_default_Deployment_nginx in the same region.
//
// Logic for managing scaling groups based on YAML configurations:
//
// - If the scaling group mapped to the YAML configuration exists, the scaling group is updated based on the YAML configuration.
//
// - If the scaling group mapped to the YAML configuration does not exist, a corresponding scaling group is created based on the YAML configuration.
//
// ### Precautions
//
// 1. If the YAML configuration does not specify the VPC, vSwitch, and security group annotations, the system performs automatic creation of a default VPC, creates a default vSwitch in the VPC, and creates the default Auto Scaling security group ess-default-sg. By default, the security group policy allows traffic on TCP ports 22 and 3389, and the ICMP (IPv4) protocol. If you have other port and protocol requirements, you can modify the security group policy.
//
// 2. When you use a public image, you must enable public network access and configure the Settings for the k8s.aliyun.com/eci-with-eip pod annotation to enable the elastic IP address (EIP).
//
// 3. After the ApplyScalingGroup operation applies the YAML configuration, the scaling group and scaling configuration take effect immediately. If replicas is set to a value greater than 0, automatic creation of resources is triggered.
//
// ### Supported annotations
//
// The annotations are referenced from [ECI Pod Annotation](https://help.aliyun.com/document_detail/186939.html).
//
// |Parameter|Example|Description|
//
// |-----|-----|-----|
//
// |k8s.aliyun.com/ess-scaling-group-min-size|1|The minimum number of instances in the scaling group. Default value: 0.
//
// |k8s.aliyun.com/ess-scaling-group-max-size|20|The maximum number of instances in the scaling group. Default value: max(replicas, 30).
//
// |k8s.aliyun.com/eci-ntp-server|100.100.*.*|The NTP server.
//
// |k8s.aliyun.com/eci-use-specs|2-4Gi|The instance specifications, such as 2 vCPUs and 4 GiB of memory. For more information, see [Create a pod by using multiple instance types](https://help.aliyun.com/document_detail/451267.html).
//
// |k8s.aliyun.com/eci-vswitch|vsw-bp1xpiowfm5vo8o3c****|The ID of the vSwitch. You can specify multiple vSwitches to implement the multi-active zone feature.
//
// |k8s.aliyun.com/eci-security-group|sg-bp1dktddjsg5nktv****|The ID of the security group. Requirements: <ul><li>You can specify one or more security groups. A maximum of five security groups can be specified.</li><li>The specified security groups must belong to the same VPC.</li><li>The specified security groups must be of the same type.</li></ul>
//
// |k8s.aliyun.com/eci-sls-enable|"false"|Specifies whether to shutdown log collection. A value of false indicates that log collection is shutdown. <p>When you collect logs by using Simple Log Service CRDs, if you do not need to collect logs for specific pods, you can configure this annotation to shutdown the log collection feature. This prevents the system from performing automatic creation of Logtail and wasting resources.</p>
//
// |k8s.aliyun.com/eci-spot-strategy|SpotAsPriceGo|The bidding policy for the spot instance. Valid values: <ul><li>SpotWithPriceLimit: You use custom Settings for the maximum hourly price of the spot instance. In this case, you must set k8s.aliyun.com/eci-spot-price-limit.</li><li>SpotAsPriceGo: The system automatically places bids based on the actual price of the current marketplace.</li></ul>
//
// |k8s.aliyun.com/eci-spot-price-limit|"0.5"|The maximum hourly price for the spot instance. The value is accurate to three decimal places. <p>This parameter takes effect only when the Settings of k8s.aliyun.com/eci-spot-strategy are SpotWithPriceLimit.</p>
//
// |k8s.aliyun.com/eci-with-eip|"true"|Specifies whether to perform automatic creation and attach an EIP. A value of true indicates that an EIP is automatically created and attached.
//
// |k8s.aliyun.com/eci-data-cache-bucket|default|The bucket of DataCache. This parameter is required when you use DataCache to create a pod.
//
// |k8s.aliyun.com/eci-data-cache-pl|PL1|The performance level (PL) of the cloud disk created based on DataCache. <p>By default, an ESSD standard SSD is used, and the default PL is PL1.</p>
//
// |k8s.aliyun.com/eci-data-cache-provisionedIops|"40000"|The provisioned read and write IOPS of the ESSD AutoPL cloud disk. Valid values: 0 to min{50000, 1000 × Capacity - Baseline performance}. Baseline performance = min{1800 + 50 × Capacity, 50000}. For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html). <p>If this annotation is added, the cloud disk created based on DataCache is an ESSD AutoPL cloud disk.</p>
//
// |k8s.aliyun.com/eci-data-cache-burstingEnabled|"true"|Specifies whether to enable the burst feature for the ESSD AutoPL cloud disk. For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html). <p>If this annotation is added, the cloud disk created based on DataCache is an ESSD AutoPL cloud disk.</p>
//
// |k8s.aliyun.com/eci-custom-tags|"env:test,name:alice"|The label string attached to the pod. You can attach up to three labels. Separate the label key and label value with a half-width colon character, and separate multiple labels with a half-width comma character.
//
// @param request - ApplyScalingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyScalingGroupResponse
func (client *Client) ApplyScalingGroupWithContext(ctx context.Context, request *ApplyScalingGroupRequest, runtime *dara.RuntimeOptions) (_result *ApplyScalingGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.Format) {
		query["Format"] = request.Format
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyScalingGroup"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyScalingGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more Application Load Balancer (ALB) server groups to a scaling group. This operation is applicable to scenarios where you need to dynamically adjust the number of instances to handle fluctuating workloads and ensure high availability of applications. It enables automatic elastic scaling during business fluctuations, dynamically distributes traffic based on loads, and improves application stability.
//
// Description:
//
// When you add ALB server groups to a scaling group, the following conditions must be met:
//
// - The network type of the scaling group must be VPC, and the scaling group must be in the same VPC as the ALB server groups.
//
// - The ALB server groups must be in the Available state.
//
// @param request - AttachAlbServerGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachAlbServerGroupsResponse
func (client *Client) AttachAlbServerGroupsWithContext(ctx context.Context, request *AttachAlbServerGroupsRequest, runtime *dara.RuntimeOptions) (_result *AttachAlbServerGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlbServerGroups) {
		query["AlbServerGroups"] = request.AlbServerGroups
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ForceAttach) {
		query["ForceAttach"] = request.ForceAttach
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachAlbServerGroups"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachAlbServerGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches one or more RDS instances to a scaling group, which allows you to adjust the number of RDS instances in the scaling group at any time to increase flexibility and usability.
//
// Description:
//
// The following conditions must be met when you attach RDS instances to a scaling group:
//
// - The RDS instances and the scaling group must belong to the same account.
//
// - The RDS instances must not be locked. For more information about the locking policy, see [Usage notes for ApsaraDB RDS](https://help.aliyun.com/document_detail/41872.html).
//
// - The RDS instances must be in the Running state.
//
// - The RDS instances must be owned by your current account and must not be deleted.
//
// - If you attach an RDS instance that already exists in the scaling group to the scaling group again, the number of RDS instances in the scaling group does not change. However, the system attempts to add the private IP addresses of all ECS instances in the current scaling group to the IP address whitelist of the RDS instance.
//
// > After you attach RDS instances, the default whitelist group of the RDS instances can contain a maximum of 1,000 IP addresses. For more information about IP address whitelists, see [Configure an IP address whitelist](https://help.aliyun.com/document_detail/96118.html).
//
// @param request - AttachDBInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachDBInstancesResponse
func (client *Client) AttachDBInstancesWithContext(ctx context.Context, request *AttachDBInstancesRequest, runtime *dara.RuntimeOptions) (_result *AttachDBInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachMode) {
		query["AttachMode"] = request.AttachMode
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstances) {
		query["DBInstances"] = request.DBInstances
	}

	if !dara.IsNil(request.ForceAttach) {
		query["ForceAttach"] = request.ForceAttach
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachDBInstances"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachDBInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually adds instances to a scaling group. You can add ECS instances, Elastic Container Instance (ECI) instances, or managed instances outside the scaling group to the scaling group and put them into service. You can also restart ECS instances in the scaling group that are in the Economical Mode to the Running state and put them into service.
//
// Description:
//
// - Before you invoke this operation, make sure that the following conditions are met:
//
//   - The scaling group is in the Active state.
//
//   - No scaling activities are in progress in the scaling group.
//
// - Limits on ECS instances or ECI instances to be added to the scaling group:
//
//   - The instances must be in the same region as the scaling group.
//
//   - The instances must be in the Running state.
//
//   - The instances must not be added to other scaling groups.
//
//   - The billing method of the instances must be subscription, pay-as-you-go, or spot instance.
//
//   - If a vSwitch ID is specified for the scaling group, classic network-connected ECS instances or ECI instances, or ECS instances or ECI instances in other VPCs cannot be added to the scaling group.
//
//   - If no vSwitch ID is specified for the scaling group, VPC-connected ECS instances or ECI instances cannot be added to the scaling group.
//
// - Limits on managed instances to be added to the scaling group:
//
//   - The instances must be in the same region as the scaling group.
//
//   - Only managed instances that are not added to any scaling group can be added to an ECS scaling group.
//
//   - Managed instances do not support lifecycle hooks, association with load balancing, or the health check feature.
//
//   - After managed instances are added to an ECS scaling group, the scaling group does not manage the lifecycle of the managed instances.
//
//   - Managed instances can only be manually added to or removed from the scaling group. Removed managed instances are not released.
//
// - Precautions for using this operation:
//
//   - If no scaling activities are in progress in the scaling group, this operation can bypass the cooldown time and be executed directly.
//
//   - A successful response only indicates that Auto Scaling accepted the request of the operation invocation and the scaling activity can be executed. It does not mean that the scaling activity will be executed successfully. You must use the returned ScalingActivityId to check the execution status of the scaling activity.
//
//   - If the sum of the number of instances specified in this operation and the current number of instances in the scaling group (Total Capacity) is greater than the maximum number of instances (MaxSize), the operation invocation fails.
//
//   - ECS instances, ECI instances, or managed instances manually added by using this operation are not associated with the active scaling configuration of the scaling group.
//
// @param request - AttachInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachInstancesResponse
func (client *Client) AttachInstancesWithContext(ctx context.Context, request *AttachInstancesRequest, runtime *dara.RuntimeOptions) (_result *AttachInstancesResponse, _err error) {
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

	if !dara.IsNil(request.Entrusted) {
		query["Entrusted"] = request.Entrusted
	}

	if !dara.IsNil(request.IgnoreInvalidInstance) {
		query["IgnoreInvalidInstance"] = request.IgnoreInvalidInstance
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.LifecycleHook) {
		query["LifecycleHook"] = request.LifecycleHook
	}

	if !dara.IsNil(request.LoadBalancerWeights) {
		query["LoadBalancerWeights"] = request.LoadBalancerWeights
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachInstances"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates one or more load balancing instances with a scaling group to distribute access traffic to multiple instances in the scaling group and enhance the service capability of the scaling group. You can call the AttachLoadBalancers API operation to add one or more load balancing instances.
//
// Description:
//
// The following conditions must be met when adding load balancing instances to a scaling group:
//
// - The load balancing instances and the scaling group must belong to the same account and region.
//
// - The load balancing instances must be in the `Running` state.
//
// - Each load balancing instance must be configured with at least one listener and have health checks enabled.
//
// - If the network type of both the load balancing instances and the scaling group is VPC, they must be in the same VPC.
//
// - If the network type of the scaling group is VPC and the network type of the load balancing instances is classic network, and the backend servers of the load balancing instances include VPC-connected instances, the VPC-connected instances must be in the same VPC as the scaling group.
//
// - The number of load balancing instances associated with the scaling group does not exceed the quota. For more information about the quota, see [Limits](https://help.aliyun.com/document_detail/25863.html).
//
// @param request - AttachLoadBalancersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachLoadBalancersResponse
func (client *Client) AttachLoadBalancersWithContext(ctx context.Context, request *AttachLoadBalancersRequest, runtime *dara.RuntimeOptions) (_result *AttachLoadBalancersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Async) {
		query["Async"] = request.Async
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ForceAttach) {
		query["ForceAttach"] = request.ForceAttach
	}

	if !dara.IsNil(request.LoadBalancerConfigs) {
		query["LoadBalancerConfigs"] = request.LoadBalancerConfigs
	}

	if !dara.IsNil(request.LoadBalancers) {
		query["LoadBalancers"] = request.LoadBalancers
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachLoadBalancers"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachLoadBalancersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the AttachServerGroups API to attach one or more load balancing server groups to a scaling group. This applies to scenarios that require dynamic instance scaling to handle fluctuating workloads and ensure application programs remain active with high availability. Supported server group types include Application Load Balancer (ALB), Network Load Balancer (NLB), and Gateway Load Balancer (GWLB). This enables automatic elastic extension, allows you to dynamically allocate traffic based on load conditions, and improves the stability of application programs.
//
// @param request - AttachServerGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachServerGroupsResponse
func (client *Client) AttachServerGroupsWithContext(ctx context.Context, request *AttachServerGroupsRequest, runtime *dara.RuntimeOptions) (_result *AttachServerGroupsResponse, _err error) {
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

	if !dara.IsNil(request.ForceAttach) {
		query["ForceAttach"] = request.ForceAttach
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.ServerGroups) {
		query["ServerGroups"] = request.ServerGroups
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachServerGroups"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachServerGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more vServer groups of a Classic Load Balancer (CLB) instance to a scaling group. After a scaling group is associated with a CLB instance, instances added to the scaling group are automatically added as backend servers of the CLB instance to process access requests forwarded by the CLB instance. To forward different requests to different backend servers or forward requests based on domain names and URLs, you can call the AttachVServerGroups API operation to add one or more vServer groups of the CLB instance and manage different backend servers by using the vServer groups.
//
// Description:
//
// - To add vServer groups to a scaling group, the following conditions must be met:
//
//   - The CLB instance and the scaling group must belong to the same account.
//
//   - The CLB instance and the scaling group must be in the same region.
//
//   - The CLB instance must be in the Running state.
//
//   - The CLB instance must be configured with at least one listener and health checks must be enabled.
//
//   - If the network types of both the CLB instance and the scaling group are VPC, they must be in the same VPC.
//
//   - If the network type of the scaling group is VPC and the network type of the CLB instance is classic network, and the vServer group contains VPC-connected instances, the instances must be in the same VPC as the scaling group.
//
//   - The vServer groups to be added must belong to the corresponding CLB instance.
//
//   - The number of vServer groups in the scaling group does not exceed the quota. For more information about the quota, see [Limits](https://help.aliyun.com/document_detail/25863.html).
//
// - When you add vServer groups to a scaling group, you must specify the following properties:
//
//   - CLB instance ID (LoadBalancerId)
//
//   - vServer group ID (VServerGroupId)
//
//   - vServer group port number (Port)
//
//     >If you add the same vServer group to a scaling group by using different port numbers, the system considers that multiple vServer groups are added to the scaling group. If both the vServer group ID and the port number in the request parameters are duplicated, the system uses the first configured group by default and ignores the other vServer groups.
//
// @param request - AttachVServerGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachVServerGroupsResponse
func (client *Client) AttachVServerGroupsWithContext(ctx context.Context, request *AttachVServerGroupsRequest, runtime *dara.RuntimeOptions) (_result *AttachVServerGroupsResponse, _err error) {
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

	if !dara.IsNil(request.ForceAttach) {
		query["ForceAttach"] = request.ForceAttach
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.VServerGroups) {
		query["VServerGroups"] = request.VServerGroups
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachVServerGroups"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachVServerGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an instance refresh task and retains new instances that have been replaced and scaled out.
//
// Description:
//
// - You cannot cancel an instance refresh task that is being rolled back.
//
// @param request - CancelInstanceRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelInstanceRefreshResponse
func (client *Client) CancelInstanceRefreshWithContext(ctx context.Context, request *CancelInstanceRefreshRequest, runtime *dara.RuntimeOptions) (_result *CancelInstanceRefreshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceRefreshTaskId) {
		query["InstanceRefreshTaskId"] = request.InstanceRefreshTaskId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelInstanceRefresh"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelInstanceRefreshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the zone fault degradation for a scaling group.
//
// @param request - CancelScalingGroupZoneShiftRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelScalingGroupZoneShiftResponse
func (client *Client) CancelScalingGroupZoneShiftWithContext(ctx context.Context, request *CancelScalingGroupZoneShiftRequest, runtime *dara.RuntimeOptions) (_result *CancelScalingGroupZoneShiftResponse, _err error) {
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelScalingGroupZoneShift"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelScalingGroupZoneShiftResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a scaling group belongs. Resource groups allow you to group your cloud resources, such as scaling groups, by purpose, permissions, and ownership to implement hierarchical resource management for multiple users and projects within an enterprise. You can call the ChangeResourceGroup operation to move a scaling group from one resource group to another. This allows you to centrally monitor and manage the scaling group in the new resource group, eliminating the need to check resources across multiple Alibaba Cloud services.
//
// Description:
//
// - A resource is a cloud service entity that you create in Alibaba Cloud. For example, a scaling group is a resource.
//
// - A resource group is a mechanism for grouping and managing resources within an Alibaba Cloud account. It helps you resolve the complexities of resource grouping and authorization management within a single cloud account. For more information, see [What is Resource Management](https://help.aliyun.com/document_detail/94475.html).
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
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2022-02-22"),
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
// Ends the wait state of a scaling activity in advance to execute the next action if a lifecycle hook is configured for a scaling group.
//
// Description:
//
// Allows you to specify the next action after the wait state ends: continue the scaling activity (CONTINUE), abandon the scaling activity (ABANDON), or perform a rollback (ROLLBACK).
//
// @param request - CompleteLifecycleActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteLifecycleActionResponse
func (client *Client) CompleteLifecycleActionWithContext(ctx context.Context, request *CompleteLifecycleActionRequest, runtime *dara.RuntimeOptions) (_result *CompleteLifecycleActionResponse, _err error) {
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

	if !dara.IsNil(request.LifecycleActionResult) {
		query["LifecycleActionResult"] = request.LifecycleActionResult
	}

	if !dara.IsNil(request.LifecycleActionToken) {
		query["LifecycleActionToken"] = request.LifecycleActionToken
	}

	if !dara.IsNil(request.LifecycleHookId) {
		query["LifecycleHookId"] = request.LifecycleHookId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompleteLifecycleAction"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompleteLifecycleActionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// For workloads with bursty or unpredictable traffic patterns, call the CreateAlarm operation to create an event-triggered task. After the task is associated with CloudMonitor metrics, you can dynamically adjust the number of ECS or ECI instances in a scaling group and monitor instance metrics in real time to help you better manage and maintain cloud resources.
//
// Description:
//
// - If you create an event-triggered task of the custom monitoring type (MetricType=custom), you must first report custom metric data to CloudMonitor. For more information, see [Custom monitoring event-triggered tasks](https://help.aliyun.com/document_detail/74861.html).
//
// - When you create an event-triggered task, you must specify a metric (MetricName) and configure dimension information (Dimensions.DimensionKey and Dimensions.DimensionValue) to determine the aggregation scope of the metric data in the scaling group. For example, you can specify user_id and scaling_group to aggregate the metric data of all ECS or ECI instances for a specific user in a specific scaling group.
//
//   - The metrics for custom monitoring event-triggered tasks (MetricType=custom) depend on your custom metrics.
//
//   - For more information about the metrics supported by system monitoring event-triggered tasks (MetricType=system), see [System monitoring event-triggered tasks](https://help.aliyun.com/document_detail/74854.html).
//
// > In the dimension information, user_id and scaling_group are automatically populated by the system, while device and state must be manually specified. For more information, see the Dimensions.DimensionKey and Dimensions.DimensionValue parameters.
//
// @param request - CreateAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlarmResponse
func (client *Client) CreateAlarmWithContext(ctx context.Context, request *CreateAlarmRequest, runtime *dara.RuntimeOptions) (_result *CreateAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmActions) {
		query["AlarmActions"] = request.AlarmActions
	}

	if !dara.IsNil(request.ComparisonOperator) {
		query["ComparisonOperator"] = request.ComparisonOperator
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.Effective) {
		query["Effective"] = request.Effective
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.Expressions) {
		query["Expressions"] = request.Expressions
	}

	if !dara.IsNil(request.ExpressionsLogicOperator) {
		query["ExpressionsLogicOperator"] = request.ExpressionsLogicOperator
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.Statistics) {
		query["Statistics"] = request.Statistics
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlarm"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAlarmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a diagnostic report.
//
// @param request - CreateDiagnoseReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiagnoseReportResponse
func (client *Client) CreateDiagnoseReportWithContext(ctx context.Context, request *CreateDiagnoseReportRequest, runtime *dara.RuntimeOptions) (_result *CreateDiagnoseReportResponse, _err error) {
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
		Action:      dara.String("CreateDiagnoseReport"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDiagnoseReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Elastic Container Instance (ECI) scaling configuration. Auto Scaling uses this scaling configuration as a template for ECI instances and performs automatic creation of ECI instances to meet your business needs when a scale-out activity is triggered.
//
// Description:
//
// The scaling configuration specifies the template for ECI instances used during scale-out activities.
//
// The Cpu and Memory parameters specify the range of vCPUs and memory to define the instance type. Auto Scaling determines the available instance types based on factors such as I/O optimization and zone, and creates the most cost-effective instance based on price sorting. This method applies only when the scaling policy of the scaling group is set to cost optimization and no instance type is specified in the scaling configuration.
//
// @param request - CreateEciScalingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEciScalingConfigurationResponse
func (client *Client) CreateEciScalingConfigurationWithContext(ctx context.Context, request *CreateEciScalingConfigurationRequest, runtime *dara.RuntimeOptions) (_result *CreateEciScalingConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcrRegistryInfos) {
		query["AcrRegistryInfos"] = request.AcrRegistryInfos
	}

	if !dara.IsNil(request.ActiveDeadlineSeconds) {
		query["ActiveDeadlineSeconds"] = request.ActiveDeadlineSeconds
	}

	if !dara.IsNil(request.AutoCreateEip) {
		query["AutoCreateEip"] = request.AutoCreateEip
	}

	if !dara.IsNil(request.AutoMatchImageCache) {
		query["AutoMatchImageCache"] = request.AutoMatchImageCache
	}

	if !dara.IsNil(request.ContainerGroupName) {
		query["ContainerGroupName"] = request.ContainerGroupName
	}

	if !dara.IsNil(request.Containers) {
		query["Containers"] = request.Containers
	}

	if !dara.IsNil(request.CostOptimization) {
		query["CostOptimization"] = request.CostOptimization
	}

	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.CpuOptionsCore) {
		query["CpuOptionsCore"] = request.CpuOptionsCore
	}

	if !dara.IsNil(request.CpuOptionsThreadsPerCore) {
		query["CpuOptionsThreadsPerCore"] = request.CpuOptionsThreadsPerCore
	}

	if !dara.IsNil(request.DataCacheBucket) {
		query["DataCacheBucket"] = request.DataCacheBucket
	}

	if !dara.IsNil(request.DataCacheBurstingEnabled) {
		query["DataCacheBurstingEnabled"] = request.DataCacheBurstingEnabled
	}

	if !dara.IsNil(request.DataCachePL) {
		query["DataCachePL"] = request.DataCachePL
	}

	if !dara.IsNil(request.DataCacheProvisionedIops) {
		query["DataCacheProvisionedIops"] = request.DataCacheProvisionedIops
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DnsConfigNameServers) {
		query["DnsConfigNameServers"] = request.DnsConfigNameServers
	}

	if !dara.IsNil(request.DnsConfigOptions) {
		query["DnsConfigOptions"] = request.DnsConfigOptions
	}

	if !dara.IsNil(request.DnsConfigSearchs) {
		query["DnsConfigSearchs"] = request.DnsConfigSearchs
	}

	if !dara.IsNil(request.DnsPolicy) {
		query["DnsPolicy"] = request.DnsPolicy
	}

	if !dara.IsNil(request.EgressBandwidth) {
		query["EgressBandwidth"] = request.EgressBandwidth
	}

	if !dara.IsNil(request.EipBandwidth) {
		query["EipBandwidth"] = request.EipBandwidth
	}

	if !dara.IsNil(request.EnableSls) {
		query["EnableSls"] = request.EnableSls
	}

	if !dara.IsNil(request.EphemeralStorage) {
		query["EphemeralStorage"] = request.EphemeralStorage
	}

	if !dara.IsNil(request.GpuDriverVersion) {
		query["GpuDriverVersion"] = request.GpuDriverVersion
	}

	if !dara.IsNil(request.HostAliases) {
		query["HostAliases"] = request.HostAliases
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.ImageRegistryCredentials) {
		query["ImageRegistryCredentials"] = request.ImageRegistryCredentials
	}

	if !dara.IsNil(request.ImageSnapshotId) {
		query["ImageSnapshotId"] = request.ImageSnapshotId
	}

	if !dara.IsNil(request.IngressBandwidth) {
		query["IngressBandwidth"] = request.IngressBandwidth
	}

	if !dara.IsNil(request.InitContainers) {
		query["InitContainers"] = request.InitContainers
	}

	if !dara.IsNil(request.InstanceFamilyLevel) {
		query["InstanceFamilyLevel"] = request.InstanceFamilyLevel
	}

	if !dara.IsNil(request.InstanceTypes) {
		query["InstanceTypes"] = request.InstanceTypes
	}

	if !dara.IsNil(request.Ipv6AddressCount) {
		query["Ipv6AddressCount"] = request.Ipv6AddressCount
	}

	if !dara.IsNil(request.LoadBalancerWeight) {
		query["LoadBalancerWeight"] = request.LoadBalancerWeight
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.NtpServers) {
		query["NtpServers"] = request.NtpServers
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RamRoleName) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.RestartPolicy) {
		query["RestartPolicy"] = request.RestartPolicy
	}

	if !dara.IsNil(request.ScalingConfigurationName) {
		query["ScalingConfigurationName"] = request.ScalingConfigurationName
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.SecurityContextSysctls) {
		query["SecurityContextSysctls"] = request.SecurityContextSysctls
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SpotPriceLimit) {
		query["SpotPriceLimit"] = request.SpotPriceLimit
	}

	if !dara.IsNil(request.SpotStrategy) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TerminationGracePeriodSeconds) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	if !dara.IsNil(request.Volumes) {
		query["Volumes"] = request.Volumes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEciScalingConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEciScalingConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The CreateLifecycleHook API creates one or more lifecycle hooks in a specified scaling group. These hooks intervene at critical stages of the instance lifecycle, such as before startup or termination, and allow you to perform custom operations, such as sending notifications or executing automation scripts. This feature enables finer-grained control and management. For example, you can perform validation, configuration, or data backup tasks before and after instances are affected by scaling activities, which enhances the flexibility and reliability of application deployments.
//
// Description:
//
// You can create up to 10 lifecycle hooks in a scaling group. When a scaling activity occurs in a scaling group that has lifecycle hooks, the lifecycle hooks suspend the scaling activity for a period of time. You can specify the wait time by using the HeartbeatTimeout parameter. During the wait time, you can perform operations such as initializing the configurations of ECS instances and obtaining ECS instance data.
//
// During a scale-out activity, ECS instances enter the wait state after they are added to the whitelist of an ApsaraDB RDS instance, and are added to the backend server group of a Server Load Balancer (SLB) instance after the wait state ends. During a scale-in activity, ECS instances enter the wait state after they are removed from the backend server group of the SLB instance, and are removed from the whitelist of the ApsaraDB RDS instance after the wait state ends.
//
// You can configure notification methods for lifecycle hooks. Supported notification methods include Message Service (MNS) topics, MNS queues, and Operation Orchestration Service (OOS) templates. If you configure an OOS template, you must create a Resource Access Management (RAM) role for OOS. For more information, see [Grant RAM permissions to OOS](https://help.aliyun.com/document_detail/120810.html).
//
// > If the scaling group contains existing ECS instances and the configured OOS template is used to add or remove instances from the whitelist of a cloud database other than ApsaraDB RDS, you must manually add the existing ECS instances to the whitelist of the cloud database.
//
// @param request - CreateLifecycleHookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLifecycleHookResponse
func (client *Client) CreateLifecycleHookWithContext(ctx context.Context, request *CreateLifecycleHookRequest, runtime *dara.RuntimeOptions) (_result *CreateLifecycleHookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultResult) {
		query["DefaultResult"] = request.DefaultResult
	}

	if !dara.IsNil(request.HeartbeatTimeout) {
		query["HeartbeatTimeout"] = request.HeartbeatTimeout
	}

	if !dara.IsNil(request.LifecycleActionSources) {
		query["LifecycleActionSources"] = request.LifecycleActionSources
	}

	if !dara.IsNil(request.LifecycleHookName) {
		query["LifecycleHookName"] = request.LifecycleHookName
	}

	if !dara.IsNil(request.LifecycleTransition) {
		query["LifecycleTransition"] = request.LifecycleTransition
	}

	if !dara.IsNil(request.NotificationArn) {
		query["NotificationArn"] = request.NotificationArn
	}

	if !dara.IsNil(request.NotificationMetadata) {
		query["NotificationMetadata"] = request.NotificationMetadata
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLifecycleHook"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLifecycleHookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a notification configuration for Auto Scaling events and resource changes to help you track scaling group dynamics and automate the management of Auto Scaling events.
//
// Description:
//
// - You can configure the notification Settings to receive messages from cloud monitoring system events, Simple Message Queue (formerly MNS) queues, or Simple Message Queue (formerly MNS) topics. When a scaling event or resource change of a specified type occurs in a scaling group, Auto Scaling notifies cloud monitoring or Simple Message Queue (formerly MNS).
//
// - Different event notifications within the same scaling group cannot use the same destination.
//
//	For example, different event notifications cannot use the same cloud monitoring system event, the same MNS topic, or the same MNS queue at the same time.
//
// @param request - CreateNotificationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNotificationConfigurationResponse
func (client *Client) CreateNotificationConfigurationWithContext(ctx context.Context, request *CreateNotificationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *CreateNotificationConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MessageEncoding) {
		query["MessageEncoding"] = request.MessageEncoding
	}

	if !dara.IsNil(request.NotificationArn) {
		query["NotificationArn"] = request.NotificationArn
	}

	if !dara.IsNil(request.NotificationTypes) {
		query["NotificationTypes"] = request.NotificationTypes
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.TimeZone) {
		query["TimeZone"] = request.TimeZone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNotificationConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNotificationConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ECS scaling configuration for a scaling group.
//
// Description:
//
// A scaling configuration specifies the ECS instance template used during scale-out. The following instance configuration methods are supported:
//
// - InstancePatternInfos: the intelligent configuration mode. You do not need to specify exact network types. You only need to specify the number of vCPUs, memory size, instance family level, and maximum price based on your business requirements. The system automatically filters network types that meet the requirements and preferentially selects the network type with the lowest price to create ECS instances. This mode applies only to scaling groups in a VPC and effectively reduces the chance of insufficient network type inventory during scale-out.
//
// - InstanceType: specifies a single instance type.
//
// - InstanceTypes: specifies multiple instance types.
//
// - InstanceTypeOverrides: specifies multiple instance types and configures weights for different instance types.
//
// - Cpu and Memory: specifies the range of instance types by defining the number of vCPUs and memory size. Auto Scaling determines the available instance types based on factors such as I/O optimization and zone, and creates instances with the lowest price. This method applies only when the scaling policy of the scaling group is set to cost optimization and no instance type is specified in the scaling configuration.
//
// - You can use one of the following methods to bind the primary ENI. Note that you can use only one method to configure the primary ENI in a single call. Using both methods at the same time causes the call to fail and returns an error message.
//
//   - Directly set the configuration information of the primary ENI by using parameters such as SecurityGroupId, SecurityGroupIds, and Ipv6AddressCount.
//
//   - Set the configuration information of the primary ENI and secondary ENIs by using NetworkInterfaces. When you use the NetworkInterface parameter, you must set the primary ENI information. If NetworkInterface.InstanceType is set to Primary, it indicates the primary ENI. If NetworkInterface.InstanceType is set to Secondary or left empty, it indicates a secondary ENI.
//
// > You cannot specify InstanceType, InstanceTypes, InstanceTypeOverrides, and Cpu and Memory at the same time. However, you can specify InstanceType or InstanceTypes together with InstancePatternInfos. When InstanceType or InstanceTypes is specified together with InstancePatternInfos, Auto Scaling preferentially uses the specified instance types for scale-out. If the specified instance types are out of stock, Auto Scaling uses the instance types that match the InstancePatternInfos configuration.
//
// @param tmpReq - CreateScalingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScalingConfigurationResponse
func (client *Client) CreateScalingConfigurationWithContext(ctx context.Context, tmpReq *CreateScalingConfigurationRequest, runtime *dara.RuntimeOptions) (_result *CreateScalingConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateScalingConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SchedulerOptions) {
		request.SchedulerOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SchedulerOptions, dara.String("SchedulerOptions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Affinity) {
		query["Affinity"] = request.Affinity
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.CreditSpecification) {
		query["CreditSpecification"] = request.CreditSpecification
	}

	if !dara.IsNil(request.CustomPriorities) {
		query["CustomPriorities"] = request.CustomPriorities
	}

	if !dara.IsNil(request.DataDisks) {
		query["DataDisks"] = request.DataDisks
	}

	if !dara.IsNil(request.DedicatedHostClusterId) {
		query["DedicatedHostClusterId"] = request.DedicatedHostClusterId
	}

	if !dara.IsNil(request.DedicatedHostId) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.DeploymentSetId) {
		query["DeploymentSetId"] = request.DeploymentSetId
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.HpcClusterId) {
		query["HpcClusterId"] = request.HpcClusterId
	}

	if !dara.IsNil(request.HttpEndpoint) {
		query["HttpEndpoint"] = request.HttpEndpoint
	}

	if !dara.IsNil(request.HttpTokens) {
		query["HttpTokens"] = request.HttpTokens
	}

	if !dara.IsNil(request.ImageFamily) {
		query["ImageFamily"] = request.ImageFamily
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.InstanceDescription) {
		query["InstanceDescription"] = request.InstanceDescription
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstancePatternInfos) {
		query["InstancePatternInfos"] = request.InstancePatternInfos
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InstanceTypeCandidateOptions) {
		query["InstanceTypeCandidateOptions"] = request.InstanceTypeCandidateOptions
	}

	if !dara.IsNil(request.InstanceTypeOverrides) {
		query["InstanceTypeOverrides"] = request.InstanceTypeOverrides
	}

	if !dara.IsNil(request.InstanceTypes) {
		query["InstanceTypes"] = request.InstanceTypes
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.InternetMaxBandwidthIn) {
		query["InternetMaxBandwidthIn"] = request.InternetMaxBandwidthIn
	}

	if !dara.IsNil(request.InternetMaxBandwidthOut) {
		query["InternetMaxBandwidthOut"] = request.InternetMaxBandwidthOut
	}

	if !dara.IsNil(request.IoOptimized) {
		query["IoOptimized"] = request.IoOptimized
	}

	if !dara.IsNil(request.Ipv6AddressCount) {
		query["Ipv6AddressCount"] = request.Ipv6AddressCount
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.LoadBalancerWeight) {
		query["LoadBalancerWeight"] = request.LoadBalancerWeight
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.NetworkInterfaces) {
		query["NetworkInterfaces"] = request.NetworkInterfaces
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PasswordInherit) {
		query["PasswordInherit"] = request.PasswordInherit
	}

	if !dara.IsNil(request.RamRoleName) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourcePoolOptions) {
		query["ResourcePoolOptions"] = request.ResourcePoolOptions
	}

	if !dara.IsNil(request.ScalingConfigurationName) {
		query["ScalingConfigurationName"] = request.ScalingConfigurationName
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.SchedulerOptionsShrink) {
		query["SchedulerOptions"] = request.SchedulerOptionsShrink
	}

	if !dara.IsNil(request.SecondaryEniVSwitchConstraint) {
		query["SecondaryEniVSwitchConstraint"] = request.SecondaryEniVSwitchConstraint
	}

	if !dara.IsNil(request.SecurityEnhancementStrategy) {
		query["SecurityEnhancementStrategy"] = request.SecurityEnhancementStrategy
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityGroupIds) {
		query["SecurityGroupIds"] = request.SecurityGroupIds
	}

	if !dara.IsNil(request.SecurityOptions) {
		query["SecurityOptions"] = request.SecurityOptions
	}

	if !dara.IsNil(request.SpotDuration) {
		query["SpotDuration"] = request.SpotDuration
	}

	if !dara.IsNil(request.SpotInterruptionBehavior) {
		query["SpotInterruptionBehavior"] = request.SpotInterruptionBehavior
	}

	if !dara.IsNil(request.SpotPriceLimits) {
		query["SpotPriceLimits"] = request.SpotPriceLimits
	}

	if !dara.IsNil(request.SpotStrategy) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !dara.IsNil(request.StorageSetId) {
		query["StorageSetId"] = request.StorageSetId
	}

	if !dara.IsNil(request.StorageSetPartitionNumber) {
		query["StorageSetPartitionNumber"] = request.StorageSetPartitionNumber
	}

	if !dara.IsNil(request.SystemDiskCategories) {
		query["SystemDiskCategories"] = request.SystemDiskCategories
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Tenancy) {
		query["Tenancy"] = request.Tenancy
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	if !dara.IsNil(request.ImageOptions) {
		query["ImageOptions"] = request.ImageOptions
	}

	if !dara.IsNil(request.PrivatePoolOptions) {
		query["PrivatePoolOptions"] = request.PrivatePoolOptions
	}

	if !dara.IsNil(request.SystemDisk) {
		query["SystemDisk"] = request.SystemDisk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScalingConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScalingConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scaling group to automatically adjust your computing capacity by adding or removing instances based on business demands and policies.
//
// Description:
//
// A scaling group is a collection of ECS instances that share the same application scenarios.
//
// The number of scaling groups that you can create in a region depends on your Auto Scaling usage. Go to Quota Center to check the quota for the total number of scaling groups.
//
// A scaling group does not take effect immediately after it is created. You must invoke the [EnableScalingGroup](https://help.aliyun.com/document_detail/25939.html) operation to enable the scaling group before it can trigger scaling activities and execute scaling rules.
//
// The scaling group, the associated Classic Load Balancer (CLB) instances, and the associated ApsaraDB RDS instances must reside in the same region. For more information, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
//
// If you associate CLB instances with the scaling group, Auto Scaling automatically adds the ECS instances in the scaling group to the backend server groups of the CLB instances. You can specify the server groups to which the ECS instances are added. The following types of server groups are supported:
//
// - Default server group: contains ECS instances that accept frontend requests. If a listener is not configured with a virtual server group or a primary/secondary server group, requests are forwarded to the ECS instances in the default server group by default.
//
// - Virtual server group: used when you need to forward different requests to different backend servers, or forward requests based on domain names and URLs.
//
// > If you specify both the default server group and multiple virtual server groups, the ECS instances are added to all these server groups.
//
// After an instance is added to a backend server group of a CLB instance, the default weight of the instance is 50. The CLB instance must meet the following conditions:
//
// - The CLB instance must be in the Active state. You can call the [DescribeLoadBalancers](https://help.aliyun.com/document_detail/2401696.html) operation to check the status of the CLB instance.
//
// - Health checks must be enabled for all listeners configured on the CLB instance. Otherwise, the scaling group fails to be created.
//
// If you associate Application Load Balancer (ALB) server groups, Network Load Balancer (NLB) server groups, or Gateway Load Balancer (GWLB) server groups with the scaling group, Auto Scaling automatically adds the ECS instances in the scaling group as backend servers to the ALB, NLB, or GWLB server groups to process access requests distributed by the ALB, NLB, or GWLB instances. You can specify multiple ALB, NLB, or GWLB server groups, but the server groups must belong to the same VPC as the scaling group. For more information, see [AttachAlbServerGroups](https://help.aliyun.com/document_detail/266800.html) or [AttachServerGroups](https://help.aliyun.com/document_detail/600535.html).
//
// If you associate ApsaraDB RDS instances with the scaling group, Auto Scaling automatically adds the internal IP addresses of the ECS instances in the scaling group to the whitelists of the ApsaraDB RDS instances. The ApsaraDB RDS instances must meet the following conditions:
//
// - The ApsaraDB RDS instance must be in the Running state. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/610396.html) operation to check the status of the ApsaraDB RDS instance.
//
// - The number of IP addresses in the whitelist of the ApsaraDB RDS instance cannot exceed the upper limit. For more information, see [Configure a whitelist](https://help.aliyun.com/document_detail/43185.html) in the ApsaraDB RDS documentation.
//
// If the MultiAZPolicy parameter of the scaling group is set to COST_OPTIMIZED:
//
// - If the OnDemandBaseCapacity, OnDemandPercentageAboveBaseCapacity, and SpotInstancePools parameters are specified, the instance allocation method under the cost optimization policy is specified. This allocation method is prioritized during scale-out and scale-in.
//
// - If the OnDemandBaseCapacity, OnDemandPercentageAboveBaseCapacity, or SpotInstancePools parameter is not specified, instances are created based on the lowest cost under the cost optimization policy. In this mode, you cannot create instances by using elasticity assurances or capacity reservations, and you cannot specify the vSwitch to which secondary elastic network interfaces (ENIs) belong.
//
// If you configure propagatable tags for the scaling group, which means `Tags.Propagate` is set to true:
//
// - Tags on the scaling group are propagated only to newly created instances, not to existing instances in the scaling group.
//
// - If you specify instance tags in the scaling configuration and choose to propagate the tags of the scaling group to the instances, all tags coexist.
//
// - If a tag in the scaling configuration and a propagatable tag in the scaling group have the same tag key, the tag value in the scaling configuration takes precedence.
//
// @param request - CreateScalingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScalingGroupResponse
func (client *Client) CreateScalingGroupWithContext(ctx context.Context, request *CreateScalingGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateScalingGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlbServerGroups) {
		query["AlbServerGroups"] = request.AlbServerGroups
	}

	if !dara.IsNil(request.AllocationStrategy) {
		query["AllocationStrategy"] = request.AllocationStrategy
	}

	if !dara.IsNil(request.AutoRebalance) {
		query["AutoRebalance"] = request.AutoRebalance
	}

	if !dara.IsNil(request.AzBalance) {
		query["AzBalance"] = request.AzBalance
	}

	if !dara.IsNil(request.BalanceMode) {
		query["BalanceMode"] = request.BalanceMode
	}

	if !dara.IsNil(request.CapacityOptions) {
		query["CapacityOptions"] = request.CapacityOptions
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CompensateWithOnDemand) {
		query["CompensateWithOnDemand"] = request.CompensateWithOnDemand
	}

	if !dara.IsNil(request.ContainerGroupId) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !dara.IsNil(request.CustomPolicyARN) {
		query["CustomPolicyARN"] = request.CustomPolicyARN
	}

	if !dara.IsNil(request.DBInstanceIds) {
		query["DBInstanceIds"] = request.DBInstanceIds
	}

	if !dara.IsNil(request.DBInstances) {
		query["DBInstances"] = request.DBInstances
	}

	if !dara.IsNil(request.DefaultCooldown) {
		query["DefaultCooldown"] = request.DefaultCooldown
	}

	if !dara.IsNil(request.DesiredCapacity) {
		query["DesiredCapacity"] = request.DesiredCapacity
	}

	if !dara.IsNil(request.GroupDeletionProtection) {
		query["GroupDeletionProtection"] = request.GroupDeletionProtection
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.HealthCheckType) {
		query["HealthCheckType"] = request.HealthCheckType
	}

	if !dara.IsNil(request.HealthCheckTypes) {
		query["HealthCheckTypes"] = request.HealthCheckTypes
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LaunchTemplateId) {
		query["LaunchTemplateId"] = request.LaunchTemplateId
	}

	if !dara.IsNil(request.LaunchTemplateOverrides) {
		query["LaunchTemplateOverrides"] = request.LaunchTemplateOverrides
	}

	if !dara.IsNil(request.LaunchTemplateVersion) {
		query["LaunchTemplateVersion"] = request.LaunchTemplateVersion
	}

	if !dara.IsNil(request.LifecycleHooks) {
		query["LifecycleHooks"] = request.LifecycleHooks
	}

	if !dara.IsNil(request.LoadBalancerConfigs) {
		query["LoadBalancerConfigs"] = request.LoadBalancerConfigs
	}

	if !dara.IsNil(request.LoadBalancerIds) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !dara.IsNil(request.MaxInstanceLifetime) {
		query["MaxInstanceLifetime"] = request.MaxInstanceLifetime
	}

	if !dara.IsNil(request.MaxSize) {
		query["MaxSize"] = request.MaxSize
	}

	if !dara.IsNil(request.MinSize) {
		query["MinSize"] = request.MinSize
	}

	if !dara.IsNil(request.MultiAZPolicy) {
		query["MultiAZPolicy"] = request.MultiAZPolicy
	}

	if !dara.IsNil(request.OnDemandBaseCapacity) {
		query["OnDemandBaseCapacity"] = request.OnDemandBaseCapacity
	}

	if !dara.IsNil(request.OnDemandPercentageAboveBaseCapacity) {
		query["OnDemandPercentageAboveBaseCapacity"] = request.OnDemandPercentageAboveBaseCapacity
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemovalPolicies) {
		query["RemovalPolicies"] = request.RemovalPolicies
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupName) {
		query["ScalingGroupName"] = request.ScalingGroupName
	}

	if !dara.IsNil(request.ScalingPolicy) {
		query["ScalingPolicy"] = request.ScalingPolicy
	}

	if !dara.IsNil(request.ServerGroups) {
		query["ServerGroups"] = request.ServerGroups
	}

	if !dara.IsNil(request.SpotAllocationStrategy) {
		query["SpotAllocationStrategy"] = request.SpotAllocationStrategy
	}

	if !dara.IsNil(request.SpotInstancePools) {
		query["SpotInstancePools"] = request.SpotInstancePools
	}

	if !dara.IsNil(request.SpotInstanceRemedy) {
		query["SpotInstanceRemedy"] = request.SpotInstanceRemedy
	}

	if !dara.IsNil(request.StopInstanceTimeout) {
		query["StopInstanceTimeout"] = request.StopInstanceTimeout
	}

	if !dara.IsNil(request.SyncAlarmRuleToCms) {
		query["SyncAlarmRuleToCms"] = request.SyncAlarmRuleToCms
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.VServerGroups) {
		query["VServerGroups"] = request.VServerGroups
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	if !dara.IsNil(request.ZoneShiftPolicy) {
		query["ZoneShiftPolicy"] = request.ZoneShiftPolicy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScalingGroup"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScalingGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scaling rule. The purpose of a scaling rule is determined by its type and can be used to trigger a scaling activity or intelligently configure the boundary values of a scaling group. You can create different types of scaling rules based on your business requirements. For example, if you only need to configure the boundary values of a scaling group, select the predictive scaling rule type.
//
// Description:
//
// A scaling rule defines specific scale-out or scale-in operations, such as adding or removing N instances. If executing a scaling rule causes the number of ECS or ECI instances in the scaling group to fall below MinSize or exceed MaxSize, Auto Scaling automatically adjusts the number of ECS or ECI instances to be added or removed so that the total number of instances reaches MinSize or MaxSize. However, the configured value of the scaling rule remains unchanged. Examples:
//
// - For a scaling group with MaxSize set to 3 and a current total capacity of 2, if the scaling rule specifies adding 3 ECS instances, only 1 ECS instance is added during actual execution, but the configured value of the scaling rule remains 3.
//
// - For a scaling group with MinSize set to 2 and a current total capacity of 3, if the scaling rule specifies removing 5 ECS instances, only 1 ECS instance is removed during actual execution, but the configured value of the scaling rule remains 5.
//
// Note the following parameter descriptions:
//
// - If AdjustmentType is set to TotalCapacity, the number of ECS or ECI instances in the scaling group is adjusted to the specified number. The corresponding AdjustmentValue must be greater than or equal to 0.
//
// - If AdjustmentType is set to QuantityChangeInCapacity or PercentChangeInCapacity, a positive AdjustmentValue indicates to increase instances, and a negative value indicates to decrease instances.
//
// - If AdjustmentType is set to PercentChangeInCapacity, Auto Scaling calculates the number of ECS instances or ECI instances to be added or removed by using the formula: Current total capacity × AdjustmentValue / 100, and uses the rounding principle to confirm the number of ECS instances or ECI instances to increase or decrease.
//
// - If a cooldown time (Cooldown) is specified for the scaling rule, the scaling group enters the cooldown state for the specified duration after the scaling activity triggered by the rule is completed. If no cooldown time is specified for the scaling rule, the default cooldown time (DefaultCooldown) of the scaling group is used.
//
// - The number of scaling rules that can be created in a scaling group is limited. For more information, see [Limits](https://help.aliyun.com/document_detail/25863.html).
//
// - The returned unique identifier of the scaling rule (ScalingRuleAri) can be used in the following API operations:
//
//   - Specify it in the ScalingRuleAri parameter of the ExecuteScalingRule operation to manually execute the scaling rule.
//
//   - Specify it in the ScheduledAction parameter of the CreateScheduledTask operation to execute the scaling rule on a schedule.
//
// @param request - CreateScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScalingRuleResponse
func (client *Client) CreateScalingRuleWithContext(ctx context.Context, request *CreateScalingRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdjustmentType) {
		query["AdjustmentType"] = request.AdjustmentType
	}

	if !dara.IsNil(request.AdjustmentValue) {
		query["AdjustmentValue"] = request.AdjustmentValue
	}

	if !dara.IsNil(request.AlarmDimensions) {
		query["AlarmDimensions"] = request.AlarmDimensions
	}

	if !dara.IsNil(request.AlarmOptions) {
		query["AlarmOptions"] = request.AlarmOptions
	}

	if !dara.IsNil(request.Cooldown) {
		query["Cooldown"] = request.Cooldown
	}

	if !dara.IsNil(request.DisableScaleIn) {
		query["DisableScaleIn"] = request.DisableScaleIn
	}

	if !dara.IsNil(request.EstimatedInstanceWarmup) {
		query["EstimatedInstanceWarmup"] = request.EstimatedInstanceWarmup
	}

	if !dara.IsNil(request.HybridMetrics) {
		query["HybridMetrics"] = request.HybridMetrics
	}

	if !dara.IsNil(request.HybridMonitorNamespace) {
		query["HybridMonitorNamespace"] = request.HybridMonitorNamespace
	}

	if !dara.IsNil(request.InitialMaxSize) {
		query["InitialMaxSize"] = request.InitialMaxSize
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.MinAdjustmentMagnitude) {
		query["MinAdjustmentMagnitude"] = request.MinAdjustmentMagnitude
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PredictiveScalingMode) {
		query["PredictiveScalingMode"] = request.PredictiveScalingMode
	}

	if !dara.IsNil(request.PredictiveTaskBufferTime) {
		query["PredictiveTaskBufferTime"] = request.PredictiveTaskBufferTime
	}

	if !dara.IsNil(request.PredictiveValueBehavior) {
		query["PredictiveValueBehavior"] = request.PredictiveValueBehavior
	}

	if !dara.IsNil(request.PredictiveValueBuffer) {
		query["PredictiveValueBuffer"] = request.PredictiveValueBuffer
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScaleInEvaluationCount) {
		query["ScaleInEvaluationCount"] = request.ScaleInEvaluationCount
	}

	if !dara.IsNil(request.ScaleOutEvaluationCount) {
		query["ScaleOutEvaluationCount"] = request.ScaleOutEvaluationCount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	if !dara.IsNil(request.ScalingRuleType) {
		query["ScalingRuleType"] = request.ScalingRuleType
	}

	if !dara.IsNil(request.StepAdjustments) {
		query["StepAdjustments"] = request.StepAdjustments
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScalingRule"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduled task. A scheduled task is a preset node. You can invoke the CreateScheduledTask API to create a scheduled task that executes a specified scaling rule at a specified time to automatically scale the compute resource, meeting business requirements while controlling costs.
//
// Description:
//
// - If a scheduled node has failed to trigger the execution of the scaling rule because a scaling activity is in progress or the scaling group is disabled, the scheduled node automatically retries within the `LaunchExpirationTime` period. Otherwise, the scheduled trigger is canceled.
//
// - If multiple scheduled nodes trigger scaling rules for the same scaling group at approximately the same time, the details of the triggered scaling activities are as follows:
//
//   - **Scaling groups with Settings for the expected number of instances**: Scaling activities triggered by multiple scheduled nodes are of the parallel scaling activity type. For example, if multiple scheduled nodes are executed at approximately the same time, the scaling group can trigger multiple parallel scaling activity executions simultaneously.
//
//   - **Scaling groups without Settings for the expected number of instances**: The earliest triggered scheduled node executes the scaling activity first. Because a scaling group can run only one scaling activity at a time, subsequent scheduled nodes automatically retry within the `LaunchExpirationTime` period. If a subsequent scheduled node is still retrying within the `LaunchExpirationTime` period after the previous scaling activity is completed, the scaling rule of the scheduled node is executed and the corresponding scaling activity is triggered.
//
// - Scheduled nodes support the following two scaling methods:
//
//   - Specify the scaling rule to be executed by using the ScheduledAction parameter.
//
//   - Specify the number of instances in the scaling group by using the ScalingGroupId parameter.
//
//     > You cannot specify both ScheduledAction and ScalingGroupId at the same time.
//
// @param request - CreateScheduledTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledTaskResponse
func (client *Client) CreateScheduledTaskWithContext(ctx context.Context, request *CreateScheduledTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateScheduledTaskResponse, _err error) {
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

	if !dara.IsNil(request.DesiredCapacity) {
		query["DesiredCapacity"] = request.DesiredCapacity
	}

	if !dara.IsNil(request.LaunchExpirationTime) {
		query["LaunchExpirationTime"] = request.LaunchExpirationTime
	}

	if !dara.IsNil(request.LaunchTime) {
		query["LaunchTime"] = request.LaunchTime
	}

	if !dara.IsNil(request.MaxValue) {
		query["MaxValue"] = request.MaxValue
	}

	if !dara.IsNil(request.MinValue) {
		query["MinValue"] = request.MinValue
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecurrenceEndTime) {
		query["RecurrenceEndTime"] = request.RecurrenceEndTime
	}

	if !dara.IsNil(request.RecurrenceType) {
		query["RecurrenceType"] = request.RecurrenceType
	}

	if !dara.IsNil(request.RecurrenceValue) {
		query["RecurrenceValue"] = request.RecurrenceValue
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.ScheduledAction) {
		query["ScheduledAction"] = request.ScheduledAction
	}

	if !dara.IsNil(request.ScheduledTaskName) {
		query["ScheduledTaskName"] = request.ScheduledTaskName
	}

	if !dara.IsNil(request.TaskEnabled) {
		query["TaskEnabled"] = request.TaskEnabled
	}

	if !dara.IsNil(request.WarmPoolCapacityOptions) {
		query["WarmPoolCapacityOptions"] = request.WarmPoolCapacityOptions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScheduledTask"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 使指定伸缩配置进入失效状态
//
// @param request - DeactivateScalingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactivateScalingConfigurationResponse
func (client *Client) DeactivateScalingConfigurationWithContext(ctx context.Context, request *DeactivateScalingConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DeactivateScalingConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingConfigurationId) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeactivateScalingConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeactivateScalingConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an event-triggered task. If your business pattern is unpredictable or prone to unforeseen traffic spikes, you can create event-triggered tasks by associating CloudMonitor metrics to effectively monitor fluctuations in your business workload. Upon detecting that the criteria for alerts, as specified in event-triggered tasks, are fulfilled, Auto Scaling promptly issues alerts and executes the scaling rules predefined within those tasks. This process occurs within the predefined effective time windows of the tasks, thereby facilitating the automatic increase or decrease of Elastic Compute Service (ECS) instances or elastic container instances within your scaling groups. Ultimately, this mechanism ensures the dynamic optimization of resources based on real-time workload demands. If you no longer need an event-triggered task, you can call the DeleteAlarm operation to delete it.
//
// @param request - DeleteAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlarmResponse
func (client *Client) DeleteAlarmWithContext(ctx context.Context, request *DeleteAlarmRequest, runtime *dara.RuntimeOptions) (_result *DeleteAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmTaskId) {
		query["AlarmTaskId"] = request.AlarmTaskId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlarm"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlarmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a health diagnosis report for a scaling group.
//
// @param request - DeleteDiagnoseReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDiagnoseReportResponse
func (client *Client) DeleteDiagnoseReportWithContext(ctx context.Context, request *DeleteDiagnoseReportRequest, runtime *dara.RuntimeOptions) (_result *DeleteDiagnoseReportResponse, _err error) {
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
		Action:      dara.String("DeleteDiagnoseReport"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDiagnoseReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scaling configuration of the Elastic Container Instance type. If the scaling configuration of a scaling group is in the Inactive state and the scaling group contains no elastic container instances created from the scaling configuration, you can call the DeleteEciScalingConfiguration operation to delete the scaling configuration to free up the scaling configuration quota.
//
// Description:
//
// You cannot call this operation to delete a scaling configuration in the following scenarios:
//
//   - The scaling configuration is in the Active state.
//
//   - The scaling group contains elastic container instances created from the scaling configuration.
//
// @param request - DeleteEciScalingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEciScalingConfigurationResponse
func (client *Client) DeleteEciScalingConfigurationWithContext(ctx context.Context, request *DeleteEciScalingConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DeleteEciScalingConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingConfigurationId) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEciScalingConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEciScalingConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a lifecycle hook. If you no longer require a lifecycle hook, you can call the DeleteLifecycleHook operation to delete it, which frees up the lifecycle hook quota.
//
// Description:
//
// If you delete an effective lifecycle hook before its timeout period ends, the instances on which the lifecycle hook takes effect exits the Pending state ahead of schedule. You can use the following methods to delete a lifecycle hook:
//
//   - Include `LifecycleHookId` within your request to specify the lifecycle hook that you want to delete. In this case, `ScalingGroupId` and `LifecycleHookName` are ignored.
//
//   - Include `ScalingGroupId` and `LifecycleHookName` within your request to specify the lifecycle hook that you want to delete.
//
// @param request - DeleteLifecycleHookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLifecycleHookResponse
func (client *Client) DeleteLifecycleHookWithContext(ctx context.Context, request *DeleteLifecycleHookRequest, runtime *dara.RuntimeOptions) (_result *DeleteLifecycleHookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LifecycleHookId) {
		query["LifecycleHookId"] = request.LifecycleHookId
	}

	if !dara.IsNil(request.LifecycleHookName) {
		query["LifecycleHookName"] = request.LifecycleHookName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLifecycleHook"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLifecycleHookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes event notification rules. The event notification feature facilitates efficient issue identification and event management by automatically forwarding notifications from Auto Scaling to designated endpoints such as CloudMonitor or Message Service (MNS) topics and queues. If you no longer require an event notification rule, you can call the DeleteNotificationConfiguration operation to delete it.
//
// @param request - DeleteNotificationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNotificationConfigurationResponse
func (client *Client) DeleteNotificationConfigurationWithContext(ctx context.Context, request *DeleteNotificationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DeleteNotificationConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NotificationArn) {
		query["NotificationArn"] = request.NotificationArn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNotificationConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNotificationConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scaling configuration that is used to create Elastic Compute Service (ECS) instances.
//
// Description:
//
// You cannot delete a scaling configuration in one of the following scenarios:
//
//   - The scaling configuration in your scaling group is in the Active state.
//
//   - The scaling group contains ECS instances that were created based on the scaling configuration.
//
// @param request - DeleteScalingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScalingConfigurationResponse
func (client *Client) DeleteScalingConfigurationWithContext(ctx context.Context, request *DeleteScalingConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DeleteScalingConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingConfigurationId) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScalingConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScalingConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scaling group. If you want to enable policy-based automatic addition or removal of instances of a specific type to meet evolving business requirements, you can create scaling groups to manage your computing power with ease. The computing power refers to the instances that provide the computing capability. If you no longer require a scaling group, you can call the DeleteScalingGroup operation to delete it to free up the scaling group quota.
//
// Description:
//
// Before you call the DeleteScalingGroup operation, take note of the following items:
//
//   - If you delete a scaling group, the scaling configurations, scaling rules, scaling activities, and scaling requests related to the scaling group are also deleted.
//
//   - If you delete a scaling group, the scheduled tasks and event-triggered tasks of the scaling group are not deleted. The Server Load Balancer (SLB) instances and ApsaraDB RDS instances that are attached to the scaling group are also not deleted.
//
//   - If the scaling group that you want to delete contains ECS instances or elastic container instances that are in the In Service state, Auto Scaling stops the instances and then removes all manually added instances from the scaling group or releases all automatically created instances in the scaling group before the scaling group is deleted.
//
//     **
//
//     **Note*	- Before you delete a scaling group, make sure that the Deletion Protection feature is disabled. If you have enabled the Deletion Protection feature for a scaling group, disable the feature on the Modify Scaling Group page before you delete the scaling group.
//
//   - If you do not disable the Deletion Protection feature for a scaling group, you cannot delete the scaling group by using the Auto Scaling console or calling this operation. The Deletion Protection feature is an effective measure to safeguard scaling groups against unintended deletion.
//
//   - Prior to deleting a scaling group, make sure that your ECS instances within the scaling group are safeguarded against unintended release. Even if you have already enabled the Release Protection feature for the ECS instances, you must manually put these ECS instances into the Protected state. Doing so guarantees that the ECS instances will not be forcibly released during the deletion process of the scaling group, providing an extra layer of security.
//
//     **
//
//     **Note*	- Before you delete a scaling group, we recommend that you enable the Deletion Protection feature for ECS instances that you want to retain. This action guarantees that the ECS instances are not forcibly released after you delete the scaling group. For more information, see [SetInstancesProtection](https://help.aliyun.com/document_detail/459342.html).
//
// @param request - DeleteScalingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScalingGroupResponse
func (client *Client) DeleteScalingGroupWithContext(ctx context.Context, request *DeleteScalingGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteScalingGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceDelete) {
		query["ForceDelete"] = request.ForceDelete
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScalingGroup"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScalingGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scaling rule.
//
// @param request - DeleteScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScalingRuleResponse
func (client *Client) DeleteScalingRuleWithContext(ctx context.Context, request *DeleteScalingRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingRuleId) {
		query["ScalingRuleId"] = request.ScalingRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScalingRule"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes scheduled tasks. For workloads with predictable patterns, you can create scheduled tasks to align with your business requirements and optimize resource utilization for cost savings. These tasks automatically ensure that sufficient computing resources are provisioned in anticipation of peak hours and efficiently release unused resources during off-peak hours, thereby streamlining operational efficiency and reducing expenses. If you no longer require a scheduled task, you can call the DeleteScheduledTask operation to delete it.
//
// @param request - DeleteScheduledTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduledTaskResponse
func (client *Client) DeleteScheduledTaskWithContext(ctx context.Context, request *DeleteScheduledTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteScheduledTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScheduledTaskId) {
		query["ScheduledTaskId"] = request.ScheduledTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScheduledTask"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScheduledTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the warmup pool of a scaling group and releases the instances in the warmup pool.
//
// @param request - DeleteWarmPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWarmPoolResponse
func (client *Client) DeleteWarmPoolWithContext(ctx context.Context, request *DeleteWarmPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteWarmPoolResponse, _err error) {
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWarmPool"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWarmPoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about event-triggered tasks. This helps you track the configurations and metric data trends of existing event-triggered tasks, detect system resource anomalies in real time, and ensure system stability and reliability.
//
// @param request - DescribeAlarmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlarmsResponse
func (client *Client) DescribeAlarmsWithContext(ctx context.Context, request *DescribeAlarmsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlarmsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmTaskId) {
		query["AlarmTaskId"] = request.AlarmTaskId
	}

	if !dara.IsNil(request.IsEnable) {
		query["IsEnable"] = request.IsEnable
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlarms"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlarmsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of scaling activities that trigger text message, internal message, or email-based notifications. When you call the DescribeAlertConfiguration operation, you must specify the scaling group ID and region ID to query the status of the desired scaling activities. A scaling activity can be in one of the following states: Successful, Failed, and Rejected.
//
// @param request - DescribeAlertConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertConfigurationResponse
func (client *Client) DescribeAlertConfigurationWithContext(ctx context.Context, request *DescribeAlertConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists candidate instance types.
//
// @param request - DescribeCandidateInstanceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCandidateInstanceTypeResponse
func (client *Client) DescribeCandidateInstanceTypeWithContext(ctx context.Context, request *DescribeCandidateInstanceTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeCandidateInstanceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowCrossAz) {
		query["AllowCrossAz"] = request.AllowCrossAz
	}

	if !dara.IsNil(request.AllowDifferentGeneration) {
		query["AllowDifferentGeneration"] = request.AllowDifferentGeneration
	}

	if !dara.IsNil(request.DataDiskCategories) {
		query["DataDiskCategories"] = request.DataDiskCategories
	}

	if !dara.IsNil(request.ImageFamily) {
		query["ImageFamily"] = request.ImageFamily
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.InstanceTypes) {
		query["InstanceTypes"] = request.InstanceTypes
	}

	if !dara.IsNil(request.Ipv6AddressCount) {
		query["Ipv6AddressCount"] = request.Ipv6AddressCount
	}

	if !dara.IsNil(request.MaxPrice) {
		query["MaxPrice"] = request.MaxPrice
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.SpotStrategy) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !dara.IsNil(request.SystemDiskCategories) {
		query["SystemDiskCategories"] = request.SystemDiskCategories
	}

	if !dara.IsNil(request.ZoneIds) {
		query["ZoneIds"] = request.ZoneIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCandidateInstanceType"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCandidateInstanceTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries health diagnostic reports.
//
// @param request - DescribeDiagnoseReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnoseReportsResponse
func (client *Client) DescribeDiagnoseReportsWithContext(ctx context.Context, request *DescribeDiagnoseReportsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnoseReportsResponse, _err error) {
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
		Action:      dara.String("DescribeDiagnoseReports"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiagnoseReportsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an Elastic Container Instance (ECI) scaling configuration based on the scaling configuration ID. The details include the temporary storage size, the number of IPv6 addresses, and the elastic IP address (EIP) bandwidth. The output supports the YAML format.
//
// @param request - DescribeEciScalingConfigurationDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEciScalingConfigurationDetailResponse
func (client *Client) DescribeEciScalingConfigurationDetailWithContext(ctx context.Context, request *DescribeEciScalingConfigurationDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeEciScalingConfigurationDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OutputFormat) {
		query["OutputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScalingConfigurationId) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEciScalingConfigurationDetail"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEciScalingConfigurationDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries ECI scaling configurations. You can invoke DescribeEciScalingConfigurations and specify parameters such as the scaling group ID, and the scaling configuration IDs or names. This parameter query retrieves information such as the instance preemption policy, container group instance types, container image pull policy, and the weight of ECI instances as backend servers.
//
// @param request - DescribeEciScalingConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEciScalingConfigurationsResponse
func (client *Client) DescribeEciScalingConfigurationsWithContext(ctx context.Context, request *DescribeEciScalingConfigurationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEciScalingConfigurationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScalingConfigurationIds) {
		query["ScalingConfigurationIds"] = request.ScalingConfigurationIds
	}

	if !dara.IsNil(request.ScalingConfigurationNames) {
		query["ScalingConfigurationNames"] = request.ScalingConfigurationNames
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEciScalingConfigurations"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEciScalingConfigurationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A scaling group\\"s configuration is key to successful scale-outs. Auto Scaling uses elasticity strength to quantify the effectiveness of a Scaling Configuration, helping you make targeted optimizations.
//
// @param request - DescribeElasticStrengthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticStrengthResponse
func (client *Client) DescribeElasticStrengthWithContext(ctx context.Context, request *DescribeElasticStrengthRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticStrengthResponse, _err error) {
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
		Action:      dara.String("DescribeElasticStrength"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeElasticStrengthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Instance creation and deletion statistics.
//
// @param request - DescribeInstanceCreateAndDeleteStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceCreateAndDeleteStatisticsResponse
func (client *Client) DescribeInstanceCreateAndDeleteStatisticsWithContext(ctx context.Context, request *DescribeInstanceCreateAndDeleteStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceCreateAndDeleteStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceCreateAndDeleteStatistics"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceCreateAndDeleteStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instance refresh tasks of a scaling group to view basic information and task progress.
//
// @param request - DescribeInstanceRefreshesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceRefreshesResponse
func (client *Client) DescribeInstanceRefreshesWithContext(ctx context.Context, request *DescribeInstanceRefreshesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceRefreshesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceRefreshTaskIds) {
		query["InstanceRefreshTaskIds"] = request.InstanceRefreshTaskIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceRefreshes"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceRefreshesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries lifecycle hook actions. When you call the DescribeLifecycleActions operation, you can specify parameters such as ScalingActivityId, NextToken, and MaxResults to query the details such as the action status and ID of a lifecycle hook. You can also call this operation to query the IDs of Elastic Compute Service (ECS) instances on which the lifecycle hook takes effect.
//
// Description:
//
// If a scaling activity is executed and a lifecycle hook is created for the scaling activity, the lifecycle hook triggers a lifecycle action. A lifecycle action can be in one of the following states:
//
//   - If a lifecycle action is in the Pending state, Elastic Compute Service (ECS) instances are waiting to be added to a scaling group or waiting to be removed from a scaling group.
//
//   - If a lifecycle action is in the Timeout state, the lifecycle hook that triggers the lifecycle action expires and ECS instances are added to or removed from a scaling group.
//
//   - If a lifecycle action is in the Completed state, you manually end the lifecycle hook that triggers the lifecycle action ahead of schedule.
//
// If you do not specify the action to perform, such as execute a specific OOS template, after a lifecycle hook ends, you can call this operation to obtain the token of the lifecycle action that corresponds to the lifecycle hook. Then, you can specify a custom action to perform after the lifecycle hook ends.
//
// @param request - DescribeLifecycleActionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLifecycleActionsResponse
func (client *Client) DescribeLifecycleActionsWithContext(ctx context.Context, request *DescribeLifecycleActionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLifecycleActionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LifecycleActionStatus) {
		query["LifecycleActionStatus"] = request.LifecycleActionStatus
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingActivityId) {
		query["ScalingActivityId"] = request.ScalingActivityId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLifecycleActions"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLifecycleActionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of lifecycle hooks by specifying lifecycle hook IDs or scaling group IDs. The information includes the next action to perform after the wait state ends, the scaling activity type associated with the lifecycle hook, the notification recipient identifier, and the suspension wait time.
//
// Description:
//
// You can query lifecycle hooks in the following three ways:
//
// - Specify a list of lifecycle hook IDs (LifecycleHookIds). In this case, the ScalingGroupId and LifecycleHookName parameters are ignored.
//
// - Specify a scaling group ID (ScalingGroupId).
//
// - Specify both a scaling group ID (ScalingGroupId) and a lifecycle hook name (LifecycleHookName).
//
// @param request - DescribeLifecycleHooksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLifecycleHooksResponse
func (client *Client) DescribeLifecycleHooksWithContext(ctx context.Context, request *DescribeLifecycleHooksRequest, runtime *dara.RuntimeOptions) (_result *DescribeLifecycleHooksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LifecycleHookIds) {
		query["LifecycleHookIds"] = request.LifecycleHookIds
	}

	if !dara.IsNil(request.LifecycleHookName) {
		query["LifecycleHookName"] = request.LifecycleHookName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLifecycleHooks"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLifecycleHooksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource quotas. For example, you can call the DescribeLimitation operation to query the maximum numbers of scheduled tasks that you can create in a scaling group, load balancers that you can attach to a scaling group, instances that you can add to a scaling group, and scaling configurations that you can create in a scaling group.
//
// @param request - DescribeLimitationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLimitationResponse
func (client *Client) DescribeLimitationWithContext(ctx context.Context, request *DescribeLimitationRequest, runtime *dara.RuntimeOptions) (_result *DescribeLimitationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLimitation"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLimitationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries notification settings. You can call the DescribeNotificationConfiguration operation to query notification settings of scaling events or resource changes, including the notification types and methods.
//
// @param request - DescribeNotificationConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNotificationConfigurationsResponse
func (client *Client) DescribeNotificationConfigurationsWithContext(ctx context.Context, request *DescribeNotificationConfigurationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeNotificationConfigurationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNotificationConfigurations"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNotificationConfigurationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries notification types. You can call the DescribeNotificationTypes operation to query the types of notifications on scaling events or resource changes occurred in your scaling groups. Notifications are triggered in scenarios such as successful scale-out events, successful scale-in events, expiration of scheduled tasks, and partially successful scale-out events.
//
// @param request - DescribeNotificationTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNotificationTypesResponse
func (client *Client) DescribeNotificationTypesWithContext(ctx context.Context, request *DescribeNotificationTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeNotificationTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNotificationTypes"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNotificationTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instance types that meet specified requirements by filtering based on parameters such as the vSwitch ID, number of vCPUs, memory size, instance family level, and maximum acceptable price.
//
// @param request - DescribePatternTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePatternTypesResponse
func (client *Client) DescribePatternTypesWithContext(ctx context.Context, request *DescribePatternTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribePatternTypesResponse, _err error) {
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
		Action:      dara.String("DescribePatternTypes"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePatternTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries regions. Before you activate Auto Scaling, you can call the DescribeRegions operation to query the regions where Auto Scaling is officially launched. This preliminary step facilitates the strategic selection of both the optimal region and availability zones for activating Auto Scaling, thereby guaranteeing the finest access speeds and operational efficiency within your chosen geographical area.
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2022-02-22"),
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
// Queries the details of scaling activities to check the results of scaling activities triggered by scheduled tasks, event-triggered tasks, or manual execution of scaling rules. For example, you can view the status, cause, and total number of instances after the scaling activity on the scaling activity details page.
//
// Description:
//
// - You can specify a scaling group ID to query all scaling activity records in the scaling group.
//
// - You can filter query results by the status of the scaling activity.
//
// - You can query the scaling activity that occurred within the last 30 days.
//
// @param request - DescribeScalingActivitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingActivitiesResponse
func (client *Client) DescribeScalingActivitiesWithContext(ctx context.Context, request *DescribeScalingActivitiesRequest, runtime *dara.RuntimeOptions) (_result *DescribeScalingActivitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceRefreshTaskId) {
		query["InstanceRefreshTaskId"] = request.InstanceRefreshTaskId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScalingActivityIds) {
		query["ScalingActivityIds"] = request.ScalingActivityIds
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.StatusCode) {
		query["StatusCode"] = request.StatusCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScalingActivities"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScalingActivitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a scaling activity. You can query a scaling activity by its ID. The scaling activity details include the scaling activity status, error code, and error message. You can efficiently troubleshoot issues and analyze service performance based on the error message.
//
// @param request - DescribeScalingActivityDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingActivityDetailResponse
func (client *Client) DescribeScalingActivityDetailWithContext(ctx context.Context, request *DescribeScalingActivityDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeScalingActivityDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScalingActivityId) {
		query["ScalingActivityId"] = request.ScalingActivityId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScalingActivityDetail"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScalingActivityDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The distribution of scaling activities in different states over a period of time.
//
// @param request - DescribeScalingActivityStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingActivityStatisticsResponse
func (client *Client) DescribeScalingActivityStatisticsWithContext(ctx context.Context, request *DescribeScalingActivityStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeScalingActivityStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScalingActivityStatistics"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScalingActivityStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about ECS-type scaling configurations, including instance types and images, to help you understand the details of ECS instance templates used during scale-out.
//
// @param request - DescribeScalingConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingConfigurationsResponse
func (client *Client) DescribeScalingConfigurationsWithContext(ctx context.Context, request *DescribeScalingConfigurationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeScalingConfigurationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScalingConfigurationIds) {
		query["ScalingConfigurationIds"] = request.ScalingConfigurationIds
	}

	if !dara.IsNil(request.ScalingConfigurationNames) {
		query["ScalingConfigurationNames"] = request.ScalingConfigurationNames
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScalingConfigurations"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScalingConfigurationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified scaling group, such as the basic information, instances in the scaling group, and scaling configurations. If the OutputFormat parameter is set to YAML for an Elastic Container Instance (ECI) scaling group, the content in the Kubernetes Deployment YAML format is returned.
//
// @param request - DescribeScalingGroupDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingGroupDetailResponse
func (client *Client) DescribeScalingGroupDetailWithContext(ctx context.Context, request *DescribeScalingGroupDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeScalingGroupDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OutputFormat) {
		query["OutputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScalingGroupDetail"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScalingGroupDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeScalingGroupDiagnoseDetails
//
// @param request - DescribeScalingGroupDiagnoseDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingGroupDiagnoseDetailsResponse
func (client *Client) DescribeScalingGroupDiagnoseDetailsWithContext(ctx context.Context, request *DescribeScalingGroupDiagnoseDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeScalingGroupDiagnoseDetailsResponse, _err error) {
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
		Action:      dara.String("DescribeScalingGroupDiagnoseDetails"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScalingGroupDiagnoseDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of scaling groups, including basic information, instances in the scaling groups, and scaling configurations.
//
// @param request - DescribeScalingGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingGroupsResponse
func (client *Client) DescribeScalingGroupsWithContext(ctx context.Context, request *DescribeScalingGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeScalingGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScalingGroupIds) {
		query["ScalingGroupIds"] = request.ScalingGroupIds
	}

	if !dara.IsNil(request.ScalingGroupName) {
		query["ScalingGroupName"] = request.ScalingGroupName
	}

	if !dara.IsNil(request.ScalingGroupNames) {
		query["ScalingGroupNames"] = request.ScalingGroupNames
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScalingGroups"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScalingGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of ECS instances and their details in a scaling group by scaling group ID. The details include the total number of running spot instances, the total number of ECS instances, the warmup status, and the lifecycle status of the ECS instances in the scaling group. This operation also allows you to filter ECS instances by health status, lifecycle status, and creation type.
//
// @param request - DescribeScalingInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingInstancesResponse
func (client *Client) DescribeScalingInstancesWithContext(ctx context.Context, request *DescribeScalingInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeScalingInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreationType) {
		query["CreationType"] = request.CreationType
	}

	if !dara.IsNil(request.CreationTypes) {
		query["CreationTypes"] = request.CreationTypes
	}

	if !dara.IsNil(request.HealthStatus) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.LifecycleState) {
		query["LifecycleState"] = request.LifecycleState
	}

	if !dara.IsNil(request.LifecycleStates) {
		query["LifecycleStates"] = request.LifecycleStates
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScalingActivityId) {
		query["ScalingActivityId"] = request.ScalingActivityId
	}

	if !dara.IsNil(request.ScalingConfigurationId) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScalingInstances"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScalingInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries scaling rules and their details by specifying parameters such as the scaling rule ID, scaling rule name, and unique identifier of the scaling rule. The details include the adjustment type, instance warmup time, and event-triggered tasks associated with the scaling rules.
//
// @param request - DescribeScalingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingRulesResponse
func (client *Client) DescribeScalingRulesWithContext(ctx context.Context, request *DescribeScalingRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeScalingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.ScalingRuleAris) {
		query["ScalingRuleAris"] = request.ScalingRuleAris
	}

	if !dara.IsNil(request.ScalingRuleIds) {
		query["ScalingRuleIds"] = request.ScalingRuleIds
	}

	if !dara.IsNil(request.ScalingRuleNames) {
		query["ScalingRuleNames"] = request.ScalingRuleNames
	}

	if !dara.IsNil(request.ScalingRuleType) {
		query["ScalingRuleType"] = request.ScalingRuleType
	}

	if !dara.IsNil(request.ShowAlarmRules) {
		query["ShowAlarmRules"] = request.ShowAlarmRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScalingRules"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScalingRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of scheduled tasks. A scheduled task is a preset task that executes a specified scaling rule at a specified time to implement automatic scale-out and scale-in of compute resources. After you create a scheduled task, you can call the DescribeScheduledTasks API to query the details of the scheduled task, such as the trigger time, the total number of scheduled tasks, and the ID of the scaling group that executes the scheduled task.
//
// Description:
//
// You can query scheduled tasks by scheduled task ID, scheduled task name, or the scaling rule executed by the scheduled task.
//
// @param request - DescribeScheduledTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScheduledTasksResponse
func (client *Client) DescribeScheduledTasksWithContext(ctx context.Context, request *DescribeScheduledTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeScheduledTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RecurrenceType) {
		query["RecurrenceType"] = request.RecurrenceType
	}

	if !dara.IsNil(request.RecurrenceValue) {
		query["RecurrenceValue"] = request.RecurrenceValue
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.ScheduledActions) {
		query["ScheduledActions"] = request.ScheduledActions
	}

	if !dara.IsNil(request.ScheduledTaskIds) {
		query["ScheduledTaskIds"] = request.ScheduledTaskIds
	}

	if !dara.IsNil(request.ScheduledTaskNames) {
		query["ScheduledTaskNames"] = request.ScheduledTaskNames
	}

	if !dara.IsNil(request.TaskEnabled) {
		query["TaskEnabled"] = request.TaskEnabled
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScheduledTasks"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScheduledTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of ECS instances in the warm pool of a scaling group.
//
// @param request - DescribeWarmPoolInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWarmPoolInstancesResponse
func (client *Client) DescribeWarmPoolInstancesWithContext(ctx context.Context, request *DescribeWarmPoolInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeWarmPoolInstancesResponse, _err error) {
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWarmPoolInstances"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWarmPoolInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches Application Load Balancer (ALB) server groups from a scaling group. When you call the DetachAlbServerGroups operation, you can specify the scaling group ID, ALB server group ID, and port used by Elastic Compute Service (ECS) instances in the desired ALB server group to detach one or more ALB server groups from your scaling group.
//
// @param request - DetachAlbServerGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachAlbServerGroupsResponse
func (client *Client) DetachAlbServerGroupsWithContext(ctx context.Context, request *DetachAlbServerGroupsRequest, runtime *dara.RuntimeOptions) (_result *DetachAlbServerGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlbServerGroups) {
		query["AlbServerGroups"] = request.AlbServerGroups
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ForceDetach) {
		query["ForceDetach"] = request.ForceDetach
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachAlbServerGroups"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachAlbServerGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches one or more ApsaraDB RDS instances from a scaling group to remove the mandatory association between the scaling group and the RDS instances and improve flexibility.
//
// @param request - DetachDBInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachDBInstancesResponse
func (client *Client) DetachDBInstancesWithContext(ctx context.Context, request *DetachDBInstancesRequest, runtime *dara.RuntimeOptions) (_result *DetachDBInstancesResponse, _err error) {
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

	if !dara.IsNil(request.DBInstances) {
		query["DBInstances"] = request.DBInstances
	}

	if !dara.IsNil(request.ForceDetach) {
		query["ForceDetach"] = request.ForceDetach
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoveSecurityGroup) {
		query["RemoveSecurityGroup"] = request.RemoveSecurityGroup
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachDBInstances"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachDBInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes instances from a scaling group. If an enabled scaling group has no ongoing scaling activities, you can call the DetachInstances operation to remove one or more Elastic Compute Service (ECS) instances, elastic container instances, or Alibaba Cloud-hosted third-party instances from the scaling group.
//
// Description:
//
//	  Before you call the DetachInstances operation, make sure that the following conditions are met:
//
//	    	- The specified scaling group is enabled.
//
//	    	- The specified scaling group does not have any ongoing scaling activities.
//
//	    **
//
//	    **Note*	- If the specified scaling group does not have any ongoing scaling activities, the operation can bypass the cooldown period of the scaling group and immediately trigger scaling activities.
//
//		- Before you call this operation, take note of the following items:
//
//	    	- The ECS instances, elastic container instances, or Alibaba Cloud-hosted third-party instances that are removed from a scaling group by using this operation can run independently. If needed, you can call the [AttachInstances](https://help.aliyun.com/document_detail/25954.html) operation to re-add these instances to a scaling group.
//
//	    	- The ECS instances, elastic container instances, or Alibaba Cloud-hosted third-party instances that are removed from a scaling group by using this operation are not stopped or released.
//
//	    	- A successful call only means that Auto Scaling accepts your request. Scaling activities can be triggered as expected, but their successful execution is not guaranteed. You can query the status of a scaling activity based on the ScalingActivityId response parameter.
//
//	    	- The removal of ECS instances, elastic container instances, or Alibaba Cloud-hosted third-party instances must not result in the overall number of instances within the specified scaling group falling below the minimum capacity threshold (MinSize); otherwise, an error will be reported.
//
// @param request - DetachInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachInstancesResponse
func (client *Client) DetachInstancesWithContext(ctx context.Context, request *DetachInstancesRequest, runtime *dara.RuntimeOptions) (_result *DetachInstancesResponse, _err error) {
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

	if !dara.IsNil(request.DecreaseDesiredCapacity) {
		query["DecreaseDesiredCapacity"] = request.DecreaseDesiredCapacity
	}

	if !dara.IsNil(request.DetachOption) {
		query["DetachOption"] = request.DetachOption
	}

	if !dara.IsNil(request.IgnoreInvalidInstance) {
		query["IgnoreInvalidInstance"] = request.IgnoreInvalidInstance
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.LifecycleHook) {
		query["LifecycleHook"] = request.LifecycleHook
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachInstances"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches Classic Load Balancer (CLB, formerly known as Server Load Balancer or SLB) instances from a scaling group. If the current CLB instance no longer meets your business requirements, you can call the DetachLoadBalancers operation to detach it from your scaling group. When you call this operation, you can use ScalingGroupId, LoadBalancer.N, and ForceDetach to specify one or more CLB instances to detach. You can also determine whether to call this operation asynchronously and whether to remove the Elastic Compute Service (ECS) instances acting as backend servers from the backend server groups of the CLB instance. You can call this operation to detach only CLB instances from a scaling group.
//
// @param request - DetachLoadBalancersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachLoadBalancersResponse
func (client *Client) DetachLoadBalancersWithContext(ctx context.Context, request *DetachLoadBalancersRequest, runtime *dara.RuntimeOptions) (_result *DetachLoadBalancersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Async) {
		query["Async"] = request.Async
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ForceDetach) {
		query["ForceDetach"] = request.ForceDetach
	}

	if !dara.IsNil(request.LoadBalancers) {
		query["LoadBalancers"] = request.LoadBalancers
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachLoadBalancers"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachLoadBalancersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Remove one or more SLB server groups from a scaling group. Supported server group types include Application Load Balancer (ALB), Network Load Balancer (NLB), and Gateway Load Balancer (GWLB), which can dynamically adjust SLB policies and help improve high availability of the system.
//
// @param request - DetachServerGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachServerGroupsResponse
func (client *Client) DetachServerGroupsWithContext(ctx context.Context, request *DetachServerGroupsRequest, runtime *dara.RuntimeOptions) (_result *DetachServerGroupsResponse, _err error) {
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

	if !dara.IsNil(request.ForceDetach) {
		query["ForceDetach"] = request.ForceDetach
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.ServerGroups) {
		query["ServerGroups"] = request.ServerGroups
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachServerGroups"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachServerGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches vServer groups from a scaling group. Auto Scaling supports the attachment of load balancers to scaling groups to improve service performance. If the load balancer currently attached to your scaling group is no longer needed to distribute the access traffic to the instances in your scaling group, you can call the DetachVServerGroups operation to detach one or more vServer groups of this load balancer from the scaling group.
//
// Description:
//
//	  When you call the DetachVServerGroups operation, you must use the following parameters to specify the vServer groups that you want to detach from your scaling group:
//
//	    	- LoadBalancerId: the ID of the load balancer
//
//	    	- VServerGroupId: the ID of the vServer group
//
//	    	- Port: the port number of the vServer group
//
//		- When the vServer group specified by the request parameters matches that attached to your scaling group, this operation yields a favorable result. Otherwise, the request is ignored and no error is reported.
//
//		- Before you call this operation, you must make sure that the load balancer has ceased routing the access traffic to the instances in the scaling group. Failure to do so may lead to service requests being dropped or lost during the detachment process.
//
// @param request - DetachVServerGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachVServerGroupsResponse
func (client *Client) DetachVServerGroupsWithContext(ctx context.Context, request *DetachVServerGroupsRequest, runtime *dara.RuntimeOptions) (_result *DetachVServerGroupsResponse, _err error) {
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

	if !dara.IsNil(request.ForceDetach) {
		query["ForceDetach"] = request.ForceDetach
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.VServerGroups) {
		query["VServerGroups"] = request.VServerGroups
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachVServerGroups"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachVServerGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an event-triggered task. If your business pattern is unpredictable or prone to unforeseen traffic spikes, you can create event-triggered tasks by associating CloudMonitor metrics to effectively monitor fluctuations in your business workload. Upon detecting that the criteria for alerts, as specified in event-triggered tasks, are fulfilled, Auto Scaling promptly issues alerts and executes the scaling rules predefined within those tasks. This process occurs within the predefined effective time windows of the tasks, thereby facilitating the automatic increase or decrease of Elastic Compute Service (ECS) instances or elastic container instances within your scaling groups. Ultimately, this mechanism ensures the dynamic optimization of resources based on real-time workload demands. If you currently do not need an event-triggered task, you can call the DisableAlarm operation to disable it.
//
// Description:
//
// Before you disable an event-triggered task, make sure that the task is in the `Normal`, `Alert`, or `Insufficient Data` state.
//
// @param request - DisableAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAlarmResponse
func (client *Client) DisableAlarmWithContext(ctx context.Context, request *DisableAlarmRequest, runtime *dara.RuntimeOptions) (_result *DisableAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmTaskId) {
		query["AlarmTaskId"] = request.AlarmTaskId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableAlarm"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAlarmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a scaling group. If you temporarily do not require a scaling group that is in the Enabled state, you can call the DisableScalingGroup operation to disable it.
//
// Description:
//
// Before you call this operation to disable a scaling group, take note of the following items:
//
//   - If scaling activities are being executed in the specified scaling group when you call this operation, these activities will continue until they are complete. However, scaling activities that are triggered after this operation is called will be rejected.
//
//   - This operation can be called only when the scaling group is in the Active state.
//
// @param request - DisableScalingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableScalingGroupResponse
func (client *Client) DisableScalingGroupWithContext(ctx context.Context, request *DisableScalingGroupRequest, runtime *dara.RuntimeOptions) (_result *DisableScalingGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableScalingGroup"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableScalingGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an event-triggered task. If your business pattern is unpredictable or prone to unforeseen traffic spikes, you can create event-triggered tasks by associating CloudMonitor metrics to effectively monitor fluctuations in your business workload. Upon detecting that the criteria for alerts, as specified in event-triggered tasks, are fulfilled, Auto Scaling promptly issues alerts and executes the scaling rules predefined within those tasks. This process occurs within the predefined effective time windows of the tasks, thereby facilitating the automatic increase or decrease of Elastic Compute Service (ECS) instances or elastic container instances within your scaling groups. Ultimately, this mechanism ensures the dynamic optimization of resources based on real-time workload demands. If you want to reuse an event-triggered task that is in the Disabled state, you can call the EnableAlarm operation to enable it.
//
// @param request - EnableAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAlarmResponse
func (client *Client) EnableAlarmWithContext(ctx context.Context, request *EnableAlarmRequest, runtime *dara.RuntimeOptions) (_result *EnableAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmTaskId) {
		query["AlarmTaskId"] = request.AlarmTaskId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableAlarm"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableAlarmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a scaling group. If a scaling group is in the Disabled state and contains an instance configuration source such as a launch template or a scaling configuration, you can call the EnableScalingGroup operation to enable the scaling group. This operation permits Auto Scaling to dynamically adjust the computing power (also known as the number of instances) in the scaling group based on your business requirements.
//
// Description:
//
//	  You can call this operation to enable a scaling group only if the scaling group is in the Inactive state and contains an instance configuration source such as a launch temple or a scaling configuration. The instance configuration source can also be the Elastic Compute Service (ECS) instance that you specified when you created the scaling group. If the preceding requirements are not met, the operation will fail.
//
//	    **
//
//	    **Note*	- A scaling group can have only one active instance configuration source at a time. When you call this operation to enable a scaling group, you can specify a scaling configuration or a launch template for the scaling group. If the scaling group already have an instance configuration source defined prior to your calling, the scaling configuration or launch template specified within your request will supersede the existing scaling configuration or launch template.
//
//		- If you specify InstanceId.N to add to the scaling group within your request, Auto Scaling will check whether the addition of InstanceId.N will cause the total number of ECS instances in the scaling group to fall outside the boundaries specified by MinSize and MaxSize after you call this operation.
//
//	    	- If the call results in the total number of ECS instances dropping below the value of MinSize, Auto Scaling proactively creates pay-as-you-go ECS instances to ensure that the total number reaches the minimum threshold. For example, if you set MinSize to 5 when you created a scaling group and include InstanceId.N within your request to add two ECS instances when you attempt to enable the scaling group, Auto Scaling creates three more ECS instances in the scaling group after the two ECS instances are added.
//
//	    	- If the call results in the total number of ECS instances exceeding the value of MaxSize, the operation fails.
//
// @param request - EnableScalingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableScalingGroupResponse
func (client *Client) EnableScalingGroupWithContext(ctx context.Context, request *EnableScalingGroupRequest, runtime *dara.RuntimeOptions) (_result *EnableScalingGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveScalingConfigurationId) {
		query["ActiveScalingConfigurationId"] = request.ActiveScalingConfigurationId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.LaunchTemplateId) {
		query["LaunchTemplateId"] = request.LaunchTemplateId
	}

	if !dara.IsNil(request.LaunchTemplateOverrides) {
		query["LaunchTemplateOverrides"] = request.LaunchTemplateOverrides
	}

	if !dara.IsNil(request.LaunchTemplateVersion) {
		query["LaunchTemplateVersion"] = request.LaunchTemplateVersion
	}

	if !dara.IsNil(request.LoadBalancerWeights) {
		query["LoadBalancerWeights"] = request.LoadBalancerWeights
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableScalingGroup"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableScalingGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Puts an Elastic Compute Service (ECS) instance into the Standby state.
//
// Description:
//
// ## Description
//
//   - If you call the operation to put an ECS instance in a scaling group that is associated with a Classic Load Balancer (CLB) instance into the Standby state, the weight of the ECS instance as a backend server of the CLB instance is set to 0.
//
//   - You can remove an instance that is in the Standby state from a scaling group, and then release the instance.
//
//   - ECS instances that are in the Standby state are not removed from the scaling group during scale-in activities triggered by event-triggered tasks.
//
//   - If Auto Scaling considers an ECS instance that is in the Standby state unhealthy, for example, the ECS instance is being stopped or being restarted, Auto Scaling does not update the health status of the ECS instance or trigger scale-in activities to remove the ECS instance from the scaling group. Auto Scaling updates the health status of the ECS instance only when the ECS instance is no longer in the Standby state.
//
// @param request - EnterStandbyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterStandbyResponse
func (client *Client) EnterStandbyWithContext(ctx context.Context, request *EnterStandbyRequest, runtime *dara.RuntimeOptions) (_result *EnterStandbyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Async) {
		query["Async"] = request.Async
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterStandby"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterStandbyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a scaling rule. To adjust the number of Elastic Compute Service (ECS) instances or elastic container instances, you can manually execute a scaling rule or enable Auto Scaling to execute a scaling rule. You can call the ExecuteScalingRule operation to execute simple scaling rules or step scaling rules. Auto Scaling automatically executes target tracking scaling rules and predictive scaling rules on your behalf without requiring explicit execution calls.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - The scaling group is in the Active state.
//
//   - No scaling activities in the scaling group are in progress.
//
// If no scaling activities in the scaling group are in progress, the operation can trigger scaling activities even before the cooldown time expires.
//
// A successful call indicates that Auto Scaling accepts the request. However, the scaling activity may still fail. You can obtain the status of a scaling activity by using the value of the ScalingActivityId parameter in the response.
//
// If the addition of a specified number of Elastic Compute Service (ECS) instances to a scaling group causes the total number of ECS instances in the scaling group to exceed the maximum number of instances allowed, Auto Scaling adds only a specific number of ECS instances to ensure that the total number of instances is equal to the maximum number of instances.
//
// If the removal of a specified number of ECS instances from a scaling group causes the total number of ECS instances in the scaling group to drop below the minimum number of instances allowed, Auto Scaling removes only a specific number of ECS instances to ensure that the total number of instances is equal to the minimum number of instances.
//
// You can specify only a limited number of ECS instances in each adjustment. For more information, see the description of the AdjustmentValue parameter in the CreateScalingRule topic.
//
// @param request - ExecuteScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteScalingRuleResponse
func (client *Client) ExecuteScalingRuleWithContext(ctx context.Context, request *ExecuteScalingRuleRequest, runtime *dara.RuntimeOptions) (_result *ExecuteScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BreachThreshold) {
		query["BreachThreshold"] = request.BreachThreshold
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.MetricValue) {
		query["MetricValue"] = request.MetricValue
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ScalingRuleAri) {
		query["ScalingRuleAri"] = request.ScalingRuleAri
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteScalingRule"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an instance from the Standby state. When a scale-in event is triggered in a scaling group, Auto Scaling does not remove Elastic Compute Service (ECS) instances or elastic container instances that are in the Standby state from the scaling group. If you want to restart the instances that are in the Standby state, you can call the ExitStandby operation to remove the instances from the Standby state and put them into the In Service state.
//
// Description:
//
// After ECS instances or elastic container instances are removed from the Standby state, the following rules apply:
//
//   - The ECS instances or elastic container instances enter the In Service state.
//
//   - The default weight of each ECS instance or elastic container instance as a backend server of the attached load balancer is 50.
//
//   - If you stop or restart the ECS instances or elastic container instances, the health check status of the instances will be updated.
//
//   - When a scale-in event is triggered, Auto Scaling may remove the ECS instances or elastic container instances from the scaling group.
//
// @param request - ExitStandbyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExitStandbyResponse
func (client *Client) ExitStandbyWithContext(ctx context.Context, request *ExitStandbyRequest, runtime *dara.RuntimeOptions) (_result *ExitStandbyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Async) {
		query["Async"] = request.Async
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExitStandby"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExitStandbyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag keys added to Auto Scaling resources. Querying tag keys facilitates easier classification, identification, and monitoring of your Auto Scaling resources, thereby enhancing the flexibility and convenience of your resource management processes.
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
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagKeys"),
		Version:     dara.String("2022-02-22"),
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
// Queries tags. You can call the ListTagResources operation to query tags that are added to Auto Scaling resources, thereby clarifying resource utilization and facilitating efficient resource management. This operation aids in the automation of resource categorization and permission management processes.
//
// Description:
//
//	  Specify at least one of the following request parameters: `ResourceIds` and `Tags`. `Tags.Key` and `Tags.Value` are used to specify the query objects.
//
//		- If you provide both `ResourceIds` and `Tags` in your request, the response will exclusively include Auto Scaling resources that satisfy the criteria set by these parameters, ensuring targeted and precise information retrieval.
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
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2022-02-22"),
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
// Queries the tag keys associated with Auto Scaling resources to facilitate a deeper comprehension of those resources. By doing so, you can categorize and manage your Auto Scaling resources more efficiently.
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
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagValues"),
		Version:     dara.String("2022-02-22"),
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
// Modifies an event-triggered task.
//
// Description:
//
// - If you create an event-triggered task of the custom monitoring type (MetricType=custom), you must first report custom monitoring metrics to CloudMonitor. For more information, see [Event-triggered tasks for custom monitoring](https://help.aliyun.com/document_detail/74861.html).
//
// - When you create an event-triggered task, you must specify a metric (MetricName) and configure dimension information (Dimension.N.DimensionKey and Dimension.N.DimensionValue) to determine the aggregation scope of the metric data in the scaling group. For example, you can specify user_id and scaling_group to aggregate the metric data of all ECS instances or elastic container instances in the scaling group of the specified user.
//
//   - The metrics of a custom monitoring event-triggered task (MetricType=custom) depend on your custom monitoring metrics.
//
//   - For more information about the metrics supported by system monitoring event-triggered tasks (MetricType=system), see [Event-triggered tasks for system monitoring](https://help.aliyun.com/document_detail/74854.html).
//
// > In the dimension information, user_id and scaling_group are automatically populated by the system. You must manually specify device and state. For more information, see the Dimension.N.DimensionKey and Dimension.N.DimensionValue parameters.
//
// @param request - ModifyAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAlarmResponse
func (client *Client) ModifyAlarmWithContext(ctx context.Context, request *ModifyAlarmRequest, runtime *dara.RuntimeOptions) (_result *ModifyAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmActions) {
		query["AlarmActions"] = request.AlarmActions
	}

	if !dara.IsNil(request.AlarmTaskId) {
		query["AlarmTaskId"] = request.AlarmTaskId
	}

	if !dara.IsNil(request.ComparisonOperator) {
		query["ComparisonOperator"] = request.ComparisonOperator
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.Effective) {
		query["Effective"] = request.Effective
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.Expressions) {
		query["Expressions"] = request.Expressions
	}

	if !dara.IsNil(request.ExpressionsLogicOperator) {
		query["ExpressionsLogicOperator"] = request.ExpressionsLogicOperator
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.Statistics) {
		query["Statistics"] = request.Statistics
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAlarm"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAlarmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the status of scaling activities that prompt text message or email notifications.
//
// @param request - ModifyAlertConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAlertConfigurationResponse
func (client *Client) ModifyAlertConfigurationWithContext(ctx context.Context, request *ModifyAlertConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifyAlertConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScaleStatuses) {
		query["ScaleStatuses"] = request.ScaleStatuses
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAlertConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAlertConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a scaling configuration by specifying parameters such as the ID of the scaling configuration to be modified, the name of the scaling configuration, and the attributes of the Elastic Container Instance (ECI), such as the instance restart policy, the preemption policy, and the bandwidth of the elastic IP address (EIP).
//
// Description:
//
// - If you modify the name of the scaling configuration, ensure that no other scaling configuration in the same scaling group has the same name.
//
// - You can call the [DescribeEciScalingConfigurations](https://help.aliyun.com/document_detail/459374.html) operation to query and verify the modification result.
//
// @param request - ModifyEciScalingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEciScalingConfigurationResponse
func (client *Client) ModifyEciScalingConfigurationWithContext(ctx context.Context, request *ModifyEciScalingConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifyEciScalingConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcrRegistryInfos) {
		query["AcrRegistryInfos"] = request.AcrRegistryInfos
	}

	if !dara.IsNil(request.ActiveDeadlineSeconds) {
		query["ActiveDeadlineSeconds"] = request.ActiveDeadlineSeconds
	}

	if !dara.IsNil(request.AutoCreateEip) {
		query["AutoCreateEip"] = request.AutoCreateEip
	}

	if !dara.IsNil(request.AutoMatchImageCache) {
		query["AutoMatchImageCache"] = request.AutoMatchImageCache
	}

	if !dara.IsNil(request.ContainerGroupName) {
		query["ContainerGroupName"] = request.ContainerGroupName
	}

	if !dara.IsNil(request.Containers) {
		query["Containers"] = request.Containers
	}

	if !dara.IsNil(request.ContainersUpdateType) {
		query["ContainersUpdateType"] = request.ContainersUpdateType
	}

	if !dara.IsNil(request.CostOptimization) {
		query["CostOptimization"] = request.CostOptimization
	}

	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.CpuOptionsCore) {
		query["CpuOptionsCore"] = request.CpuOptionsCore
	}

	if !dara.IsNil(request.CpuOptionsThreadsPerCore) {
		query["CpuOptionsThreadsPerCore"] = request.CpuOptionsThreadsPerCore
	}

	if !dara.IsNil(request.DataCacheBucket) {
		query["DataCacheBucket"] = request.DataCacheBucket
	}

	if !dara.IsNil(request.DataCacheBurstingEnabled) {
		query["DataCacheBurstingEnabled"] = request.DataCacheBurstingEnabled
	}

	if !dara.IsNil(request.DataCachePL) {
		query["DataCachePL"] = request.DataCachePL
	}

	if !dara.IsNil(request.DataCacheProvisionedIops) {
		query["DataCacheProvisionedIops"] = request.DataCacheProvisionedIops
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DnsConfigNameServers) {
		query["DnsConfigNameServers"] = request.DnsConfigNameServers
	}

	if !dara.IsNil(request.DnsConfigOptions) {
		query["DnsConfigOptions"] = request.DnsConfigOptions
	}

	if !dara.IsNil(request.DnsConfigSearchs) {
		query["DnsConfigSearchs"] = request.DnsConfigSearchs
	}

	if !dara.IsNil(request.DnsPolicy) {
		query["DnsPolicy"] = request.DnsPolicy
	}

	if !dara.IsNil(request.EgressBandwidth) {
		query["EgressBandwidth"] = request.EgressBandwidth
	}

	if !dara.IsNil(request.EipBandwidth) {
		query["EipBandwidth"] = request.EipBandwidth
	}

	if !dara.IsNil(request.EnableSls) {
		query["EnableSls"] = request.EnableSls
	}

	if !dara.IsNil(request.EphemeralStorage) {
		query["EphemeralStorage"] = request.EphemeralStorage
	}

	if !dara.IsNil(request.GpuDriverVersion) {
		query["GpuDriverVersion"] = request.GpuDriverVersion
	}

	if !dara.IsNil(request.HostAliases) {
		query["HostAliases"] = request.HostAliases
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.ImageRegistryCredentials) {
		query["ImageRegistryCredentials"] = request.ImageRegistryCredentials
	}

	if !dara.IsNil(request.ImageSnapshotId) {
		query["ImageSnapshotId"] = request.ImageSnapshotId
	}

	if !dara.IsNil(request.IngressBandwidth) {
		query["IngressBandwidth"] = request.IngressBandwidth
	}

	if !dara.IsNil(request.InitContainers) {
		query["InitContainers"] = request.InitContainers
	}

	if !dara.IsNil(request.InstanceFamilyLevel) {
		query["InstanceFamilyLevel"] = request.InstanceFamilyLevel
	}

	if !dara.IsNil(request.InstanceTypes) {
		query["InstanceTypes"] = request.InstanceTypes
	}

	if !dara.IsNil(request.Ipv6AddressCount) {
		query["Ipv6AddressCount"] = request.Ipv6AddressCount
	}

	if !dara.IsNil(request.LoadBalancerWeight) {
		query["LoadBalancerWeight"] = request.LoadBalancerWeight
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.NtpServers) {
		query["NtpServers"] = request.NtpServers
	}

	if !dara.IsNil(request.Override) {
		query["Override"] = request.Override
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RamRoleName) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.RestartPolicy) {
		query["RestartPolicy"] = request.RestartPolicy
	}

	if !dara.IsNil(request.ScalingConfigurationId) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	if !dara.IsNil(request.ScalingConfigurationName) {
		query["ScalingConfigurationName"] = request.ScalingConfigurationName
	}

	if !dara.IsNil(request.SecurityContextSysCtls) {
		query["SecurityContextSysCtls"] = request.SecurityContextSysCtls
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SpotPriceLimit) {
		query["SpotPriceLimit"] = request.SpotPriceLimit
	}

	if !dara.IsNil(request.SpotStrategy) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TerminationGracePeriodSeconds) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	if !dara.IsNil(request.Volumes) {
		query["Volumes"] = request.Volumes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEciScalingConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEciScalingConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an ECS instance in a scaling group. Currently, only the lifecycle management attribute of instances manually added to the scaling group can be modified.
//
// @param request - ModifyInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceAttributeResponse
func (client *Client) ModifyInstanceAttributeWithContext(ctx context.Context, request *ModifyInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Entrusted) {
		query["Entrusted"] = request.Entrusted
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceAttribute"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a lifecycle hook, such as the applicable scaling activity type, the wait time for the scaling activity, and the subsequent action after the wait state ends. You can specify the lifecycle hook by its ID or name and the scaling group ID without the need to delete the lifecycle hook if it no longer meets your business requirements.
//
// Description:
//
// You can specify the lifecycle hook to be modified in one of the following ways:
//
// - Specify only the lifecycle hook ID (LifecycleHookId). In this case, the ScalingGroupId and LifecycleHookName parameters are ignored.
//
// - Specify both the scaling group ID (ScalingGroupId) and the lifecycle hook name (LifecycleHookName). The lifecycle hook name is unique within a scaling group.
//
// @param request - ModifyLifecycleHookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLifecycleHookResponse
func (client *Client) ModifyLifecycleHookWithContext(ctx context.Context, request *ModifyLifecycleHookRequest, runtime *dara.RuntimeOptions) (_result *ModifyLifecycleHookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultResult) {
		query["DefaultResult"] = request.DefaultResult
	}

	if !dara.IsNil(request.HeartbeatTimeout) {
		query["HeartbeatTimeout"] = request.HeartbeatTimeout
	}

	if !dara.IsNil(request.LifecycleActionSources) {
		query["LifecycleActionSources"] = request.LifecycleActionSources
	}

	if !dara.IsNil(request.LifecycleHookId) {
		query["LifecycleHookId"] = request.LifecycleHookId
	}

	if !dara.IsNil(request.LifecycleHookName) {
		query["LifecycleHookName"] = request.LifecycleHookName
	}

	if !dara.IsNil(request.LifecycleHookStatus) {
		query["LifecycleHookStatus"] = request.LifecycleHookStatus
	}

	if !dara.IsNil(request.LifecycleTransition) {
		query["LifecycleTransition"] = request.LifecycleTransition
	}

	if !dara.IsNil(request.NotificationArn) {
		query["NotificationArn"] = request.NotificationArn
	}

	if !dara.IsNil(request.NotificationMetadata) {
		query["NotificationMetadata"] = request.NotificationMetadata
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLifecycleHook"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLifecycleHookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies event notification rules. Event notification rules support automatic sending of notifications to CloudMonitor, Message Service (MNS) topics, or MNS queues when a specified type of events occur. This helps you learn about the dynamics of your scaling group at the earliest opportunity and further automate resource management. If an existing event notification rule does not meet your business requirements, you can call the ModifyNotificationConfiguration operation to modify the event notification rule, without the need to create a new rule. Take not that you cannot modify the notification method of an event notification rule by calling this operation.
//
// @param request - ModifyNotificationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNotificationConfigurationResponse
func (client *Client) ModifyNotificationConfigurationWithContext(ctx context.Context, request *ModifyNotificationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifyNotificationConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MessageEncoding) {
		query["MessageEncoding"] = request.MessageEncoding
	}

	if !dara.IsNil(request.NotificationArn) {
		query["NotificationArn"] = request.NotificationArn
	}

	if !dara.IsNil(request.NotificationTypes) {
		query["NotificationTypes"] = request.NotificationTypes
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.TimeZone) {
		query["TimeZone"] = request.TimeZone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNotificationConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNotificationConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an ECS-type scaling configuration (the ECS instance template for scale-out) if the current scaling configuration does not meet your business requirements.
//
// Description:
//
// - If you modify the name of a scaling configuration, ensure that no other scaling configuration in the same scaling group has the same name.
//
// - You can bind a primary network interface card (NIC) by using one of the following methods. However, you can use only one method in a single request. If you use both methods, the request fails and an error message is returned. When you modify the scaling configuration by using one method, the NIC configuration set by the other method is cleared.
//
//   - Directly configure the primary NIC by using parameters such as SecurityGroupId, SecurityGroupIds, and Ipv6AddressCount.
//
//   - Configure the primary NIC and secondary NICs by using the NetworkInterfaces parameter. When you use the NetworkInterfaces parameter, you must configure the primary NIC. If NetworkInterface.InstanceType is set to Primary, the primary NIC is configured. If NetworkInterface.InstanceType is set to Secondary or left empty, a secondary NIC is configured.
//
// @param tmpReq - ModifyScalingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScalingConfigurationResponse
func (client *Client) ModifyScalingConfigurationWithContext(ctx context.Context, tmpReq *ModifyScalingConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifyScalingConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyScalingConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SchedulerOptions) {
		request.SchedulerOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SchedulerOptions, dara.String("SchedulerOptions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Affinity) {
		query["Affinity"] = request.Affinity
	}

	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.CreditSpecification) {
		query["CreditSpecification"] = request.CreditSpecification
	}

	if !dara.IsNil(request.CustomPriorities) {
		query["CustomPriorities"] = request.CustomPriorities
	}

	if !dara.IsNil(request.DataDisks) {
		query["DataDisks"] = request.DataDisks
	}

	if !dara.IsNil(request.DedicatedHostClusterId) {
		query["DedicatedHostClusterId"] = request.DedicatedHostClusterId
	}

	if !dara.IsNil(request.DedicatedHostId) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.DeploymentSetId) {
		query["DeploymentSetId"] = request.DeploymentSetId
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.HpcClusterId) {
		query["HpcClusterId"] = request.HpcClusterId
	}

	if !dara.IsNil(request.HttpEndpoint) {
		query["HttpEndpoint"] = request.HttpEndpoint
	}

	if !dara.IsNil(request.HttpTokens) {
		query["HttpTokens"] = request.HttpTokens
	}

	if !dara.IsNil(request.ImageFamily) {
		query["ImageFamily"] = request.ImageFamily
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.InstanceDescription) {
		query["InstanceDescription"] = request.InstanceDescription
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstancePatternInfos) {
		query["InstancePatternInfos"] = request.InstancePatternInfos
	}

	if !dara.IsNil(request.InstanceTypeCandidateOptions) {
		query["InstanceTypeCandidateOptions"] = request.InstanceTypeCandidateOptions
	}

	if !dara.IsNil(request.InstanceTypeOverrides) {
		query["InstanceTypeOverrides"] = request.InstanceTypeOverrides
	}

	if !dara.IsNil(request.InstanceTypes) {
		query["InstanceTypes"] = request.InstanceTypes
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.InternetMaxBandwidthIn) {
		query["InternetMaxBandwidthIn"] = request.InternetMaxBandwidthIn
	}

	if !dara.IsNil(request.InternetMaxBandwidthOut) {
		query["InternetMaxBandwidthOut"] = request.InternetMaxBandwidthOut
	}

	if !dara.IsNil(request.IoOptimized) {
		query["IoOptimized"] = request.IoOptimized
	}

	if !dara.IsNil(request.Ipv6AddressCount) {
		query["Ipv6AddressCount"] = request.Ipv6AddressCount
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.LoadBalancerWeight) {
		query["LoadBalancerWeight"] = request.LoadBalancerWeight
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.NetworkInterfaces) {
		query["NetworkInterfaces"] = request.NetworkInterfaces
	}

	if !dara.IsNil(request.Override) {
		query["Override"] = request.Override
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PasswordInherit) {
		query["PasswordInherit"] = request.PasswordInherit
	}

	if !dara.IsNil(request.RamRoleName) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourcePoolOptions) {
		query["ResourcePoolOptions"] = request.ResourcePoolOptions
	}

	if !dara.IsNil(request.ScalingConfigurationId) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	if !dara.IsNil(request.ScalingConfigurationName) {
		query["ScalingConfigurationName"] = request.ScalingConfigurationName
	}

	if !dara.IsNil(request.SchedulerOptionsShrink) {
		query["SchedulerOptions"] = request.SchedulerOptionsShrink
	}

	if !dara.IsNil(request.SecondaryEniVSwitchConstraint) {
		query["SecondaryEniVSwitchConstraint"] = request.SecondaryEniVSwitchConstraint
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityGroupIds) {
		query["SecurityGroupIds"] = request.SecurityGroupIds
	}

	if !dara.IsNil(request.SecurityOptions) {
		query["SecurityOptions"] = request.SecurityOptions
	}

	if !dara.IsNil(request.SpotDuration) {
		query["SpotDuration"] = request.SpotDuration
	}

	if !dara.IsNil(request.SpotInterruptionBehavior) {
		query["SpotInterruptionBehavior"] = request.SpotInterruptionBehavior
	}

	if !dara.IsNil(request.SpotPriceLimits) {
		query["SpotPriceLimits"] = request.SpotPriceLimits
	}

	if !dara.IsNil(request.SpotStrategy) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !dara.IsNil(request.StorageSetId) {
		query["StorageSetId"] = request.StorageSetId
	}

	if !dara.IsNil(request.StorageSetPartitionNumber) {
		query["StorageSetPartitionNumber"] = request.StorageSetPartitionNumber
	}

	if !dara.IsNil(request.SystemDiskCategories) {
		query["SystemDiskCategories"] = request.SystemDiskCategories
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Tenancy) {
		query["Tenancy"] = request.Tenancy
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	if !dara.IsNil(request.ImageOptions) {
		query["ImageOptions"] = request.ImageOptions
	}

	if !dara.IsNil(request.PrivatePoolOptions) {
		query["PrivatePoolOptions"] = request.PrivatePoolOptions
	}

	if !dara.IsNil(request.SystemDisk) {
		query["SystemDisk"] = request.SystemDisk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyScalingConfiguration"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScalingConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a scaling group. You can configure a scaling group to automatically adjust your computing capacity by increasing or decreasing the number of instances based on business demands and policies. If the current attributes of the scaling group no longer meet your business requirements, you can call this operation to modify attributes such as the maximum number of instances, minimum number of instances, or expected number of instances. This eliminates the need to recreate the scaling group and reconfigure related attributes, saving time and resource costs.
//
// Description:
//
// - The following parameters cannot be modified:
//
//   - RegionId.
//
//   - LoadBalancerId.
//
//     > To modify load balancing instances, use the AttachLoadBalancers and DetachLoadBalancers operations.
//
//   - DBInstanceId.
//
//     > To modify ApsaraDB RDS instances, use the AttachDBInstances and DetachDBInstances operations.
//
// - You can call this operation only when the scaling group is in the Active or Inactive state.
//
// - Activating a new scaling configuration does not affect the running ECS instances or elastic container instances that were created by using earlier scaling configurations.
//
// - If you modify the MaxSize parameter and the number of ECS instances or elastic container instances in the scaling group exceeds the value of MaxSize, Auto Scaling automatically removes ECS instances or elastic container instances from the scaling group until the number of instances equals the value of MaxSize.
//
// - If you modify the MinSize parameter and the number of ECS instances or elastic container instances in the scaling group is less than the value of MinSize, Auto Scaling automatically adds ECS instances or elastic container instances to the scaling group until the number of instances equals the value of MinSize.
//
// - If the DesiredCapacity parameter is configured for the scaling group and you modify the value of DesiredCapacity, and the current number of ECS instances or elastic container instances in the scaling group is different from the value of DesiredCapacity, Auto Scaling automatically adds or removes ECS instances or elastic container instances until the number of instances equals the value of DesiredCapacity.
//
// @param request - ModifyScalingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScalingGroupResponse
func (client *Client) ModifyScalingGroupWithContext(ctx context.Context, request *ModifyScalingGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyScalingGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveScalingConfigurationId) {
		query["ActiveScalingConfigurationId"] = request.ActiveScalingConfigurationId
	}

	if !dara.IsNil(request.AllocationStrategy) {
		query["AllocationStrategy"] = request.AllocationStrategy
	}

	if !dara.IsNil(request.AutoRebalance) {
		query["AutoRebalance"] = request.AutoRebalance
	}

	if !dara.IsNil(request.AzBalance) {
		query["AzBalance"] = request.AzBalance
	}

	if !dara.IsNil(request.BalanceMode) {
		query["BalanceMode"] = request.BalanceMode
	}

	if !dara.IsNil(request.CapacityOptions) {
		query["CapacityOptions"] = request.CapacityOptions
	}

	if !dara.IsNil(request.CompensateWithOnDemand) {
		query["CompensateWithOnDemand"] = request.CompensateWithOnDemand
	}

	if !dara.IsNil(request.CustomPolicyARN) {
		query["CustomPolicyARN"] = request.CustomPolicyARN
	}

	if !dara.IsNil(request.DefaultCooldown) {
		query["DefaultCooldown"] = request.DefaultCooldown
	}

	if !dara.IsNil(request.DesiredCapacity) {
		query["DesiredCapacity"] = request.DesiredCapacity
	}

	if !dara.IsNil(request.DisableDesiredCapacity) {
		query["DisableDesiredCapacity"] = request.DisableDesiredCapacity
	}

	if !dara.IsNil(request.GroupDeletionProtection) {
		query["GroupDeletionProtection"] = request.GroupDeletionProtection
	}

	if !dara.IsNil(request.HealthCheckType) {
		query["HealthCheckType"] = request.HealthCheckType
	}

	if !dara.IsNil(request.HealthCheckTypes) {
		query["HealthCheckTypes"] = request.HealthCheckTypes
	}

	if !dara.IsNil(request.LaunchTemplateId) {
		query["LaunchTemplateId"] = request.LaunchTemplateId
	}

	if !dara.IsNil(request.LaunchTemplateOverrides) {
		query["LaunchTemplateOverrides"] = request.LaunchTemplateOverrides
	}

	if !dara.IsNil(request.LaunchTemplateVersion) {
		query["LaunchTemplateVersion"] = request.LaunchTemplateVersion
	}

	if !dara.IsNil(request.MaxInstanceLifetime) {
		query["MaxInstanceLifetime"] = request.MaxInstanceLifetime
	}

	if !dara.IsNil(request.MaxSize) {
		query["MaxSize"] = request.MaxSize
	}

	if !dara.IsNil(request.MinSize) {
		query["MinSize"] = request.MinSize
	}

	if !dara.IsNil(request.MultiAZPolicy) {
		query["MultiAZPolicy"] = request.MultiAZPolicy
	}

	if !dara.IsNil(request.OnDemandBaseCapacity) {
		query["OnDemandBaseCapacity"] = request.OnDemandBaseCapacity
	}

	if !dara.IsNil(request.OnDemandPercentageAboveBaseCapacity) {
		query["OnDemandPercentageAboveBaseCapacity"] = request.OnDemandPercentageAboveBaseCapacity
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RemovalPolicies) {
		query["RemovalPolicies"] = request.RemovalPolicies
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.ScalingGroupName) {
		query["ScalingGroupName"] = request.ScalingGroupName
	}

	if !dara.IsNil(request.ScalingPolicy) {
		query["ScalingPolicy"] = request.ScalingPolicy
	}

	if !dara.IsNil(request.SpotAllocationStrategy) {
		query["SpotAllocationStrategy"] = request.SpotAllocationStrategy
	}

	if !dara.IsNil(request.SpotInstancePools) {
		query["SpotInstancePools"] = request.SpotInstancePools
	}

	if !dara.IsNil(request.SpotInstanceRemedy) {
		query["SpotInstanceRemedy"] = request.SpotInstanceRemedy
	}

	if !dara.IsNil(request.StopInstanceTimeout) {
		query["StopInstanceTimeout"] = request.StopInstanceTimeout
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	if !dara.IsNil(request.ZoneShiftPolicy) {
		query["ZoneShiftPolicy"] = request.ZoneShiftPolicy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyScalingGroup"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScalingGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies one or more attributes of a scaling rule. If an existing scaling rule no longer meets your requirements, you can call this operation to modify its attributes instead of creating a new one. This avoids unnecessary repetitive operations and saves time and resources.
//
// @param request - ModifyScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScalingRuleResponse
func (client *Client) ModifyScalingRuleWithContext(ctx context.Context, request *ModifyScalingRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdjustmentType) {
		query["AdjustmentType"] = request.AdjustmentType
	}

	if !dara.IsNil(request.AdjustmentValue) {
		query["AdjustmentValue"] = request.AdjustmentValue
	}

	if !dara.IsNil(request.AlarmDimensions) {
		query["AlarmDimensions"] = request.AlarmDimensions
	}

	if !dara.IsNil(request.AlarmOptions) {
		query["AlarmOptions"] = request.AlarmOptions
	}

	if !dara.IsNil(request.Cooldown) {
		query["Cooldown"] = request.Cooldown
	}

	if !dara.IsNil(request.DisableScaleIn) {
		query["DisableScaleIn"] = request.DisableScaleIn
	}

	if !dara.IsNil(request.EstimatedInstanceWarmup) {
		query["EstimatedInstanceWarmup"] = request.EstimatedInstanceWarmup
	}

	if !dara.IsNil(request.HybridMetrics) {
		query["HybridMetrics"] = request.HybridMetrics
	}

	if !dara.IsNil(request.HybridMonitorNamespace) {
		query["HybridMonitorNamespace"] = request.HybridMonitorNamespace
	}

	if !dara.IsNil(request.InitialMaxSize) {
		query["InitialMaxSize"] = request.InitialMaxSize
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.MinAdjustmentMagnitude) {
		query["MinAdjustmentMagnitude"] = request.MinAdjustmentMagnitude
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PredictiveScalingMode) {
		query["PredictiveScalingMode"] = request.PredictiveScalingMode
	}

	if !dara.IsNil(request.PredictiveTaskBufferTime) {
		query["PredictiveTaskBufferTime"] = request.PredictiveTaskBufferTime
	}

	if !dara.IsNil(request.PredictiveValueBehavior) {
		query["PredictiveValueBehavior"] = request.PredictiveValueBehavior
	}

	if !dara.IsNil(request.PredictiveValueBuffer) {
		query["PredictiveValueBuffer"] = request.PredictiveValueBuffer
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScaleInEvaluationCount) {
		query["ScaleInEvaluationCount"] = request.ScaleInEvaluationCount
	}

	if !dara.IsNil(request.ScaleOutEvaluationCount) {
		query["ScaleOutEvaluationCount"] = request.ScaleOutEvaluationCount
	}

	if !dara.IsNil(request.ScalingRuleId) {
		query["ScalingRuleId"] = request.ScalingRuleId
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	if !dara.IsNil(request.StepAdjustments) {
		query["StepAdjustments"] = request.StepAdjustments
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyScalingRule"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration parameters of an existing scheduled task, such as the scaling rule to execute, the number of instances in the scaling group, the task execution frequency, and the retry policy.
//
// Description:
//
// Scheduled tasks support two scaling methods:
//
// - Specify the scaling rule to execute by using the ScheduledAction parameter.
//
// - Specify the number of instances in the scaling group by using the ScalingGroupId parameter.
//
// > You cannot specify both ScheduledAction and ScalingGroupId at the same time.
//
// @param request - ModifyScheduledTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScheduledTaskResponse
func (client *Client) ModifyScheduledTaskWithContext(ctx context.Context, request *ModifyScheduledTaskRequest, runtime *dara.RuntimeOptions) (_result *ModifyScheduledTaskResponse, _err error) {
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

	if !dara.IsNil(request.DesiredCapacity) {
		query["DesiredCapacity"] = request.DesiredCapacity
	}

	if !dara.IsNil(request.LaunchExpirationTime) {
		query["LaunchExpirationTime"] = request.LaunchExpirationTime
	}

	if !dara.IsNil(request.LaunchTime) {
		query["LaunchTime"] = request.LaunchTime
	}

	if !dara.IsNil(request.MaxValue) {
		query["MaxValue"] = request.MaxValue
	}

	if !dara.IsNil(request.MinValue) {
		query["MinValue"] = request.MinValue
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecurrenceEndTime) {
		query["RecurrenceEndTime"] = request.RecurrenceEndTime
	}

	if !dara.IsNil(request.RecurrenceType) {
		query["RecurrenceType"] = request.RecurrenceType
	}

	if !dara.IsNil(request.RecurrenceValue) {
		query["RecurrenceValue"] = request.RecurrenceValue
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.ScheduledAction) {
		query["ScheduledAction"] = request.ScheduledAction
	}

	if !dara.IsNil(request.ScheduledTaskId) {
		query["ScheduledTaskId"] = request.ScheduledTaskId
	}

	if !dara.IsNil(request.ScheduledTaskName) {
		query["ScheduledTaskName"] = request.ScheduledTaskName
	}

	if !dara.IsNil(request.TaskEnabled) {
		query["TaskEnabled"] = request.TaskEnabled
	}

	if !dara.IsNil(request.WarmPoolCapacityOptions) {
		query["WarmPoolCapacityOptions"] = request.WarmPoolCapacityOptions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyScheduledTask"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a warm pool for a scaling group. If the scaling group does not have a warm pool, this operation creates one. If the scaling group already has a warm pool, this operation updates the configurations of the warm pool.
//
// Description:
//
// - Warm pools are supported only for ECS scaling groups.
//
// - Warm pools are supported only when the reclaim mode of the scaling group is set to release or forcerelease.
//
// @param request - PutWarmPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutWarmPoolResponse
func (client *Client) PutWarmPoolWithContext(ctx context.Context, request *PutWarmPoolRequest, runtime *dara.RuntimeOptions) (_result *PutWarmPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CapacityOptions) {
		query["CapacityOptions"] = request.CapacityOptions
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceReusePolicy) {
		query["InstanceReusePolicy"] = request.InstanceReusePolicy
	}

	if !dara.IsNil(request.PoolState) {
		query["PoolState"] = request.PoolState
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutWarmPool"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutWarmPoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical monitoring data of a scaling group that uses predictive scaling rules.
//
// @param request - QueryHistoricalMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHistoricalMetricResponse
func (client *Client) QueryHistoricalMetricWithContext(ctx context.Context, request *QueryHistoricalMetricRequest, runtime *dara.RuntimeOptions) (_result *QueryHistoricalMetricResponse, _err error) {
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

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryHistoricalMetric"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryHistoricalMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries predictive monitoring data.
//
// @param request - QueryPredictiveMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPredictiveMetricResponse
func (client *Client) QueryPredictiveMetricWithContext(ctx context.Context, request *QueryPredictiveMetricRequest, runtime *dara.RuntimeOptions) (_result *QueryPredictiveMetricResponse, _err error) {
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

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPredictiveMetric"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPredictiveMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a prediction task.
//
// @param request - QueryPredictiveTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPredictiveTaskInfoResponse
func (client *Client) QueryPredictiveTaskInfoWithContext(ctx context.Context, request *QueryPredictiveTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryPredictiveTaskInfoResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ScalingRuleId) {
		query["ScalingRuleId"] = request.ScalingRuleId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPredictiveTaskInfo"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPredictiveTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the predicted number of instances in a scaling group.
//
// @param request - QueryPredictiveValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPredictiveValueResponse
func (client *Client) QueryPredictiveValueWithContext(ctx context.Context, request *QueryPredictiveValueRequest, runtime *dara.RuntimeOptions) (_result *QueryPredictiveValueResponse, _err error) {
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

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPredictiveValue"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPredictiveValueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rebalances the distribution of ECS instances in a multi-zone scaling group if the number of instances across zones is unbalanced due to insufficient inventory or other reasons.
//
// Description:
//
// ## Operation description
//
// Rebalancing compensates for unbalanced zones by creating new ECS instances to replace existing ones. New ECS instances are started before existing ones are stopped. Rebalancing does not affect the performance or availability of your applications.
//
// - This operation is supported only for multi-zone scaling groups whose `MultiAZPolicy` is set to `BALANCE`. It is used to balance the distribution of ECS instances across multiple zones.
//
// - You can perform a rebalancing operation only when the instance distribution in the scaling group is severely unbalanced. A single rebalancing activity can replace up to 20 ECS instances.
//
// - During a rebalancing activity, if the number of instances in the scaling group approaches or reaches the specified maximum number of ECS instances (MaxSize) but the distribution still needs to be balanced, Auto Scaling allows the number of ECS instances to temporarily exceed MaxSize by 10%. This temporary excess lasts until the distribution is balanced, typically for 1 to 6 minutes.
//
//	> If 10% of the maximum number of instances in the scaling group is not an integer, the decimal part is rounded up to allow an excess of 1 instance. For example, if the maximum number of instances in the scaling group is 15, 10% is 1.5, and the temporary excess allowed is 2 instances.
//
// @param request - RebalanceInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebalanceInstancesResponse
func (client *Client) RebalanceInstancesWithContext(ctx context.Context, request *RebalanceInstancesRequest, runtime *dara.RuntimeOptions) (_result *RebalanceInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebalanceInstances"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebalanceInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Extends the suspension period of an ECS instance that is suspended after a lifecycle hook is triggered in a scaling group. If the default wait time of the suspended ECS instance is insufficient for your custom operations, you can call this operation to extend the suspension period by specifying the lifecycle hook ID, the wait state token of the scaling activity, and the new wait time.
//
// Description:
//
// You can call this operation to extend the wait time only when the ECS instance is suspended.
//
// The wait time of an ECS instance cannot exceed 6 hours, and each wait state can be extended up to 20 times.
//
// @param request - RecordLifecycleActionHeartbeatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecordLifecycleActionHeartbeatResponse
func (client *Client) RecordLifecycleActionHeartbeatWithContext(ctx context.Context, request *RecordLifecycleActionHeartbeatRequest, runtime *dara.RuntimeOptions) (_result *RecordLifecycleActionHeartbeatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.HeartbeatTimeout) {
		query["heartbeatTimeout"] = request.HeartbeatTimeout
	}

	if !dara.IsNil(request.LifecycleActionToken) {
		query["lifecycleActionToken"] = request.LifecycleActionToken
	}

	if !dara.IsNil(request.LifecycleHookId) {
		query["lifecycleHookId"] = request.LifecycleHookId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecordLifecycleActionHeartbeat"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecordLifecycleActionHeartbeatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more unnecessary instances from a scaling group. You can invoke the RemoveInstances API to delete instances if the scaling group is enabled and no scaling activity is being executed.
//
// Description:
//
// - Before you invoke this operation, ensure that the following conditions are met:
//
//   - The scaling group is enabled (Active).
//
//   - No scaling activity is being executed in the scaling group.
//
// > If no scaling activity is being executed in the scaling group, this operation can bypass the cooldown time (DefaultCooldown) and be executed directly.
//
// - If an ECS instance is created through automatic creation by Auto Scaling, or was manually added and managed by the scaling group, the ECS instance triggers the economical mode or is released when you delete it from the scaling group.
//
// - If an ECS instance was manually added and is not managed by the scaling group, the ECS instance is not stopped or released when you delete it from the scaling group.
//
// - If the number of instances in the scaling group (TotalCapacity) minus the number of instances specified in this operation is less than the minimum number of instances (MinSize) in the scaling group, the request to invoke this operation fails.
//
// A successful response only indicates that the Auto Scaling service has accepted the request and the scaling activity can be executed. It does not mean that the scaling activity will be executed successfully. You must check the execution status of the scaling activity by using the returned ScalingActivityId.
//
// @param tmpReq - RemoveInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveInstancesResponse
func (client *Client) RemoveInstancesWithContext(ctx context.Context, tmpReq *RemoveInstancesRequest, runtime *dara.RuntimeOptions) (_result *RemoveInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LifecycleHookContext) {
		request.LifecycleHookContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LifecycleHookContext, dara.String("LifecycleHookContext"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DecreaseDesiredCapacity) {
		query["DecreaseDesiredCapacity"] = request.DecreaseDesiredCapacity
	}

	if !dara.IsNil(request.IgnoreInvalidInstance) {
		query["IgnoreInvalidInstance"] = request.IgnoreInvalidInstance
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.LifecycleHookContextShrink) {
		query["LifecycleHookContext"] = request.LifecycleHookContextShrink
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemovePolicy) {
		query["RemovePolicy"] = request.RemovePolicy
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.StopInstanceTimeout) {
		query["StopInstanceTimeout"] = request.StopInstanceTimeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveInstances"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes an instance refresh task.
//
// @param request - ResumeInstanceRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeInstanceRefreshResponse
func (client *Client) ResumeInstanceRefreshWithContext(ctx context.Context, request *ResumeInstanceRefreshRequest, runtime *dara.RuntimeOptions) (_result *ResumeInstanceRefreshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceRefreshTaskId) {
		query["InstanceRefreshTaskId"] = request.InstanceRefreshTaskId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeInstanceRefresh"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeInstanceRefreshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes suspended processes in a scaling group so that the scaling group continues to execute the specified processes based on functional logic. For example, if you resume the health check process, an instance in the scaling group is automatically removed from the scaling group when it is considered unhealthy.
//
// @param request - ResumeProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeProcessesResponse
func (client *Client) ResumeProcessesWithContext(ctx context.Context, request *ResumeProcessesRequest, runtime *dara.RuntimeOptions) (_result *ResumeProcessesResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Processes) {
		query["Processes"] = request.Processes
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeProcesses"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeProcessesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back an instance refresh task if the desired configuration does not meet your requirements during task execution. During the rollback, the system uses the active configuration of the scaling group to create new instances to replace those created based on the desired configuration.
//
// Description:
//
// ## 前提条件
//
// 被回滚的实例刷新任务必须在启动时（调用 StartInstanceRefresh）设置了 `DesiredConfiguration` 参数。未设置 `DesiredConfiguration` 的实例刷新任务不支持回滚，调用本接口会返回 `UnsupportedOperation.DesiredConfiguration` 错误；如需回滚此类任务，请先取消该任务，再以 `DesiredConfiguration` 重新启动实例刷新后执行回滚。
//
// @param request - RollbackInstanceRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackInstanceRefreshResponse
func (client *Client) RollbackInstanceRefreshWithContext(ctx context.Context, request *RollbackInstanceRefreshRequest, runtime *dara.RuntimeOptions) (_result *RollbackInstanceRefreshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceRefreshTaskId) {
		query["InstanceRefreshTaskId"] = request.InstanceRefreshTaskId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackInstanceRefresh"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackInstanceRefreshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers auto scaling based on a specified adjustment rule by calling the ScaleWithAdjustment API, eliminating the need to create a scaling rule in advance as required by the ExecuteScalingRule API.
//
// Description:
//
// - Before you call this operation, ensure that the following conditions are met:
//
//   - The scaling group is in the Active state.
//
//   - No scaling activities are in progress in the scaling group.
//
// - If no scaling activities are in progress in the scaling group, this operation can bypass the cooldown time and directly execute a scaling activity.
//
// - If the number of ECS instances to be added specified in the scaling rule plus the current number of instances in the scaling group (Total Capacity) exceeds the maximum number of instances (MaxSize), the scaling activity is executed with Total Capacity set to MaxSize.
//
// - If the current number of instances in the scaling group (Total Capacity) minus the number of ECS instances to be removed specified in the scaling rule is less than the minimum number of instances (MinSize), the scaling activity is executed with Total Capacity set to MinSize.
//
// A successful response only indicates that Auto Scaling accepted the request and can execute the scaling activity. It does not guarantee that the scaling activity will be successfully executed. You must use the returned `ScalingActivityId` to check the execution status of the scaling activity.
//
// @param tmpReq - ScaleWithAdjustmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleWithAdjustmentResponse
func (client *Client) ScaleWithAdjustmentWithContext(ctx context.Context, tmpReq *ScaleWithAdjustmentRequest, runtime *dara.RuntimeOptions) (_result *ScaleWithAdjustmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ScaleWithAdjustmentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LifecycleHookContext) {
		request.LifecycleHookContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LifecycleHookContext, dara.String("LifecycleHookContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Overrides) {
		request.OverridesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Overrides, dara.String("Overrides"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityMetadata) {
		query["ActivityMetadata"] = request.ActivityMetadata
	}

	if !dara.IsNil(request.AdjustmentType) {
		query["AdjustmentType"] = request.AdjustmentType
	}

	if !dara.IsNil(request.AdjustmentValue) {
		query["AdjustmentValue"] = request.AdjustmentValue
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ExecutionMode) {
		query["ExecutionMode"] = request.ExecutionMode
	}

	if !dara.IsNil(request.LifecycleHookContextShrink) {
		query["LifecycleHookContext"] = request.LifecycleHookContextShrink
	}

	if !dara.IsNil(request.MinAdjustmentMagnitude) {
		query["MinAdjustmentMagnitude"] = request.MinAdjustmentMagnitude
	}

	if !dara.IsNil(request.OverridesShrink) {
		query["Overrides"] = request.OverridesShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParallelTask) {
		query["ParallelTask"] = request.ParallelTask
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.SyncActivity) {
		query["SyncActivity"] = request.SyncActivity
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScaleWithAdjustment"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ScaleWithAdjustmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets deletion protection for a scaling group. If you enable deletion protection for a scaling group, you cannot delete the scaling group. If you disable deletion protection for a scaling group, you can directly delete the scaling group. You can call the SetGroupDeletionProtection operation to enable or disable deletion protection.
//
// @param request - SetGroupDeletionProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetGroupDeletionProtectionResponse
func (client *Client) SetGroupDeletionProtectionWithContext(ctx context.Context, request *SetGroupDeletionProtectionRequest, runtime *dara.RuntimeOptions) (_result *SetGroupDeletionProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupDeletionProtection) {
		query["GroupDeletionProtection"] = request.GroupDeletionProtection
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetGroupDeletionProtection"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetGroupDeletionProtectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the health status of an ECS instance or an Elastic Container Instance (ECI) in a scaling group. In some cases, automatic health checks may not accurately detect the health status of ECS or ECI instances. Manually setting the health status enables quick identification and handling of faulty instances to better meet actual business and O&M requirements.
//
// Description:
//
// A scaling group detects and removes ECS instances or ECI instances that are in the unhealthy state. To retain specific ECS or ECI instances, put the instances into the standby state or protection status. For more information, see [EnterStandby](https://help.aliyun.com/document_detail/459345.html) and [SetInstancesProtection](https://help.aliyun.com/document_detail/459342.html).
//
// @param request - SetInstanceHealthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetInstanceHealthResponse
func (client *Client) SetInstanceHealthWithContext(ctx context.Context, request *SetInstanceHealthRequest, runtime *dara.RuntimeOptions) (_result *SetInstanceHealthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HealthStatus) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetInstanceHealth"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetInstanceHealthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets or removes the protection status for one or more ECS instances in a scaling group. After the protection feature is enabled for an ECS instance, the instance is not terminated or released during automatic scale-in of the scaling group. The instance can only be manually removed from the console or by calling the RemoveInstances operation.
//
// Description:
//
// After the protection status is enabled for an ECS instance:
//
// - The instance remains in this state until you disable the protection status.
//
// - The protected ECS instance is not removed even if changes in the number of instances in the scaling group or monitoring tasks trigger a scaling activity for automatic scale-in. You must remove the ECS instance before you can release it. For more information, see [Remove ECS instances](https://help.aliyun.com/document_detail/459393.html).
//
// - The health check status of the ECS instance is not updated when the instance is stopped or restarted.
//
// @param request - SetInstancesProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetInstancesProtectionResponse
func (client *Client) SetInstancesProtectionWithContext(ctx context.Context, request *SetInstancesProtectionRequest, runtime *dara.RuntimeOptions) (_result *SetInstancesProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProtectedFromScaleIn) {
		query["ProtectedFromScaleIn"] = request.ProtectedFromScaleIn
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetInstancesProtection"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetInstancesProtectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refreshes the instances in a scaling group to apply a new scaling configuration or update the images of the instances.
//
// Description:
//
// - A scaling group can have only one running instance refresh node at a time.
//
// - This feature is supported only for ECS scaling groups that use the **priority policy*	- as the scale-out policy. Scaling groups that use the **number of vCPUs to calculate the capacity of the scaling group*	- and scaling groups whose instance revoke pattern is set to **shutdown revoke pattern*	- or **forced shutdown revoke pattern*	- do not support this feature.
//
// - During an instance refresh, normal scale-out and scale-in activities can be performed. However, scale-out activities use the desired configuration specified for the instance refresh.
//
// - Instances that are manually added to the scaling group, or are in the standby state or protection status are ignored by the instance refresh node and are not refreshed.
//
// @param request - StartInstanceRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstanceRefreshResponse
func (client *Client) StartInstanceRefreshWithContext(ctx context.Context, request *StartInstanceRefreshRequest, runtime *dara.RuntimeOptions) (_result *StartInstanceRefreshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckpointPauseTime) {
		query["CheckpointPauseTime"] = request.CheckpointPauseTime
	}

	if !dara.IsNil(request.Checkpoints) {
		query["Checkpoints"] = request.Checkpoints
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DesiredConfiguration) {
		query["DesiredConfiguration"] = request.DesiredConfiguration
	}

	if !dara.IsNil(request.MaxHealthyPercentage) {
		query["MaxHealthyPercentage"] = request.MaxHealthyPercentage
	}

	if !dara.IsNil(request.MinHealthyPercentage) {
		query["MinHealthyPercentage"] = request.MinHealthyPercentage
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !dara.IsNil(request.SkipMatching) {
		query["SkipMatching"] = request.SkipMatching
	}

	if !dara.IsNil(request.Strategy) {
		query["Strategy"] = request.Strategy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartInstanceRefresh"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartInstanceRefreshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables zone fault degradation for a scaling group.
//
// @param request - StartScalingGroupZoneShiftRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartScalingGroupZoneShiftResponse
func (client *Client) StartScalingGroupZoneShiftWithContext(ctx context.Context, request *StartScalingGroupZoneShiftRequest, runtime *dara.RuntimeOptions) (_result *StartScalingGroupZoneShiftResponse, _err error) {
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

	if !dara.IsNil(request.ExpiryTime) {
		query["ExpiryTime"] = request.ExpiryTime
	}

	if !dara.IsNil(request.ImpairedZoneIds) {
		query["ImpairedZoneIds"] = request.ImpairedZoneIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartScalingGroupZoneShift"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartScalingGroupZoneShiftResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends an instance refresh task. If you are not sure that you want to roll back an ongoing instance refresh task whose configurations you think has an issue, you can call the SuspendInstanceRefresh operation to suspend the task.
//
// Description:
//
//	You cannot call this operation to suspend an instance refresh task that is being rolled back.
//
// @param request - SuspendInstanceRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendInstanceRefreshResponse
func (client *Client) SuspendInstanceRefreshWithContext(ctx context.Context, request *SuspendInstanceRefreshRequest, runtime *dara.RuntimeOptions) (_result *SuspendInstanceRefreshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceRefreshTaskId) {
		query["InstanceRefreshTaskId"] = request.InstanceRefreshTaskId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendInstanceRefresh"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendInstanceRefreshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends scaling processes. This operation empowers you to selectively pause distinct scaling processes within a particular scaling group, enabling you to carry out alternative tasks and achieve more granular management over your scaling operations.
//
// @param request - SuspendProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendProcessesResponse
func (client *Client) SuspendProcessesWithContext(ctx context.Context, request *SuspendProcessesRequest, runtime *dara.RuntimeOptions) (_result *SuspendProcessesResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Processes) {
		query["Processes"] = request.Processes
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScalingGroupId) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendProcesses"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendProcessesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and adds tags. You can call the TagResources operation to uniformly create and attach tags to your Auto Scaling resources, streamlining resource management. This capability empowers you to categorize resources based on tags, thereby enhancing the overall efficiency of resource allocation and utilization.
//
// Description:
//
//	  You can attach up to 20 tags to a scaling group.
//
//	    **
//
//	    **Note*	- Before you attach tags to a specific Auto Scaling resource, Alibaba Cloud automatically verifies the current number of tags attached to that resource. In the event the proposed addition would exceed the maximum allowed number of tags, an error message will be promptly returned after you call this operation.
//
//		- If you set `Tags.Propagate` to `true`, any tags attached to your scaling group will be automatically propagated to new instances that are subsequently created in the scaling group, without affecting existing instances.
//
//		- If both the scaling configuration and the scaling group have tags attached, and tag propagation from the scaling group is enabled, the tags of newly created instances comply with the following rules:
//
//	    	- Instances set to join the scaling group will inherit the following tags: tags attached to the scaling configuration that initiates the instance creation and tags attached to the scaling group that are allowed to propagate to these instances upon instance creation.
//
//	    	- If the tag keys of the scaling configuration and those attached to the scaling group and propagated to the instances are identical, the tags attached to the scaling group and propagated to the instances will be overwritten by the tags of the scaling configuration.
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
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2022-02-22"),
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
// Removes tags from Auto Scaling resources simultaneously. This operation streamlines resource management activities, enhances system efficiency, and mitigates potential security vulnerabilities. Once a tag is removed from a particular resource, and if it is not re-added to any other resource, the system will automatically delete the unused tag.
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeys) {
		query["TagKeys"] = request.TagKeys
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2022-02-22"),
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
// Checks whether the account has the AliyunServiceRoleForAutoScaling service-linked role by calling VerifyAuthentication and specifying the Alibaba Cloud account ID. This role is used to authorize Auto Scaling to manage ECS and ECI resources.
//
// @param request - VerifyAuthenticationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyAuthenticationResponse
func (client *Client) VerifyAuthenticationWithContext(ctx context.Context, request *VerifyAuthenticationRequest, runtime *dara.RuntimeOptions) (_result *VerifyAuthenticationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OnlyCheck) {
		query["OnlyCheck"] = request.OnlyCheck
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyAuthentication"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyAuthenticationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether Auto Scaling is activated. This operation guarantees that in response to shifts in business workloads or variations in incoming traffic, the system will automatically adjust resource provisioning. This auto-scaling capability enhances the overall system performance, ensuring high availability and improved flexibility to accommodate dynamic demands.
//
// @param request - VerifyUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyUserResponse
func (client *Client) VerifyUserWithContext(ctx context.Context, request *VerifyUserRequest, runtime *dara.RuntimeOptions) (_result *VerifyUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyUser"),
		Version:     dara.String("2022-02-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
