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
		"us-west-1":                   dara.String("cs.us-west-1.aliyuncs.com"),
		"us-southeast-1":              dara.String("cs.us-southeast-1.aliyuncs.com"),
		"us-east-1":                   dara.String("cs.us-east-1.aliyuncs.com"),
		"na-south-1":                  dara.String("cs.na-south-1.aliyuncs.com"),
		"me-east-1":                   dara.String("cs.me-east-1.aliyuncs.com"),
		"me-central-1":                dara.String("cs.me-central-1.aliyuncs.com"),
		"eu-west-1":                   dara.String("cs.eu-west-1.aliyuncs.com"),
		"eu-central-1":                dara.String("cs.eu-central-1.aliyuncs.com"),
		"cn-zhengzhou-jva":            dara.String("cs.cn-zhengzhou-jva.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("cs.cn-zhangjiakou.aliyuncs.com"),
		"cn-wulanchabu-gic-1":         dara.String("cs.cn-wulanchabu-gic-1.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("cs.cn-wulanchabu.aliyuncs.com"),
		"cn-wuhan-lr":                 dara.String("cs.cn-wuhan-lr.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("cs.cn-shenzhen-finance-1.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("cs.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("cs.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-shanghai":                 dara.String("cs.cn-shanghai.aliyuncs.com"),
		"cn-qingdao":                  dara.String("cs.cn-qingdao.aliyuncs.com"),
		"cn-nanjing":                  dara.String("cs.cn-nanjing.aliyuncs.com"),
		"cn-huhehaote":                dara.String("cs.cn-huhehaote.aliyuncs.com"),
		"cn-hongkong":                 dara.String("cs.cn-hongkong.aliyuncs.com"),
		"cn-heyuan-acdr-1":            dara.String("cs.cn-heyuan-acdr-1.aliyuncs.com"),
		"cn-heyuan":                   dara.String("cs.cn-heyuan.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("cs.cn-hangzhou-finance.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("cs.cn-hangzhou.aliyuncs.com"),
		"cn-guangzhou":                dara.String("cs.cn-guangzhou.aliyuncs.com"),
		"cn-fuzhou":                   dara.String("cs.cn-fuzhou.aliyuncs.com"),
		"cn-chengdu":                  dara.String("cs.cn-chengdu.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("cs.cn-beijing-finance-1.aliyuncs.com"),
		"cn-beijing":                  dara.String("cs.cn-beijing.aliyuncs.com"),
		"ap-southeast-7":              dara.String("cs.ap-southeast-7.aliyuncs.com"),
		"ap-southeast-5":              dara.String("cs.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-3":              dara.String("cs.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-2":              dara.String("cs.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-1":              dara.String("cs.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":                  dara.String("cs.ap-south-1.aliyuncs.com"),
		"ap-northeast-2":              dara.String("cs.ap-northeast-2.aliyuncs.com"),
		"ap-northeast-1":              dara.String("cs.ap-northeast-1.aliyuncs.com"),
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
// Adds existing Elastic Compute Service (ECS) instances to a Container Service for Kubernetes (ACK) cluster. You can call the AttachInstances operation to add ECS instances to an ACK cluster as worker nodes after purchasing the instances, or to re-add node instances to a node pool after they are removed.
//
// Description:
//
// Before you invoke this operation, read [Limits](https://help.aliyun.com/document_detail/86919.html) to make sure that the ECS instances to be added meet the requirements.
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
// Adds existing Elastic Compute Service (ECS) instances to a Container Service for Kubernetes (ACK) cluster. You can call the AttachInstances operation to add ECS instances to an ACK cluster as worker nodes after purchasing the instances, or to re-add node instances to a node pool after they are removed.
//
// Description:
//
// Before you invoke this operation, read [Limits](https://help.aliyun.com/document_detail/86919.html) to make sure that the ECS instances to be added meet the requirements.
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
// Adds existing ECS instances to a node pool in an ACK cluster. You can call this operation to add existing ECS instances as worker nodes to an ACK cluster or to re-add node instances to a node pool after they have been removed.
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
// Adds existing ECS instances to a node pool in an ACK cluster. You can call this operation to add existing ECS instances as worker nodes to an ACK cluster or to re-add node instances to a node pool after they have been removed.
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
// Cancels the upgrade of an ACK cluster that is in the upgrading state.
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
// Cancels the upgrade of an ACK cluster that is in the upgrading state.
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
// Cancels the upgrade of a cluster component.
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
// Cancels the upgrade of a cluster component.
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
// Cancels an existing but unexecuted automated O&M task execution plan by calling CancelOperationPlan.
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
// Cancels an existing but unexecuted automated O&M task execution plan by calling CancelOperationPlan.
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
// Queries the log configuration of control plane components for an ACK managed cluster. Control plane components include Kube API Server, Kube Scheduler, Kube Controller Manager, and Cloud Controller Manager. You can call the CheckControlPlaneLogEnable operation to query the current log configuration of control plane components, including the log retention period and the components from which logs are collected.
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
// Queries the log configuration of control plane components for an ACK managed cluster. Control plane components include Kube API Server, Kube Scheduler, Kube Controller Manager, and Cloud Controller Manager. You can call the CheckControlPlaneLogEnable operation to query the current log configuration of control plane components, including the log retention period and the components from which logs are collected.
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
// Checks whether the current service account has been granted the specified service role permissions. Container Service for Kubernetes (ACK) can call other cloud services (such as ECS, OSS, NAS, and SLB) that are associated with service roles only after the corresponding role permissions are granted.
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
// Checks whether the current service account has been granted the specified service role permissions. Container Service for Kubernetes (ACK) can call other cloud services (such as ECS, OSS, NAS, and SLB) that are associated with service roles only after the corresponding role permissions are granted.
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
// Cleans up KubeConfig credentials and RBAC permissions that have been issued to a specified user in a specified cluster. You can call this operation to revoke authorization for KubeConfig credentials that pose security risks.
//
// Description:
//
// > 1. You must have the permission to manage Container Service for Kubernetes (ACK) (AliyunCSFullAccess).
//
// > 2. Cleaning up permissions of an Alibaba Cloud account is not supported.
//
// > 3. Cleaning up the permissions of the user who performs this operation is not supported.
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
// Cleans up KubeConfig credentials and RBAC permissions that have been issued to a specified user in a specified cluster. You can call this operation to revoke authorization for KubeConfig credentials that pose security risks.
//
// Description:
//
// > 1. You must have the permission to manage Container Service for Kubernetes (ACK) (AliyunCSFullAccess).
//
// > 2. Cleaning up permissions of an Alibaba Cloud account is not supported.
//
// > 3. Cleaning up the permissions of the user who performs this operation is not supported.
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
// Cleans up KubeConfig credentials and revokes RBAC permissions for a specified user. If you want to clean up KubeConfig credentials and revoke authorization for risky users such as those who have left the organization or whose accounts have been frozen, call CleanUserPermissions to clean up the issued KubeConfig credentials and RBAC permissions for the specified user.
//
// Description:
//
// >- You must have full access permissions on Container Service for Kubernetes (ACK) (AliyunCSFullAccess).
//
// >- Cleaning up permissions of an Alibaba Cloud account is not supported.
//
// >- Cleaning up the permissions of the user who performs this operation is not supported.
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
// Cleans up KubeConfig credentials and revokes RBAC permissions for a specified user. If you want to clean up KubeConfig credentials and revoke authorization for risky users such as those who have left the organization or whose accounts have been frozen, call CleanUserPermissions to clean up the issued KubeConfig credentials and RBAC permissions for the specified user.
//
// Description:
//
// >- You must have full access permissions on Container Service for Kubernetes (ACK) (AliyunCSFullAccess).
//
// >- Cleaning up permissions of an Alibaba Cloud account is not supported.
//
// >- Cleaning up the permissions of the user who performs this operation is not supported.
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
// Creates a self-healing rule.
//
// Description:
//
//	Notice: This API is not yet available. Stay tuned..
//
// @param request - CreateAutoRepairPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoRepairPolicyResponse
func (client *Client) CreateAutoRepairPolicyWithOptions(clusterId *string, request *CreateAutoRepairPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAutoRepairPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	Notice: This API is not yet available. Stay tuned..
//
// @param request - CreateAutoRepairPolicyRequest
//
// @return CreateAutoRepairPolicyResponse
func (client *Client) CreateAutoRepairPolicy(clusterId *string, request *CreateAutoRepairPolicyRequest) (_result *CreateAutoRepairPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAutoRepairPolicyResponse{}
	_body, _err := client.CreateAutoRepairPolicyWithOptions(clusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an elastic scaling configuration that allows the system to automatically increase or decrease compute resources based on the configured scaling rules to meet the workload demands of your cluster. During the creation procedure, you can specify scaling metrics and thresholds, scale-out order, cool-down period, and more.
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
// Creates an elastic scaling configuration that allows the system to automatically increase or decrease compute resources based on the configured scaling rules to meet the workload demands of your cluster. During the creation procedure, you can specify scaling metrics and thresholds, scale-out order, cool-down period, and more.
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
// You can create ACK clusters through OpenAPI, including ACK managed clusters, ACK Serverless clusters, ACK Edge clusters, and registered clusters. When creating a cluster, you will configure the cluster information, cluster components, and ACK-related cloud resources.
//
// Description:
//
// ### Generate OpenAPI request parameters through the console
//
// When calling the CreateCluster API to create a cluster, if the API call fails due to incorrect request parameter combinations, you can generate the required request parameter combinations through the console. Follow these steps:
//
// 1. Log in to the [Container Service management console](https://csnew.console.aliyun.com) and choose **Clusters*	- in the left navigation pane.
//
// 1. On the **Clusters*	- page, click **Cluster Template**.
//
// 1. In the dialog box, select the cluster type to create, click Create, and then configure the cluster information on the cluster configuration page.
//
// 1. After the configuration is complete, on the **Confirm Configuration*	- page, click **Equivalent Code*	- in the upper-right corner. The dialog box will display the parameter combinations required for creating the cluster, which you can copy and use.
//
//	Notice: Starting from July 4, 2026, some request parameters will no longer take effect. For change details and alternative parameter descriptions, see [Announcement on OpenAPI parameter changes and API deprecation for ACK cluster management](https://help.aliyun.com/document_detail/2932733.html).</notice>
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can create ACK clusters through OpenAPI, including ACK managed clusters, ACK Serverless clusters, ACK Edge clusters, and registered clusters. When creating a cluster, you will configure the cluster information, cluster components, and ACK-related cloud resources.
//
// Description:
//
// ### Generate OpenAPI request parameters through the console
//
// When calling the CreateCluster API to create a cluster, if the API call fails due to incorrect request parameter combinations, you can generate the required request parameter combinations through the console. Follow these steps:
//
// 1. Log in to the [Container Service management console](https://csnew.console.aliyun.com) and choose **Clusters*	- in the left navigation pane.
//
// 1. On the **Clusters*	- page, click **Cluster Template**.
//
// 1. In the dialog box, select the cluster type to create, click Create, and then configure the cluster information on the cluster configuration page.
//
// 1. After the configuration is complete, on the **Confirm Configuration*	- page, click **Equivalent Code*	- in the upper-right corner. The dialog box will display the parameter combinations required for creating the cluster, which you can copy and use.
//
//	Notice: Starting from July 4, 2026, some request parameters will no longer take effect. For change details and alternative parameter descriptions, see [Announcement on OpenAPI parameter changes and API deprecation for ACK cluster management](https://help.aliyun.com/document_detail/2932733.html).</notice>
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
// Initiates a cluster diagnosis.
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
// Initiates a cluster diagnosis.
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
// Creates a cluster inspection configuration.
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
// Creates a cluster inspection configuration.
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
// A node pool is a logical collection of nodes with the same attributes, allowing unified management and O&M of nodes, such as node upgrades and auto scaling. You can further leverage the automated O&M capabilities of node pools to use features such as automatic OS CVE vulnerability patching, automatic faulty node recovery, and automatic kubelet and containerd version upgrades, reducing O&M costs. You can call CreateClusterNodePool to create a node pool for a cluster.
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
// A node pool is a logical collection of nodes with the same attributes, allowing unified management and O&M of nodes, such as node upgrades and auto scaling. You can further leverage the automated O&M capabilities of node pools to use features such as automatic OS CVE vulnerability patching, automatic faulty node recovery, and automatic kubelet and containerd version upgrades, reducing O&M costs. You can call CreateClusterNodePool to create a node pool for a cluster.
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
// Creates a trigger for an application.
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
// Creates a trigger for an application.
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
// An orchestration template defines and describes a set of Kubernetes cluster resources in a declarative manner, specifying how applications should run or be configured. You can use these templates to automate the deployment and cluster management of resources, such as Pods, Services, Deployments, ConfigMaps, and PersistentVolumes. You can invoke the CreateTemplate operation to create an orchestration template.
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
// An orchestration template defines and describes a set of Kubernetes cluster resources in a declarative manner, specifying how applications should run or be configured. You can use these templates to automate the deployment and cluster management of resources, such as Pods, Services, Deployments, ConfigMaps, and PersistentVolumes. You can invoke the CreateTemplate operation to create an orchestration template.
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
// Configures a trigger for a cluster application. The trigger redeploys pods when specified conditions are met.
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
// Configures a trigger for a cluster application. The trigger redeploys pods when specified conditions are met.
//
// Description:
//
// > Creating a trigger only supports pod redeployment.
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
// Deletes alert contacts from ACK.
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
// Deletes alert contacts from ACK.
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
// # Delete ACK alert contact group
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
// # Delete ACK alert contact group
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
// Deletes an automatic fault recovery rule.
//
// Description:
//
//	Notice: This API is not yet available. Stay tuned..
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoRepairPolicyResponse
func (client *Client) DeleteAutoRepairPolicyWithOptions(clusterId *string, policyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAutoRepairPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an automatic fault recovery rule.
//
// Description:
//
//	Notice: This API is not yet available. Stay tuned..
//
// @return DeleteAutoRepairPolicyResponse
func (client *Client) DeleteAutoRepairPolicy(clusterId *string, policyId *string) (_result *DeleteAutoRepairPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAutoRepairPolicyResponse{}
	_body, _err := client.DeleteAutoRepairPolicyWithOptions(clusterId, policyId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If you no longer need a cluster, delete the cluster and choose whether to delete or retain the associated resources. Before you delete a cluster, manually clean up workloads (Deployments, StatefulSets, Jobs, and CronJobs). Otherwise, the cluster deletion may fail.
//
// Description:
//
// Risk notice:
//
// - Subscription ECS instances, Lingjun compute nodes, and other subscription resources in the cluster cannot be subject to automatic release. To avoid unnecessary billing, perform manual release of these resources. For more information, refer to Cluster deletion and node release rules.
//
// - Subscription APIServer SLB resources cannot be subject to automatic release. To avoid unnecessary billing, perform manual release of these resources.
//
// - VPCs, vSwitches, security groups, and RAM roles that are used by other resources cannot be deleted and are retained by default. Perform manual release of these resources.
//
// - Elastic Container Instances (ECIs) created through virtual nodes are subject to automatic release.
//
// - Some cloud resources created through the cluster are not subject to automatic release when the cluster is deleted. These resources continue to incur billing after the cluster is deleted. Release or retain them as needed. These resources include: Simple Log Service (SLS) projects automatically created by the cluster, and cloud disks purchased through dynamic storage volumes.
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
// If you no longer need a cluster, delete the cluster and choose whether to delete or retain the associated resources. Before you delete a cluster, manually clean up workloads (Deployments, StatefulSets, Jobs, and CronJobs). Otherwise, the cluster deletion may fail.
//
// Description:
//
// Risk notice:
//
// - Subscription ECS instances, Lingjun compute nodes, and other subscription resources in the cluster cannot be subject to automatic release. To avoid unnecessary billing, perform manual release of these resources. For more information, refer to Cluster deletion and node release rules.
//
// - Subscription APIServer SLB resources cannot be subject to automatic release. To avoid unnecessary billing, perform manual release of these resources.
//
// - VPCs, vSwitches, security groups, and RAM roles that are used by other resources cannot be deleted and are retained by default. Perform manual release of these resources.
//
// - Elastic Container Instances (ECIs) created through virtual nodes are subject to automatic release.
//
// - Some cloud resources created through the cluster are not subject to automatic release when the cluster is deleted. These resources continue to incur billing after the cluster is deleted. Release or retain them as needed. These resources include: Simple Log Service (SLS) projects automatically created by the cluster, and cloud disks purchased through dynamic storage volumes.
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
// Deletes a cluster inspection configuration.
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
// Deletes a cluster inspection configuration.
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
// Deletes a node pool that is no longer in use. When a node pool is deleted, all pods on the nodes are deleted, which may trigger descheduling. If descheduling cannot be performed, your services may be affected. Make sure that the cluster has sufficient resources for descheduling.
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
// Deletes a node pool that is no longer in use. When a node pool is deleted, all pods on the nodes are deleted, which may trigger descheduling. If descheduling cannot be performed, your services may be affected. Make sure that the cluster has sufficient resources for descheduling.
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
// When you no longer need cluster nodes to continue working, you can call the DeleteClusterNodes operation to remove nodes from the cluster. When removing nodes, you can choose whether to simultaneously release the ECS instances and whether to automatically drain the nodes.
//
// Description:
//
// - When removing nodes, use this API or the [Container Service console](https://cs.console.aliyun.com) to perform standardized operations. Do not manually remove nodes by using the `kubectl delete node` command.
//
// - Do not directly release nodes, remove instances in the ECS or ESS console (or through related APIs), or allow nodes to be passively released due to the expiration of subscription instances. In these cases, the nodes will be directly shut down and automatically removed from the Container Service management console.
//
// - If the node pool has a desired number of nodes configured, the node pool will automatically scale out other instances according to the corresponding configuration to always maintain the node count at the desired number.
//
// - Removing nodes involves Pod migration, which may affect your business. Perform this operation during off-peak hours. Unexpected risks may exist during the operation, so back up relevant data in advance.
//
// - When removing nodes, ACK performs a drain operation. Ensure that the resources on other nodes in the cluster are sufficient to prevent business Pods from being unable to be scheduled.
//
// - Check the node affinity rules and scheduling policies of the Pods on the nodes to be removed to ensure that Pods can still be scheduled to other nodes after node removal.
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
// When you no longer need cluster nodes to continue working, you can call the DeleteClusterNodes operation to remove nodes from the cluster. When removing nodes, you can choose whether to simultaneously release the ECS instances and whether to automatically drain the nodes.
//
// Description:
//
// - When removing nodes, use this API or the [Container Service console](https://cs.console.aliyun.com) to perform standardized operations. Do not manually remove nodes by using the `kubectl delete node` command.
//
// - Do not directly release nodes, remove instances in the ECS or ESS console (or through related APIs), or allow nodes to be passively released due to the expiration of subscription instances. In these cases, the nodes will be directly shut down and automatically removed from the Container Service management console.
//
// - If the node pool has a desired number of nodes configured, the node pool will automatically scale out other instances according to the corresponding configuration to always maintain the node count at the desired number.
//
// - Removing nodes involves Pod migration, which may affect your business. Perform this operation during off-peak hours. Unexpected risks may exist during the operation, so back up relevant data in advance.
//
// - When removing nodes, ACK performs a drain operation. Ensure that the resources on other nodes in the cluster are sufficient to prevent business Pods from being unable to be scheduled.
//
// - Check the node affinity rules and scheduling policies of the Pods on the nodes to be removed to ensure that Pods can still be scheduled to other nodes after node removal.
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
// Deletes an application trigger by trigger ID.
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
// Deletes an application trigger by trigger ID.
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
// Deletes a policy rule instance from a specified cluster.
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
// Deletes a policy rule instance from a specified cluster.
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
// Deletes an orchestration template when you no longer need it.
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
// Deletes an orchestration template when you no longer need it.
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
// Deletes an application trigger when it is no longer needed.
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
// Deletes an application trigger when it is no longer needed.
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
// Deploys a policy instance in a specified namespace of a specified cluster. You can select a security policy type in an ACK cluster, configure the governance action (alerting or blocking) and the namespace scope for the policy instance to create and deploy a policy instance.
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
// Deploys a policy instance in a specified namespace of a specified cluster. You can select a security policy type in an ACK cluster, configure the governance action (alerting or blocking) and the namespace scope for the policy instance to create and deploy a policy instance.
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
// You can call the DescribeAddon operation to query information about a specified component based on parameters such as region, cluster type, cluster subtype (profile), cluster version, and component name. The returned information includes whether the component is managed, component category, supported custom parameter schema, compatible OS architectures, and the minimum cluster version required by the component version.
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
// You can call the DescribeAddon operation to query information about a specified component based on parameters such as region, cluster type, cluster subtype (profile), cluster version, and component name. The returned information includes whether the component is managed, component category, supported custom parameter schema, compatible OS architectures, and the minimum cluster version required by the component version.
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
// Queries the details of all components supported by the platform.
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
// Queries the details of all components supported by the platform.
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

// Summary:
//
// Queries an auto-repair rule.
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
func (client *Client) DescribeAutoRepairPolicyWithOptions(clusterId *string, policyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAutoRepairPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an auto-repair rule.
//
// Description:
//
//	Notice: This API is not yet available. Stay tuned.
//
// @return DescribeAutoRepairPolicyResponse
func (client *Client) DescribeAutoRepairPolicy(clusterId *string, policyId *string) (_result *DescribeAutoRepairPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAutoRepairPolicyResponse{}
	_body, _err := client.DescribeAutoRepairPolicyWithOptions(clusterId, policyId, headers, runtime)
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
// Calls DescribeClusterAddonInstance to query information about an installed cluster component, such as its version, status, and configuration.
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
// Calls DescribeClusterAddonInstance to query information about an installed cluster component, such as its version, status, and configuration.
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
// Queries the version information of a specified component that can be used in a specific cluster, including the component version and configurable parameters.
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
// Queries the version information of a specified component that can be used in a specific cluster, including the component version and configurable parameters.
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
// Queries the upgrade status of a cluster component by calling DescribeClusterAddonUpgradeStatus.
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
// Queries the upgrade status of a cluster component by calling DescribeClusterAddonUpgradeStatus.
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
// Queries the upgrade status of a component by component name.
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
// Queries the upgrade status of a component by component name.
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
// Queries the details of all installed components in a cluster by cluster ID.
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
// Queries the details of all installed components in a cluster by cluster ID.
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
// If you need to add existing ECS instances to an ACK cluster as worker nodes, or re-add node instances to a node pool after they are removed, ACK allows you to manually add existing nodes to a node pool. Calls the DescribeClusterAttachScripts operation to obtain the script for adding existing nodes.
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
// If you need to add existing ECS instances to an ACK cluster as worker nodes, or re-add node instances to a node pool after they are removed, ACK allows you to manually add existing nodes to a node pool. Calls the DescribeClusterAttachScripts operation to obtain the script for adding existing nodes.
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
// You can call the DescribeClusterDetail API to query the detailed information of a specified cluster by cluster ID.
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
// You can call the DescribeClusterDetail API to query the detailed information of a specified cluster by cluster ID.
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
// Cluster operation events include cluster creation, cluster modification, node pool creation, node pool scale-out, addon installation, and cluster upgrade. You can call the DescribeClusterEvents operation to retrieve the list of events that occurred in a specified cluster and query event details, including the event level, event status, and event time.
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cluster operation events include cluster creation, cluster modification, node pool creation, node pool scale-out, addon installation, and cluster upgrade. You can call the DescribeClusterEvents operation to retrieve the list of events that occurred in a specified cluster and query event details, including the event level, event status, and event time.
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
// Retrieves the log data of a cluster for root cause analysis and tracing when cluster issues occur.
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
// Retrieves the log data of a cluster for root cause analysis and tracing when cluster issues occur.
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
// You can call the DescribeClusterNodePoolDetail operation to query the configuration of a specified node pool in a cluster by node pool ID.
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
// You can call the DescribeClusterNodePoolDetail operation to query the configuration of a specified node pool in a cluster by node pool ID.
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
// Queries the list of all node pools in a cluster.
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
// Queries the list of all node pools in a cluster.
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
// Queries the list of nodes that meet the specified conditions in a cluster.
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of nodes that meet the specified conditions in a cluster.
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
// When you use Container Service for Kubernetes (ACK), you also use resources from other associated Alibaba Cloud services. You can call the DescribeClusterResources operation to query the associated resources of a specified cluster, such as VPCs and SLBs. To query node pool or node resources, call the DescribeClusterNodePools or DescribeClusterNodes operation.
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
// When you use Container Service for Kubernetes (ACK), you also use resources from other associated Alibaba Cloud services. You can call the DescribeClusterResources operation to query the associated resources of a specified cluster, such as VPCs and SLBs. To query node pool or node resources, call the DescribeClusterNodePools or DescribeClusterNodes operation.
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
// Queries the task list of a cluster by calling the DescribeClusterTasks operation.
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the task list of a cluster by calling the DescribeClusterTasks operation.
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
// KubeConfig is used to configure access credentials for ACK clusters on the client. It contains identity and authentication data for accessing the target cluster. When you use kubectl for cluster management, you must first connect to the cluster by using KubeConfig. You can invoke the DescribeClusterUserKubeconfig operation to query the KubeConfig of a specified cluster.
//
// Description:
//
// - The default validity period of the certificate issued by KubeConfig is 3 years. Within 180 days before the certificate expires, you can obtain a rotated and refreshed KubeConfig by using the Container Service console or the DescribeClusterUserKubeconfig operation. The new KubeConfig certificate is valid for 3 years. The old KubeConfig credential remains valid until the certificate expires. Obtain the rotated credential promptly based on the KubeConfig expiration time displayed in the console or returned by the operation.
//
// - Properly manage the KubeConfig credentials of your cluster and revoke them when they are no longer needed to avoid security risks such as sensitive data leakage caused by KubeConfig exposure.
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
// KubeConfig is used to configure access credentials for ACK clusters on the client. It contains identity and authentication data for accessing the target cluster. When you use kubectl for cluster management, you must first connect to the cluster by using KubeConfig. You can invoke the DescribeClusterUserKubeconfig operation to query the KubeConfig of a specified cluster.
//
// Description:
//
// - The default validity period of the certificate issued by KubeConfig is 3 years. Within 180 days before the certificate expires, you can obtain a rotated and refreshed KubeConfig by using the Container Service console or the DescribeClusterUserKubeconfig operation. The new KubeConfig certificate is valid for 3 years. The old KubeConfig credential remains valid until the certificate expires. Obtain the rotated credential promptly based on the KubeConfig expiration time displayed in the console or returned by the operation.
//
// - Properly manage the KubeConfig credentials of your cluster and revoke them when they are no longer needed to avoid security risks such as sensitive data leakage caused by KubeConfig exposure.
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
// Retrieves the kubeconfig file for a cluster.
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
// Retrieves the kubeconfig file for a cluster.
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
// You can call the DescribeClusterVuls operation to query the details of security vulnerabilities in a cluster by cluster ID, including vulnerability names, types, and severity levels. Regularly scan your cluster for security vulnerabilities to improve cluster security.
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
// You can call the DescribeClusterVuls operation to query the details of security vulnerabilities in a cluster by cluster ID, including vulnerability names, types, and severity levels. Regularly scan your cluster for security vulnerabilities to improve cluster security.
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
// View all clusters created in Container Service (including Swarm and Kubernetes clusters).
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
// View all clusters created in Container Service (including Swarm and Kubernetes clusters).
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
// You can invoke the DescribeClustersV1 operation to query the list of ACK clusters that meet conditional criteria (such as cluster type and cluster specification) under the current account.
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
// You can invoke the DescribeClustersV1 operation to query the list of ACK clusters that meet conditional criteria (such as cluster type and cluster specification) under the current account.
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
// Cluster operation events include cluster creation, cluster modification, node pool creation, node pool scale-out, addon installation, and cluster upgrade. You can call the DescribeEvents operation to query the details of a specific type of event, including the event level, event status, and event time.
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cluster operation events include cluster creation, cluster modification, node pool creation, node pool scale-out, addon installation, and cluster upgrade. You can call the DescribeEvents operation to query the details of a specific type of event, including the event level, event status, and event time.
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
// Queries the agent configuration for a registered cluster by cluster ID.
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
// Queries the agent configuration for a registered cluster by cluster ID.
//
// Description:
//
// For more information about cluster registration, see [Register an external Kubernetes cluster](https://help.aliyun.com/document_detail/121053.html).
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
// You can call the DescribeKubernetesVersionMetadata operation to query detailed information about Kubernetes versions, including version information, release dates and expiration dates, compatible operating systems, and container runtimes.
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
// You can call the DescribeKubernetesVersionMetadata operation to query detailed information about Kubernetes versions, including version information, release dates and expiration dates, compatible operating systems, and container runtimes.
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
// Queries the security vulnerability details of a node pool by node pool ID by calling the DescribeNodePoolVuls operation. The details include vulnerability names and severity levels. Regularly scan node pools for security vulnerabilities to improve cluster security.
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
// Queries the security vulnerability details of a node pool by node pool ID by calling the DescribeNodePoolVuls operation. The details include vulnerability names and severity levels. Regularly scan node pools for security vulnerabilities to improve cluster security.
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
// ACK cluster container security policies provide a comprehensive built-in rule library that includes Compliance, Infra, K8s-general, and PSP categories to ensure the secure operation of containers in production environments. You can call the DescribePolicies operation to query the list of policy governance rule libraries.
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
// ACK cluster container security policies provide a comprehensive built-in rule library that includes Compliance, Infra, K8s-general, and PSP categories to ensure the secure operation of containers in production environments. You can call the DescribePolicies operation to query the list of policy governance rule libraries.
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
// ACK cluster container security policies provide a rich set of built-in rule libraries, including Compliance, Infra, K8s-general, and PSP, to ensure the secure operation of containers in production environments. You can call the DescribePolicyDetails operation to query the details of a specified policy governance rule, such as the rule template description, governance action, and governance severity level.
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
// ACK cluster container security policies provide a rich set of built-in rule libraries, including Compliance, Infra, K8s-general, and PSP, to ensure the secure operation of containers in production environments. You can call the DescribePolicyDetails operation to query the details of a specified policy governance rule, such as the rule template description, governance action, and governance severity level.
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
// ACK cluster container security policies provide a rich set of built-in rule libraries, including Compliance, Infra, K8s-general, and PSP, to ensure the secure operation of containers in production environments. You can call the DescribePolicyGovernanceInCluster operation to query detailed policy governance information for a specified cluster, such as the count of enabled policies at different severity levels, policy governance audit logs, and interception and alert details.
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
// ACK cluster container security policies provide a rich set of built-in rule libraries, including Compliance, Infra, K8s-general, and PSP, to ensure the secure operation of containers in production environments. You can call the DescribePolicyGovernanceInCluster operation to query detailed policy governance information for a specified cluster, such as the count of enabled policies at different severity levels, policy governance audit logs, and interception and alert details.
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
// You can select a security policy type in an ACK cluster, configure the related enforcement actions and scope, and create and deploy a policy instance. You can call the DescribePolicyInstances operation to retrieve the details of specified policy instances in a cluster, such as the policy description and governance level.
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
// You can select a security policy type in an ACK cluster, configure the related enforcement actions and scope, and create and deploy a policy instance. You can call the DescribePolicyInstances operation to retrieve the details of specified policy instances in a cluster, such as the policy description and governance level.
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
// Queries the deployment status of policy instances for different policy types in a cluster, including the number of enabled instances for each policy rule and the number of enabled policy types at different governance levels.
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
// Queries the deployment status of policy instances for different policy types in a cluster, including the number of enabled instances for each policy rule and the number of enabled policy types at different governance levels.
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
// Queries the list of regions.
//
// @param request - DescribeRegionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether deletion protection is enabled for a specified resource in a cluster. Resources that support deletion protection include namespaces and services.
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
// Queries whether deletion protection is enabled for a specified resource in a cluster. Resources that support deletion protection include namespaces and services.
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
// As a cluster permission management administrator, you can use an Alibaba Cloud account to issue KubeConfig credentials that contain identity information for a specified Resource Access Management (RAM) user or RAM role within the account. These credentials are used to connect to ACK clusters. You can invoke the DescribeSubaccountK8sClusterUserConfig operation to issue or retrieve the KubeConfig for any RAM user or role within the account.
//
// Description:
//
// This operation can be called only by an Alibaba Cloud account.
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
// As a cluster permission management administrator, you can use an Alibaba Cloud account to issue KubeConfig credentials that contain identity information for a specified Resource Access Management (RAM) user or RAM role within the account. These credentials are used to connect to ACK clusters. You can invoke the DescribeSubaccountK8sClusterUserConfig operation to issue or retrieve the KubeConfig for any RAM user or role within the account.
//
// Description:
//
// This operation can be called only by an Alibaba Cloud account.
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
// Queries the details of a cluster task, such as the task type, running status, and running stage.
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
// Queries the details of a cluster task, such as the task type, running status, and running stage.
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
// An orchestration template defines and describes a set of Kubernetes cluster resources in a declarative manner, specifying how applications should run or be configured. You can call the DescribeTemplateAttribute operation to query the details of a specified orchestration template, including access permissions, YAML content, and labels.
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
// An orchestration template defines and describes a set of Kubernetes cluster resources in a declarative manner, specifying how applications should run or be configured. You can call the DescribeTemplateAttribute operation to query the details of a specified orchestration template, including access permissions, YAML content, and labels.
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
// An orchestration template defines and describes a set of Kubernetes cluster resources in a declarative manner, specifying how applications should run or be configured. You can call the DescribeTemplates operation to retrieve a list of created orchestration templates and query detailed information about the templates, including access permissions, YAML content, and tags.
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
// An orchestration template defines and describes a set of Kubernetes cluster resources in a declarative manner, specifying how applications should run or be configured. You can call the DescribeTemplates operation to retrieve a list of created orchestration templates and query detailed information about the templates, including access permissions, YAML content, and tags.
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
// Queries triggers that meet the specified conditions.
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
// Queries triggers that meet the specified conditions.
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
// In ACK clusters, you can use Kubernetes namespaces to achieve logically isolated permissions and resources for cluster users. Users who are granted RBAC permissions only for a specified namespace cannot access resources in other namespaces of the cluster. You can invoke the DescribeUserClusterNamespaces operation to query the namespaces for which the current Resource Access Management (RAM) user or role has been granted RBAC access permissions in a specified ACK cluster.
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
// In ACK clusters, you can use Kubernetes namespaces to achieve logically isolated permissions and resources for cluster users. Users who are granted RBAC permissions only for a specified namespace cannot access resources in other namespaces of the cluster. You can invoke the DescribeUserClusterNamespaces operation to query the namespaces for which the current Resource Access Management (RAM) user or role has been granted RBAC access permissions in a specified ACK cluster.
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
// In ACK clusters, you can create and assign different access permissions to different Resource Access Management (RAM) users or roles to ensure secure access control and resource isolation. You can invoke the DescribeUserPermission operation to query the details of cluster permissions granted to a RAM user or role, including accessible resources, permission scope, preset role types, and permission sources.
//
// Description:
//
// *Before you begin**:
//
// - If the account that invokes this API operation is a Resource Access Management (RAM) user or RAM role, the API operation returns only the permissions for clusters in which the calling account has RBAC administrator permissions. To list permissions for all clusters, the calling account must have RBAC administrator permissions on all clusters.
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
// In ACK clusters, you can create and assign different access permissions to different Resource Access Management (RAM) users or roles to ensure secure access control and resource isolation. You can invoke the DescribeUserPermission operation to query the details of cluster permissions granted to a RAM user or role, including accessible resources, permission scope, preset role types, and permission sources.
//
// Description:
//
// *Before you begin**:
//
// - If the account that invokes this API operation is a Resource Access Management (RAM) user or RAM role, the API operation returns only the permissions for clusters in which the calling account has RBAC administrator permissions. To list permissions for all clusters, the calling account must have RBAC administrator permissions on all clusters.
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
// Queries the quotas of ACK clusters, node pools, and nodes. To increase a quota, go to Quota Center to submit a request.
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
// Queries the quotas of ACK clusters, node pools, and nodes. To increase a quota, go to Quota Center to submit a request.
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
// Alibaba Cloud security products periodically scan ECS nodes for security vulnerabilities and provide corresponding remediation suggestions and methods. Some CVE fixes may require node restarts. Ensure that the cluster has sufficient nodes for drain operations. You can call the FixNodePoolVuls operation to fix node security vulnerabilities in a specified cluster node pool and improve the security of cluster nodes.
//
// Description:
//
// - CVE compatibility is ensured by Security Center. Make sure that you have activated the Ultimate edition of Security Center or [purchased vulnerability fixing (pay-as-you-go)](https://help.aliyun.com/document_detail/42308.html).
//
// - Some CVE fixes require node restarts. Container Service drains the node before restarting it. Ensure that the cluster has sufficient spare node resources for draining. For example, scale out the node pool in advance.
//
// - Pay attention to the compatibility between your applications and CVEs. CVE fixes are performed in batches. During the CVE fix process, you can pause or cancel the task. After you pause or cancel the task, batches that have already been dispatched continue to run until completion. Batches that have not been dispatched are paused or canceled.
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
// Alibaba Cloud security products periodically scan ECS nodes for security vulnerabilities and provide corresponding remediation suggestions and methods. Some CVE fixes may require node restarts. Ensure that the cluster has sufficient nodes for drain operations. You can call the FixNodePoolVuls operation to fix node security vulnerabilities in a specified cluster node pool and improve the security of cluster nodes.
//
// Description:
//
// - CVE compatibility is ensured by Security Center. Make sure that you have activated the Ultimate edition of Security Center or [purchased vulnerability fixing (pay-as-you-go)](https://help.aliyun.com/document_detail/42308.html).
//
// - Some CVE fixes require node restarts. Container Service drains the node before restarting it. Ensure that the cluster has sufficient spare node resources for draining. For example, scale out the node pool in advance.
//
// - Pay attention to the compatibility between your applications and CVEs. CVE fixes are performed in batches. During the CVE fix process, you can pause or cancel the task. After you pause or cancel the task, batches that have already been dispatched continue to run until completion. Batches that have not been dispatched are paused or canceled.
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
// Queries the details of a specified component instance in a cluster, including the version, parameter settings, and logging feature status of the component instance.
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
// Queries the details of a specified component instance in a cluster, including the version, parameter settings, and logging feature status of the component instance.
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
// Queries whether the API server audit feature is enabled for a cluster and retrieves the Simple Log Service (SLS) project that stores the API server audit logs.
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
// Queries whether the API server audit feature is enabled for a cluster and retrieves the Simple Log Service (SLS) project that stores the API server audit logs.
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
// The intelligent O&M platform for containers provides comprehensive Kubernetes cluster check capabilities, including cluster upgrade checks, cluster migration checks, component installation checks, component upgrade checks, and node pool checks. You can call the GetClusterCheck operation to query information about a specified check task based on the cluster ID and check task ID, such as the check status, specific check items, and check creation and completion time.
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
// The intelligent O&M platform for containers provides comprehensive Kubernetes cluster check capabilities, including cluster upgrade checks, cluster migration checks, component installation checks, component upgrade checks, and node pool checks. You can call the GetClusterCheck operation to query information about a specified check task based on the cluster ID and check task ID, such as the check status, specific check items, and check creation and completion time.
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
// Retrieves the diagnostic check items of a cluster.
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
// Retrieves the diagnostic check items of a cluster.
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
// Retrieves the diagnosis result of a cluster.
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
// Retrieves the diagnosis result of a cluster.
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
// Retrieves the inspection configuration of a cluster.
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
// Retrieves the inspection configuration of a cluster.
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
// Retrieves the details of a cluster inspection report.
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
// Retrieves the details of a cluster inspection report.
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
// Queries the triggers of an application by application name.
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
// Queries the triggers of an application by application name.
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
// Queries the upgrade status of a cluster by cluster ID.
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
// Queries the upgrade status of a cluster by cluster ID.
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
// By default, Resource Access Management (RAM) users or roles that are not cluster creators and have not been granted access permissions across all cluster dimensions do not have any RBAC permissions in a cluster. You can invoke the GrantPermissions operation to update the RBAC access permissions of a RAM user or role, including accessible resources, permission scope, and preset role types, to better manage cluster management security and access control.
//
// Description:
//
// - If the account that invokes this API operation is a RAM user, make sure that the account has been granted the permission to modify the RBAC authorization information of other Resource Access Management (RAM) users or RAM roles. Otherwise, the API operation returns the `StatusForbidden` or `ForbiddenGrantPermissions` fault. For more information, see [Grant RBAC permissions to a RAM user](https://help.aliyun.com/document_detail/119035.html).
//
// - The operation of fully updating the RBAC authorization information of a RAM user or RAM role overwrites the existing cluster permissions of the target RAM user or RAM role. Include all permission configurations that you want to grant to the target RAM user or RAM role in the request.
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
// By default, Resource Access Management (RAM) users or roles that are not cluster creators and have not been granted access permissions across all cluster dimensions do not have any RBAC permissions in a cluster. You can invoke the GrantPermissions operation to update the RBAC access permissions of a RAM user or role, including accessible resources, permission scope, and preset role types, to better manage cluster management security and access control.
//
// Description:
//
// - If the account that invokes this API operation is a RAM user, make sure that the account has been granted the permission to modify the RBAC authorization information of other Resource Access Management (RAM) users or RAM roles. Otherwise, the API operation returns the `StatusForbidden` or `ForbiddenGrantPermissions` fault. For more information, see [Grant RBAC permissions to a RAM user](https://help.aliyun.com/document_detail/119035.html).
//
// - The operation of fully updating the RBAC authorization information of a RAM user or RAM role overwrites the existing cluster permissions of the target RAM user or RAM role. Include all permission configurations that you want to grant to the target RAM user or RAM role in the request.
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
// To enhance Kubernetes capabilities, ACK clusters support various components, such as managed core components, application components, logging and monitoring components, networking components, storage components, and security components. You can call the InstallClusterAddons operation to install components by specifying the component name and version.
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
// To enhance Kubernetes capabilities, ACK clusters support various components, such as managed core components, application components, logging and monitoring components, networking components, storage components, and security components. You can call the InstallClusterAddons operation to install components by specifying the component name and version.
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
// Installs components on nodes. You can configure and specify nodes on which to install components.
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
// Installs components on nodes. You can configure and specify nodes on which to install components.
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
// Queries the list of available components based on specified parameters such as region, cluster type, cluster subtype (profile), and cluster version. You can also retrieve detailed component information, including whether a component is managed, the schema of supported custom parameters, and compatible operating system architectures.
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
// Queries the list of available components based on specified parameters such as region, cluster type, cluster subtype (profile), and cluster version. You can also retrieve detailed component information, including whether a component is managed, the schema of supported custom parameters, and compatible operating system architectures.
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
// Retrieves the list of auto-repair policies.
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
func (client *Client) ListAutoRepairPoliciesWithOptions(clusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAutoRepairPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of auto-repair policies.
//
// Description:
//
//	Notice: This API is not yet available. Stay tuned.
//
// @return ListAutoRepairPoliciesResponse
func (client *Client) ListAutoRepairPolicies(clusterId *string) (_result *ListAutoRepairPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAutoRepairPoliciesResponse{}
	_body, _err := client.ListAutoRepairPoliciesWithOptions(clusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// Queries the list of resources contained in an installed cluster component instance, including Kubernetes cluster resources and Helm release instances.
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
// Lists the component instances installed in a specified cluster and queries related information about the component instances, such as the component version and status.
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
// Lists the component instances installed in a specified cluster and queries related information about the component instances, such as the component version and status.
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
// The intelligent O&M platform for containers provides comprehensive Kubernetes cluster check capabilities, including cluster upgrade checks, cluster migration checks, component installation checks, component upgrade checks, and node pool checks. You can call the ListClusterChecks operation to query the list of cluster checks and related information by cluster ID, such as check type, status, creation time, and completion time.
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
// The intelligent O&M platform for containers provides comprehensive Kubernetes cluster check capabilities, including cluster upgrade checks, cluster migration checks, component installation checks, component upgrade checks, and node pool checks. You can call the ListClusterChecks operation to query the list of cluster checks and related information by cluster ID, such as check type, status, creation time, and completion time.
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
// Retrieves a list of cluster inspection reports.
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
// Retrieves a list of cluster inspection reports.
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
// Queries the list and status of KubeConfigs that have been issued to users in a specified cluster. You can call this operation to view the access control status of the current cluster.
//
// Description:
//
// > - To call this operation, you must have the ram:ListUsers and ram:ListRoles permissions.
//
// > - To call this operation, you must have full access to Container Service for Kubernetes (ACK) (AliyunCSFullAccess).
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
// Queries the list and status of KubeConfigs that have been issued to users in a specified cluster. You can call this operation to view the access control status of the current cluster.
//
// Description:
//
// > - To call this operation, you must have the ram:ListUsers and ram:ListRoles permissions.
//
// > - To call this operation, you must have full access to Container Service for Kubernetes (ACK) (AliyunCSFullAccess).
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
// Retrieves a list of automated O&M execution plans.
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
// Retrieves a list of automated O&M execution plans.
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
// Queries the most recent 100 automated O&M execution plans in a specified region. When features such as cluster intelligent managed mode (Auto Mode), automatic cluster upgrade, or node pool automated O&M are enabled, you can call this operation to query the O&M plans automatically generated by the system and their execution status, such as cluster upgrades and node pool CVE fixes.
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
// Queries the most recent 100 automated O&M execution plans in a specified region. When features such as cluster intelligent managed mode (Auto Mode), automatic cluster upgrade, or node pool automated O&M are enabled, you can call this operation to query the O&M plans automatically generated by the system and their execution status, such as cluster upgrades and node pool CVE fixes.
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
// You can add tag key-value pairs to clusters so that cluster developers or O&M engineers can classify and manage clusters more flexibly, and better support requirements such as monitoring, cost analysis, and multi-tenant data isolation. You can call the ListTagResources operation to obtain a list of resource tags and query detailed tag information, such as key-value pairs and associated clusters.
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
// You can add tag key-value pairs to clusters so that cluster developers or O&M engineers can classify and manage clusters more flexibly, and better support requirements such as monitoring, cost analysis, and multi-tenant data isolation. You can call the ListTagResources operation to obtain a list of resource tags and query detailed tag information, such as key-value pairs and associated clusters.
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
// Queries the KubeConfig status list of all clusters for the current user. If you want to view the KubeConfig issuance status of each cluster for the current user, you can call this operation to retrieve the KubeConfig status list of all clusters.
//
// Description:
//
// > To call this operation, you must have full access permissions on Container Service for Kubernetes (ACK) (AliyunCSFullAccess).
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
// Queries the KubeConfig status list of all clusters for the current user. If you want to view the KubeConfig issuance status of each cluster for the current user, you can call this operation to retrieve the KubeConfig status list of all clusters.
//
// Description:
//
// > To call this operation, you must have full access permissions on Container Service for Kubernetes (ACK) (AliyunCSFullAccess).
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
// ACK ACK clusters are an evolution of ACK Basic clusters, inheriting all the advantages of managed clusters, such as managed control planes and high-availability control planes. ACK ACK clusters further enhance cluster reliability, security, and scheduling capabilities, and support SLAs with compensation standards. ACK ACK clusters are suitable for enterprise customers who run large-scale workloads in production environments and have high requirements for stability and security. You can call the MigrateCluster operation to migrate an ACK Basic cluster to an ACK ACK cluster.
//
// Description:
//
// After you migrate an ACK managed Basic cluster to an ACK managed Pro cluster, a [cluster management fee](https://help.aliyun.com/document_detail/462278.html) charged by ACK is added. Billing for other cloud resources remains unchanged.
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
// ACK ACK clusters are an evolution of ACK Basic clusters, inheriting all the advantages of managed clusters, such as managed control planes and high-availability control planes. ACK ACK clusters further enhance cluster reliability, security, and scheduling capabilities, and support SLAs with compensation standards. ACK ACK clusters are suitable for enterprise customers who run large-scale workloads in production environments and have high requirements for stability and security. You can call the MigrateCluster operation to migrate an ACK Basic cluster to an ACK ACK cluster.
//
// Description:
//
// After you migrate an ACK managed Basic cluster to an ACK managed Pro cluster, a [cluster management fee](https://help.aliyun.com/document_detail/462278.html) charged by ACK is added. Billing for other cloud resources remains unchanged.
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
// Modifies a self-healing rule.
//
// Description:
//
//	Notice: This API is not yet available. Stay tuned..
//
// @param request - ModifyAutoRepairPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAutoRepairPolicyResponse
func (client *Client) ModifyAutoRepairPolicyWithOptions(clusterId *string, policyId *string, request *ModifyAutoRepairPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyAutoRepairPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	Notice: This API is not yet available. Stay tuned..
//
// @param request - ModifyAutoRepairPolicyRequest
//
// @return ModifyAutoRepairPolicyResponse
func (client *Client) ModifyAutoRepairPolicy(clusterId *string, policyId *string, request *ModifyAutoRepairPolicyRequest) (_result *ModifyAutoRepairPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyAutoRepairPolicyResponse{}
	_body, _err := client.ModifyAutoRepairPolicyWithOptions(clusterId, policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ModifyCluster operation to modify the configuration of an ACK cluster.
//
// Description:
//
// <notice>Starting from July 04, 2026, the request parameters instance_deletion_protection, ingress_loadbalancer_id, and access_control_list will no longer take effect. For details about the changes, see [Announcement on Changes to ACK Cluster Management OpenAPI Parameters and OpenAPI Deprecation](https://help.aliyun.com/document_detail/2932733.html).</notice>
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyCluster operation to modify the configuration of an ACK cluster.
//
// Description:
//
// <notice>Starting from July 04, 2026, the request parameters instance_deletion_protection, ingress_loadbalancer_id, and access_control_list will no longer take effect. For details about the changes, see [Announcement on Changes to ACK Cluster Management OpenAPI Parameters and OpenAPI Deprecation](https://help.aliyun.com/document_detail/2932733.html).</notice>
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
// Modifies the configuration of an installed cluster component instance. Modifying configurations may affect your services. Evaluate the impact before performing this operation during off-peak hours and back up relevant data in advance.
//
// Description:
//
// You can call this API operation to modify the configuration of common clusters components and the control plane parameter settings of ACK Pro clusters:
//
// - To query the configurable parameters of common components, call the DescribeClusterAddonMetadata API operation. For details, see [Query cluster component version metadata](https://help.aliyun.com/document_detail/2667944.html).
//
// - For the configurable control plane parameter settings of ACK Pro clusters, see [Customize control plane parameters of ACK Pro clusters](https://help.aliyun.com/document_detail/199588.html).
//
// Modifying configurations may cause the component to be redeployed and restarted. Evaluate the impact before performing this operation.
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
// Modifies the configuration of an installed cluster component instance. Modifying configurations may affect your services. Evaluate the impact before performing this operation during off-peak hours and back up relevant data in advance.
//
// Description:
//
// You can call this API operation to modify the configuration of common clusters components and the control plane parameter settings of ACK Pro clusters:
//
// - To query the configurable parameters of common components, call the DescribeClusterAddonMetadata API operation. For details, see [Query cluster component version metadata](https://help.aliyun.com/document_detail/2667944.html).
//
// - For the configurable control plane parameter settings of ACK Pro clusters, see [Customize control plane parameters of ACK Pro clusters](https://help.aliyun.com/document_detail/199588.html).
//
// Modifying configurations may cause the component to be redeployed and restarted. Evaluate the impact before performing this operation.
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
// You can call the ModifyClusterNodePool API to update the configuration of a target node pool by specifying its node pool ID.
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
// You can call the ModifyClusterNodePool API to update the configuration of a target node pool by specifying its node pool ID.
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
// You can add tag key-value pairs to clusters so that cluster developers or O&M engineers can categorize and manage clusters more flexibly, and better support requirements such as monitoring, cost analysis, and tenant isolation. You can call the ModifyClusterTags operation to modify cluster tags.
//
// Description:
//
// - This operation performs a full update. When you call this operation, specify all target tags as input parameters to avoid losing existing tags. To perform an incremental update, use [TagResources to bind tags to a cluster](https://help.aliyun.com/document_detail/2667969.html).
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
// You can add tag key-value pairs to clusters so that cluster developers or O&M engineers can categorize and manage clusters more flexibly, and better support requirements such as monitoring, cost analysis, and tenant isolation. You can call the ModifyClusterTags operation to modify cluster tags.
//
// Description:
//
// - This operation performs a full update. When you call this operation, specify all target tags as input parameters to avoid losing existing tags. To perform an incremental update, use [TagResources to bind tags to a cluster](https://help.aliyun.com/document_detail/2667969.html).
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
// You can call the ModifyNodePoolNodeConfig operation to modify the node configuration in a cluster node pool, such as kubelet configuration and node rolling update configuration. Modifying node configuration changes the node configuration in batches and restarts kubelet, which may affect node operations and workload operations. We recommend that you perform this operation during off-peak hours.
//
// Description:
//
// > ACK allows you to modify the kubelet configuration of nodes in a node pool. After the modification is complete, the changes automatically take effect on the nodes in the node pool, and newly added nodes in the node pool also use the new configuration.
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyNodePoolNodeConfig operation to modify the node configuration in a cluster node pool, such as kubelet configuration and node rolling update configuration. Modifying node configuration changes the node configuration in batches and restarts kubelet, which may affect node operations and workload operations. We recommend that you perform this operation during off-peak hours.
//
// Description:
//
// > ACK allows you to modify the kubelet configuration of nodes in a node pool. After the modification is complete, the changes automatically take effect on the nodes in the node pool, and newly added nodes in the node pool also use the new configuration.
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
// Updates a policy rule instance in a specified cluster. You can modify the governance action (alert or block) and the scope of namespaces to which the policy instance applies.
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
// Updates a policy rule instance in a specified cluster. You can modify the governance action (alert or block) and the scope of namespaces to which the policy instance applies.
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
// When you use Container Service for Kubernetes (ACK) for the first time, you must call the OpenAckService operation to activate the service.
//
// Description:
//
// - An Alibaba Cloud account can activate ACK.
//
// - A Resource Access Management (RAM) user that has the AdministratorAccess permission can activate the service.
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
// When you use Container Service for Kubernetes (ACK) for the first time, you must call the OpenAckService operation to activate the service.
//
// Description:
//
// - An Alibaba Cloud account can activate ACK.
//
// - A Resource Access Management (RAM) user that has the AdministratorAccess permission can activate the service.
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
// Pauses a cluster upgrade.
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
// Pauses a cluster upgrade.
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
// Pauses a component upgrade.
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
// Pauses a component upgrade.
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
// You can call the PauseTask operation to pause a running cluster task.
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
// You can call the PauseTask operation to pause a running cluster task.
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
// 移除集群节点
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
// 移除集群节点
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
// Removes nodes from a node pool in a cluster and adjusts the expected number of nodes. When removing nodes, you can specify whether to release the associated ECS instances and whether to drain the nodes. Removing nodes involves pod migration, which may affect your services. Perform this operation during off-peak hours and back up your data in advance.
//
// Description:
//
// - Removing nodes involves pod migration, which may affect your services. Perform this operation during off-peak hours.
//
// - Unexpected risks may occur during the operation. Back up your data in advance.
//
// - During the operation, the nodes being removed are set to the unschedulable state in the background.
//
// - This operation removes only worker nodes, not master nodes.
//
//   - Even if you choose to release nodes (nodes for which `release_node` is set to `true`), subscription nodes are not released. After removing the nodes, release them in the [ECS console](https://ecs.console.aliyun.com/).
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
// Removes nodes from a node pool in a cluster and adjusts the expected number of nodes. When removing nodes, you can specify whether to release the associated ECS instances and whether to drain the nodes. Removing nodes involves pod migration, which may affect your services. Perform this operation during off-peak hours and back up your data in advance.
//
// Description:
//
// - Removing nodes involves pod migration, which may affect your services. Perform this operation during off-peak hours.
//
// - Unexpected risks may occur during the operation. Back up your data in advance.
//
// - During the operation, the nodes being removed are set to the unschedulable state in the background.
//
// - This operation removes only worker nodes, not master nodes.
//
//   - Even if you choose to release nodes (nodes for which `release_node` is set to `true`), subscription nodes are not released. After removing the nodes, release them in the [ECS console](https://ecs.console.aliyun.com/).
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
// Repairs nodes in a cluster node pool.
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
// Repairs nodes in a cluster node pool.
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
// Calls ResumeComponentUpgrade to restart a paused component upgrade task.
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
// Calls ResumeComponentUpgrade to restart a paused component upgrade task.
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
// Resumes a paused cluster task.
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
// Resumes a paused cluster task.
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
// Resumes the upgrade of a cluster that is in the upgrade-paused state based on the cluster ID.
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
// Resumes the upgrade of a cluster that is in the upgrade-paused state based on the cluster ID.
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
// If you want to revoke the cluster KubeConfig credential owned by the currently logged-on Alibaba Cloud account or Resource Access Management (RAM) user, you can call the RevokeK8sClusterKubeConfig operation to revoke it. After the revocation succeeds, the cluster generates a new KubeConfig, and the original KubeConfig becomes invalid.
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
// If you want to revoke the cluster KubeConfig credential owned by the currently logged-on Alibaba Cloud account or Resource Access Management (RAM) user, you can call the RevokeK8sClusterKubeConfig operation to revoke it. After the revocation succeeds, the cluster generates a new KubeConfig, and the original KubeConfig becomes invalid.
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
// The intelligent O&M platform for containers provides a wide range of Kubernetes cluster check capabilities, including cluster upgrade checks, cluster migration checks, component installation checks, component upgrade checks, and node pool checks. Before an upgrade, migration, or installation operation is performed, the platform automatically triggers a check. You can perform the change operation only after the check is passed. You can also manually call the RunClusterCheck operation to perform a cluster check. Periodically check and maintain your clusters to prevent security risks.
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
// The intelligent O&M platform for containers provides a wide range of Kubernetes cluster check capabilities, including cluster upgrade checks, cluster migration checks, component installation checks, component upgrade checks, and node pool checks. Before an upgrade, migration, or installation operation is performed, the platform automatically triggers a check. You can perform the change operation only after the check is passed. You can also manually call the RunClusterCheck operation to perform a cluster check. Periodically check and maintain your clusters to prevent security risks.
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
// Initiates a cluster inspection and creates an inspection report.
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
// Initiates a cluster inspection and creates an inspection report.
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
// 执行节点上的运维操作
//
// @param request - RunNodeOperationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunNodeOperationResponse
func (client *Client) RunNodeOperationWithOptions(clusterId *string, nodepoolId *string, nodeName *string, request *RunNodeOperationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunNodeOperationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行节点上的运维操作
//
// @param request - RunNodeOperationRequest
//
// @return RunNodeOperationResponse
func (client *Client) RunNodeOperation(clusterId *string, nodepoolId *string, nodeName *string, request *RunNodeOperationRequest) (_result *RunNodeOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunNodeOperationResponse{}
	_body, _err := client.RunNodeOperationWithOptions(clusterId, nodepoolId, nodeName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scales out a node pool by adding nodes to ensure that the number of nodes is sufficient to support your workloads.
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
// Scales out a node pool by adding nodes to ensure that the number of nodes is sufficient to support your workloads.
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
// 扩容Kubernetes集群
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
// 扩容Kubernetes集群
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
// You can call the ScanClusterVuls operation to scan for potential security vulnerabilities in an ACK cluster, including container workload vulnerabilities, third-party software vulnerabilities, CVE vulnerabilities, WebCMS vulnerabilities, and Windows operating system vulnerabilities. Regularly scan your cluster for security vulnerabilities to improve cluster security.
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
// You can call the ScanClusterVuls operation to scan for potential security vulnerabilities in an ACK cluster, including container workload vulnerabilities, third-party software vulnerabilities, CVE vulnerabilities, WebCMS vulnerabilities, and Windows operating system vulnerabilities. Regularly scan your cluster for security vulnerabilities to improve cluster security.
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
// Starts a specified alert rule.
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
// Starts a specified alert rule.
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
// Stops alert rules in the ACK alert center. You can stop an entire alert rule group or a single alert rule.
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
// Stops alert rules in the ACK alert center. You can stop an entire alert rule group or a single alert rule.
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
// Synchronizes a cluster node pool, including node pool metadata and information about the nodes in the node pool.
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
// Synchronizes a cluster node pool, including node pool metadata and information about the nodes in the node pool.
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
// Adds tag key-value pairs to clusters so that cluster developers or O&M engineers can categorize and manage clusters more flexibly, and better support monitoring, cost analysis, and tenant isolation requirements. You can call the TagResources operation to attach tags to clusters.
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
// Adds tag key-value pairs to clusters so that cluster developers or O&M engineers can categorize and manage clusters more flexibly, and better support monitoring, cost analysis, and tenant isolation requirements. You can call the TagResources operation to attach tags to clusters.
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
// Uninstalls specified components from a cluster when they are no longer needed, with the option to delete associated Alibaba Cloud resources.
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
// Uninstalls specified components from a cluster when they are no longer needed, with the option to delete associated Alibaba Cloud resources.
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
// Deletes resource tags when you no longer need the tag key-value pairs for a cluster. You can call the UntagResources operation to delete resource tags.
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
// Deletes resource tags when you no longer need the tag key-value pairs for a cluster. You can call the UntagResources operation to delete resource tags.
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
// When you need to record Kubernetes API requests and their results to trace cluster operation history or troubleshoot cluster issues, you can invoke the UpdateClusterAuditLogConfig operation to enable or shutdown the audit log feature for a specified ACK cluster and update the audit log configuration.
//
// Description:
//
// Before you use this operation, make sure that you fully understand the billing methods and pricing of <props="china">[Simple Log Service](https://www.aliyun.com/price/product#/sls/detail/sls)<props="intl">[Simple Log Service](https://www.alibabacloud.com/product/log-service/pricing).
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
// When you need to record Kubernetes API requests and their results to trace cluster operation history or troubleshoot cluster issues, you can invoke the UpdateClusterAuditLogConfig operation to enable or shutdown the audit log feature for a specified ACK cluster and update the audit log configuration.
//
// Description:
//
// Before you use this operation, make sure that you fully understand the billing methods and pricing of <props="china">[Simple Log Service](https://www.aliyun.com/price/product#/sls/detail/sls)<props="intl">[Simple Log Service](https://www.alibabacloud.com/product/log-service/pricing).
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
// Updates the cluster inspection configuration.
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
// Updates the cluster inspection configuration.
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
// Updates the contact group for an alert rule set in an ACK cluster.
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
// Updates the contact group for an alert rule set in an ACK cluster.
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
// ACK managed clusters support collecting control plane component logs and delivering them to your Simple Log Service (SLS) Log Project. Control plane components include Kube API Server, Kube Scheduler, Kube Controller Manager, Cloud Controller Manager, and other core components. You can call the UpdateControlPlaneLog operation to modify the control plane component log configuration, such as the log retention period and the components from which logs are collected.
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
// ACK managed clusters support collecting control plane component logs and delivering them to your Simple Log Service (SLS) Log Project. Control plane components include Kube API Server, Kube Scheduler, Kube Controller Manager, Cloud Controller Manager, and other core components. You can call the UpdateControlPlaneLog operation to modify the control plane component log configuration, such as the log retention period and the components from which logs are collected.
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
// The default expiration time of the KubeConfig issued by an ACK cluster is 3 years. You can use an Alibaba Cloud account to customize the configuration by invoking the UpdateK8sClusterUserConfigExpire operation to specify the expiration time (1 to 876,000 hours) of the KubeConfig issued to a Resource Access Management (RAM) user or role in an ACK cluster.
//
// Description:
//
// - This operation can be called only by an Alibaba Cloud account.
//
// - If you revoke the KubeConfig credential used in the cluster, the custom expiration time configured for the KubeConfig of the cluster is also reset. You must call this operation to reconfigure the expiration time.
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
// The default expiration time of the KubeConfig issued by an ACK cluster is 3 years. You can use an Alibaba Cloud account to customize the configuration by invoking the UpdateK8sClusterUserConfigExpire operation to specify the expiration time (1 to 876,000 hours) of the KubeConfig issued to a Resource Access Management (RAM) user or role in an ACK cluster.
//
// Description:
//
// - This operation can be called only by an Alibaba Cloud account.
//
// - If you revoke the KubeConfig credential used in the cluster, the custom expiration time configured for the KubeConfig of the cluster is also reset. You must call this operation to reconfigure the expiration time.
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
// Updates the Secret encryption at rest configuration for a specified cluster by cluster ID.
//
// Description:
//
// While enabling or disabling encryption at rest and after the feature is enabled, do not disable or delete the KMS key used by this feature in the KMS console or through OpenAPI. Otherwise, the cluster API Server becomes unavailable, which prevents normal retrieval of objects such as Secrets and ServiceAccounts and affects the normal operation of business applications. For more information, see [Encrypt Secrets at rest by using China KMS](https://help.aliyun.com/document_detail/177372.html).
//
//   - The user or role that calls this API operation must be granted additional cluster RBAC permissions (O&M engineer or administrator permissions). Otherwise, the ForbiddenUpdateKMSState error code is returned.
//
//   - After this API operation is successfully called, the cluster status changes to updating. After the update is complete, the cluster status changes back to running. After a change is complete for a cluster, wait at least 5 minutes before calling this API operation again. Otherwise, HTTP status code 409 is returned.
//
// @param request - UpdateKMSEncryptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKMSEncryptionResponse
func (client *Client) UpdateKMSEncryptionWithOptions(ClusterId *string, request *UpdateKMSEncryptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKMSEncryptionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the Secret encryption at rest configuration for a specified cluster by cluster ID.
//
// Description:
//
// While enabling or disabling encryption at rest and after the feature is enabled, do not disable or delete the KMS key used by this feature in the KMS console or through OpenAPI. Otherwise, the cluster API Server becomes unavailable, which prevents normal retrieval of objects such as Secrets and ServiceAccounts and affects the normal operation of business applications. For more information, see [Encrypt Secrets at rest by using China KMS](https://help.aliyun.com/document_detail/177372.html).
//
//   - The user or role that calls this API operation must be granted additional cluster RBAC permissions (O&M engineer or administrator permissions). Otherwise, the ForbiddenUpdateKMSState error code is returned.
//
//   - After this API operation is successfully called, the cluster status changes to updating. After the update is complete, the cluster status changes back to running. After a change is complete for a cluster, wait at least 5 minutes before calling this API operation again. Otherwise, HTTP status code 409 is returned.
//
// @param request - UpdateKMSEncryptionRequest
//
// @return UpdateKMSEncryptionResponse
func (client *Client) UpdateKMSEncryption(ClusterId *string, request *UpdateKMSEncryptionRequest) (_result *UpdateKMSEncryptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKMSEncryptionResponse{}
	_body, _err := client.UpdateKMSEncryptionWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// Updates a node component.
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
// Updates the deletion protection status of a specified resource. Currently supported resource types include namespaces and services.
//
// You can call this operation to enable deletion protection for namespaces or services that involve critical business or sensitive data to avoid maintenance costs caused by accidental deletion.
//
// Description:
//
// Before calling this operation, install or upgrade the security policy component for the cluster. For more information, see [Enable security policy management](https://help.aliyun.com/document_detail/359818.html).
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
// Updates the deletion protection status of a specified resource. Currently supported resource types include namespaces and services.
//
// You can call this operation to enable deletion protection for namespaces or services that involve critical business or sensitive data to avoid maintenance costs caused by accidental deletion.
//
// Description:
//
// Before calling this operation, install or upgrade the security policy component for the cluster. For more information, see [Enable security policy management](https://help.aliyun.com/document_detail/359818.html).
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
// An orchestration template defines and describes a set of Kubernetes cluster resources in a declarative manner, specifying how applications should run or be configured. Calls the UpdateTemplate operation to update an orchestration template configuration.
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
// An orchestration template defines and describes a set of Kubernetes cluster resources in a declarative manner, specifying how applications should run or be configured. Calls the UpdateTemplate operation to update an orchestration template configuration.
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
// In an ACK cluster, non-cluster creators, Resource Access Management (RAM) users, and RAM roles have no RBAC permissions by default. You can invoke the UpdateUserPermissions operation to update the RBAC access permissions of a RAM user or role, including accessible resources, permission scope, and preset role types, to better manage cluster management and secure access control.
//
// Description:
//
// You can update the cluster authorization information of a target Resource Access Management (RAM) user or RAM role by using full update or incremental update. A full update overwrites all existing cluster permissions of the target RAM user or RAM role. The request must include all permission configurations that you want to grant to the target RAM user or RAM role. An incremental update includes add and delete operations. Only the cluster authorization information included in the request is changed, and other cluster permissions of the RAM user or RAM role are not affected.
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
// In an ACK cluster, non-cluster creators, Resource Access Management (RAM) users, and RAM roles have no RBAC permissions by default. You can invoke the UpdateUserPermissions operation to update the RBAC access permissions of a RAM user or role, including accessible resources, permission scope, and preset role types, to better manage cluster management and secure access control.
//
// Description:
//
// You can update the cluster authorization information of a target Resource Access Management (RAM) user or RAM role by using full update or incremental update. A full update overwrites all existing cluster permissions of the target RAM user or RAM role. The request must include all permission configurations that you want to grant to the target RAM user or RAM role. An incremental update includes add and delete operations. Only the cluster authorization information included in the request is changed, and other cluster permissions of the RAM user or RAM role are not affected.
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
// To avoid potential security and stability risks of expired cluster versions and to use new features of the latest cluster versions, upgrade your clusters in accordance with the ACK cluster version release schedule. You can call the UpgradeCluster operation to manually upgrade a cluster.
//
// Description:
//
// After you successfully call the UpgradeCluster operation, the API returns the `task_id` of the upgrade task. You can manage the task by calling the following task API operations:
//
// - [Call DescribeTaskInfo to query task details](https://help.aliyun.com/document_detail/2667985.html)
//
// - [Call PauseTask to pause a running task](https://help.aliyun.com/document_detail/2667986.html)
//
// - [Call ResumeTask to resume a paused task](https://help.aliyun.com/document_detail/2667987.html)
//
// - [Call CancelTask to cancel a running task](https://help.aliyun.com/document_detail/2667988.html).
//
//	Notice: Starting July 4, 2026, the request parameters rolling_policy and rolling_policy.max_parallelism will no longer take effect. Use [UpgradeClusterNodepool](https://help.aliyun.com/document_detail/2667922.html) to upgrade worker nodes instead. For more information about the changes, see [Notice on changes to ACK cluster management OpenAPI request and response parameters and OpenAPI deprecation](https://help.aliyun.com/document_detail/2932733.html).</notice>.
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
// To avoid potential security and stability risks of expired cluster versions and to use new features of the latest cluster versions, upgrade your clusters in accordance with the ACK cluster version release schedule. You can call the UpgradeCluster operation to manually upgrade a cluster.
//
// Description:
//
// After you successfully call the UpgradeCluster operation, the API returns the `task_id` of the upgrade task. You can manage the task by calling the following task API operations:
//
// - [Call DescribeTaskInfo to query task details](https://help.aliyun.com/document_detail/2667985.html)
//
// - [Call PauseTask to pause a running task](https://help.aliyun.com/document_detail/2667986.html)
//
// - [Call ResumeTask to resume a paused task](https://help.aliyun.com/document_detail/2667987.html)
//
// - [Call CancelTask to cancel a running task](https://help.aliyun.com/document_detail/2667988.html).
//
//	Notice: Starting July 4, 2026, the request parameters rolling_policy and rolling_policy.max_parallelism will no longer take effect. Use [UpgradeClusterNodepool](https://help.aliyun.com/document_detail/2667922.html) to upgrade worker nodes instead. For more information about the changes, see [Notice on changes to ACK cluster management OpenAPI request and response parameters and OpenAPI deprecation](https://help.aliyun.com/document_detail/2932733.html).</notice>.
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
// Upgrades the versions of cluster component instances so that you can benefit from the feature optimizations in the new versions.
//
// Description:
//
// - Upgrading cluster component instance versions may affect your services. Assess the impact before performing the upgrade during off-peak hours, and back up relevant data in advance.
//
// - Before upgrading a component, refer to [Component release notes](https://help.aliyun.com/document_detail/176087.html) to learn about the changes and their impact for the specified component.
//
// - Upgrade components one at a time. Confirm that one component has been upgraded successfully before upgrading the next one.
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
// Upgrades the versions of cluster component instances so that you can benefit from the feature optimizations in the new versions.
//
// Description:
//
// - Upgrading cluster component instance versions may affect your services. Assess the impact before performing the upgrade during off-peak hours, and back up relevant data in advance.
//
// - Before upgrading a component, refer to [Component release notes](https://help.aliyun.com/document_detail/176087.html) to learn about the changes and their impact for the specified component.
//
// - Upgrade components one at a time. Confirm that one component has been upgraded successfully before upgrading the next one.
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
// Upgrades the kubelet version (recommended to match the control plane version), operating system version, or container runtime version of a specified cluster node pool.
//
// Description:
//
// Upgrades the Kubernetes version, operating system version, or container runtime version of nodes in a specified cluster node pool. After you call the UpgradeClusterNodepool operation, the API returns a task_id for the upgrade task. You can manage the task by calling the following task API operations:
//
// - [Call DescribeTaskInfo to query task details](https://help.aliyun.com/document_detail/2667985.html)
//
// - [Call PauseTask to pause a running task](https://help.aliyun.com/document_detail/2667986.html)
//
// - [Call ResumeTask to resume a paused task](https://help.aliyun.com/document_detail/2667987.html)
//
// - [Call CancelTask to cancel a running task](https://help.aliyun.com/document_detail/2667988.html).
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
// Upgrades the kubelet version (recommended to match the control plane version), operating system version, or container runtime version of a specified cluster node pool.
//
// Description:
//
// Upgrades the Kubernetes version, operating system version, or container runtime version of nodes in a specified cluster node pool. After you call the UpgradeClusterNodepool operation, the API returns a task_id for the upgrade task. You can manage the task by calling the following task API operations:
//
// - [Call DescribeTaskInfo to query task details](https://help.aliyun.com/document_detail/2667985.html)
//
// - [Call PauseTask to pause a running task](https://help.aliyun.com/document_detail/2667986.html)
//
// - [Call ResumeTask to resume a paused task](https://help.aliyun.com/document_detail/2667987.html)
//
// - [Call CancelTask to cancel a running task](https://help.aliyun.com/document_detail/2667988.html).
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
