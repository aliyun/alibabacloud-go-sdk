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
	client.Endpoint, _err = client.GetEndpoint(dara.String("ehpc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds Elastic Compute Service (ECS) instances as compute nodes to Elastic High Performance Computing (E-HPC) clusters.
//
// Description:
//
// The ECS instances must meet the following requirements:
//
//   - The ECS instances do not belong to any E-HPC cluster.
//
//   - The ECS instances reside in the same virtual private cloud (VPC) as the cluster.
//
//   - The ECS instances are in the Stopped state.
//
// Take of the following limits:
//
//   - You can specify multiple instance IDs to add them at a time. However, the instances must be of the same type.
//
//   - When an instance is added to the cluster, [the system disk is reset](https://help.aliyun.com/zh/ecs/user-guide/re-initialize-a-system-disk) by using the image specified by the input parameters.
//
//   - If the instance has data disks, they are not automatically created and mounted after the instance is added.
//
//   - The hostname of the instance remains the same. Therefore, you must ensure that the hostname of the instance to be added is different from the hostname of an existing node in the cluster.
//
// @param tmpReq - AttachNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachNodesResponse
func (client *Client) AttachNodesWithOptions(tmpReq *AttachNodesRequest, runtime *dara.RuntimeOptions) (_result *AttachNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AttachNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ComputeNode) {
		request.ComputeNodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComputeNode, dara.String("ComputeNode"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ComputeNodeShrink) {
		query["ComputeNode"] = request.ComputeNodeShrink
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachNodes"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds Elastic Compute Service (ECS) instances as compute nodes to Elastic High Performance Computing (E-HPC) clusters.
//
// Description:
//
// The ECS instances must meet the following requirements:
//
//   - The ECS instances do not belong to any E-HPC cluster.
//
//   - The ECS instances reside in the same virtual private cloud (VPC) as the cluster.
//
//   - The ECS instances are in the Stopped state.
//
// Take of the following limits:
//
//   - You can specify multiple instance IDs to add them at a time. However, the instances must be of the same type.
//
//   - When an instance is added to the cluster, [the system disk is reset](https://help.aliyun.com/zh/ecs/user-guide/re-initialize-a-system-disk) by using the image specified by the input parameters.
//
//   - If the instance has data disks, they are not automatically created and mounted after the instance is added.
//
//   - The hostname of the instance remains the same. Therefore, you must ensure that the hostname of the instance to be added is different from the hostname of an existing node in the cluster.
//
// @param request - AttachNodesRequest
//
// @return AttachNodesResponse
func (client *Client) AttachNodes(request *AttachNodesRequest) (_result *AttachNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachNodesResponse{}
	_body, _err := client.AttachNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches shared storage to an Elastic High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// When you call this operation, take note of the following items:
//
//   - The file system that you want to attach must be created in advance in the same virtual private cloud (VPC) as the destination cluster. For more information, see [Create a file system](https://help.aliyun.com/document_detail/27530.html) and [Manage mount targets](https://help.aliyun.com/document_detail/27531.html).
//
//   - E-HPC clusters support Apsara File Storage NAS file systems.
//
// @param tmpReq - AttachSharedStoragesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachSharedStoragesResponse
func (client *Client) AttachSharedStoragesWithOptions(tmpReq *AttachSharedStoragesRequest, runtime *dara.RuntimeOptions) (_result *AttachSharedStoragesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AttachSharedStoragesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SharedStorages) {
		request.SharedStoragesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SharedStorages, dara.String("SharedStorages"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.SharedStoragesShrink) {
		query["SharedStorages"] = request.SharedStoragesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachSharedStorages"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachSharedStoragesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches shared storage to an Elastic High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// When you call this operation, take note of the following items:
//
//   - The file system that you want to attach must be created in advance in the same virtual private cloud (VPC) as the destination cluster. For more information, see [Create a file system](https://help.aliyun.com/document_detail/27530.html) and [Manage mount targets](https://help.aliyun.com/document_detail/27531.html).
//
//   - E-HPC clusters support Apsara File Storage NAS file systems.
//
// @param request - AttachSharedStoragesRequest
//
// @return AttachSharedStoragesResponse
func (client *Client) AttachSharedStorages(request *AttachSharedStoragesRequest) (_result *AttachSharedStoragesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachSharedStoragesResponse{}
	_body, _err := client.AttachSharedStoragesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go or subscription Elastic High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you call this operation, make sure that you are familiar with the billing and pricing of E-HPC. For more information, see [Overview](https://help.aliyun.com/document_detail/2842985.html).
//
// @param tmpReq - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithOptions(tmpReq *CreateClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdditionalPackages) {
		request.AdditionalPackagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdditionalPackages, dara.String("AdditionalPackages"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Addons) {
		request.AddonsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Addons, dara.String("Addons"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ClusterCredentials) {
		request.ClusterCredentialsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterCredentials, dara.String("ClusterCredentials"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ClusterCustomConfiguration) {
		request.ClusterCustomConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterCustomConfiguration, dara.String("ClusterCustomConfiguration"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Manager) {
		request.ManagerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Manager, dara.String("Manager"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Queues) {
		request.QueuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Queues, dara.String("Queues"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SharedStorages) {
		request.SharedStoragesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SharedStorages, dara.String("SharedStorages"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalPackagesShrink) {
		query["AdditionalPackages"] = request.AdditionalPackagesShrink
	}

	if !dara.IsNil(request.AddonsShrink) {
		query["Addons"] = request.AddonsShrink
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.ClusterCategory) {
		query["ClusterCategory"] = request.ClusterCategory
	}

	if !dara.IsNil(request.ClusterCredentialsShrink) {
		query["ClusterCredentials"] = request.ClusterCredentialsShrink
	}

	if !dara.IsNil(request.ClusterCustomConfigurationShrink) {
		query["ClusterCustomConfiguration"] = request.ClusterCustomConfigurationShrink
	}

	if !dara.IsNil(request.ClusterDescription) {
		query["ClusterDescription"] = request.ClusterDescription
	}

	if !dara.IsNil(request.ClusterMode) {
		query["ClusterMode"] = request.ClusterMode
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ClusterVSwitchId) {
		query["ClusterVSwitchId"] = request.ClusterVSwitchId
	}

	if !dara.IsNil(request.ClusterVpcId) {
		query["ClusterVpcId"] = request.ClusterVpcId
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.IsEnterpriseSecurityGroup) {
		query["IsEnterpriseSecurityGroup"] = request.IsEnterpriseSecurityGroup
	}

	if !dara.IsNil(request.ManagerShrink) {
		query["Manager"] = request.ManagerShrink
	}

	if !dara.IsNil(request.MaxCoreCount) {
		query["MaxCoreCount"] = request.MaxCoreCount
	}

	if !dara.IsNil(request.MaxCount) {
		query["MaxCount"] = request.MaxCount
	}

	if !dara.IsNil(request.QueuesShrink) {
		query["Queues"] = request.QueuesShrink
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SharedStoragesShrink) {
		query["SharedStorages"] = request.SharedStoragesShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCluster"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Creates a pay-as-you-go or subscription Elastic High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you call this operation, make sure that you are familiar with the billing and pricing of E-HPC. For more information, see [Overview](https://help.aliyun.com/document_detail/2842985.html).
//
// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a job for a cluster.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing and [pricing](https://www.aliyun.com/price/product#/ecs/detail) of E-HPC.
//
// @param tmpReq - CreateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithOptions(tmpReq *CreateJobRequest, runtime *dara.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobSpec) {
		request.JobSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSpec, dara.String("JobSpec"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.JobSpecShrink) {
		query["JobSpec"] = request.JobSpecShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJob"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a job for a cluster.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing and [pricing](https://www.aliyun.com/price/product#/ecs/detail) of E-HPC.
//
// @param request - CreateJobRequest
//
// @return CreateJobResponse
func (client *Client) CreateJob(request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates compute nodes for an Elastic High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)
//
// @param tmpReq - CreateNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodesResponse
func (client *Client) CreateNodesWithOptions(tmpReq *CreateNodesRequest, runtime *dara.RuntimeOptions) (_result *CreateNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ComputeNode) {
		request.ComputeNodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComputeNode, dara.String("ComputeNode"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Hostnames) {
		request.HostnamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hostnames, dara.String("Hostnames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ComputeNodeShrink) {
		query["ComputeNode"] = request.ComputeNodeShrink
	}

	if !dara.IsNil(request.Count) {
		query["Count"] = request.Count
	}

	if !dara.IsNil(request.DeploymentSetId) {
		query["DeploymentSetId"] = request.DeploymentSetId
	}

	if !dara.IsNil(request.HPCInterConnect) {
		query["HPCInterConnect"] = request.HPCInterConnect
	}

	if !dara.IsNil(request.HostnamePrefix) {
		query["HostnamePrefix"] = request.HostnamePrefix
	}

	if !dara.IsNil(request.HostnameSuffix) {
		query["HostnameSuffix"] = request.HostnameSuffix
	}

	if !dara.IsNil(request.HostnamesShrink) {
		query["Hostnames"] = request.HostnamesShrink
	}

	if !dara.IsNil(request.KeepAlive) {
		query["KeepAlive"] = request.KeepAlive
	}

	if !dara.IsNil(request.MinCount) {
		query["MinCount"] = request.MinCount
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.RamRole) {
		query["RamRole"] = request.RamRole
	}

	if !dara.IsNil(request.ReservedNodePoolId) {
		query["ReservedNodePoolId"] = request.ReservedNodePoolId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNodes"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates compute nodes for an Elastic High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)
//
// @param request - CreateNodesRequest
//
// @return CreateNodesResponse
func (client *Client) CreateNodes(request *CreateNodesRequest) (_result *CreateNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNodesResponse{}
	_body, _err := client.CreateNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a queue for an Enterprise High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - CreateQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQueueResponse
func (client *Client) CreateQueueWithOptions(tmpReq *CreateQueueRequest, runtime *dara.RuntimeOptions) (_result *CreateQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateQueueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Queue) {
		request.QueueShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Queue, dara.String("Queue"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.QueueShrink) {
		query["Queue"] = request.QueueShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQueue"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a queue for an Enterprise High Performance Computing (E-HPC) cluster.
//
// @param request - CreateQueueRequest
//
// @return CreateQueueResponse
func (client *Client) CreateQueue(request *CreateQueueRequest) (_result *CreateQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQueueResponse{}
	_body, _err := client.CreateQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建预设节点池
//
// @param request - CreateReservedNodePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReservedNodePoolResponse
func (client *Client) CreateReservedNodePoolWithOptions(request *CreateReservedNodePoolRequest, runtime *dara.RuntimeOptions) (_result *CreateReservedNodePoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Count) {
		query["Count"] = request.Count
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.HostPostfix) {
		query["HostPostfix"] = request.HostPostfix
	}

	if !dara.IsNil(request.HostPrefix) {
		query["HostPrefix"] = request.HostPrefix
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateReservedNodePool"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateReservedNodePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建预设节点池
//
// @param request - CreateReservedNodePoolRequest
//
// @return CreateReservedNodePoolResponse
func (client *Client) CreateReservedNodePool(request *CreateReservedNodePoolRequest) (_result *CreateReservedNodePoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateReservedNodePoolResponse{}
	_body, _err := client.CreateReservedNodePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds users to an Elastic High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - CreateUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUsersResponse
func (client *Client) CreateUsersWithOptions(tmpReq *CreateUsersRequest, runtime *dara.RuntimeOptions) (_result *CreateUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.User) {
		request.UserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.User, dara.String("User"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.UserShrink) {
		query["User"] = request.UserShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUsers"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds users to an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - CreateUsersRequest
//
// @return CreateUsersResponse
func (client *Client) CreateUsers(request *CreateUsersRequest) (_result *CreateUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUsersResponse{}
	_body, _err := client.CreateUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Make sure that data of the cluster to be deleted is backed up before you call this operation.
//
// > After a cluster is released, you cannot restore the data stored in the cluster. Exercise caution when you release a cluster.
//
// @param request - DeleteClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCluster"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Releases an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Make sure that data of the cluster to be deleted is backed up before you call this operation.
//
// > After a cluster is released, you cannot restore the data stored in the cluster. Exercise caution when you release a cluster.
//
// @param request - DeleteClusterRequest
//
// @return DeleteClusterResponse
func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes compute nodes from an Enterprise High Performance Computing (E-HPC) cluster at a time.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you delete a compute node, we recommend that you export all job data from the node to prevent data loss.
//
// @param tmpReq - DeleteNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodesResponse
func (client *Client) DeleteNodesWithOptions(tmpReq *DeleteNodesRequest, runtime *dara.RuntimeOptions) (_result *DeleteNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNodes"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes compute nodes from an Enterprise High Performance Computing (E-HPC) cluster at a time.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you delete a compute node, we recommend that you export all job data from the node to prevent data loss.
//
// @param request - DeleteNodesRequest
//
// @return DeleteNodesResponse
func (client *Client) DeleteNodes(request *DeleteNodesRequest) (_result *DeleteNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNodesResponse{}
	_body, _err := client.DeleteNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes queues from an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you delete a queue, you must delete all compute nodes in the queue.
//
// @param tmpReq - DeleteQueuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQueuesResponse
func (client *Client) DeleteQueuesWithOptions(tmpReq *DeleteQueuesRequest, runtime *dara.RuntimeOptions) (_result *DeleteQueuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteQueuesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueueNames) {
		request.QueueNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueueNames, dara.String("QueueNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.QueueNamesShrink) {
		query["QueueNames"] = request.QueueNamesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQueues"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQueuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes queues from an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you delete a queue, you must delete all compute nodes in the queue.
//
// @param request - DeleteQueuesRequest
//
// @return DeleteQueuesResponse
func (client *Client) DeleteQueues(request *DeleteQueuesRequest) (_result *DeleteQueuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQueuesResponse{}
	_body, _err := client.DeleteQueuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除集群预设节点池
//
// @param request - DeleteReservedNodePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReservedNodePoolResponse
func (client *Client) DeleteReservedNodePoolWithOptions(request *DeleteReservedNodePoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteReservedNodePoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteReservedNodePool"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteReservedNodePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除集群预设节点池
//
// @param request - DeleteReservedNodePoolRequest
//
// @return DeleteReservedNodePoolResponse
func (client *Client) DeleteReservedNodePool(request *DeleteReservedNodePoolRequest) (_result *DeleteReservedNodePoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteReservedNodePoolResponse{}
	_body, _err := client.DeleteReservedNodePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes users from a cluster.
//
// @param tmpReq - DeleteUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUsersResponse
func (client *Client) DeleteUsersWithOptions(tmpReq *DeleteUsersRequest, runtime *dara.RuntimeOptions) (_result *DeleteUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.User) {
		request.UserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.User, dara.String("User"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUsers"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes users from a cluster.
//
// @param request - DeleteUsersRequest
//
// @return DeleteUsersResponse
func (client *Client) DeleteUsers(request *DeleteUsersRequest) (_result *DeleteUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUsersResponse{}
	_body, _err := client.DeleteUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an addon template.
//
// @param request - DescribeAddonTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAddonTemplateResponse
func (client *Client) DescribeAddonTemplateWithOptions(request *DescribeAddonTemplateRequest, runtime *dara.RuntimeOptions) (_result *DescribeAddonTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["AddonName"] = request.AddonName
	}

	if !dara.IsNil(request.AddonVersion) {
		query["AddonVersion"] = request.AddonVersion
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

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAddonTemplate"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAddonTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an addon template.
//
// @param request - DescribeAddonTemplateRequest
//
// @return DescribeAddonTemplateResponse
func (client *Client) DescribeAddonTemplate(request *DescribeAddonTemplateRequest) (_result *DescribeAddonTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAddonTemplateResponse{}
	_body, _err := client.DescribeAddonTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unmounts shared storage from the mount directory of a cluster.
//
// @param tmpReq - DetachSharedStoragesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachSharedStoragesResponse
func (client *Client) DetachSharedStoragesWithOptions(tmpReq *DetachSharedStoragesRequest, runtime *dara.RuntimeOptions) (_result *DetachSharedStoragesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetachSharedStoragesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SharedStorages) {
		request.SharedStoragesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SharedStorages, dara.String("SharedStorages"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.SharedStoragesShrink) {
		query["SharedStorages"] = request.SharedStoragesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachSharedStorages"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachSharedStoragesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unmounts shared storage from the mount directory of a cluster.
//
// @param request - DetachSharedStoragesRequest
//
// @return DetachSharedStoragesResponse
func (client *Client) DetachSharedStorages(request *DetachSharedStoragesRequest) (_result *DetachSharedStoragesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachSharedStoragesResponse{}
	_body, _err := client.DetachSharedStoragesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an installed addon.
//
// @param request - GetAddonRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddonResponse
func (client *Client) GetAddonWithOptions(request *GetAddonRequest, runtime *dara.RuntimeOptions) (_result *GetAddonResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonId) {
		query["AddonId"] = request.AddonId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAddon"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAddonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an installed addon.
//
// @param request - GetAddonRequest
//
// @return GetAddonResponse
func (client *Client) GetAddon(request *GetAddonRequest) (_result *GetAddonResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAddonResponse{}
	_body, _err := client.GetAddonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - GetClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterResponse
func (client *Client) GetClusterWithOptions(request *GetClusterRequest, runtime *dara.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCluster"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - GetClusterRequest
//
// @return GetClusterResponse
func (client *Client) GetCluster(request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query logs based on a request ID. Logs for specific actions can be queried thanks to an Action-Stage-Method three-layer log splitting structure.
//
// @param request - GetCommonLogDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommonLogDetailResponse
func (client *Client) GetCommonLogDetailWithOptions(request *GetCommonLogDetailRequest, runtime *dara.RuntimeOptions) (_result *GetCommonLogDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.HiddenProcess) {
		query["HiddenProcess"] = request.HiddenProcess
	}

	if !dara.IsNil(request.LogRequestId) {
		query["LogRequestId"] = request.LogRequestId
	}

	if !dara.IsNil(request.To) {
		query["To"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCommonLogDetail"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCommonLogDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query logs based on a request ID. Logs for specific actions can be queried thanks to an Action-Stage-Method three-layer log splitting structure.
//
// @param request - GetCommonLogDetailRequest
//
// @return GetCommonLogDetailResponse
func (client *Client) GetCommonLogDetail(request *GetCommonLogDetailRequest) (_result *GetCommonLogDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCommonLogDetailResponse{}
	_body, _err := client.GetCommonLogDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a job.
//
// @param request - GetJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithOptions(request *GetJobRequest, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJob"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a job.
//
// @param request - GetJobRequest
//
// @return GetJobResponse
func (client *Client) GetJob(request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the output logs of a job, including standard output logs and error output logs.
//
// Description:
//
// ## [](#)Usage notes
//
// Currently, only Slurm and PBS Pro schedulers for Standard Edition clusters are supported.
//
// @param request - GetJobLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobLogResponse
func (client *Client) GetJobLogWithOptions(request *GetJobLogRequest, runtime *dara.RuntimeOptions) (_result *GetJobLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobLog"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the output logs of a job, including standard output logs and error output logs.
//
// Description:
//
// ## [](#)Usage notes
//
// Currently, only Slurm and PBS Pro schedulers for Standard Edition clusters are supported.
//
// @param request - GetJobLogRequest
//
// @return GetJobLogResponse
func (client *Client) GetJobLog(request *GetJobLogRequest) (_result *GetJobLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobLogResponse{}
	_body, _err := client.GetJobLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a queue in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - GetQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueResponse
func (client *Client) GetQueueWithOptions(request *GetQueueRequest, runtime *dara.RuntimeOptions) (_result *GetQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueue"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a queue in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - GetQueueRequest
//
// @return GetQueueResponse
func (client *Client) GetQueue(request *GetQueueRequest) (_result *GetQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueueResponse{}
	_body, _err := client.GetQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs an addon.
//
// Description:
//
// ## [](#)Usage notes
//
// Take note of the following items when you call this operation:
//
//   - The cluster must be in the `Running` state.
//
//   - Clusters fall into two types:
//
//   - Regular clusters on Alibaba Cloud Public Cloud
//
//   - Managed clusters on Alibaba Cloud Public Cloud
//
// @param request - InstallAddonRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAddonResponse
func (client *Client) InstallAddonWithOptions(request *InstallAddonRequest, runtime *dara.RuntimeOptions) (_result *InstallAddonResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["AddonName"] = request.AddonName
	}

	if !dara.IsNil(request.AddonVersion) {
		query["AddonVersion"] = request.AddonVersion
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ResourcesSpec) {
		query["ResourcesSpec"] = request.ResourcesSpec
	}

	if !dara.IsNil(request.ServicesSpec) {
		query["ServicesSpec"] = request.ServicesSpec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallAddon"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallAddonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an addon.
//
// Description:
//
// ## [](#)Usage notes
//
// Take note of the following items when you call this operation:
//
//   - The cluster must be in the `Running` state.
//
//   - Clusters fall into two types:
//
//   - Regular clusters on Alibaba Cloud Public Cloud
//
//   - Managed clusters on Alibaba Cloud Public Cloud
//
// @param request - InstallAddonRequest
//
// @return InstallAddonResponse
func (client *Client) InstallAddon(request *InstallAddonRequest) (_result *InstallAddonResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InstallAddonResponse{}
	_body, _err := client.InstallAddonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Install software for the specified cluster.
//
// Description:
//
// ## Interface Description
//
// When calling this interface, please note the following:
//
// - The cluster status must be `Running`.
//
// - If the cluster series is `Serverless`, ensure that there is at least one login node or compute node in the cluster; otherwise, software cannot be added to the target cluster.
//
// @param tmpReq - InstallSoftwaresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallSoftwaresResponse
func (client *Client) InstallSoftwaresWithOptions(tmpReq *InstallSoftwaresRequest, runtime *dara.RuntimeOptions) (_result *InstallSoftwaresResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InstallSoftwaresShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdditionalPackages) {
		request.AdditionalPackagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdditionalPackages, dara.String("AdditionalPackages"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallSoftwares"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallSoftwaresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Install software for the specified cluster.
//
// Description:
//
// ## Interface Description
//
// When calling this interface, please note the following:
//
// - The cluster status must be `Running`.
//
// - If the cluster series is `Serverless`, ensure that there is at least one login node or compute node in the cluster; otherwise, software cannot be added to the target cluster.
//
// @param request - InstallSoftwaresRequest
//
// @return InstallSoftwaresResponse
func (client *Client) InstallSoftwares(request *InstallSoftwaresRequest) (_result *InstallSoftwaresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InstallSoftwaresResponse{}
	_body, _err := client.InstallSoftwaresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries supported addon templates.
//
// @param request - ListAddonTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonTemplatesResponse
func (client *Client) ListAddonTemplatesWithOptions(request *ListAddonTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListAddonTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonNames) {
		query["AddonNames"] = request.AddonNames
	}

	if !dara.IsNil(request.ClusterCategory) {
		query["ClusterCategory"] = request.ClusterCategory
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
		Action:      dara.String("ListAddonTemplates"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAddonTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries supported addon templates.
//
// @param request - ListAddonTemplatesRequest
//
// @return ListAddonTemplatesResponse
func (client *Client) ListAddonTemplates(request *ListAddonTemplatesRequest) (_result *ListAddonTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAddonTemplatesResponse{}
	_body, _err := client.ListAddonTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries installed addons of an Elastic High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - ListAddonsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonsResponse
func (client *Client) ListAddonsWithOptions(tmpReq *ListAddonsRequest, runtime *dara.RuntimeOptions) (_result *ListAddonsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAddonsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddonIds) {
		request.AddonIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddonIds, dara.String("AddonIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonIdsShrink) {
		query["AddonIds"] = request.AddonIdsShrink
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAddons"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Queries installed addons of an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - ListAddonsRequest
//
// @return ListAddonsResponse
func (client *Client) ListAddons(request *ListAddonsRequest) (_result *ListAddonsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAddonsResponse{}
	_body, _err := client.ListAddonsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the file systems that can be attached in a region.
//
// @param request - ListAvailableFileSystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableFileSystemsResponse
func (client *Client) ListAvailableFileSystemsWithOptions(request *ListAvailableFileSystemsRequest, runtime *dara.RuntimeOptions) (_result *ListAvailableFileSystemsResponse, _err error) {
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
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAvailableFileSystems"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAvailableFileSystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the file systems that can be attached in a region.
//
// @param request - ListAvailableFileSystemsRequest
//
// @return ListAvailableFileSystemsResponse
func (client *Client) ListAvailableFileSystems(request *ListAvailableFileSystemsRequest) (_result *ListAvailableFileSystemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAvailableFileSystemsResponse{}
	_body, _err := client.ListAvailableFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries images that are available for Elastic High Performance Computing (E-HPC) clusters.
//
// @param tmpReq - ListAvailableImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableImagesResponse
func (client *Client) ListAvailableImagesWithOptions(tmpReq *ListAvailableImagesRequest, runtime *dara.RuntimeOptions) (_result *ListAvailableImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAvailableImagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DirectoryService) {
		request.DirectoryServiceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DirectoryService, dara.String("DirectoryService"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Scheduler) {
		request.SchedulerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scheduler, dara.String("Scheduler"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAvailableImages"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAvailableImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries images that are available for Elastic High Performance Computing (E-HPC) clusters.
//
// @param request - ListAvailableImagesRequest
//
// @return ListAvailableImagesResponse
func (client *Client) ListAvailableImages(request *ListAvailableImagesRequest) (_result *ListAvailableImagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAvailableImagesResponse{}
	_body, _err := client.ListAvailableImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all clusters of a user in each region.
//
// @param tmpReq - ListClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(tmpReq *ListClustersRequest, runtime *dara.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListClustersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ClusterIds) {
		request.ClusterIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterIds, dara.String("ClusterIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ClusterNames) {
		request.ClusterNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterNames, dara.String("ClusterNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterIdsShrink) {
		query["ClusterIds"] = request.ClusterIdsShrink
	}

	if !dara.IsNil(request.ClusterNamesShrink) {
		query["ClusterNames"] = request.ClusterNamesShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusters"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all clusters of a user in each region.
//
// @param request - ListClustersRequest
//
// @return ListClustersResponse
func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logs of a cluster that are generated within a time range.
//
// @param tmpReq - ListCommonLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCommonLogsResponse
func (client *Client) ListCommonLogsWithOptions(tmpReq *ListCommonLogsRequest, runtime *dara.RuntimeOptions) (_result *ListCommonLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCommonLogsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ActionName) {
		request.ActionNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionName, dara.String("ActionName"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionNameShrink) {
		query["ActionName"] = request.ActionNameShrink
	}

	if !dara.IsNil(request.ActionStatus) {
		query["ActionStatus"] = request.ActionStatus
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.IsReverse) {
		query["IsReverse"] = request.IsReverse
	}

	if !dara.IsNil(request.LogRequestId) {
		query["LogRequestId"] = request.LogRequestId
	}

	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.OperatorUid) {
		query["OperatorUid"] = request.OperatorUid
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.To) {
		query["To"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCommonLogs"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCommonLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of a cluster that are generated within a time range.
//
// @param request - ListCommonLogsRequest
//
// @return ListCommonLogsResponse
func (client *Client) ListCommonLogs(request *ListCommonLogsRequest) (_result *ListCommonLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCommonLogsResponse{}
	_body, _err := client.ListCommonLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the installed software of a cluster.
//
// @param request - ListInstalledSoftwaresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstalledSoftwaresResponse
func (client *Client) ListInstalledSoftwaresWithOptions(request *ListInstalledSoftwaresRequest, runtime *dara.RuntimeOptions) (_result *ListInstalledSoftwaresResponse, _err error) {
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
		Action:      dara.String("ListInstalledSoftwares"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstalledSoftwaresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the installed software of a cluster.
//
// @param request - ListInstalledSoftwaresRequest
//
// @return ListInstalledSoftwaresResponse
func (client *Client) ListInstalledSoftwares(request *ListInstalledSoftwaresRequest) (_result *ListInstalledSoftwaresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstalledSoftwaresResponse{}
	_body, _err := client.ListInstalledSoftwaresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the jobs in a cluster.
//
// @param tmpReq - ListJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithOptions(tmpReq *ListJobsRequest, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobFilter) {
		request.JobFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobFilter, dara.String("JobFilter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobFilterShrink) {
		query["JobFilter"] = request.JobFilterShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the jobs in a cluster.
//
// @param request - ListJobsRequest
//
// @return ListJobsResponse
func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the nodes of an Elastic High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(tmpReq *ListNodesRequest, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Hostnames) {
		request.HostnamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hostnames, dara.String("Hostnames"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PrivateIpAddress) {
		request.PrivateIpAddressShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PrivateIpAddress, dara.String("PrivateIpAddress"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.QueueNames) {
		request.QueueNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueueNames, dara.String("QueueNames"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Status) {
		request.StatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Status, dara.String("Status"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.HostnamesShrink) {
		query["Hostnames"] = request.HostnamesShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PrivateIpAddressShrink) {
		query["PrivateIpAddress"] = request.PrivateIpAddressShrink
	}

	if !dara.IsNil(request.QueueNamesShrink) {
		query["QueueNames"] = request.QueueNamesShrink
	}

	if !dara.IsNil(request.Sequence) {
		query["Sequence"] = request.Sequence
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StatusShrink) {
		query["Status"] = request.StatusShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodes"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the nodes of an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries queues in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - ListQueuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueuesResponse
func (client *Client) ListQueuesWithOptions(tmpReq *ListQueuesRequest, runtime *dara.RuntimeOptions) (_result *ListQueuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListQueuesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueueNames) {
		request.QueueNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueueNames, dara.String("QueueNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.QueueNamesShrink) {
		query["QueueNames"] = request.QueueNamesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQueues"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQueuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries queues in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - ListQueuesRequest
//
// @return ListQueuesResponse
func (client *Client) ListQueues(request *ListQueuesRequest) (_result *ListQueuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQueuesResponse{}
	_body, _err := client.ListQueuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询产品支持的地域列表。
//
// @param request - ListRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(request *ListRegionsRequest, runtime *dara.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SpecCode) {
		query["SpecCode"] = request.SpecCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegions"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询产品支持的地域列表。
//
// @param request - ListRegionsRequest
//
// @return ListRegionsResponse
func (client *Client) ListRegions(request *ListRegionsRequest) (_result *ListRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the shared storage that is attached to a cluster.
//
// @param request - ListSharedStoragesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSharedStoragesResponse
func (client *Client) ListSharedStoragesWithOptions(request *ListSharedStoragesRequest, runtime *dara.RuntimeOptions) (_result *ListSharedStoragesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSharedStorages"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSharedStoragesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the shared storage that is attached to a cluster.
//
// @param request - ListSharedStoragesRequest
//
// @return ListSharedStoragesResponse
func (client *Client) ListSharedStorages(request *ListSharedStoragesRequest) (_result *ListSharedStoragesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSharedStoragesResponse{}
	_body, _err := client.ListSharedStoragesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the software that can be installed in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - ListSoftwaresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSoftwaresResponse
func (client *Client) ListSoftwaresWithOptions(request *ListSoftwaresRequest, runtime *dara.RuntimeOptions) (_result *ListSoftwaresResponse, _err error) {
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
		Action:      dara.String("ListSoftwares"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSoftwaresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the software that can be installed in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - ListSoftwaresRequest
//
// @return ListSoftwaresResponse
func (client *Client) ListSoftwares(request *ListSoftwaresRequest) (_result *ListSoftwaresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSoftwaresResponse{}
	_body, _err := client.ListSoftwaresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the users of a cluster.
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
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
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the users of a cluster.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops uncompleted jobs in a batch in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - StopJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopJobsResponse
func (client *Client) StopJobsWithOptions(tmpReq *StopJobsRequest, runtime *dara.RuntimeOptions) (_result *StopJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StopJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobIds) {
		request.JobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIds, dara.String("JobIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobIdsShrink) {
		query["JobIds"] = request.JobIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopJobs"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops uncompleted jobs in a batch in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - StopJobsRequest
//
// @return StopJobsResponse
func (client *Client) StopJobs(request *StopJobsRequest) (_result *StopJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopJobsResponse{}
	_body, _err := client.StopJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls an addon.
//
// Description:
//
// ## [](#)Usage notes
//
// Take note of the following items when you call this operation:
//
//   - The cluster must be in the `Running` state.
//
//   - Clusters fall into the following types:
//
//   - Regular clusters on Alibaba Cloud Public Cloud
//
//   - Managed clusters on Alibaba Cloud Public Cloud
//
// @param request - UnInstallAddonRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnInstallAddonResponse
func (client *Client) UnInstallAddonWithOptions(request *UnInstallAddonRequest, runtime *dara.RuntimeOptions) (_result *UnInstallAddonResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonId) {
		query["AddonId"] = request.AddonId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnInstallAddon"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnInstallAddonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls an addon.
//
// Description:
//
// ## [](#)Usage notes
//
// Take note of the following items when you call this operation:
//
//   - The cluster must be in the `Running` state.
//
//   - Clusters fall into the following types:
//
//   - Regular clusters on Alibaba Cloud Public Cloud
//
//   - Managed clusters on Alibaba Cloud Public Cloud
//
// @param request - UnInstallAddonRequest
//
// @return UnInstallAddonResponse
func (client *Client) UnInstallAddon(request *UnInstallAddonRequest) (_result *UnInstallAddonResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnInstallAddonResponse{}
	_body, _err := client.UnInstallAddonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls software systems from an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## Interface Description
//
// When calling this interface, please note:
//
// The cluster status must be `Running`.
//
// @param tmpReq - UninstallSoftwaresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallSoftwaresResponse
func (client *Client) UninstallSoftwaresWithOptions(tmpReq *UninstallSoftwaresRequest, runtime *dara.RuntimeOptions) (_result *UninstallSoftwaresResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UninstallSoftwaresShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdditionalPackages) {
		request.AdditionalPackagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdditionalPackages, dara.String("AdditionalPackages"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallSoftwares"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallSoftwaresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls software systems from an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## Interface Description
//
// When calling this interface, please note:
//
// The cluster status must be `Running`.
//
// @param request - UninstallSoftwaresRequest
//
// @return UninstallSoftwaresResponse
func (client *Client) UninstallSoftwares(request *UninstallSoftwaresRequest) (_result *UninstallSoftwaresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UninstallSoftwaresResponse{}
	_body, _err := client.UninstallSoftwaresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify the basic information of a specified cluster.
//
// @param tmpReq - UpdateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterResponse
func (client *Client) UpdateClusterWithOptions(tmpReq *UpdateClusterRequest, runtime *dara.RuntimeOptions) (_result *UpdateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ClusterCustomConfiguration) {
		request.ClusterCustomConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterCustomConfiguration, dara.String("ClusterCustomConfiguration"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MonitorSpec) {
		request.MonitorSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MonitorSpec, dara.String("MonitorSpec"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SchedulerSpec) {
		request.SchedulerSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SchedulerSpec, dara.String("SchedulerSpec"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.ClusterCustomConfigurationShrink) {
		query["ClusterCustomConfiguration"] = request.ClusterCustomConfigurationShrink
	}

	if !dara.IsNil(request.ClusterDescription) {
		query["ClusterDescription"] = request.ClusterDescription
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.EnableScaleIn) {
		query["EnableScaleIn"] = request.EnableScaleIn
	}

	if !dara.IsNil(request.EnableScaleOut) {
		query["EnableScaleOut"] = request.EnableScaleOut
	}

	if !dara.IsNil(request.GrowInterval) {
		query["GrowInterval"] = request.GrowInterval
	}

	if !dara.IsNil(request.IdleInterval) {
		query["IdleInterval"] = request.IdleInterval
	}

	if !dara.IsNil(request.MaxCoreCount) {
		query["MaxCoreCount"] = request.MaxCoreCount
	}

	if !dara.IsNil(request.MaxCount) {
		query["MaxCount"] = request.MaxCount
	}

	if !dara.IsNil(request.MonitorSpecShrink) {
		query["MonitorSpec"] = request.MonitorSpecShrink
	}

	if !dara.IsNil(request.SchedulerSpecShrink) {
		query["SchedulerSpec"] = request.SchedulerSpecShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCluster"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the basic information of a specified cluster.
//
// @param request - UpdateClusterRequest
//
// @return UpdateClusterResponse
func (client *Client) UpdateCluster(request *UpdateClusterRequest) (_result *UpdateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateClusterResponse{}
	_body, _err := client.UpdateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of compute nodes in an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you delete a compute node, we recommend that you export all job data from the node to prevent data loss.
//
// @param tmpReq - UpdateNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodesResponse
func (client *Client) UpdateNodesWithOptions(tmpReq *UpdateNodesRequest, runtime *dara.RuntimeOptions) (_result *UpdateNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Instances) {
		request.InstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Instances, dara.String("Instances"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstancesShrink) {
		query["Instances"] = request.InstancesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNodes"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of compute nodes in an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you delete a compute node, we recommend that you export all job data from the node to prevent data loss.
//
// @param request - UpdateNodesRequest
//
// @return UpdateNodesResponse
func (client *Client) UpdateNodes(request *UpdateNodesRequest) (_result *UpdateNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNodesResponse{}
	_body, _err := client.UpdateNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a queue in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - UpdateQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQueueResponse
func (client *Client) UpdateQueueWithOptions(tmpReq *UpdateQueueRequest, runtime *dara.RuntimeOptions) (_result *UpdateQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateQueueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Queue) {
		request.QueueShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Queue, dara.String("Queue"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.QueueShrink) {
		query["Queue"] = request.QueueShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQueue"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a queue in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - UpdateQueueRequest
//
// @return UpdateQueueResponse
func (client *Client) UpdateQueue(request *UpdateQueueRequest) (_result *UpdateQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateQueueResponse{}
	_body, _err := client.UpdateQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information of a user in an Elastic High Performance Computing (E-HPC) cluster, including the user group and password.
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2024-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information of a user in an Elastic High Performance Computing (E-HPC) cluster, including the user group and password.
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
