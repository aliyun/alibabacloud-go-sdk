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
		"ap-northeast-2-pop":          dara.String("cs.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("cs.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("cs.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("cs.aliyuncs.com"),
		"cn-edge-1":                   dara.String("cs.aliyuncs.com"),
		"cn-fujian":                   dara.String("cs.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("cs.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("cs.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("cs.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("cs.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("cs.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("cs.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("cs.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("cs.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("cs.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("cs.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("cs.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("cs.aliyuncs.com"),
		"cn-wuhan":                    dara.String("cs.aliyuncs.com"),
		"cn-yushanfang":               dara.String("cs.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("cs.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("cs.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("cs.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("cs.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("cs.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("cs.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds existing Elastic Compute Service (ECS) instances to a Container Service for Kubernetes (ACK) cluster.
//
// @param request - AttachInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachInstancesResponse
func (client *Client) AttachInstancesWithOptions(ClusterId *string, request *AttachInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AttachInstancesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds existing Elastic Compute Service (ECS) instances to a Container Service for Kubernetes (ACK) cluster.
//
// @param request - AttachInstancesRequest
//
// @return AttachInstancesResponse
func (client *Client) AttachInstances(ClusterId *string, request *AttachInstancesRequest) (_result *AttachInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AttachInstancesResponse{}
	_body, _err := client.AttachInstancesWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds existing nodes to a specific node pool. You can add existing ECS instances to a specific node pool in a Container Service for Kubernetes (ACK) cluster as worker nodes. You can also add removed worker nodes back to the node pool.
//
// @param request - AttachInstancesToNodePoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachInstancesToNodePoolResponse
func (client *Client) AttachInstancesToNodePoolWithOptions(ClusterId *string, NodepoolId *string, request *AttachInstancesToNodePoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AttachInstancesToNodePoolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds existing nodes to a specific node pool. You can add existing ECS instances to a specific node pool in a Container Service for Kubernetes (ACK) cluster as worker nodes. You can also add removed worker nodes back to the node pool.
//
// @param request - AttachInstancesToNodePoolRequest
//
// @return AttachInstancesToNodePoolResponse
func (client *Client) AttachInstancesToNodePool(ClusterId *string, NodepoolId *string, request *AttachInstancesToNodePoolRequest) (_result *AttachInstancesToNodePoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AttachInstancesToNodePoolResponse{}
	_body, _err := client.AttachInstancesToNodePoolWithOptions(ClusterId, NodepoolId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CancelClusterUpgrade is deprecated
//
// Summary:
//
// You can call the CancelClusterUpgrade operation to cancel the update of a cluster.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelClusterUpgradeResponse
func (client *Client) CancelClusterUpgradeWithOptions(ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelClusterUpgradeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the CancelClusterUpgrade operation to cancel the update of a cluster.
//
// @return CancelClusterUpgradeResponse
// Deprecated
func (client *Client) CancelClusterUpgrade(ClusterId *string) (_result *CancelClusterUpgradeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelClusterUpgradeResponse{}
	_body, _err := client.CancelClusterUpgradeWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CancelComponentUpgrade is deprecated
//
// Summary:
//
// You can call the CancelComponentUpgrade operation to cancel the update of a component.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelComponentUpgradeResponse
func (client *Client) CancelComponentUpgradeWithOptions(clusterId *string, componentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelComponentUpgradeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the CancelComponentUpgrade operation to cancel the update of a component.
//
// @return CancelComponentUpgradeResponse
// Deprecated
func (client *Client) CancelComponentUpgrade(clusterId *string, componentId *string) (_result *CancelComponentUpgradeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelComponentUpgradeResponse{}
	_body, _err := client.CancelComponentUpgradeWithOptions(clusterId, componentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the CancelOperationPlan operation to cancel a pending auto O\\\\\\\\\\\\&M plan.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelOperationPlanResponse
func (client *Client) CancelOperationPlanWithOptions(planId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelOperationPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the CancelOperationPlan operation to cancel a pending auto O\\\\\\\\\\\\&M plan.
//
// @return CancelOperationPlanResponse
func (client *Client) CancelOperationPlan(planId *string) (_result *CancelOperationPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelOperationPlanResponse{}
	_body, _err := client.CancelOperationPlanWithOptions(planId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelTaskWithOptions(taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CancelTaskResponse
func (client *Client) CancelTask(taskId *string) (_result *CancelTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelTaskResponse{}
	_body, _err := client.CancelTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the current log configuration of control plane components, including the log retention period and the log collection component. Container Service for Kubernetes (ACK) managed clusters can collect the logs of control plane components and deliver the logs to projects in Simple Log Service. These control plane components include Kube API Server, Kube Scheduler, Kube Controller Manager, and Cloud Controller Manager.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckControlPlaneLogEnableResponse
func (client *Client) CheckControlPlaneLogEnableWithOptions(ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckControlPlaneLogEnableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current log configuration of control plane components, including the log retention period and the log collection component. Container Service for Kubernetes (ACK) managed clusters can collect the logs of control plane components and deliver the logs to projects in Simple Log Service. These control plane components include Kube API Server, Kube Scheduler, Kube Controller Manager, and Cloud Controller Manager.
//
// @return CheckControlPlaneLogEnableResponse
func (client *Client) CheckControlPlaneLogEnable(ClusterId *string) (_result *CheckControlPlaneLogEnableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckControlPlaneLogEnableResponse{}
	_body, _err := client.CheckControlPlaneLogEnableWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the specified service roles are granted to Container Service for Kubernetes (ACK) within the current Alibaba Cloud account. ACK can access other cloud services, such as Elastic Compute Service (ECS), Object Storage Service (OSS), File Storage NAS (NAS), and Server Load Balancer (SLB), only after ACK is assigned the required service roles.
//
// @param request - CheckServiceRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceRoleResponse
func (client *Client) CheckServiceRoleWithOptions(request *CheckServiceRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckServiceRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the specified service roles are granted to Container Service for Kubernetes (ACK) within the current Alibaba Cloud account. ACK can access other cloud services, such as Elastic Compute Service (ECS), Object Storage Service (OSS), File Storage NAS (NAS), and Server Load Balancer (SLB), only after ACK is assigned the required service roles.
//
// @param request - CheckServiceRoleRequest
//
// @return CheckServiceRoleResponse
func (client *Client) CheckServiceRole(request *CheckServiceRoleRequest) (_result *CheckServiceRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckServiceRoleResponse{}
	_body, _err := client.CheckServiceRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes kubeconfig files that may pose potential risks from a user and revokes Role-Based Access Control (RBAC) permissions on a cluster.
//
// Description:
//
// >
//
//   - To call this operation, make sure that you have the AliyunCSFullAccess permission.
//
//   - You cannot revoke the permissions of an Alibaba Cloud account.
//
//   - You cannot revoke the permissions of the account that you use to call this operation.
//
// @param request - CleanClusterUserPermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CleanClusterUserPermissionsResponse
func (client *Client) CleanClusterUserPermissionsWithOptions(ClusterId *string, Uid *string, request *CleanClusterUserPermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CleanClusterUserPermissionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes kubeconfig files that may pose potential risks from a user and revokes Role-Based Access Control (RBAC) permissions on a cluster.
//
// Description:
//
// >
//
//   - To call this operation, make sure that you have the AliyunCSFullAccess permission.
//
//   - You cannot revoke the permissions of an Alibaba Cloud account.
//
//   - You cannot revoke the permissions of the account that you use to call this operation.
//
// @param request - CleanClusterUserPermissionsRequest
//
// @return CleanClusterUserPermissionsResponse
func (client *Client) CleanClusterUserPermissions(ClusterId *string, Uid *string, request *CleanClusterUserPermissionsRequest) (_result *CleanClusterUserPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CleanClusterUserPermissionsResponse{}
	_body, _err := client.CleanClusterUserPermissionsWithOptions(ClusterId, Uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the CleanUserPermissions operation to delete the kubeconfig files of the specified users and revoke the relevant Role-Based Access Control (RBAC) permissions. This API operation is suitable for scenarios where employees have resigned or the accounts of employees are locked.
//
// Description:
//
// > - To call this operation, make sure that you have the AliyunCSFullAccess permission.
//
// > - You cannot revoke the permissions of an Alibaba Cloud account.
//
// > - You cannot revoke the permissions of the account that you use to call this operation.
//
// @param tmpReq - CleanUserPermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CleanUserPermissionsResponse
func (client *Client) CleanUserPermissionsWithOptions(Uid *string, tmpReq *CleanUserPermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CleanUserPermissionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the CleanUserPermissions operation to delete the kubeconfig files of the specified users and revoke the relevant Role-Based Access Control (RBAC) permissions. This API operation is suitable for scenarios where employees have resigned or the accounts of employees are locked.
//
// Description:
//
// > - To call this operation, make sure that you have the AliyunCSFullAccess permission.
//
// > - You cannot revoke the permissions of an Alibaba Cloud account.
//
// > - You cannot revoke the permissions of the account that you use to call this operation.
//
// @param request - CleanUserPermissionsRequest
//
// @return CleanUserPermissionsResponse
func (client *Client) CleanUserPermissions(Uid *string, request *CleanUserPermissionsRequest) (_result *CleanUserPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CleanUserPermissionsResponse{}
	_body, _err := client.CleanUserPermissionsWithOptions(Uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a scaling configuration to allow the system to scale resources based on the given scaling rules. When you create a scaling configuration, you can specify the scaling metrics, thresholds, scaling order, and scaling interval.
//
// @param request - CreateAutoscalingConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoscalingConfigResponse
func (client *Client) CreateAutoscalingConfigWithOptions(ClusterId *string, request *CreateAutoscalingConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAutoscalingConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scaling configuration to allow the system to scale resources based on the given scaling rules. When you create a scaling configuration, you can specify the scaling metrics, thresholds, scaling order, and scaling interval.
//
// @param request - CreateAutoscalingConfigRequest
//
// @return CreateAutoscalingConfigResponse
func (client *Client) CreateAutoscalingConfig(ClusterId *string, request *CreateAutoscalingConfigRequest) (_result *CreateAutoscalingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAutoscalingConfigResponse{}
	_body, _err := client.CreateAutoscalingConfigWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Container Service for Kubernetes (ACK) cluster. For example, you can create an ACK managed cluster, ACK Serverless cluster, ACK Edge cluster, or registered cluster. When you create an ACK cluster, you need to configure the cluster information, components, and cloud resources used by ACK.
//
// Description:
//
// ### [](#-openapi-)Generate API request parameters through the ACK console
//
// When calling the CreateCluster operation to create a cluster, if the API call fails due to invalid parameter settings, you can generate valid request parameters through the ACK console. Follow these steps:
//
// 1.  Log on to the [ACK console](https://csnew.console.aliyun.com). In the left-side navigation pane, click **Clusters**.
//
// 2.  On the **Clusters*	- page, click **Cluster Templates**.
//
// 3.  In the Select Cluster Template dialog box, select the type of cluster you want to create and click Create. Then, configure the cluster parameters.
//
// 4.  In the **Confirm*	- step, click **Generate API Request Parameters**.
//
//	The API request parameters are displayed in the API Request Parameters dialog box.
//
// @param request - CreateClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Container Service for Kubernetes (ACK) cluster. For example, you can create an ACK managed cluster, ACK Serverless cluster, ACK Edge cluster, or registered cluster. When you create an ACK cluster, you need to configure the cluster information, components, and cloud resources used by ACK.
//
// Description:
//
// ### [](#-openapi-)Generate API request parameters through the ACK console
//
// When calling the CreateCluster operation to create a cluster, if the API call fails due to invalid parameter settings, you can generate valid request parameters through the ACK console. Follow these steps:
//
// 1.  Log on to the [ACK console](https://csnew.console.aliyun.com). In the left-side navigation pane, click **Clusters**.
//
// 2.  On the **Clusters*	- page, click **Cluster Templates**.
//
// 3.  In the Select Cluster Template dialog box, select the type of cluster you want to create and click Create. Then, configure the cluster parameters.
//
// 4.  In the **Confirm*	- step, click **Generate API Request Parameters**.
//
//	The API request parameters are displayed in the API Request Parameters dialog box.
//
// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a cluster diagnostic.
//
// @param request - CreateClusterDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterDiagnosisResponse
func (client *Client) CreateClusterDiagnosisWithOptions(clusterId *string, request *CreateClusterDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterDiagnosisResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a cluster diagnostic.
//
// @param request - CreateClusterDiagnosisRequest
//
// @return CreateClusterDiagnosisResponse
func (client *Client) CreateClusterDiagnosis(clusterId *string, request *CreateClusterDiagnosisRequest) (_result *CreateClusterDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateClusterDiagnosisResponse{}
	_body, _err := client.CreateClusterDiagnosisWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures cluster inspection.
//
// @param request - CreateClusterInspectConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterInspectConfigResponse
func (client *Client) CreateClusterInspectConfigWithOptions(clusterId *string, request *CreateClusterInspectConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterInspectConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures cluster inspection.
//
// @param request - CreateClusterInspectConfigRequest
//
// @return CreateClusterInspectConfigResponse
func (client *Client) CreateClusterInspectConfig(clusterId *string, request *CreateClusterInspectConfigRequest) (_result *CreateClusterInspectConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateClusterInspectConfigResponse{}
	_body, _err := client.CreateClusterInspectConfigWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a node pool for a Container Service for Kubernetes (ACK) cluster. You can use node pools to facilitate node management. For example, you can schedule, configure, or maintain nodes by node pool, and enable auto scaling for a node pool. We recommend that you use a managed node pool, which can help automate specific O\\\\\\&M tasks for nodes, such as Common Vulnerabilities and Exposures (CVE) patching and node repair. This reduces your O\\\\\\&M workload.
//
// @param request - CreateClusterNodePoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterNodePoolResponse
func (client *Client) CreateClusterNodePoolWithOptions(ClusterId *string, request *CreateClusterNodePoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterNodePoolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a node pool for a Container Service for Kubernetes (ACK) cluster. You can use node pools to facilitate node management. For example, you can schedule, configure, or maintain nodes by node pool, and enable auto scaling for a node pool. We recommend that you use a managed node pool, which can help automate specific O\\\\\\&M tasks for nodes, such as Common Vulnerabilities and Exposures (CVE) patching and node repair. This reduces your O\\\\\\&M workload.
//
// @param request - CreateClusterNodePoolRequest
//
// @return CreateClusterNodePoolResponse
func (client *Client) CreateClusterNodePool(ClusterId *string, request *CreateClusterNodePoolRequest) (_result *CreateClusterNodePoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateClusterNodePoolResponse{}
	_body, _err := client.CreateClusterNodePoolWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateKubernetesTrigger is deprecated
//
// Summary:
//
// You can call the CreateKubernetesTrigger operation to create a trigger for an application.
//
// @param request - CreateKubernetesTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKubernetesTriggerResponse
func (client *Client) CreateKubernetesTriggerWithOptions(request *CreateKubernetesTriggerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKubernetesTriggerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the CreateKubernetesTrigger operation to create a trigger for an application.
//
// @param request - CreateKubernetesTriggerRequest
//
// @return CreateKubernetesTriggerResponse
// Deprecated
func (client *Client) CreateKubernetesTrigger(request *CreateKubernetesTriggerRequest) (_result *CreateKubernetesTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateKubernetesTriggerResponse{}
	_body, _err := client.CreateKubernetesTriggerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an orchestration template. An orchestration template defines and describes a group of Kubernetes resources. It declaratively describes the configuration of an application or how an application runs. You can use orchestration templates to manage resources in Kubernetes clusters and automate resource deployment, such as pods, Services, Deployments, ConfigMaps, and persistent volumes (PVs).
//
// @param request - CreateTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an orchestration template. An orchestration template defines and describes a group of Kubernetes resources. It declaratively describes the configuration of an application or how an application runs. You can use orchestration templates to manage resources in Kubernetes clusters and automate resource deployment, such as pods, Services, Deployments, ConfigMaps, and persistent volumes (PVs).
//
// @param request - CreateTemplateRequest
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a trigger for an application to redeploy the application pods when specific conditions are met.
//
// @param request - CreateTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTriggerResponse
func (client *Client) CreateTriggerWithOptions(clusterId *string, request *CreateTriggerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTriggerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a trigger for an application to redeploy the application pods when specific conditions are met.
//
// @param request - CreateTriggerRequest
//
// @return CreateTriggerResponse
func (client *Client) CreateTrigger(clusterId *string, request *CreateTriggerRequest) (_result *CreateTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTriggerResponse{}
	_body, _err := client.CreateTriggerWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes one or more ACK alert contacts.
//
// @param tmpReq - DeleteAlertContactRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertContactResponse
func (client *Client) DeleteAlertContactWithOptions(tmpReq *DeleteAlertContactRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlertContactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more ACK alert contacts.
//
// @param request - DeleteAlertContactRequest
//
// @return DeleteAlertContactResponse
func (client *Client) DeleteAlertContact(request *DeleteAlertContactRequest) (_result *DeleteAlertContactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAlertContactResponse{}
	_body, _err := client.DeleteAlertContactWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an ACK alert contact group.
//
// @param tmpReq - DeleteAlertContactGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertContactGroupResponse
func (client *Client) DeleteAlertContactGroupWithOptions(tmpReq *DeleteAlertContactGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlertContactGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an ACK alert contact group.
//
// @param request - DeleteAlertContactGroupRequest
//
// @return DeleteAlertContactGroupResponse
func (client *Client) DeleteAlertContactGroup(request *DeleteAlertContactGroupRequest) (_result *DeleteAlertContactGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAlertContactGroupResponse{}
	_body, _err := client.DeleteAlertContactGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DeleteCluster operation to delete a cluster and specify whether to delete or retain the relevant cluster resources. Before you delete a cluster, you must manually delete workloads in the cluster, such as Deployments, StatefulSets, Jobs, and CronJobs. Otherwise, you may fail to delete the cluster.
//
// Description:
//
// Warning:
//
//   - Subscription ECS instances and Lingjun nodes in a cluster cannot be automatically released. To avoid unnecessary costs, we recommend that you manually release the resources. For more information, see \\<a href="{0}" target="_blank">Rules for deleting clusters and releasing nodes\\</a>.
//
//   - If the SLB instance of the API server uses the subscription billing method, it cannot be automatically released. To avoid unnecessary costs, we recommend that you manually release it.
//
//   - By default, virtual private clouds (VPCs), vSwitches, security groups, and RAM roles are retained if they are used by other resources. To avoid unnecessary costs, we recommend that you manually release the resources.
//
//   - Elastic container instances created on virtual nodes are automatically released.
//
//   - Some resources created together with a cluster are not automatically released when the cluster is deleted. After the cluster is deleted, you are still charged for the resources. Release or retain the resources based on your actual needs. The resources include Simple Log Service projects automatically created by the cluster and dynamically provisioned disks.
//
// @param tmpReq - DeleteClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithOptions(ClusterId *string, tmpReq *DeleteClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DeleteCluster operation to delete a cluster and specify whether to delete or retain the relevant cluster resources. Before you delete a cluster, you must manually delete workloads in the cluster, such as Deployments, StatefulSets, Jobs, and CronJobs. Otherwise, you may fail to delete the cluster.
//
// Description:
//
// Warning:
//
//   - Subscription ECS instances and Lingjun nodes in a cluster cannot be automatically released. To avoid unnecessary costs, we recommend that you manually release the resources. For more information, see \\<a href="{0}" target="_blank">Rules for deleting clusters and releasing nodes\\</a>.
//
//   - If the SLB instance of the API server uses the subscription billing method, it cannot be automatically released. To avoid unnecessary costs, we recommend that you manually release it.
//
//   - By default, virtual private clouds (VPCs), vSwitches, security groups, and RAM roles are retained if they are used by other resources. To avoid unnecessary costs, we recommend that you manually release the resources.
//
//   - Elastic container instances created on virtual nodes are automatically released.
//
//   - Some resources created together with a cluster are not automatically released when the cluster is deleted. After the cluster is deleted, you are still charged for the resources. Release or retain the resources based on your actual needs. The resources include Simple Log Service projects automatically created by the cluster and dynamically provisioned disks.
//
// @param request - DeleteClusterRequest
//
// @return DeleteClusterResponse
func (client *Client) DeleteCluster(ClusterId *string, request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes cluster inspection configurations.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterInspectConfigResponse
func (client *Client) DeleteClusterInspectConfigWithOptions(clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteClusterInspectConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes cluster inspection configurations.
//
// @return DeleteClusterInspectConfigResponse
func (client *Client) DeleteClusterInspectConfig(clusterId *string) (_result *DeleteClusterInspectConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteClusterInspectConfigResponse{}
	_body, _err := client.DeleteClusterInspectConfigWithOptions(clusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// null
//
// @param request - DeleteClusterNodepoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterNodepoolResponse
func (client *Client) DeleteClusterNodepoolWithOptions(ClusterId *string, NodepoolId *string, request *DeleteClusterNodepoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteClusterNodepoolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// null
//
// @param request - DeleteClusterNodepoolRequest
//
// @return DeleteClusterNodepoolResponse
func (client *Client) DeleteClusterNodepool(ClusterId *string, NodepoolId *string, request *DeleteClusterNodepoolRequest) (_result *DeleteClusterNodepoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteClusterNodepoolResponse{}
	_body, _err := client.DeleteClusterNodepoolWithOptions(ClusterId, NodepoolId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes nodes from a Container Service for Kubernetes (ACK) cluster when they are no longer required through the DeleteClusterNodes interface. When removing nodes, you can specify whether to release the Elastic Compute Service (ECS) instances and drain the nodes.
//
// Description:
//
//	  Use this API or the [ACK console](https://cs.console.aliyun.com) for node removal. Do not remove a node by running the `kubectl delete node` command.
//
//		- Never directly release or remove ECS instances through the ECS or Auto Scaling console/APIs. Renew subscription ECS instances before they expire. Violations may cause node termination and removal from the ACK console.
//
//		- If a node pool has the Expected Nodes parameter configured, the node pool automatically scales other ECS instances to maintain the expected number of nodes.
//
//		- When you remove a node, the pods on the node are migrated to other nodes. To prevent service interruptions, remove nodes during off-peak hours. Unexpected risks may arise during node removal. Back up the data in advance.
//
//		- ACK drains the node during node removal. Make sure that other nodes in the cluster have sufficient resources to host the evicted pods.
//
//		- To ensure the pods on the node you want to remove can be successfully scheduled to other nodes, check whether the node affinity rules and scheduling policies of the pods meet the requirements.
//
// @param request - DeleteClusterNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterNodesResponse
func (client *Client) DeleteClusterNodesWithOptions(ClusterId *string, request *DeleteClusterNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteClusterNodesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes nodes from a Container Service for Kubernetes (ACK) cluster when they are no longer required through the DeleteClusterNodes interface. When removing nodes, you can specify whether to release the Elastic Compute Service (ECS) instances and drain the nodes.
//
// Description:
//
//	  Use this API or the [ACK console](https://cs.console.aliyun.com) for node removal. Do not remove a node by running the `kubectl delete node` command.
//
//		- Never directly release or remove ECS instances through the ECS or Auto Scaling console/APIs. Renew subscription ECS instances before they expire. Violations may cause node termination and removal from the ACK console.
//
//		- If a node pool has the Expected Nodes parameter configured, the node pool automatically scales other ECS instances to maintain the expected number of nodes.
//
//		- When you remove a node, the pods on the node are migrated to other nodes. To prevent service interruptions, remove nodes during off-peak hours. Unexpected risks may arise during node removal. Back up the data in advance.
//
//		- ACK drains the node during node removal. Make sure that other nodes in the cluster have sufficient resources to host the evicted pods.
//
//		- To ensure the pods on the node you want to remove can be successfully scheduled to other nodes, check whether the node affinity rules and scheduling policies of the pods meet the requirements.
//
// @param request - DeleteClusterNodesRequest
//
// @return DeleteClusterNodesResponse
func (client *Client) DeleteClusterNodes(ClusterId *string, request *DeleteClusterNodesRequest) (_result *DeleteClusterNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteClusterNodesResponse{}
	_body, _err := client.DeleteClusterNodesWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteKubernetesTrigger is deprecated
//
// Summary:
//
// # You can call the DeleteKubernetesTrigger operation to delete an application trigger by trigger ID
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKubernetesTriggerResponse
func (client *Client) DeleteKubernetesTriggerWithOptions(Id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteKubernetesTriggerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// # You can call the DeleteKubernetesTrigger operation to delete an application trigger by trigger ID
//
// @return DeleteKubernetesTriggerResponse
// Deprecated
func (client *Client) DeleteKubernetesTrigger(Id *string) (_result *DeleteKubernetesTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteKubernetesTriggerResponse{}
	_body, _err := client.DeleteKubernetesTriggerWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes policy instances in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - DeletePolicyInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyInstanceResponse
func (client *Client) DeletePolicyInstanceWithOptions(clusterId *string, policyName *string, request *DeletePolicyInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePolicyInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes policy instances in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - DeletePolicyInstanceRequest
//
// @return DeletePolicyInstanceResponse
func (client *Client) DeletePolicyInstance(clusterId *string, policyName *string, request *DeletePolicyInstanceRequest) (_result *DeletePolicyInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePolicyInstanceResponse{}
	_body, _err := client.DeletePolicyInstanceWithOptions(clusterId, policyName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the orchestration templates that you no longer need.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithOptions(TemplateId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the orchestration templates that you no longer need.
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplate(TemplateId *string) (_result *DeleteTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(TemplateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteTriggerWithOptions(clusterId *string, Id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTriggerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteTriggerResponse
func (client *Client) DeleteTrigger(clusterId *string, Id *string) (_result *DeleteTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTriggerResponse{}
	_body, _err := client.DeleteTriggerWithOptions(clusterId, Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deploys a policy in the specified namespaces of a specific Container Service for Kubernetes (ACK) cluster. You can create and deploy a security policy by specifying the policy type, action of the policy such as alerting or denying, and namespaces to which the policy applies.
//
// @param request - DeployPolicyInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployPolicyInstanceResponse
func (client *Client) DeployPolicyInstanceWithOptions(clusterId *string, policyName *string, request *DeployPolicyInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeployPolicyInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys a policy in the specified namespaces of a specific Container Service for Kubernetes (ACK) cluster. You can create and deploy a security policy by specifying the policy type, action of the policy such as alerting or denying, and namespaces to which the policy applies.
//
// @param request - DeployPolicyInstanceRequest
//
// @return DeployPolicyInstanceResponse
func (client *Client) DeployPolicyInstance(clusterId *string, policyName *string, request *DeployPolicyInstanceRequest) (_result *DeployPolicyInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeployPolicyInstanceResponse{}
	_body, _err := client.DeployPolicyInstanceWithOptions(clusterId, policyName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a component based on specific conditions such as the region, cluster type, cluster subtype defined by cluster profile, cluster version, and component name. The information includes whether the component is managed, the component type, supported custom parameter schema, compatible operating system architecture, and earliest supported cluster version.
//
// @param request - DescribeAddonRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAddonResponse
func (client *Client) DescribeAddonWithOptions(addonName *string, request *DescribeAddonRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAddonResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a component based on specific conditions such as the region, cluster type, cluster subtype defined by cluster profile, cluster version, and component name. The information includes whether the component is managed, the component type, supported custom parameter schema, compatible operating system architecture, and earliest supported cluster version.
//
// @param request - DescribeAddonRequest
//
// @return DescribeAddonResponse
func (client *Client) DescribeAddon(addonName *string, request *DescribeAddonRequest) (_result *DescribeAddonResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAddonResponse{}
	_body, _err := client.DescribeAddonWithOptions(addonName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeAddons is deprecated
//
// Summary:
//
// You can call the DescribeAddons operation to query the details about all components that are supported by Container Service for Kubernetes (ACK).
//
// @param request - DescribeAddonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAddonsResponse
func (client *Client) DescribeAddonsWithOptions(request *DescribeAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAddonsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the DescribeAddons operation to query the details about all components that are supported by Container Service for Kubernetes (ACK).
//
// @param request - DescribeAddonsRequest
//
// @return DescribeAddonsResponse
// Deprecated
func (client *Client) DescribeAddons(request *DescribeAddonsRequest) (_result *DescribeAddonsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAddonsResponse{}
	_body, _err := client.DescribeAddonsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusterAddonInstance is deprecated
//
// Summary:
//
// You can call the DescribeClusterAddonInstance operation to query the information about a cluster component, including the version, status, and configuration of the component.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAddonInstanceResponse
func (client *Client) DescribeClusterAddonInstanceWithOptions(ClusterID *string, AddonName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterAddonInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the DescribeClusterAddonInstance operation to query the information about a cluster component, including the version, status, and configuration of the component.
//
// @return DescribeClusterAddonInstanceResponse
// Deprecated
func (client *Client) DescribeClusterAddonInstance(ClusterID *string, AddonName *string) (_result *DescribeClusterAddonInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterAddonInstanceResponse{}
	_body, _err := client.DescribeClusterAddonInstanceWithOptions(ClusterID, AddonName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusterAddonMetadata is deprecated
//
// Summary:
//
// You can call the DescribeClusterAddonMetadata operation to query the metadata of a component version. The metadata includes the component version and available parameters.
//
// @param request - DescribeClusterAddonMetadataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAddonMetadataResponse
func (client *Client) DescribeClusterAddonMetadataWithOptions(clusterId *string, componentId *string, request *DescribeClusterAddonMetadataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterAddonMetadataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the DescribeClusterAddonMetadata operation to query the metadata of a component version. The metadata includes the component version and available parameters.
//
// @param request - DescribeClusterAddonMetadataRequest
//
// @return DescribeClusterAddonMetadataResponse
// Deprecated
func (client *Client) DescribeClusterAddonMetadata(clusterId *string, componentId *string, request *DescribeClusterAddonMetadataRequest) (_result *DescribeClusterAddonMetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterAddonMetadataResponse{}
	_body, _err := client.DescribeClusterAddonMetadataWithOptions(clusterId, componentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusterAddonUpgradeStatus is deprecated
//
// Summary:
//
// You can call the DescribeClusterAddonUpgradeStatus operation to query the update progress of a cluster component.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAddonUpgradeStatusResponse
func (client *Client) DescribeClusterAddonUpgradeStatusWithOptions(ClusterId *string, ComponentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterAddonUpgradeStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the DescribeClusterAddonUpgradeStatus operation to query the update progress of a cluster component.
//
// @return DescribeClusterAddonUpgradeStatusResponse
// Deprecated
func (client *Client) DescribeClusterAddonUpgradeStatus(ClusterId *string, ComponentId *string) (_result *DescribeClusterAddonUpgradeStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterAddonUpgradeStatusResponse{}
	_body, _err := client.DescribeClusterAddonUpgradeStatusWithOptions(ClusterId, ComponentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusterAddonsUpgradeStatus is deprecated
//
// Summary:
//
// You can call the DescribeClusterAddonsUpgradeStatus operation to query the update progress of a component by component name.
//
// @param tmpReq - DescribeClusterAddonsUpgradeStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAddonsUpgradeStatusResponse
func (client *Client) DescribeClusterAddonsUpgradeStatusWithOptions(ClusterId *string, tmpReq *DescribeClusterAddonsUpgradeStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterAddonsUpgradeStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the DescribeClusterAddonsUpgradeStatus operation to query the update progress of a component by component name.
//
// @param request - DescribeClusterAddonsUpgradeStatusRequest
//
// @return DescribeClusterAddonsUpgradeStatusResponse
// Deprecated
func (client *Client) DescribeClusterAddonsUpgradeStatus(ClusterId *string, request *DescribeClusterAddonsUpgradeStatusRequest) (_result *DescribeClusterAddonsUpgradeStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterAddonsUpgradeStatusResponse{}
	_body, _err := client.DescribeClusterAddonsUpgradeStatusWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusterAddonsVersion is deprecated
//
// Summary:
//
// You can call the DescribeClusterAddonsVersion operation to query the details about all components in a cluster by cluster ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAddonsVersionResponse
func (client *Client) DescribeClusterAddonsVersionWithOptions(ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterAddonsVersionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the DescribeClusterAddonsVersion operation to query the details about all components in a cluster by cluster ID.
//
// @return DescribeClusterAddonsVersionResponse
// Deprecated
func (client *Client) DescribeClusterAddonsVersion(ClusterId *string) (_result *DescribeClusterAddonsVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterAddonsVersionResponse{}
	_body, _err := client.DescribeClusterAddonsVersionWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the scripts used to add existing nodes to a Container Service for Kubernetes (ACK) cluster. ACK allows you to manually add existing Elastic Compute Service (ECS) instances to an ACK cluster as worker nodes or re-add worker nodes that you remove from the cluster to a node pool.
//
// @param request - DescribeClusterAttachScriptsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAttachScriptsResponse
func (client *Client) DescribeClusterAttachScriptsWithOptions(ClusterId *string, request *DescribeClusterAttachScriptsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterAttachScriptsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scripts used to add existing nodes to a Container Service for Kubernetes (ACK) cluster. ACK allows you to manually add existing Elastic Compute Service (ECS) instances to an ACK cluster as worker nodes or re-add worker nodes that you remove from the cluster to a node pool.
//
// @param request - DescribeClusterAttachScriptsRequest
//
// @return DescribeClusterAttachScriptsResponse
func (client *Client) DescribeClusterAttachScripts(ClusterId *string, request *DescribeClusterAttachScriptsRequest) (_result *DescribeClusterAttachScriptsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterAttachScriptsResponse{}
	_body, _err := client.DescribeClusterAttachScriptsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeClusterDetail operation to query the details of a Container Service for Kubernetes (ACK) cluster by cluster ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterDetailResponse
func (client *Client) DescribeClusterDetailWithOptions(ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeClusterDetail operation to query the details of a Container Service for Kubernetes (ACK) cluster by cluster ID.
//
// @return DescribeClusterDetailResponse
func (client *Client) DescribeClusterDetail(ClusterId *string) (_result *DescribeClusterDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterDetailResponse{}
	_body, _err := client.DescribeClusterDetailWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries events and event details in a Container Service for Kubernetes (ACK) cluster, including the severity level, status, and start time of each event. Events are generated when clusters created, modified, and updated, node pools are created and scaled out, and components are installed.
//
// @param request - DescribeClusterEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterEventsResponse
func (client *Client) DescribeClusterEventsWithOptions(ClusterId *string, request *DescribeClusterEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterEventsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries events and event details in a Container Service for Kubernetes (ACK) cluster, including the severity level, status, and start time of each event. Events are generated when clusters created, modified, and updated, node pools are created and scaled out, and components are installed.
//
// @param request - DescribeClusterEventsRequest
//
// @return DescribeClusterEventsResponse
func (client *Client) DescribeClusterEvents(ClusterId *string, request *DescribeClusterEventsRequest) (_result *DescribeClusterEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterEventsResponse{}
	_body, _err := client.DescribeClusterEventsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the cluster log to help analyze cluster issues and locate the cause.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterLogsResponse
func (client *Client) DescribeClusterLogsWithOptions(ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterLogsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cluster log to help analyze cluster issues and locate the cause.
//
// @return DescribeClusterLogsResponse
func (client *Client) DescribeClusterLogs(ClusterId *string) (_result *DescribeClusterLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterLogsResponse{}
	_body, _err := client.DescribeClusterLogsWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeClusterNodePoolDetail.html operation to query the details about a node pool in a cluster by node pool ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterNodePoolDetailResponse
func (client *Client) DescribeClusterNodePoolDetailWithOptions(ClusterId *string, NodepoolId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterNodePoolDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeClusterNodePoolDetail.html operation to query the details about a node pool in a cluster by node pool ID.
//
// @return DescribeClusterNodePoolDetailResponse
func (client *Client) DescribeClusterNodePoolDetail(ClusterId *string, NodepoolId *string) (_result *DescribeClusterNodePoolDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterNodePoolDetailResponse{}
	_body, _err := client.DescribeClusterNodePoolDetailWithOptions(ClusterId, NodepoolId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all node pools in a cluster.
//
// @param request - DescribeClusterNodePoolsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterNodePoolsResponse
func (client *Client) DescribeClusterNodePoolsWithOptions(ClusterId *string, request *DescribeClusterNodePoolsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterNodePoolsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all node pools in a cluster.
//
// @param request - DescribeClusterNodePoolsRequest
//
// @return DescribeClusterNodePoolsResponse
func (client *Client) DescribeClusterNodePools(ClusterId *string, request *DescribeClusterNodePoolsRequest) (_result *DescribeClusterNodePoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterNodePoolsResponse{}
	_body, _err := client.DescribeClusterNodePoolsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// null
//
// @param request - DescribeClusterNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterNodesResponse
func (client *Client) DescribeClusterNodesWithOptions(ClusterId *string, request *DescribeClusterNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterNodesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// null
//
// @param request - DescribeClusterNodesRequest
//
// @return DescribeClusterNodesResponse
func (client *Client) DescribeClusterNodes(ClusterId *string, request *DescribeClusterNodesRequest) (_result *DescribeClusterNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterNodesResponse{}
	_body, _err := client.DescribeClusterNodesWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeClusterResources operation to query all resources in a cluster by cluster ID.
//
// @param request - DescribeClusterResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterResourcesResponse
func (client *Client) DescribeClusterResourcesWithOptions(ClusterId *string, request *DescribeClusterResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeClusterResources operation to query all resources in a cluster by cluster ID.
//
// @param request - DescribeClusterResourcesRequest
//
// @return DescribeClusterResourcesResponse
func (client *Client) DescribeClusterResources(ClusterId *string, request *DescribeClusterResourcesRequest) (_result *DescribeClusterResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterResourcesResponse{}
	_body, _err := client.DescribeClusterResourcesWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tasks in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - DescribeClusterTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterTasksResponse
func (client *Client) DescribeClusterTasksWithOptions(clusterId *string, request *DescribeClusterTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tasks in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - DescribeClusterTasksRequest
//
// @return DescribeClusterTasksResponse
func (client *Client) DescribeClusterTasks(clusterId *string, request *DescribeClusterTasksRequest) (_result *DescribeClusterTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterTasksResponse{}
	_body, _err := client.DescribeClusterTasksWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Kubeconfig files store identity and authentication information that is used by clients to access Container Service for Kubernetes (ACK) clusters. To use a kubectl client to manage an ACK cluster, you need to use the corresponding kubeconfig file to connect to the ACK cluster. We recommend that you keep kubeconfig files confidential and revoke kubeconfig files that are not in use. This helps prevent data leaks caused by the disclosure of kubeconfig files.
//
// Description:
//
//	  The default validity period of a kubeconfig file is 3 years. 180 days before a kubeconfig file expires, you can renew it in the Container Service for Kubernetes (ACK) console or by calling API operations. After a kubeconfig file is renewed, the kubeconfig file is valid for 3 years. The previous kubeconfig file still remains valid until expiration. We recommend that you renew your kubeconfig file at the earliest opportunity.
//
//		- We recommend that you keep kubeconfig files confidential and revoke kubeconfig files that are not in use. This helps prevent data leaks caused by the disclosure of kubeconfig files.
//
// @param request - DescribeClusterUserKubeconfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterUserKubeconfigResponse
func (client *Client) DescribeClusterUserKubeconfigWithOptions(ClusterId *string, request *DescribeClusterUserKubeconfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterUserKubeconfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Kubeconfig files store identity and authentication information that is used by clients to access Container Service for Kubernetes (ACK) clusters. To use a kubectl client to manage an ACK cluster, you need to use the corresponding kubeconfig file to connect to the ACK cluster. We recommend that you keep kubeconfig files confidential and revoke kubeconfig files that are not in use. This helps prevent data leaks caused by the disclosure of kubeconfig files.
//
// Description:
//
//	  The default validity period of a kubeconfig file is 3 years. 180 days before a kubeconfig file expires, you can renew it in the Container Service for Kubernetes (ACK) console or by calling API operations. After a kubeconfig file is renewed, the kubeconfig file is valid for 3 years. The previous kubeconfig file still remains valid until expiration. We recommend that you renew your kubeconfig file at the earliest opportunity.
//
//		- We recommend that you keep kubeconfig files confidential and revoke kubeconfig files that are not in use. This helps prevent data leaks caused by the disclosure of kubeconfig files.
//
// @param request - DescribeClusterUserKubeconfigRequest
//
// @return DescribeClusterUserKubeconfigResponse
func (client *Client) DescribeClusterUserKubeconfig(ClusterId *string, request *DescribeClusterUserKubeconfigRequest) (_result *DescribeClusterUserKubeconfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterUserKubeconfigResponse{}
	_body, _err := client.DescribeClusterUserKubeconfigWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusterV2UserKubeconfig is deprecated
//
// Summary:
//
// 获取集群kubeconfig接口
//
// @param request - DescribeClusterV2UserKubeconfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterV2UserKubeconfigResponse
func (client *Client) DescribeClusterV2UserKubeconfigWithOptions(ClusterId *string, request *DescribeClusterV2UserKubeconfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterV2UserKubeconfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 获取集群kubeconfig接口
//
// @param request - DescribeClusterV2UserKubeconfigRequest
//
// @return DescribeClusterV2UserKubeconfigResponse
// Deprecated
func (client *Client) DescribeClusterV2UserKubeconfig(ClusterId *string, request *DescribeClusterV2UserKubeconfigRequest) (_result *DescribeClusterV2UserKubeconfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterV2UserKubeconfigResponse{}
	_body, _err := client.DescribeClusterV2UserKubeconfigWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the security vulnerability details of a cluster by cluster ID. The details include vulnerability name, vulnerability type, and vulnerability severity. We recommend that you scan your cluster on a regular basis to ensure cluster security.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterVulsResponse
func (client *Client) DescribeClusterVulsWithOptions(clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterVulsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security vulnerability details of a cluster by cluster ID. The details include vulnerability name, vulnerability type, and vulnerability severity. We recommend that you scan your cluster on a regular basis to ensure cluster security.
//
// @return DescribeClusterVulsResponse
func (client *Client) DescribeClusterVuls(clusterId *string) (_result *DescribeClusterVulsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterVulsResponse{}
	_body, _err := client.DescribeClusterVulsWithOptions(clusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeClusters is deprecated
//
// Summary:
//
// Queries all the clusters that belong to the current Alibaba Cloud account, including Kubernetes clusters and Swarm clusters.
//
// @param request - DescribeClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClustersResponse
func (client *Client) DescribeClustersWithOptions(request *DescribeClustersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClustersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Queries all the clusters that belong to the current Alibaba Cloud account, including Kubernetes clusters and Swarm clusters.
//
// @param request - DescribeClustersRequest
//
// @return DescribeClustersResponse
// Deprecated
func (client *Client) DescribeClusters(request *DescribeClustersRequest) (_result *DescribeClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClustersResponse{}
	_body, _err := client.DescribeClustersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all clusters in a specified region.
//
// @param request - DescribeClustersForRegionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClustersForRegionResponse
func (client *Client) DescribeClustersForRegionWithOptions(regionId *string, request *DescribeClustersForRegionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClustersForRegionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all clusters in a specified region.
//
// @param request - DescribeClustersForRegionRequest
//
// @return DescribeClustersForRegionResponse
func (client *Client) DescribeClustersForRegion(regionId *string, request *DescribeClustersForRegionRequest) (_result *DescribeClustersForRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClustersForRegionResponse{}
	_body, _err := client.DescribeClustersForRegionWithOptions(regionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about Container Service for Kubernetes (ACK) clusters of specified types or specifications within an account.
//
// @param request - DescribeClustersV1Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClustersV1Response
func (client *Client) DescribeClustersV1WithOptions(request *DescribeClustersV1Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClustersV1Response, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about Container Service for Kubernetes (ACK) clusters of specified types or specifications within an account.
//
// @param request - DescribeClustersV1Request
//
// @return DescribeClustersV1Response
func (client *Client) DescribeClustersV1(request *DescribeClustersV1Request) (_result *DescribeClustersV1Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClustersV1Response{}
	_body, _err := client.DescribeClustersV1WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a type of events, including the severity level, status, and time. Events are generated when clusters are created, modified, and updated, node pools are created and scaled out, and components are installed.
//
// @param request - DescribeEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsResponse
func (client *Client) DescribeEventsWithOptions(request *DescribeEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a type of events, including the severity level, status, and time. Events are generated when clusters are created, modified, and updated, node pools are created and scaled out, and components are installed.
//
// @param request - DescribeEventsRequest
//
// @return DescribeEventsResponse
func (client *Client) DescribeEvents(request *DescribeEventsRequest) (_result *DescribeEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeEventsResponse{}
	_body, _err := client.DescribeEventsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all events in a specified region.
//
// @param request - DescribeEventsForRegionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsForRegionResponse
func (client *Client) DescribeEventsForRegionWithOptions(regionId *string, request *DescribeEventsForRegionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeEventsForRegionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all events in a specified region.
//
// @param request - DescribeEventsForRegionRequest
//
// @return DescribeEventsForRegionResponse
func (client *Client) DescribeEventsForRegion(regionId *string, request *DescribeEventsForRegionRequest) (_result *DescribeEventsForRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeEventsForRegionResponse{}
	_body, _err := client.DescribeEventsForRegionWithOptions(regionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeExternalAgent is deprecated
//
// Summary:
//
// Queries the proxy configurations of a registered cluster by cluster ID.
//
// Description:
//
// For more information, see [Register an external Kubernetes cluster](https://help.aliyun.com/document_detail/121053.html).
//
// @param request - DescribeExternalAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExternalAgentResponse
func (client *Client) DescribeExternalAgentWithOptions(ClusterId *string, request *DescribeExternalAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeExternalAgentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Queries the proxy configurations of a registered cluster by cluster ID.
//
// Description:
//
// For more information, see [Register an external Kubernetes cluster](https://help.aliyun.com/document_detail/121053.html).
//
// @param request - DescribeExternalAgentRequest
//
// @return DescribeExternalAgentResponse
// Deprecated
func (client *Client) DescribeExternalAgent(ClusterId *string, request *DescribeExternalAgentRequest) (_result *DescribeExternalAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeExternalAgentResponse{}
	_body, _err := client.DescribeExternalAgentWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about Kubernetes versions, including the version number, release date, expiration date, compatible OSs, and runtime.
//
// @param request - DescribeKubernetesVersionMetadataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKubernetesVersionMetadataResponse
func (client *Client) DescribeKubernetesVersionMetadataWithOptions(request *DescribeKubernetesVersionMetadataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeKubernetesVersionMetadataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about Kubernetes versions, including the version number, release date, expiration date, compatible OSs, and runtime.
//
// @param request - DescribeKubernetesVersionMetadataRequest
//
// @return DescribeKubernetesVersionMetadataResponse
func (client *Client) DescribeKubernetesVersionMetadata(request *DescribeKubernetesVersionMetadataRequest) (_result *DescribeKubernetesVersionMetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeKubernetesVersionMetadataResponse{}
	_body, _err := client.DescribeKubernetesVersionMetadataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the vulnerability information of a node pool, such as vulnerability names and severity levels, by specifying the ID of the node pool. We recommend that you periodically scan node pools for vulnerabilities to enhance cluster security.
//
// @param request - DescribeNodePoolVulsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodePoolVulsResponse
func (client *Client) DescribeNodePoolVulsWithOptions(clusterId *string, nodepoolId *string, request *DescribeNodePoolVulsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeNodePoolVulsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the vulnerability information of a node pool, such as vulnerability names and severity levels, by specifying the ID of the node pool. We recommend that you periodically scan node pools for vulnerabilities to enhance cluster security.
//
// @param request - DescribeNodePoolVulsRequest
//
// @return DescribeNodePoolVulsResponse
func (client *Client) DescribeNodePoolVuls(clusterId *string, nodepoolId *string, request *DescribeNodePoolVulsRequest) (_result *DescribeNodePoolVulsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeNodePoolVulsResponse{}
	_body, _err := client.DescribeNodePoolVulsWithOptions(clusterId, nodepoolId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of security policies. Container Service for Kubernetes (ACK) clusters offer a variety of built-in container security policies, such as Compliance, Infra, K8s-general, and pod security policy (PSP). You can use these policies to ensure the security of containers running in a production environment.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePoliciesResponse
func (client *Client) DescribePoliciesWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of security policies. Container Service for Kubernetes (ACK) clusters offer a variety of built-in container security policies, such as Compliance, Infra, K8s-general, and pod security policy (PSP). You can use these policies to ensure the security of containers running in a production environment.
//
// @return DescribePoliciesResponse
func (client *Client) DescribePolicies() (_result *DescribePoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePoliciesResponse{}
	_body, _err := client.DescribePoliciesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a policy. The information includes the content, action, and severity level of the policy. Container Service for Kubernetes (ACK) provides the following types of predefined security policies: Compliance, Infra, K8s-general, and pod security policy (PSP). These policies ensure that containers are running in the production environment in a secure manner.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyDetailsResponse
func (client *Client) DescribePolicyDetailsWithOptions(policyName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePolicyDetailsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a policy. The information includes the content, action, and severity level of the policy. Container Service for Kubernetes (ACK) provides the following types of predefined security policies: Compliance, Infra, K8s-general, and pod security policy (PSP). These policies ensure that containers are running in the production environment in a secure manner.
//
// @return DescribePolicyDetailsResponse
func (client *Client) DescribePolicyDetails(policyName *string) (_result *DescribePolicyDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePolicyDetailsResponse{}
	_body, _err := client.DescribePolicyDetailsWithOptions(policyName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Container Service for Kubernetes (ACK) clusters offer a variety of built-in container security policies, such as Compliance, Infra, K8s-general, and pod security policy (PSP). You can use these policies to ensure the security of containers running in a production environment. You can call the DescribePolicyGovernanceInCluster operation to query the details of policies for an ACK cluster. For example, you can query the number of policies that are enabled per severity level, the audit logs of policies, and the blocking and alerting information.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyGovernanceInClusterResponse
func (client *Client) DescribePolicyGovernanceInClusterWithOptions(clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePolicyGovernanceInClusterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Container Service for Kubernetes (ACK) clusters offer a variety of built-in container security policies, such as Compliance, Infra, K8s-general, and pod security policy (PSP). You can use these policies to ensure the security of containers running in a production environment. You can call the DescribePolicyGovernanceInCluster operation to query the details of policies for an ACK cluster. For example, you can query the number of policies that are enabled per severity level, the audit logs of policies, and the blocking and alerting information.
//
// @return DescribePolicyGovernanceInClusterResponse
func (client *Client) DescribePolicyGovernanceInCluster(clusterId *string) (_result *DescribePolicyGovernanceInClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePolicyGovernanceInClusterResponse{}
	_body, _err := client.DescribePolicyGovernanceInClusterWithOptions(clusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about policy instances of the specified type in a Container Service for Kubernetes (ACK) cluster, such as the policy description and severity level. You can choose a type of security policy for an ACK cluster, specify the action and applicable scope of the policy, and then create and deploy a policy instance.
//
// @param request - DescribePolicyInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyInstancesResponse
func (client *Client) DescribePolicyInstancesWithOptions(clusterId *string, request *DescribePolicyInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePolicyInstancesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about policy instances of the specified type in a Container Service for Kubernetes (ACK) cluster, such as the policy description and severity level. You can choose a type of security policy for an ACK cluster, specify the action and applicable scope of the policy, and then create and deploy a policy instance.
//
// @param request - DescribePolicyInstancesRequest
//
// @return DescribePolicyInstancesResponse
func (client *Client) DescribePolicyInstances(clusterId *string, request *DescribePolicyInstancesRequest) (_result *DescribePolicyInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePolicyInstancesResponse{}
	_body, _err := client.DescribePolicyInstancesWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the deployment of policy instances in the current Container Service for Kubernetes (ACK) cluster, including the number of policy instances of each type and the number of policy types of each severity level.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyInstancesStatusResponse
func (client *Client) DescribePolicyInstancesStatusWithOptions(clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePolicyInstancesStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deployment of policy instances in the current Container Service for Kubernetes (ACK) cluster, including the number of policy instances of each type and the number of policy types of each severity level.
//
// @return DescribePolicyInstancesStatusResponse
func (client *Client) DescribePolicyInstancesStatus(clusterId *string) (_result *DescribePolicyInstancesStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePolicyInstancesStatusResponse{}
	_body, _err := client.DescribePolicyInstancesStatusWithOptions(clusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether the deletion protection feature is enabled for the specified resources in the cluster. The resources that you can query include namespaces and Services.
//
// @param request - DescribeResourcesDeleteProtectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourcesDeleteProtectionResponse
func (client *Client) DescribeResourcesDeleteProtectionWithOptions(ClusterId *string, ResourceType *string, request *DescribeResourcesDeleteProtectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeResourcesDeleteProtectionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the deletion protection feature is enabled for the specified resources in the cluster. The resources that you can query include namespaces and Services.
//
// @param request - DescribeResourcesDeleteProtectionRequest
//
// @return DescribeResourcesDeleteProtectionResponse
func (client *Client) DescribeResourcesDeleteProtection(ClusterId *string, ResourceType *string, request *DescribeResourcesDeleteProtectionRequest) (_result *DescribeResourcesDeleteProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeResourcesDeleteProtectionResponse{}
	_body, _err := client.DescribeResourcesDeleteProtectionWithOptions(ClusterId, ResourceType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries or issues the kubeconfig credentials of a Resource Access Management (RAM) user or RAM role of the account. If you are the permission manager of a Container Service for Kubernetes (ACK) cluster, you can issue the kubeconfig credentials to a specific RAM user or RAM role of the account by using the Alibaba Cloud account. The kubeconfig credentials, which are used to connect to the ACK cluster, contain the identity information about the RAM user or RAM role.
//
// Description:
//
// You can call this operation only by using an Alibaba Cloud account.
//
// @param request - DescribeSubaccountK8sClusterUserConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubaccountK8sClusterUserConfigResponse
func (client *Client) DescribeSubaccountK8sClusterUserConfigWithOptions(ClusterId *string, Uid *string, request *DescribeSubaccountK8sClusterUserConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeSubaccountK8sClusterUserConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries or issues the kubeconfig credentials of a Resource Access Management (RAM) user or RAM role of the account. If you are the permission manager of a Container Service for Kubernetes (ACK) cluster, you can issue the kubeconfig credentials to a specific RAM user or RAM role of the account by using the Alibaba Cloud account. The kubeconfig credentials, which are used to connect to the ACK cluster, contain the identity information about the RAM user or RAM role.
//
// Description:
//
// You can call this operation only by using an Alibaba Cloud account.
//
// @param request - DescribeSubaccountK8sClusterUserConfigRequest
//
// @return DescribeSubaccountK8sClusterUserConfigResponse
func (client *Client) DescribeSubaccountK8sClusterUserConfig(ClusterId *string, Uid *string, request *DescribeSubaccountK8sClusterUserConfigRequest) (_result *DescribeSubaccountK8sClusterUserConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSubaccountK8sClusterUserConfigResponse{}
	_body, _err := client.DescribeSubaccountK8sClusterUserConfigWithOptions(ClusterId, Uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries detailed information about a task, such as the task type, status, and progress.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskInfoResponse
func (client *Client) DescribeTaskInfoWithOptions(taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTaskInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries detailed information about a task, such as the task type, status, and progress.
//
// @return DescribeTaskInfoResponse
func (client *Client) DescribeTaskInfo(taskId *string) (_result *DescribeTaskInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTaskInfoResponse{}
	_body, _err := client.DescribeTaskInfoWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// An orchestration template defines and describes a group of Kubernetes resources. It declaratively describes the configuration of an application or how an application runs. You can call the DescribeTemplates API operation to query orchestration templates and their detailed information, including access permissions, YAML content, and labels.
//
// @param request - DescribeTemplateAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplateAttributeResponse
func (client *Client) DescribeTemplateAttributeWithOptions(TemplateId *string, request *DescribeTemplateAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTemplateAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// An orchestration template defines and describes a group of Kubernetes resources. It declaratively describes the configuration of an application or how an application runs. You can call the DescribeTemplates API operation to query orchestration templates and their detailed information, including access permissions, YAML content, and labels.
//
// @param request - DescribeTemplateAttributeRequest
//
// @return DescribeTemplateAttributeResponse
func (client *Client) DescribeTemplateAttribute(TemplateId *string, request *DescribeTemplateAttributeRequest) (_result *DescribeTemplateAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTemplateAttributeResponse{}
	_body, _err := client.DescribeTemplateAttributeWithOptions(TemplateId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// An orchestration template defines and describes a group of Kubernetes resources. It declaratively describes the configuration of an application or how an application runs. You can call the DescribeTemplates API operation to query orchestration templates and their detailed information, including access permissions, YAML content, and labels.
//
// @param request - DescribeTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplatesResponse
func (client *Client) DescribeTemplatesWithOptions(request *DescribeTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// An orchestration template defines and describes a group of Kubernetes resources. It declaratively describes the configuration of an application or how an application runs. You can call the DescribeTemplates API operation to query orchestration templates and their detailed information, including access permissions, YAML content, and labels.
//
// @param request - DescribeTemplatesRequest
//
// @return DescribeTemplatesResponse
func (client *Client) DescribeTemplates(request *DescribeTemplatesRequest) (_result *DescribeTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.DescribeTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries triggers that match specific conditions.
//
// @param request - DescribeTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTriggerResponse
func (client *Client) DescribeTriggerWithOptions(clusterId *string, request *DescribeTriggerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTriggerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries triggers that match specific conditions.
//
// @param request - DescribeTriggerRequest
//
// @return DescribeTriggerResponse
func (client *Client) DescribeTrigger(clusterId *string, request *DescribeTriggerRequest) (_result *DescribeTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTriggerResponse{}
	_body, _err := client.DescribeTriggerWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can use Kubernetes namespaces to limit users from accessing resources in a Container Service for Kubernetes (ACK) cluster. Users that are granted Role-Based Access Control (RBAC) permissions only on one namespace cannot access resources in other namespaces. Queries the RBAC permissions that are granted to the current Resource Access Management (RAM) user or RAM role on an ACK cluster.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserClusterNamespacesResponse
func (client *Client) DescribeUserClusterNamespacesWithOptions(ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeUserClusterNamespacesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can use Kubernetes namespaces to limit users from accessing resources in a Container Service for Kubernetes (ACK) cluster. Users that are granted Role-Based Access Control (RBAC) permissions only on one namespace cannot access resources in other namespaces. Queries the RBAC permissions that are granted to the current Resource Access Management (RAM) user or RAM role on an ACK cluster.
//
// @return DescribeUserClusterNamespacesResponse
func (client *Client) DescribeUserClusterNamespaces(ClusterId *string) (_result *DescribeUserClusterNamespacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeUserClusterNamespacesResponse{}
	_body, _err := client.DescribeUserClusterNamespacesWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// In an Container Service for Kubernetes (ACK) cluster, you can create and specify different Resource Access Management (RAM) users or roles to have different access permissions. This ensures access control and resource isolation. You can call the DescribeUserPermission operation to query the permissions that are granted to a RAM user or RAM role on ACK clusters, including the resources that are allowed to access, the scope of the permissions, the predefined role, and the permission source.
//
// Description:
//
// *Precautions**:
//
//   - If you call this operation as a Resource Access Management (RAM) user or by assuming a RAM role, only the permissions granted on the clusters on which the current account has the role-based access control (RBAC) administrator permissions are returned. If you want to query the permissions on all clusters, you must use an account that has the RBAC administrator permissions on all clusters.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserPermissionResponse
func (client *Client) DescribeUserPermissionWithOptions(uid *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeUserPermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// In an Container Service for Kubernetes (ACK) cluster, you can create and specify different Resource Access Management (RAM) users or roles to have different access permissions. This ensures access control and resource isolation. You can call the DescribeUserPermission operation to query the permissions that are granted to a RAM user or RAM role on ACK clusters, including the resources that are allowed to access, the scope of the permissions, the predefined role, and the permission source.
//
// Description:
//
// *Precautions**:
//
//   - If you call this operation as a Resource Access Management (RAM) user or by assuming a RAM role, only the permissions granted on the clusters on which the current account has the role-based access control (RBAC) administrator permissions are returned. If you want to query the permissions on all clusters, you must use an account that has the RBAC administrator permissions on all clusters.
//
// @return DescribeUserPermissionResponse
func (client *Client) DescribeUserPermission(uid *string) (_result *DescribeUserPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeUserPermissionResponse{}
	_body, _err := client.DescribeUserPermissionWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries quotas related to Container Service for Kubernetes (ACK) clusters, node pools, and nodes. To increase a quota, submit an application in the Quota Center console.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserQuotaResponse
func (client *Client) DescribeUserQuotaWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeUserQuotaResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quotas related to Container Service for Kubernetes (ACK) clusters, node pools, and nodes. To increase a quota, submit an application in the Quota Center console.
//
// @return DescribeUserQuotaResponse
func (client *Client) DescribeUserQuota() (_result *DescribeUserQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeUserQuotaResponse{}
	_body, _err := client.DescribeUserQuotaWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Patches node vulnerabilities in a node pool to enhance node security. Cloud Security provided by Alibaba Cloud periodically scans Elastic Compute Service (ECS) instances for vulnerabilities and provides suggestions on how to patch the detected vulnerabilities. Vulnerability patching may require node restarts. Make sure that your cluster has sufficient idle nodes for node draining.
//
// Description:
//
// 1.  The Common Vulnerabilities and Exposures (CVE) patching feature is developed based on Security Center. To use this feature, you must purchase the Security Center Ultimate Edition that supports Container Service for Kubernetes (ACK).
//
// 2.  ACK may need to restart nodes to patch certain vulnerabilities. ACK drains a node before the node restarts. Make sure that the ACK cluster has sufficient idle nodes to host the pods evicted from the trained nodes. For example, you can scale out a node pool before you patch vulnerabilities for the nodes in the node pool.
//
// 3.  Security Center ensures the compatibility of CVE patches. We recommend that you check the compatibility of a CVE patch with your application before you install the patch. You can pause or cancel a CVE patching task anytime.
//
// 4.  CVE patching is a progressive task that consists of multiple batches. After you pause or cancel a CVE patching task, ACK continues to process the dispatched batches. Only the batches that have not been dispatched are paused or canceled.
//
// @param request - FixNodePoolVulsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FixNodePoolVulsResponse
func (client *Client) FixNodePoolVulsWithOptions(clusterId *string, nodepoolId *string, request *FixNodePoolVulsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FixNodePoolVulsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Patches node vulnerabilities in a node pool to enhance node security. Cloud Security provided by Alibaba Cloud periodically scans Elastic Compute Service (ECS) instances for vulnerabilities and provides suggestions on how to patch the detected vulnerabilities. Vulnerability patching may require node restarts. Make sure that your cluster has sufficient idle nodes for node draining.
//
// Description:
//
// 1.  The Common Vulnerabilities and Exposures (CVE) patching feature is developed based on Security Center. To use this feature, you must purchase the Security Center Ultimate Edition that supports Container Service for Kubernetes (ACK).
//
// 2.  ACK may need to restart nodes to patch certain vulnerabilities. ACK drains a node before the node restarts. Make sure that the ACK cluster has sufficient idle nodes to host the pods evicted from the trained nodes. For example, you can scale out a node pool before you patch vulnerabilities for the nodes in the node pool.
//
// 3.  Security Center ensures the compatibility of CVE patches. We recommend that you check the compatibility of a CVE patch with your application before you install the patch. You can pause or cancel a CVE patching task anytime.
//
// 4.  CVE patching is a progressive task that consists of multiple batches. After you pause or cancel a CVE patching task, ACK continues to process the dispatched batches. Only the batches that have not been dispatched are paused or canceled.
//
// @param request - FixNodePoolVulsRequest
//
// @return FixNodePoolVulsResponse
func (client *Client) FixNodePoolVuls(clusterId *string, nodepoolId *string, request *FixNodePoolVulsRequest) (_result *FixNodePoolVulsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FixNodePoolVulsResponse{}
	_body, _err := client.FixNodePoolVulsWithOptions(clusterId, nodepoolId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the GetClusterAddonInstance operation to query the information of a component instance in a cluster, including the version, configurations, and log status of the component instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterAddonInstanceResponse
func (client *Client) GetClusterAddonInstanceWithOptions(clusterId *string, instanceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterAddonInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the GetClusterAddonInstance operation to query the information of a component instance in a cluster, including the version, configurations, and log status of the component instance.
//
// @return GetClusterAddonInstanceResponse
func (client *Client) GetClusterAddonInstance(clusterId *string, instanceName *string) (_result *GetClusterAddonInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterAddonInstanceResponse{}
	_body, _err := client.GetClusterAddonInstanceWithOptions(clusterId, instanceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the GetClusterAuditProject operation to check whether the cluster has API Server auditing enabled and the corresponding Simple Log Service project that stores API Server audit logs.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterAuditProjectResponse
func (client *Client) GetClusterAuditProjectWithOptions(clusterid *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterAuditProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the GetClusterAuditProject operation to check whether the cluster has API Server auditing enabled and the corresponding Simple Log Service project that stores API Server audit logs.
//
// @return GetClusterAuditProjectResponse
func (client *Client) GetClusterAuditProject(clusterid *string) (_result *GetClusterAuditProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterAuditProjectResponse{}
	_body, _err := client.GetClusterAuditProjectWithOptions(clusterid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a cluster check task by cluster ID and task ID. You can view the status, check items, creation time, and end time of the task. Container Intelligence Service (CIS) provides a variety of Kubernetes cluster check features, including cluster update check, cluster migration check, component installation check, component update check, and node pool check.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterCheckResponse
func (client *Client) GetClusterCheckWithOptions(clusterId *string, checkId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterCheckResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a cluster check task by cluster ID and task ID. You can view the status, check items, creation time, and end time of the task. Container Intelligence Service (CIS) provides a variety of Kubernetes cluster check features, including cluster update check, cluster migration check, component installation check, component update check, and node pool check.
//
// @return GetClusterCheckResponse
func (client *Client) GetClusterCheck(clusterId *string, checkId *string) (_result *GetClusterCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterCheckResponse{}
	_body, _err := client.GetClusterCheckWithOptions(clusterId, checkId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries cluster diagnostic items.
//
// @param request - GetClusterDiagnosisCheckItemsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterDiagnosisCheckItemsResponse
func (client *Client) GetClusterDiagnosisCheckItemsWithOptions(clusterId *string, diagnosisId *string, request *GetClusterDiagnosisCheckItemsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterDiagnosisCheckItemsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cluster diagnostic items.
//
// @param request - GetClusterDiagnosisCheckItemsRequest
//
// @return GetClusterDiagnosisCheckItemsResponse
func (client *Client) GetClusterDiagnosisCheckItems(clusterId *string, diagnosisId *string, request *GetClusterDiagnosisCheckItemsRequest) (_result *GetClusterDiagnosisCheckItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterDiagnosisCheckItemsResponse{}
	_body, _err := client.GetClusterDiagnosisCheckItemsWithOptions(clusterId, diagnosisId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries cluster diagnostic results.
//
// @param request - GetClusterDiagnosisResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterDiagnosisResultResponse
func (client *Client) GetClusterDiagnosisResultWithOptions(clusterId *string, diagnosisId *string, request *GetClusterDiagnosisResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterDiagnosisResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cluster diagnostic results.
//
// @param request - GetClusterDiagnosisResultRequest
//
// @return GetClusterDiagnosisResultResponse
func (client *Client) GetClusterDiagnosisResult(clusterId *string, diagnosisId *string, request *GetClusterDiagnosisResultRequest) (_result *GetClusterDiagnosisResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterDiagnosisResultResponse{}
	_body, _err := client.GetClusterDiagnosisResultWithOptions(clusterId, diagnosisId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves cluster inspection configuration.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterInspectConfigResponse
func (client *Client) GetClusterInspectConfigWithOptions(clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterInspectConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves cluster inspection configuration.
//
// @return GetClusterInspectConfigResponse
func (client *Client) GetClusterInspectConfig(clusterId *string) (_result *GetClusterInspectConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterInspectConfigResponse{}
	_body, _err := client.GetClusterInspectConfigWithOptions(clusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Obtain the details of the inspection report for the cluster
//
// @param request - GetClusterInspectReportDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterInspectReportDetailResponse
func (client *Client) GetClusterInspectReportDetailWithOptions(clusterId *string, reportId *string, request *GetClusterInspectReportDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterInspectReportDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtain the details of the inspection report for the cluster
//
// @param request - GetClusterInspectReportDetailRequest
//
// @return GetClusterInspectReportDetailResponse
func (client *Client) GetClusterInspectReportDetail(clusterId *string, reportId *string, request *GetClusterInspectReportDetailRequest) (_result *GetClusterInspectReportDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterInspectReportDetailResponse{}
	_body, _err := client.GetClusterInspectReportDetailWithOptions(clusterId, reportId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetKubernetesTrigger is deprecated
//
// Summary:
//
// You can call the GetKubernetesTrigger operationto query the triggers of an application by application name.
//
// @param request - GetKubernetesTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKubernetesTriggerResponse
func (client *Client) GetKubernetesTriggerWithOptions(ClusterId *string, request *GetKubernetesTriggerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKubernetesTriggerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the GetKubernetesTrigger operationto query the triggers of an application by application name.
//
// @param request - GetKubernetesTriggerRequest
//
// @return GetKubernetesTriggerResponse
// Deprecated
func (client *Client) GetKubernetesTrigger(ClusterId *string, request *GetKubernetesTriggerRequest) (_result *GetKubernetesTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetKubernetesTriggerResponse{}
	_body, _err := client.GetKubernetesTriggerWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetUpgradeStatus is deprecated
//
// Summary:
//
// You can call the GetUpgradeStatus operation to query the update progress of a cluster by cluster ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUpgradeStatusResponse
func (client *Client) GetUpgradeStatusWithOptions(ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUpgradeStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the GetUpgradeStatus operation to query the update progress of a cluster by cluster ID.
//
// @return GetUpgradeStatusResponse
// Deprecated
func (client *Client) GetUpgradeStatus(ClusterId *string) (_result *GetUpgradeStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUpgradeStatusResponse{}
	_body, _err := client.GetUpgradeStatusWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the role-based access control (RBAC) permissions of a Resource Access Management (RAM) user or RAM role. By default, you do not have the RBAC permissions on a Container Service for Kubernetes (ACK) cluster if you are not the cluster owner or you are not using an Alibaba Cloud account. You can call this operation to specify the resources that can be accessed, permission scope, and predefined roles. This helps you better manage the access control on resources in ACK clusters.
//
// Description:
//
//	  If you use a Resource Access Management (RAM) account to call this operation, make sure it has permissions to modify cluster authorization information for other RAM users or RAM roles. Otherwise, the `StatusForbidden` or `ForbiddenGrantPermissions` error code is returned. For more information, see [Use a RAM user to grant RBAC permissions to other RAM users](https://help.aliyun.com/document_detail/119035.html).
//
//		- This operation overwrites all existing cluster permissions for the target RAM user or RAM role. You must specify all the permissions you want to grant to the RAM user or RAM role in the request.
//
// @param request - GrantPermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantPermissionsResponse
func (client *Client) GrantPermissionsWithOptions(uid *string, request *GrantPermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantPermissionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the role-based access control (RBAC) permissions of a Resource Access Management (RAM) user or RAM role. By default, you do not have the RBAC permissions on a Container Service for Kubernetes (ACK) cluster if you are not the cluster owner or you are not using an Alibaba Cloud account. You can call this operation to specify the resources that can be accessed, permission scope, and predefined roles. This helps you better manage the access control on resources in ACK clusters.
//
// Description:
//
//	  If you use a Resource Access Management (RAM) account to call this operation, make sure it has permissions to modify cluster authorization information for other RAM users or RAM roles. Otherwise, the `StatusForbidden` or `ForbiddenGrantPermissions` error code is returned. For more information, see [Use a RAM user to grant RBAC permissions to other RAM users](https://help.aliyun.com/document_detail/119035.html).
//
//		- This operation overwrites all existing cluster permissions for the target RAM user or RAM role. You must specify all the permissions you want to grant to the RAM user or RAM role in the request.
//
// @param request - GrantPermissionsRequest
//
// @return GrantPermissionsResponse
func (client *Client) GrantPermissions(uid *string, request *GrantPermissionsRequest) (_result *GrantPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GrantPermissionsResponse{}
	_body, _err := client.GrantPermissionsWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为了增强Kubernetes能力，ACK集群支持了多种组件，例如托管的核心组件，应用、日志和监控、网络、存储、安全组件等。您可以调用InstallClusterAddons接口，通过组件名称和版本安装组件。
//
// @param request - InstallClusterAddonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallClusterAddonsResponse
func (client *Client) InstallClusterAddonsWithOptions(ClusterId *string, request *InstallClusterAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallClusterAddonsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为了增强Kubernetes能力，ACK集群支持了多种组件，例如托管的核心组件，应用、日志和监控、网络、存储、安全组件等。您可以调用InstallClusterAddons接口，通过组件名称和版本安装组件。
//
// @param request - InstallClusterAddonsRequest
//
// @return InstallClusterAddonsResponse
func (client *Client) InstallClusterAddons(ClusterId *string, request *InstallClusterAddonsRequest) (_result *InstallClusterAddonsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallClusterAddonsResponse{}
	_body, _err := client.InstallClusterAddonsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为ACK集群节点池安装节点组件
//
// @param request - InstallNodePoolComponentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallNodePoolComponentsResponse
func (client *Client) InstallNodePoolComponentsWithOptions(clusterId *string, nodePoolId *string, request *InstallNodePoolComponentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallNodePoolComponentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为ACK集群节点池安装节点组件
//
// @param request - InstallNodePoolComponentsRequest
//
// @return InstallNodePoolComponentsResponse
func (client *Client) InstallNodePoolComponents(clusterId *string, nodePoolId *string, request *InstallNodePoolComponentsRequest) (_result *InstallNodePoolComponentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallNodePoolComponentsResponse{}
	_body, _err := client.InstallNodePoolComponentsWithOptions(clusterId, nodePoolId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available components based on specific conditions such as the region, cluster type, cluster subtype defined by cluster profile, and cluster version and queries the detailed information about a component. The information includes whether the component is managed, the supported custom parameter schema, and compatible operating system architecture.
//
// @param request - ListAddonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonsResponse
func (client *Client) ListAddonsWithOptions(request *ListAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAddonsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available components based on specific conditions such as the region, cluster type, cluster subtype defined by cluster profile, and cluster version and queries the detailed information about a component. The information includes whether the component is managed, the supported custom parameter schema, and compatible operating system architecture.
//
// @param request - ListAddonsRequest
//
// @return ListAddonsResponse
func (client *Client) ListAddons(request *ListAddonsRequest) (_result *ListAddonsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAddonsResponse{}
	_body, _err := client.ListAddonsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取集群组件实例的资源列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterAddonInstanceResourcesResponse
func (client *Client) ListClusterAddonInstanceResourcesWithOptions(clusterId *string, instanceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterAddonInstanceResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取集群组件实例的资源列表
//
// @return ListClusterAddonInstanceResourcesResponse
func (client *Client) ListClusterAddonInstanceResources(clusterId *string, instanceName *string) (_result *ListClusterAddonInstanceResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterAddonInstanceResourcesResponse{}
	_body, _err := client.ListClusterAddonInstanceResourcesWithOptions(clusterId, instanceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the component instances that are running in the specified cluster and the information about the component instances. The information includes the component version and status.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterAddonInstancesResponse
func (client *Client) ListClusterAddonInstancesWithOptions(clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterAddonInstancesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the component instances that are running in the specified cluster and the information about the component instances. The information includes the component version and status.
//
// @return ListClusterAddonInstancesResponse
func (client *Client) ListClusterAddonInstances(clusterId *string) (_result *ListClusterAddonInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterAddonInstancesResponse{}
	_body, _err := client.ListClusterAddonInstancesWithOptions(clusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ListClusterChecks operation to query all the cluster check results of a cluster.
//
// @param request - ListClusterChecksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterChecksResponse
func (client *Client) ListClusterChecksWithOptions(clusterId *string, request *ListClusterChecksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterChecksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ListClusterChecks operation to query all the cluster check results of a cluster.
//
// @param request - ListClusterChecksRequest
//
// @return ListClusterChecksResponse
func (client *Client) ListClusterChecks(clusterId *string, request *ListClusterChecksRequest) (_result *ListClusterChecksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterChecksResponse{}
	_body, _err := client.ListClusterChecksWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of the cluster inspection report.
//
// @param request - ListClusterInspectReportsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterInspectReportsResponse
func (client *Client) ListClusterInspectReportsWithOptions(clusterId *string, request *ListClusterInspectReportsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterInspectReportsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of the cluster inspection report.
//
// @param request - ListClusterInspectReportsRequest
//
// @return ListClusterInspectReportsResponse
func (client *Client) ListClusterInspectReports(clusterId *string, request *ListClusterInspectReportsRequest) (_result *ListClusterInspectReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterInspectReportsResponse{}
	_body, _err := client.ListClusterInspectReportsWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the kubeconfig files that are issued to users for the current cluster and the status of the kubeconfig files.
//
// Description:
//
// > - To call this operation, make sure that you have ram:ListUsers and ram:ListRoles permissions.
//
// > - To call this operation, make sure that you have the AliyunCSFullAccess permissions.
//
// @param request - ListClusterKubeconfigStatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterKubeconfigStatesResponse
func (client *Client) ListClusterKubeconfigStatesWithOptions(ClusterId *string, request *ListClusterKubeconfigStatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterKubeconfigStatesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the kubeconfig files that are issued to users for the current cluster and the status of the kubeconfig files.
//
// Description:
//
// > - To call this operation, make sure that you have ram:ListUsers and ram:ListRoles permissions.
//
// > - To call this operation, make sure that you have the AliyunCSFullAccess permissions.
//
// @param request - ListClusterKubeconfigStatesRequest
//
// @return ListClusterKubeconfigStatesResponse
func (client *Client) ListClusterKubeconfigStates(ClusterId *string, request *ListClusterKubeconfigStatesRequest) (_result *ListClusterKubeconfigStatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterKubeconfigStatesResponse{}
	_body, _err := client.ListClusterKubeconfigStatesWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the automated maintenance schedules of a cluster.
//
// @param request - ListOperationPlansRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationPlansResponse
func (client *Client) ListOperationPlansWithOptions(request *ListOperationPlansRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOperationPlansResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the automated maintenance schedules of a cluster.
//
// @param request - ListOperationPlansRequest
//
// @return ListOperationPlansResponse
func (client *Client) ListOperationPlans(request *ListOperationPlansRequest) (_result *ListOperationPlansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOperationPlansResponse{}
	_body, _err := client.ListOperationPlansWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个地域的自动运维执行计划列表
//
// @param request - ListOperationPlansForRegionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationPlansForRegionResponse
func (client *Client) ListOperationPlansForRegionWithOptions(regionId *string, request *ListOperationPlansForRegionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOperationPlansForRegionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个地域的自动运维执行计划列表
//
// @param request - ListOperationPlansForRegionRequest
//
// @return ListOperationPlansForRegionResponse
func (client *Client) ListOperationPlansForRegion(regionId *string, request *ListOperationPlansForRegionRequest) (_result *ListOperationPlansForRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOperationPlansForRegionResponse{}
	_body, _err := client.ListOperationPlansForRegionWithOptions(regionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries resource labels and the detailed information, such as the key-value pairs of the labels and the clusters to which the labels are added. You can use labels to classify and manage Container Service for Kubernetes (ACK) clusters in order to meet monitoring, cost analysis, and tenant isolation requirements.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource labels and the detailed information, such as the key-value pairs of the labels and the clusters to which the labels are added. You can use labels to classify and manage Container Service for Kubernetes (ACK) clusters in order to meet monitoring, cost analysis, and tenant isolation requirements.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ListUserKubeConfigStates operation to query the status of the kubeconfig files of all clusters managed by the current user.
//
// Description:
//
// >  To call this operation, make sure that you have the AliyunCSFullAccess permissions.
//
// @param request - ListUserKubeConfigStatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserKubeConfigStatesResponse
func (client *Client) ListUserKubeConfigStatesWithOptions(Uid *string, request *ListUserKubeConfigStatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserKubeConfigStatesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ListUserKubeConfigStates operation to query the status of the kubeconfig files of all clusters managed by the current user.
//
// Description:
//
// >  To call this operation, make sure that you have the AliyunCSFullAccess permissions.
//
// @param request - ListUserKubeConfigStatesRequest
//
// @return ListUserKubeConfigStatesResponse
func (client *Client) ListUserKubeConfigStates(Uid *string, request *ListUserKubeConfigStatesRequest) (_result *ListUserKubeConfigStatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserKubeConfigStatesResponse{}
	_body, _err := client.ListUserKubeConfigStatesWithOptions(Uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The Container Service for Kubernetes (ACK) managed Pro cluster type is developed based on the ACK managed Basic cluster type. It inherits all benefits of ACK managed clusters, such as fully-managed control planes and control plane high availability. It further enhances reliability, security, scheduling capabilities, and offers service level agreement (SLA)-backed guarantees, making it ideal for enterprise customers with large-scale production workloads requiring high stability and security. You can call the MigrateCluster operation to migrate an ACK managed Basic cluster to an ACK managed Pro cluster.
//
// @param request - MigrateClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateClusterResponse
func (client *Client) MigrateClusterWithOptions(clusterId *string, request *MigrateClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MigrateClusterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The Container Service for Kubernetes (ACK) managed Pro cluster type is developed based on the ACK managed Basic cluster type. It inherits all benefits of ACK managed clusters, such as fully-managed control planes and control plane high availability. It further enhances reliability, security, scheduling capabilities, and offers service level agreement (SLA)-backed guarantees, making it ideal for enterprise customers with large-scale production workloads requiring high stability and security. You can call the MigrateCluster operation to migrate an ACK managed Basic cluster to an ACK managed Pro cluster.
//
// @param request - MigrateClusterRequest
//
// @return MigrateClusterResponse
func (client *Client) MigrateCluster(clusterId *string, request *MigrateClusterRequest) (_result *MigrateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MigrateClusterResponse{}
	_body, _err := client.MigrateClusterWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ModifyCluster operation to modify the cluster configurations by cluster ID.
//
// @param request - ModifyClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterResponse
func (client *Client) ModifyClusterWithOptions(ClusterId *string, request *ModifyClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyClusterResponse, _err error) {
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

	if !dara.IsNil(request.ClusterName) {
		body["cluster_name"] = request.ClusterName
	}

	if !dara.IsNil(request.ControlPlaneConfig) {
		body["control_plane_config"] = request.ControlPlaneConfig
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyCluster operation to modify the cluster configurations by cluster ID.
//
// @param request - ModifyClusterRequest
//
// @return ModifyClusterResponse
func (client *Client) ModifyCluster(ClusterId *string, request *ModifyClusterRequest) (_result *ModifyClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterResponse{}
	_body, _err := client.ModifyClusterWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a cluster component. This operation may affect your businesses. We recommend that you assess the impact, back up data, and perform the operation during off-peak hours.
//
// Description:
//
// You can call this API operation to modify the component parameters of an ACK Basic cluster or the control plane parameters of an ACK Pro cluster:
//
//   - To view the component parameters of an ACK Basic cluster, call the DescribeClusterAddonMetadata API operation. For more information, see [Query the metadata of a cluster component](https://help.aliyun.com/document_detail/2667944.html).
//
//   - To view the control plane parameters of an ACK Pro cluster, see [Customize the control plane parameters of an ACK Pro cluster](https://help.aliyun.com/document_detail/199588.html).
//
// After you call this operation, the component may be redeployed and restarted. We recommend that you assess the impact before you call this operation.
//
// @param request - ModifyClusterAddonRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterAddonResponse
func (client *Client) ModifyClusterAddonWithOptions(clusterId *string, componentId *string, request *ModifyClusterAddonRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyClusterAddonResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a cluster component. This operation may affect your businesses. We recommend that you assess the impact, back up data, and perform the operation during off-peak hours.
//
// Description:
//
// You can call this API operation to modify the component parameters of an ACK Basic cluster or the control plane parameters of an ACK Pro cluster:
//
//   - To view the component parameters of an ACK Basic cluster, call the DescribeClusterAddonMetadata API operation. For more information, see [Query the metadata of a cluster component](https://help.aliyun.com/document_detail/2667944.html).
//
//   - To view the control plane parameters of an ACK Pro cluster, see [Customize the control plane parameters of an ACK Pro cluster](https://help.aliyun.com/document_detail/199588.html).
//
// After you call this operation, the component may be redeployed and restarted. We recommend that you assess the impact before you call this operation.
//
// @param request - ModifyClusterAddonRequest
//
// @return ModifyClusterAddonResponse
func (client *Client) ModifyClusterAddon(clusterId *string, componentId *string, request *ModifyClusterAddonRequest) (_result *ModifyClusterAddonResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterAddonResponse{}
	_body, _err := client.ModifyClusterAddonWithOptions(clusterId, componentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ModifyClusterNodePool operation to modify the configuration of a node pool with the specified node pool ID.
//
// @param request - ModifyClusterNodePoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterNodePoolResponse
func (client *Client) ModifyClusterNodePoolWithOptions(ClusterId *string, NodepoolId *string, request *ModifyClusterNodePoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyClusterNodePoolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyClusterNodePool operation to modify the configuration of a node pool with the specified node pool ID.
//
// @param request - ModifyClusterNodePoolRequest
//
// @return ModifyClusterNodePoolResponse
func (client *Client) ModifyClusterNodePool(ClusterId *string, NodepoolId *string, request *ModifyClusterNodePoolRequest) (_result *ModifyClusterNodePoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterNodePoolResponse{}
	_body, _err := client.ModifyClusterNodePoolWithOptions(ClusterId, NodepoolId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can add labels in key-value pairs to clusters. This allows cluster developers or O\\\\\\&M engineers to classify and manage clusters in a more flexible manner. This also meets the requirements for monitoring, cost analysis, and tenant isolation. You can call the ModifyClusterTags operation to modify the labels of a cluster.
//
// @param request - ModifyClusterTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterTagsResponse
func (client *Client) ModifyClusterTagsWithOptions(ClusterId *string, request *ModifyClusterTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyClusterTagsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can add labels in key-value pairs to clusters. This allows cluster developers or O\\\\\\&M engineers to classify and manage clusters in a more flexible manner. This also meets the requirements for monitoring, cost analysis, and tenant isolation. You can call the ModifyClusterTags operation to modify the labels of a cluster.
//
// @param request - ModifyClusterTagsRequest
//
// @return ModifyClusterTagsResponse
func (client *Client) ModifyClusterTags(ClusterId *string, request *ModifyClusterTagsRequest) (_result *ModifyClusterTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterTagsResponse{}
	_body, _err := client.ModifyClusterTagsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a node pool, such as the kubelet configuration and node rolling update configuration. After you modify the node pool configuration, nodes are batch updated and the kubelet on each node is restarted. This may adversely affect the nodes and workloads. We recommend that you perform this operation during off-peak hours.
//
// Description:
//
// >  Container Service for Kubernetes (ACK) allows you to modify the kubelet configuration of nodes in a node pool. After you modify the kubelet configuration, the new configuration immediately takes effect on existing nodes in the node pool and is automatically applied to newly added nodes.
//
// @param request - ModifyNodePoolNodeConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodePoolNodeConfigResponse
func (client *Client) ModifyNodePoolNodeConfigWithOptions(ClusterId *string, NodepoolId *string, request *ModifyNodePoolNodeConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyNodePoolNodeConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a node pool, such as the kubelet configuration and node rolling update configuration. After you modify the node pool configuration, nodes are batch updated and the kubelet on each node is restarted. This may adversely affect the nodes and workloads. We recommend that you perform this operation during off-peak hours.
//
// Description:
//
// >  Container Service for Kubernetes (ACK) allows you to modify the kubelet configuration of nodes in a node pool. After you modify the kubelet configuration, the new configuration immediately takes effect on existing nodes in the node pool and is automatically applied to newly added nodes.
//
// @param request - ModifyNodePoolNodeConfigRequest
//
// @return ModifyNodePoolNodeConfigResponse
func (client *Client) ModifyNodePoolNodeConfig(ClusterId *string, NodepoolId *string, request *ModifyNodePoolNodeConfigRequest) (_result *ModifyNodePoolNodeConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyNodePoolNodeConfigResponse{}
	_body, _err := client.ModifyNodePoolNodeConfigWithOptions(ClusterId, NodepoolId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a policy in a specific Container Service for Kubernetes (ACK) cluster. You can modify the action of the policy such as alerting or denying and namespaces to which the policy applies.
//
// @param request - ModifyPolicyInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyInstanceResponse
func (client *Client) ModifyPolicyInstanceWithOptions(clusterId *string, policyName *string, request *ModifyPolicyInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyPolicyInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a policy in a specific Container Service for Kubernetes (ACK) cluster. You can modify the action of the policy such as alerting or denying and namespaces to which the policy applies.
//
// @param request - ModifyPolicyInstanceRequest
//
// @return ModifyPolicyInstanceResponse
func (client *Client) ModifyPolicyInstance(clusterId *string, policyName *string, request *ModifyPolicyInstanceRequest) (_result *ModifyPolicyInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyPolicyInstanceResponse{}
	_body, _err := client.ModifyPolicyInstanceWithOptions(clusterId, policyName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// When using Container Service for Kubernetes (ACK) for the first time, you must call the OpenAckService operation to activate the service.
//
// Description:
//
//	  You can activate ACK by using Alibaba Cloud accounts.
//
//		- To activate ACK by using RAM users, you need to grant the AdministratorAccess permission to the RAM users.
//
// @param request - OpenAckServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenAckServiceResponse
func (client *Client) OpenAckServiceWithOptions(request *OpenAckServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenAckServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When using Container Service for Kubernetes (ACK) for the first time, you must call the OpenAckService operation to activate the service.
//
// Description:
//
//	  You can activate ACK by using Alibaba Cloud accounts.
//
//		- To activate ACK by using RAM users, you need to grant the AdministratorAccess permission to the RAM users.
//
// @param request - OpenAckServiceRequest
//
// @return OpenAckServiceResponse
func (client *Client) OpenAckService(request *OpenAckServiceRequest) (_result *OpenAckServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenAckServiceResponse{}
	_body, _err := client.OpenAckServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI PauseClusterUpgrade is deprecated
//
// Summary:
//
// You can call the PauseClusterUpgrade operation to pause the update of a Container Service for Kubernetes (ACK) cluster.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseClusterUpgradeResponse
func (client *Client) PauseClusterUpgradeWithOptions(ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PauseClusterUpgradeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the PauseClusterUpgrade operation to pause the update of a Container Service for Kubernetes (ACK) cluster.
//
// @return PauseClusterUpgradeResponse
// Deprecated
func (client *Client) PauseClusterUpgrade(ClusterId *string) (_result *PauseClusterUpgradeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PauseClusterUpgradeResponse{}
	_body, _err := client.PauseClusterUpgradeWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI PauseComponentUpgrade is deprecated
//
// Summary:
//
// You can call the PauseComponentUpgrade operation to pause the update of a component.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseComponentUpgradeResponse
func (client *Client) PauseComponentUpgradeWithOptions(clusterid *string, componentid *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PauseComponentUpgradeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the PauseComponentUpgrade operation to pause the update of a component.
//
// @return PauseComponentUpgradeResponse
// Deprecated
func (client *Client) PauseComponentUpgrade(clusterid *string, componentid *string) (_result *PauseComponentUpgradeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PauseComponentUpgradeResponse{}
	_body, _err := client.PauseComponentUpgradeWithOptions(clusterid, componentid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pauses an on-going task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseTaskResponse
func (client *Client) PauseTaskWithOptions(taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PauseTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses an on-going task.
//
// @return PauseTaskResponse
func (client *Client) PauseTask(taskId *string) (_result *PauseTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PauseTaskResponse{}
	_body, _err := client.PauseTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI RemoveClusterNodes is deprecated
//
// Summary:
//
// You can call the RemoveClusterNodes operation to remove nodes from a Container Service for Kubernetes (ACK) cluster.
//
// Description:
//
// ***
//
//   - When you remove a node, the pods that run on the node are migrated to other nodes. This may cause service interruptions. We recommend that you remove nodes during off-peak hours.
//
//   - Unknown errors may occur when you remove nodes. Before you remove nodes, back up the data on the nodes.
//
//   - Nodes remain in the Unschedulable state when they are being removed.
//
//   - You can remove only worker nodes. You cannot remove master nodes.
//
// @param request - RemoveClusterNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveClusterNodesResponse
func (client *Client) RemoveClusterNodesWithOptions(ClusterId *string, request *RemoveClusterNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveClusterNodesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the RemoveClusterNodes operation to remove nodes from a Container Service for Kubernetes (ACK) cluster.
//
// Description:
//
// ***
//
//   - When you remove a node, the pods that run on the node are migrated to other nodes. This may cause service interruptions. We recommend that you remove nodes during off-peak hours.
//
//   - Unknown errors may occur when you remove nodes. Before you remove nodes, back up the data on the nodes.
//
//   - Nodes remain in the Unschedulable state when they are being removed.
//
//   - You can remove only worker nodes. You cannot remove master nodes.
//
// @param request - RemoveClusterNodesRequest
//
// @return RemoveClusterNodesResponse
// Deprecated
func (client *Client) RemoveClusterNodes(ClusterId *string, request *RemoveClusterNodesRequest) (_result *RemoveClusterNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveClusterNodesResponse{}
	_body, _err := client.RemoveClusterNodesWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes nodes from a node pool.
//
// Description:
//
//	  When you remove a node, the pods on the node are migrated to other nodes. This may cause service interruptions. We recommend that you remove nodes during off-peak hours.
//
//		- The operation may have unexpected risks. Back up the data before you perform this operation.
//
//		- Nodes remain in the Unschedulable state when they are being removed.
//
//		- The system removes only worker nodes. It does not remove master nodes.
//
//		- Even if you set the `release_node` parameter to `true`, subscription nodes are not released. You must release the subscription nodes in the [ECS console](https://ecs.console.aliyun.com/) after you remove the nodes.
//
// @param tmpReq - RemoveNodePoolNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveNodePoolNodesResponse
func (client *Client) RemoveNodePoolNodesWithOptions(ClusterId *string, NodepoolId *string, tmpReq *RemoveNodePoolNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveNodePoolNodesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes nodes from a node pool.
//
// Description:
//
//	  When you remove a node, the pods on the node are migrated to other nodes. This may cause service interruptions. We recommend that you remove nodes during off-peak hours.
//
//		- The operation may have unexpected risks. Back up the data before you perform this operation.
//
//		- Nodes remain in the Unschedulable state when they are being removed.
//
//		- The system removes only worker nodes. It does not remove master nodes.
//
//		- Even if you set the `release_node` parameter to `true`, subscription nodes are not released. You must release the subscription nodes in the [ECS console](https://ecs.console.aliyun.com/) after you remove the nodes.
//
// @param request - RemoveNodePoolNodesRequest
//
// @return RemoveNodePoolNodesResponse
func (client *Client) RemoveNodePoolNodes(ClusterId *string, NodepoolId *string, request *RemoveNodePoolNodesRequest) (_result *RemoveNodePoolNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveNodePoolNodesResponse{}
	_body, _err := client.RemoveNodePoolNodesWithOptions(ClusterId, NodepoolId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Repairs a node pool.
//
// @param request - RepairClusterNodePoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RepairClusterNodePoolResponse
func (client *Client) RepairClusterNodePoolWithOptions(clusterId *string, nodepoolId *string, request *RepairClusterNodePoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RepairClusterNodePoolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Repairs a node pool.
//
// @param request - RepairClusterNodePoolRequest
//
// @return RepairClusterNodePoolResponse
func (client *Client) RepairClusterNodePool(clusterId *string, nodepoolId *string, request *RepairClusterNodePoolRequest) (_result *RepairClusterNodePoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RepairClusterNodePoolResponse{}
	_body, _err := client.RepairClusterNodePoolWithOptions(clusterId, nodepoolId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ResumeComponentUpgrade is deprecated
//
// Summary:
//
// You can call the ResumeComponentUpgrade operation to resume the update of a component.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeComponentUpgradeResponse
func (client *Client) ResumeComponentUpgradeWithOptions(clusterid *string, componentid *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeComponentUpgradeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the ResumeComponentUpgrade operation to resume the update of a component.
//
// @return ResumeComponentUpgradeResponse
// Deprecated
func (client *Client) ResumeComponentUpgrade(clusterid *string, componentid *string) (_result *ResumeComponentUpgradeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeComponentUpgradeResponse{}
	_body, _err := client.ResumeComponentUpgradeWithOptions(clusterid, componentid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes a task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeTaskResponse
func (client *Client) ResumeTaskWithOptions(taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes a task.
//
// @return ResumeTaskResponse
func (client *Client) ResumeTask(taskId *string) (_result *ResumeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeTaskResponse{}
	_body, _err := client.ResumeTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ResumeUpgradeCluster is deprecated
//
// Summary:
//
// You can call the ResumeUpgradeCluster operation to resume the update of a cluster by cluster ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeUpgradeClusterResponse
func (client *Client) ResumeUpgradeClusterWithOptions(ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeUpgradeClusterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can call the ResumeUpgradeCluster operation to resume the update of a cluster by cluster ID.
//
// @return ResumeUpgradeClusterResponse
// Deprecated
func (client *Client) ResumeUpgradeCluster(ClusterId *string) (_result *ResumeUpgradeClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeUpgradeClusterResponse{}
	_body, _err := client.ResumeUpgradeClusterWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the RevokeK8sClusterKubeConfig operation to revoke the kubeconfig file of a cluster that belongs to the current Alibaba Cloud account or RAM user. After the kubeconfig file is revoked, the cluster generates a new kubeconfig file, and the original kubeconfig file becomes invalid.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeK8sClusterKubeConfigResponse
func (client *Client) RevokeK8sClusterKubeConfigWithOptions(ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeK8sClusterKubeConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the RevokeK8sClusterKubeConfig operation to revoke the kubeconfig file of a cluster that belongs to the current Alibaba Cloud account or RAM user. After the kubeconfig file is revoked, the cluster generates a new kubeconfig file, and the original kubeconfig file becomes invalid.
//
// @return RevokeK8sClusterKubeConfigResponse
func (client *Client) RevokeK8sClusterKubeConfig(ClusterId *string) (_result *RevokeK8sClusterKubeConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeK8sClusterKubeConfigResponse{}
	_body, _err := client.RevokeK8sClusterKubeConfigWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Container Intelligence Service (CIS) provides a variety of cluster check capabilities to allow you to perform cluster update check, cluster migration check, component installation check, component update check, and node pool check. A precheck is automatically triggered before an update, migration, or installation is performed. You can perform changes only if the cluster passes the precheck. You can also manually call the RunClusterCheck operation to initiate cluster checks. We recommend that you periodically check and maintain your cluster to mitigate potential risks.
//
// @param request - RunClusterCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunClusterCheckResponse
func (client *Client) RunClusterCheckWithOptions(clusterId *string, request *RunClusterCheckRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunClusterCheckResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Container Intelligence Service (CIS) provides a variety of cluster check capabilities to allow you to perform cluster update check, cluster migration check, component installation check, component update check, and node pool check. A precheck is automatically triggered before an update, migration, or installation is performed. You can perform changes only if the cluster passes the precheck. You can also manually call the RunClusterCheck operation to initiate cluster checks. We recommend that you periodically check and maintain your cluster to mitigate potential risks.
//
// @param request - RunClusterCheckRequest
//
// @return RunClusterCheckResponse
func (client *Client) RunClusterCheck(clusterId *string, request *RunClusterCheckRequest) (_result *RunClusterCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunClusterCheckResponse{}
	_body, _err := client.RunClusterCheckWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Triggers a cluster inspection and generates a report.
//
// @param request - RunClusterInspectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunClusterInspectResponse
func (client *Client) RunClusterInspectWithOptions(clusterId *string, request *RunClusterInspectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunClusterInspectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers a cluster inspection and generates a report.
//
// @param request - RunClusterInspectRequest
//
// @return RunClusterInspectResponse
func (client *Client) RunClusterInspect(clusterId *string, request *RunClusterInspectRequest) (_result *RunClusterInspectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunClusterInspectResponse{}
	_body, _err := client.RunClusterInspectWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scales out a node pool.
//
// @param request - ScaleClusterNodePoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleClusterNodePoolResponse
func (client *Client) ScaleClusterNodePoolWithOptions(ClusterId *string, NodepoolId *string, request *ScaleClusterNodePoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleClusterNodePoolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales out a node pool.
//
// @param request - ScaleClusterNodePoolRequest
//
// @return ScaleClusterNodePoolResponse
func (client *Client) ScaleClusterNodePool(ClusterId *string, NodepoolId *string, request *ScaleClusterNodePoolRequest) (_result *ScaleClusterNodePoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleClusterNodePoolResponse{}
	_body, _err := client.ScaleClusterNodePoolWithOptions(ClusterId, NodepoolId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ScaleOutCluster operation to scale out a cluster by cluster ID.
//
// Description:
//
// *
//
// ****The ScaleOutCluster API operation is phased out. You must call the node pool-related API operations to manage nodes. If you want to add worker nodes to a Container Service for Kubernetes (ACK) cluster, call the ScaleClusterNodePool API operation. For more information, see [ScaleClusterNodePool](https://help.aliyun.com/document_detail/184928.html).
//
// @param request - ScaleOutClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleOutClusterResponse
func (client *Client) ScaleOutClusterWithOptions(ClusterId *string, request *ScaleOutClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleOutClusterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ScaleOutCluster operation to scale out a cluster by cluster ID.
//
// Description:
//
// *
//
// ****The ScaleOutCluster API operation is phased out. You must call the node pool-related API operations to manage nodes. If you want to add worker nodes to a Container Service for Kubernetes (ACK) cluster, call the ScaleClusterNodePool API operation. For more information, see [ScaleClusterNodePool](https://help.aliyun.com/document_detail/184928.html).
//
// @param request - ScaleOutClusterRequest
//
// @return ScaleOutClusterResponse
func (client *Client) ScaleOutCluster(ClusterId *string, request *ScaleOutClusterRequest) (_result *ScaleOutClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleOutClusterResponse{}
	_body, _err := client.ScaleOutClusterWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scans for vulnerabilities in a Container Service for Kubernetes (ACK) cluster, including workload vulnerabilities, third-party software vulnerabilities, CVE vulnerabilities, WebCMS vulnerabilities, and Windows vulnerabilities. We recommend that you scan your cluster on a regular basis to ensure cluster security.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScanClusterVulsResponse
func (client *Client) ScanClusterVulsWithOptions(clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScanClusterVulsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scans for vulnerabilities in a Container Service for Kubernetes (ACK) cluster, including workload vulnerabilities, third-party software vulnerabilities, CVE vulnerabilities, WebCMS vulnerabilities, and Windows vulnerabilities. We recommend that you scan your cluster on a regular basis to ensure cluster security.
//
// @return ScanClusterVulsResponse
func (client *Client) ScanClusterVuls(clusterId *string) (_result *ScanClusterVulsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScanClusterVulsResponse{}
	_body, _err := client.ScanClusterVulsWithOptions(clusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates the specified alert rule(s).
//
// @param request - StartAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAlertResponse
func (client *Client) StartAlertWithOptions(ClusterId *string, request *StartAlertRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartAlertResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates the specified alert rule(s).
//
// @param request - StartAlertRequest
//
// @return StartAlertResponse
func (client *Client) StartAlert(ClusterId *string, request *StartAlertRequest) (_result *StartAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartAlertResponse{}
	_body, _err := client.StartAlertWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the StopAlert operation to disable an alert rule or an alert rule set in the alert center of Container Service for Kubernetes (ACK).
//
// @param request - StopAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAlertResponse
func (client *Client) StopAlertWithOptions(ClusterId *string, request *StopAlertRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopAlertResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the StopAlert operation to disable an alert rule or an alert rule set in the alert center of Container Service for Kubernetes (ACK).
//
// @param request - StopAlertRequest
//
// @return StopAlertResponse
func (client *Client) StopAlert(ClusterId *string, request *StopAlertRequest) (_result *StopAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopAlertResponse{}
	_body, _err := client.StopAlertWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes the information about a node pool, including the metadata and node information of the node pool.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncClusterNodePoolResponse
func (client *Client) SyncClusterNodePoolWithOptions(ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SyncClusterNodePoolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes the information about a node pool, including the metadata and node information of the node pool.
//
// @return SyncClusterNodePoolResponse
func (client *Client) SyncClusterNodePool(ClusterId *string) (_result *SyncClusterNodePoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SyncClusterNodePoolResponse{}
	_body, _err := client.SyncClusterNodePoolWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can add labels in key-value pairs to clusters. This allows cluster developers or O\\&M engineers to classify and manage clusters in a more flexible manner. This also meets the requirements for monitoring, cost analysis, and tenant isolation. You can call the TagResources operation to add labels to a cluster.
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can add labels in key-value pairs to clusters. This allows cluster developers or O\\&M engineers to classify and manage clusters in a more flexible manner. This also meets the requirements for monitoring, cost analysis, and tenant isolation. You can call the TagResources operation to add labels to a cluster.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls components that you no longer need from a cluster. You must specify the name of the components and specify whether to release associated Alibaba Cloud resources from the cluster.
//
// @param request - UnInstallClusterAddonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnInstallClusterAddonsResponse
func (client *Client) UnInstallClusterAddonsWithOptions(ClusterId *string, request *UnInstallClusterAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnInstallClusterAddonsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls components that you no longer need from a cluster. You must specify the name of the components and specify whether to release associated Alibaba Cloud resources from the cluster.
//
// @param request - UnInstallClusterAddonsRequest
//
// @return UnInstallClusterAddonsResponse
func (client *Client) UnInstallClusterAddons(ClusterId *string, request *UnInstallClusterAddonsRequest) (_result *UnInstallClusterAddonsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnInstallClusterAddonsResponse{}
	_body, _err := client.UnInstallClusterAddonsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If you no longer need the labels (key-value pairs) of a cluster, you can call the UntagResources operation to delete the labels.
//
// @param tmpReq - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(tmpReq *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If you no longer need the labels (key-value pairs) of a cluster, you can call the UntagResources operation to delete the labels.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the UpdateClusterAuditLogConfig operation to enable or disable the audit log feature in a Container Service for Kubernetes (ACK) cluster and update the audit log configuration. This operation also allows you to record requests to the Kubernetes API and the responses, which can be used to trace cluster operation history and troubleshoot cluster issues.
//
// Description:
//
// Before you call this operation, ensure that you understand the billing methods and pricing of [Simple Log Service](https://www.alibabacloud.com/product/log-service/pricing).
//
// @param request - UpdateClusterAuditLogConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterAuditLogConfigResponse
func (client *Client) UpdateClusterAuditLogConfigWithOptions(clusterid *string, request *UpdateClusterAuditLogConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateClusterAuditLogConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the UpdateClusterAuditLogConfig operation to enable or disable the audit log feature in a Container Service for Kubernetes (ACK) cluster and update the audit log configuration. This operation also allows you to record requests to the Kubernetes API and the responses, which can be used to trace cluster operation history and troubleshoot cluster issues.
//
// Description:
//
// Before you call this operation, ensure that you understand the billing methods and pricing of [Simple Log Service](https://www.alibabacloud.com/product/log-service/pricing).
//
// @param request - UpdateClusterAuditLogConfigRequest
//
// @return UpdateClusterAuditLogConfigResponse
func (client *Client) UpdateClusterAuditLogConfig(clusterid *string, request *UpdateClusterAuditLogConfigRequest) (_result *UpdateClusterAuditLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateClusterAuditLogConfigResponse{}
	_body, _err := client.UpdateClusterAuditLogConfigWithOptions(clusterid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies cluster inspection configurations.
//
// @param request - UpdateClusterInspectConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterInspectConfigResponse
func (client *Client) UpdateClusterInspectConfigWithOptions(clusterId *string, request *UpdateClusterInspectConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateClusterInspectConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies cluster inspection configurations.
//
// @param request - UpdateClusterInspectConfigRequest
//
// @return UpdateClusterInspectConfigResponse
func (client *Client) UpdateClusterInspectConfig(clusterId *string, request *UpdateClusterInspectConfigRequest) (_result *UpdateClusterInspectConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateClusterInspectConfigResponse{}
	_body, _err := client.UpdateClusterInspectConfigWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the UpdateContactGroupForAlert operation to specify a contact group for an alert rule in an ACK cluster.
//
// @param request - UpdateContactGroupForAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContactGroupForAlertResponse
func (client *Client) UpdateContactGroupForAlertWithOptions(ClusterId *string, request *UpdateContactGroupForAlertRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateContactGroupForAlertResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the UpdateContactGroupForAlert operation to specify a contact group for an alert rule in an ACK cluster.
//
// @param request - UpdateContactGroupForAlertRequest
//
// @return UpdateContactGroupForAlertResponse
func (client *Client) UpdateContactGroupForAlert(ClusterId *string, request *UpdateContactGroupForAlertRequest) (_result *UpdateContactGroupForAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateContactGroupForAlertResponse{}
	_body, _err := client.UpdateContactGroupForAlertWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the log configurations of control plane components. The configurations include the log retention period and components whose logs that you want to collect. Container Service for Kubernetes (ACK) managed clusters can collect the logs of control plane components and deliver the logs to projects in Simple Log Service. These control plane components include Kube-apiserver, kube-scheduler, Kubernetes controller manager, and cloud controller manager (CCM).
//
// @param request - UpdateControlPlaneLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateControlPlaneLogResponse
func (client *Client) UpdateControlPlaneLogWithOptions(ClusterId *string, request *UpdateControlPlaneLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateControlPlaneLogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the log configurations of control plane components. The configurations include the log retention period and components whose logs that you want to collect. Container Service for Kubernetes (ACK) managed clusters can collect the logs of control plane components and deliver the logs to projects in Simple Log Service. These control plane components include Kube-apiserver, kube-scheduler, Kubernetes controller manager, and cloud controller manager (CCM).
//
// @param request - UpdateControlPlaneLogRequest
//
// @return UpdateControlPlaneLogResponse
func (client *Client) UpdateControlPlaneLog(ClusterId *string, request *UpdateControlPlaneLogRequest) (_result *UpdateControlPlaneLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateControlPlaneLogResponse{}
	_body, _err := client.UpdateControlPlaneLogWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the validity period of a kubeconfig file used by a Resource Access Management (RAM) user or RAM role to connect to a Container Service for Kubernetes (ACK) cluster. The validity period ranges from 1 to 876,000 hours. You can call this API operation when you customize configurations by using an Alibaba Cloud account. The default validity period of a kubeconfig file is three years.
//
// Description:
//
//	  You can call this operation only with an Alibaba Cloud account.
//
//		- If the kubeconfig file used by your cluster is revoked, the custom validity period of the kubeconfig file is reset. In this case, you need to call this API operation to reconfigure the validity period of the kubeconfig file.
//
// @param request - UpdateK8sClusterUserConfigExpireRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateK8sClusterUserConfigExpireResponse
func (client *Client) UpdateK8sClusterUserConfigExpireWithOptions(ClusterId *string, request *UpdateK8sClusterUserConfigExpireRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateK8sClusterUserConfigExpireResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the validity period of a kubeconfig file used by a Resource Access Management (RAM) user or RAM role to connect to a Container Service for Kubernetes (ACK) cluster. The validity period ranges from 1 to 876,000 hours. You can call this API operation when you customize configurations by using an Alibaba Cloud account. The default validity period of a kubeconfig file is three years.
//
// Description:
//
//	  You can call this operation only with an Alibaba Cloud account.
//
//		- If the kubeconfig file used by your cluster is revoked, the custom validity period of the kubeconfig file is reset. In this case, you need to call this API operation to reconfigure the validity period of the kubeconfig file.
//
// @param request - UpdateK8sClusterUserConfigExpireRequest
//
// @return UpdateK8sClusterUserConfigExpireResponse
func (client *Client) UpdateK8sClusterUserConfigExpire(ClusterId *string, request *UpdateK8sClusterUserConfigExpireRequest) (_result *UpdateK8sClusterUserConfigExpireResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateK8sClusterUserConfigExpireResponse{}
	_body, _err := client.UpdateK8sClusterUserConfigExpireWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新节点组件
//
// @param request - UpdateNodePoolComponentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodePoolComponentResponse
func (client *Client) UpdateNodePoolComponentWithOptions(clusterId *string, nodepoolId *string, request *UpdateNodePoolComponentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateNodePoolComponentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新节点组件
//
// @param request - UpdateNodePoolComponentRequest
//
// @return UpdateNodePoolComponentResponse
func (client *Client) UpdateNodePoolComponent(clusterId *string, nodepoolId *string, request *UpdateNodePoolComponentRequest) (_result *UpdateNodePoolComponentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateNodePoolComponentResponse{}
	_body, _err := client.UpdateNodePoolComponentWithOptions(clusterId, nodepoolId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the deletion protection status of the specified resources. You can enable or disable deletion protection for namespaces and Services. You can call this operation to enable deletion protection for namespaces or Services that involve businesses-critical and sensitive data to avoid incurring maintenance costs caused by accidental namespace or Service deletion.
//
// @param request - UpdateResourcesDeleteProtectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourcesDeleteProtectionResponse
func (client *Client) UpdateResourcesDeleteProtectionWithOptions(ClusterId *string, request *UpdateResourcesDeleteProtectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourcesDeleteProtectionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the deletion protection status of the specified resources. You can enable or disable deletion protection for namespaces and Services. You can call this operation to enable deletion protection for namespaces or Services that involve businesses-critical and sensitive data to avoid incurring maintenance costs caused by accidental namespace or Service deletion.
//
// @param request - UpdateResourcesDeleteProtectionRequest
//
// @return UpdateResourcesDeleteProtectionResponse
func (client *Client) UpdateResourcesDeleteProtection(ClusterId *string, request *UpdateResourcesDeleteProtectionRequest) (_result *UpdateResourcesDeleteProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourcesDeleteProtectionResponse{}
	_body, _err := client.UpdateResourcesDeleteProtectionWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of an orchestration template. An orchestration template defines and describes a group of Container Service for Kubernetes (ACK) resources. An orchestration template describes the configurations of an application or how an application runs in a declarative manner.
//
// @param request - UpdateTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithOptions(TemplateId *string, request *UpdateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of an orchestration template. An orchestration template defines and describes a group of Container Service for Kubernetes (ACK) resources. An orchestration template describes the configurations of an application or how an application runs in a declarative manner.
//
// @param request - UpdateTemplateRequest
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplate(TemplateId *string, request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(TemplateId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the role-based access control (RBAC) permissions of a Resource Access Management (RAM) user or RAM role. By default, you do not have the RBAC permissions on a Container Service for Kubernetes (ACK) cluster if you are not the cluster owner or you are not using an Alibaba Cloud account. You can call this operation to specify the resources that can be accessed, permission scope, and predefined roles. This helps you better manage the access control on resources in ACK clusters.
//
// Description:
//
// *Precautions**:
//
//   - You can update the permissions of a RAM user or RAM role on a cluster by using full update or incremental update. If you use full update, the existing permissions of the RAM user or RAM role on the cluster are overwritten. You must specify all the permissions that you want to grant to the RAM user or RAM role in the request parameters when you call the operation. If you use incremental update, you can grant permissions to or revoke permissions from the RAM user or RAM role on the cluster. In this case, only the permissions that you specify in the request parameters when you call the operation are granted or revoked, other permissions of the RAM user or RAM role on the cluster are not affected.
//
// @param request - UpdateUserPermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserPermissionsResponse
func (client *Client) UpdateUserPermissionsWithOptions(uid *string, request *UpdateUserPermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUserPermissionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the role-based access control (RBAC) permissions of a Resource Access Management (RAM) user or RAM role. By default, you do not have the RBAC permissions on a Container Service for Kubernetes (ACK) cluster if you are not the cluster owner or you are not using an Alibaba Cloud account. You can call this operation to specify the resources that can be accessed, permission scope, and predefined roles. This helps you better manage the access control on resources in ACK clusters.
//
// Description:
//
// *Precautions**:
//
//   - You can update the permissions of a RAM user or RAM role on a cluster by using full update or incremental update. If you use full update, the existing permissions of the RAM user or RAM role on the cluster are overwritten. You must specify all the permissions that you want to grant to the RAM user or RAM role in the request parameters when you call the operation. If you use incremental update, you can grant permissions to or revoke permissions from the RAM user or RAM role on the cluster. In this case, only the permissions that you specify in the request parameters when you call the operation are granted or revoked, other permissions of the RAM user or RAM role on the cluster are not affected.
//
// @param request - UpdateUserPermissionsRequest
//
// @return UpdateUserPermissionsResponse
func (client *Client) UpdateUserPermissions(uid *string, request *UpdateUserPermissionsRequest) (_result *UpdateUserPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateUserPermissionsResponse{}
	_body, _err := client.UpdateUserPermissionsWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Outdated Kubernetes versions may have security and stability issues. We recommend that you update the Kubernetes version of your cluster at the earliest opportunity to enjoy the new features of the new Kubernetes version. You can call the UpgradeCluster operation to manually upgrade a cluster.
//
// Description:
//
// After successfully calling the UpgradeCluster interface, this API returns the `task_id` of the upgrade task. You can manage this operation task by calling the following task APIs:
//
// - [Call DescribeTaskInfo to query task details](https://help.aliyun.com/document_detail/2667985.html)
//
// - [Call PauseTask to pause a running task](https://help.aliyun.com/document_detail/2667986.html)
//
// - [Call ResumeTask to resume a task that has been paused](https://help.aliyun.com/document_detail/2667987.html)
//
// - [Call CancelTask to cancel a running task](https://help.aliyun.com/document_detail/2667988.html)
//
// @param request - UpgradeClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeClusterResponse
func (client *Client) UpgradeClusterWithOptions(ClusterId *string, request *UpgradeClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeClusterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Outdated Kubernetes versions may have security and stability issues. We recommend that you update the Kubernetes version of your cluster at the earliest opportunity to enjoy the new features of the new Kubernetes version. You can call the UpgradeCluster operation to manually upgrade a cluster.
//
// Description:
//
// After successfully calling the UpgradeCluster interface, this API returns the `task_id` of the upgrade task. You can manage this operation task by calling the following task APIs:
//
// - [Call DescribeTaskInfo to query task details](https://help.aliyun.com/document_detail/2667985.html)
//
// - [Call PauseTask to pause a running task](https://help.aliyun.com/document_detail/2667986.html)
//
// - [Call ResumeTask to resume a task that has been paused](https://help.aliyun.com/document_detail/2667987.html)
//
// - [Call CancelTask to cancel a running task](https://help.aliyun.com/document_detail/2667988.html)
//
// @param request - UpgradeClusterRequest
//
// @return UpgradeClusterResponse
func (client *Client) UpgradeCluster(ClusterId *string, request *UpgradeClusterRequest) (_result *UpgradeClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeClusterResponse{}
	_body, _err := client.UpgradeClusterWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates cluster components to use new features and patch vulnerabilities. You must update cluster components one after one and update a component only after the previous one is successfully updated. Before you update a component, we recommend that you read the update notes for each component. Cluster component updates may affect your businesses. Assess the impact, back up data, and perform the update during off-peak hours.
//
// @param request - UpgradeClusterAddonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeClusterAddonsResponse
func (client *Client) UpgradeClusterAddonsWithOptions(ClusterId *string, request *UpgradeClusterAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeClusterAddonsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates cluster components to use new features and patch vulnerabilities. You must update cluster components one after one and update a component only after the previous one is successfully updated. Before you update a component, we recommend that you read the update notes for each component. Cluster component updates may affect your businesses. Assess the impact, back up data, and perform the update during off-peak hours.
//
// @param request - UpgradeClusterAddonsRequest
//
// @return UpgradeClusterAddonsResponse
func (client *Client) UpgradeClusterAddons(ClusterId *string, request *UpgradeClusterAddonsRequest) (_result *UpgradeClusterAddonsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeClusterAddonsResponse{}
	_body, _err := client.UpgradeClusterAddonsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the UpgradeClusterNodepool operation to update the Kubernetes version, OS version, or container runtime version of the nodes in a node pool.
//
// Description:
//
// This operation allows you to update the Kubernetes version, OS version, or container runtime version of the nodes in a node pool.
//
// @param request - UpgradeClusterNodepoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeClusterNodepoolResponse
func (client *Client) UpgradeClusterNodepoolWithOptions(ClusterId *string, NodepoolId *string, request *UpgradeClusterNodepoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeClusterNodepoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the UpgradeClusterNodepool operation to update the Kubernetes version, OS version, or container runtime version of the nodes in a node pool.
//
// Description:
//
// This operation allows you to update the Kubernetes version, OS version, or container runtime version of the nodes in a node pool.
//
// @param request - UpgradeClusterNodepoolRequest
//
// @return UpgradeClusterNodepoolResponse
func (client *Client) UpgradeClusterNodepool(ClusterId *string, NodepoolId *string, request *UpgradeClusterNodepoolRequest) (_result *UpgradeClusterNodepoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeClusterNodepoolResponse{}
	_body, _err := client.UpgradeClusterNodepoolWithOptions(ClusterId, NodepoolId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
