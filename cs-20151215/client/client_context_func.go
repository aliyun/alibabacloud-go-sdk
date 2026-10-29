// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds existing ECS instances to an ACK cluster. You can invoke this operation if you need to add a purchased cloud server to an ACK cluster as a worker node, or re-add a removed worker node instance to a node pool.
//
// Description:
//
// Read [Limits](https://help.aliyun.com/document_detail/86919.html) before calling this operation to ensure that the ECS instances to be added meet the requirements.
//
// @param request - AttachInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachInstancesResponse
func (client *Client) AttachInstancesWithContext(ctx context.Context, ClusterId *string, request *AttachInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AttachInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CpuPolicy) {
		body["cpu_policy"] = request.CpuPolicy
	}

	if !dara.IsNil(request.FormatDisk) {
		body["format_disk"] = request.FormatDisk
	}

	if !dara.IsNil(request.ImageId) {
		body["image_id"] = request.ImageId
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	if !dara.IsNil(request.IsEdgeWorker) {
		body["is_edge_worker"] = request.IsEdgeWorker
	}

	if !dara.IsNil(request.KeepInstanceName) {
		body["keep_instance_name"] = request.KeepInstanceName
	}

	if !dara.IsNil(request.KeyPair) {
		body["key_pair"] = request.KeyPair
	}

	if !dara.IsNil(request.NodepoolId) {
		body["nodepool_id"] = request.NodepoolId
	}

	if !dara.IsNil(request.Password) {
		body["password"] = request.Password
	}

	if !dara.IsNil(request.RdsInstances) {
		body["rds_instances"] = request.RdsInstances
	}

	if !dara.IsNil(request.Runtime) {
		body["runtime"] = request.Runtime
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.UserData) {
		body["user_data"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachInstances"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/attach"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Adds existing ECS instances to a node pool as worker nodes in an ACK cluster, or re-adds removed worker nodes to a node pool.
//
// @param request - AttachInstancesToNodePoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachInstancesToNodePoolResponse
func (client *Client) AttachInstancesToNodePoolWithContext(ctx context.Context, ClusterId *string, NodepoolId *string, request *AttachInstancesToNodePoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AttachInstancesToNodePoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FormatDisk) {
		body["format_disk"] = request.FormatDisk
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	if !dara.IsNil(request.KeepInstanceName) {
		body["keep_instance_name"] = request.KeepInstanceName
	}

	if !dara.IsNil(request.Password) {
		body["password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachInstancesToNodePool"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(NodepoolId)) + "/attach"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachInstancesToNodePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CancelClusterUpgrade is deprecated
//
// Summary:
//
// Cancels the upgrade of an ACK cluster.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelClusterUpgradeResponse
func (client *Client) CancelClusterUpgradeWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelClusterUpgradeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelClusterUpgrade"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/upgrade/cancel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CancelClusterUpgradeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CancelComponentUpgrade is deprecated
//
// Summary:
//
// Cancels the upgrade of a cluster component.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelComponentUpgradeResponse
func (client *Client) CancelComponentUpgradeWithContext(ctx context.Context, clusterId *string, componentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelComponentUpgradeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelComponentUpgrade"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/components/" + dara.PercentEncode(dara.StringValue(componentId)) + "/cancel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CancelComponentUpgradeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an existing but unexecuted automated O&M task execution plan.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelOperationPlanResponse
func (client *Client) CancelOperationPlanWithContext(ctx context.Context, planId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelOperationPlanResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelOperationPlan"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/operation/plans/" + dara.PercentEncode(dara.StringValue(planId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelOperationPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the execution of a cluster task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTaskResponse
func (client *Client) CancelTaskWithContext(ctx context.Context, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelTask"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks/" + dara.PercentEncode(dara.StringValue(taskId)) + "/cancel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log configuration of control plane components in an ACK managed cluster. ACK managed clusters support collecting control plane component logs and delivering them to your Simple Log Service (SLS) project. Control plane components include core components such as Kube API Server, Kube Scheduler, Kube Controller Manager, and Cloud Controller Manager. You can call this operation to query the current log configuration of control plane components, including the log retention period and collected components.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckControlPlaneLogEnableResponse
func (client *Client) CheckControlPlaneLogEnableWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckControlPlaneLogEnableResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckControlPlaneLogEnable"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/controlplanelog"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckControlPlaneLogEnableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the current service account is granted the specified service role permissions. Container Service for Kubernetes can call other cloud services related to the service role, such as ECS, OSS, NAS, and SLB, only after the corresponding role permissions are granted.
//
// @param request - CheckServiceRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceRoleResponse
func (client *Client) CheckServiceRoleWithContext(ctx context.Context, request *CheckServiceRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckServiceRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Roles) {
		body["roles"] = request.Roles
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckServiceRole"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ram/check-service-role"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckServiceRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the issued KubeConfig credentials and Role-Based Access Control (RBAC) permissions of a specified user in a specified cluster. You can call this operation to clean up and revoke authorization for KubeConfig files that pose security risks.
//
// Description:
//
// > 1. You must have the permissions to manage Container Service for Kubernetes (CS). The required policy is AliyunCSFullAccess.
//
// > 2. You cannot clear the permissions of an Alibaba Cloud account.
//
// > 3. You cannot clear the permissions of the user who performs this operation.
//
// @param request - CleanClusterUserPermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CleanClusterUserPermissionsResponse
func (client *Client) CleanClusterUserPermissionsWithContext(ctx context.Context, ClusterId *string, Uid *string, request *CleanClusterUserPermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CleanClusterUserPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CleanClusterUserPermissions"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cluster/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/user/" + dara.PercentEncode(dara.StringValue(Uid)) + "/permissions"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CleanClusterUserPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cleans up and revokes the KubeConfig credentials and RBAC permissions issued to a specified user, such as resigned employees or users with frozen accounts.
//
// Description:
//
// >- Requires the permissions to manage Container Service for Kubernetes (CS) (AliyunCSFullAccess).
//
// >- Clearing the permissions of Alibaba Cloud accounts is not supported.
//
// >- Clearing the permissions of the calling user is not supported.
//
// @param tmpReq - CleanUserPermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CleanUserPermissionsResponse
func (client *Client) CleanUserPermissionsWithContext(ctx context.Context, Uid *string, tmpReq *CleanUserPermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CleanUserPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CleanUserPermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ClusterIds) {
		request.ClusterIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterIds, dara.String("ClusterIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterIdsShrink) {
		query["ClusterIds"] = request.ClusterIdsShrink
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CleanUserPermissions"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/users/" + dara.PercentEncode(dara.StringValue(Uid)) + "/permissions"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CleanUserPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a self-healing rule.
//
// Description:
//
//	Notice: This API is not yet available. Stay tuned.
//
// @param request - CreateAutoRepairPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoRepairPolicyResponse
func (client *Client) CreateAutoRepairPolicyWithContext(ctx context.Context, clusterId *string, request *CreateAutoRepairPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAutoRepairPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ResourceSubType) {
		body["resource_sub_type"] = request.ResourceSubType
	}

	if !dara.IsNil(request.ResourceType) {
		body["resource_type"] = request.ResourceType
	}

	if !dara.IsNil(request.Rules) {
		body["rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAutoRepairPolicy"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/auto_repair_policies"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAutoRepairPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an elastic scaling configuration that enables the system to automatically increase or decrease compute resources based on the configured scaling rules to meet the workload demands of your cluster. During the creation procedure, you can specify scaling metrics and thresholds, scale-out order, cool-down period, and more.
//
// @param request - CreateAutoscalingConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoscalingConfigResponse
func (client *Client) CreateAutoscalingConfigWithContext(ctx context.Context, ClusterId *string, request *CreateAutoscalingConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAutoscalingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CoolDownDuration) {
		body["cool_down_duration"] = request.CoolDownDuration
	}

	if !dara.IsNil(request.DaemonsetEvictionForNodes) {
		body["daemonset_eviction_for_nodes"] = request.DaemonsetEvictionForNodes
	}

	if !dara.IsNil(request.Expander) {
		body["expander"] = request.Expander
	}

	if !dara.IsNil(request.GpuUtilizationThreshold) {
		body["gpu_utilization_threshold"] = request.GpuUtilizationThreshold
	}

	if !dara.IsNil(request.MaxGracefulTerminationSec) {
		body["max_graceful_termination_sec"] = request.MaxGracefulTerminationSec
	}

	if !dara.IsNil(request.MinReplicaCount) {
		body["min_replica_count"] = request.MinReplicaCount
	}

	if !dara.IsNil(request.Priorities) {
		body["priorities"] = request.Priorities
	}

	if !dara.IsNil(request.RecycleNodeDeletionEnabled) {
		body["recycle_node_deletion_enabled"] = request.RecycleNodeDeletionEnabled
	}

	if !dara.IsNil(request.ScaleDownEnabled) {
		body["scale_down_enabled"] = request.ScaleDownEnabled
	}

	if !dara.IsNil(request.ScaleUpFromZero) {
		body["scale_up_from_zero"] = request.ScaleUpFromZero
	}

	if !dara.IsNil(request.ScalerType) {
		body["scaler_type"] = request.ScalerType
	}

	if !dara.IsNil(request.ScanInterval) {
		body["scan_interval"] = request.ScanInterval
	}

	if !dara.IsNil(request.SkipNodesWithLocalStorage) {
		body["skip_nodes_with_local_storage"] = request.SkipNodesWithLocalStorage
	}

	if !dara.IsNil(request.SkipNodesWithSystemPods) {
		body["skip_nodes_with_system_pods"] = request.SkipNodesWithSystemPods
	}

	if !dara.IsNil(request.UnneededDuration) {
		body["unneeded_duration"] = request.UnneededDuration
	}

	if !dara.IsNil(request.UtilizationThreshold) {
		body["utilization_threshold"] = request.UtilizationThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAutoscalingConfig"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cluster/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/autoscale/config/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAutoscalingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ACK cluster, including ACK managed clusters, ACK Serverless clusters, ACK Edge clusters, and registered clusters. When you create a cluster, you configure cluster information, cluster components, and ACK-related cloud resources.
//
// Description:
//
// ### Generate OpenAPI request parameters by using the console
//
// When you call the CreateCluster operation to create a cluster, if the API call fails due to incorrect combinations of request parameters, you can use the console to generate the required request parameter combinations. Perform the following steps:
//
// 1. Log on to the [Container Service for Kubernetes (ACK) console](https://csnew.console.aliyun.com) and choose **Clusters*	- in the left-side navigation pane.
//
// 1. On the **Clusters*	- page, click **Cluster Template**.
//
// 1. In the dialog box, select the cluster type that you want to create, click Create, and then configure the cluster information on the cluster configuration page.
//
// 1. After the configuration is complete, on the **Confirm Configuration*	- page, click **Equivalent Code*	- in the upper-right corner. The dialog box displays the parameter combinations required for creating the cluster. You can copy and use them.
//
//	Notice: Starting from July 04, 2026, some request parameters will no longer take effect. For details about the changes and replacement parameter descriptions, see [Notice on changes to OpenAPI request and response parameters for ACK cluster management and OpenAPI offline](https://help.aliyun.com/document_detail/2932733.html).</notice>
//
// @param request - CreateClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessControlList) {
		body["access_control_list"] = request.AccessControlList
	}

	if !dara.IsNil(request.Addons) {
		body["addons"] = request.Addons
	}

	if !dara.IsNil(request.ApiAudiences) {
		body["api_audiences"] = request.ApiAudiences
	}

	if !dara.IsNil(request.AuditLogConfig) {
		body["audit_log_config"] = request.AuditLogConfig
	}

	if !dara.IsNil(request.AutoMode) {
		body["auto_mode"] = request.AutoMode
	}

	if !dara.IsNil(request.AutoRenew) {
		body["auto_renew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		body["auto_renew_period"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.ChargeType) {
		body["charge_type"] = request.ChargeType
	}

	if !dara.IsNil(request.CisEnabled) {
		body["cis_enabled"] = request.CisEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		body["client_token"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudMonitorFlags) {
		body["cloud_monitor_flags"] = request.CloudMonitorFlags
	}

	if !dara.IsNil(request.ClusterDomain) {
		body["cluster_domain"] = request.ClusterDomain
	}

	if !dara.IsNil(request.ClusterSpec) {
		body["cluster_spec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.ClusterType) {
		body["cluster_type"] = request.ClusterType
	}

	if !dara.IsNil(request.ContainerCidr) {
		body["container_cidr"] = request.ContainerCidr
	}

	if !dara.IsNil(request.ControlPlaneConfig) {
		body["control_plane_config"] = request.ControlPlaneConfig
	}

	if !dara.IsNil(request.ControlPlaneEndpointsConfig) {
		body["control_plane_endpoints_config"] = request.ControlPlaneEndpointsConfig
	}

	if !dara.IsNil(request.ControlplaneLogComponents) {
		body["controlplane_log_components"] = request.ControlplaneLogComponents
	}

	if !dara.IsNil(request.ControlplaneLogProject) {
		body["controlplane_log_project"] = request.ControlplaneLogProject
	}

	if !dara.IsNil(request.ControlplaneLogTtl) {
		body["controlplane_log_ttl"] = request.ControlplaneLogTtl
	}

	if !dara.IsNil(request.CpuPolicy) {
		body["cpu_policy"] = request.CpuPolicy
	}

	if !dara.IsNil(request.CustomSan) {
		body["custom_san"] = request.CustomSan
	}

	if !dara.IsNil(request.DeletionProtection) {
		body["deletion_protection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.DisableRollback) {
		body["disable_rollback"] = request.DisableRollback
	}

	if !dara.IsNil(request.EnableRrsa) {
		body["enable_rrsa"] = request.EnableRrsa
	}

	if !dara.IsNil(request.EncryptionProviderKey) {
		body["encryption_provider_key"] = request.EncryptionProviderKey
	}

	if !dara.IsNil(request.EndpointPublicAccess) {
		body["endpoint_public_access"] = request.EndpointPublicAccess
	}

	if !dara.IsNil(request.ExtraSans) {
		body["extra_sans"] = request.ExtraSans
	}

	if !dara.IsNil(request.FormatDisk) {
		body["format_disk"] = request.FormatDisk
	}

	if !dara.IsNil(request.ImageId) {
		body["image_id"] = request.ImageId
	}

	if !dara.IsNil(request.ImageType) {
		body["image_type"] = request.ImageType
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	if !dara.IsNil(request.IpStack) {
		body["ip_stack"] = request.IpStack
	}

	if !dara.IsNil(request.IsEnterpriseSecurityGroup) {
		body["is_enterprise_security_group"] = request.IsEnterpriseSecurityGroup
	}

	if !dara.IsNil(request.KeepInstanceName) {
		body["keep_instance_name"] = request.KeepInstanceName
	}

	if !dara.IsNil(request.KeyPair) {
		body["key_pair"] = request.KeyPair
	}

	if !dara.IsNil(request.KubernetesVersion) {
		body["kubernetes_version"] = request.KubernetesVersion
	}

	if !dara.IsNil(request.LoadBalancerId) {
		body["load_balancer_id"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LoadBalancerSpec) {
		body["load_balancer_spec"] = request.LoadBalancerSpec
	}

	if !dara.IsNil(request.LoggingType) {
		body["logging_type"] = request.LoggingType
	}

	if !dara.IsNil(request.LoginPassword) {
		body["login_password"] = request.LoginPassword
	}

	if !dara.IsNil(request.MaintenanceWindow) {
		body["maintenance_window"] = request.MaintenanceWindow
	}

	if !dara.IsNil(request.MasterAutoRenew) {
		body["master_auto_renew"] = request.MasterAutoRenew
	}

	if !dara.IsNil(request.MasterAutoRenewPeriod) {
		body["master_auto_renew_period"] = request.MasterAutoRenewPeriod
	}

	if !dara.IsNil(request.MasterCount) {
		body["master_count"] = request.MasterCount
	}

	if !dara.IsNil(request.MasterInstanceChargeType) {
		body["master_instance_charge_type"] = request.MasterInstanceChargeType
	}

	if !dara.IsNil(request.MasterInstanceTypes) {
		body["master_instance_types"] = request.MasterInstanceTypes
	}

	if !dara.IsNil(request.MasterPeriod) {
		body["master_period"] = request.MasterPeriod
	}

	if !dara.IsNil(request.MasterPeriodUnit) {
		body["master_period_unit"] = request.MasterPeriodUnit
	}

	if !dara.IsNil(request.MasterSystemDiskCategory) {
		body["master_system_disk_category"] = request.MasterSystemDiskCategory
	}

	if !dara.IsNil(request.MasterSystemDiskPerformanceLevel) {
		body["master_system_disk_performance_level"] = request.MasterSystemDiskPerformanceLevel
	}

	if !dara.IsNil(request.MasterSystemDiskSize) {
		body["master_system_disk_size"] = request.MasterSystemDiskSize
	}

	if !dara.IsNil(request.MasterSystemDiskSnapshotPolicyId) {
		body["master_system_disk_snapshot_policy_id"] = request.MasterSystemDiskSnapshotPolicyId
	}

	if !dara.IsNil(request.MasterVswitchIds) {
		body["master_vswitch_ids"] = request.MasterVswitchIds
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.NatGateway) {
		body["nat_gateway"] = request.NatGateway
	}

	if !dara.IsNil(request.NodeCidrMask) {
		body["node_cidr_mask"] = request.NodeCidrMask
	}

	if !dara.IsNil(request.NodeNameMode) {
		body["node_name_mode"] = request.NodeNameMode
	}

	if !dara.IsNil(request.NodePortRange) {
		body["node_port_range"] = request.NodePortRange
	}

	if !dara.IsNil(request.Nodepools) {
		body["nodepools"] = request.Nodepools
	}

	if !dara.IsNil(request.NumOfNodes) {
		body["num_of_nodes"] = request.NumOfNodes
	}

	if !dara.IsNil(request.OperationPolicy) {
		body["operation_policy"] = request.OperationPolicy
	}

	if !dara.IsNil(request.OsType) {
		body["os_type"] = request.OsType
	}

	if !dara.IsNil(request.Period) {
		body["period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["period_unit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.Platform) {
		body["platform"] = request.Platform
	}

	if !dara.IsNil(request.PodVswitchIds) {
		body["pod_vswitch_ids"] = request.PodVswitchIds
	}

	if !dara.IsNil(request.Profile) {
		body["profile"] = request.Profile
	}

	if !dara.IsNil(request.ProxyMode) {
		body["proxy_mode"] = request.ProxyMode
	}

	if !dara.IsNil(request.RdsInstances) {
		body["rds_instances"] = request.RdsInstances
	}

	if !dara.IsNil(request.RegionId) {
		body["region_id"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resource_group_id"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RrsaConfig) {
		body["rrsa_config"] = request.RrsaConfig
	}

	if !dara.IsNil(request.Runtime) {
		body["runtime"] = request.Runtime
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["security_group_id"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityHardeningOs) {
		body["security_hardening_os"] = request.SecurityHardeningOs
	}

	if !dara.IsNil(request.ServiceAccountIssuer) {
		body["service_account_issuer"] = request.ServiceAccountIssuer
	}

	if !dara.IsNil(request.ServiceCidr) {
		body["service_cidr"] = request.ServiceCidr
	}

	if !dara.IsNil(request.ServiceDiscoveryTypes) {
		body["service_discovery_types"] = request.ServiceDiscoveryTypes
	}

	if !dara.IsNil(request.SnatEntry) {
		body["snat_entry"] = request.SnatEntry
	}

	if !dara.IsNil(request.SocEnabled) {
		body["soc_enabled"] = request.SocEnabled
	}

	if !dara.IsNil(request.SshFlags) {
		body["ssh_flags"] = request.SshFlags
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.Taints) {
		body["taints"] = request.Taints
	}

	if !dara.IsNil(request.TimeoutMins) {
		body["timeout_mins"] = request.TimeoutMins
	}

	if !dara.IsNil(request.Timezone) {
		body["timezone"] = request.Timezone
	}

	if !dara.IsNil(request.UserCa) {
		body["user_ca"] = request.UserCa
	}

	if !dara.IsNil(request.UserData) {
		body["user_data"] = request.UserData
	}

	if !dara.IsNil(request.Vpcid) {
		body["vpcid"] = request.Vpcid
	}

	if !dara.IsNil(request.VswitchIds) {
		body["vswitch_ids"] = request.VswitchIds
	}

	if !dara.IsNil(request.WorkerAutoRenew) {
		body["worker_auto_renew"] = request.WorkerAutoRenew
	}

	if !dara.IsNil(request.WorkerAutoRenewPeriod) {
		body["worker_auto_renew_period"] = request.WorkerAutoRenewPeriod
	}

	if !dara.IsNil(request.WorkerDataDisks) {
		body["worker_data_disks"] = request.WorkerDataDisks
	}

	if !dara.IsNil(request.WorkerInstanceChargeType) {
		body["worker_instance_charge_type"] = request.WorkerInstanceChargeType
	}

	if !dara.IsNil(request.WorkerInstanceTypes) {
		body["worker_instance_types"] = request.WorkerInstanceTypes
	}

	if !dara.IsNil(request.WorkerPeriod) {
		body["worker_period"] = request.WorkerPeriod
	}

	if !dara.IsNil(request.WorkerPeriodUnit) {
		body["worker_period_unit"] = request.WorkerPeriodUnit
	}

	if !dara.IsNil(request.WorkerSystemDiskCategory) {
		body["worker_system_disk_category"] = request.WorkerSystemDiskCategory
	}

	if !dara.IsNil(request.WorkerSystemDiskPerformanceLevel) {
		body["worker_system_disk_performance_level"] = request.WorkerSystemDiskPerformanceLevel
	}

	if !dara.IsNil(request.WorkerSystemDiskSize) {
		body["worker_system_disk_size"] = request.WorkerSystemDiskSize
	}

	if !dara.IsNil(request.WorkerSystemDiskSnapshotPolicyId) {
		body["worker_system_disk_snapshot_policy_id"] = request.WorkerSystemDiskSnapshotPolicyId
	}

	if !dara.IsNil(request.WorkerVswitchIds) {
		body["worker_vswitch_ids"] = request.WorkerVswitchIds
	}

	if !dara.IsNil(request.ZoneId) {
		body["zone_id"] = request.ZoneId
	}

	if !dara.IsNil(request.ZoneIds) {
		body["zone_ids"] = request.ZoneIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCluster"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a cluster diagnosis.
//
// @param request - CreateClusterDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterDiagnosisResponse
func (client *Client) CreateClusterDiagnosisWithContext(ctx context.Context, clusterId *string, request *CreateClusterDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterDiagnosisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Target) {
		body["target"] = request.Target
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClusterDiagnosis"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/diagnosis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterDiagnosisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cluster inspection configuration.
//
// @param request - CreateClusterInspectConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterInspectConfigResponse
func (client *Client) CreateClusterInspectConfigWithContext(ctx context.Context, clusterId *string, request *CreateClusterInspectConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterInspectConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DisabledCheckItems) {
		body["disabledCheckItems"] = request.DisabledCheckItems
	}

	if !dara.IsNil(request.Enabled) {
		body["enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Recurrence) {
		body["recurrence"] = request.Recurrence
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClusterInspectConfig"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/inspectConfig"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterInspectConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a node pool. A node pool is a logical collection of nodes with the same property, which allows unified management and O&M of nodes, such as node upgrades and elastic scaling. You can further use the automated O&M features of node pools, such as automatic repair of OS CVE vulnerabilities, automatic recovery of faulty nodes, and automatic upgrades of kubelet and containerd versions, to reduce O&M costs. You can invoke CreateClusterNodePool to create a node pool for a cluster.
//
// @param request - CreateClusterNodePoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterNodePoolResponse
func (client *Client) CreateClusterNodePoolWithContext(ctx context.Context, ClusterId *string, request *CreateClusterNodePoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterNodePoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoMode) {
		body["auto_mode"] = request.AutoMode
	}

	if !dara.IsNil(request.AutoScaling) {
		body["auto_scaling"] = request.AutoScaling
	}

	if !dara.IsNil(request.Count) {
		body["count"] = request.Count
	}

	if !dara.IsNil(request.EfloNodeGroup) {
		body["eflo_node_group"] = request.EfloNodeGroup
	}

	if !dara.IsNil(request.HostNetwork) {
		body["host_network"] = request.HostNetwork
	}

	if !dara.IsNil(request.InterconnectConfig) {
		body["interconnect_config"] = request.InterconnectConfig
	}

	if !dara.IsNil(request.InterconnectMode) {
		body["interconnect_mode"] = request.InterconnectMode
	}

	if !dara.IsNil(request.Intranet) {
		body["intranet"] = request.Intranet
	}

	if !dara.IsNil(request.KubernetesConfig) {
		body["kubernetes_config"] = request.KubernetesConfig
	}

	if !dara.IsNil(request.Management) {
		body["management"] = request.Management
	}

	if !dara.IsNil(request.MaxNodes) {
		body["max_nodes"] = request.MaxNodes
	}

	if !dara.IsNil(request.NodeComponents) {
		body["node_components"] = request.NodeComponents
	}

	if !dara.IsNil(request.NodeConfig) {
		body["node_config"] = request.NodeConfig
	}

	if !dara.IsNil(request.NodepoolInfo) {
		body["nodepool_info"] = request.NodepoolInfo
	}

	if !dara.IsNil(request.ScalingGroup) {
		body["scaling_group"] = request.ScalingGroup
	}

	if !dara.IsNil(request.TeeConfig) {
		body["tee_config"] = request.TeeConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClusterNodePool"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterNodePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateKubernetesTrigger is deprecated
//
// Summary:
//
// Creates a trigger for an application.
//
// @param request - CreateKubernetesTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKubernetesTriggerResponse
func (client *Client) CreateKubernetesTriggerWithContext(ctx context.Context, request *CreateKubernetesTriggerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKubernetesTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["action"] = request.Action
	}

	if !dara.IsNil(request.ClusterId) {
		body["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.ProjectId) {
		body["project_id"] = request.ProjectId
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKubernetesTrigger"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/triggers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKubernetesTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs node components.
//
// @param request - CreateNodePoolComponentInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodePoolComponentInstancesResponse
func (client *Client) CreateNodePoolComponentInstancesWithContext(ctx context.Context, clusterId *string, nodepoolId *string, request *CreateNodePoolComponentInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateNodePoolComponentInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Components) {
		body["components"] = request.Components
	}

	if !dara.IsNil(request.NodeNames) {
		body["node_names"] = request.NodeNames
	}

	if !dara.IsNil(request.RollingPolicy) {
		body["rolling_policy"] = request.RollingPolicy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNodePoolComponentInstances"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(nodepoolId)) + "/component_instances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNodePoolComponentInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an orchestration template. An orchestration template defines and describes a set of Kubernetes cluster resources in a declarative manner to specify how applications should run or be configured. You can use these templates to automate deployment and cluster management of resources, such as Pods, Services, Deployments, ConfigMaps, and PersistentVolumes. You can invoke the CreateTemplate API operation to create an orchestration template.
//
// @param request - CreateTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithContext(ctx context.Context, request *CreateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.Template) {
		body["template"] = request.Template
	}

	if !dara.IsNil(request.TemplateType) {
		body["template_type"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTemplate"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/templates"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Configures a trigger for a cluster application to redeploy pods when specific conditions are met.
//
// Description:
//
// > Creating a trigger only supports pod redeployment.
//
// @param request - CreateTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTriggerResponse
func (client *Client) CreateTriggerWithContext(ctx context.Context, clusterId *string, request *CreateTriggerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["action"] = request.Action
	}

	if !dara.IsNil(request.ClusterId) {
		body["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.ProjectId) {
		body["project_id"] = request.ProjectId
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrigger"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/triggers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes ACK alert contacts.
//
// @param tmpReq - DeleteAlertContactRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertContactResponse
func (client *Client) DeleteAlertContactWithContext(ctx context.Context, tmpReq *DeleteAlertContactRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlertContactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAlertContactShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ContactIds) {
		request.ContactIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContactIds, dara.String("contact_ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactIdsShrink) {
		query["contact_ids"] = request.ContactIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlertContact"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/alert/contacts"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DeleteAlertContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes ACK alert contact groups.
//
// @param tmpReq - DeleteAlertContactGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertContactGroupResponse
func (client *Client) DeleteAlertContactGroupWithContext(ctx context.Context, tmpReq *DeleteAlertContactGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlertContactGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAlertContactGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ContactGroupIds) {
		request.ContactGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContactGroupIds, dara.String("contact_group_ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupIdsShrink) {
		query["contact_group_ids"] = request.ContactGroupIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlertContactGroup"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/alert/contact_groups"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DeleteAlertContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a self-healing rule.
//
// Description:
//
//	Notice: This API is not yet available. Stay tuned.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoRepairPolicyResponse
func (client *Client) DeleteAutoRepairPolicyWithContext(ctx context.Context, clusterId *string, policyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAutoRepairPolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAutoRepairPolicy"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/auto_repair_policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAutoRepairPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a cluster and specifies whether to delete or retain the associated resources of the cluster. Before you delete a cluster, manually clean up workloads such as stateless, stateful, job, and cron job workloads. Otherwise, the cluster deletion may fail.
//
// Description:
//
// Risk warnings:
//
// - Subscription ECS instances and Lingjun computing nodes in the cluster cannot be automatically released. To avoid unnecessary costs, perform a manual release. For more information, see Cluster deletion and node release rules.
//
// - Subscription API server SLB resources cannot be automatically released. To avoid unnecessary costs, perform a manual release.
//
// - VPCs, vSwitches, security groups, and RAM roles are retained by default if they are used by other resources and cannot be deleted. Perform a manual release.
//
// - Elastic Container Instance (ECI) pods created by using virtual nodes are automatically released.
//
// - Some cloud resources created by the cluster are not automatically released when the cluster is deleted. These resources continue to incur charges after the cluster is deleted. Release or retain them based on your business requirements. These resources include: Simple Log Service (SLS) projects automatically created by the cluster, and cloud disks purchased by using dynamic volumes.
//
// @param tmpReq - DeleteClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithContext(ctx context.Context, ClusterId *string, tmpReq *DeleteClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteOptions) {
		request.DeleteOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteOptions, dara.String("delete_options"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RetainResources) {
		request.RetainResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RetainResources, dara.String("retain_resources"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteOptionsShrink) {
		query["delete_options"] = request.DeleteOptionsShrink
	}

	if !dara.IsNil(request.KeepSlb) {
		query["keep_slb"] = request.KeepSlb
	}

	if !dara.IsNil(request.RetainAllResources) {
		query["retain_all_resources"] = request.RetainAllResources
	}

	if !dara.IsNil(request.RetainResourcesShrink) {
		query["retain_resources"] = request.RetainResourcesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCluster"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the cluster inspection configuration.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterInspectConfigResponse
func (client *Client) DeleteClusterInspectConfigWithContext(ctx context.Context, clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteClusterInspectConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClusterInspectConfig"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/inspectConfig"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClusterInspectConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a node pool that is no longer in use. When a node pool is deleted, all pods on its nodes are deleted, which may trigger rescheduling. If rescheduling fails, your services may be affected. Ensure that the cluster has sufficient resources for rescheduling.
//
// @param request - DeleteClusterNodepoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterNodepoolResponse
func (client *Client) DeleteClusterNodepoolWithContext(ctx context.Context, ClusterId *string, NodepoolId *string, request *DeleteClusterNodepoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteClusterNodepoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClusterNodepool"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(NodepoolId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClusterNodepoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes nodes from a cluster. When you remove nodes, you can specify whether to release the associated ECS instances and whether to automatically drain the nodes.
//
// Description:
//
// - When removing nodes, perform standard operations by calling this API operation or using the [Container Service console](https://cs.console.aliyun.com). Do not manually remove nodes by running the `kubectl delete node` command.
//
// - Do not manually release nodes or remove instances in the ECS or Auto Scaling (ESS) console or by calling related API operations, or passively release nodes when a subscription instance expires. In these cases, the nodes are directly shut down and automatically removed from the Container Service console.
//
// - If the expected number of nodes is configured for a node pool, the node pool automatically scales out other instances based on the corresponding configurations to maintain the expected number of nodes.
//
// - Removing nodes involves pod migration and may affect your services. Perform this operation during off-peak hours. Unexpected risks may occur during the operation. Back up your data in advance.
//
// - When you remove nodes, ACK executes drain operations. Make sure that other nodes in the cluster have sufficient resources to avoid application pod scheduling failures.
//
// - Check the node affinity rules and scheduling policies of the pods on the nodes to be removed to ensure that the pods can be scheduled to other nodes after the nodes are removed.
//
// @param request - DeleteClusterNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterNodesResponse
func (client *Client) DeleteClusterNodesWithContext(ctx context.Context, ClusterId *string, request *DeleteClusterNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteClusterNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DrainNode) {
		body["drain_node"] = request.DrainNode
	}

	if !dara.IsNil(request.Nodes) {
		body["nodes"] = request.Nodes
	}

	if !dara.IsNil(request.ReleaseNode) {
		body["release_node"] = request.ReleaseNode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClusterNodes"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClusterNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteKubernetesTrigger is deprecated
//
// Summary:
//
// Deletes an application trigger based on the application trigger ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKubernetesTriggerResponse
func (client *Client) DeleteKubernetesTriggerWithContext(ctx context.Context, Id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteKubernetesTriggerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteKubernetesTrigger"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/triggers/revoke/" + dara.PercentEncode(dara.StringValue(Id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteKubernetesTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls a node component.
//
// @param tmpReq - DeleteNodePoolComponentInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodePoolComponentInstanceResponse
func (client *Client) DeleteNodePoolComponentInstanceWithContext(ctx context.Context, clusterId *string, nodepoolId *string, name *string, tmpReq *DeleteNodePoolComponentInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteNodePoolComponentInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteNodePoolComponentInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeNames) {
		request.NodeNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeNames, dara.String("node_names"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchInterval) {
		query["batch_interval"] = request.BatchInterval
	}

	if !dara.IsNil(request.MaxFailedNodes) {
		query["max_failed_nodes"] = request.MaxFailedNodes
	}

	if !dara.IsNil(request.MaxParallelism) {
		query["max_parallelism"] = request.MaxParallelism
	}

	if !dara.IsNil(request.NodeNamesShrink) {
		query["node_names"] = request.NodeNamesShrink
	}

	if !dara.IsNil(request.PausePolicy) {
		query["pause_policy"] = request.PausePolicy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNodePoolComponentInstance"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(nodepoolId)) + "/component_instances/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNodePoolComponentInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a policy rule instance in a specified cluster.
//
// @param request - DeletePolicyInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyInstanceResponse
func (client *Client) DeletePolicyInstanceWithContext(ctx context.Context, clusterId *string, policyName *string, request *DeletePolicyInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePolicyInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		query["instance_name"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicyInstance"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/policies/" + dara.PercentEncode(dara.StringValue(policyName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an orchestration template.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithContext(ctx context.Context, TemplateId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplate"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/templates/" + dara.PercentEncode(dara.StringValue(TemplateId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application trigger.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTriggerResponse
func (client *Client) DeleteTriggerWithContext(ctx context.Context, clusterId *string, Id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTriggerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrigger"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/triggers/" + dara.PercentEncode(dara.StringValue(Id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys a policy rule instance in a specified namespace of an ACK cluster. You can select the security policy type, configure the governance action (alert or block) for the policy instance, and specify the namespace scope.
//
// @param request - DeployPolicyInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployPolicyInstanceResponse
func (client *Client) DeployPolicyInstanceWithContext(ctx context.Context, clusterId *string, policyName *string, request *DeployPolicyInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeployPolicyInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["action"] = request.Action
	}

	if !dara.IsNil(request.Namespaces) {
		body["namespaces"] = request.Namespaces
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployPolicyInstance"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/policies/" + dara.PercentEncode(dara.StringValue(policyName))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployPolicyInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a target component based on parameters such as the specified region, cluster type, cluster profile, cluster version, and component name. The information includes whether the component is managed, the component category, the supported custom parameter schema, the compatible operating system architecture, and the minimum cluster version required by the component version.
//
// @param request - DescribeAddonRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAddonResponse
func (client *Client) DescribeAddonWithContext(ctx context.Context, addonName *string, request *DescribeAddonRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAddonResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterSpec) {
		query["cluster_spec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.ClusterType) {
		query["cluster_type"] = request.ClusterType
	}

	if !dara.IsNil(request.ClusterVersion) {
		query["cluster_version"] = request.ClusterVersion
	}

	if !dara.IsNil(request.Profile) {
		query["profile"] = request.Profile
	}

	if !dara.IsNil(request.RegionId) {
		query["region_id"] = request.RegionId
	}

	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAddon"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/addons/" + dara.PercentEncode(dara.StringValue(addonName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAddonResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeAddons is deprecated
//
// Summary:
//
// Queries the details of all components supported by the platform.
//
// @param request - DescribeAddonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAddonsResponse
func (client *Client) DescribeAddonsWithContext(ctx context.Context, request *DescribeAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAddonsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterProfile) {
		query["cluster_profile"] = request.ClusterProfile
	}

	if !dara.IsNil(request.ClusterSpec) {
		query["cluster_spec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.ClusterType) {
		query["cluster_type"] = request.ClusterType
	}

	if !dara.IsNil(request.ClusterVersion) {
		query["cluster_version"] = request.ClusterVersion
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAddons"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/components/metadata"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAddonsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries self-healing rules.
//
// Description:
//
//	Notice: This API is not yet available. Stay tuned.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoRepairPolicyResponse
func (client *Client) DescribeAutoRepairPolicyWithContext(ctx context.Context, clusterId *string, policyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAutoRepairPolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAutoRepairPolicy"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/auto_repair_policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoRepairPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusterAddonInstance is deprecated
//
// Summary:
//
// Queries the version, state, and configuration of installed cluster add-ons.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAddonInstanceResponse
func (client *Client) DescribeClusterAddonInstanceWithContext(ctx context.Context, ClusterID *string, AddonName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterAddonInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterAddonInstance"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterID)) + "/components/" + dara.PercentEncode(dara.StringValue(AddonName)) + "/instance"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterAddonInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusterAddonMetadata is deprecated
//
// Summary:
//
// Queries the metadata of a specified component version available for a specific cluster, including the component version and configurable parameters.
//
// @param request - DescribeClusterAddonMetadataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAddonMetadataResponse
func (client *Client) DescribeClusterAddonMetadataWithContext(ctx context.Context, clusterId *string, componentId *string, request *DescribeClusterAddonMetadataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterAddonMetadataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterAddonMetadata"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/components/" + dara.PercentEncode(dara.StringValue(componentId)) + "/metadata"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterAddonMetadataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusterAddonUpgradeStatus is deprecated
//
// Summary:
//
// Queries the upgrade status of cluster add-ons.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAddonUpgradeStatusResponse
func (client *Client) DescribeClusterAddonUpgradeStatusWithContext(ctx context.Context, ClusterId *string, ComponentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterAddonUpgradeStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterAddonUpgradeStatus"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/components/" + dara.PercentEncode(dara.StringValue(ComponentId)) + "/upgradestatus"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterAddonUpgradeStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusterAddonsUpgradeStatus is deprecated
//
// Summary:
//
// Queries the upgrade status of a component by name.
//
// @param tmpReq - DescribeClusterAddonsUpgradeStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAddonsUpgradeStatusResponse
func (client *Client) DescribeClusterAddonsUpgradeStatusWithContext(ctx context.Context, ClusterId *string, tmpReq *DescribeClusterAddonsUpgradeStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterAddonsUpgradeStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeClusterAddonsUpgradeStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ComponentIds) {
		request.ComponentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComponentIds, dara.String("componentIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ComponentIdsShrink) {
		query["componentIds"] = request.ComponentIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterAddonsUpgradeStatus"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/components/upgradestatus"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterAddonsUpgradeStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusterAddonsVersion is deprecated
//
// Summary:
//
// Queries the details of all installed components in a cluster by cluster ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAddonsVersionResponse
func (client *Client) DescribeClusterAddonsVersionWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterAddonsVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterAddonsVersion"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/components/version"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterAddonsVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the script used to add existing ECS instances to an ACK cluster as worker nodes, or to re-add removed worker nodes to a node pool. ACK supports manually adding existing nodes to a node pool. You can invoke the DescribeClusterAttachScripts operation to obtain the script for adding existing nodes.
//
// @param request - DescribeClusterAttachScriptsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAttachScriptsResponse
func (client *Client) DescribeClusterAttachScriptsWithContext(ctx context.Context, ClusterId *string, request *DescribeClusterAttachScriptsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterAttachScriptsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Arch) {
		body["arch"] = request.Arch
	}

	if !dara.IsNil(request.Expired) {
		body["expired"] = request.Expired
	}

	if !dara.IsNil(request.FormatDisk) {
		body["format_disk"] = request.FormatDisk
	}

	if !dara.IsNil(request.KeepInstanceName) {
		body["keep_instance_name"] = request.KeepInstanceName
	}

	if !dara.IsNil(request.NodepoolId) {
		body["nodepool_id"] = request.NodepoolId
	}

	if !dara.IsNil(request.OneTimeToken) {
		body["one_time_token"] = request.OneTimeToken
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	if !dara.IsNil(request.RdsInstances) {
		body["rds_instances"] = request.RdsInstances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterAttachScripts"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/attachscript"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("string"),
	}
	_result = &DescribeClusterAttachScriptsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified cluster by cluster ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterDetailResponse
func (client *Client) DescribeClusterDetailWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterDetailResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterDetail"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the events of a specified cluster. Cluster operation events include cluster creation, cluster modification, node pool creation, node pool scaling, add-on installation, and cluster upgrades. You can call this operation to retrieve a list of events that occurred in the cluster and query event details, including the event level, event status, and event time.
//
// @param request - DescribeClusterEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterEventsResponse
func (client *Client) DescribeClusterEventsWithContext(ctx context.Context, ClusterId *string, request *DescribeClusterEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["max_results"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["next_token"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["page_number"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterEvents"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the log data of a cluster for root cause analysis and tracing when cluster issues occur.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterLogsResponse
func (client *Client) DescribeClusterLogsWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterLogsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterLogs"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DescribeClusterLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of a specified node pool in a cluster.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterNodePoolDetailResponse
func (client *Client) DescribeClusterNodePoolDetailWithContext(ctx context.Context, ClusterId *string, NodepoolId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterNodePoolDetailResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterNodePoolDetail"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(NodepoolId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterNodePoolDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all node pools in a cluster.
//
// Description:
//
// - When you remove a node, perform standard operations by using this API operation or the [Container Service console](https://cs.console.aliyun.com). Do not manually remove nodes by running the `kubectl delete node` command.
//
// - Do not release nodes or remove instances in the ECS or Auto Scaling console, or by calling related API operations, or passively release nodes when a subscription instance expires. In these cases, the nodes are directly shut down and automatically removed from the Container Service console.
//
// - If the expected number of nodes is configured for a node pool, the node pool automatically scales out other instances based on the corresponding configurations to maintain the expected number of nodes.
//
// - Removing a node involves pod migration, which may affect your services. Perform this operation during off-peak hours. Unexpected risks may occur during the operation. Back up your data in advance.
//
// - When a node is removed, ACK drains the node. Make sure that other nodes in the cluster have sufficient resources to prevent application pods from failing to be scheduled.
//
// - Check the node affinity rules and scheduling policies of the pods on the node to be removed. Make sure that the pods can be scheduled to other nodes after the node is removed.
//
// @param request - DescribeClusterNodePoolsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterNodePoolsResponse
func (client *Client) DescribeClusterNodePoolsWithContext(ctx context.Context, ClusterId *string, request *DescribeClusterNodePoolsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterNodePoolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodepoolName) {
		query["NodepoolName"] = request.NodepoolName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterNodePools"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterNodePoolsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of nodes that meet the specified conditions.
//
// @param request - DescribeClusterNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterNodesResponse
func (client *Client) DescribeClusterNodesWithContext(ctx context.Context, ClusterId *string, request *DescribeClusterNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["instanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.NodeIps) {
		query["nodeIps"] = request.NodeIps
	}

	if !dara.IsNil(request.NodeLabels) {
		query["nodeLabels"] = request.NodeLabels
	}

	if !dara.IsNil(request.NodeNames) {
		query["nodeNames"] = request.NodeNames
	}

	if !dara.IsNil(request.NodepoolId) {
		query["nodepool_id"] = request.NodepoolId
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.State) {
		query["state"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterNodes"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the associated resources of a specified cluster. When you use Container Service for Kubernetes (ACK), you also use resources of other associated Alibaba Cloud services. You can call this operation to query these associated resources, such as VPCs and SLB instances. To query node pools or nodes, call the DescribeClusterNodePools or DescribeClusterNodes operation.
//
// @param request - DescribeClusterResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterResourcesResponse
func (client *Client) DescribeClusterResourcesWithContext(ctx context.Context, ClusterId *string, request *DescribeClusterResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WithAddonResources) {
		query["with_addon_resources"] = request.WithAddonResources
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterResources"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DescribeClusterResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tasks in a cluster.
//
// @param request - DescribeClusterTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterTasksResponse
func (client *Client) DescribeClusterTasksWithContext(ctx context.Context, clusterId *string, request *DescribeClusterTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["max_results"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["next_token"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["page_number"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterTasks"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the kubeconfig of a cluster.
//
// Description:
//
// - The default validity period of a certificate issued for a kubeconfig file is three years. Within 180 days before expiration, you can obtain a refreshed kubeconfig file with rotated certificates through the Container Service for Kubernetes console or by calling the API operation to retrieve the cluster kubeconfig. The new kubeconfig certificate remains valid for three years. The old kubeconfig credentials remain valid until the certificate expires. Obtain the rotated credentials promptly based on the kubeconfig expiration time displayed in the console or returned by the API operation.
//
// - Properly manage the kubeconfig credentials for cluster management and revoke them promptly when they are no longer needed to prevent security risks such as data leaks caused by kubeconfig exposure.
//
// @param request - DescribeClusterUserKubeconfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterUserKubeconfigResponse
func (client *Client) DescribeClusterUserKubeconfigWithContext(ctx context.Context, ClusterId *string, request *DescribeClusterUserKubeconfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterUserKubeconfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.TemporaryDurationMinutes) {
		query["TemporaryDurationMinutes"] = request.TemporaryDurationMinutes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterUserKubeconfig"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/k8s/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/user_config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterUserKubeconfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusterV2UserKubeconfig is deprecated
//
// Summary:
//
// Retrieves the kubeconfig for a cluster.
//
// @param request - DescribeClusterV2UserKubeconfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterV2UserKubeconfigResponse
func (client *Client) DescribeClusterV2UserKubeconfigWithContext(ctx context.Context, ClusterId *string, request *DescribeClusterV2UserKubeconfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterV2UserKubeconfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.TemporaryDurationMinutes) {
		query["TemporaryDurationMinutes"] = request.TemporaryDurationMinutes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterV2UserKubeconfig"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/k8s/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/user_config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterV2UserKubeconfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the DescribeClusterVuls operation to query the security vulnerability details of a cluster by cluster ID, including the vulnerability name, type, and severity level. Regularly scan clusters for security vulnerabilities to improve cluster security.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterVulsResponse
func (client *Client) DescribeClusterVulsWithContext(ctx context.Context, clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterVulsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterVuls"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/vuls"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterVulsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusters is deprecated
//
// Summary:
//
// Queries all clusters created in Container Service, including Swarm and Kubernetes clusters.
//
// @param request - DescribeClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClustersResponse
func (client *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterType) {
		query["clusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resource_group_id"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusters"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DescribeClustersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of all clusters in a specified region.
//
// @param request - DescribeClustersForRegionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClustersForRegionResponse
func (client *Client) DescribeClustersForRegionWithContext(ctx context.Context, regionId *string, request *DescribeClustersForRegionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClustersForRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterSpec) {
		query["cluster_spec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.ClusterType) {
		query["cluster_type"] = request.ClusterType
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["page_number"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.Profile) {
		query["profile"] = request.Profile
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClustersForRegion"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/regions/" + dara.PercentEncode(dara.StringValue(regionId)) + "/clusters"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClustersForRegionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of ACK clusters that match specified conditions, such as cluster type and cluster specifications, under the current account.
//
// @param request - DescribeClustersV1Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClustersV1Response
func (client *Client) DescribeClustersV1WithContext(ctx context.Context, request *DescribeClustersV1Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClustersV1Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterSpec) {
		query["cluster_spec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.ClusterType) {
		query["cluster_type"] = request.ClusterType
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["page_number"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.Profile) {
		query["profile"] = request.Profile
	}

	if !dara.IsNil(request.RegionId) {
		query["region_id"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClustersV1"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/clusters"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClustersV1Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of cluster operation events, such as cluster creation, cluster modification, node pool creation, node pool scale-out, add-on installation, and cluster upgrade. This operation allows you to query the details of specific events, including the event level, status, and occurrence time.
//
// @param request - DescribeEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsResponse
func (client *Client) DescribeEventsWithContext(ctx context.Context, request *DescribeEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.MaxResults) {
		query["max_results"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["next_token"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["page_number"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEvents"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of all events in a specified region.
//
// @param request - DescribeEventsForRegionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsForRegionResponse
func (client *Client) DescribeEventsForRegionWithContext(ctx context.Context, regionId *string, request *DescribeEventsForRegionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeEventsForRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.MaxResults) {
		query["max_results"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["next_token"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["page_number"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventsForRegion"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/regions/" + dara.PercentEncode(dara.StringValue(regionId)) + "/events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventsForRegionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeExternalAgent is deprecated
//
// Summary:
//
// Queries the agent configuration deployed in a registered cluster based on the cluster ID.
//
// Description:
//
// For more information about cluster registration, see [Register an external Kubernetes cluster](https://help.aliyun.com/document_detail/121053.html).
//
// @param request - DescribeExternalAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExternalAgentResponse
func (client *Client) DescribeExternalAgentWithContext(ctx context.Context, ClusterId *string, request *DescribeExternalAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeExternalAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentMode) {
		query["AgentMode"] = request.AgentMode
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExternalAgent"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/k8s/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/external/agent/deployment"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExternalAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of Kubernetes versions, including version information, release dates, expiration dates, compatible operating systems, and supported runtimes.
//
// @param request - DescribeKubernetesVersionMetadataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKubernetesVersionMetadataResponse
func (client *Client) DescribeKubernetesVersionMetadataWithContext(ctx context.Context, request *DescribeKubernetesVersionMetadataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeKubernetesVersionMetadataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.KubernetesVersion) {
		query["KubernetesVersion"] = request.KubernetesVersion
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Profile) {
		query["Profile"] = request.Profile
	}

	if !dara.IsNil(request.QueryUpgradableVersion) {
		query["QueryUpgradableVersion"] = request.QueryUpgradableVersion
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Runtime) {
		query["runtime"] = request.Runtime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKubernetesVersionMetadata"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/metadata/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DescribeKubernetesVersionMetadataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the DescribeNodePoolVuls API operation to query the security vulnerability details of a node pool by node pool ID, including the vulnerability name and level. Regularly scan node pools for security vulnerabilities to improve cluster security.
//
// @param request - DescribeNodePoolVulsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodePoolVulsResponse
func (client *Client) DescribeNodePoolVulsWithContext(ctx context.Context, clusterId *string, nodepoolId *string, request *DescribeNodePoolVulsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeNodePoolVulsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Necessity) {
		query["necessity"] = request.Necessity
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNodePoolVuls"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(nodepoolId)) + "/vuls"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNodePoolVulsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides a comprehensive set of built-in rule libraries for container security policies in ACK clusters, including Compliance, Infra, K8s-general, and PSP, to ensure secure container operations in production environments. You can call the DescribePolicies operation to query the list of policy governance rule libraries.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePoliciesResponse
func (client *Client) DescribePoliciesWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePoliciesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolicies"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/policies"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides a comprehensive built-in rule library for container security policies in ACK clusters, including Compliance, Infra, K8s-general, and PSP, to ensure the secure operation of containers in production environments. You can call the DescribePolicyDetails operation to query the details of specific policy governance rules, such as the rule template description, rule governance action, and rule governance level.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyDetailsResponse
func (client *Client) DescribePolicyDetailsWithContext(ctx context.Context, policyName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePolicyDetailsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolicyDetails"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/policies/" + dara.PercentEncode(dara.StringValue(policyName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolicyDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of policy governance for a specified ACK cluster. Container security policies for ACK clusters provide a comprehensive set of built-in rule libraries, including Compliance, Infra, K8s-general, and PSP, to ensure the secure operation of containers in production environments. You can use this operation to retrieve information such as the counts of enabled policies at different severity levels, policy governance audit logs, and interception and alerting statistics.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyGovernanceInClusterResponse
func (client *Client) DescribePolicyGovernanceInClusterWithContext(ctx context.Context, clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePolicyGovernanceInClusterResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolicyGovernanceInCluster"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/policygovernance"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolicyGovernanceInClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of specified policy rule instances in an ACK cluster, such as policy descriptions and governance levels. You can select a security policy type, configure related enforcement actions and scopes, and create and deploy a policy instance in the ACK cluster.
//
// @param request - DescribePolicyInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyInstancesResponse
func (client *Client) DescribePolicyInstancesWithContext(ctx context.Context, clusterId *string, request *DescribePolicyInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePolicyInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		query["instance_name"] = request.InstanceName
	}

	if !dara.IsNil(request.PolicyName) {
		query["policy_name"] = request.PolicyName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolicyInstances"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/policies"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DescribePolicyInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instance deployment status for different policy types in a cluster, including the number of enabled instances for each policy rule and the number of enabled policy types at different governance levels.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyInstancesStatusResponse
func (client *Client) DescribePolicyInstancesStatusWithContext(ctx context.Context, clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePolicyInstancesStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolicyInstancesStatus"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/policies/status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolicyInstancesStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of regions.
//
// @param request - DescribeRegionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterType) {
		query["clusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.Profile) {
		query["profile"] = request.Profile
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/regions"),
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
// Queries whether deletion protection is enabled for specified resources in a cluster. Supported resources include namespaces and services.
//
// @param request - DescribeResourcesDeleteProtectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourcesDeleteProtectionResponse
func (client *Client) DescribeResourcesDeleteProtectionWithContext(ctx context.Context, ClusterId *string, ResourceType *string, request *DescribeResourcesDeleteProtectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeResourcesDeleteProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Resources) {
		query["resources"] = request.Resources
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourcesDeleteProtection"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/resources/" + dara.PercentEncode(dara.StringValue(ResourceType)) + "/protection"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DescribeResourcesDeleteProtectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// As an administrator responsible for cluster permission management, you can use an Alibaba Cloud account to issue KubeConfig credentials that contain identity information for a specified Resource Access Management (RAM) user or RAM role within the account to connect to an ACK cluster. You can invoke the DescribeSubaccountK8sClusterUserConfig operation to issue or obtain the KubeConfig for any RAM user or role within the account.
//
// Description:
//
// This API operation can be called only by an Alibaba Cloud account.
//
// @param request - DescribeSubaccountK8sClusterUserConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubaccountK8sClusterUserConfigResponse
func (client *Client) DescribeSubaccountK8sClusterUserConfigWithContext(ctx context.Context, ClusterId *string, Uid *string, request *DescribeSubaccountK8sClusterUserConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeSubaccountK8sClusterUserConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.TemporaryDurationMinutes) {
		query["TemporaryDurationMinutes"] = request.TemporaryDurationMinutes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSubaccountK8sClusterUserConfig"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/k8s/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/users/" + dara.PercentEncode(dara.StringValue(Uid)) + "/user_config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSubaccountK8sClusterUserConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a cluster task, such as the task type, running status, and running stage.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskInfoResponse
func (client *Client) DescribeTaskInfoWithContext(ctx context.Context, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTaskInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTaskInfo"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified orchestration template, including access permissions, YAML content, and tags. An orchestration template defines and describes a set of Kubernetes cluster resources in a declarative manner to specify how applications should run or be configured.
//
// @param request - DescribeTemplateAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplateAttributeResponse
func (client *Client) DescribeTemplateAttributeWithContext(ctx context.Context, TemplateId *string, request *DescribeTemplateAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTemplateAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateType) {
		query["template_type"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTemplateAttribute"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/templates/" + dara.PercentEncode(dara.StringValue(TemplateId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DescribeTemplateAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of orchestration templates and retrieves their details, including access permissions, YAML content, and tags. An orchestration template defines and describes a set of Kubernetes cluster resources in a declarative manner to specify how an application should run or be configured.
//
// @param request - DescribeTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplatesResponse
func (client *Client) DescribeTemplatesWithContext(ctx context.Context, request *DescribeTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		query["page_num"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.TemplateType) {
		query["template_type"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTemplates"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/templates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Queries triggers that meet the specified conditions.
//
// @param request - DescribeTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTriggerResponse
func (client *Client) DescribeTriggerWithContext(ctx context.Context, clusterId *string, request *DescribeTriggerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTriggerResponse, _err error) {
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

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Action) {
		query["action"] = request.Action
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrigger"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/triggers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DescribeTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the RBAC access permissions to resources in namespaces granted to the current Resource Access Management (RAM) user or role in a specified ACK cluster. In ACK clusters, Kubernetes namespaces ensure that permissions and resources for cluster users are logically isolated. Users granted RBAC permissions only for a specific namespace cannot access resources in other namespaces of the cluster. You can invoke the DescribeUserClusterNamespaces operation to view the namespaces for which the current RAM user or role has been granted access permissions.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserClusterNamespacesResponse
func (client *Client) DescribeUserClusterNamespacesWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeUserClusterNamespacesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserClusterNamespaces"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/k8s/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/namespaces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DescribeUserClusterNamespacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// In ACK clusters, you can create and assign different access permissions to different Resource Access Management (RAM) users or roles to ensure secure access control and resource isolation. You can invoke the DescribeUserPermission operation to query the details of cluster permissions granted to a Resource Access Management (RAM) user or role, including accessible resources, permission scope, preset role types, and permission sources.
//
// Description:
//
// *Precautions**:
//
// - If the account that invokes this API operation is a Resource Access Management (RAM) user or RAM role, the API operation returns only the permissions for clusters on which the calling account has RBAC administrator permissions. To list permissions for all clusters, the calling account must have RBAC administrator permissions on all clusters.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserPermissionResponse
func (client *Client) DescribeUserPermissionWithContext(ctx context.Context, uid *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeUserPermissionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserPermission"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/permissions/users/" + dara.PercentEncode(dara.StringValue(uid))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DescribeUserPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quotas for ACK clusters, node pools, and nodes. To increase a quota, apply in Quota Center.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserQuotaResponse
func (client *Client) DescribeUserQuotaWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeUserQuotaResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserQuota"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/quota"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Fixes security vulnerability on nodes in a specified node pool of a cluster to improve cluster security. Alibaba Cloud security products regularly scan Elastic Compute Service (ECS) nodes for security vulnerability and provide corresponding fix suggestions and methods. Fixing certain Common Vulnerabilities and Exposures (CVEs) may require node restarts. Ensure that the cluster has sufficient nodes for drain operations. You can invoke the FixNodePoolVuls operation to fix security vulnerability on nodes in a specified node pool and improve the security of cluster nodes.
//
// Description:
//
// - CVE compatibility is guaranteed by Security Center. Ensure that you have activated the Ultimate edition of Security Center or [purchased vulnerability fixes (pay-as-you-go)](https://help.aliyun.com/document_detail/42308.html).
//
// - Fixing certain CVEs requires node restarts. Container Service performs drain operations on a node before it restarts. Ensure that the cluster has sufficient idle resources for drain operations, such as scaling out the node pool in advance.
//
// - Pay attention to the compatibility between your applications and CVEs. CVE fixes are performed in batches. During the fix process, you can pause or cancel the task. After you pause or cancel the task, the batches to which fix tasks have been delivered continue to run until completion. The delivery of fix tasks to undelivered batches is paused or canceled.
//
// @param request - FixNodePoolVulsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FixNodePoolVulsResponse
func (client *Client) FixNodePoolVulsWithContext(ctx context.Context, clusterId *string, nodepoolId *string, request *FixNodePoolVulsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FixNodePoolVulsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRestart) {
		body["auto_restart"] = request.AutoRestart
	}

	if !dara.IsNil(request.Nodes) {
		body["nodes"] = request.Nodes
	}

	if !dara.IsNil(request.RolloutPolicy) {
		body["rollout_policy"] = request.RolloutPolicy
	}

	if !dara.IsNil(request.Vuls) {
		body["vuls"] = request.Vuls
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FixNodePoolVuls"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(nodepoolId)) + "/vuls/fix"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FixNodePoolVulsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the GetClusterAddonInstance operation to query the details of a specified add-on in a cluster, including the version, parameter settings, and log feature enabling status of the add-on instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterAddonInstanceResponse
func (client *Client) GetClusterAddonInstanceWithContext(ctx context.Context, clusterId *string, instanceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterAddonInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClusterAddonInstance"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/addon_instances/" + dara.PercentEncode(dara.StringValue(instanceName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterAddonInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the API server audit feature is enabled for a cluster and the Simple Log Service (SLS) project associated with the API server audit logs.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterAuditProjectResponse
func (client *Client) GetClusterAuditProjectWithContext(ctx context.Context, clusterid *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterAuditProjectResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClusterAuditProject"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterid)) + "/audit"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterAuditProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific check task for a Kubernetes cluster. The intelligent O&M platform for containers provides comprehensive cluster check capabilities, including cluster upgrade, cluster migration, component installation, component upgrade, and node pool checks. You can retrieve information such as the check status, specific check items, and the creation and completion time of the check task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterCheckResponse
func (client *Client) GetClusterCheckWithContext(ctx context.Context, clusterId *string, checkId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterCheckResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClusterCheck"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/checks/" + dara.PercentEncode(dara.StringValue(checkId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the check items for cluster diagnostics.
//
// @param request - GetClusterDiagnosisCheckItemsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterDiagnosisCheckItemsResponse
func (client *Client) GetClusterDiagnosisCheckItemsWithContext(ctx context.Context, clusterId *string, diagnosisId *string, request *GetClusterDiagnosisCheckItemsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterDiagnosisCheckItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClusterDiagnosisCheckItems"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/diagnosis/" + dara.PercentEncode(dara.StringValue(diagnosisId)) + "/check_items"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterDiagnosisCheckItemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the cluster diagnosis results.
//
// @param request - GetClusterDiagnosisResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterDiagnosisResultResponse
func (client *Client) GetClusterDiagnosisResultWithContext(ctx context.Context, clusterId *string, diagnosisId *string, request *GetClusterDiagnosisResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterDiagnosisResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClusterDiagnosisResult"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/diagnosis/" + dara.PercentEncode(dara.StringValue(diagnosisId)) + "/result"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterDiagnosisResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the cluster inspection configuration.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterInspectConfigResponse
func (client *Client) GetClusterInspectConfigWithContext(ctx context.Context, clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterInspectConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClusterInspectConfig"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/inspectConfig"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterInspectConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a cluster inspection report.
//
// @param request - GetClusterInspectReportDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterInspectReportDetailResponse
func (client *Client) GetClusterInspectReportDetailWithContext(ctx context.Context, clusterId *string, reportId *string, request *GetClusterInspectReportDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterInspectReportDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["category"] = request.Category
	}

	if !dara.IsNil(request.EnableFilter) {
		query["enableFilter"] = request.EnableFilter
	}

	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

	if !dara.IsNil(request.Level) {
		query["level"] = request.Level
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TargetType) {
		query["targetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClusterInspectReportDetail"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/inspectReports/" + dara.PercentEncode(dara.StringValue(reportId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterInspectReportDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetKubernetesTrigger is deprecated
//
// Summary:
//
// Queries the triggers of an application by name.
//
// @param request - GetKubernetesTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKubernetesTriggerResponse
func (client *Client) GetKubernetesTriggerWithContext(ctx context.Context, ClusterId *string, request *GetKubernetesTriggerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKubernetesTriggerResponse, _err error) {
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

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Action) {
		query["action"] = request.Action
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKubernetesTrigger"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/triggers/" + dara.PercentEncode(dara.StringValue(ClusterId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &GetKubernetesTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetUpgradeStatus is deprecated
//
// Summary:
//
// Queries the upgrade status of a cluster based on the cluster ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUpgradeStatusResponse
func (client *Client) GetUpgradeStatusWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUpgradeStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUpgradeStatus"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/upgrade/status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUpgradeStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// By default, Resource Access Management (RAM) users or roles that are not cluster creators and have not been granted access permissions on all cluster dimensions do not have any RBAC permissions in a cluster. You can invoke the GrantPermissions operation to update the RBAC access permissions of a RAM user or role, including accessible resources, permission scope, and preset role types, to better manage cluster management security access control.
//
// Description:
//
// - If the account that invokes this API operation is a Resource Access Management (RAM) user, make sure that the account has been granted permissions to modify the cluster authorization information of other RAM users or RAM roles. Otherwise, the API returns the `StatusForbidden` or `ForbiddenGrantPermissions` fault code. For more information, see [Grant RBAC authorization to other RAM users](https://help.aliyun.com/document_detail/119035.html).
//
// - The operation of updating all RBAC permissions of a RAM user or RAM role overwrites the existing cluster permissions of the target RAM user or RAM role. The request must include all permission configurations that you want to grant to the target RAM user or RAM role.
//
// @param request - GrantPermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantPermissionsResponse
func (client *Client) GrantPermissionsWithContext(ctx context.Context, uid *string, request *GrantPermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantPermissions"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/permissions/users/" + dara.PercentEncode(dara.StringValue(uid))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &GrantPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// To enhance Kubernetes capabilities, ACK clusters support various components, such as managed core components, and application, logging and monitoring, networking, storage, and security components. You can call the InstallClusterAddons operation to install components by specifying the component name and version.
//
// @param request - InstallClusterAddonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallClusterAddonsResponse
func (client *Client) InstallClusterAddonsWithContext(ctx context.Context, ClusterId *string, request *InstallClusterAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallClusterAddonsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallClusterAddons"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/components/install"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallClusterAddonsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs node components on specified nodes with custom configurations.
//
// @param request - InstallNodePoolComponentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallNodePoolComponentsResponse
func (client *Client) InstallNodePoolComponentsWithContext(ctx context.Context, clusterId *string, nodePoolId *string, request *InstallNodePoolComponentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallNodePoolComponentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Components) {
		body["components"] = request.Components
	}

	if !dara.IsNil(request.NodeNames) {
		body["nodeNames"] = request.NodeNames
	}

	if !dara.IsNil(request.RollingPolicy) {
		body["rollingPolicy"] = request.RollingPolicy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallNodePoolComponents"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(nodePoolId)) + "/components"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallNodePoolComponentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of available components and retrieves detailed component information based on parameters such as the region, cluster type, cluster profile, and cluster version. The detailed information includes whether the component is managed, the schema of supported custom parameters, and compatible operating system architectures.
//
// @param request - ListAddonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonsResponse
func (client *Client) ListAddonsWithContext(ctx context.Context, request *ListAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAddonsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterSpec) {
		query["cluster_spec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.ClusterType) {
		query["cluster_type"] = request.ClusterType
	}

	if !dara.IsNil(request.ClusterVersion) {
		query["cluster_version"] = request.ClusterVersion
	}

	if !dara.IsNil(request.Profile) {
		query["profile"] = request.Profile
	}

	if !dara.IsNil(request.RegionId) {
		query["region_id"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAddons"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/addons"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAddonsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of self-healing rules.
//
// Description:
//
//	Notice: This API is not yet available. Stay tuned.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoRepairPoliciesResponse
func (client *Client) ListAutoRepairPoliciesWithContext(ctx context.Context, clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAutoRepairPoliciesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAutoRepairPolicies"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/auto_repair_policies"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAutoRepairPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of resources contained in an installed cluster component instance, including Kubernetes cluster resources and Helm release instances.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterAddonInstanceResourcesResponse
func (client *Client) ListClusterAddonInstanceResourcesWithContext(ctx context.Context, clusterId *string, instanceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterAddonInstanceResourcesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterAddonInstanceResources"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/addon_instances/" + dara.PercentEncode(dara.StringValue(instanceName)) + "/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterAddonInstanceResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of installed add-on instances in a specified cluster and retrieves related information about the add-on instances, such as the add-on version and status.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterAddonInstancesResponse
func (client *Client) ListClusterAddonInstancesWithContext(ctx context.Context, clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterAddonInstancesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterAddonInstances"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/addon_instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterAddonInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of cluster checks and related information, such as the check type, status, creation time, and completion time, based on the cluster ID. The container intelligent O&M platform provides comprehensive Kubernetes cluster check capabilities, including cluster upgrade checks, cluster migration checks, component installation checks, component upgrade checks, and node pool checks.
//
// @param request - ListClusterChecksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterChecksResponse
func (client *Client) ListClusterChecksWithContext(ctx context.Context, clusterId *string, request *ListClusterChecksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterChecksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Target) {
		query["target"] = request.Target
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterChecks"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/checks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterChecksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of cluster inspection reports.
//
// @param request - ListClusterInspectReportsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterInspectReportsResponse
func (client *Client) ListClusterInspectReportsWithContext(ctx context.Context, clusterId *string, request *ListClusterInspectReportsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterInspectReportsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterInspectReports"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/inspectReports"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterInspectReportsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list and status of user KubeConfigs issued for the current cluster.
//
// Description:
//
// > - You must have the ram:ListUsers and ram:ListRoles permissions.
//
// > - You must have the permissions to manage Container Service for Kubernetes (CS) (AliyunCSFullAccess).
//
// @param request - ListClusterKubeconfigStatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterKubeconfigStatesResponse
func (client *Client) ListClusterKubeconfigStatesWithContext(ctx context.Context, ClusterId *string, request *ListClusterKubeconfigStatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterKubeconfigStatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudServiceKubeConfig) {
		query["cloudServiceKubeConfig"] = request.CloudServiceKubeConfig
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterKubeconfigStates"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/kubeconfig/states"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterKubeconfigStatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of node components on nodes.
//
// @param tmpReq - ListNodePoolComponentInstanceNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodePoolComponentInstanceNodesResponse
func (client *Client) ListNodePoolComponentInstanceNodesWithContext(ctx context.Context, clusterId *string, nodepoolId *string, name *string, tmpReq *ListNodePoolComponentInstanceNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNodePoolComponentInstanceNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListNodePoolComponentInstanceNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeNames) {
		request.NodeNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeNames, dara.String("node_names"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigRevision) {
		query["config_revision"] = request.ConfigRevision
	}

	if !dara.IsNil(request.MaxResults) {
		query["max_results"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["next_token"] = request.NextToken
	}

	if !dara.IsNil(request.NodeNamesShrink) {
		query["node_names"] = request.NodeNamesShrink
	}

	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodePoolComponentInstanceNodes"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(nodepoolId)) + "/component_instances/" + dara.PercentEncode(dara.StringValue(name)) + "/nodes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodePoolComponentInstanceNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of node component instances.
//
// @param request - ListNodePoolComponentInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodePoolComponentInstancesResponse
func (client *Client) ListNodePoolComponentInstancesWithContext(ctx context.Context, clusterId *string, nodepoolId *string, request *ListNodePoolComponentInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNodePoolComponentInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["max_results"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["next_token"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodePoolComponentInstances"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(nodepoolId)) + "/component_instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodePoolComponentInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the node components that can be installed in a node pool.
//
// @param tmpReq - ListNodePoolComponentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodePoolComponentsResponse
func (client *Client) ListNodePoolComponentsWithContext(ctx context.Context, clusterId *string, tmpReq *ListNodePoolComponentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNodePoolComponentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListNodePoolComponentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceTypes) {
		request.InstanceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceTypes, dara.String("instance_types"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["image_id"] = request.ImageId
	}

	if !dara.IsNil(request.ImageType) {
		query["image_type"] = request.ImageType
	}

	if !dara.IsNil(request.InstanceTypesShrink) {
		query["instance_types"] = request.InstanceTypesShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["max_results"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["next_token"] = request.NextToken
	}

	if !dara.IsNil(request.NodepoolId) {
		query["nodepool_id"] = request.NodepoolId
	}

	if !dara.IsNil(request.NodepoolType) {
		query["nodepool_type"] = request.NodepoolType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodePoolComponents"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/nodepool_components"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodePoolComponentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of automated O&M execution plans.
//
// @param request - ListOperationPlansRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationPlansResponse
func (client *Client) ListOperationPlansWithContext(ctx context.Context, request *ListOperationPlansRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOperationPlansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationPlans"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/operation/plans"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationPlansResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the latest 100 automated O&M execution plans in a specified region. When features such as Auto Mode for intelligent cluster management, automatic cluster upgrades, and automated O&M for node pools are enabled, you can use this operation to query system-generated O&M plans and their execution status, such as cluster upgrades and node pool CVE fixes.
//
// @param request - ListOperationPlansForRegionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationPlansForRegionResponse
func (client *Client) ListOperationPlansForRegionWithContext(ctx context.Context, regionId *string, request *ListOperationPlansForRegionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOperationPlansForRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.State) {
		query["state"] = request.State
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationPlansForRegion"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/regions/" + dara.PercentEncode(dara.StringValue(regionId)) + "/operation/plans"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationPlansForRegionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag list of clusters. You can add tag key-value pairs to clusters so that developers and O&M engineers can classify and manage clusters more flexibly. This better supports requirements such as monitoring, cost analysis, and tenant isolation. You can call the ListTagResources operation to obtain the list of resource tags and query tag details, such as key-value pairs and associated clusters.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceIds) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, dara.String("resource_ids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["next_token"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["region_id"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIdsShrink) {
		query["resource_ids"] = request.ResourceIdsShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["resource_type"] = request.ResourceType
	}

	if !dara.IsNil(request.TagsShrink) {
		query["tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tags"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Queries the KubeConfig status list of all clusters for the current user.
//
// Description:
//
// > To call this operation, you must have the permissions to manage Container Service (CS) (AliyunCSFullAccess).
//
// @param request - ListUserKubeConfigStatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserKubeConfigStatesResponse
func (client *Client) ListUserKubeConfigStatesWithContext(ctx context.Context, Uid *string, request *ListUserKubeConfigStatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserKubeConfigStatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["page_number"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserKubeConfigStates"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/users/" + dara.PercentEncode(dara.StringValue(Uid)) + "/kubeconfig/states"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserKubeConfigStatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates a cluster.
//
// Description:
//
// After you migrate an ACK basic managed cluster to an ACK Pro managed cluster, a [cluster management fee](https://help.aliyun.com/document_detail/462278.html) charged by ACK is added. The billing of other cloud resources remains unchanged.
//
// @param request - MigrateClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateClusterResponse
func (client *Client) MigrateClusterWithContext(ctx context.Context, clusterId *string, request *MigrateClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MigrateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OssBucketEndpoint) {
		body["oss_bucket_endpoint"] = request.OssBucketEndpoint
	}

	if !dara.IsNil(request.OssBucketName) {
		body["oss_bucket_name"] = request.OssBucketName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateCluster"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/migrate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a self-healing rule.
//
// Description:
//
//	Notice: This API is not yet available. Stay tuned.
//
// @param request - ModifyAutoRepairPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAutoRepairPolicyResponse
func (client *Client) ModifyAutoRepairPolicyWithContext(ctx context.Context, clusterId *string, policyId *string, request *ModifyAutoRepairPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyAutoRepairPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Rules) {
		body["rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAutoRepairPolicy"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/auto_repair_policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAutoRepairPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of an ACK cluster.
//
// Description:
//
// <notice>Starting from July 4, 2026, the request parameters instance_deletion_protection, ingress_loadbalancer_id, and access_control_list will no longer take effect. For more information about the changes, see [Announcement on changes to OpenAPI request and response parameters and OpenAPI deprecation for ACK cluster management](https://help.aliyun.com/document_detail/2932733.html).</notice>
//
// @param request - ModifyClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterResponse
func (client *Client) ModifyClusterWithContext(ctx context.Context, ClusterId *string, request *ModifyClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessControlList) {
		body["access_control_list"] = request.AccessControlList
	}

	if !dara.IsNil(request.ApiServerCustomCertSans) {
		body["api_server_custom_cert_sans"] = request.ApiServerCustomCertSans
	}

	if !dara.IsNil(request.ApiServerEip) {
		body["api_server_eip"] = request.ApiServerEip
	}

	if !dara.IsNil(request.ApiServerEipId) {
		body["api_server_eip_id"] = request.ApiServerEipId
	}

	if !dara.IsNil(request.ClientToken) {
		body["client_token"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterName) {
		body["cluster_name"] = request.ClusterName
	}

	if !dara.IsNil(request.ClusterSpec) {
		body["cluster_spec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.ControlPlaneConfig) {
		body["control_plane_config"] = request.ControlPlaneConfig
	}

	if !dara.IsNil(request.ControlPlaneEndpointsConfig) {
		body["control_plane_endpoints_config"] = request.ControlPlaneEndpointsConfig
	}

	if !dara.IsNil(request.DeletionProtection) {
		body["deletion_protection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.EnableRrsa) {
		body["enable_rrsa"] = request.EnableRrsa
	}

	if !dara.IsNil(request.IngressDomainRebinding) {
		body["ingress_domain_rebinding"] = request.IngressDomainRebinding
	}

	if !dara.IsNil(request.IngressLoadbalancerId) {
		body["ingress_loadbalancer_id"] = request.IngressLoadbalancerId
	}

	if !dara.IsNil(request.InstanceDeletionProtection) {
		body["instance_deletion_protection"] = request.InstanceDeletionProtection
	}

	if !dara.IsNil(request.MaintenanceWindow) {
		body["maintenance_window"] = request.MaintenanceWindow
	}

	if !dara.IsNil(request.OperationPolicy) {
		body["operation_policy"] = request.OperationPolicy
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resource_group_id"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["security_group_id"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SystemEventsLogging) {
		body["system_events_logging"] = request.SystemEventsLogging
	}

	if !dara.IsNil(request.Timezone) {
		body["timezone"] = request.Timezone
	}

	if !dara.IsNil(request.VswitchIds) {
		body["vswitch_ids"] = request.VswitchIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCluster"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of an installed cluster add-on instance. Modifying the configuration may affect your services. Evaluate the impact, perform the operation during off-peak hours, and back up your data in advance.
//
// Description:
//
// You can call this operation to modify the configuration of add-ons in common clusters and the control plane parameter settings of ACK Pro clusters:
//
// - You can call the DescribeClusterAddonMetadata API operation to query the list of configurable parameters for common add-ons. For details, see [Query cluster add-on version metadata](https://help.aliyun.com/document_detail/2667944.html).
//
// - For the list of configurable parameters for the control plane components of ACK Pro clusters, see [Customize the control plane parameter settings of an ACK Pro cluster](https://help.aliyun.com/document_detail/199588.html).
//
// Modifying the configuration may cause add-on redeployment and restart. Evaluate the impact before you proceed.
//
// @param request - ModifyClusterAddonRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterAddonResponse
func (client *Client) ModifyClusterAddonWithContext(ctx context.Context, clusterId *string, componentId *string, request *ModifyClusterAddonRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyClusterAddonResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyClusterAddon"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/components/" + dara.PercentEncode(dara.StringValue(componentId)) + "/config"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &ModifyClusterAddonResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a node pool based on the node pool ID by calling the ModifyClusterNodePool operation.
//
// Description:
//
// <notice>Starting July 4, 2026, the request parameters instance_deletion_protection, ingress_loadbalancer_id, and access_control_list will no longer take effect. For more information about the changes, see [Announcement on changes to input and output parameters of ACK cluster management OpenAPI operations and deprecation of OpenAPI operations](https://help.aliyun.com/document_detail/2932733.html).</notice>
//
// @param request - ModifyClusterNodePoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterNodePoolResponse
func (client *Client) ModifyClusterNodePoolWithContext(ctx context.Context, ClusterId *string, NodepoolId *string, request *ModifyClusterNodePoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyClusterNodePoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoScaling) {
		body["auto_scaling"] = request.AutoScaling
	}

	if !dara.IsNil(request.Concurrency) {
		body["concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.EfloNodeGroup) {
		body["eflo_node_group"] = request.EfloNodeGroup
	}

	if !dara.IsNil(request.KubernetesConfig) {
		body["kubernetes_config"] = request.KubernetesConfig
	}

	if !dara.IsNil(request.Management) {
		body["management"] = request.Management
	}

	if !dara.IsNil(request.NodepoolInfo) {
		body["nodepool_info"] = request.NodepoolInfo
	}

	if !dara.IsNil(request.ScalingGroup) {
		body["scaling_group"] = request.ScalingGroup
	}

	if !dara.IsNil(request.TeeConfig) {
		body["tee_config"] = request.TeeConfig
	}

	if !dara.IsNil(request.UpdateNodes) {
		body["update_nodes"] = request.UpdateNodes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyClusterNodePool"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(NodepoolId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyClusterNodePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the tags of a cluster. You can add tag key-value pairs to a cluster to help developers and O&M engineers classify and manage clusters more flexibly, and better support requirements such as monitoring, cost analysis, and tenant isolation. You can call the ModifyClusterTags operation to modify cluster tags.
//
// Description:
//
// - This operation performs a full update. When you call this operation, specify all target tags in the request parameters to avoid losing existing tags. To perform an incremental update, use the [TagResources](https://help.aliyun.com/document_detail/2667969.html) operation to add tags to the cluster.
//
// @param request - ModifyClusterTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterTagsResponse
func (client *Client) ModifyClusterTagsWithContext(ctx context.Context, ClusterId *string, request *ModifyClusterTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyClusterTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyClusterTags"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/tags"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &ModifyClusterTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the node configurations in a cluster node pool, such as kubelet configurations and rolling update configurations. Modifying node configurations applies changes in batches and restarts kubelet, which may affect node operations and workload execution. Perform this operation during off-peak hours.
//
// Description:
//
// > ACK supports modifying the kubelet configurations of nodes in a node pool. After the modification is complete, the configurations automatically take effect on existing nodes in the node pool, and new nodes added to the node pool also use the new configurations.
//
// @param request - ModifyNodePoolNodeConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodePoolNodeConfigResponse
func (client *Client) ModifyNodePoolNodeConfigWithContext(ctx context.Context, ClusterId *string, NodepoolId *string, request *ModifyNodePoolNodeConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyNodePoolNodeConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContainerdConfig) {
		body["containerd_config"] = request.ContainerdConfig
	}

	if !dara.IsNil(request.KubeletConfig) {
		body["kubelet_config"] = request.KubeletConfig
	}

	if !dara.IsNil(request.NodeNames) {
		body["node_names"] = request.NodeNames
	}

	if !dara.IsNil(request.OsConfig) {
		body["os_config"] = request.OsConfig
	}

	if !dara.IsNil(request.RollingPolicy) {
		body["rolling_policy"] = request.RollingPolicy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNodePoolNodeConfig"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(NodepoolId)) + "/node_config"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNodePoolNodeConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a policy rule instance in a specified cluster. You can modify the governance action, such as alert or block, and the namespace scope of the policy instance.
//
// @param request - ModifyPolicyInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyInstanceResponse
func (client *Client) ModifyPolicyInstanceWithContext(ctx context.Context, clusterId *string, policyName *string, request *ModifyPolicyInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyPolicyInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["action"] = request.Action
	}

	if !dara.IsNil(request.InstanceName) {
		body["instance_name"] = request.InstanceName
	}

	if !dara.IsNil(request.Namespaces) {
		body["namespaces"] = request.Namespaces
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPolicyInstance"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/policies/" + dara.PercentEncode(dara.StringValue(policyName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPolicyInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates Container Service for Kubernetes (ACK) for first-time users.
//
// Description:
//
// - An Alibaba Cloud account can activate Container Service for Kubernetes (ACK).
//
// - A Resource Access Management (RAM) user to which the AdministratorAccess policy is attached can activate the service.
//
// @param request - OpenAckServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenAckServiceResponse
func (client *Client) OpenAckServiceWithContext(ctx context.Context, request *OpenAckServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenAckServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenAckService"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/service/open"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenAckServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI PauseClusterUpgrade is deprecated
//
// Summary:
//
// Pauses a cluster upgrade.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseClusterUpgradeResponse
func (client *Client) PauseClusterUpgradeWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PauseClusterUpgradeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseClusterUpgrade"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/upgrade/pause"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &PauseClusterUpgradeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI PauseComponentUpgrade is deprecated
//
// Summary:
//
// Pauses a component upgrade.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseComponentUpgradeResponse
func (client *Client) PauseComponentUpgradeWithContext(ctx context.Context, clusterid *string, componentid *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PauseComponentUpgradeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseComponentUpgrade"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterid)) + "/components/" + dara.PercentEncode(dara.StringValue(componentid)) + "/pause"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &PauseComponentUpgradeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses a running cluster task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseTaskResponse
func (client *Client) PauseTaskWithContext(ctx context.Context, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PauseTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseTask"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks/" + dara.PercentEncode(dara.StringValue(taskId)) + "/pause"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &PauseTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI RemoveClusterNodes is deprecated
//
// Summary:
//
// 移除集群节点
//
// @param request - RemoveClusterNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveClusterNodesResponse
func (client *Client) RemoveClusterNodesWithContext(ctx context.Context, ClusterId *string, request *RemoveClusterNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveClusterNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DrainNode) {
		body["drain_node"] = request.DrainNode
	}

	if !dara.IsNil(request.Nodes) {
		body["nodes"] = request.Nodes
	}

	if !dara.IsNil(request.ReleaseNode) {
		body["release_node"] = request.ReleaseNode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveClusterNodes"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodes/remove"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &RemoveClusterNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes nodes from a node pool in a cluster and adjusts the expected number of nodes when the nodes are no longer needed. When removing nodes, you can specify whether to release the associated ECS instances and whether to automatically drain the nodes. Node removal involves pod migration, which may affect your workloads. Perform this operation during off-peak hours and back up your data in advance.
//
// Description:
//
// - Node removal involves pod migration, which may affect your workloads. Perform this operation during off-peak hours.
//
// - Unexpected risks may occur during the operation. Back up your data in advance.
//
// - During the operation, the system sets the nodes to the unschedulable state.
//
// - Only worker nodes are removed. Master nodes are not removed.
//
// - Even if you choose to release nodes (by setting release_node to true), subscription nodes are not released. You must release them manually in the [ECS console](https://ecs.console.aliyun.com/) after removing them from the node pool.
//
// @param tmpReq - RemoveNodePoolNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveNodePoolNodesResponse
func (client *Client) RemoveNodePoolNodesWithContext(ctx context.Context, ClusterId *string, NodepoolId *string, tmpReq *RemoveNodePoolNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveNodePoolNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveNodePoolNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("instance_ids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Nodes) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, dara.String("nodes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Concurrency) {
		query["concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.DrainNode) {
		query["drain_node"] = request.DrainNode
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["instance_ids"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.NodesShrink) {
		query["nodes"] = request.NodesShrink
	}

	if !dara.IsNil(request.ReleaseNode) {
		query["release_node"] = request.ReleaseNode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveNodePoolNodes"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(NodepoolId)) + "/nodes"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveNodePoolNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Repairs a node pool in a cluster.
//
// @param request - RepairClusterNodePoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RepairClusterNodePoolResponse
func (client *Client) RepairClusterNodePoolWithContext(ctx context.Context, clusterId *string, nodepoolId *string, request *RepairClusterNodePoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RepairClusterNodePoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRestart) {
		body["auto_restart"] = request.AutoRestart
	}

	if !dara.IsNil(request.Nodes) {
		body["nodes"] = request.Nodes
	}

	if !dara.IsNil(request.Operations) {
		body["operations"] = request.Operations
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RepairClusterNodePool"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(nodepoolId)) + "/repair"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RepairClusterNodePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ResumeComponentUpgrade is deprecated
//
// Summary:
//
// Resumes a paused component upgrade task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeComponentUpgradeResponse
func (client *Client) ResumeComponentUpgradeWithContext(ctx context.Context, clusterid *string, componentid *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeComponentUpgradeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeComponentUpgrade"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterid)) + "/components/" + dara.PercentEncode(dara.StringValue(componentid)) + "/resume"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &ResumeComponentUpgradeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes a paused cluster task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeTaskResponse
func (client *Client) ResumeTaskWithContext(ctx context.Context, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeTask"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks/" + dara.PercentEncode(dara.StringValue(taskId)) + "/resume"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &ResumeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ResumeUpgradeCluster is deprecated
//
// Summary:
//
// Resumes the upgrade of a cluster that is in the upgrade-paused state based on the cluster ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeUpgradeClusterResponse
func (client *Client) ResumeUpgradeClusterWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeUpgradeClusterResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeUpgradeCluster"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/upgrade/resume"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &ResumeUpgradeClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If you want to revoke the cluster KubeConfig credential owned by the currently logged-on Alibaba Cloud account or Resource Access Management (RAM) user, you can call the RevokeK8sClusterKubeConfig operation to revoke it. After the revocation succeeds, the cluster generates a new KubeConfig, and the original KubeConfig becomes invalid.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeK8sClusterKubeConfigResponse
func (client *Client) RevokeK8sClusterKubeConfigWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeK8sClusterKubeConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeK8sClusterKubeConfig"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/k8s/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/certs"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeK8sClusterKubeConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a check for a Kubernetes cluster. The Container Intelligent O&M Platform provides comprehensive check capabilities for Kubernetes clusters, including cluster upgrade, cluster migration, component installation, component upgrade, and node pool checks. The platform automatically triggers a check before an upgrade, migration, or installation. You can perform the change operation only after the check passes. You can also manually call the RunClusterCheck operation to perform a cluster check. Regularly check and maintain your clusters to prevent security risks.
//
// @param request - RunClusterCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunClusterCheckResponse
func (client *Client) RunClusterCheckWithContext(ctx context.Context, clusterId *string, request *RunClusterCheckRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunClusterCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	if !dara.IsNil(request.Target) {
		body["target"] = request.Target
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunClusterCheck"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/checks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunClusterCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a cluster inspection and creates an inspection report.
//
// @param request - RunClusterInspectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunClusterInspectResponse
func (client *Client) RunClusterInspectWithContext(ctx context.Context, clusterId *string, request *RunClusterInspectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunClusterInspectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunClusterInspect"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/inspectReports"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunClusterInspectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes O&M operations on a node.
//
// @param request - RunNodeOperationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunNodeOperationResponse
func (client *Client) RunNodeOperationWithContext(ctx context.Context, clusterId *string, nodepoolId *string, nodeName *string, request *RunNodeOperationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunNodeOperationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OperationAction) {
		body["operationAction"] = request.OperationAction
	}

	if !dara.IsNil(request.OperationArgs) {
		body["operationArgs"] = request.OperationArgs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunNodeOperation"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(nodepoolId)) + "/nodes/" + dara.PercentEncode(dara.StringValue(nodeName)) + "/operation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunNodeOperationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales out the nodes in a node pool to ensure sufficient node capacity for workloads.
//
// @param request - ScaleClusterNodePoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleClusterNodePoolResponse
func (client *Client) ScaleClusterNodePoolWithContext(ctx context.Context, ClusterId *string, NodepoolId *string, request *ScaleClusterNodePoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleClusterNodePoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Count) {
		body["count"] = request.Count
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScaleClusterNodePool"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(NodepoolId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScaleClusterNodePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 扩容Kubernetes集群
//
// @param request - ScaleOutClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleOutClusterResponse
func (client *Client) ScaleOutClusterWithContext(ctx context.Context, ClusterId *string, request *ScaleOutClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleOutClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloudMonitorFlags) {
		body["cloud_monitor_flags"] = request.CloudMonitorFlags
	}

	if !dara.IsNil(request.Count) {
		body["count"] = request.Count
	}

	if !dara.IsNil(request.CpuPolicy) {
		body["cpu_policy"] = request.CpuPolicy
	}

	if !dara.IsNil(request.ImageId) {
		body["image_id"] = request.ImageId
	}

	if !dara.IsNil(request.KeyPair) {
		body["key_pair"] = request.KeyPair
	}

	if !dara.IsNil(request.LoginPassword) {
		body["login_password"] = request.LoginPassword
	}

	if !dara.IsNil(request.RdsInstances) {
		body["rds_instances"] = request.RdsInstances
	}

	if !dara.IsNil(request.Runtime) {
		body["runtime"] = request.Runtime
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.Taints) {
		body["taints"] = request.Taints
	}

	if !dara.IsNil(request.UserData) {
		body["user_data"] = request.UserData
	}

	if !dara.IsNil(request.VswitchIds) {
		body["vswitch_ids"] = request.VswitchIds
	}

	if !dara.IsNil(request.WorkerAutoRenew) {
		body["worker_auto_renew"] = request.WorkerAutoRenew
	}

	if !dara.IsNil(request.WorkerAutoRenewPeriod) {
		body["worker_auto_renew_period"] = request.WorkerAutoRenewPeriod
	}

	if !dara.IsNil(request.WorkerDataDisks) {
		body["worker_data_disks"] = request.WorkerDataDisks
	}

	if !dara.IsNil(request.WorkerInstanceChargeType) {
		body["worker_instance_charge_type"] = request.WorkerInstanceChargeType
	}

	if !dara.IsNil(request.WorkerInstanceTypes) {
		body["worker_instance_types"] = request.WorkerInstanceTypes
	}

	if !dara.IsNil(request.WorkerPeriod) {
		body["worker_period"] = request.WorkerPeriod
	}

	if !dara.IsNil(request.WorkerPeriodUnit) {
		body["worker_period_unit"] = request.WorkerPeriodUnit
	}

	if !dara.IsNil(request.WorkerSystemDiskCategory) {
		body["worker_system_disk_category"] = request.WorkerSystemDiskCategory
	}

	if !dara.IsNil(request.WorkerSystemDiskSize) {
		body["worker_system_disk_size"] = request.WorkerSystemDiskSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScaleOutCluster"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScaleOutClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the ScanClusterVuls API operation to scan ACK clusters for potential security vulnerabilities, including container workload vulnerabilities, third-party software vulnerabilities, CVE vulnerabilities, WebCMS vulnerabilities, and Windows operating system vulnerabilities. Regularly scan clusters for security vulnerabilities to improve cluster security.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScanClusterVulsResponse
func (client *Client) ScanClusterVulsWithContext(ctx context.Context, clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScanClusterVulsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScanClusterVuls"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/vuls/scan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScanClusterVulsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts the specified alert rule.
//
// @param request - StartAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAlertResponse
func (client *Client) StartAlertWithContext(ctx context.Context, ClusterId *string, request *StartAlertRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertRuleGroupName) {
		body["alert_rule_group_name"] = request.AlertRuleGroupName
	}

	if !dara.IsNil(request.AlertRuleName) {
		body["alert_rule_name"] = request.AlertRuleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAlert"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/alert/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/alert_rule/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables alert rules in the ACK Alert Center. You can disable an entire alert rule group or an individual alert rule.
//
// @param request - StopAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAlertResponse
func (client *Client) StopAlertWithContext(ctx context.Context, ClusterId *string, request *StopAlertRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertRuleGroupName) {
		body["alert_rule_group_name"] = request.AlertRuleGroupName
	}

	if !dara.IsNil(request.AlertRuleName) {
		body["alert_rule_name"] = request.AlertRuleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAlert"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/alert/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/alert_rule/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes a cluster node pool, including node pool metadata and node information.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncClusterNodePoolResponse
func (client *Client) SyncClusterNodePoolWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SyncClusterNodePoolResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncClusterNodePool"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/sync_nodepools"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncClusterNodePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds tags to a cluster. Adding tag key-value pairs allows cluster developers and O&M engineers to classify and manage clusters more flexibly, and better supports requirements such as monitoring, cost analysis, and tenant isolation. You can call the TagResources operation to bind tags to a cluster.
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["region_id"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIds) {
		body["resource_ids"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		body["resource_type"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tags"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Uninstalls specified components from a cluster and specifies whether to delete the associated Alibaba Cloud resources.
//
// @param request - UnInstallClusterAddonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnInstallClusterAddonsResponse
func (client *Client) UnInstallClusterAddonsWithContext(ctx context.Context, ClusterId *string, request *UnInstallClusterAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnInstallClusterAddonsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ToArray(request.Addons),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnInstallClusterAddons"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/components/uninstall"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UnInstallClusterAddonsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes resource tags when the tag key-value pairs of a cluster are no longer needed.
//
// @param tmpReq - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, tmpReq *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UntagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceIds) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, dara.String("resource_ids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagKeys) {
		request.TagKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKeys, dara.String("tag_keys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["all"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["region_id"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIdsShrink) {
		query["resource_ids"] = request.ResourceIdsShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["resource_type"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeysShrink) {
		query["tag_keys"] = request.TagKeysShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tags"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Enables or disables the audit log feature for a specified ACK cluster, or updates the audit log configurations to record requests to the Kubernetes API and the corresponding results for tracing cluster operation history and troubleshooting cluster faults.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and pricing of <props="china">[Simple Log Service](https://www.aliyun.com/price/product#/sls/detail/sls)<props="intl">[Simple Log Service](https://www.alibabacloud.com/product/log-service/pricing).
//
// @param request - UpdateClusterAuditLogConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterAuditLogConfigResponse
func (client *Client) UpdateClusterAuditLogConfigWithContext(ctx context.Context, clusterid *string, request *UpdateClusterAuditLogConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateClusterAuditLogConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Disable) {
		body["disable"] = request.Disable
	}

	if !dara.IsNil(request.SlsProjectName) {
		body["sls_project_name"] = request.SlsProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClusterAuditLogConfig"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterid)) + "/audit_log"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClusterAuditLogConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the cluster inspection configuration.
//
// @param request - UpdateClusterInspectConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterInspectConfigResponse
func (client *Client) UpdateClusterInspectConfigWithContext(ctx context.Context, clusterId *string, request *UpdateClusterInspectConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateClusterInspectConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DisabledCheckItems) {
		body["disabledCheckItems"] = request.DisabledCheckItems
	}

	if !dara.IsNil(request.Enabled) {
		body["enabled"] = request.Enabled
	}

	if !dara.IsNil(request.ScheduleTime) {
		body["scheduleTime"] = request.ScheduleTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClusterInspectConfig"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/inspectConfig"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClusterInspectConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the contact group for an alert rule set in an ACK cluster.
//
// @param request - UpdateContactGroupForAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContactGroupForAlertResponse
func (client *Client) UpdateContactGroupForAlertWithContext(ctx context.Context, ClusterId *string, request *UpdateContactGroupForAlertRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateContactGroupForAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertRuleGroupName) {
		body["alert_rule_group_name"] = request.AlertRuleGroupName
	}

	if !dara.IsNil(request.ContactGroupIds) {
		body["contact_group_ids"] = request.ContactGroupIds
	}

	if !dara.IsNil(request.CrName) {
		body["cr_name"] = request.CrName
	}

	if !dara.IsNil(request.Namespace) {
		body["namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContactGroupForAlert"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/alert/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/alert_rule/contact_groups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContactGroupForAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the control plane component log configurations for an ACK managed cluster. ACK managed clusters support collecting control plane component logs and delivering them to your Simple Log Service (SLS) project. Control plane components include core components such as Kube API Server, Kube Scheduler, Kube Controller Manager, and Cloud Controller Manager. You can call the UpdateControlPlaneLog operation to modify the control plane component log configurations, such as the log retention period and the components to collect.
//
// @param request - UpdateControlPlaneLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateControlPlaneLogResponse
func (client *Client) UpdateControlPlaneLogWithContext(ctx context.Context, ClusterId *string, request *UpdateControlPlaneLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateControlPlaneLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Aliuid) {
		body["aliuid"] = request.Aliuid
	}

	if !dara.IsNil(request.Components) {
		body["components"] = request.Components
	}

	if !dara.IsNil(request.LogProject) {
		body["log_project"] = request.LogProject
	}

	if !dara.IsNil(request.LogTtl) {
		body["log_ttl"] = request.LogTtl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateControlPlaneLog"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/controlplanelog"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateControlPlaneLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the expiration time of the KubeConfig issued by an ACK cluster. By default, the KubeConfig expires in three years. You can use an Alibaba Cloud account to invoke the UpdateK8sClusterUserConfigExpire operation to specify the expiration time of the KubeConfig issued to a Resource Access Management (RAM) user or role in the ACK cluster. Valid values: 1 to 876000 hours.
//
// Description:
//
// - Only Alibaba Cloud accounts can call this operation.
//
// - If the KubeConfig credential used for the cluster is revoked, the custom expiration time configured for the cluster KubeConfig is also reset. You must call this operation to reconfigure the expiration time.
//
// @param request - UpdateK8sClusterUserConfigExpireRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateK8sClusterUserConfigExpireResponse
func (client *Client) UpdateK8sClusterUserConfigExpireWithContext(ctx context.Context, ClusterId *string, request *UpdateK8sClusterUserConfigExpireRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateK8sClusterUserConfigExpireResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExpireHour) {
		body["expire_hour"] = request.ExpireHour
	}

	if !dara.IsNil(request.User) {
		body["user"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateK8sClusterUserConfigExpire"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/k8s/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/user_config/expire"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &UpdateK8sClusterUserConfigExpireResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the encryption at rest configuration for Secrets in a specified cluster.
//
// Description:
//
// Do not disable or delete the KMS key used for encryption at rest in the KMS console or by calling API operations while enabling or disabling the feature, or after the feature is enabled. Otherwise, the API server of the cluster becomes unavailable. As a result, objects such as Secrets and ServiceAccounts cannot be retrieved, which affects the normal operation of your applications. For more information, see [Use Alibaba Cloud KMS to encrypt Secrets at rest](https://help.aliyun.com/document_detail/177372.html).
//
//   - The user or role that calls this API operation must be granted the required cluster RBAC permissions, such as the O&M engineer or administrator permissions. Otherwise, the "ForbiddenUpdateKMSState" error code is returned.
//
//   - After this API operation is successfully called, the cluster status changes to "Updating". After the update is complete, the cluster status changes back to "Running". You must wait at least 5 minutes after an update is complete before you can call this API operation for the same cluster again. Otherwise, the "409" status code is returned.
//
// @param request - UpdateKMSEncryptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKMSEncryptionResponse
func (client *Client) UpdateKMSEncryptionWithContext(ctx context.Context, ClusterId *string, request *UpdateKMSEncryptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKMSEncryptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DisableEncryption) {
		body["disable_encryption"] = request.DisableEncryption
	}

	if !dara.IsNil(request.KmsKeyId) {
		body["kms_key_id"] = request.KmsKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKMSEncryption"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/kms"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &UpdateKMSEncryptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a node component.
//
// @param request - UpdateNodePoolComponentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodePoolComponentResponse
func (client *Client) UpdateNodePoolComponentWithContext(ctx context.Context, clusterId *string, nodepoolId *string, request *UpdateNodePoolComponentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateNodePoolComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.DisableRolling) {
		body["disableRolling"] = request.DisableRolling
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.NodeNames) {
		body["nodeNames"] = request.NodeNames
	}

	if !dara.IsNil(request.RollingPolicy) {
		body["rollingPolicy"] = request.RollingPolicy
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNodePoolComponent"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(nodepoolId)) + "/component"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNodePoolComponentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a node component.
//
// @param request - UpdateNodePoolComponentInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodePoolComponentInstanceResponse
func (client *Client) UpdateNodePoolComponentInstanceWithContext(ctx context.Context, clusterId *string, nodepoolId *string, name *string, request *UpdateNodePoolComponentInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateNodePoolComponentInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.DisableRolling) {
		body["disable_rolling"] = request.DisableRolling
	}

	if !dara.IsNil(request.NodeNames) {
		body["node_names"] = request.NodeNames
	}

	if !dara.IsNil(request.RollingPolicy) {
		body["rolling_policy"] = request.RollingPolicy
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNodePoolComponentInstance"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(clusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(nodepoolId)) + "/component_instances/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNodePoolComponentInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the deletion protection status of a specified resource. Supported resource types include namespaces and services. You can enable deletion protection for namespaces or services that involve critical workloads or sensitive data to prevent accidental deletion and reduce maintenance costs.
//
// Description:
//
// Before you call this operation, you must install or upgrade the security policy component for the cluster. For more information, see [Enable security policy management](https://help.aliyun.com/document_detail/359818.html).
//
// @param request - UpdateResourcesDeleteProtectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourcesDeleteProtectionResponse
func (client *Client) UpdateResourcesDeleteProtectionWithContext(ctx context.Context, ClusterId *string, request *UpdateResourcesDeleteProtectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourcesDeleteProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.Namespace) {
		body["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ResourceType) {
		body["resource_type"] = request.ResourceType
	}

	if !dara.IsNil(request.Resources) {
		body["resources"] = request.Resources
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourcesDeleteProtection"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/resources/protection"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourcesDeleteProtectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of an orchestration template. An orchestration template defines and describes a set of Kubernetes cluster resources declaratively to specify how an application should run or be configured.
//
// @param request - UpdateTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithContext(ctx context.Context, TemplateId *string, request *UpdateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.Template) {
		body["template"] = request.Template
	}

	if !dara.IsNil(request.TemplateType) {
		body["template_type"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTemplate"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/templates/" + dara.PercentEncode(dara.StringValue(TemplateId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the Role-Based Access Control (RBAC) access permissions of a Resource Access Management (RAM) user or RAM role in an ACK cluster. By default, RAM users, RAM roles, and users other than the cluster creator do not have any RBAC access permissions in the cluster. You can invoke the UpdateUserPermissions operation to update the RBAC access permissions of a RAM user or RAM role, including the accessible resources, the scope of the permissions, and the preset role types, to improve secure access control in cluster management.
//
// Description:
//
// You can update the cluster authorization information of a target Resource Access Management (RAM) user or RAM role by using a full update or an incremental update. A full update overwrites the existing cluster permissions of the target RAM user or RAM role. The request must include all permission configurations that you want to grant to the target RAM user or RAM role. An incremental update includes add and delete operations. It only modifies the permissions based on the cluster authorization information included in the request and does not affect other cluster permissions of the RAM user or RAM role.
//
// @param request - UpdateUserPermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserPermissionsResponse
func (client *Client) UpdateUserPermissionsWithContext(ctx context.Context, uid *string, request *UpdateUserPermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUserPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mode) {
		query["mode"] = request.Mode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserPermissions"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/permissions/users/" + dara.PercentEncode(dara.StringValue(uid)) + "/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &UpdateUserPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades an ACK cluster to mitigate potential security and stability risks associated with outdated versions and enable new features of the latest cluster version. You can call the UpgradeCluster operation to manually upgrade the cluster.
//
// Description:
//
// After you successfully call the UpgradeCluster operation, the API returns the task_id of the upgrade task. You can manage the task by calling the following task operations:
//
// - [Call DescribeTaskInfo to query task details](https://help.aliyun.com/document_detail/2667985.html)
//
// - [Call PauseTask to pause a running task](https://help.aliyun.com/document_detail/2667986.html)
//
// - [Call ResumeTask to resume a paused task](https://help.aliyun.com/document_detail/2667987.html)
//
// - [Call CancelTask to cancel a running task](https://help.aliyun.com/document_detail/2667988.html)
//
//	Notice: Starting from July 4, 2026, the request parameters rolling_policy and rolling_policy.max_parallelism no longer take effect. Use [UpgradeClusterNodepool](https://help.aliyun.com/document_detail/2667922.html) to upgrade worker nodes. For more information about the changes, see [Announcement on changes to request and response parameters of ACK cluster management operations and operation deprecation](https://help.aliyun.com/document_detail/2932733.html).</notice>
//
// @param request - UpgradeClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeClusterResponse
func (client *Client) UpgradeClusterWithContext(ctx context.Context, ClusterId *string, request *UpgradeClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComponentName) {
		body["component_name"] = request.ComponentName
	}

	if !dara.IsNil(request.MasterOnly) {
		body["master_only"] = request.MasterOnly
	}

	if !dara.IsNil(request.NextVersion) {
		body["next_version"] = request.NextVersion
	}

	if !dara.IsNil(request.RollingPolicy) {
		body["rolling_policy"] = request.RollingPolicy
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeCluster"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/upgrade"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the versions of cluster component instances so that you can benefit from the feature optimizations in the new versions.
//
// Description:
//
// - Upgrading cluster component instance versions may affect your business. Assess the impact before performing the operation during off-peak hours and back up relevant data in advance.
//
// - Before upgrading a component, refer to [Component release notes](https://help.aliyun.com/document_detail/176087.html) to learn about the changes and their impact for the specified component.
//
// - Upgrade components one at a time. Confirm that one component has been successfully upgraded before upgrading the next.
//
// @param request - UpgradeClusterAddonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeClusterAddonsResponse
func (client *Client) UpgradeClusterAddonsWithContext(ctx context.Context, ClusterId *string, request *UpgradeClusterAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeClusterAddonsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeClusterAddons"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/components/upgrade"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeClusterAddonsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the kubelet version (recommended to match the control plane version), operating system version, or container runtime version of a specified cluster node pool.
//
// Description:
//
// Upgrades the Kubernetes version, operating system version, or container runtime version of nodes in a specified cluster node pool. After you successfully call the UpgradeClusterNodepool operation, the API returns a `task_id` for the upgrade task. You can manage the task by calling the following task API operations:
//
// - [Call DescribeTaskInfo to query task details](https://help.aliyun.com/document_detail/2667985.html)
//
// - [Call PauseTask to pause a running task](https://help.aliyun.com/document_detail/2667986.html)
//
// - [Call ResumeTask to resume a paused task](https://help.aliyun.com/document_detail/2667987.html)
//
// - [Call CancelTask to cancel a running task](https://help.aliyun.com/document_detail/2667988.html)
//
// @param request - UpgradeClusterNodepoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeClusterNodepoolResponse
func (client *Client) UpgradeClusterNodepoolWithContext(ctx context.Context, ClusterId *string, NodepoolId *string, request *UpgradeClusterNodepoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeClusterNodepoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IgnoreWarningCheck) {
		body["ignore_warning_check"] = request.IgnoreWarningCheck
	}

	if !dara.IsNil(request.ImageId) {
		body["image_id"] = request.ImageId
	}

	if !dara.IsNil(request.KubernetesVersion) {
		body["kubernetes_version"] = request.KubernetesVersion
	}

	if !dara.IsNil(request.NodeNames) {
		body["node_names"] = request.NodeNames
	}

	if !dara.IsNil(request.RollingPolicy) {
		body["rolling_policy"] = request.RollingPolicy
	}

	if !dara.IsNil(request.RuntimeType) {
		body["runtime_type"] = request.RuntimeType
	}

	if !dara.IsNil(request.RuntimeVersion) {
		body["runtime_version"] = request.RuntimeVersion
	}

	if !dara.IsNil(request.UseReplace) {
		body["use_replace"] = request.UseReplace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeClusterNodepool"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(NodepoolId)) + "/upgrade"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeClusterNodepoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
