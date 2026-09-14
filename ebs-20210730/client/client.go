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
	client.Endpoint, _err = client.GetEndpoint(dara.String("ebs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a replication pair to a replication pair-consistent group. You can use a replication pair-consistent group to batch manage replication pairs. When you call this operation, you can specify parameters, such as ReplicaGroupId, ReplicaPairId, and ClientToken, in the request.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - A replication pair and a replication pair-consistent group replicate in the same direction if they have the same primary region (production region), primary zone (production zone), secondary region (disaster recovery region), and secondary zone (disaster recovery zone). A replication pair can be added only to a replication pair-consistent group that replicates in the same direction as the replication pair.
//
//   - Before you can add a replication pair to a replication pair-consistent group, make sure that the pair and the group are in the **Created*	- (`created`) or **Stopped*	- (`stopped`) state.
//
//   - Up to 17 replication pairs can be added to a single replication pair-consistent group.
//
//   - After replication pairs are added to a replication pair-consistent group, the recovery point objective (RPO) of the group takes effect on the pairs in place of their original RPOs.
//
// @param request - AddDiskReplicaPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDiskReplicaPairResponse
func (client *Client) AddDiskReplicaPairWithOptions(request *AddDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *AddDiskReplicaPairResponse, _err error) {
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

	if !dara.IsNil(request.ReplicaGroupId) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	if !dara.IsNil(request.ReplicaPairId) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDiskReplicaPair"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a replication pair to a replication pair-consistent group. You can use a replication pair-consistent group to batch manage replication pairs. When you call this operation, you can specify parameters, such as ReplicaGroupId, ReplicaPairId, and ClientToken, in the request.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - A replication pair and a replication pair-consistent group replicate in the same direction if they have the same primary region (production region), primary zone (production zone), secondary region (disaster recovery region), and secondary zone (disaster recovery zone). A replication pair can be added only to a replication pair-consistent group that replicates in the same direction as the replication pair.
//
//   - Before you can add a replication pair to a replication pair-consistent group, make sure that the pair and the group are in the **Created*	- (`created`) or **Stopped*	- (`stopped`) state.
//
//   - Up to 17 replication pairs can be added to a single replication pair-consistent group.
//
//   - After replication pairs are added to a replication pair-consistent group, the recovery point objective (RPO) of the group takes effect on the pairs in place of their original RPOs.
//
// @param request - AddDiskReplicaPairRequest
//
// @return AddDiskReplicaPairResponse
func (client *Client) AddDiskReplicaPair(request *AddDiskReplicaPairRequest) (_result *AddDiskReplicaPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDiskReplicaPairResponse{}
	_body, _err := client.AddDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates EBS Lens (CloudLens for EBS).
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyLensServiceResponse
func (client *Client) ApplyLensServiceWithOptions(runtime *dara.RuntimeOptions) (_result *ApplyLensServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyLensService"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyLensServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates EBS Lens (CloudLens for EBS).
//
// @return ApplyLensServiceResponse
func (client *Client) ApplyLensService() (_result *ApplyLensServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyLensServiceResponse{}
	_body, _err := client.ApplyLensServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Bind disks into a enterprise-level snapshot policy.
//
// @param request - BindEnterpriseSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindEnterpriseSnapshotPolicyResponse
func (client *Client) BindEnterpriseSnapshotPolicyWithOptions(request *BindEnterpriseSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *BindEnterpriseSnapshotPolicyResponse, _err error) {
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

	if !dara.IsNil(request.DiskTargets) {
		query["DiskTargets"] = request.DiskTargets
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
		Action:      dara.String("BindEnterpriseSnapshotPolicy"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Bind disks into a enterprise-level snapshot policy.
//
// @param request - BindEnterpriseSnapshotPolicyRequest
//
// @return BindEnterpriseSnapshotPolicyResponse
func (client *Client) BindEnterpriseSnapshotPolicy(request *BindEnterpriseSnapshotPolicyRequest) (_result *BindEnterpriseSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.BindEnterpriseSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels EBS Lens.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelLensServiceResponse
func (client *Client) CancelLensServiceWithOptions(runtime *dara.RuntimeOptions) (_result *CancelLensServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("CancelLensService"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelLensServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels EBS Lens.
//
// @return CancelLensServiceResponse
func (client *Client) CancelLensService() (_result *CancelLensServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelLensServiceResponse{}
	_body, _err := client.CancelLensServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which an Elastic Block Storage (EBS) resource belongs.
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
		Version:     dara.String("2021-07-30"),
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
// Changes the resource group to which an Elastic Block Storage (EBS) resource belongs.
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
// Clears the disaster recovery drills that were initiated from the secondary disk of a replication pair and deletes the auto-created drill disks.
//
// @param request - ClearPairDrillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearPairDrillResponse
func (client *Client) ClearPairDrillWithOptions(request *ClearPairDrillRequest, runtime *dara.RuntimeOptions) (_result *ClearPairDrillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrillId) {
		query["DrillId"] = request.DrillId
	}

	if !dara.IsNil(request.PairId) {
		query["PairId"] = request.PairId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearPairDrill"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearPairDrillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the disaster recovery drills that were initiated from the secondary disk of a replication pair and deletes the auto-created drill disks.
//
// @param request - ClearPairDrillRequest
//
// @return ClearPairDrillResponse
func (client *Client) ClearPairDrill(request *ClearPairDrillRequest) (_result *ClearPairDrillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClearPairDrillResponse{}
	_body, _err := client.ClearPairDrillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clears the disaster recovery drills that were initiated from the secondary disks of a replication pair-consistent group and deletes the auto-created drill disks.
//
// @param request - ClearReplicaGroupDrillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearReplicaGroupDrillResponse
func (client *Client) ClearReplicaGroupDrillWithOptions(request *ClearReplicaGroupDrillRequest, runtime *dara.RuntimeOptions) (_result *ClearReplicaGroupDrillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrillId) {
		query["DrillId"] = request.DrillId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearReplicaGroupDrill"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearReplicaGroupDrillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the disaster recovery drills that were initiated from the secondary disks of a replication pair-consistent group and deletes the auto-created drill disks.
//
// @param request - ClearReplicaGroupDrillRequest
//
// @return ClearReplicaGroupDrillResponse
func (client *Client) ClearReplicaGroupDrill(request *ClearReplicaGroupDrillRequest) (_result *ClearReplicaGroupDrillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClearReplicaGroupDrillResponse{}
	_body, _err := client.ClearReplicaGroupDrillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an app with the centralized role.
//
// @param request - CreateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *dara.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppTags) {
		query["AppTags"] = request.AppTags
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReportSendEnabled) {
		query["ReportSendEnabled"] = request.ReportSendEnabled
	}

	if !dara.IsNil(request.SubscribePeriod) {
		query["SubscribePeriod"] = request.SubscribePeriod
	}

	if !dara.IsNil(request.SubscribeStatus) {
		query["SubscribeStatus"] = request.SubscribeStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApp"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an app with the centralized role.
//
// @param request - CreateAppRequest
//
// @return CreateAppResponse
func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a dedicated block storage cluster by calling the CreateDedicatedBlockStorageCluster operation with parameters such as Azone, Capacity, Type, and PeriodUnit.
//
// Description:
//
// ## Operation description
//
// - Dedicated Block Storage Cluster is a block storage service that is physically isolated from other public cloud block storage clusters and provides exclusive access to all cluster resources for the owner.<props="china"> For more information, see [What is Dedicated Block Storage Cluster](https://help.aliyun.com/document_detail/208883.html).
//
// - Cloud disks created on a dedicated block storage cluster can be attached only to ECS instances in the same zone. Before you create a dedicated block storage cluster, plan the region and zone for the resources.
//
// - Dedicated block storage clusters are classified into basic and performance types. Select the appropriate cluster performance type based on your business requirements when you create a cluster.
//
// - You are charged for creating a dedicated block storage cluster.<props="china"> For more information, see [Billing](https://help.aliyun.com/document_detail/208884.html).
//
// @param request - CreateDedicatedBlockStorageClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDedicatedBlockStorageClusterResponse
func (client *Client) CreateDedicatedBlockStorageClusterWithOptions(request *CreateDedicatedBlockStorageClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateDedicatedBlockStorageClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Azone) {
		query["Azone"] = request.Azone
	}

	if !dara.IsNil(request.Capacity) {
		query["Capacity"] = request.Capacity
	}

	if !dara.IsNil(request.DbscId) {
		query["DbscId"] = request.DbscId
	}

	if !dara.IsNil(request.DbscName) {
		query["DbscName"] = request.DbscName
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
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
		Action:      dara.String("CreateDedicatedBlockStorageCluster"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDedicatedBlockStorageClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dedicated block storage cluster by calling the CreateDedicatedBlockStorageCluster operation with parameters such as Azone, Capacity, Type, and PeriodUnit.
//
// Description:
//
// ## Operation description
//
// - Dedicated Block Storage Cluster is a block storage service that is physically isolated from other public cloud block storage clusters and provides exclusive access to all cluster resources for the owner.<props="china"> For more information, see [What is Dedicated Block Storage Cluster](https://help.aliyun.com/document_detail/208883.html).
//
// - Cloud disks created on a dedicated block storage cluster can be attached only to ECS instances in the same zone. Before you create a dedicated block storage cluster, plan the region and zone for the resources.
//
// - Dedicated block storage clusters are classified into basic and performance types. Select the appropriate cluster performance type based on your business requirements when you create a cluster.
//
// - You are charged for creating a dedicated block storage cluster.<props="china"> For more information, see [Billing](https://help.aliyun.com/document_detail/208884.html).
//
// @param request - CreateDedicatedBlockStorageClusterRequest
//
// @return CreateDedicatedBlockStorageClusterResponse
func (client *Client) CreateDedicatedBlockStorageCluster(request *CreateDedicatedBlockStorageClusterRequest) (_result *CreateDedicatedBlockStorageClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDedicatedBlockStorageClusterResponse{}
	_body, _err := client.CreateDedicatedBlockStorageClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a diagnostic report for a specified resource.
//
// @param request - CreateDiagnoseReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiagnoseReportResponse
func (client *Client) CreateDiagnoseReportWithOptions(request *CreateDiagnoseReportRequest, runtime *dara.RuntimeOptions) (_result *CreateDiagnoseReportResponse, _err error) {
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

	if !dara.IsNil(request.DiagnoseType) {
		query["DiagnoseType"] = request.DiagnoseType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDiagnoseReport"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDiagnoseReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a diagnostic report for a specified resource.
//
// @param request - CreateDiagnoseReportRequest
//
// @return CreateDiagnoseReportResponse
func (client *Client) CreateDiagnoseReport(request *CreateDiagnoseReportRequest) (_result *CreateDiagnoseReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDiagnoseReportResponse{}
	_body, _err := client.CreateDiagnoseReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a replication pair-consistent group.
//
// Description:
//
// ## Operation Description
//
// Replication pair-consistent groups help you manage asynchronous replication for multiple disks in disaster recovery scenarios. You can centrally manage operations for these disks. The groups ensure that data on all disks can be recovered to the same point in time. This provides disaster recovery protection for one or more instances.
//
// When you create a replication pair-consistent group, note the following:
//
// - For information about the regions that support replication pair-consistent groups, see [Async replication overview](https://help.aliyun.com/document_detail/314563.html).
//
// - Replication pair-consistent groups support asynchronous disaster recovery across zones in the same region or across regions.
//
// - A replication pair can be added to a replication pair-consistent group only if the pair and the group have the same data replication direction. This means their production regions, production zones, disaster recovery regions, and disaster recovery zones must be the same.
//
// - After a replication pair is added to a replication pair-consistent group, the original Recovery Point Object (RPO) of the pair becomes invalid. Data is then replicated based on the RPO of the group.
//
// @param request - CreateDiskReplicaGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiskReplicaGroupResponse
func (client *Client) CreateDiskReplicaGroupWithOptions(request *CreateDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDiskReplicaGroupResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationRegionId) {
		query["DestinationRegionId"] = request.DestinationRegionId
	}

	if !dara.IsNil(request.DestinationZoneId) {
		query["DestinationZoneId"] = request.DestinationZoneId
	}

	if !dara.IsNil(request.EnableRtc) {
		query["EnableRtc"] = request.EnableRtc
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.RPO) {
		query["RPO"] = request.RPO
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceZoneId) {
		query["SourceZoneId"] = request.SourceZoneId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDiskReplicaGroup"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a replication pair-consistent group.
//
// Description:
//
// ## Operation Description
//
// Replication pair-consistent groups help you manage asynchronous replication for multiple disks in disaster recovery scenarios. You can centrally manage operations for these disks. The groups ensure that data on all disks can be recovered to the same point in time. This provides disaster recovery protection for one or more instances.
//
// When you create a replication pair-consistent group, note the following:
//
// - For information about the regions that support replication pair-consistent groups, see [Async replication overview](https://help.aliyun.com/document_detail/314563.html).
//
// - Replication pair-consistent groups support asynchronous disaster recovery across zones in the same region or across regions.
//
// - A replication pair can be added to a replication pair-consistent group only if the pair and the group have the same data replication direction. This means their production regions, production zones, disaster recovery regions, and disaster recovery zones must be the same.
//
// - After a replication pair is added to a replication pair-consistent group, the original Recovery Point Object (RPO) of the pair becomes invalid. Data is then replicated based on the RPO of the group.
//
// @param request - CreateDiskReplicaGroupRequest
//
// @return CreateDiskReplicaGroupResponse
func (client *Client) CreateDiskReplicaGroup(request *CreateDiskReplicaGroupRequest) (_result *CreateDiskReplicaGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDiskReplicaGroupResponse{}
	_body, _err := client.CreateDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a disk replication pair.
//
// Description:
//
// ## Description
//
// <props="china">
//
// Asynchronous replication is a feature that provides cross-region data protection using the data replication capabilities of Elastic Block Storage. The feature asynchronously replicates data from a disk in one region to another disk in a different region to provide cross-region disaster recovery and backup for your data. You can use this feature to build disaster recovery capabilities for critical services, protect database data, and improve business continuity. For more information about the feature notes and billing, see [Overview of asynchronous replication for disks](https://help.aliyun.com/document_detail/314563.html). The asynchronous replication feature is billed for bandwidth on a subscription basis. Pay-as-you-go billing for traffic is also supported.
//
// <props="intl">
//
// Async replication is a feature that provides cross-region data protection based on the data replication capabilities of Elastic Block Storage (EBS). The feature asynchronously replicates data from a disk in one region to a disk in another region for disaster recovery. You can use this feature to build disaster recovery capabilities for critical services, protect database data, and improve business continuity. The async replication feature is billed for bandwidth on a subscription basis.
//
// Currently, the asynchronous replication feature is available only for Enhanced SSD (ESSD) disks and has certain limitations.
//
// Before you create a replication pair, note the following:
//
// - Prepare a source disk (primary disk) and a destination disk (secondary disk). The source disk is the disk that you want to protect with disaster recovery, and the destination disk is the replication target. You can call the [CreateDisk](https://help.aliyun.com/document_detail/25513.html) operation to create the disks.
//
// - The secondary disk must be in a different region from the primary disk. For more information about the regions that support asynchronous replication, see [Overview of asynchronous replication for disks](https://help.aliyun.com/document_detail/314563.html).
//
// - This operation creates a replication pair between the primary and secondary disks. After you create the pair, you must call the [StartDiskReplicaPair](https://help.aliyun.com/document_detail/354205.html) operation to activate it. Activating the replication pair starts the periodic, cross-region data replication from the primary disk to the secondary disk.
//
// @param request - CreateDiskReplicaPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiskReplicaPairResponse
func (client *Client) CreateDiskReplicaPairWithOptions(request *CreateDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *CreateDiskReplicaPairResponse, _err error) {
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

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationDiskId) {
		query["DestinationDiskId"] = request.DestinationDiskId
	}

	if !dara.IsNil(request.DestinationRegionId) {
		query["DestinationRegionId"] = request.DestinationRegionId
	}

	if !dara.IsNil(request.DestinationZoneId) {
		query["DestinationZoneId"] = request.DestinationZoneId
	}

	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.EnableRtc) {
		query["EnableRtc"] = request.EnableRtc
	}

	if !dara.IsNil(request.PairName) {
		query["PairName"] = request.PairName
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.RPO) {
		query["RPO"] = request.RPO
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceZoneId) {
		query["SourceZoneId"] = request.SourceZoneId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDiskReplicaPair"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a disk replication pair.
//
// Description:
//
// ## Description
//
// <props="china">
//
// Asynchronous replication is a feature that provides cross-region data protection using the data replication capabilities of Elastic Block Storage. The feature asynchronously replicates data from a disk in one region to another disk in a different region to provide cross-region disaster recovery and backup for your data. You can use this feature to build disaster recovery capabilities for critical services, protect database data, and improve business continuity. For more information about the feature notes and billing, see [Overview of asynchronous replication for disks](https://help.aliyun.com/document_detail/314563.html). The asynchronous replication feature is billed for bandwidth on a subscription basis. Pay-as-you-go billing for traffic is also supported.
//
// <props="intl">
//
// Async replication is a feature that provides cross-region data protection based on the data replication capabilities of Elastic Block Storage (EBS). The feature asynchronously replicates data from a disk in one region to a disk in another region for disaster recovery. You can use this feature to build disaster recovery capabilities for critical services, protect database data, and improve business continuity. The async replication feature is billed for bandwidth on a subscription basis.
//
// Currently, the asynchronous replication feature is available only for Enhanced SSD (ESSD) disks and has certain limitations.
//
// Before you create a replication pair, note the following:
//
// - Prepare a source disk (primary disk) and a destination disk (secondary disk). The source disk is the disk that you want to protect with disaster recovery, and the destination disk is the replication target. You can call the [CreateDisk](https://help.aliyun.com/document_detail/25513.html) operation to create the disks.
//
// - The secondary disk must be in a different region from the primary disk. For more information about the regions that support asynchronous replication, see [Overview of asynchronous replication for disks](https://help.aliyun.com/document_detail/314563.html).
//
// - This operation creates a replication pair between the primary and secondary disks. After you create the pair, you must call the [StartDiskReplicaPair](https://help.aliyun.com/document_detail/354205.html) operation to activate it. Activating the replication pair starts the periodic, cross-region data replication from the primary disk to the secondary disk.
//
// @param request - CreateDiskReplicaPairRequest
//
// @return CreateDiskReplicaPairResponse
func (client *Client) CreateDiskReplicaPair(request *CreateDiskReplicaPairRequest) (_result *CreateDiskReplicaPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDiskReplicaPairResponse{}
	_body, _err := client.CreateDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an enterprise-level snapshot policy by calling CreateEnterpriseSnapshotPolicy.
//
// @param tmpReq - CreateEnterpriseSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnterpriseSnapshotPolicyResponse
func (client *Client) CreateEnterpriseSnapshotPolicyWithOptions(tmpReq *CreateEnterpriseSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateEnterpriseSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateEnterpriseSnapshotPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CrossRegionCopyInfo) {
		request.CrossRegionCopyInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CrossRegionCopyInfo, dara.String("CrossRegionCopyInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RetainRule) {
		request.RetainRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RetainRule, dara.String("RetainRule"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Schedule) {
		request.ScheduleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Schedule, dara.String("Schedule"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SpecialRetainRules) {
		request.SpecialRetainRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SpecialRetainRules, dara.String("SpecialRetainRules"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StorageRule) {
		request.StorageRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StorageRule, dara.String("StorageRule"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossRegionCopyInfoShrink) {
		query["CrossRegionCopyInfo"] = request.CrossRegionCopyInfoShrink
	}

	if !dara.IsNil(request.Desc) {
		query["Desc"] = request.Desc
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RetainRuleShrink) {
		query["RetainRule"] = request.RetainRuleShrink
	}

	if !dara.IsNil(request.ScheduleShrink) {
		query["Schedule"] = request.ScheduleShrink
	}

	if !dara.IsNil(request.SpecialRetainRulesShrink) {
		query["SpecialRetainRules"] = request.SpecialRetainRulesShrink
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.StorageRuleShrink) {
		query["StorageRule"] = request.StorageRuleShrink
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnterpriseSnapshotPolicy"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an enterprise-level snapshot policy by calling CreateEnterpriseSnapshotPolicy.
//
// @param request - CreateEnterpriseSnapshotPolicyRequest
//
// @return CreateEnterpriseSnapshotPolicyResponse
func (client *Client) CreateEnterpriseSnapshotPolicy(request *CreateEnterpriseSnapshotPolicyRequest) (_result *CreateEnterpriseSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CreateEnterpriseSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application using the centralized role.
//
// @param request - DeleteAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApp"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application using the centralized role.
//
// @param request - DeleteAppRequest
//
// @return DeleteAppResponse
func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a replication pair-consistent group.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - Before you can delete a replication pair-consistent group, make sure that no replication pairs exist in the group.
//
//   - The replication pair-consistent group that you want to delete must be in the **Created*	- (`created`), **Creation Failed*	- (`create_failed`), **Stopped*	- (`stopped`), **Failovered*	- (`failovered`), **Deleting*	- (`deleting`), **Deletion Failed*	- (`delete_failed`), or **Invalid*	- (`invalid`) state.
//
// @param request - DeleteDiskReplicaGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDiskReplicaGroupResponse
func (client *Client) DeleteDiskReplicaGroupWithOptions(request *DeleteDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDiskReplicaGroupResponse, _err error) {
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

	if !dara.IsNil(request.ReplicaGroupId) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDiskReplicaGroup"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a replication pair-consistent group.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - Before you can delete a replication pair-consistent group, make sure that no replication pairs exist in the group.
//
//   - The replication pair-consistent group that you want to delete must be in the **Created*	- (`created`), **Creation Failed*	- (`create_failed`), **Stopped*	- (`stopped`), **Failovered*	- (`failovered`), **Deleting*	- (`deleting`), **Deletion Failed*	- (`delete_failed`), or **Invalid*	- (`invalid`) state.
//
// @param request - DeleteDiskReplicaGroupRequest
//
// @return DeleteDiskReplicaGroupResponse
func (client *Client) DeleteDiskReplicaGroup(request *DeleteDiskReplicaGroupRequest) (_result *DeleteDiskReplicaGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDiskReplicaGroupResponse{}
	_body, _err := client.DeleteDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes replication pairs.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which async replication is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - Only replication pairs that are in the **Stopped*	- (`stopped`), **Invalid*	- (`invalid`), or **Failovered*	- (`failovered`) state can be deleted. This operation deletes only replication pairs. The primary and secondary disks in the deleted replication pairs are retained.
//
//   - To delete a replication pair, you must call this operation in the region where the primary disk is located. After the replication pair is deleted, the functionality limits are lifted from the primary and secondary disks. For example, you can attach the secondary disk, resize the disk, or read data from or write data to the disk.
//
// @param request - DeleteDiskReplicaPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDiskReplicaPairResponse
func (client *Client) DeleteDiskReplicaPairWithOptions(request *DeleteDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *DeleteDiskReplicaPairResponse, _err error) {
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

	if !dara.IsNil(request.ReplicaPairId) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDiskReplicaPair"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes replication pairs.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which async replication is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - Only replication pairs that are in the **Stopped*	- (`stopped`), **Invalid*	- (`invalid`), or **Failovered*	- (`failovered`) state can be deleted. This operation deletes only replication pairs. The primary and secondary disks in the deleted replication pairs are retained.
//
//   - To delete a replication pair, you must call this operation in the region where the primary disk is located. After the replication pair is deleted, the functionality limits are lifted from the primary and secondary disks. For example, you can attach the secondary disk, resize the disk, or read data from or write data to the disk.
//
// @param request - DeleteDiskReplicaPairRequest
//
// @return DeleteDiskReplicaPairResponse
func (client *Client) DeleteDiskReplicaPair(request *DeleteDiskReplicaPairRequest) (_result *DeleteDiskReplicaPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDiskReplicaPairResponse{}
	_body, _err := client.DeleteDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a enterprise-level snapshot policy.
//
// @param request - DeleteEnterpriseSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnterpriseSnapshotPolicyResponse
func (client *Client) DeleteEnterpriseSnapshotPolicyWithOptions(request *DeleteEnterpriseSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnterpriseSnapshotPolicyResponse, _err error) {
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
		Action:      dara.String("DeleteEnterpriseSnapshotPolicy"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a enterprise-level snapshot policy.
//
// @param request - DeleteEnterpriseSnapshotPolicyRequest
//
// @return DeleteEnterpriseSnapshotPolicyResponse
func (client *Client) DeleteEnterpriseSnapshotPolicy(request *DeleteEnterpriseSnapshotPolicyRequest) (_result *DeleteEnterpriseSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.DeleteEnterpriseSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries application information as a centralized role.
//
// Description:
//
// ## Operation description
//
// Queries the tag keys and values of cloud disks and snapshots contained in a user-defined application. You can use filterTagKey to narrow the query scope.
//
// @param request - DescribeAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsResponse
func (client *Client) DescribeAppsWithOptions(request *DescribeAppsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.AppNames) {
		query["AppNames"] = request.AppNames
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApps"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries application information as a centralized role.
//
// Description:
//
// ## Operation description
//
// Queries the tag keys and values of cloud disks and snapshots contained in a user-defined application. You can use filterTagKey to narrow the query scope.
//
// @param request - DescribeAppsRequest
//
// @return DescribeAppsResponse
func (client *Client) DescribeApps(request *DescribeAppsRequest) (_result *DescribeAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppsResponse{}
	_body, _err := client.DescribeAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about one or more cloud disks that you have created in a dedicated block storage cluster.
//
// Description:
//
// - The following two methods are supported to view returned data:
//
//   - Method 1: Use `NextToken` to set a query token. Set its value to the `NextToken` value returned by the previous call to DescribeDisks, and use `MaxResults` to set the maximum number of entries per page.
//
//   - Method 2: Use `PageSize` to set the number of entries per page, and use `PageNumber` to set the page number.
//
//     You can use only one of the preceding methods. When a large number of entries are returned, we recommend that you use Method 1. If you set `NextToken`, the `PageSize` and `PageNumber` request parameters do not take effect, and `TotalCount` in the returned data is invalid.
//
// - Cloud disks with the multi-attach feature enabled can be attached to multiple instances. You can view all attachment information of a cloud disk based on the `Attachment` list in the response.
//
// When you invoke an API operation by using Cloud Assistant CLI, specify request parameters of different data types in the required formats. For more information, see [Parameter format of CLI](https://help.aliyun.com/document_detail/110340.html).
//
// @param request - DescribeDedicatedBlockStorageClusterDisksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDedicatedBlockStorageClusterDisksResponse
func (client *Client) DescribeDedicatedBlockStorageClusterDisksWithOptions(request *DescribeDedicatedBlockStorageClusterDisksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDedicatedBlockStorageClusterDisksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbscId) {
		query["DbscId"] = request.DbscId
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
		Action:      dara.String("DescribeDedicatedBlockStorageClusterDisks"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDedicatedBlockStorageClusterDisksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about one or more cloud disks that you have created in a dedicated block storage cluster.
//
// Description:
//
// - The following two methods are supported to view returned data:
//
//   - Method 1: Use `NextToken` to set a query token. Set its value to the `NextToken` value returned by the previous call to DescribeDisks, and use `MaxResults` to set the maximum number of entries per page.
//
//   - Method 2: Use `PageSize` to set the number of entries per page, and use `PageNumber` to set the page number.
//
//     You can use only one of the preceding methods. When a large number of entries are returned, we recommend that you use Method 1. If you set `NextToken`, the `PageSize` and `PageNumber` request parameters do not take effect, and `TotalCount` in the returned data is invalid.
//
// - Cloud disks with the multi-attach feature enabled can be attached to multiple instances. You can view all attachment information of a cloud disk based on the `Attachment` list in the response.
//
// When you invoke an API operation by using Cloud Assistant CLI, specify request parameters of different data types in the required formats. For more information, see [Parameter format of CLI](https://help.aliyun.com/document_detail/110340.html).
//
// @param request - DescribeDedicatedBlockStorageClusterDisksRequest
//
// @return DescribeDedicatedBlockStorageClusterDisksResponse
func (client *Client) DescribeDedicatedBlockStorageClusterDisks(request *DescribeDedicatedBlockStorageClusterDisksRequest) (_result *DescribeDedicatedBlockStorageClusterDisksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDedicatedBlockStorageClusterDisksResponse{}
	_body, _err := client.DescribeDedicatedBlockStorageClusterDisksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries dedicated clusters.
//
// Description:
//
// ## [](#)Usage notes
//
// >  The Dedicated Block Storage Cluster feature is available only in the China (Heyuan), Indonesia (Jakarta), and China (Shenzhen) regions.
//
//   - You can specify multiple request parameters to be queried. Specified parameters are evaluated by using the AND operator. Only the specified parameters are included in the filter conditions.
//
//   - We recommend that you use NextToken and MaxResults to perform paged queries. We recommend that you use MaxResults to specify the maximum number of entries to return in each request. The return value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. When you call the DescribeDedicatedBlockStorageClusters operation to retrieve a new page of results, set NextToken to the NextToken value that is returned in the previous call and specify MaxResults to limit the number of entries returned.
//
// @param request - DescribeDedicatedBlockStorageClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDedicatedBlockStorageClustersResponse
func (client *Client) DescribeDedicatedBlockStorageClustersWithOptions(request *DescribeDedicatedBlockStorageClustersRequest, runtime *dara.RuntimeOptions) (_result *DescribeDedicatedBlockStorageClustersResponse, _err error) {
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

	if !dara.IsNil(request.DedicatedBlockStorageClusterId) {
		query["DedicatedBlockStorageClusterId"] = request.DedicatedBlockStorageClusterId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AzoneId) {
		body["AzoneId"] = request.AzoneId
	}

	if !dara.IsNil(request.Category) {
		body["Category"] = request.Category
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDedicatedBlockStorageClusters"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDedicatedBlockStorageClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries dedicated clusters.
//
// Description:
//
// ## [](#)Usage notes
//
// >  The Dedicated Block Storage Cluster feature is available only in the China (Heyuan), Indonesia (Jakarta), and China (Shenzhen) regions.
//
//   - You can specify multiple request parameters to be queried. Specified parameters are evaluated by using the AND operator. Only the specified parameters are included in the filter conditions.
//
//   - We recommend that you use NextToken and MaxResults to perform paged queries. We recommend that you use MaxResults to specify the maximum number of entries to return in each request. The return value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. When you call the DescribeDedicatedBlockStorageClusters operation to retrieve a new page of results, set NextToken to the NextToken value that is returned in the previous call and specify MaxResults to limit the number of entries returned.
//
// @param request - DescribeDedicatedBlockStorageClustersRequest
//
// @return DescribeDedicatedBlockStorageClustersResponse
func (client *Client) DescribeDedicatedBlockStorageClusters(request *DescribeDedicatedBlockStorageClustersRequest) (_result *DescribeDedicatedBlockStorageClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDedicatedBlockStorageClustersResponse{}
	_body, _err := client.DescribeDedicatedBlockStorageClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries diagnostic reports.
//
// @param request - DescribeDiagnoseReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnoseReportResponse
func (client *Client) DescribeDiagnoseReportWithOptions(request *DescribeDiagnoseReportRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnoseReportResponse, _err error) {
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

	if !dara.IsNil(request.DiagnoseType) {
		query["DiagnoseType"] = request.DiagnoseType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.ReportIds) {
		query["ReportIds"] = request.ReportIds
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiagnoseReport"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiagnoseReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries diagnostic reports.
//
// @param request - DescribeDiagnoseReportRequest
//
// @return DescribeDiagnoseReportResponse
func (client *Client) DescribeDiagnoseReport(request *DescribeDiagnoseReportRequest) (_result *DescribeDiagnoseReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDiagnoseReportResponse{}
	_body, _err := client.DescribeDiagnoseReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries fine-grained monitoring data of a disk.
//
// Description:
//
// ## Usage notes
//
//   - CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
//
//   - Up to 400 monitoring data entries can be returned at a time. An error is returned if the value calculated based on the following formula is greater than 400: `(EndTime - StartTime)/Period`.
//
//   - You can query the monitoring data collected in the last three days. An error is returned if the time specified by `StartTime` is more than three days prior to the current time.
//
// @param request - DescribeDiskMonitorDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiskMonitorDataResponse
func (client *Client) DescribeDiskMonitorDataWithOptions(request *DescribeDiskMonitorDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiskMonitorDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiskMonitorData"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiskMonitorDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries fine-grained monitoring data of a disk.
//
// Description:
//
// ## Usage notes
//
//   - CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
//
//   - Up to 400 monitoring data entries can be returned at a time. An error is returned if the value calculated based on the following formula is greater than 400: `(EndTime - StartTime)/Period`.
//
//   - You can query the monitoring data collected in the last three days. An error is returned if the time specified by `StartTime` is more than three days prior to the current time.
//
// @param request - DescribeDiskMonitorDataRequest
//
// @return DescribeDiskMonitorDataResponse
func (client *Client) DescribeDiskMonitorData(request *DescribeDiskMonitorDataRequest) (_result *DescribeDiskMonitorDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDiskMonitorDataResponse{}
	_body, _err := client.DescribeDiskMonitorDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about one or more replication pair-consistent groups in a specified region.
//
// Description:
//
// ## Description
//
// To perform a paged query, you can use the MaxResults and NextToken parameters.
//
// When you query the first page, set `MaxResults` to specify the number of entries to return. The `NextToken` value in the response is the token used to query subsequent pages. To query a subsequent page, set `NextToken` to the value from the previous response and set MaxResults to specify the number of entries for the current page.
//
// @param request - DescribeDiskReplicaGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiskReplicaGroupsResponse
func (client *Client) DescribeDiskReplicaGroupsWithOptions(request *DescribeDiskReplicaGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiskReplicaGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
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

	if !dara.IsNil(request.Site) {
		query["Site"] = request.Site
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiskReplicaGroups"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiskReplicaGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about one or more replication pair-consistent groups in a specified region.
//
// Description:
//
// ## Description
//
// To perform a paged query, you can use the MaxResults and NextToken parameters.
//
// When you query the first page, set `MaxResults` to specify the number of entries to return. The `NextToken` value in the response is the token used to query subsequent pages. To query a subsequent page, set `NextToken` to the value from the previous response and set MaxResults to specify the number of entries for the current page.
//
// @param request - DescribeDiskReplicaGroupsRequest
//
// @return DescribeDiskReplicaGroupsResponse
func (client *Client) DescribeDiskReplicaGroups(request *DescribeDiskReplicaGroupsRequest) (_result *DescribeDiskReplicaGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDiskReplicaGroupsResponse{}
	_body, _err := client.DescribeDiskReplicaGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the replication progress of a replication pair.
//
// @param request - DescribeDiskReplicaPairProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiskReplicaPairProgressResponse
func (client *Client) DescribeDiskReplicaPairProgressWithOptions(request *DescribeDiskReplicaPairProgressRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiskReplicaPairProgressResponse, _err error) {
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

	if !dara.IsNil(request.ReplicaPairId) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiskReplicaPairProgress"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiskReplicaPairProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the replication progress of a replication pair.
//
// @param request - DescribeDiskReplicaPairProgressRequest
//
// @return DescribeDiskReplicaPairProgressResponse
func (client *Client) DescribeDiskReplicaPairProgress(request *DescribeDiskReplicaPairProgressRequest) (_result *DescribeDiskReplicaPairProgressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDiskReplicaPairProgressResponse{}
	_body, _err := client.DescribeDiskReplicaPairProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about one or more replication pairs in a specified region.
//
// Description:
//
// ## Description
//
// - For information about the regions that support asynchronous replication, see [Overview of asynchronous replication for disks](https://help.aliyun.com/document_detail/314563.html).
//
// - If the primary or secondary disk of an asynchronous replication pair is in the specified region, the query returns information about that replication pair.
//
// - To perform a paged query, use the MaxResults and NextToken parameters. For the first page, set MaxResults to limit the number of entries returned. The `NextToken` value in the response is the token for the next page. For subsequent pages, set `NextToken` to the value from the previous response and set `MaxResults` to limit the number of entries.
//
// @param request - DescribeDiskReplicaPairsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiskReplicaPairsResponse
func (client *Client) DescribeDiskReplicaPairsWithOptions(request *DescribeDiskReplicaPairsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiskReplicaPairsResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PairIds) {
		query["PairIds"] = request.PairIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplicaGroupId) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Site) {
		query["Site"] = request.Site
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiskReplicaPairs"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiskReplicaPairsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about one or more replication pairs in a specified region.
//
// Description:
//
// ## Description
//
// - For information about the regions that support asynchronous replication, see [Overview of asynchronous replication for disks](https://help.aliyun.com/document_detail/314563.html).
//
// - If the primary or secondary disk of an asynchronous replication pair is in the specified region, the query returns information about that replication pair.
//
// - To perform a paged query, use the MaxResults and NextToken parameters. For the first page, set MaxResults to limit the number of entries returned. The `NextToken` value in the response is the token for the next page. For subsequent pages, set `NextToken` to the value from the previous response and set `MaxResults` to limit the number of entries.
//
// @param request - DescribeDiskReplicaPairsRequest
//
// @return DescribeDiskReplicaPairsResponse
func (client *Client) DescribeDiskReplicaPairs(request *DescribeDiskReplicaPairsRequest) (_result *DescribeDiskReplicaPairsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDiskReplicaPairsResponse{}
	_body, _err := client.DescribeDiskReplicaPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries enterprise-level snapshot policies by specifying parameters such as PolicyIds, ResourceGroupId, and Tag.
//
// @param request - DescribeEnterpriseSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnterpriseSnapshotPolicyResponse
func (client *Client) DescribeEnterpriseSnapshotPolicyWithOptions(request *DescribeEnterpriseSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnterpriseSnapshotPolicyResponse, _err error) {
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

	if !dara.IsNil(request.DiskIds) {
		query["DiskIds"] = request.DiskIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyIds) {
		query["PolicyIds"] = request.PolicyIds
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
		Action:      dara.String("DescribeEnterpriseSnapshotPolicy"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries enterprise-level snapshot policies by specifying parameters such as PolicyIds, ResourceGroupId, and Tag.
//
// @param request - DescribeEnterpriseSnapshotPolicyRequest
//
// @return DescribeEnterpriseSnapshotPolicyResponse
func (client *Client) DescribeEnterpriseSnapshotPolicy(request *DescribeEnterpriseSnapshotPolicyRequest) (_result *DescribeEnterpriseSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.DescribeEnterpriseSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries cloud disk risk events.
//
// @param request - DescribeEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsResponse
func (client *Client) DescribeEventsWithOptions(request *DescribeEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
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

	if !dara.IsNil(request.EventLevel) {
		query["EventLevel"] = request.EventLevel
	}

	if !dara.IsNil(request.EventName) {
		query["EventName"] = request.EventName
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

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEvents"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Queries cloud disk risk events.
//
// @param request - DescribeEventsRequest
//
// @return DescribeEventsResponse
func (client *Client) DescribeEvents(request *DescribeEventsRequest) (_result *DescribeEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventsResponse{}
	_body, _err := client.DescribeEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more block storage cloud disks that you have created.
//
// @param request - DescribeLensMonitorDisksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLensMonitorDisksResponse
func (client *Client) DescribeLensMonitorDisksWithOptions(request *DescribeLensMonitorDisksRequest, runtime *dara.RuntimeOptions) (_result *DescribeLensMonitorDisksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskCategory) {
		query["DiskCategory"] = request.DiskCategory
	}

	if !dara.IsNil(request.DiskIdPattern) {
		query["DiskIdPattern"] = request.DiskIdPattern
	}

	if !dara.IsNil(request.DiskIds) {
		query["DiskIds"] = request.DiskIds
	}

	if !dara.IsNil(request.EcsInstanceId) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !dara.IsNil(request.LensTags) {
		query["LensTags"] = request.LensTags
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
		Action:      dara.String("DescribeLensMonitorDisks"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLensMonitorDisksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more block storage cloud disks that you have created.
//
// @param request - DescribeLensMonitorDisksRequest
//
// @return DescribeLensMonitorDisksResponse
func (client *Client) DescribeLensMonitorDisks(request *DescribeLensMonitorDisksRequest) (_result *DescribeLensMonitorDisksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLensMonitorDisksResponse{}
	_body, _err := client.DescribeLensMonitorDisksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the EBS Lens service activated by the user.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLensServiceStatusResponse
func (client *Client) DescribeLensServiceStatusWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeLensServiceStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLensServiceStatus"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLensServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the EBS Lens service activated by the user.
//
// @return DescribeLensServiceStatusResponse
func (client *Client) DescribeLensServiceStatus() (_result *DescribeLensServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLensServiceStatusResponse{}
	_body, _err := client.DescribeLensServiceStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the monitoring metrics of a resource.
//
// @param tmpReq - DescribeMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricDataResponse
func (client *Client) DescribeMetricDataWithOptions(tmpReq *DescribeMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeMetricDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GroupByLabels) {
		request.GroupByLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupByLabels, dara.String("GroupByLabels"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AggreOps) {
		query["AggreOps"] = request.AggreOps
	}

	if !dara.IsNil(request.AggreOverLineOps) {
		query["AggreOverLineOps"] = request.AggreOverLineOps
	}

	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupByLabelsShrink) {
		query["GroupByLabels"] = request.GroupByLabelsShrink
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricData"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring metrics of a resource.
//
// @param request - DescribeMetricDataRequest
//
// @return DescribeMetricDataResponse
func (client *Client) DescribeMetricData(request *DescribeMetricDataRequest) (_result *DescribeMetricDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricDataResponse{}
	_body, _err := client.DescribeMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the disaster recovery drills that were performed on the replication pair whose secondary disk resides in a specific region.
//
// @param request - DescribePairDrillsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePairDrillsResponse
func (client *Client) DescribePairDrillsWithOptions(request *DescribePairDrillsRequest, runtime *dara.RuntimeOptions) (_result *DescribePairDrillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrillId) {
		query["DrillId"] = request.DrillId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PairId) {
		query["PairId"] = request.PairId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePairDrills"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePairDrillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the disaster recovery drills that were performed on the replication pair whose secondary disk resides in a specific region.
//
// @param request - DescribePairDrillsRequest
//
// @return DescribePairDrillsResponse
func (client *Client) DescribePairDrills(request *DescribePairDrillsRequest) (_result *DescribePairDrillsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePairDrillsResponse{}
	_body, _err := client.DescribePairDrillsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions supported by Elastic Block Storage (EBS) features, including asynchronous replication, EBS Lens, and dedicated block storage clusters.
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2021-07-30"),
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
// Queries the regions supported by Elastic Block Storage (EBS) features, including asynchronous replication, EBS Lens, and dedicated block storage clusters.
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
// Queries the disaster recovery drills that were performed on the replication pair-consistent group whose secondary disk resides in a specific region.
//
// @param request - DescribeReplicaGroupDrillsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeReplicaGroupDrillsResponse
func (client *Client) DescribeReplicaGroupDrillsWithOptions(request *DescribeReplicaGroupDrillsRequest, runtime *dara.RuntimeOptions) (_result *DescribeReplicaGroupDrillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrillId) {
		query["DrillId"] = request.DrillId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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
		Action:      dara.String("DescribeReplicaGroupDrills"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeReplicaGroupDrillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the disaster recovery drills that were performed on the replication pair-consistent group whose secondary disk resides in a specific region.
//
// @param request - DescribeReplicaGroupDrillsRequest
//
// @return DescribeReplicaGroupDrillsResponse
func (client *Client) DescribeReplicaGroupDrills(request *DescribeReplicaGroupDrillsRequest) (_result *DescribeReplicaGroupDrillsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeReplicaGroupDrillsResponse{}
	_body, _err := client.DescribeReplicaGroupDrillsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default configurations of a solution.
//
// @param request - DescribeSolutionInstanceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSolutionInstanceConfigurationResponse
func (client *Client) DescribeSolutionInstanceConfigurationWithOptions(request *DescribeSolutionInstanceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DescribeSolutionInstanceConfigurationResponse, _err error) {
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

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SolutionId) {
		query["SolutionId"] = request.SolutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSolutionInstanceConfiguration"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSolutionInstanceConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default configurations of a solution.
//
// @param request - DescribeSolutionInstanceConfigurationRequest
//
// @return DescribeSolutionInstanceConfigurationResponse
func (client *Client) DescribeSolutionInstanceConfiguration(request *DescribeSolutionInstanceConfigurationRequest) (_result *DescribeSolutionInstanceConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSolutionInstanceConfigurationResponse{}
	_body, _err := client.DescribeSolutionInstanceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Centralized Role: Query User Disk Snapshot tagKeys
//
// Description:
//
// ## Interface Description
//
// Query the tag key-value pairs of user\\"s cloud disk and snapshot. The search scope can be narrowed down by using filterTagKey.
//
// @param request - DescribeUserTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserTagKeysResponse
func (client *Client) DescribeUserTagKeysWithOptions(request *DescribeUserTagKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserTagKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TagFilterKey) {
		body["TagFilterKey"] = request.TagFilterKey
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserTagKeys"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Centralized Role: Query User Disk Snapshot tagKeys
//
// Description:
//
// ## Interface Description
//
// Query the tag key-value pairs of user\\"s cloud disk and snapshot. The search scope can be narrowed down by using filterTagKey.
//
// @param request - DescribeUserTagKeysRequest
//
// @return DescribeUserTagKeysResponse
func (client *Client) DescribeUserTagKeys(request *DescribeUserTagKeysRequest) (_result *DescribeUserTagKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserTagKeysResponse{}
	_body, _err := client.DescribeUserTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Centralized Role: Query User Disk and Snapshot tagValues
//
// Description:
//
// ## Interface Description
//
// > The dedicated block storage cluster feature is currently supported in the following regions: South China 2 (Heyuan), Indonesia (Jakarta), and South China 1 (Shenzhen).
//
// - The request parameters act as a filter, with a logical AND relationship. If any parameter is empty, the filter does not take effect.
//
// - For paginated queries, it is recommended to use the MaxResults and NextToken parameters. Usage instructions: When querying the first page, set only MaxResults to limit the number of returned entries. The NextToken in the response will serve as the token for querying subsequent pages. When querying subsequent pages, set the NextToken parameter to the value obtained from the previous response, and set MaxResults to limit the number of returned entries.
//
// @param request - DescribeUserTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserTagValuesResponse
func (client *Client) DescribeUserTagValuesWithOptions(request *DescribeUserTagValuesRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserTagValuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TagFilterValue) {
		body["TagFilterValue"] = request.TagFilterValue
	}

	if !dara.IsNil(request.TagKey) {
		body["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserTagValues"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Centralized Role: Query User Disk and Snapshot tagValues
//
// Description:
//
// ## Interface Description
//
// > The dedicated block storage cluster feature is currently supported in the following regions: South China 2 (Heyuan), Indonesia (Jakarta), and South China 1 (Shenzhen).
//
// - The request parameters act as a filter, with a logical AND relationship. If any parameter is empty, the filter does not take effect.
//
// - For paginated queries, it is recommended to use the MaxResults and NextToken parameters. Usage instructions: When querying the first page, set only MaxResults to limit the number of returned entries. The NextToken in the response will serve as the token for querying subsequent pages. When querying subsequent pages, set the NextToken parameter to the value obtained from the previous response, and set MaxResults to limit the number of returned entries.
//
// @param request - DescribeUserTagValuesRequest
//
// @return DescribeUserTagValuesResponse
func (client *Client) DescribeUserTagValues(request *DescribeUserTagValuesRequest) (_result *DescribeUserTagValuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserTagValuesResponse{}
	_body, _err := client.DescribeUserTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the failover feature for replication pairs in a replication pair-consistent group. When the primary disks of specific replication pairs in a replication pair-consistent group fail, you can call this operation to enable the read and write permissions on the secondary disks.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair-consistent group must be in the **One-time Syncing*	- (`manual_syncing`), **Syncing*	- (`syncing`), **Normal*	- (`normal`), **Stopping*	- (`stopping`), **Stop Failed*	- (`stop_failed`), **Stopped*	- (`stopped`), **In Failover*	- (`failovering`), **Failover Failed*	- (`failover_failed`), or **Failovered*	- (`failovered`) state.
//
//   - After a failover is performed, the replication pair-consistent group enters the **Failovered*	- (`failovered`) state.
//
//   - Before you perform a failover, make sure that the first full data synchronization is completed between the primary site and secondary site.
//
// @param request - FailoverDiskReplicaGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FailoverDiskReplicaGroupResponse
func (client *Client) FailoverDiskReplicaGroupWithOptions(request *FailoverDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *FailoverDiskReplicaGroupResponse, _err error) {
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

	if !dara.IsNil(request.ReplicaGroupId) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FailoverDiskReplicaGroup"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FailoverDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the failover feature for replication pairs in a replication pair-consistent group. When the primary disks of specific replication pairs in a replication pair-consistent group fail, you can call this operation to enable the read and write permissions on the secondary disks.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair-consistent group must be in the **One-time Syncing*	- (`manual_syncing`), **Syncing*	- (`syncing`), **Normal*	- (`normal`), **Stopping*	- (`stopping`), **Stop Failed*	- (`stop_failed`), **Stopped*	- (`stopped`), **In Failover*	- (`failovering`), **Failover Failed*	- (`failover_failed`), or **Failovered*	- (`failovered`) state.
//
//   - After a failover is performed, the replication pair-consistent group enters the **Failovered*	- (`failovered`) state.
//
//   - Before you perform a failover, make sure that the first full data synchronization is completed between the primary site and secondary site.
//
// @param request - FailoverDiskReplicaGroupRequest
//
// @return FailoverDiskReplicaGroupResponse
func (client *Client) FailoverDiskReplicaGroup(request *FailoverDiskReplicaGroupRequest) (_result *FailoverDiskReplicaGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FailoverDiskReplicaGroupResponse{}
	_body, _err := client.FailoverDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the failover feature for replication pairs.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which async replication is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair for which you want to enable failover cannot be in the **Invalid*	- (`invalid`) or **Deleted*	- (`deleted`) state.
//
//   - After a failover is performed, the replication pair enters the **Failovered*	- (`failovered`) state.
//
// @param request - FailoverDiskReplicaPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FailoverDiskReplicaPairResponse
func (client *Client) FailoverDiskReplicaPairWithOptions(request *FailoverDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *FailoverDiskReplicaPairResponse, _err error) {
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

	if !dara.IsNil(request.ReplicaPairId) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FailoverDiskReplicaPair"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FailoverDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the failover feature for replication pairs.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which async replication is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair for which you want to enable failover cannot be in the **Invalid*	- (`invalid`) or **Deleted*	- (`deleted`) state.
//
//   - After a failover is performed, the replication pair enters the **Failovered*	- (`failovered`) state.
//
// @param request - FailoverDiskReplicaPairRequest
//
// @return FailoverDiskReplicaPairResponse
func (client *Client) FailoverDiskReplicaPair(request *FailoverDiskReplicaPairRequest) (_result *FailoverDiskReplicaPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FailoverDiskReplicaPairResponse{}
	_body, _err := client.FailoverDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a user usage report by report ID using the centralized role.
//
// @param request - GetReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReportResponse
func (client *Client) GetReportWithOptions(request *GetReportRequest, runtime *dara.RuntimeOptions) (_result *GetReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ReportType) {
		query["ReportType"] = request.ReportType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReportId) {
		body["ReportId"] = request.ReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetReport"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a user usage report by report ID using the centralized role.
//
// @param request - GetReportRequest
//
// @return GetReportResponse
func (client *Client) GetReport(request *GetReportRequest) (_result *GetReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetReportResponse{}
	_body, _err := client.GetReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the destination regions and zones that support asynchronous replication for a specified zone.
//
// @param request - ListReplicaEdgeSupportedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReplicaEdgeSupportedResponse
func (client *Client) ListReplicaEdgeSupportedWithOptions(request *ListReplicaEdgeSupportedRequest, runtime *dara.RuntimeOptions) (_result *ListReplicaEdgeSupportedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Azone) {
		query["Azone"] = request.Azone
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
		Action:      dara.String("ListReplicaEdgeSupported"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListReplicaEdgeSupportedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the destination regions and zones that support asynchronous replication for a specified zone.
//
// @param request - ListReplicaEdgeSupportedRequest
//
// @return ListReplicaEdgeSupportedResponse
func (client *Client) ListReplicaEdgeSupported(request *ListReplicaEdgeSupportedRequest) (_result *ListReplicaEdgeSupportedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListReplicaEdgeSupportedResponse{}
	_body, _err := client.ListReplicaEdgeSupportedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical report list for a specified application using a centralized role.
//
// @param request - ListReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReportsResponse
func (client *Client) ListReportsWithOptions(request *ListReportsRequest, runtime *dara.RuntimeOptions) (_result *ListReportsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListReports"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical report list for a specified application using a centralized role.
//
// @param request - ListReportsRequest
//
// @return ListReportsResponse
func (client *Client) ListReports(request *ListReportsRequest) (_result *ListReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListReportsResponse{}
	_body, _err := client.ListReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to EBS resources, or queries the IDs and tags of resources in a non-default resource group.
//
// Description:
//
// Specify at least one of the following parameters or parameter pairs in a request to determine a query object:
//
//   - `ResourceId.N`
//
//   - `Tag.N` parameter pair (`Tag.N.Key` and `Tag.N.Value`)
//
// If you set `Tag.N` and `ResourceId.N` at the same time, the EBS resources that match both the parameters are returned.
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
		Version:     dara.String("2021-07-30"),
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
// Queries the tags that are added to EBS resources, or queries the IDs and tags of resources in a non-default resource group.
//
// Description:
//
// Specify at least one of the following parameters or parameter pairs in a request to determine a query object:
//
//   - `ResourceId.N`
//
//   - `Tag.N` parameter pair (`Tag.N.Key` and `Tag.N.Value`)
//
// If you set `Tag.N` and `ResourceId.N` at the same time, the EBS resources that match both the parameters are returned.
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
// Modifies an app from a centralized role.
//
// @param request - ModifyAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppResponse
func (client *Client) ModifyAppWithOptions(request *ModifyAppRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppTags) {
		query["AppTags"] = request.AppTags
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReportSendEnabled) {
		query["ReportSendEnabled"] = request.ReportSendEnabled
	}

	if !dara.IsNil(request.SubscribePeriod) {
		query["SubscribePeriod"] = request.SubscribePeriod
	}

	if !dara.IsNil(request.SubscribeStatus) {
		query["SubscribeStatus"] = request.SubscribeStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApp"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an app from a centralized role.
//
// @param request - ModifyAppRequest
//
// @return ModifyAppResponse
func (client *Client) ModifyApp(request *ModifyAppRequest) (_result *ModifyAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppResponse{}
	_body, _err := client.ModifyAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information of a dedicated elastic block storage cluster.
//
// Description:
//
// You can call this operation to modify the information of a dedicated block storage cluster. The information includes the name and description of the cluster.
//
// @param request - ModifyDedicatedBlockStorageClusterAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDedicatedBlockStorageClusterAttributeResponse
func (client *Client) ModifyDedicatedBlockStorageClusterAttributeWithOptions(request *ModifyDedicatedBlockStorageClusterAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDedicatedBlockStorageClusterAttributeResponse, _err error) {
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

	if !dara.IsNil(request.DbscId) {
		query["DbscId"] = request.DbscId
	}

	if !dara.IsNil(request.DbscName) {
		query["DbscName"] = request.DbscName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDedicatedBlockStorageClusterAttribute"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDedicatedBlockStorageClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of a dedicated elastic block storage cluster.
//
// Description:
//
// You can call this operation to modify the information of a dedicated block storage cluster. The information includes the name and description of the cluster.
//
// @param request - ModifyDedicatedBlockStorageClusterAttributeRequest
//
// @return ModifyDedicatedBlockStorageClusterAttributeResponse
func (client *Client) ModifyDedicatedBlockStorageClusterAttribute(request *ModifyDedicatedBlockStorageClusterAttributeRequest) (_result *ModifyDedicatedBlockStorageClusterAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDedicatedBlockStorageClusterAttributeResponse{}
	_body, _err := client.ModifyDedicatedBlockStorageClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name, description, or Recovery Point Objective (RPO) of a replication pair-consistent group.
//
// Description:
//
// ## Description
//
// - For information about the regions that support the replication pair-consistent group feature, see [Overview of asynchronous replication](https://help.aliyun.com/document_detail/314563.html).
//
// - The replication pair-consistent group must be in the **Created*	- (`created`) or **Stopped*	- (`stopped`) state.
//
// @param request - ModifyDiskReplicaGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskReplicaGroupResponse
func (client *Client) ModifyDiskReplicaGroupWithOptions(request *ModifyDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDiskReplicaGroupResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableRtc) {
		query["EnableRtc"] = request.EnableRtc
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.RPO) {
		query["RPO"] = request.RPO
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplicaGroupId) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDiskReplicaGroup"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, or Recovery Point Objective (RPO) of a replication pair-consistent group.
//
// Description:
//
// ## Description
//
// - For information about the regions that support the replication pair-consistent group feature, see [Overview of asynchronous replication](https://help.aliyun.com/document_detail/314563.html).
//
// - The replication pair-consistent group must be in the **Created*	- (`created`) or **Stopped*	- (`stopped`) state.
//
// @param request - ModifyDiskReplicaGroupRequest
//
// @return ModifyDiskReplicaGroupResponse
func (client *Client) ModifyDiskReplicaGroup(request *ModifyDiskReplicaGroupRequest) (_result *ModifyDiskReplicaGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDiskReplicaGroupResponse{}
	_body, _err := client.ModifyDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an async replication pair.
//
// Description:
//
// ## Description
//
// - For information about the regions that support the asynchronous replication feature, see [Overview of asynchronous replication](https://help.aliyun.com/document_detail/314563.html).
//
// - You can modify the name or description of a replication pair only when the pair is in the **Created*	- (`created`) or **Stopped*	- (`stopped`) state.
//
// @param request - ModifyDiskReplicaPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskReplicaPairResponse
func (client *Client) ModifyDiskReplicaPairWithOptions(request *ModifyDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *ModifyDiskReplicaPairResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableRtc) {
		query["EnableRtc"] = request.EnableRtc
	}

	if !dara.IsNil(request.PairName) {
		query["PairName"] = request.PairName
	}

	if !dara.IsNil(request.RPO) {
		query["RPO"] = request.RPO
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplicaPairId) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDiskReplicaPair"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an async replication pair.
//
// Description:
//
// ## Description
//
// - For information about the regions that support the asynchronous replication feature, see [Overview of asynchronous replication](https://help.aliyun.com/document_detail/314563.html).
//
// - You can modify the name or description of a replication pair only when the pair is in the **Created*	- (`created`) or **Stopped*	- (`stopped`) state.
//
// @param request - ModifyDiskReplicaPairRequest
//
// @return ModifyDiskReplicaPairResponse
func (client *Client) ModifyDiskReplicaPair(request *ModifyDiskReplicaPairRequest) (_result *ModifyDiskReplicaPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDiskReplicaPairResponse{}
	_body, _err := client.ModifyDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the throughput status of a dedicated block storage cluster disk which has been set through the SetDedicatedBlockStorageClusterDiskThroughput API.
//
// @param request - QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse
func (client *Client) QueryDedicatedBlockStorageClusterDiskThroughputStatusWithOptions(request *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.QosRequestId) {
		body["QosRequestId"] = request.QosRequestId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDedicatedBlockStorageClusterDiskThroughputStatus"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the throughput status of a dedicated block storage cluster disk which has been set through the SetDedicatedBlockStorageClusterDiskThroughput API.
//
// @param request - QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest
//
// @return QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse
func (client *Client) QueryDedicatedBlockStorageClusterDiskThroughputStatus(request *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) (_result *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse{}
	_body, _err := client.QueryDedicatedBlockStorageClusterDiskThroughputStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the capacity trend data of a dedicated Elastic Block Storage (EBS) cluster within a specific time period, including the available capacity and total capacity.
//
// Description:
//
// Period is the time interval between data retrieval points. When set to 60 (minute interval), a maximum of 1440 data points can be returned; when set to 3600 (hour interval), a maximum of 744 data points can be returned; and when set to 86400 (day interval), a maximum of 366 data points can be returned.
//
// @param request - QueryDedicatedBlockStorageClusterInventoryDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDedicatedBlockStorageClusterInventoryDataResponse
func (client *Client) QueryDedicatedBlockStorageClusterInventoryDataWithOptions(request *QueryDedicatedBlockStorageClusterInventoryDataRequest, runtime *dara.RuntimeOptions) (_result *QueryDedicatedBlockStorageClusterInventoryDataResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.DbscId) {
		body["DbscId"] = request.DbscId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDedicatedBlockStorageClusterInventoryData"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDedicatedBlockStorageClusterInventoryDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the capacity trend data of a dedicated Elastic Block Storage (EBS) cluster within a specific time period, including the available capacity and total capacity.
//
// Description:
//
// Period is the time interval between data retrieval points. When set to 60 (minute interval), a maximum of 1440 data points can be returned; when set to 3600 (hour interval), a maximum of 744 data points can be returned; and when set to 86400 (day interval), a maximum of 366 data points can be returned.
//
// @param request - QueryDedicatedBlockStorageClusterInventoryDataRequest
//
// @return QueryDedicatedBlockStorageClusterInventoryDataResponse
func (client *Client) QueryDedicatedBlockStorageClusterInventoryData(request *QueryDedicatedBlockStorageClusterInventoryDataRequest) (_result *QueryDedicatedBlockStorageClusterInventoryDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDedicatedBlockStorageClusterInventoryDataResponse{}
	_body, _err := client.QueryDedicatedBlockStorageClusterInventoryDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a replication pair from a replication pair-consistent group. After a replication pair is removed from a replication pair-consistent group, the pair is disassociated from the group but is not deleted.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair-consistent group from which you want to remove a replication pair must be in the **Created*	- (`created`), **Stopped*	- (`stopped`), or **Invalid*	- (`invalid`) state.
//
// @param request - RemoveDiskReplicaPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDiskReplicaPairResponse
func (client *Client) RemoveDiskReplicaPairWithOptions(request *RemoveDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *RemoveDiskReplicaPairResponse, _err error) {
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

	if !dara.IsNil(request.ReplicaGroupId) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	if !dara.IsNil(request.ReplicaPairId) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDiskReplicaPair"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a replication pair from a replication pair-consistent group. After a replication pair is removed from a replication pair-consistent group, the pair is disassociated from the group but is not deleted.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair-consistent group from which you want to remove a replication pair must be in the **Created*	- (`created`), **Stopped*	- (`stopped`), or **Invalid*	- (`invalid`) state.
//
// @param request - RemoveDiskReplicaPairRequest
//
// @return RemoveDiskReplicaPairResponse
func (client *Client) RemoveDiskReplicaPair(request *RemoveDiskReplicaPairRequest) (_result *RemoveDiskReplicaPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveDiskReplicaPairResponse{}
	_body, _err := client.RemoveDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the reverse replication feature for replication pairs that belong to a replication pair-consistent group. After reverse replication is enabled, data stored on the original secondary disks is replicated to the original primary disks. When a reverse replication is being performed, the primary and secondary sites of the replication pair-consistent group remain unchanged, but data is replicated from the secondary site to the primary site.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair-consistent group for which you want to enable reverse replication must be in the **Failovered*	- (`failovered`) state. You can call the `FailoverDiskReplicaPair` operation to enable failover.
//
//   - Before a reverse replication is performed, the primary disks must be detached from its associated Elastic Compute Service (ECS) instance and must be in the Unattached state. You can call the [DetachDisk](https://help.aliyun.com/document_detail/25516.html) operation to detach the disks.
//
//   - After you enable reverse replication, you must call the `StartDiskReplicaPair` operation again to enable the async replication feature before data can be replicated from the original secondary disks to the original primary disks.
//
//   - You can set the ReverseReplicate parameter to false to cancel the **Failovered*	- (`failovered`) state and restore the original replication direction.
//
// @param request - ReprotectDiskReplicaGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReprotectDiskReplicaGroupResponse
func (client *Client) ReprotectDiskReplicaGroupWithOptions(request *ReprotectDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *ReprotectDiskReplicaGroupResponse, _err error) {
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

	if !dara.IsNil(request.ReplicaGroupId) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	if !dara.IsNil(request.ReverseReplicate) {
		query["ReverseReplicate"] = request.ReverseReplicate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReprotectDiskReplicaGroup"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReprotectDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the reverse replication feature for replication pairs that belong to a replication pair-consistent group. After reverse replication is enabled, data stored on the original secondary disks is replicated to the original primary disks. When a reverse replication is being performed, the primary and secondary sites of the replication pair-consistent group remain unchanged, but data is replicated from the secondary site to the primary site.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair-consistent group for which you want to enable reverse replication must be in the **Failovered*	- (`failovered`) state. You can call the `FailoverDiskReplicaPair` operation to enable failover.
//
//   - Before a reverse replication is performed, the primary disks must be detached from its associated Elastic Compute Service (ECS) instance and must be in the Unattached state. You can call the [DetachDisk](https://help.aliyun.com/document_detail/25516.html) operation to detach the disks.
//
//   - After you enable reverse replication, you must call the `StartDiskReplicaPair` operation again to enable the async replication feature before data can be replicated from the original secondary disks to the original primary disks.
//
//   - You can set the ReverseReplicate parameter to false to cancel the **Failovered*	- (`failovered`) state and restore the original replication direction.
//
// @param request - ReprotectDiskReplicaGroupRequest
//
// @return ReprotectDiskReplicaGroupResponse
func (client *Client) ReprotectDiskReplicaGroup(request *ReprotectDiskReplicaGroupRequest) (_result *ReprotectDiskReplicaGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReprotectDiskReplicaGroupResponse{}
	_body, _err := client.ReprotectDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the reverse replication feature for a replication pair.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which async replication is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair for which you want to enable reverse replication must be in the **Failovered*	- (`failovered`) state. You can call the [FailoverDiskReplicaPair](https://help.aliyun.com/document_detail/354358.html) operation to enable failover.
//
//   - The primary disk must be detached from its associated Elastic Compute Service (ECS) instance and is in the Unattached state. You can call the [DetachDisk](https://help.aliyun.com/document_detail/25516.html) operation to detach the disk.
//
//   - After you enable reverse replication, you must call the [StartDiskReplicaPair](https://help.aliyun.com/document_detail/354205.html) operation again to activate the replication pair before data can be replicated from the original secondary disk to the original primary disk.
//
//   - You can set the ReverseReplicate parameter to false to cancel the **Failovered*	- (`failovered`) state and restore the original replication direction.
//
// @param request - ReprotectDiskReplicaPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReprotectDiskReplicaPairResponse
func (client *Client) ReprotectDiskReplicaPairWithOptions(request *ReprotectDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *ReprotectDiskReplicaPairResponse, _err error) {
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

	if !dara.IsNil(request.ReplicaPairId) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	if !dara.IsNil(request.ReverseReplicate) {
		query["ReverseReplicate"] = request.ReverseReplicate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReprotectDiskReplicaPair"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReprotectDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the reverse replication feature for a replication pair.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which async replication is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair for which you want to enable reverse replication must be in the **Failovered*	- (`failovered`) state. You can call the [FailoverDiskReplicaPair](https://help.aliyun.com/document_detail/354358.html) operation to enable failover.
//
//   - The primary disk must be detached from its associated Elastic Compute Service (ECS) instance and is in the Unattached state. You can call the [DetachDisk](https://help.aliyun.com/document_detail/25516.html) operation to detach the disk.
//
//   - After you enable reverse replication, you must call the [StartDiskReplicaPair](https://help.aliyun.com/document_detail/354205.html) operation again to activate the replication pair before data can be replicated from the original secondary disk to the original primary disk.
//
//   - You can set the ReverseReplicate parameter to false to cancel the **Failovered*	- (`failovered`) state and restore the original replication direction.
//
// @param request - ReprotectDiskReplicaPairRequest
//
// @return ReprotectDiskReplicaPairResponse
func (client *Client) ReprotectDiskReplicaPair(request *ReprotectDiskReplicaPairRequest) (_result *ReprotectDiskReplicaPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReprotectDiskReplicaPairResponse{}
	_body, _err := client.ReprotectDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// In the elastic type dedicated block storage cluster, you can easily achieve the specified throughput (Bps) for the target disk. You only need to set the cloud disk ID and the target throughput, simplifying the process of configuring.
//
// @param request - SetDedicatedBlockStorageClusterDiskThroughputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDedicatedBlockStorageClusterDiskThroughputResponse
func (client *Client) SetDedicatedBlockStorageClusterDiskThroughputWithOptions(request *SetDedicatedBlockStorageClusterDiskThroughputRequest, runtime *dara.RuntimeOptions) (_result *SetDedicatedBlockStorageClusterDiskThroughputResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Bps) {
		body["Bps"] = request.Bps
	}

	if !dara.IsNil(request.DiskId) {
		body["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDedicatedBlockStorageClusterDiskThroughput"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDedicatedBlockStorageClusterDiskThroughputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// In the elastic type dedicated block storage cluster, you can easily achieve the specified throughput (Bps) for the target disk. You only need to set the cloud disk ID and the target throughput, simplifying the process of configuring.
//
// @param request - SetDedicatedBlockStorageClusterDiskThroughputRequest
//
// @return SetDedicatedBlockStorageClusterDiskThroughputResponse
func (client *Client) SetDedicatedBlockStorageClusterDiskThroughput(request *SetDedicatedBlockStorageClusterDiskThroughputRequest) (_result *SetDedicatedBlockStorageClusterDiskThroughputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDedicatedBlockStorageClusterDiskThroughputResponse{}
	_body, _err := client.SetDedicatedBlockStorageClusterDiskThroughputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the async replication feature for replication pairs that belong to a replication pair-consistent group. The first time the async replication feature is enabled for the replication pairs, the system performs a full synchronization to synchronize all data from disks at the primary site (primary disks) to disks at the secondary site (secondary disks). Then, the system periodically synchronizes incremental data based on the recovery point objective (RPO) of the replication pair-consistent group.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - If you set the `OneShot` to `false`, the replication pair-consistent group must be in the **Created*	- (`created` ), **Synchronizing*	- (`syncing` ), **Normal*	- (`normal` ), or **Stopped*	- (`stopped`) state.
//
//   - If you set `OneShot` to `true`, the replication pair-consistent group must be in the **Created*	- (`created` ), **One-time Syncing*	- (`manual_syncing` ), or **Stopped*	- (`stopped`) state. The time interval between two consecutive one-time synchronizations must be longer than one half of the recovery point objective (RPO).
//
//   - After a replication pair-consistent group is activated, the group enters the **Initial Syncing*	- (`initial_syncing`) state and the system performs the first async replication to replicate all data from the primary disks to secondary disks.
//
// @param request - StartDiskReplicaGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDiskReplicaGroupResponse
func (client *Client) StartDiskReplicaGroupWithOptions(request *StartDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *StartDiskReplicaGroupResponse, _err error) {
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

	if !dara.IsNil(request.OneShot) {
		query["OneShot"] = request.OneShot
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplicaGroupId) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDiskReplicaGroup"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the async replication feature for replication pairs that belong to a replication pair-consistent group. The first time the async replication feature is enabled for the replication pairs, the system performs a full synchronization to synchronize all data from disks at the primary site (primary disks) to disks at the secondary site (secondary disks). Then, the system periodically synchronizes incremental data based on the recovery point objective (RPO) of the replication pair-consistent group.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - If you set the `OneShot` to `false`, the replication pair-consistent group must be in the **Created*	- (`created` ), **Synchronizing*	- (`syncing` ), **Normal*	- (`normal` ), or **Stopped*	- (`stopped`) state.
//
//   - If you set `OneShot` to `true`, the replication pair-consistent group must be in the **Created*	- (`created` ), **One-time Syncing*	- (`manual_syncing` ), or **Stopped*	- (`stopped`) state. The time interval between two consecutive one-time synchronizations must be longer than one half of the recovery point objective (RPO).
//
//   - After a replication pair-consistent group is activated, the group enters the **Initial Syncing*	- (`initial_syncing`) state and the system performs the first async replication to replicate all data from the primary disks to secondary disks.
//
// @param request - StartDiskReplicaGroupRequest
//
// @return StartDiskReplicaGroupResponse
func (client *Client) StartDiskReplicaGroup(request *StartDiskReplicaGroupRequest) (_result *StartDiskReplicaGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDiskReplicaGroupResponse{}
	_body, _err := client.StartDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates the async replication feature for a specific replication pair.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which async replication is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - Only replication pairs that are in the **Created*	- (`created`) or **Stopped*	- (`stopped`) state can be activated.
//
//   - After a replication pair is activated, it enters the **Initial Syncing*	- (`initial_syncing`) state and the system performs the first asynchronous replication to replicate all data from the primary disk to the secondary disk.
//
// @param request - StartDiskReplicaPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDiskReplicaPairResponse
func (client *Client) StartDiskReplicaPairWithOptions(request *StartDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *StartDiskReplicaPairResponse, _err error) {
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

	if !dara.IsNil(request.OneShot) {
		query["OneShot"] = request.OneShot
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplicaPairId) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDiskReplicaPair"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates the async replication feature for a specific replication pair.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which async replication is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - Only replication pairs that are in the **Created*	- (`created`) or **Stopped*	- (`stopped`) state can be activated.
//
//   - After a replication pair is activated, it enters the **Initial Syncing*	- (`initial_syncing`) state and the system performs the first asynchronous replication to replicate all data from the primary disk to the secondary disk.
//
// @param request - StartDiskReplicaPairRequest
//
// @return StartDiskReplicaPairResponse
func (client *Client) StartDiskReplicaPair(request *StartDiskReplicaPairRequest) (_result *StartDiskReplicaPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDiskReplicaPairResponse{}
	_body, _err := client.StartDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a disaster recovery drill to ensure the continued replication and clone the data from the last recovery point of the secondary disk to a new disk. This helps you test the completeness and correctness of applications that are deployed on the disaster recovery site on a regular basis.
//
// Description:
//
// After the disaster recovery drill is complete on the secondary disk, a pay-as-you-go drill disk that has the same capacity and category as the secondary disk is created in the zone where the secondary disk resides. The drill disk contains last-recovery-point data that can be used to test the completeness and correctness of applications.
//
// @param request - StartPairDrillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartPairDrillResponse
func (client *Client) StartPairDrillWithOptions(request *StartPairDrillRequest, runtime *dara.RuntimeOptions) (_result *StartPairDrillResponse, _err error) {
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

	if !dara.IsNil(request.PairId) {
		query["PairId"] = request.PairId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartPairDrill"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartPairDrillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a disaster recovery drill to ensure the continued replication and clone the data from the last recovery point of the secondary disk to a new disk. This helps you test the completeness and correctness of applications that are deployed on the disaster recovery site on a regular basis.
//
// Description:
//
// After the disaster recovery drill is complete on the secondary disk, a pay-as-you-go drill disk that has the same capacity and category as the secondary disk is created in the zone where the secondary disk resides. The drill disk contains last-recovery-point data that can be used to test the completeness and correctness of applications.
//
// @param request - StartPairDrillRequest
//
// @return StartPairDrillResponse
func (client *Client) StartPairDrill(request *StartPairDrillRequest) (_result *StartPairDrillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartPairDrillResponse{}
	_body, _err := client.StartPairDrillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a disaster recovery drill in a replication pair-consistent group to ensure the continued replication and restores data from the latest recovery point of secondary disks to new disks. This helps test the completeness and correctness of applications that are deployed on the disaster recovery site on a regular basis.
//
// Description:
//
// After the disaster recovery drill is complete on secondary disks, a pay-as-you-go drill disk is created in the zone where the secondary disk of each replication pair resides. The latest-recovery-point data is restored to the drill disks to test the completeness and correctness of applications.
//
// @param request - StartReplicaGroupDrillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartReplicaGroupDrillResponse
func (client *Client) StartReplicaGroupDrillWithOptions(request *StartReplicaGroupDrillRequest, runtime *dara.RuntimeOptions) (_result *StartReplicaGroupDrillResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartReplicaGroupDrill"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartReplicaGroupDrillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a disaster recovery drill in a replication pair-consistent group to ensure the continued replication and restores data from the latest recovery point of secondary disks to new disks. This helps test the completeness and correctness of applications that are deployed on the disaster recovery site on a regular basis.
//
// Description:
//
// After the disaster recovery drill is complete on secondary disks, a pay-as-you-go drill disk is created in the zone where the secondary disk of each replication pair resides. The latest-recovery-point data is restored to the drill disks to test the completeness and correctness of applications.
//
// @param request - StartReplicaGroupDrillRequest
//
// @return StartReplicaGroupDrillResponse
func (client *Client) StartReplicaGroupDrill(request *StartReplicaGroupDrillRequest) (_result *StartReplicaGroupDrillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartReplicaGroupDrillResponse{}
	_body, _err := client.StartReplicaGroupDrillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a replication pair-consistent group. This operation stops all replication pairs in the replication pair-consistent group.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair-consistent group that you want to stop must be in the **One-time Syncing*	- (`manual_syncing`), **Syncing*	- (`syncing`), **Normal*	- (`normal`), **Stopping*	- (`stopping`), **Stop Failed*	- (`stop_failed`), or **Stopped*	- (`stopped`) state.
//
//   - When a replication pair-consistent group is stopped, it enters the **Stopped*	- (`stopped`) state. If a replication pair-consistent group cannot be stopped, the state of the group remains unchanged or changes to **Stop Failed*	- (`stop_failed`). In this case, try again later.
//
// @param request - StopDiskReplicaGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDiskReplicaGroupResponse
func (client *Client) StopDiskReplicaGroupWithOptions(request *StopDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *StopDiskReplicaGroupResponse, _err error) {
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

	if !dara.IsNil(request.ReplicaGroupId) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDiskReplicaGroup"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a replication pair-consistent group. This operation stops all replication pairs in the replication pair-consistent group.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair-consistent group that you want to stop must be in the **One-time Syncing*	- (`manual_syncing`), **Syncing*	- (`syncing`), **Normal*	- (`normal`), **Stopping*	- (`stopping`), **Stop Failed*	- (`stop_failed`), or **Stopped*	- (`stopped`) state.
//
//   - When a replication pair-consistent group is stopped, it enters the **Stopped*	- (`stopped`) state. If a replication pair-consistent group cannot be stopped, the state of the group remains unchanged or changes to **Stop Failed*	- (`stop_failed`). In this case, try again later.
//
// @param request - StopDiskReplicaGroupRequest
//
// @return StopDiskReplicaGroupResponse
func (client *Client) StopDiskReplicaGroup(request *StopDiskReplicaGroupRequest) (_result *StopDiskReplicaGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDiskReplicaGroupResponse{}
	_body, _err := client.StopDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a replication pair.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which async replication is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - Only replication pairs that are in the **Initial Syncing*	- (`initial_syncing`), **Syncing*	- (`syncing`), **One-time Syncing*	- (`manual_syncing`), or **Normal*	- (`normal`) state can be stopped. When a replication pair is stopped, it enters the Stopped (`stopped`) state. The secondary disk rolls back to the point in time when the last async replication was complete and drops all the data that is being replicated from the primary disk.
//
// @param request - StopDiskReplicaPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDiskReplicaPairResponse
func (client *Client) StopDiskReplicaPairWithOptions(request *StopDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *StopDiskReplicaPairResponse, _err error) {
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

	if !dara.IsNil(request.ReplicaPairId) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDiskReplicaPair"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a replication pair.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which async replication is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - Only replication pairs that are in the **Initial Syncing*	- (`initial_syncing`), **Syncing*	- (`syncing`), **One-time Syncing*	- (`manual_syncing`), or **Normal*	- (`normal`) state can be stopped. When a replication pair is stopped, it enters the Stopped (`stopped`) state. The secondary disk rolls back to the point in time when the last async replication was complete and drops all the data that is being replicated from the primary disk.
//
// @param request - StopDiskReplicaPairRequest
//
// @return StopDiskReplicaPairResponse
func (client *Client) StopDiskReplicaPair(request *StopDiskReplicaPairRequest) (_result *StopDiskReplicaPairResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDiskReplicaPairResponse{}
	_body, _err := client.StopDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates tags and adds the tags to Elastic Block Storage (EBS) resources.
//
// Description:
//
// Before you add tags to a resource, Alibaba Cloud checks the number of existing tags of the resource. If the maximum number of tags is reached, an error message is returned. For more information, see the "Tag limits" section in [Limits](https://help.aliyun.com/document_detail/25412.html).
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
		Version:     dara.String("2021-07-30"),
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
// Creates tags and adds the tags to Elastic Block Storage (EBS) resources.
//
// Description:
//
// Before you add tags to a resource, Alibaba Cloud checks the number of existing tags of the resource. If the maximum number of tags is reached, an error message is returned. For more information, see the "Tag limits" section in [Limits](https://help.aliyun.com/document_detail/25412.html).
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
// Unbind disks from a enterprise-level snapshot policy.
//
// @param request - UnbindEnterpriseSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindEnterpriseSnapshotPolicyResponse
func (client *Client) UnbindEnterpriseSnapshotPolicyWithOptions(request *UnbindEnterpriseSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *UnbindEnterpriseSnapshotPolicyResponse, _err error) {
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

	if !dara.IsNil(request.DiskTargets) {
		query["DiskTargets"] = request.DiskTargets
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
		Action:      dara.String("UnbindEnterpriseSnapshotPolicy"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbind disks from a enterprise-level snapshot policy.
//
// @param request - UnbindEnterpriseSnapshotPolicyRequest
//
// @return UnbindEnterpriseSnapshotPolicyResponse
func (client *Client) UnbindEnterpriseSnapshotPolicy(request *UnbindEnterpriseSnapshotPolicyRequest) (_result *UnbindEnterpriseSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.UnbindEnterpriseSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from specified Elastic Block Storage (EBS) resources.
//
// Description:
//
//	  You can remove up to 20 tags at a time.
//
//		- After a tag is removed from an EBS resource, the tag is automatically deleted if the tag is not added to any instance.
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
		Version:     dara.String("2021-07-30"),
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
// Removes tags from specified Elastic Block Storage (EBS) resources.
//
// Description:
//
//	  You can remove up to 20 tags at a time.
//
//		- After a tag is removed from an EBS resource, the tag is automatically deleted if the tag is not added to any instance.
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
// Modifies the configuration of an enterprise-level snapshot policy.
//
// @param tmpReq - UpdateEnterpriseSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnterpriseSnapshotPolicyResponse
func (client *Client) UpdateEnterpriseSnapshotPolicyWithOptions(tmpReq *UpdateEnterpriseSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateEnterpriseSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateEnterpriseSnapshotPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CrossRegionCopyInfo) {
		request.CrossRegionCopyInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CrossRegionCopyInfo, dara.String("CrossRegionCopyInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RetainRule) {
		request.RetainRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RetainRule, dara.String("RetainRule"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Schedule) {
		request.ScheduleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Schedule, dara.String("Schedule"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SpecialRetainRules) {
		request.SpecialRetainRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SpecialRetainRules, dara.String("SpecialRetainRules"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StorageRule) {
		request.StorageRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StorageRule, dara.String("StorageRule"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossRegionCopyInfoShrink) {
		query["CrossRegionCopyInfo"] = request.CrossRegionCopyInfoShrink
	}

	if !dara.IsNil(request.Desc) {
		query["Desc"] = request.Desc
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RetainRuleShrink) {
		query["RetainRule"] = request.RetainRuleShrink
	}

	if !dara.IsNil(request.ScheduleShrink) {
		query["Schedule"] = request.ScheduleShrink
	}

	if !dara.IsNil(request.SpecialRetainRulesShrink) {
		query["SpecialRetainRules"] = request.SpecialRetainRulesShrink
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.StorageRuleShrink) {
		query["StorageRule"] = request.StorageRuleShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEnterpriseSnapshotPolicy"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of an enterprise-level snapshot policy.
//
// @param request - UpdateEnterpriseSnapshotPolicyRequest
//
// @return UpdateEnterpriseSnapshotPolicyResponse
func (client *Client) UpdateEnterpriseSnapshotPolicy(request *UpdateEnterpriseSnapshotPolicyRequest) (_result *UpdateEnterpriseSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.UpdateEnterpriseSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the attributes of a solution instance.
//
// @param request - UpdateSolutionInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSolutionInstanceAttributeResponse
func (client *Client) UpdateSolutionInstanceAttributeWithOptions(request *UpdateSolutionInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateSolutionInstanceAttributeResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SolutionInstanceId) {
		query["SolutionInstanceId"] = request.SolutionInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSolutionInstanceAttribute"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSolutionInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attributes of a solution instance.
//
// @param request - UpdateSolutionInstanceAttributeRequest
//
// @return UpdateSolutionInstanceAttributeResponse
func (client *Client) UpdateSolutionInstanceAttribute(request *UpdateSolutionInstanceAttributeRequest) (_result *UpdateSolutionInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSolutionInstanceAttributeResponse{}
	_body, _err := client.UpdateSolutionInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
