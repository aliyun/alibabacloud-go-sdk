// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) AddDiskReplicaPairWithContext(ctx context.Context, request *AddDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *AddDiskReplicaPairResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindEnterpriseSnapshotPolicyResponse
func (client *Client) BindEnterpriseSnapshotPolicyWithContext(ctx context.Context, request *BindEnterpriseSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *BindEnterpriseSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearPairDrillResponse
func (client *Client) ClearPairDrillWithContext(ctx context.Context, request *ClearPairDrillRequest, runtime *dara.RuntimeOptions) (_result *ClearPairDrillResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearReplicaGroupDrillResponse
func (client *Client) ClearReplicaGroupDrillWithContext(ctx context.Context, request *ClearReplicaGroupDrillRequest, runtime *dara.RuntimeOptions) (_result *ClearReplicaGroupDrillResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 中心化角色：创建App
//
// @param request - CreateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithContext(ctx context.Context, request *CreateAppRequest, runtime *dara.RuntimeOptions) (_result *CreateAppResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dedicated block storage cluster. When you call this operation, you can specify parameters, such as Azone, Capacity, Type, and PeriodUnit, in the request.
//
// Description:
//
// ## [](#)Usage notes
//
//   - Dedicated block storage clusters are physically isolated from public block storage clusters. The owner of each dedicated block storage cluster has exclusive access to all resources in the cluster.
//
//   - Disks created in a dedicated block storage cluster can be attached only to Elastic Compute Service (ECS) instances that reside in the same zone as the cluster. Before you create a dedicated block storage cluster, decide the regions and zones in which to deploy your cloud resources.
//
//   - Dedicated block storage clusters are classified into basic and performance types. When you create a dedicated block storage cluster, select a cluster type based on your business requirements.
//
//   - You are charged for creating dedicated block storage clusters.
//
// @param request - CreateDedicatedBlockStorageClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDedicatedBlockStorageClusterResponse
func (client *Client) CreateDedicatedBlockStorageClusterWithContext(ctx context.Context, request *CreateDedicatedBlockStorageClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateDedicatedBlockStorageClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// ## [](#)Usage notes
//
// The replication pair-consistent group feature allows you to batch manage multiple disks in disaster recovery scenarios. You can restore the data of all disks in the same replication pair-consistent group to the same point in time to allow for disaster recovery of instances.
//
// Take note of the following items:
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - Replication pair-consistent groups can be used to implement disaster recovery across zones within the same region and disaster recovery across regions.
//
//   - A replication pair and a replication pair-consistent group can replicate in the same direction if they have the same primary region (production region), primary zone (production zone), secondary region (disaster recovery region), and secondary zone (disaster recovery zone). A replication pair can be added to only a replication pair-consistent group that replicates in the same direction as the replication pair.
//
//   - After replication pairs are added to a replication pair-consistent group, the recovery point objective (RPO) of the group takes effect on the pairs instead of their original RPOs.
//
// @param request - CreateDiskReplicaGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiskReplicaGroupResponse
func (client *Client) CreateDiskReplicaGroupWithContext(ctx context.Context, request *CreateDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDiskReplicaGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a replication pair to asynchronously replicate data between disks.
//
// Description:
//
// ## [](#)Usage notes
//
// Async replication is a feature that protects data across regions by using the data replication capability of Elastic Block Storage (EBS). This feature can be used to asynchronously replicate data from a disk in one region to a disk in another region for disaster recovery purposes. You can use this feature to implement disaster recovery for critical business to protect data in your databases and improve business continuity. You are charged on a subscription basis for the bandwidth that is used by the async replication feature.
//
// Currently, the async replication feature can asynchronously replicate data only between enhanced SSDs (ESSDs). The functionality of disks in replication pairs is limited.
//
// Take note of the following items:
//
//   - Make sure that the source disk (primary disk) from which to replicate data and the destination disk (secondary disk) to which to replicate data are created. You can call the [CreateDisk](https://help.aliyun.com/document_detail/25513.html) operation to create disks.
//
//   - The secondary disk cannot reside in the same region as the primary disk. For information about the regions that support async replication, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - After you call this operation to create a replication pair for the primary disk and the secondary disk, you must call the [StartDiskReplicaPair](https://help.aliyun.com/document_detail/354205.html) operation to enable async replication to replicate data from the primary disk to the secondary disk cross regions on a periodic basis.
//
// @param request - CreateDiskReplicaPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiskReplicaPairResponse
func (client *Client) CreateDiskReplicaPairWithContext(ctx context.Context, request *CreateDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *CreateDiskReplicaPairResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create an enterprise-level snapshot policy
//
// @param tmpReq - CreateEnterpriseSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnterpriseSnapshotPolicyResponse
func (client *Client) CreateEnterpriseSnapshotPolicyWithContext(ctx context.Context, tmpReq *CreateEnterpriseSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateEnterpriseSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 中心化角色：删除App
//
// @param request - DeleteAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
func (client *Client) DeleteAppWithContext(ctx context.Context, request *DeleteAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDiskReplicaGroupResponse
func (client *Client) DeleteDiskReplicaGroupWithContext(ctx context.Context, request *DeleteDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDiskReplicaGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDiskReplicaPairResponse
func (client *Client) DeleteDiskReplicaPairWithContext(ctx context.Context, request *DeleteDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *DeleteDiskReplicaPairResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnterpriseSnapshotPolicyResponse
func (client *Client) DeleteEnterpriseSnapshotPolicyWithContext(ctx context.Context, request *DeleteEnterpriseSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnterpriseSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of one or more disks in a dedicated block storage cluster.
//
// Description:
//
//	  You can use one of the following methods to check the responses:
//
//	    	- Method 1: Use `NextToken` to configure the query token. Set the value to the `NextToken` value that is returned in the last call to the DescribeDisks operation. Then, use `MaxResults` to specify the maximum number of entries to return on each page.
//
//	    	- Method 2: Use `PageSize` to specify the number of entries to return on each page and then use `PageNumber` to specify the number of the page to return.
//
//	        You can use only one of the preceding methods. If a large number of entries are to be returned, we recommend that you use method 1. When `NextToken` is specified, `PageSize` and `PageNumber` do not take effect and `TotalCount` in the response is invalid.
//
//		- A disk that has the multi-attach feature enabled can be attached to multiple instances. You can query the attachment information of the disk based on the `Attachment` values in the response.
//
// When you call an API operation by using Alibaba Cloud CLI, you must specify request parameter values of different data types in the required formats. For more information, see [Parameter format overview](https://help.aliyun.com/document_detail/110340.html).
//
// @param request - DescribeDedicatedBlockStorageClusterDisksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDedicatedBlockStorageClusterDisksResponse
func (client *Client) DescribeDedicatedBlockStorageClusterDisksWithContext(ctx context.Context, request *DescribeDedicatedBlockStorageClusterDisksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDedicatedBlockStorageClusterDisksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dedicated block storage clusters that are created.
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
func (client *Client) DescribeDedicatedBlockStorageClustersWithContext(ctx context.Context, request *DescribeDedicatedBlockStorageClustersRequest, runtime *dara.RuntimeOptions) (_result *DescribeDedicatedBlockStorageClustersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the risk events of a disk.
//
// Description:
//
// ## Usage notes
//
// CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
//
// @param request - DescribeDiskEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiskEventsResponse
func (client *Client) DescribeDiskEventsWithContext(ctx context.Context, request *DescribeDiskEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiskEventsResponse, _err error) {
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

	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
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
		Action:      dara.String("DescribeDiskEvents"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiskEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the near real-time monitoring data of a disk.
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
func (client *Client) DescribeDiskMonitorDataWithContext(ctx context.Context, request *DescribeDiskMonitorDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiskMonitorDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the near real-time monitoring data of disks. You can query only the burst performance data of ESSD AutoPL disks. The data is aggregated by hour.
//
// Description:
//
// ## Usage notes
//
// CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
//
// @param request - DescribeDiskMonitorDataListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiskMonitorDataListResponse
func (client *Client) DescribeDiskMonitorDataListWithContext(ctx context.Context, request *DescribeDiskMonitorDataListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiskMonitorDataListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskIds) {
		query["DiskIds"] = request.DiskIds
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
		Action:      dara.String("DescribeDiskMonitorDataList"),
		Version:     dara.String("2021-07-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiskMonitorDataListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of replication pair-consistent groups in a specific region.
//
// Description:
//
// ## [](#)Usage notes
//
// To perform a paged query, specify the MaxResults and NextToken parameters.
//
// During a paged query, when you call the DescribeDiskReplicaGroups operation to retrieve the first page of results, set `MaxResults` to specify the maximum number of entries to return in the call. The return value of `NextToken` is a pagination token, which can be used in the next call to retrieve a new page of results. When you call the DescribeDiskReplicaGroups operation to retrieve a new page of results, set NextToken to the NextToken value returned in the previous call and set MaxResults to specify the maximum number of entries to return in this call.
//
// @param request - DescribeDiskReplicaGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiskReplicaGroupsResponse
func (client *Client) DescribeDiskReplicaGroupsWithContext(ctx context.Context, request *DescribeDiskReplicaGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiskReplicaGroupsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiskReplicaPairProgressResponse
func (client *Client) DescribeDiskReplicaPairProgressWithContext(ctx context.Context, request *DescribeDiskReplicaPairProgressRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiskReplicaPairProgressResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about replication pairs in a specific region.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which async replication is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - When you call this operation for a specific region, if the primary disk (source disk) or secondary disk (destination disk) of a replication pair resides in the region, information about the replication pair is displayed in the response.
//
//   - If you want to perform a paged query, configure the `NextToken` and `MaxResults` parameters. During a paged query, when you call the DescribeDiskReplicaPairs operation to retrieve the first page of results, set `MaxResults` to limit the maximum number of entries to return in the call. The return value of NextToken is a pagination token, which can be used in the next call to retrieve a new page of results. When you call the DescribeDiskReplicaPairs operation to retrieve a new page of results, set NextToken to the NextToken value returned in the previous call and set MaxResults to specify the maximum number of entries to return in this call.
//
// @param request - DescribeDiskReplicaPairsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiskReplicaPairsResponse
func (client *Client) DescribeDiskReplicaPairsWithContext(ctx context.Context, request *DescribeDiskReplicaPairsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiskReplicaPairsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about enterprise-level snapshot policies. When you call this operation, you can specify parameters, such as PolicyIds, ResourceGroupId, and Tag, in the request.
//
// @param request - DescribeEnterpriseSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnterpriseSnapshotPolicyResponse
func (client *Client) DescribeEnterpriseSnapshotPolicyWithContext(ctx context.Context, request *DescribeEnterpriseSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnterpriseSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the risk events of a disk.
//
// @param request - DescribeEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsResponse
func (client *Client) DescribeEventsWithContext(ctx context.Context, request *DescribeEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more Elastic Block Storage (EBS) devices that you created.
//
// @param request - DescribeLensMonitorDisksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLensMonitorDisksResponse
func (client *Client) DescribeLensMonitorDisksWithContext(ctx context.Context, request *DescribeLensMonitorDisksRequest, runtime *dara.RuntimeOptions) (_result *DescribeLensMonitorDisksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query single metric monitoring information
//
// @param tmpReq - DescribeMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricDataResponse
func (client *Client) DescribeMetricDataWithContext(ctx context.Context, tmpReq *DescribeMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePairDrillsResponse
func (client *Client) DescribePairDrillsWithContext(ctx context.Context, request *DescribePairDrillsRequest, runtime *dara.RuntimeOptions) (_result *DescribePairDrillsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of regions in which Elastic Block Storage (EBS) features (such as async replication, CloudLens for EBS, and Dedicated Block Storage Cluster) are supported.
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeReplicaGroupDrillsResponse
func (client *Client) DescribeReplicaGroupDrillsWithContext(ctx context.Context, request *DescribeReplicaGroupDrillsRequest, runtime *dara.RuntimeOptions) (_result *DescribeReplicaGroupDrillsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询解决方案实例默认配置
//
// @param request - DescribeSolutionInstanceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSolutionInstanceConfigurationResponse
func (client *Client) DescribeSolutionInstanceConfigurationWithContext(ctx context.Context, request *DescribeSolutionInstanceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DescribeSolutionInstanceConfigurationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserTagKeysResponse
func (client *Client) DescribeUserTagKeysWithContext(ctx context.Context, request *DescribeUserTagKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserTagKeysResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserTagValuesResponse
func (client *Client) DescribeUserTagValuesWithContext(ctx context.Context, request *DescribeUserTagValuesRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserTagValuesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FailoverDiskReplicaGroupResponse
func (client *Client) FailoverDiskReplicaGroupWithContext(ctx context.Context, request *FailoverDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *FailoverDiskReplicaGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FailoverDiskReplicaPairResponse
func (client *Client) FailoverDiskReplicaPairWithContext(ctx context.Context, request *FailoverDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *FailoverDiskReplicaPairResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Centralized Role: Obtain User Usage Report with reportId
//
// @param request - GetReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReportResponse
func (client *Client) GetReportWithContext(ctx context.Context, request *GetReportRequest, runtime *dara.RuntimeOptions) (_result *GetReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询异步复制支持的目的地域和可用区
//
// @param request - ListReplicaEdgeSupportedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReplicaEdgeSupportedResponse
func (client *Client) ListReplicaEdgeSupportedWithContext(ctx context.Context, request *ListReplicaEdgeSupportedRequest, runtime *dara.RuntimeOptions) (_result *ListReplicaEdgeSupportedResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Centralized Role: Query Historical Reports
//
// @param request - ListReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReportsResponse
func (client *Client) ListReportsWithContext(ctx context.Context, request *ListReportsRequest, runtime *dara.RuntimeOptions) (_result *ListReportsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to one or more Elastic Block Storage (EBS) resources, or queries the IDs and tags of resources in a specified non-default resource group.
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 中心化角色：修改App
//
// @param request - ModifyAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppResponse
func (client *Client) ModifyAppWithContext(ctx context.Context, request *ModifyAppRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改专属集群属性OpenApi
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
func (client *Client) ModifyDedicatedBlockStorageClusterAttributeWithContext(ctx context.Context, request *ModifyDedicatedBlockStorageClusterAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDedicatedBlockStorageClusterAttributeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, or recovery point objective (RPO) of a replication pair-consistent group.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which the replication pair-consistent group feature is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - The replication pair-consistent group must be in the **Created*	- (`created`) or **Stopped*	- (`stopped`) state.
//
// @param request - ModifyDiskReplicaGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskReplicaGroupResponse
func (client *Client) ModifyDiskReplicaGroupWithContext(ctx context.Context, request *ModifyDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDiskReplicaGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a replication pair.
//
// Description:
//
// ## [](#)Usage notes
//
//   - For information about the regions in which async replication is available, see [Overview](https://help.aliyun.com/document_detail/314563.html).
//
//   - Only replication pairs that are in the **Created*	- (`created`) or **Stopped*	- (`stopped`) state can have their names or descriptions modified.
//
// @param request - ModifyDiskReplicaPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskReplicaPairResponse
func (client *Client) ModifyDiskReplicaPairWithContext(ctx context.Context, request *ModifyDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *ModifyDiskReplicaPairResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse
func (client *Client) QueryDedicatedBlockStorageClusterDiskThroughputStatusWithContext(ctx context.Context, request *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query dedicated block storage cluster capacity trend data, includ available capacity size and total capacity size.
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
func (client *Client) QueryDedicatedBlockStorageClusterInventoryDataWithContext(ctx context.Context, request *QueryDedicatedBlockStorageClusterInventoryDataRequest, runtime *dara.RuntimeOptions) (_result *QueryDedicatedBlockStorageClusterInventoryDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDiskReplicaPairResponse
func (client *Client) RemoveDiskReplicaPairWithContext(ctx context.Context, request *RemoveDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *RemoveDiskReplicaPairResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReprotectDiskReplicaGroupResponse
func (client *Client) ReprotectDiskReplicaGroupWithContext(ctx context.Context, request *ReprotectDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *ReprotectDiskReplicaGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReprotectDiskReplicaPairResponse
func (client *Client) ReprotectDiskReplicaPairWithContext(ctx context.Context, request *ReprotectDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *ReprotectDiskReplicaPairResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDedicatedBlockStorageClusterDiskThroughputResponse
func (client *Client) SetDedicatedBlockStorageClusterDiskThroughputWithContext(ctx context.Context, request *SetDedicatedBlockStorageClusterDiskThroughputRequest, runtime *dara.RuntimeOptions) (_result *SetDedicatedBlockStorageClusterDiskThroughputResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the async replication feature for replication pairs that belong to a replication pair-consistent group. When the async replication feature is enabled for the pairs for the first time, the system first performs a full synchronization to synchronize all data from disks at the primary site (primary disks) to disks at the secondary site (secondary disks) and then periodically synchronizes incremental data based on the recovery point objective (RPO) of the replication pair-consistent group.
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
func (client *Client) StartDiskReplicaGroupWithContext(ctx context.Context, request *StartDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *StartDiskReplicaGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates a replication pair.
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
func (client *Client) StartDiskReplicaPairWithContext(ctx context.Context, request *StartDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *StartDiskReplicaPairResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartPairDrillResponse
func (client *Client) StartPairDrillWithContext(ctx context.Context, request *StartPairDrillRequest, runtime *dara.RuntimeOptions) (_result *StartPairDrillResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartReplicaGroupDrillResponse
func (client *Client) StartReplicaGroupDrillWithContext(ctx context.Context, request *StartReplicaGroupDrillRequest, runtime *dara.RuntimeOptions) (_result *StartReplicaGroupDrillResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDiskReplicaGroupResponse
func (client *Client) StopDiskReplicaGroupWithContext(ctx context.Context, request *StopDiskReplicaGroupRequest, runtime *dara.RuntimeOptions) (_result *StopDiskReplicaGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDiskReplicaPairResponse
func (client *Client) StopDiskReplicaPairWithContext(ctx context.Context, request *StopDiskReplicaPairRequest, runtime *dara.RuntimeOptions) (_result *StopDiskReplicaPairResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindEnterpriseSnapshotPolicyResponse
func (client *Client) UnbindEnterpriseSnapshotPolicyWithContext(ctx context.Context, request *UnbindEnterpriseSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *UnbindEnterpriseSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Search for a enterprise-level snapshot policy.
//
// @param tmpReq - UpdateEnterpriseSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnterpriseSnapshotPolicyResponse
func (client *Client) UpdateEnterpriseSnapshotPolicyWithContext(ctx context.Context, tmpReq *UpdateEnterpriseSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateEnterpriseSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新解决方案实例属性
//
// @param request - UpdateSolutionInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSolutionInstanceAttributeResponse
func (client *Client) UpdateSolutionInstanceAttributeWithContext(ctx context.Context, request *UpdateSolutionInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateSolutionInstanceAttributeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
