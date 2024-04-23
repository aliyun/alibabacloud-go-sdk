// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddDiskReplicaPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the replication pair-consistent group.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group.
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// The ID of the replication pair. You can call the [DescribeDiskReplicaPairs](~~354206~~) operation to query the IDs of existing replication pairs.
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s AddDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *AddDiskReplicaPairRequest) SetClientToken(v string) *AddDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *AddDiskReplicaPairRequest) SetRegionId(v string) *AddDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *AddDiskReplicaPairRequest) SetReplicaGroupId(v string) *AddDiskReplicaPairRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *AddDiskReplicaPairRequest) SetReplicaPairId(v string) *AddDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type AddDiskReplicaPairResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *AddDiskReplicaPairResponseBody) SetRequestId(v string) *AddDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type AddDiskReplicaPairResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *AddDiskReplicaPairResponse) SetHeaders(v map[string]*string) *AddDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *AddDiskReplicaPairResponse) SetStatusCode(v int32) *AddDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDiskReplicaPairResponse) SetBody(v *AddDiskReplicaPairResponseBody) *AddDiskReplicaPairResponse {
	s.Body = v
	return s
}

type ApplyLensServiceResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyLensServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyLensServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyLensServiceResponseBody) SetRequestId(v string) *ApplyLensServiceResponseBody {
	s.RequestId = &v
	return s
}

type ApplyLensServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyLensServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyLensServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyLensServiceResponse) GoString() string {
	return s.String()
}

func (s *ApplyLensServiceResponse) SetHeaders(v map[string]*string) *ApplyLensServiceResponse {
	s.Headers = v
	return s
}

func (s *ApplyLensServiceResponse) SetStatusCode(v int32) *ApplyLensServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyLensServiceResponse) SetBody(v *ApplyLensServiceResponseBody) *ApplyLensServiceResponse {
	s.Body = v
	return s
}

type BindEnterpriseSnapshotPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of disks.
	DiskTargets []*string `json:"DiskTargets,omitempty" xml:"DiskTargets,omitempty" type:"Repeated"`
	// The id of the policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID . You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which snapshot policy is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BindEnterpriseSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s BindEnterpriseSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *BindEnterpriseSnapshotPolicyRequest) SetClientToken(v string) *BindEnterpriseSnapshotPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *BindEnterpriseSnapshotPolicyRequest) SetDiskTargets(v []*string) *BindEnterpriseSnapshotPolicyRequest {
	s.DiskTargets = v
	return s
}

func (s *BindEnterpriseSnapshotPolicyRequest) SetPolicyId(v string) *BindEnterpriseSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *BindEnterpriseSnapshotPolicyRequest) SetRegionId(v string) *BindEnterpriseSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

type BindEnterpriseSnapshotPolicyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindEnterpriseSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindEnterpriseSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *BindEnterpriseSnapshotPolicyResponseBody) SetRequestId(v string) *BindEnterpriseSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type BindEnterpriseSnapshotPolicyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindEnterpriseSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindEnterpriseSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s BindEnterpriseSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *BindEnterpriseSnapshotPolicyResponse) SetHeaders(v map[string]*string) *BindEnterpriseSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *BindEnterpriseSnapshotPolicyResponse) SetStatusCode(v int32) *BindEnterpriseSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *BindEnterpriseSnapshotPolicyResponse) SetBody(v *BindEnterpriseSnapshotPolicyResponseBody) *BindEnterpriseSnapshotPolicyResponse {
	s.Body = v
	return s
}

type CancelLensServiceResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelLensServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelLensServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CancelLensServiceResponseBody) SetRequestId(v string) *CancelLensServiceResponseBody {
	s.RequestId = &v
	return s
}

type CancelLensServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelLensServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelLensServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelLensServiceResponse) GoString() string {
	return s.String()
}

func (s *CancelLensServiceResponse) SetHeaders(v map[string]*string) *CancelLensServiceResponse {
	s.Headers = v
	return s
}

func (s *CancelLensServiceResponse) SetStatusCode(v int32) *CancelLensServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelLensServiceResponse) SetBody(v *CancelLensServiceResponseBody) *CancelLensServiceResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the new resource group. You can view the available resource groups in the Resource Management console. For more information, see [View basic information of a resource group](~~151181~~).
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The region ID of the resource. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource. For example, if you set ResourceType to diskreplicapair, set this parameter to the ID of a replication pair.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// *   dedicatedblockstoragecluster: dedicated block storage cluster.
	// *   diskreplicapair: replication pair.
	// *   diskreplicagroup: replication pair-consistent group.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetClientToken(v string) *ChangeResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type ClearPairDrillRequest struct {
	// The ID of the drill. You can call the [DescribePairDrills](~~2584480~~) operation to query the disaster recovery drills that were performed on replication pairs in a specific region.
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The ID of the replication pair. You can call the [DescribeDiskReplicaPairs](~~354206~~) operation to query the most recent list of replication pairs, including replication pair IDs.
	PairId *string `json:"PairId,omitempty" xml:"PairId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which async replication is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ClearPairDrillRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearPairDrillRequest) GoString() string {
	return s.String()
}

func (s *ClearPairDrillRequest) SetDrillId(v string) *ClearPairDrillRequest {
	s.DrillId = &v
	return s
}

func (s *ClearPairDrillRequest) SetPairId(v string) *ClearPairDrillRequest {
	s.PairId = &v
	return s
}

func (s *ClearPairDrillRequest) SetRegionId(v string) *ClearPairDrillRequest {
	s.RegionId = &v
	return s
}

type ClearPairDrillResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearPairDrillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearPairDrillResponseBody) GoString() string {
	return s.String()
}

func (s *ClearPairDrillResponseBody) SetRequestId(v string) *ClearPairDrillResponseBody {
	s.RequestId = &v
	return s
}

type ClearPairDrillResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearPairDrillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearPairDrillResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearPairDrillResponse) GoString() string {
	return s.String()
}

func (s *ClearPairDrillResponse) SetHeaders(v map[string]*string) *ClearPairDrillResponse {
	s.Headers = v
	return s
}

func (s *ClearPairDrillResponse) SetStatusCode(v int32) *ClearPairDrillResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearPairDrillResponse) SetBody(v *ClearPairDrillResponseBody) *ClearPairDrillResponse {
	s.Body = v
	return s
}

type ClearReplicaGroupDrillRequest struct {
	// The ID of the drill. You can call the [DescribeReplicaGroupDrills](~~2584481~~) operation to query disaster recovery drills that were performed on replication pairs in a specific region.
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](~~426614~~) operation to query the most recent list of replication pair-consistent groups, including group IDs.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which async replication is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ClearReplicaGroupDrillRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearReplicaGroupDrillRequest) GoString() string {
	return s.String()
}

func (s *ClearReplicaGroupDrillRequest) SetDrillId(v string) *ClearReplicaGroupDrillRequest {
	s.DrillId = &v
	return s
}

func (s *ClearReplicaGroupDrillRequest) SetGroupId(v string) *ClearReplicaGroupDrillRequest {
	s.GroupId = &v
	return s
}

func (s *ClearReplicaGroupDrillRequest) SetRegionId(v string) *ClearReplicaGroupDrillRequest {
	s.RegionId = &v
	return s
}

type ClearReplicaGroupDrillResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearReplicaGroupDrillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearReplicaGroupDrillResponseBody) GoString() string {
	return s.String()
}

func (s *ClearReplicaGroupDrillResponseBody) SetRequestId(v string) *ClearReplicaGroupDrillResponseBody {
	s.RequestId = &v
	return s
}

type ClearReplicaGroupDrillResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearReplicaGroupDrillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearReplicaGroupDrillResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearReplicaGroupDrillResponse) GoString() string {
	return s.String()
}

func (s *ClearReplicaGroupDrillResponse) SetHeaders(v map[string]*string) *ClearReplicaGroupDrillResponse {
	s.Headers = v
	return s
}

func (s *ClearReplicaGroupDrillResponse) SetStatusCode(v int32) *ClearReplicaGroupDrillResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearReplicaGroupDrillResponse) SetBody(v *ClearReplicaGroupDrillResponseBody) *ClearReplicaGroupDrillResponse {
	s.Body = v
	return s
}

type CreateDedicatedBlockStorageClusterRequest struct {
	// The ID of the zone in which to create the dedicated block storage cluster. You can call the [DescribeZones](~~25610~~) operation to query the most recent zone list.
	Azone *string `json:"Azone,omitempty" xml:"Azone,omitempty"`
	// The capacity of the dedicated block storage cluster. Valid values: 61440 to 2334720. Unit: GiB. 2,334,720 GiB is equal to 2,280 TiB. The capacity increases in a minimum increment of 12,288 GB.
	//
	// >  If the capacity of a dedicated block storage cluster is less than 576 TiB, the maximum throughput supported per TiB does not exceed 52 MB/s. If the capacity of a dedicated block storage cluster is greater than 576 TiB, the maximum throughput supported per TiB does not exceed 26 MB/s.
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// Deprecated
	// This parameter is no longer used.
	DbscId *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	// The name of the dedicated block storage cluster.
	DbscName *string `json:"DbscName,omitempty" xml:"DbscName,omitempty"`
	// The specific number of PeriodUnit.
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration for the dedicated block storage cluster,
	// Valid values:
	//
	// *   **Month**
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the region in which to create the dedicated block storage cluster. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the dedicated block storage cluster.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource tags. You can specify up to 20 tags.
	Tag []*CreateDedicatedBlockStorageClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the dedicated block storage cluster. Valid values:
	//
	// *   Standard: basic type. When you set Type to Standard, enhanced SSDs (ESSDs) at performance level 0 (PL0 ESSDs) can be created in the dedicated block storage cluster.
	// *   Premium: performance type. When you set Type to Premium, ESSDs at performance level 1 (PL1 ESSDs) can be created in the dedicated block storage cluster.
	//
	// Default value: Premium.
	//
	// For more information about ESSDs, see [ESSDs](~~122389~~).
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDedicatedBlockStorageClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedBlockStorageClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetAzone(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.Azone = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetCapacity(v int64) *CreateDedicatedBlockStorageClusterRequest {
	s.Capacity = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetDbscId(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.DbscId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetDbscName(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.DbscName = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetPeriod(v int32) *CreateDedicatedBlockStorageClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetPeriodUnit(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetRegionId(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetResourceGroupId(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetTag(v []*CreateDedicatedBlockStorageClusterRequestTag) *CreateDedicatedBlockStorageClusterRequest {
	s.Tag = v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetType(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.Type = &v
	return s
}

type CreateDedicatedBlockStorageClusterRequestTag struct {
	// The key of tag N used for exact search of EBS resources. The tag key must be 1 to 128 characters in length. Valid values of N: 1 to 20.
	//
	// The `Tag.N` parameter pair (Tag.N.Key and Tag.N.Value) is used for exact search of EBS resources that have specified tags added. Each tag is a key-value pair.
	//
	// *   If you specify only `Tag.N.Key`, all EBS resources whose tags contain the specified tag key are returned.
	// *   If you specify only `Tag.N.Value`, the `InvalidParameter.TagValue` error is returned.
	// *   If you specify multiple tag key-value pairs at the same time, only EBS resources that match all tag key-value pairs are returned.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:` or contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDedicatedBlockStorageClusterRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedBlockStorageClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDedicatedBlockStorageClusterRequestTag) SetKey(v string) *CreateDedicatedBlockStorageClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequestTag) SetValue(v string) *CreateDedicatedBlockStorageClusterRequestTag {
	s.Value = &v
	return s
}

type CreateDedicatedBlockStorageClusterResponseBody struct {
	// The ID of the dedicated block storage cluster.
	DbscId *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	// The ID of the order.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDedicatedBlockStorageClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedBlockStorageClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) SetDbscId(v string) *CreateDedicatedBlockStorageClusterResponseBody {
	s.DbscId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) SetOrderId(v string) *CreateDedicatedBlockStorageClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) SetRequestId(v string) *CreateDedicatedBlockStorageClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateDedicatedBlockStorageClusterResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDedicatedBlockStorageClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDedicatedBlockStorageClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedBlockStorageClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateDedicatedBlockStorageClusterResponse) SetHeaders(v map[string]*string) *CreateDedicatedBlockStorageClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponse) SetStatusCode(v int32) *CreateDedicatedBlockStorageClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponse) SetBody(v *CreateDedicatedBlockStorageClusterResponseBody) *CreateDedicatedBlockStorageClusterResponse {
	s.Body = v
	return s
}

type CreateDiskReplicaGroupRequest struct {
	// The bandwidth value. Unit: Mbit/s.
	//
	// >  This parameter is not publicly available.
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the replication pair-consistent group. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID of the secondary site.
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The zone ID of the secondary site.
	DestinationZoneId *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	// The name of the replication pair-consistent group. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (\_), and hyphens (-).
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The RPO of the replication pair-consistent group. Unit: seconds. Valid value: 900.
	RPO *int64 `json:"RPO,omitempty" xml:"RPO,omitempty"`
	// The ID of the region in which to create the replication pair-consistent group. The primary site is deployed in the specified region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the replication pair-consistent group belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The zone ID of the primary site.
	SourceZoneId *string `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
	// The tags. Up to 20 tags are supported.
	Tag []*CreateDiskReplicaGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaGroupRequest) SetBandwidth(v int64) *CreateDiskReplicaGroupRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetClientToken(v string) *CreateDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetDescription(v string) *CreateDiskReplicaGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetDestinationRegionId(v string) *CreateDiskReplicaGroupRequest {
	s.DestinationRegionId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetDestinationZoneId(v string) *CreateDiskReplicaGroupRequest {
	s.DestinationZoneId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetGroupName(v string) *CreateDiskReplicaGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetRPO(v int64) *CreateDiskReplicaGroupRequest {
	s.RPO = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetRegionId(v string) *CreateDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetResourceGroupId(v string) *CreateDiskReplicaGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetSourceZoneId(v string) *CreateDiskReplicaGroupRequest {
	s.SourceZoneId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetTag(v []*CreateDiskReplicaGroupRequestTag) *CreateDiskReplicaGroupRequest {
	s.Tag = v
	return s
}

type CreateDiskReplicaGroupRequestTag struct {
	// The key of tag N of the replication pair-consistent group.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the replication pair-consistent group.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDiskReplicaGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaGroupRequestTag) SetKey(v string) *CreateDiskReplicaGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDiskReplicaGroupRequestTag) SetValue(v string) *CreateDiskReplicaGroupRequestTag {
	s.Value = &v
	return s
}

type CreateDiskReplicaGroupResponseBody struct {
	// The ID of the replication pair-consistent group.
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaGroupResponseBody) SetReplicaGroupId(v string) *CreateDiskReplicaGroupResponseBody {
	s.ReplicaGroupId = &v
	return s
}

func (s *CreateDiskReplicaGroupResponseBody) SetRequestId(v string) *CreateDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateDiskReplicaGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *CreateDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDiskReplicaGroupResponse) SetStatusCode(v int32) *CreateDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiskReplicaGroupResponse) SetBody(v *CreateDiskReplicaGroupResponseBody) *CreateDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type CreateDiskReplicaPairRequest struct {
	// The bandwidth to use to asynchronously replicate data between the primary disk and secondary disk. Unit: Kbit/s. Valid values:
	//
	// *   10240 : equal to 10 Mbit/s
	// *   20480 : equal to 20 Mbit/s
	// *   51200 : equal to 50 Mbit/s
	// *   102400 : equal to 100 Mbit/s
	//
	// Default value: 10240.
	//
	// When you set the ChargeType parameter to POSTPAY, the Bandwidth parameter is automatically set to 0 and cannot be modified. The value 0 indicates that bandwidth is dynamically allocated based on the volume of data that is asynchronously replicated from the primary disk to the secondary disk.
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method of the replication pair. Valid values:
	//
	// *   PREPAY: subscription
	// *   POSTPAY: pay-as-you-go
	//
	// Default value: POSTPAY.
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the replication pair. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the secondary disk.
	DestinationDiskId *string `json:"DestinationDiskId,omitempty" xml:"DestinationDiskId,omitempty"`
	// The region ID of the secondary disk. You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which async replication is supported.
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The zone ID of the secondary disk.
	DestinationZoneId *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	// The ID of the primary disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The name of the replication pair. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (\_), periods (.), and hyphens (-).
	PairName *string `json:"PairName,omitempty" xml:"PairName,omitempty"`
	// The subscription duration of the replication pair. This parameter is required when the `ChargeType` parameter is set to PREPAY. The unit of the subscription duration is specified by the `PeriodUnit` parameter.
	//
	// *   Valid values when the `PeriodUnit` parameter is set to Week: 1, 2, 3, and 4.
	// *   Valid values when the `PeriodUnit` parameter is set to Month: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration of the replication pair. Valid values:
	//
	// *   Week.
	// *   Month
	//
	// Default value: Month.
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The recovery point objective (RPO) of the replication pair. Unit: seconds. Set the value to 900.
	RPO *int64 `json:"RPO,omitempty" xml:"RPO,omitempty"`
	// The ID of the region in which to create the replication pair.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the replication group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The zone ID of the primary disk.
	SourceZoneId *string `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
	// The resource tags. You can specify up to 20 tags.
	Tag []*CreateDiskReplicaPairRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairRequest) SetBandwidth(v int64) *CreateDiskReplicaPairRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetChargeType(v string) *CreateDiskReplicaPairRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetClientToken(v string) *CreateDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDescription(v string) *CreateDiskReplicaPairRequest {
	s.Description = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDestinationDiskId(v string) *CreateDiskReplicaPairRequest {
	s.DestinationDiskId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDestinationRegionId(v string) *CreateDiskReplicaPairRequest {
	s.DestinationRegionId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDestinationZoneId(v string) *CreateDiskReplicaPairRequest {
	s.DestinationZoneId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDiskId(v string) *CreateDiskReplicaPairRequest {
	s.DiskId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetPairName(v string) *CreateDiskReplicaPairRequest {
	s.PairName = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetPeriod(v int64) *CreateDiskReplicaPairRequest {
	s.Period = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetPeriodUnit(v string) *CreateDiskReplicaPairRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetRPO(v int64) *CreateDiskReplicaPairRequest {
	s.RPO = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetRegionId(v string) *CreateDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetResourceGroupId(v string) *CreateDiskReplicaPairRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetSourceZoneId(v string) *CreateDiskReplicaPairRequest {
	s.SourceZoneId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetTag(v []*CreateDiskReplicaPairRequestTag) *CreateDiskReplicaPairRequest {
	s.Tag = v
	return s
}

type CreateDiskReplicaPairRequestTag struct {
	// The key of tag N to add to the resource. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:` or contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDiskReplicaPairRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaPairRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairRequestTag) SetKey(v string) *CreateDiskReplicaPairRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDiskReplicaPairRequestTag) SetValue(v string) *CreateDiskReplicaPairRequestTag {
	s.Value = &v
	return s
}

type CreateDiskReplicaPairResponseBody struct {
	// The ID of the order.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the replication pair.
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairResponseBody) SetOrderId(v string) *CreateDiskReplicaPairResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDiskReplicaPairResponseBody) SetReplicaPairId(v string) *CreateDiskReplicaPairResponseBody {
	s.ReplicaPairId = &v
	return s
}

func (s *CreateDiskReplicaPairResponseBody) SetRequestId(v string) *CreateDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type CreateDiskReplicaPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairResponse) SetHeaders(v map[string]*string) *CreateDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *CreateDiskReplicaPairResponse) SetStatusCode(v int32) *CreateDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiskReplicaPairResponse) SetBody(v *CreateDiskReplicaPairResponseBody) *CreateDiskReplicaPairResponse {
	s.Body = v
	return s
}

type CreateEnterpriseSnapshotPolicyRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Snapshot replication destination information.
	CrossRegionCopyInfo *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo `json:"CrossRegionCopyInfo,omitempty" xml:"CrossRegionCopyInfo,omitempty" type:"Struct"`
	// The description of the policy.
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The name of the policy.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID . You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which snapshot policy is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the snapshot policy.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The snapshot retention rule.
	RetainRule *CreateEnterpriseSnapshotPolicyRequestRetainRule `json:"RetainRule,omitempty" xml:"RetainRule,omitempty" type:"Struct"`
	// The rule for scheduling.
	Schedule *CreateEnterpriseSnapshotPolicyRequestSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Struct"`
	// The special snapshot retention rules.
	SpecialRetainRules *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules `json:"SpecialRetainRules,omitempty" xml:"SpecialRetainRules,omitempty" type:"Struct"`
	// The status of the policy. Valid values:
	//
	// - ENABLED: Enable snapshot policy execution.
	//
	// - DISABLED: Disable snapshot policy execution.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Advanced snapshot features.
	StorageRule *CreateEnterpriseSnapshotPolicyRequestStorageRule `json:"StorageRule,omitempty" xml:"StorageRule,omitempty" type:"Struct"`
	// The list of tags.
	Tag []*CreateEnterpriseSnapshotPolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Binding target type, valid value:
	//
	// - DISK
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetClientToken(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetCrossRegionCopyInfo(v *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) *CreateEnterpriseSnapshotPolicyRequest {
	s.CrossRegionCopyInfo = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetDesc(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.Desc = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetName(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetRegionId(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetResourceGroupId(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetRetainRule(v *CreateEnterpriseSnapshotPolicyRequestRetainRule) *CreateEnterpriseSnapshotPolicyRequest {
	s.RetainRule = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetSchedule(v *CreateEnterpriseSnapshotPolicyRequestSchedule) *CreateEnterpriseSnapshotPolicyRequest {
	s.Schedule = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetSpecialRetainRules(v *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) *CreateEnterpriseSnapshotPolicyRequest {
	s.SpecialRetainRules = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetState(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.State = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetStorageRule(v *CreateEnterpriseSnapshotPolicyRequestStorageRule) *CreateEnterpriseSnapshotPolicyRequest {
	s.StorageRule = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetTag(v []*CreateEnterpriseSnapshotPolicyRequestTag) *CreateEnterpriseSnapshotPolicyRequest {
	s.Tag = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetTargetType(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.TargetType = &v
	return s
}

type CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo struct {
	// Whether cross-region replication is enabled. The range of values:
	//
	// - true
	//
	// - false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of destination regions.
	Regions []*CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) SetEnabled(v bool) *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo {
	s.Enabled = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) SetRegions(v []*CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo {
	s.Regions = v
	return s
}

type CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions struct {
	// The region ID of the destination. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Number of days to retain the destination snapshot. The range of values is greater than 1.
	RetainDays *int32 `json:"RetainDays,omitempty" xml:"RetainDays,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) SetRegionId(v string) *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions {
	s.RegionId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) SetRetainDays(v int32) *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions {
	s.RetainDays = &v
	return s
}

type CreateEnterpriseSnapshotPolicyRequestRetainRule struct {
	// Maximum number of retained snapshots.
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// The time interval , valid value greater than 1.
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// The unit of time, valid values:
	//
	// - DAYS
	//
	// - WEEKS
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequestRetainRule) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestRetainRule) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestRetainRule) SetNumber(v int32) *CreateEnterpriseSnapshotPolicyRequestRetainRule {
	s.Number = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestRetainRule) SetTimeInterval(v int32) *CreateEnterpriseSnapshotPolicyRequestRetainRule {
	s.TimeInterval = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestRetainRule) SetTimeUnit(v string) *CreateEnterpriseSnapshotPolicyRequestRetainRule {
	s.TimeUnit = &v
	return s
}

type CreateEnterpriseSnapshotPolicyRequestSchedule struct {
	// The time when the policy will to be scheduled. Valid values: Set the parameter in a cron expression.
	//
	// For example, you can use 0 0 4 1/1 * ? to specify 04:00:00 (UTC+8) on the first day of each month.
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequestSchedule) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestSchedule) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestSchedule) SetCronExpression(v string) *CreateEnterpriseSnapshotPolicyRequestSchedule {
	s.CronExpression = &v
	return s
}

type CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules struct {
	// Indicates whether the special retention is enabled.
	//
	// *   true: enable
	// *   false: disable
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The special retention rules.
	Rules []*CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) SetEnabled(v bool) *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules {
	s.Enabled = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) SetRules(v []*CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules {
	s.Rules = v
	return s
}

type CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules struct {
	// The periodic unit for specially retained snapshots. If configured to WEEKS, it provides special retention for the first snapshot of each week. The retention period is determined by TimeUnit and TimeInterval. The range of values are:
	// - WEEKS
	// - MONTHS
	// - YEARS
	SpecialPeriodUnit *string `json:"SpecialPeriodUnit,omitempty" xml:"SpecialPeriodUnit,omitempty"`
	// Retention Time Value. The range of values is greater than 1.
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// Retention time unit for special snapshots. The range of values:
	//
	// - DAYS
	//
	// - WEEKS
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) SetSpecialPeriodUnit(v string) *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	s.SpecialPeriodUnit = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) SetTimeInterval(v int32) *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	s.TimeInterval = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) SetTimeUnit(v string) *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	s.TimeUnit = &v
	return s
}

type CreateEnterpriseSnapshotPolicyRequestStorageRule struct {
	// Whether to enable the rapid availability of snapshots. The range of values:
	//
	// - true
	//
	// - false
	EnableImmediateAccess *bool `json:"EnableImmediateAccess,omitempty" xml:"EnableImmediateAccess,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequestStorageRule) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestStorageRule) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestStorageRule) SetEnableImmediateAccess(v bool) *CreateEnterpriseSnapshotPolicyRequestStorageRule {
	s.EnableImmediateAccess = &v
	return s
}

type CreateEnterpriseSnapshotPolicyRequestTag struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be 0 to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key must have a unique tag value. You can specify at most 20 tag values in each call.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestTag) SetKey(v string) *CreateEnterpriseSnapshotPolicyRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestTag) SetValue(v string) *CreateEnterpriseSnapshotPolicyRequestTag {
	s.Value = &v
	return s
}

type CreateEnterpriseSnapshotPolicyShrinkRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Snapshot replication destination information.
	CrossRegionCopyInfoShrink *string `json:"CrossRegionCopyInfo,omitempty" xml:"CrossRegionCopyInfo,omitempty"`
	// The description of the policy.
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The name of the policy.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID . You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which snapshot policy is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the snapshot policy.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The snapshot retention rule.
	RetainRuleShrink *string `json:"RetainRule,omitempty" xml:"RetainRule,omitempty"`
	// The rule for scheduling.
	ScheduleShrink *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The special snapshot retention rules.
	SpecialRetainRulesShrink *string `json:"SpecialRetainRules,omitempty" xml:"SpecialRetainRules,omitempty"`
	// The status of the policy. Valid values:
	//
	// - ENABLED: Enable snapshot policy execution.
	//
	// - DISABLED: Disable snapshot policy execution.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Advanced snapshot features.
	StorageRuleShrink *string `json:"StorageRule,omitempty" xml:"StorageRule,omitempty"`
	// The list of tags.
	Tag []*CreateEnterpriseSnapshotPolicyShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Binding target type, valid value:
	//
	// - DISK
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetClientToken(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetCrossRegionCopyInfoShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.CrossRegionCopyInfoShrink = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetDesc(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.Desc = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetName(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetRegionId(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetResourceGroupId(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetRetainRuleShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.RetainRuleShrink = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetScheduleShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.ScheduleShrink = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetSpecialRetainRulesShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.SpecialRetainRulesShrink = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetState(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.State = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetStorageRuleShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.StorageRuleShrink = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetTag(v []*CreateEnterpriseSnapshotPolicyShrinkRequestTag) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetTargetType(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.TargetType = &v
	return s
}

type CreateEnterpriseSnapshotPolicyShrinkRequestTag struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be 0 to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key must have a unique tag value. You can specify at most 20 tag values in each call.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequestTag) SetKey(v string) *CreateEnterpriseSnapshotPolicyShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequestTag) SetValue(v string) *CreateEnterpriseSnapshotPolicyShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateEnterpriseSnapshotPolicyResponseBody struct {
	// The id of a policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyResponseBody) SetPolicyId(v string) *CreateEnterpriseSnapshotPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyResponseBody) SetRequestId(v string) *CreateEnterpriseSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateEnterpriseSnapshotPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnterpriseSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyResponse) SetHeaders(v map[string]*string) *CreateEnterpriseSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyResponse) SetStatusCode(v int32) *CreateEnterpriseSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyResponse) SetBody(v *CreateEnterpriseSnapshotPolicyResponseBody) *CreateEnterpriseSnapshotPolicyResponse {
	s.Body = v
	return s
}

type DeleteDiskRequest struct {
	// The ID of the disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of the primary or secondary disk in the replication pair. You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which async replication is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDiskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiskRequest) SetDiskId(v string) *DeleteDiskRequest {
	s.DiskId = &v
	return s
}

func (s *DeleteDiskRequest) SetRegionId(v string) *DeleteDiskRequest {
	s.RegionId = &v
	return s
}

type DeleteDiskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiskResponseBody) SetRequestId(v string) *DeleteDiskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDiskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDiskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiskResponse) SetHeaders(v map[string]*string) *DeleteDiskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiskResponse) SetStatusCode(v int32) *DeleteDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiskResponse) SetBody(v *DeleteDiskResponseBody) *DeleteDiskResponse {
	s.Body = v
	return s
}

type DeleteDiskReplicaGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the replication pair-consistent group.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](~~426614~~) operation to query the IDs of replication pair-consistent groups.
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s DeleteDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaGroupRequest) SetClientToken(v string) *DeleteDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDiskReplicaGroupRequest) SetRegionId(v string) *DeleteDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDiskReplicaGroupRequest) SetReplicaGroupId(v string) *DeleteDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

type DeleteDiskReplicaGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaGroupResponseBody) SetRequestId(v string) *DeleteDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDiskReplicaGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *DeleteDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiskReplicaGroupResponse) SetStatusCode(v int32) *DeleteDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiskReplicaGroupResponse) SetBody(v *DeleteDiskReplicaGroupResponseBody) *DeleteDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type DeleteDiskReplicaPairRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the primary disk in the replication pair. You can call the [DescribeDiskReplicaPairs](~~354206~~) operation to query the region information of replication pairs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair.
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s DeleteDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaPairRequest) SetClientToken(v string) *DeleteDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDiskReplicaPairRequest) SetRegionId(v string) *DeleteDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDiskReplicaPairRequest) SetReplicaPairId(v string) *DeleteDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type DeleteDiskReplicaPairResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaPairResponseBody) SetRequestId(v string) *DeleteDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDiskReplicaPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaPairResponse) SetHeaders(v map[string]*string) *DeleteDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiskReplicaPairResponse) SetStatusCode(v int32) *DeleteDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiskReplicaPairResponse) SetBody(v *DeleteDiskReplicaPairResponseBody) *DeleteDiskReplicaPairResponse {
	s.Body = v
	return s
}

type DeleteEnterpriseSnapshotPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The id of the policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID . You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which snapshot policy is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEnterpriseSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseSnapshotPolicyRequest) SetClientToken(v string) *DeleteEnterpriseSnapshotPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEnterpriseSnapshotPolicyRequest) SetPolicyId(v string) *DeleteEnterpriseSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeleteEnterpriseSnapshotPolicyRequest) SetRegionId(v string) *DeleteEnterpriseSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

type DeleteEnterpriseSnapshotPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnterpriseSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseSnapshotPolicyResponseBody) SetRequestId(v string) *DeleteEnterpriseSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEnterpriseSnapshotPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnterpriseSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnterpriseSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseSnapshotPolicyResponse) SetHeaders(v map[string]*string) *DeleteEnterpriseSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnterpriseSnapshotPolicyResponse) SetStatusCode(v int32) *DeleteEnterpriseSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnterpriseSnapshotPolicyResponse) SetBody(v *DeleteEnterpriseSnapshotPolicyResponseBody) *DeleteEnterpriseSnapshotPolicyResponse {
	s.Body = v
	return s
}

type DescribeDedicatedBlockStorageClusterDisksRequest struct {
	// The ID of the dedicated block storage cluster.
	DbscId *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	// The maximum number of entries to return on each page. Maximum value: 500.
	//
	// Default value: 10.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the NextToken value returned in the previous call to the DescribeDedicatedBlockStorageClusterDisks operation. Leave this parameter empty the first time you call this operation.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the dedicated block storage cluster resides. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) SetDbscId(v string) *DescribeDedicatedBlockStorageClusterDisksRequest {
	s.DbscId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) SetMaxResults(v int64) *DescribeDedicatedBlockStorageClusterDisksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) SetNextToken(v string) *DescribeDedicatedBlockStorageClusterDisksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) SetRegionId(v string) *DescribeDedicatedBlockStorageClusterDisksRequest {
	s.RegionId = &v
	return s
}

type DescribeDedicatedBlockStorageClusterDisksResponseBody struct {
	// Details about the cloud disks.
	Disks *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
	// The query token returned in this call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) SetDisks(v *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) *DescribeDedicatedBlockStorageClusterDisksResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) SetNextToken(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) SetRequestId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks struct {
	// Details about the cloud disks.
	Disk []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) SetDisk(v []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks {
	s.Disk = v
	return s
}

type DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk struct {
	// The time when the cloud disk was last attached. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mmZ format. The time is displayed in UTC.
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	// This parameter is currently in invitational preview and unavailable for general users.
	BdfId *string `json:"BdfId,omitempty" xml:"BdfId,omitempty"`
	// Whether the ESSD AutoPL disk is enabled burst IOPS / BPS. This parameter is available only if the DiskCategory parameter is set to cloud_auto. For more information, see [ESSD AutoPL disks](~~368372~~).
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of the disk. A value of cloud_essd indicates that the disk is an ESSD.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Indicates whether the automatic snapshots of the cloud disk are deleted when the disk is released. Valid values:
	//
	// *   true: The automatic snapshots of the cloud disk are deleted when the disk is released.
	// *   false: The automatic snapshots of the cloud disk are retained when the disk is released.
	//
	// Snapshots that are created by calling the [CreateSnapshot](~~25524~~) operation or by using the Elastic Compute Service (ECS) console are retained and not affected by this parameter.
	DeleteAutoSnapshot *bool `json:"DeleteAutoSnapshot,omitempty" xml:"DeleteAutoSnapshot,omitempty"`
	// Indicates whether the cloud disk is released when its associated instance is released. Valid values:
	//
	// *   true: The cloud disk is released when its associated instance is released.
	// *   false: The cloud disk is retained when its associated instance is released.
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the cloud disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the cloud disk was last detached.
	DetachedTime *string `json:"DetachedTime,omitempty" xml:"DetachedTime,omitempty"`
	// The device name of the cloud disk on its associated instance. Example: /dev/xvdb. Take note of the following items:
	//
	// *   This parameter has a value only when the `Status` value is `In_use`.
	// *   This parameter is empty for cloud disks that have the multi-attach feature enabled. You can query the attachment information of the cloud disk based on the `Attachment` values.
	//
	// >  This parameter will be removed in the future. We recommend that you use other parameters to ensure future compatibility.
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The billing method of the cloud disk. Valid values:
	//
	// *   PrePaid: subscription
	// *   PostPaid: pay-as-you-go
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The ID of the cloud disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The name of the cloud disk.
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Indicates whether the automatic snapshot policy feature is enabled for the cloud disk.
	EnableAutoSnapshot *bool `json:"EnableAutoSnapshot,omitempty" xml:"EnableAutoSnapshot,omitempty"`
	// Indicates whether the cloud disk is encrypted.
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The maximum number of IOPS.
	IOPS *int64 `json:"IOPS,omitempty" xml:"IOPS,omitempty"`
	// The ID of the image that was used to create the instance. This parameter is empty unless the cloud disk was created from an image. The value of this parameter remains unchanged throughout the lifecycle of the cloud disk.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the instance to which the cloud disk is attached. Take note of the following items:
	//
	// *   This parameter has a value only when the `Status` value is `In_use`.
	// *   This parameter is empty for cloud disks that have the multi-attach feature enabled. You can query the attachment information of the cloud disk based on the `Attachment` values.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Key Management Service (KMS) key used by the cloud disk.
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The number of instances to which the Shared Block Storage device is attached.
	MountInstanceNum *int32 `json:"MountInstanceNum,omitempty" xml:"MountInstanceNum,omitempty"`
	// Indicates whether the multi-attach feature was enabled for the cloud disk.
	MultiAttach *string `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	// The performance level of the enhanced SSD (ESSD). Valid values:
	//
	// *   PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	// *   PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// *   PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// *   PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// Indicates whether the cloud disk is removable.
	Portable *bool `json:"Portable,omitempty" xml:"Portable,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk.
	// >  This parameter is available only if the DiskCategory parameter is set to cloud_auto. For more information, see [ESSD AutoPL disks](~~368372~~) and [Modify the performance configurations of an ESSD AutoPL disk](~~413275~~).
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The region ID of cloud disk.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The size of the disk. Unit: GiB.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot that was used to create the cloud disk.
	//
	// This parameter is empty unless the cloud disk was created from a snapshot. The value of this parameter remains unchanged throughout the lifecycle of the cloud disk.
	SourceSnapshotId *string `json:"SourceSnapshotId,omitempty" xml:"SourceSnapshotId,omitempty"`
	// The state of the cloud disk. For more information, see [Disk states](~~25689~~). Valid values:
	//
	// *   In_use
	// *   Available
	// *   Attaching
	// *   Detaching
	// *   Creating
	// *   ReIniting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the dedicated block storage cluster to which the cloud disk belongs. If your cloud disk belongs to the public block storage cluster, an empty value is returned.
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
	// The ID of the storage set.
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The maximum number of partitions in the storage set.
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The tags of the cloud disk.
	Tags []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The maximum number of BPS.
	Throughput *int64 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	// The type of the disk. Valid values:
	//
	// *   system: system disk
	// *   data: data disk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The zone ID of cloud disk.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetAttachedTime(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.AttachedTime = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetBdfId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.BdfId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetBurstingEnabled(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetCategory(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDeleteAutoSnapshot(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DeleteAutoSnapshot = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDeleteWithInstance(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDescription(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Description = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDetachedTime(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DetachedTime = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDevice(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Device = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDiskChargeType(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DiskChargeType = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDiskId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDiskName(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DiskName = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetEnableAutoSnapshot(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.EnableAutoSnapshot = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetEncrypted(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Encrypted = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetIOPS(v int64) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.IOPS = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetImageId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.ImageId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetInstanceId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetKMSKeyId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetMountInstanceNum(v int32) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.MountInstanceNum = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetMultiAttach(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.MultiAttach = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetPerformanceLevel(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetPortable(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Portable = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetProvisionedIops(v int64) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetRegionId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetSize(v int32) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Size = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetSourceSnapshotId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.SourceSnapshotId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetStatus(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetStorageClusterId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.StorageClusterId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetStorageSetId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.StorageSetId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetStorageSetPartitionNumber(v int32) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetTags(v []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Tags = v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetThroughput(v int64) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Throughput = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetType(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Type = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetZoneId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.ZoneId = &v
	return s
}

type DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags struct {
	// The tag key of the cloud disk.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the cloud disk.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) SetTagKey(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) SetTagValue(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags {
	s.TagValue = &v
	return s
}

type DescribeDedicatedBlockStorageClusterDisksResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDedicatedBlockStorageClusterDisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) SetHeaders(v map[string]*string) *DescribeDedicatedBlockStorageClusterDisksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) SetStatusCode(v int32) *DescribeDedicatedBlockStorageClusterDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) SetBody(v *DescribeDedicatedBlockStorageClusterDisksResponseBody) *DescribeDedicatedBlockStorageClusterDisksResponse {
	s.Body = v
	return s
}

type DescribeDedicatedBlockStorageClustersRequest struct {
	// The zone ID of the dedicated block storage cluster. You can call the [DescribeZones](~~25610~~) operation to query the most recent zone list.
	AzoneId *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	// The category of disks that can be created in the dedicated block storage cluster.
	//
	// Set the value to cloud_essd. Only enhanced SSDs (ESSDs) can be created in dedicated block storage clusters.
	Category                       *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	ClientToken                    *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DedicatedBlockStorageClusterId []*string `json:"DedicatedBlockStorageClusterId,omitempty" xml:"DedicatedBlockStorageClusterId,omitempty" type:"Repeated"`
	MaxResults                     *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                      *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the dedicated block storage cluster. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the dedicated block storage cluster belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The states of dedicated block storage clusters. Valid values:
	//
	// *   Preparing
	// *   Running
	// *   Expired
	// *   Offline
	//
	// Multiple states can be specified. Valid values of N: 1, 2, 3, and 4.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The tags. Up to 20 tags are supported.
	Tag []*DescribeDedicatedBlockStorageClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedBlockStorageClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetAzoneId(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.AzoneId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetCategory(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetClientToken(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetDedicatedBlockStorageClusterId(v []*string) *DescribeDedicatedBlockStorageClustersRequest {
	s.DedicatedBlockStorageClusterId = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetMaxResults(v int32) *DescribeDedicatedBlockStorageClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetNextToken(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetPageNumber(v int32) *DescribeDedicatedBlockStorageClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetPageSize(v int32) *DescribeDedicatedBlockStorageClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetRegionId(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetResourceGroupId(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetStatus(v []*string) *DescribeDedicatedBlockStorageClustersRequest {
	s.Status = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetTag(v []*DescribeDedicatedBlockStorageClustersRequestTag) *DescribeDedicatedBlockStorageClustersRequest {
	s.Tag = v
	return s
}

type DescribeDedicatedBlockStorageClustersRequestTag struct {
	// The tag key of the dedicated block storage cluster.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the dedicated block storage cluster.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersRequestTag) SetKey(v string) *DescribeDedicatedBlockStorageClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequestTag) SetValue(v string) *DescribeDedicatedBlockStorageClustersRequestTag {
	s.Value = &v
	return s
}

type DescribeDedicatedBlockStorageClustersResponseBody struct {
	// The queried dedicated block storage clusters.
	DedicatedBlockStorageClusters []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters `json:"DedicatedBlockStorageClusters,omitempty" xml:"DedicatedBlockStorageClusters,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetDedicatedBlockStorageClusters(v []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.DedicatedBlockStorageClusters = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetNextToken(v string) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetPageNumber(v int32) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetPageSize(v int32) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetRequestId(v string) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetTotalCount(v int64) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters struct {
	// The unique ID (UID) of the Alibaba Cloud account.
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The category of disks that can be created in the dedicated block storage cluster.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the dedicated block storage cluster was created. The value is a UNIX timestamp. Unit: seconds.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The storage capacity of the dedicated block storage cluster.
	DedicatedBlockStorageClusterCapacity *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity `json:"DedicatedBlockStorageClusterCapacity,omitempty" xml:"DedicatedBlockStorageClusterCapacity,omitempty" type:"Struct"`
	// The ID of the dedicated block storage cluster.
	DedicatedBlockStorageClusterId *string `json:"DedicatedBlockStorageClusterId,omitempty" xml:"DedicatedBlockStorageClusterId,omitempty"`
	// The name of the dedicated block storage cluster.
	DedicatedBlockStorageClusterName *string `json:"DedicatedBlockStorageClusterName,omitempty" xml:"DedicatedBlockStorageClusterName,omitempty"`
	// The description of the dedicated block storage cluster.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The dedicated block storage cluster enable thin provison.
	EnableThinProvision *bool `json:"EnableThinProvision,omitempty" xml:"EnableThinProvision,omitempty"`
	// The time when the dedicated block storage cluster expires. The value is a UNIX timestamp. Unit: seconds.
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The performance level of disks. Valid values:
	//
	// *   PL0
	// *   PL1
	// *   PL2
	// *   PL3
	//
	// >  This parameter takes effect only if Category is set to cloud_essd.
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The region ID of the dedicated block storage cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the dedicated block storage cluster belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Oversold ratio of cluster space capacity for creating disk.
	SizeOverSoldRatio *float64 `json:"SizeOverSoldRatio,omitempty" xml:"SizeOverSoldRatio,omitempty"`
	// The state of the dedicated block storage cluster. Valid values:
	//
	// *   Preparing
	// *   Running
	// *   Expired
	// *   Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// StorageDomain.
	StorageDomain *string `json:"StorageDomain,omitempty" xml:"StorageDomain,omitempty"`
	// This parameter is not supported.
	SupportedCategory *string `json:"SupportedCategory,omitempty" xml:"SupportedCategory,omitempty"`
	// The tags of the dedicated block storage cluster.
	Tags []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the dedicated block storage cluster. Valid values:
	//
	// *   Standard: basic dedicated block storage cluster. ESSDs at performance level 0 (PL0 ESSDs) can be created in basic dedicated block storage clusters.
	// *   Premium: performance dedicated block storage cluster. ESSDs at performance level 1 (PL1 ESSDs) can be created in performance dedicated block storage clusters.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The zone ID of the dedicated block storage cluster.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetAliUid(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.AliUid = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetCategory(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetCreateTime(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.CreateTime = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetDedicatedBlockStorageClusterCapacity(v *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.DedicatedBlockStorageClusterCapacity = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetDedicatedBlockStorageClusterId(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.DedicatedBlockStorageClusterId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetDedicatedBlockStorageClusterName(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.DedicatedBlockStorageClusterName = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetDescription(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Description = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetEnableThinProvision(v bool) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.EnableThinProvision = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetExpiredTime(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetPerformanceLevel(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetRegionId(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetResourceGroupId(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetSizeOverSoldRatio(v float64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.SizeOverSoldRatio = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetStatus(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetStorageDomain(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.StorageDomain = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetSupportedCategory(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.SupportedCategory = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetTags(v []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Tags = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetType(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Type = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetZoneId(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.ZoneId = &v
	return s
}

type DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity struct {
	// The available capacity of the dedicated block storage cluster. Unit: GiB.
	AvailableCapacity *int64 `json:"AvailableCapacity,omitempty" xml:"AvailableCapacity,omitempty"`
	// The total capacity of the dedicated block storage cluster that was delivered in disk creation orders. Unit: GB.
	AvailableDeviceCapacity *int64 `json:"AvailableDeviceCapacity,omitempty" xml:"AvailableDeviceCapacity,omitempty"`
	// This parameter is displayed only if Thin Provision is enabled.
	AvailableSpaceCapacity *float64 `json:"AvailableSpaceCapacity,omitempty" xml:"AvailableSpaceCapacity,omitempty"`
	// The capacity of the dedicated block storage cluster that was delivered in orders. Unit: GB.
	ClusterAvailableCapacity *int64 `json:"ClusterAvailableCapacity,omitempty" xml:"ClusterAvailableCapacity,omitempty"`
	// The capacity of the dedicated block storage cluster that is to be delivered in orders. Unit: GB.
	ClusterDeliveryCapacity *int64 `json:"ClusterDeliveryCapacity,omitempty" xml:"ClusterDeliveryCapacity,omitempty"`
	// The to-be-delivered capacity of the dedicated block storage cluster. Unit: GB.
	DeliveryCapacity *int64 `json:"DeliveryCapacity,omitempty" xml:"DeliveryCapacity,omitempty"`
	// The total capacity of the dedicated block storage cluster. Unit: GiB.
	TotalCapacity *int64 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	// The total capacity of the dedicated block storage cluster that is to be delivered in disk creation orders. Unit: GB.
	TotalDeviceCapacity *int64 `json:"TotalDeviceCapacity,omitempty" xml:"TotalDeviceCapacity,omitempty"`
	// This parameter is displayed only if Thin Provision is enabled.
	TotalSpaceCapacity *int64 `json:"TotalSpaceCapacity,omitempty" xml:"TotalSpaceCapacity,omitempty"`
	// The used capacity of the dedicated block storage cluster. Unit: GB.
	UsedCapacity *int64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
	// The capacity of the dedicated block storage cluster that was used to create disks. Unit: GB.
	UsedDeviceCapacity *int64 `json:"UsedDeviceCapacity,omitempty" xml:"UsedDeviceCapacity,omitempty"`
	// This parameter is displayed only if Thin Provision is enabled.
	UsedSpaceCapacity *float64 `json:"UsedSpaceCapacity,omitempty" xml:"UsedSpaceCapacity,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetAvailableCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.AvailableCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetAvailableDeviceCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.AvailableDeviceCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetAvailableSpaceCapacity(v float64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.AvailableSpaceCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetClusterAvailableCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.ClusterAvailableCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetClusterDeliveryCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.ClusterDeliveryCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetDeliveryCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.DeliveryCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetTotalCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.TotalCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetTotalDeviceCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.TotalDeviceCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetTotalSpaceCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.TotalSpaceCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetUsedCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.UsedCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetUsedDeviceCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.UsedDeviceCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetUsedSpaceCapacity(v float64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.UsedSpaceCapacity = &v
	return s
}

type DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags struct {
	// The tag key of the dedicated block storage cluster.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the dedicated block storage cluster.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) SetTagKey(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) SetTagValue(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags {
	s.TagValue = &v
	return s
}

type DescribeDedicatedBlockStorageClustersResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDedicatedBlockStorageClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponse) SetHeaders(v map[string]*string) *DescribeDedicatedBlockStorageClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponse) SetStatusCode(v int32) *DescribeDedicatedBlockStorageClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponse) SetBody(v *DescribeDedicatedBlockStorageClustersResponseBody) *DescribeDedicatedBlockStorageClustersResponse {
	s.Body = v
	return s
}

type DescribeDiskEventsRequest struct {
	// The type of the disk. Valid values:
	//
	// *   cloud_efficiency: ultra disk.
	// *   cloud_ssd: standard SSD.
	// *   cloud_essd: enhanced SSD (ESSD).
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The ID of the disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The end of the time range to query. Specify the time in the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100.
	//
	// Default values:
	//
	// *   If this parameter is not specified or is set to a value smaller than 10, the default value is 10.
	// *   If this parameter is set to a value greater than 100, the default value is 100.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in this request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the disk. You can call the [DescribeRegions](~~354276~~) operation to query the list of regions that support CloudLens for EBS.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The event type. Set the value to DataNeedProtect, which indicates that the disk data needs to be protected.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDiskEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskEventsRequest) SetDiskCategory(v string) *DescribeDiskEventsRequest {
	s.DiskCategory = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetDiskId(v string) *DescribeDiskEventsRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetEndTime(v string) *DescribeDiskEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetMaxResults(v int64) *DescribeDiskEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetNextToken(v string) *DescribeDiskEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetRegionId(v string) *DescribeDiskEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetStartTime(v string) *DescribeDiskEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetType(v string) *DescribeDiskEventsRequest {
	s.Type = &v
	return s
}

type DescribeDiskEventsResponseBody struct {
	// The risk events of the disk.
	DiskEvents []*DescribeDiskEventsResponseBodyDiskEvents `json:"DiskEvents,omitempty" xml:"DiskEvents,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiskEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskEventsResponseBody) SetDiskEvents(v []*DescribeDiskEventsResponseBodyDiskEvents) *DescribeDiskEventsResponseBody {
	s.DiskEvents = v
	return s
}

func (s *DescribeDiskEventsResponseBody) SetNextToken(v string) *DescribeDiskEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskEventsResponseBody) SetRequestId(v string) *DescribeDiskEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskEventsResponseBody) SetTotalCount(v int64) *DescribeDiskEventsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDiskEventsResponseBodyDiskEvents struct {
	// The description of the event.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The recommended action after the event occurred. Valid values:
	//
	// *   Resize: resizes the disk.
	// *   ModifyDiskSpec: changes the category of the disk.
	// *   NoAction: performs no operation.
	RecommendAction *string `json:"RecommendAction,omitempty" xml:"RecommendAction,omitempty"`
	// The region ID of the disk.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The state of the event. Valid values:
	//
	// *   Solved
	// *   UnSolved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the event occurred. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The type of the event. Only DataNeedProtect can be returned.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDiskEventsResponseBodyDiskEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskEventsResponseBodyDiskEvents) GoString() string {
	return s.String()
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetDescription(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.Description = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetDiskId(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetRecommendAction(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.RecommendAction = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetRegionId(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetStatus(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.Status = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetTimestamp(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.Timestamp = &v
	return s
}

func (s *DescribeDiskEventsResponseBodyDiskEvents) SetType(v string) *DescribeDiskEventsResponseBodyDiskEvents {
	s.Type = &v
	return s
}

type DescribeDiskEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskEventsResponse) SetHeaders(v map[string]*string) *DescribeDiskEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskEventsResponse) SetStatusCode(v int32) *DescribeDiskEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskEventsResponse) SetBody(v *DescribeDiskEventsResponseBody) *DescribeDiskEventsResponse {
	s.Body = v
	return s
}

type DescribeDiskMonitorDataRequest struct {
	// The ID of the disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The end of the time range during which you want to query the near real-time monitoring data of the disk. Specify the time in the [ISO 8601](~~25696~~) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The interval at which the near real-time monitoring data is collected. Unit: seconds. Valid values:
	//
	// *   5
	// *   60
	//
	// Default value: 5.
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the disk.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range during which you want to query the near real-time monitoring data of the disk. Specify the time in the [ISO 8601](~~25696~~) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the monitoring data. Valid values:
	//
	// *   basic: baseline performance data.
	// *   pro: burst performance data, such as burst I/O operations.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDiskMonitorDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataRequest) SetDiskId(v string) *DescribeDiskMonitorDataRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetEndTime(v string) *DescribeDiskMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetPeriod(v int64) *DescribeDiskMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetRegionId(v string) *DescribeDiskMonitorDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetStartTime(v string) *DescribeDiskMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetType(v string) *DescribeDiskMonitorDataRequest {
	s.Type = &v
	return s
}

type DescribeDiskMonitorDataResponseBody struct {
	// The near real-time monitoring data of the disk.
	MonitorData []*DescribeDiskMonitorDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiskMonitorDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataResponseBody) SetMonitorData(v []*DescribeDiskMonitorDataResponseBodyMonitorData) *DescribeDiskMonitorDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeDiskMonitorDataResponseBody) SetRequestId(v string) *DescribeDiskMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBody) SetTotalCount(v int64) *DescribeDiskMonitorDataResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDiskMonitorDataResponseBodyMonitorData struct {
	// The percentage of BPS.
	BPSPercent *int64 `json:"BPSPercent,omitempty" xml:"BPSPercent,omitempty"`
	// The number of burst I/O operations.
	BurstIOCount *int64 `json:"BurstIOCount,omitempty" xml:"BurstIOCount,omitempty"`
	// The ID of the disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The percentage of IOPS.
	IOPSPercent *int64 `json:"IOPSPercent,omitempty" xml:"IOPSPercent,omitempty"`
	// The read bandwidth of the disk. Unit: MByte/s.
	ReadBPS *int64 `json:"ReadBPS,omitempty" xml:"ReadBPS,omitempty"`
	// Read IO block size. Unit: Bytes
	ReadBlockSize *int64 `json:"ReadBlockSize,omitempty" xml:"ReadBlockSize,omitempty"`
	// The maximum number of read IOPS.
	ReadIOPS *int64 `json:"ReadIOPS,omitempty" xml:"ReadIOPS,omitempty"`
	// Read IO latency. Unit:  microsecond
	ReadLatency *int64 `json:"ReadLatency,omitempty" xml:"ReadLatency,omitempty"`
	// The timestamp that is used to query the near real-time monitoring data of the disk. The time follows the [ISO 8601](~~25696~~) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The write bandwidth of the disk. Unit: MByte/s.
	WriteBPS *int64 `json:"WriteBPS,omitempty" xml:"WriteBPS,omitempty"`
	// Write IO block size. Unit: Bytes
	WriteBlockSize *int64 `json:"WriteBlockSize,omitempty" xml:"WriteBlockSize,omitempty"`
	// The maximum number of write IOPS.
	WriteIOPS *int64 `json:"WriteIOPS,omitempty" xml:"WriteIOPS,omitempty"`
	// Write IO latency. Unit: microsecond
	WriteLatency *int64 `json:"WriteLatency,omitempty" xml:"WriteLatency,omitempty"`
}

func (s DescribeDiskMonitorDataResponseBodyMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskMonitorDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetBPSPercent(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.BPSPercent = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetBurstIOCount(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.BurstIOCount = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetDiskId(v string) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetIOPSPercent(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.IOPSPercent = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetReadBPS(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.ReadBPS = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetReadBlockSize(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.ReadBlockSize = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetReadIOPS(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.ReadIOPS = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetReadLatency(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.ReadLatency = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetTimestamp(v string) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.Timestamp = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetWriteBPS(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.WriteBPS = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetWriteBlockSize(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.WriteBlockSize = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetWriteIOPS(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.WriteIOPS = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetWriteLatency(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.WriteLatency = &v
	return s
}

type DescribeDiskMonitorDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskMonitorDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataResponse) SetHeaders(v map[string]*string) *DescribeDiskMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskMonitorDataResponse) SetStatusCode(v int32) *DescribeDiskMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskMonitorDataResponse) SetBody(v *DescribeDiskMonitorDataResponseBody) *DescribeDiskMonitorDataResponse {
	s.Body = v
	return s
}

type DescribeDiskMonitorDataListRequest struct {
	// The IDs of the disks. The value is a JSON array that contains multiple disk IDs. Separate the IDs with commas (,).
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The end of the time range during which you want to query the near real-time monitoring data of the disks. Specify the time in the [ISO 8601](~~25696~~) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries per page. If you specify this parameter, both `MaxResults` and `NextToken` are used for a paged query.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in this request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~354276~~) operation to query the list of regions that support CloudLens for EBS.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range during which you want to query the near real-time monitoring data of the disks. Specify the time in the [ISO 8601](~~25696~~) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the monitoring data. Set the value to pro.
	//
	// pro: burst performance data, such as burst I/O operations.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDiskMonitorDataListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskMonitorDataListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataListRequest) SetDiskIds(v string) *DescribeDiskMonitorDataListRequest {
	s.DiskIds = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) SetEndTime(v string) *DescribeDiskMonitorDataListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) SetMaxResults(v string) *DescribeDiskMonitorDataListRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) SetNextToken(v string) *DescribeDiskMonitorDataListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) SetRegionId(v string) *DescribeDiskMonitorDataListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) SetStartTime(v string) *DescribeDiskMonitorDataListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) SetType(v string) *DescribeDiskMonitorDataListRequest {
	s.Type = &v
	return s
}

type DescribeDiskMonitorDataListResponseBody struct {
	// The near real-time monitoring data of the disks.
	MonitorData []*DescribeDiskMonitorDataListResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiskMonitorDataListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskMonitorDataListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataListResponseBody) SetMonitorData(v []*DescribeDiskMonitorDataListResponseBodyMonitorData) *DescribeDiskMonitorDataListResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeDiskMonitorDataListResponseBody) SetNextToken(v string) *DescribeDiskMonitorDataListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskMonitorDataListResponseBody) SetRequestId(v string) *DescribeDiskMonitorDataListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskMonitorDataListResponseBody) SetTotalCount(v int64) *DescribeDiskMonitorDataListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDiskMonitorDataListResponseBodyMonitorData struct {
	// The number of burst I/O operations.
	BurstIOCount *int64 `json:"BurstIOCount,omitempty" xml:"BurstIOCount,omitempty"`
	// The ID of the disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The beginning of the time range during which the performance of the disk bursts. The time follows the [ISO 8601](~~25696~~) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeDiskMonitorDataListResponseBodyMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskMonitorDataListResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataListResponseBodyMonitorData) SetBurstIOCount(v int64) *DescribeDiskMonitorDataListResponseBodyMonitorData {
	s.BurstIOCount = &v
	return s
}

func (s *DescribeDiskMonitorDataListResponseBodyMonitorData) SetDiskId(v string) *DescribeDiskMonitorDataListResponseBodyMonitorData {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskMonitorDataListResponseBodyMonitorData) SetTimestamp(v string) *DescribeDiskMonitorDataListResponseBodyMonitorData {
	s.Timestamp = &v
	return s
}

type DescribeDiskMonitorDataListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskMonitorDataListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskMonitorDataListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskMonitorDataListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataListResponse) SetHeaders(v map[string]*string) *DescribeDiskMonitorDataListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskMonitorDataListResponse) SetStatusCode(v int32) *DescribeDiskMonitorDataListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskMonitorDataListResponse) SetBody(v *DescribeDiskMonitorDataListResponseBody) *DescribeDiskMonitorDataListResponse {
	s.Body = v
	return s
}

type DescribeDiskReplicaGroupsRequest struct {
	// The IDs of replication pair-consistent groups. You can specify the IDs of one or more replication pair-consistent groups. Separate the IDs with commas (,).
	//
	// This parameter is empty by default, which indicates that all replication pair-consistent groups in the specified region are queried.
	GroupIds *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// The maximum number of entries to return on each page. Valid values: 1 to 500.
	//
	// Default value: 10.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the NextToken value returned in the previous call to the DescribeDiskReplicaGroups operation. Leave this parameter empty the first time you call this operation. When NextToken is specified, the PageSize and PageNumber request parameters do not take effect and the TotalCount response parameter is invalid.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of the page to return.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the replication pair-consistent group.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the replication group belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the site from which the information of replication pair-consistent groups is retrieved. This parameter is used for scenarios where data is replicated across zones in replication pairs.
	//
	// *   If the Site parameter is not specified, information such as the state of replication pair-consistent groups at the primary site is queried and returned.
	//
	// *   Otherwise, information such as the state of replication pair-consistent groups at the site specified by the Site parameter is queried and returned. Valid values:
	//
	//     *   production: primary site
	//     *   backup: secondary site
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// The resource tags. You can specify up to 20 tags.
	Tag []*DescribeDiskReplicaGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDiskReplicaGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsRequest) SetGroupIds(v string) *DescribeDiskReplicaGroupsRequest {
	s.GroupIds = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetMaxResults(v int64) *DescribeDiskReplicaGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetNextToken(v string) *DescribeDiskReplicaGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetPageNumber(v int32) *DescribeDiskReplicaGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetPageSize(v int32) *DescribeDiskReplicaGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetRegionId(v string) *DescribeDiskReplicaGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetResourceGroupId(v string) *DescribeDiskReplicaGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetSite(v string) *DescribeDiskReplicaGroupsRequest {
	s.Site = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetTag(v []*DescribeDiskReplicaGroupsRequestTag) *DescribeDiskReplicaGroupsRequest {
	s.Tag = v
	return s
}

type DescribeDiskReplicaGroupsRequestTag struct {
	// The key of tag N of the replication group.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the replication group.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDiskReplicaGroupsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsRequestTag) SetKey(v string) *DescribeDiskReplicaGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequestTag) SetValue(v string) *DescribeDiskReplicaGroupsRequestTag {
	s.Value = &v
	return s
}

type DescribeDiskReplicaGroupsResponseBody struct {
	// The query token returned in this call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Details about the replication pair-consistent groups.
	ReplicaGroups []*DescribeDiskReplicaGroupsResponseBodyReplicaGroups `json:"ReplicaGroups,omitempty" xml:"ReplicaGroups,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiskReplicaGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetNextToken(v string) *DescribeDiskReplicaGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetPageNumber(v int32) *DescribeDiskReplicaGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetPageSize(v int32) *DescribeDiskReplicaGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetReplicaGroups(v []*DescribeDiskReplicaGroupsResponseBodyReplicaGroups) *DescribeDiskReplicaGroupsResponseBody {
	s.ReplicaGroups = v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetRequestId(v string) *DescribeDiskReplicaGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetTotalCount(v int64) *DescribeDiskReplicaGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDiskReplicaGroupsResponseBodyReplicaGroups struct {
	// The bandwidth value. Unit: Mbit/s. This parameter is unavailable and has a system-preset value.
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The description of the replication pair-consistent group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the region in which the secondary site is deployed.
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The ID of the zone in which the secondary site is deployed.
	DestinationZoneId *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	// The name of the replication pair-consistent group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when data was last replicated from the primary disks to the secondary disks in the replication pair-consistent group. The value of this parameter is a timestamp. Unit: seconds.
	LastRecoverPoint *int64 `json:"LastRecoverPoint,omitempty" xml:"LastRecoverPoint,omitempty"`
	// The IDs of the replications pairs that belong to the replication pair-consistent group.
	PairIds [][]byte `json:"PairIds,omitempty" xml:"PairIds,omitempty" type:"Repeated"`
	// The number of replications pairs that belong to the replication pair-consistent group.
	PairNumber *int64 `json:"PairNumber,omitempty" xml:"PairNumber,omitempty"`
	// The initial source region (primary region) of the replication pair-consistent group.
	PrimaryRegion *string `json:"PrimaryRegion,omitempty" xml:"PrimaryRegion,omitempty"`
	// The initial source zone (primary zone) of the replication pair-consistent group.
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// The recovery point objective (RPO) of the replication pair-consistent group. Unit: seconds.
	RPO *int64 `json:"RPO,omitempty" xml:"RPO,omitempty"`
	// The ID of the replication pair-consistent group.
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// The ID of the resource group to which the replication group belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the site from which the information of the replication pair and replication pair-consistent group is obtained. Valid values:
	//
	// *   production: primary site
	// *   backup: secondary site
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// The ID of the region in which the primary site is deployed.
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// The ID of the zone in which the primary site is deployed.
	SourceZoneId *string `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
	// The initial destination region (secondary region) of the replication pair-consistent group.
	StandbyRegion *string `json:"StandbyRegion,omitempty" xml:"StandbyRegion,omitempty"`
	// The initial destination zone (secondary zone) of the replication pair-consistent group.
	StandbyZone *string `json:"StandbyZone,omitempty" xml:"StandbyZone,omitempty"`
	// The state of the replication pair-consistent group. Valid values:
	//
	// *   invalid: The replication pair-consistent group is invalid, which indicates that abnormal replication pairs are present in the replication pair-consistent group.
	// *   creating: The replication pair-consistent group is being created.
	// *   created: The replication pair-consistent group is created.
	// *   create_failed: The replication pair-consistent group cannot be created.
	// *   manual_syncing: Data is being manually synchronized between the disks in the replication pair-consistent group. The first time data is being manually synchronized between the disks in a replication pair-consistent group, the replication pair-consistent group is in this state.
	// *   syncing: Data is being synchronized between the disks in the replication pair-consistent group. While data is being asynchronously replicated from the primary disks to the secondary disks not for the first time, the replication pair-consistent group is in this state.
	// *   normal: The replication pair-consistent group is working as expected. When the system finishes replicating data from the primary disks to the secondary disks within the current replication cycle, the replication pair-consistent group enters this state.
	// *   stopping: The replication pair-consistent group is being stopped.
	// *   stopped: The replication pair-consistent group is stopped.
	// *   stop_failed: The replication pair-consistent group cannot be stopped.
	// *   failovering: A failover is being performed.
	// *   failovered: A failover is performed.
	// *   failover_failed: A failover cannot be performed.
	// *   reprotecting: A reverse replication is being performed.
	// *   reprotect_failed: A reverse replication cannot be performed.
	// *   deleting: The replication pair-consistent group is being deleted.
	// *   delete_failed: The replication pair-consistent group cannot be deleted.
	// *   deleted: The replication pair-consistent group is deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the replication pair.
	Tags []*DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeDiskReplicaGroupsResponseBodyReplicaGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetBandwidth(v int64) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Bandwidth = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetDescription(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Description = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetDestinationRegionId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.DestinationRegionId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetDestinationZoneId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.DestinationZoneId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetGroupName(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetLastRecoverPoint(v int64) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.LastRecoverPoint = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetPairIds(v [][]byte) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.PairIds = v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetPairNumber(v int64) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.PairNumber = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetPrimaryRegion(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.PrimaryRegion = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetPrimaryZone(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetRPO(v int64) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.RPO = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetReplicaGroupId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.ReplicaGroupId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetResourceGroupId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetSite(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Site = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetSourceRegionId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.SourceRegionId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetSourceZoneId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.SourceZoneId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetStandbyRegion(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.StandbyRegion = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetStandbyZone(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.StandbyZone = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetStatus(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Status = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetTags(v []*DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Tags = v
	return s
}

type DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags struct {
	// The tag key of the replication group.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the replication group.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) SetTagKey(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) SetTagValue(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags {
	s.TagValue = &v
	return s
}

type DescribeDiskReplicaGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskReplicaGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskReplicaGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsResponse) SetHeaders(v map[string]*string) *DescribeDiskReplicaGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskReplicaGroupsResponse) SetStatusCode(v int32) *DescribeDiskReplicaGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponse) SetBody(v *DescribeDiskReplicaGroupsResponseBody) *DescribeDiskReplicaGroupsResponse {
	s.Body = v
	return s
}

type DescribeDiskReplicaPairProgressRequest struct {
	// The region ID of the replication pair.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair. You can call the [DescribeDiskReplicaPairs](~~354206~~)operation to query the IDs of existing replication pairs.
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s DescribeDiskReplicaPairProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairProgressRequest) SetRegionId(v string) *DescribeDiskReplicaPairProgressRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskReplicaPairProgressRequest) SetReplicaPairId(v string) *DescribeDiskReplicaPairProgressRequest {
	s.ReplicaPairId = &v
	return s
}

type DescribeDiskReplicaPairProgressResponseBody struct {
	// The replication progress of the replication pair.
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The timestamp that indicates the last recovery point in time. The value is returned only after the replication pair works for replicating data.
	RecoverPoint *int64 `json:"RecoverPoint,omitempty" xml:"RecoverPoint,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiskReplicaPairProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairProgressResponseBody) SetProgress(v int32) *DescribeDiskReplicaPairProgressResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeDiskReplicaPairProgressResponseBody) SetRecoverPoint(v int64) *DescribeDiskReplicaPairProgressResponseBody {
	s.RecoverPoint = &v
	return s
}

func (s *DescribeDiskReplicaPairProgressResponseBody) SetRequestId(v string) *DescribeDiskReplicaPairProgressResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDiskReplicaPairProgressResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskReplicaPairProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskReplicaPairProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairProgressResponse) SetHeaders(v map[string]*string) *DescribeDiskReplicaPairProgressResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskReplicaPairProgressResponse) SetStatusCode(v int32) *DescribeDiskReplicaPairProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskReplicaPairProgressResponse) SetBody(v *DescribeDiskReplicaPairProgressResponseBody) *DescribeDiskReplicaPairProgressResponse {
	s.Body = v
	return s
}

type DescribeDiskReplicaPairsRequest struct {
	// The maximum number of entries per page. You can use this parameter together with NextToken.
	//
	// Valid values: 1 to 500.
	//
	// Default value: 10.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken. If you specify NextToken, the PageSize and PageNumber request parameters do not take effect, and the TotalCount response parameter is invalid.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of replication pairs. You can specify the IDs of one or more replication pairs and separate the IDs with commas (,). Example: `pair-cn-dsa****,pair-cn-asd****`.
	//
	// This parameter is empty by default, which indicates that all replication pairs in the specified region are queried. You can specify a maximum of 100 replication pair IDs.
	PairIds *string `json:"PairIds,omitempty" xml:"PairIds,omitempty"`
	// The region ID of the primary or secondary disk in the replication pair. You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which async replication is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group. You can specify the ID of a replication pair-consistent group to query the replication pairs in the group. Example: `pg-****`.
	//
	// This parameter is empty by default, which indicates that all replication pairs in the specified region are queried.
	//
	// >  If this parameter is set to`-`, replication pairs that are not added to any replication pair-consistent groups are returned.
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// The ID of the resource group to which the replication pair belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the site from which the information of replication pairs is retrieved. Valid value:
	//
	// *   production: primary site
	// *   backup: secondary site
	//
	// Default value: production.
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// The tags. Up to 20 tags are supported.
	Tag []*DescribeDiskReplicaPairsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDiskReplicaPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsRequest) SetMaxResults(v int64) *DescribeDiskReplicaPairsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetNextToken(v string) *DescribeDiskReplicaPairsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetPageNumber(v int32) *DescribeDiskReplicaPairsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetPageSize(v int32) *DescribeDiskReplicaPairsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetPairIds(v string) *DescribeDiskReplicaPairsRequest {
	s.PairIds = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetRegionId(v string) *DescribeDiskReplicaPairsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetReplicaGroupId(v string) *DescribeDiskReplicaPairsRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetResourceGroupId(v string) *DescribeDiskReplicaPairsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetSite(v string) *DescribeDiskReplicaPairsRequest {
	s.Site = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetTag(v []*DescribeDiskReplicaPairsRequestTag) *DescribeDiskReplicaPairsRequest {
	s.Tag = v
	return s
}

type DescribeDiskReplicaPairsRequestTag struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDiskReplicaPairsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsRequestTag) SetKey(v string) *DescribeDiskReplicaPairsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequestTag) SetValue(v string) *DescribeDiskReplicaPairsRequestTag {
	s.Value = &v
	return s
}

type DescribeDiskReplicaPairsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Details about the replication pairs.
	ReplicaPairs []*DescribeDiskReplicaPairsResponseBodyReplicaPairs `json:"ReplicaPairs,omitempty" xml:"ReplicaPairs,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiskReplicaPairsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponseBody) SetNextToken(v string) *DescribeDiskReplicaPairsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetPageNumber(v int32) *DescribeDiskReplicaPairsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetPageSize(v int32) *DescribeDiskReplicaPairsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetReplicaPairs(v []*DescribeDiskReplicaPairsResponseBodyReplicaPairs) *DescribeDiskReplicaPairsResponseBody {
	s.ReplicaPairs = v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetRequestId(v string) *DescribeDiskReplicaPairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetTotalCount(v int64) *DescribeDiskReplicaPairsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDiskReplicaPairsResponseBodyReplicaPairs struct {
	// The bandwidth used to asynchronously replicate data from the primary disk to the secondary disk. Unit: Kbit/s.
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method of the replication pair. Valid values:
	//
	// *   PREPAY: subscription
	// *   POSTPAY: pay-as-you-go
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the replication pair was created. The value of this parameter is a timestamp. Unit: seconds.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the replication pair.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the secondary disk.
	DestinationDiskId *string `json:"DestinationDiskId,omitempty" xml:"DestinationDiskId,omitempty"`
	// The region ID of the secondary disk.
	DestinationRegion *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	// The zone ID of the secondary disk.
	DestinationZoneId *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	// The time when the replication pair expires. The value of this parameter is a timestamp. Unit: seconds.
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The time when data was last replicated from the primary disk to the secondary disk in the replication pair. The value of this parameter is a timestamp. Unit: seconds. 86,400 seconds is equivalent to 24 hours.
	LastRecoverPoint *int64 `json:"LastRecoverPoint,omitempty" xml:"LastRecoverPoint,omitempty"`
	// The name of the replication pair.
	PairName *string `json:"PairName,omitempty" xml:"PairName,omitempty"`
	// The initial source region (primary region) of the replication pair.
	PrimaryRegion *string `json:"PrimaryRegion,omitempty" xml:"PrimaryRegion,omitempty"`
	// The initial source zone (primary zone) of the replication pair.
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// The recovery point objective (RPO) of the replication pair. Unit: seconds.
	RPO *int64 `json:"RPO,omitempty" xml:"RPO,omitempty"`
	// The ID of the replication pair-consistent group to which the replication pair belongs.
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// The name of the replication pair-consistent group to which the replication pair belongs.
	ReplicaGroupName *string `json:"ReplicaGroupName,omitempty" xml:"ReplicaGroupName,omitempty"`
	// The ID of the replication pair.
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
	// The ID of the resource group to which the replication pair belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the site from which the information about the replication pairs and replication pair-consistent group was obtained. Valid values:
	//
	// *   production: primary site
	// *   backup: secondary site
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// The ID of the primary disk.
	SourceDiskId *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	// The region ID of the primary disk.
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// The zone ID of the primary disk.
	SourceZoneId *string `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
	// The initial destination region (secondary region) of the replication pair.
	StandbyRegion *string `json:"StandbyRegion,omitempty" xml:"StandbyRegion,omitempty"`
	// The initial destination zone (secondary zone) of the replication pair.
	StandbyZone *string `json:"StandbyZone,omitempty" xml:"StandbyZone,omitempty"`
	// The status of the replication pair. Valid values:
	//
	// *   invalid: The replication pair was invalid. When a replication pair becomes abnormal, it enters this state.
	// *   creating: The replication pair was being created.
	// *   created: The replication pair was created.
	// *   create_failed: The replication pair failed to be created.
	// *   initial_syncing: Data was synchronized from the primary disk to the secondary disk for the first time. After a replication pair is created and activated, the replication pair is in this state the first time data is synchronized from the primary disk to the secondary disk.
	// *   manual_syncing: Data was being manually synchronized from the primary disk to the secondary disk. After data is manually synchronized from the primary disk to the secondary disk, the replication pair returns to the stopped state. The first time data is manually synchronized from the primary disk to the secondary disk, the replication pair is in the manual_syncing state during the synchronization.
	// *   syncing: Data was being synchronized from the primary disk to the secondary disk. When data is being asynchronously replicated from the primary disk to the secondary disk again in subsequent operations, the replication pair is in this state.
	// *   normal: The replication pair was working as expected. When the system finishes replicating data from the primary disk to the secondary disk within the current replication cycle, the replication pair enters this state.
	// *   stopping: The replication pair was being stopped.
	// *   stopped: The replication pair was stopped.
	// *   stop_failed: The replication pair failed to be stopped.
	// *   failovering: A failover was being performed.
	// *   failovered: A failover was performed.
	// *   failover_failed: A failover failed to be performed.
	// *   reprotecting: A reverse replication was being performed.
	// *   reprotect_failed: A reverse replication failed to be performed.
	// *   deleting: The replication pair was being deleted.
	// *   delete_failed: The replication pair failed to be deleted.
	// *   deleted: The replication pair was deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message that describes the state of the replication pair. This parameter has a value when `Status` has a value of invalid or `create_failed`. Valid values:
	//
	// *   PrePayOrderExpired: The replication pair has expired.
	// *   PostPayOrderCeaseService: The pay-as-you-go replication pair has been stopped due to an overdue payment.
	// *   DeviceRemoved: The primary or secondary disk has been deleted.
	// *   DeviceKeyChanged: The `DeviceKey` mapping of the primary or secondary disk has changed.
	// *   DeviceSizeChanged: The `DeviceSize` value of the primary or secondary disk has changed.
	// *   OperationDenied.QuotaExceed: The maximum number of replication pairs that can be created has been reached.
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The tags of the replication pair.
	Tags []*DescribeDiskReplicaPairsResponseBodyReplicaPairsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeDiskReplicaPairsResponseBodyReplicaPairs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponseBodyReplicaPairs) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetBandwidth(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Bandwidth = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetChargeType(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ChargeType = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetCreateTime(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.CreateTime = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDescription(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Description = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDestinationDiskId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.DestinationDiskId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDestinationRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.DestinationRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDestinationZoneId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.DestinationZoneId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetExpiredTime(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetLastRecoverPoint(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.LastRecoverPoint = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetPairName(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.PairName = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetPrimaryRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.PrimaryRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetPrimaryZone(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetRPO(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.RPO = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetReplicaGroupId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ReplicaGroupId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetReplicaGroupName(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ReplicaGroupName = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetReplicaPairId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ReplicaPairId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetResourceGroupId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSite(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Site = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSourceDiskId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.SourceDiskId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSourceRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.SourceRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSourceZoneId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.SourceZoneId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStandbyRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.StandbyRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStandbyZone(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.StandbyZone = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStatus(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Status = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStatusMessage(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.StatusMessage = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetTags(v []*DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Tags = v
	return s
}

type DescribeDiskReplicaPairsResponseBodyReplicaPairsTags struct {
	// The key of the tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) SetTagKey(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairsTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) SetTagValue(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairsTags {
	s.TagValue = &v
	return s
}

type DescribeDiskReplicaPairsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskReplicaPairsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskReplicaPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponse) SetHeaders(v map[string]*string) *DescribeDiskReplicaPairsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskReplicaPairsResponse) SetStatusCode(v int32) *DescribeDiskReplicaPairsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponse) SetBody(v *DescribeDiskReplicaPairsResponseBody) *DescribeDiskReplicaPairsResponse {
	s.Body = v
	return s
}

type DescribeDisksRequest struct {
	// The category of disks.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The IDs of the disks. The value is a JSON array that contains multiple disk IDs. Separate the IDs with commas (,).
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The name of the disk.
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100.
	//
	// Default values:
	//
	// *   If this parameter is not specified or is set to a value smaller than 10, the default value is 10.
	// *   If this parameter is set to a value greater than 100, the default value is 100.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the disk belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The states of disk. Valid values:
	//
	// *   In_use
	// *   Available
	// *   Attaching
	// *   Detaching
	// *   Creating
	// *   ReIniting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The resource tags. You can specify up to 20 tags.
	Tag []*DescribeDisksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDisksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisksRequest) SetCategory(v string) *DescribeDisksRequest {
	s.Category = &v
	return s
}

func (s *DescribeDisksRequest) SetDiskIds(v string) *DescribeDisksRequest {
	s.DiskIds = &v
	return s
}

func (s *DescribeDisksRequest) SetDiskName(v string) *DescribeDisksRequest {
	s.DiskName = &v
	return s
}

func (s *DescribeDisksRequest) SetMaxResults(v int32) *DescribeDisksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDisksRequest) SetNextToken(v string) *DescribeDisksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDisksRequest) SetPageNumber(v int32) *DescribeDisksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDisksRequest) SetPageSize(v int32) *DescribeDisksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDisksRequest) SetRegionId(v string) *DescribeDisksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDisksRequest) SetResourceGroupId(v string) *DescribeDisksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDisksRequest) SetStatus(v string) *DescribeDisksRequest {
	s.Status = &v
	return s
}

func (s *DescribeDisksRequest) SetTag(v []*DescribeDisksRequestTag) *DescribeDisksRequest {
	s.Tag = v
	return s
}

func (s *DescribeDisksRequest) SetZoneId(v string) *DescribeDisksRequest {
	s.ZoneId = &v
	return s
}

type DescribeDisksRequestTag struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDisksRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisksRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDisksRequestTag) SetKey(v string) *DescribeDisksRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDisksRequestTag) SetValue(v string) *DescribeDisksRequestTag {
	s.Value = &v
	return s
}

type DescribeDisksResponseBody struct {
	// The data returned.
	Data []*DescribeDisksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDisksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBody) SetData(v []*DescribeDisksResponseBodyData) *DescribeDisksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDisksResponseBody) SetNextToken(v string) *DescribeDisksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDisksResponseBody) SetPageNumber(v int32) *DescribeDisksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDisksResponseBody) SetPageSize(v int32) *DescribeDisksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDisksResponseBody) SetRequestId(v string) *DescribeDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisksResponseBody) SetTotalCount(v int64) *DescribeDisksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDisksResponseBodyData struct {
	// The access authorization information of the disk.
	AccessAuthorization []*DescribeDisksResponseBodyDataAccessAuthorization `json:"AccessAuthorization,omitempty" xml:"AccessAuthorization,omitempty" type:"Repeated"`
	// The time when the disk was attached.
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	// The attachment information of the disk.
	Attachments []*DescribeDisksResponseBodyDataAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	// Specifies whether to enable the performance burst feature for the disk. Valid values:
	//
	// *   true: enable the performance burst feature.
	// *   false: disable the performance burst feature.
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of disk.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the disk was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the disk was detached.
	DetachedTime *string `json:"DetachedTime,omitempty" xml:"DetachedTime,omitempty"`
	// The maximum number of read and write operations per second. Unit: operations/s.
	DeviceNguid *string `json:"DeviceNguid,omitempty" xml:"DeviceNguid,omitempty"`
	// The ID of the disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The name of the disk.
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the disk. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false.
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// IOPS.
	Iops *int64 `json:"Iops,omitempty" xml:"Iops,omitempty"`
	// The maximum number of read operations per second. Unit: operations/s.
	IopsRead *int64 `json:"IopsRead,omitempty" xml:"IopsRead,omitempty"`
	// The maximum number of write operations per second. Unit: operations/s.
	IopsWrite *int64 `json:"IopsWrite,omitempty" xml:"IopsWrite,omitempty"`
	// The ID of the KMS key.
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The time when the disk was modified.
	ModifyAt *string `json:"ModifyAt,omitempty" xml:"ModifyAt,omitempty"`
	// The performance level of the ESSD. Valid values:
	//
	// *   PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	// *   PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// *   PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// *   PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1.
	//
	// For more information about ESSD performance levels, see [ESSDs](~~122389~~).
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}
	//
	// Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}
	//
	// >  This parameter is available only if the DiskCategory parameter is set to cloud_auto. For more information, see [ESSD AutoPL disks](~~368372~~) and [Modify the performance configurations of an ESSD AutoPL disk](~~413275~~).
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The region ID of the disk.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the disk belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The size of the disk. Unit: GB.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The states of disk. Valid values:
	//
	// *   In_use
	// *   Available
	// *   Attaching
	// *   Detaching
	// *   Creating
	// *   ReIniting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The details of the tags.
	Tags []*DescribeDisksResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The throughput of the disk.
	//
	// Unit: MB/s.
	Throughput *int64 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	// The zone ID of the disk.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDisksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyData) SetAccessAuthorization(v []*DescribeDisksResponseBodyDataAccessAuthorization) *DescribeDisksResponseBodyData {
	s.AccessAuthorization = v
	return s
}

func (s *DescribeDisksResponseBodyData) SetAttachedTime(v string) *DescribeDisksResponseBodyData {
	s.AttachedTime = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetAttachments(v []*DescribeDisksResponseBodyDataAttachments) *DescribeDisksResponseBodyData {
	s.Attachments = v
	return s
}

func (s *DescribeDisksResponseBodyData) SetBurstingEnabled(v bool) *DescribeDisksResponseBodyData {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetCategory(v string) *DescribeDisksResponseBodyData {
	s.Category = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetCreationTime(v string) *DescribeDisksResponseBodyData {
	s.CreationTime = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetDescription(v string) *DescribeDisksResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetDetachedTime(v string) *DescribeDisksResponseBodyData {
	s.DetachedTime = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetDeviceNguid(v string) *DescribeDisksResponseBodyData {
	s.DeviceNguid = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetDiskId(v string) *DescribeDisksResponseBodyData {
	s.DiskId = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetDiskName(v string) *DescribeDisksResponseBodyData {
	s.DiskName = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetEncrypted(v bool) *DescribeDisksResponseBodyData {
	s.Encrypted = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetIops(v int64) *DescribeDisksResponseBodyData {
	s.Iops = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetIopsRead(v int64) *DescribeDisksResponseBodyData {
	s.IopsRead = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetIopsWrite(v int64) *DescribeDisksResponseBodyData {
	s.IopsWrite = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetKmsKeyId(v string) *DescribeDisksResponseBodyData {
	s.KmsKeyId = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetModifyAt(v string) *DescribeDisksResponseBodyData {
	s.ModifyAt = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetPerformanceLevel(v string) *DescribeDisksResponseBodyData {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetProvisionedIops(v int64) *DescribeDisksResponseBodyData {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetRegionId(v string) *DescribeDisksResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetResourceGroupId(v string) *DescribeDisksResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetSize(v int64) *DescribeDisksResponseBodyData {
	s.Size = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetStatus(v string) *DescribeDisksResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetTags(v []*DescribeDisksResponseBodyDataTags) *DescribeDisksResponseBodyData {
	s.Tags = v
	return s
}

func (s *DescribeDisksResponseBodyData) SetThroughput(v int64) *DescribeDisksResponseBodyData {
	s.Throughput = &v
	return s
}

func (s *DescribeDisksResponseBodyData) SetZoneId(v string) *DescribeDisksResponseBodyData {
	s.ZoneId = &v
	return s
}

type DescribeDisksResponseBodyDataAccessAuthorization struct {
	// The token that is used to access the disk.
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The information of ecs.
	Ecs *DescribeDisksResponseBodyDataAccessAuthorizationEcs `json:"Ecs,omitempty" xml:"Ecs,omitempty" type:"Struct"`
	// The information of pod.
	Pod *DescribeDisksResponseBodyDataAccessAuthorizationPod `json:"Pod,omitempty" xml:"Pod,omitempty" type:"Struct"`
}

func (s DescribeDisksResponseBodyDataAccessAuthorization) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisksResponseBodyDataAccessAuthorization) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDataAccessAuthorization) SetAccessToken(v string) *DescribeDisksResponseBodyDataAccessAuthorization {
	s.AccessToken = &v
	return s
}

func (s *DescribeDisksResponseBodyDataAccessAuthorization) SetEcs(v *DescribeDisksResponseBodyDataAccessAuthorizationEcs) *DescribeDisksResponseBodyDataAccessAuthorization {
	s.Ecs = v
	return s
}

func (s *DescribeDisksResponseBodyDataAccessAuthorization) SetPod(v *DescribeDisksResponseBodyDataAccessAuthorizationPod) *DescribeDisksResponseBodyDataAccessAuthorization {
	s.Pod = v
	return s
}

type DescribeDisksResponseBodyDataAccessAuthorizationEcs struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDisksResponseBodyDataAccessAuthorizationEcs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisksResponseBodyDataAccessAuthorizationEcs) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDataAccessAuthorizationEcs) SetInstanceId(v string) *DescribeDisksResponseBodyDataAccessAuthorizationEcs {
	s.InstanceId = &v
	return s
}

type DescribeDisksResponseBodyDataAccessAuthorizationPod struct {
	// The cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The pod ID.
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
}

func (s DescribeDisksResponseBodyDataAccessAuthorizationPod) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisksResponseBodyDataAccessAuthorizationPod) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDataAccessAuthorizationPod) SetClusterId(v string) *DescribeDisksResponseBodyDataAccessAuthorizationPod {
	s.ClusterId = &v
	return s
}

func (s *DescribeDisksResponseBodyDataAccessAuthorizationPod) SetPodId(v string) *DescribeDisksResponseBodyDataAccessAuthorizationPod {
	s.PodId = &v
	return s
}

type DescribeDisksResponseBodyDataAttachments struct {
	// The time when the disk was attached.
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	NamespaceId *int32 `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeDisksResponseBodyDataAttachments) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisksResponseBodyDataAttachments) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDataAttachments) SetAttachedTime(v string) *DescribeDisksResponseBodyDataAttachments {
	s.AttachedTime = &v
	return s
}

func (s *DescribeDisksResponseBodyDataAttachments) SetInstanceId(v string) *DescribeDisksResponseBodyDataAttachments {
	s.InstanceId = &v
	return s
}

func (s *DescribeDisksResponseBodyDataAttachments) SetNamespaceId(v int32) *DescribeDisksResponseBodyDataAttachments {
	s.NamespaceId = &v
	return s
}

type DescribeDisksResponseBodyDataTags struct {
	// The tag key.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDisksResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisksResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDataTags) SetTagKey(v string) *DescribeDisksResponseBodyDataTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDisksResponseBodyDataTags) SetTagValue(v string) *DescribeDisksResponseBodyDataTags {
	s.TagValue = &v
	return s
}

type DescribeDisksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDisksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponse) SetHeaders(v map[string]*string) *DescribeDisksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisksResponse) SetStatusCode(v int32) *DescribeDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisksResponse) SetBody(v *DescribeDisksResponseBody) *DescribeDisksResponse {
	s.Body = v
	return s
}

type DescribeEnterpriseSnapshotPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of disks.
	DiskIds []*string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of enterprise-level snapshot policies.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the enterprise-level snapshot policy belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the enterprise-level snapshot policies. Valid values of N: 1 to 20.
	Tag []*DescribeEnterpriseSnapshotPolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeEnterpriseSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetClientToken(v string) *DescribeEnterpriseSnapshotPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetDiskIds(v []*string) *DescribeEnterpriseSnapshotPolicyRequest {
	s.DiskIds = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetMaxResults(v int32) *DescribeEnterpriseSnapshotPolicyRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetNextToken(v string) *DescribeEnterpriseSnapshotPolicyRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetPageNumber(v int32) *DescribeEnterpriseSnapshotPolicyRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetPageSize(v int32) *DescribeEnterpriseSnapshotPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetPolicyIds(v []*string) *DescribeEnterpriseSnapshotPolicyRequest {
	s.PolicyIds = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetRegionId(v string) *DescribeEnterpriseSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetResourceGroupId(v string) *DescribeEnterpriseSnapshotPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetTag(v []*DescribeEnterpriseSnapshotPolicyRequestTag) *DescribeEnterpriseSnapshotPolicyRequest {
	s.Tag = v
	return s
}

type DescribeEnterpriseSnapshotPolicyRequestTag struct {
	// The key of tag N of the enterprise-level snapshot policy.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the enterprise-level snapshot policy.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyRequestTag) SetKey(v string) *DescribeEnterpriseSnapshotPolicyRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequestTag) SetValue(v string) *DescribeEnterpriseSnapshotPolicyRequestTag {
	s.Value = &v
	return s
}

type DescribeEnterpriseSnapshotPolicyResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned enterprise-level snapshot policies.
	Policies []*DescribeEnterpriseSnapshotPolicyResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) SetNextToken(v string) *DescribeEnterpriseSnapshotPolicyResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) SetPageNumber(v int32) *DescribeEnterpriseSnapshotPolicyResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) SetPageSize(v int32) *DescribeEnterpriseSnapshotPolicyResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) SetPolicies(v []*DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) *DescribeEnterpriseSnapshotPolicyResponseBody {
	s.Policies = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) SetRequestId(v string) *DescribeEnterpriseSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) SetTotalCount(v int64) *DescribeEnterpriseSnapshotPolicyResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPolicies struct {
	// The time when the enterprise-level snapshot policy was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The replication rule of snapshots in the enterprise-level snapshot policy.
	CrossRegionCopyInfo *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo `json:"CrossRegionCopyInfo,omitempty" xml:"CrossRegionCopyInfo,omitempty" type:"Struct"`
	// The description of the enterprise-level snapshot policy.
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The list of disks.
	DiskIds []*string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty" type:"Repeated"`
	// Indicates whether snapshots are managed.
	ManagedForEcs *bool `json:"ManagedForEcs,omitempty" xml:"ManagedForEcs,omitempty"`
	// The name of the enterprise-level snapshot policy.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the enterprise-level snapshot policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// the resource group
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The retention rule of the enterprise-level snapshot policy.
	RetainRule *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule `json:"RetainRule,omitempty" xml:"RetainRule,omitempty" type:"Struct"`
	// The scheduling rule of the enterprise-level snapshot policy.
	Schedule *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Struct"`
	// The special retention rules of the enterprise-level snapshot policy.
	SpecialRetainRules *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules `json:"SpecialRetainRules,omitempty" xml:"SpecialRetainRules,omitempty" type:"Struct"`
	// The status of the enterprise-level snapshot policy.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The storage rule of snapshots in the enterprise-level snapshot policy.
	StorageRule *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule `json:"StorageRule,omitempty" xml:"StorageRule,omitempty" type:"Struct"`
	// the pair tags
	Tags []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The number of objects that are associated with the enterprise-level snapshot policy.
	TargetCount *int32 `json:"TargetCount,omitempty" xml:"TargetCount,omitempty"`
	// The type of the enterprise-level snapshot policy.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetCreateTime(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.CreateTime = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetCrossRegionCopyInfo(v *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.CrossRegionCopyInfo = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetDesc(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.Desc = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetDiskIds(v []*string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.DiskIds = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetManagedForEcs(v bool) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.ManagedForEcs = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetName(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.Name = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetPolicyId(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetResourceGroupId(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetRetainRule(v *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.RetainRule = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetSchedule(v *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.Schedule = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetSpecialRetainRules(v *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.SpecialRetainRules = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetState(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.State = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetStorageRule(v *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.StorageRule = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetTags(v []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.Tags = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetTargetCount(v int32) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.TargetCount = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetTargetType(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.TargetType = &v
	return s
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo struct {
	// Indicates whether the cross-region replication feature is enabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The destination regions that store snapshot copies.
	Regions []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) SetEnabled(v bool) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo {
	s.Enabled = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) SetRegions(v []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo {
	s.Regions = v
	return s
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions struct {
	// The ID of the destination region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The retention period of snapshot copies in the destination region. Unit: day.
	RetainDays *int32 `json:"RetainDays,omitempty" xml:"RetainDays,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) SetRegionId(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) SetRetainDays(v int32) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions {
	s.RetainDays = &v
	return s
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule struct {
	// The maximum number of snapshots that can be retained.
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// The value of the retention period of snapshots.
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// The unit of the retention period of snapshots.
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) SetNumber(v int32) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule {
	s.Number = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) SetTimeInterval(v int32) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule {
	s.TimeInterval = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) SetTimeUnit(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule {
	s.TimeUnit = &v
	return s
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule struct {
	// The cron expression of the enterprise-level snapshot policy.
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule) SetCronExpression(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule {
	s.CronExpression = &v
	return s
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules struct {
	// Indicates whether the special retention period is enabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The special retention rules.
	Rules []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) SetEnabled(v bool) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules {
	s.Enabled = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) SetRules(v []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules {
	s.Rules = v
	return s
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules struct {
	// The unit of the special retention period.
	SpecialPeriodUnit *string `json:"SpecialPeriodUnit,omitempty" xml:"SpecialPeriodUnit,omitempty"`
	// The value of the retention period.
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// The unit of the retention period.
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) SetSpecialPeriodUnit(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules {
	s.SpecialPeriodUnit = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) SetTimeInterval(v int32) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules {
	s.TimeInterval = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) SetTimeUnit(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules {
	s.TimeUnit = &v
	return s
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule struct {
	// Indicates whether the instant access feature is enabled.
	EnableImmediateAccess *bool `json:"EnableImmediateAccess,omitempty" xml:"EnableImmediateAccess,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule) SetEnableImmediateAccess(v bool) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule {
	s.EnableImmediateAccess = &v
	return s
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags struct {
	// The key of the tag of the enterprise-level snapshot policy.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag of the enterprise-level snapshot policy.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) SetTagKey(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags {
	s.TagKey = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) SetTagValue(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags {
	s.TagValue = &v
	return s
}

type DescribeEnterpriseSnapshotPolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnterpriseSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponse) SetHeaders(v map[string]*string) *DescribeEnterpriseSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponse) SetStatusCode(v int32) *DescribeEnterpriseSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponse) SetBody(v *DescribeEnterpriseSnapshotPolicyResponseBody) *DescribeEnterpriseSnapshotPolicyResponse {
	s.Body = v
	return s
}

type DescribeEventsRequest struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventName    *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsRequest) SetEndTime(v string) *DescribeEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventsRequest) SetEventName(v string) *DescribeEventsRequest {
	s.EventName = &v
	return s
}

func (s *DescribeEventsRequest) SetMaxResults(v int32) *DescribeEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeEventsRequest) SetNextToken(v string) *DescribeEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeEventsRequest) SetRegionId(v string) *DescribeEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventsRequest) SetResourceId(v string) *DescribeEventsRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeEventsRequest) SetResourceType(v string) *DescribeEventsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeEventsRequest) SetStartTime(v string) *DescribeEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEventsRequest) SetStatus(v string) *DescribeEventsRequest {
	s.Status = &v
	return s
}

type DescribeEventsResponseBody struct {
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceEvents []*DescribeEventsResponseBodyResourceEvents `json:"ResourceEvents,omitempty" xml:"ResourceEvents,omitempty" type:"Repeated"`
	TotalCount     *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBody) SetNextToken(v string) *DescribeEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeEventsResponseBody) SetRequestId(v string) *DescribeEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventsResponseBody) SetResourceEvents(v []*DescribeEventsResponseBodyResourceEvents) *DescribeEventsResponseBody {
	s.ResourceEvents = v
	return s
}

func (s *DescribeEventsResponseBody) SetTotalCount(v int32) *DescribeEventsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEventsResponseBodyResourceEvents struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventLevel      *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	EventName       *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	EventType       *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	RecommendAction *string `json:"RecommendAction,omitempty" xml:"RecommendAction,omitempty"`
	RecommendParams *string `json:"RecommendParams,omitempty" xml:"RecommendParams,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventsResponseBodyResourceEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyResourceEvents) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyResourceEvents) SetDescription(v string) *DescribeEventsResponseBodyResourceEvents {
	s.Description = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetEndTime(v string) *DescribeEventsResponseBodyResourceEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetEventLevel(v string) *DescribeEventsResponseBodyResourceEvents {
	s.EventLevel = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetEventName(v string) *DescribeEventsResponseBodyResourceEvents {
	s.EventName = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetEventType(v string) *DescribeEventsResponseBodyResourceEvents {
	s.EventType = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetRecommendAction(v string) *DescribeEventsResponseBodyResourceEvents {
	s.RecommendAction = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetRecommendParams(v string) *DescribeEventsResponseBodyResourceEvents {
	s.RecommendParams = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetResourceId(v string) *DescribeEventsResponseBodyResourceEvents {
	s.ResourceId = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetResourceType(v string) *DescribeEventsResponseBodyResourceEvents {
	s.ResourceType = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetStartTime(v string) *DescribeEventsResponseBodyResourceEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetStatus(v string) *DescribeEventsResponseBodyResourceEvents {
	s.Status = &v
	return s
}

type DescribeEventsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponse) SetHeaders(v map[string]*string) *DescribeEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventsResponse) SetStatusCode(v int32) *DescribeEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventsResponse) SetBody(v *DescribeEventsResponseBody) *DescribeEventsResponse {
	s.Body = v
	return s
}

type DescribeLensServiceStatusResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of CloudLens for EBS. Valid values:
	//
	// *   Applying
	// *   UnAvailable
	// *   Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeLensServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLensServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLensServiceStatusResponseBody) SetRequestId(v string) *DescribeLensServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLensServiceStatusResponseBody) SetStatus(v string) *DescribeLensServiceStatusResponseBody {
	s.Status = &v
	return s
}

type DescribeLensServiceStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLensServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLensServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLensServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeLensServiceStatusResponse) SetHeaders(v map[string]*string) *DescribeLensServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeLensServiceStatusResponse) SetStatusCode(v int32) *DescribeLensServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLensServiceStatusResponse) SetBody(v *DescribeLensServiceStatusResponseBody) *DescribeLensServiceStatusResponse {
	s.Body = v
	return s
}

type DescribeMetricDataRequest struct {
	// The dimension map in the JSON format. A dimension is a key-value pair. Valid dimension key: diskId.
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The end of the time range to query. The specified time must be later than the current time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the metric.
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The interval at which metric data is collected. Unit: seconds. Default value: 60. Valid values: 60, 300, 600, and 3600, which support queries for time ranges of up to 2 hours, 2 hours, 1 day, and 7 days, respectively. For example, if you set Period to 60, the end time is less than 2 hours from the end time.
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. You can specify a point in time that is up to one year apart from the current time. If StartTime and EndTime are both unspecified, the monitoring metric data of the last Period value is queried. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMetricDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataRequest) SetDimensions(v string) *DescribeMetricDataRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricDataRequest) SetEndTime(v string) *DescribeMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricDataRequest) SetMetricName(v string) *DescribeMetricDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricDataRequest) SetPeriod(v int32) *DescribeMetricDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricDataRequest) SetRegionId(v string) *DescribeMetricDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricDataRequest) SetStartTime(v string) *DescribeMetricDataRequest {
	s.StartTime = &v
	return s
}

type DescribeMetricDataResponseBody struct {
	// The disk monitoring data.
	DataList []*DescribeMetricDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of data entries queried.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMetricDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataResponseBody) SetDataList(v []*DescribeMetricDataResponseBodyDataList) *DescribeMetricDataResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeMetricDataResponseBody) SetRequestId(v string) *DescribeMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricDataResponseBody) SetTotalCount(v int32) *DescribeMetricDataResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeMetricDataResponseBodyDataList struct {
	// The datapoints that consist of consecutive timestamps in seconds and metric values that were recorded at these time points.
	Datapoints interface{} `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// The tags.
	Labels interface{} `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s DescribeMetricDataResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataResponseBodyDataList) SetDatapoints(v interface{}) *DescribeMetricDataResponseBodyDataList {
	s.Datapoints = v
	return s
}

func (s *DescribeMetricDataResponseBodyDataList) SetLabels(v interface{}) *DescribeMetricDataResponseBodyDataList {
	s.Labels = v
	return s
}

type DescribeMetricDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataResponse) SetHeaders(v map[string]*string) *DescribeMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricDataResponse) SetStatusCode(v int32) *DescribeMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricDataResponse) SetBody(v *DescribeMetricDataResponseBody) *DescribeMetricDataResponse {
	s.Body = v
	return s
}

type DescribePairDrillsRequest struct {
	// The ID of the drill.
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The maximum number of entries to be returned. You can use this parameter together with NextToken.
	//
	// Valid values: 1 to 500.
	//
	// Default value: 10.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Set the value to the NextToken value returned in the previous call to the DescribeDiskReplicaPairs operation. Leave this parameter empty the first time you call this operation. When you specify NextToken, the PageSize and PageNumber request parameters do not take effect and the TotalCount response parameter is invalid.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the replication pair. You can call the [DescribeDiskReplicaPairs](~~354206~~) operation to query a list of asynchronous replication pairs, including replication pair IDs.
	PairId *string `json:"PairId,omitempty" xml:"PairId,omitempty"`
	// The region ID of the primary or secondary disk in the async replication pair. You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which async replication is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePairDrillsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePairDrillsRequest) GoString() string {
	return s.String()
}

func (s *DescribePairDrillsRequest) SetDrillId(v string) *DescribePairDrillsRequest {
	s.DrillId = &v
	return s
}

func (s *DescribePairDrillsRequest) SetMaxResults(v int64) *DescribePairDrillsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePairDrillsRequest) SetNextToken(v string) *DescribePairDrillsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePairDrillsRequest) SetPageNumber(v int32) *DescribePairDrillsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePairDrillsRequest) SetPageSize(v int32) *DescribePairDrillsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePairDrillsRequest) SetPairId(v string) *DescribePairDrillsRequest {
	s.PairId = &v
	return s
}

func (s *DescribePairDrillsRequest) SetRegionId(v string) *DescribePairDrillsRequest {
	s.RegionId = &v
	return s
}

type DescribePairDrillsResponseBody struct {
	// The information of disaster recovery drills that were performed on the replication pair.
	Drills []*DescribePairDrillsResponseBodyDrills `json:"Drills,omitempty" xml:"Drills,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePairDrillsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePairDrillsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePairDrillsResponseBody) SetDrills(v []*DescribePairDrillsResponseBodyDrills) *DescribePairDrillsResponseBody {
	s.Drills = v
	return s
}

func (s *DescribePairDrillsResponseBody) SetNextToken(v string) *DescribePairDrillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePairDrillsResponseBody) SetPageNumber(v int32) *DescribePairDrillsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePairDrillsResponseBody) SetPageSize(v int32) *DescribePairDrillsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePairDrillsResponseBody) SetRequestId(v string) *DescribePairDrillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePairDrillsResponseBody) SetTotalCount(v int64) *DescribePairDrillsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribePairDrillsResponseBodyDrills struct {
	// The ID of the drill disk.
	DrillDiskId *string `json:"DrillDiskId,omitempty" xml:"DrillDiskId,omitempty"`
	// The status of the drill disk. Valid values:
	//
	// *   created
	// *   deleted
	// *   creating
	// *   deleting
	//
	// >  This parameter can also display error code details if your drill disk fails to be created or deleted.
	DrillDiskStatus *string `json:"DrillDiskStatus,omitempty" xml:"DrillDiskStatus,omitempty"`
	// The ID of the drill.
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The recovery point of the drill. The value of this parameter is a timestamp. Unit: seconds.
	RecoverPoint *int64 `json:"RecoverPoint,omitempty" xml:"RecoverPoint,omitempty"`
	// The beginning time of the drill. The value of this parameter is a timestamp. Unit: seconds.
	StartAt *int64 `json:"StartAt,omitempty" xml:"StartAt,omitempty"`
	// The status of the drill. Valid values:
	//
	// *   execute_failed
	// *   executed
	// *   executing
	// *   clear_failed
	// *   clearing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The error message that was displayed if the drill failed to be executed.
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DescribePairDrillsResponseBodyDrills) String() string {
	return tea.Prettify(s)
}

func (s DescribePairDrillsResponseBodyDrills) GoString() string {
	return s.String()
}

func (s *DescribePairDrillsResponseBodyDrills) SetDrillDiskId(v string) *DescribePairDrillsResponseBodyDrills {
	s.DrillDiskId = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) SetDrillDiskStatus(v string) *DescribePairDrillsResponseBodyDrills {
	s.DrillDiskStatus = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) SetDrillId(v string) *DescribePairDrillsResponseBodyDrills {
	s.DrillId = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) SetRecoverPoint(v int64) *DescribePairDrillsResponseBodyDrills {
	s.RecoverPoint = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) SetStartAt(v int64) *DescribePairDrillsResponseBodyDrills {
	s.StartAt = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) SetStatus(v string) *DescribePairDrillsResponseBodyDrills {
	s.Status = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) SetStatusMessage(v string) *DescribePairDrillsResponseBodyDrills {
	s.StatusMessage = &v
	return s
}

type DescribePairDrillsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePairDrillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePairDrillsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePairDrillsResponse) GoString() string {
	return s.String()
}

func (s *DescribePairDrillsResponse) SetHeaders(v map[string]*string) *DescribePairDrillsResponse {
	s.Headers = v
	return s
}

func (s *DescribePairDrillsResponse) SetStatusCode(v int32) *DescribePairDrillsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePairDrillsResponse) SetBody(v *DescribePairDrillsResponseBody) *DescribePairDrillsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The language in which the regions and zones are named. This parameter corresponds to the `LocalName` response parameter. Valid values:
	//
	// *   zh-CN: Chinese
	// *   en-US: English
	// *   ja: Japanese
	//
	// Default value: zh-CN.
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of resource. Valid values:
	//
	// *   ear: async replication
	// *   lens: CloudLens for EBS
	// *   dbsc: Dedicated Block Storage Cluster
	//
	// Default value: ear.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceType(v string) *DescribeRegionsRequest {
	s.ResourceType = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// Details about the regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The name of the region.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Details about the zones.
	Zones []*DescribeRegionsResponseBodyRegionsZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetZones(v []*DescribeRegionsResponseBodyRegionsZones) *DescribeRegionsResponseBodyRegions {
	s.Zones = v
	return s
}

type DescribeRegionsResponseBodyRegionsZones struct {
	// The name of the zone.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The type of resource list.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsZones {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetResourceTypes(v []*string) *DescribeRegionsResponseBodyRegionsZones {
	s.ResourceTypes = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsZones {
	s.ZoneId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeReplicaGroupDrillsRequest struct {
	// The ID of the drill.
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](~~426614~~) operation to query a list of async replication pair-consistent groups, including group IDs.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum number of entries to be returned. You can use this parameter together with NextToken.
	//
	// Valid values: 1 to 500.
	//
	// Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken. When you specify NextToken, the PageSize and PageNumber request parameters do not take effect and the TotalCount response parameter is invalid.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the primary or secondary disk in the async replication pair-consistent group. You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which async replication is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeReplicaGroupDrillsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicaGroupDrillsRequest) GoString() string {
	return s.String()
}

func (s *DescribeReplicaGroupDrillsRequest) SetDrillId(v string) *DescribeReplicaGroupDrillsRequest {
	s.DrillId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) SetGroupId(v string) *DescribeReplicaGroupDrillsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) SetMaxResults(v int32) *DescribeReplicaGroupDrillsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) SetNextToken(v string) *DescribeReplicaGroupDrillsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) SetPageNumber(v int32) *DescribeReplicaGroupDrillsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) SetPageSize(v int32) *DescribeReplicaGroupDrillsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) SetRegionId(v string) *DescribeReplicaGroupDrillsRequest {
	s.RegionId = &v
	return s
}

type DescribeReplicaGroupDrillsResponseBody struct {
	// The information of disaster recovery drills that were performed on the replication pair-consistent group.
	Drills []*DescribeReplicaGroupDrillsResponseBodyDrills `json:"Drills,omitempty" xml:"Drills,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeReplicaGroupDrillsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicaGroupDrillsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReplicaGroupDrillsResponseBody) SetDrills(v []*DescribeReplicaGroupDrillsResponseBodyDrills) *DescribeReplicaGroupDrillsResponseBody {
	s.Drills = v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBody) SetNextToken(v string) *DescribeReplicaGroupDrillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBody) SetPageNumber(v int32) *DescribeReplicaGroupDrillsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBody) SetPageSize(v int32) *DescribeReplicaGroupDrillsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBody) SetRequestId(v string) *DescribeReplicaGroupDrillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBody) SetTotalCount(v int64) *DescribeReplicaGroupDrillsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeReplicaGroupDrillsResponseBodyDrills struct {
	// The ID of the drill.
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The ID of the replication pair-consistent group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The information of replication pairs.
	PairsInfo []*DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo `json:"PairsInfo,omitempty" xml:"PairsInfo,omitempty" type:"Repeated"`
	// The recovery point of the drill. The value of this parameter is a timestamp. Unit: seconds.
	RecoverPoint *int64 `json:"RecoverPoint,omitempty" xml:"RecoverPoint,omitempty"`
	// The beginning time of the drill. The value of this parameter is a timestamp. Unit: seconds.
	StartAt *int64 `json:"StartAt,omitempty" xml:"StartAt,omitempty"`
	// The status of the drill. Valid values:
	//
	// *   execute_failed
	// *   executed
	// *   executing
	// *   clear_failed
	// *   clearing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The error message that appears if the drill fails to be executed.
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DescribeReplicaGroupDrillsResponseBodyDrills) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicaGroupDrillsResponseBodyDrills) GoString() string {
	return s.String()
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetDrillId(v string) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.DrillId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetGroupId(v string) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.GroupId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetPairsInfo(v []*DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.PairsInfo = v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetRecoverPoint(v int64) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.RecoverPoint = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetStartAt(v int64) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.StartAt = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetStatus(v string) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.Status = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetStatusMessage(v string) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.StatusMessage = &v
	return s
}

type DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo struct {
	// The ID of the drill disk.
	DrillDiskId *string `json:"DrillDiskId,omitempty" xml:"DrillDiskId,omitempty"`
	// The status of the drill disk. Valid values:
	//
	// *   created
	// *   deleted
	// *   creating
	// *   deleting
	//
	// >  This parameter can also display error code details if your drill disk fails to be created or deleted.
	DrillDiskStatus *string `json:"DrillDiskStatus,omitempty" xml:"DrillDiskStatus,omitempty"`
	// The ID of the replication pair.
	PairId *string `json:"PairId,omitempty" xml:"PairId,omitempty"`
}

func (s DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) GoString() string {
	return s.String()
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) SetDrillDiskId(v string) *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo {
	s.DrillDiskId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) SetDrillDiskStatus(v string) *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo {
	s.DrillDiskStatus = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) SetPairId(v string) *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo {
	s.PairId = &v
	return s
}

type DescribeReplicaGroupDrillsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReplicaGroupDrillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReplicaGroupDrillsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicaGroupDrillsResponse) GoString() string {
	return s.String()
}

func (s *DescribeReplicaGroupDrillsResponse) SetHeaders(v map[string]*string) *DescribeReplicaGroupDrillsResponse {
	s.Headers = v
	return s
}

func (s *DescribeReplicaGroupDrillsResponse) SetStatusCode(v int32) *DescribeReplicaGroupDrillsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponse) SetBody(v *DescribeReplicaGroupDrillsResponseBody) *DescribeReplicaGroupDrillsResponse {
	s.Body = v
	return s
}

type DescribeSolutionInstanceConfigurationRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The parameters.
	Parameters []*DescribeSolutionInstanceConfigurationRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the solution.
	SolutionId *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s DescribeSolutionInstanceConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSolutionInstanceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeSolutionInstanceConfigurationRequest) SetClientToken(v string) *DescribeSolutionInstanceConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSolutionInstanceConfigurationRequest) SetParameters(v []*DescribeSolutionInstanceConfigurationRequestParameters) *DescribeSolutionInstanceConfigurationRequest {
	s.Parameters = v
	return s
}

func (s *DescribeSolutionInstanceConfigurationRequest) SetRegionId(v string) *DescribeSolutionInstanceConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSolutionInstanceConfigurationRequest) SetSolutionId(v string) *DescribeSolutionInstanceConfigurationRequest {
	s.SolutionId = &v
	return s
}

type DescribeSolutionInstanceConfigurationRequestParameters struct {
	// The key of the parameter.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeSolutionInstanceConfigurationRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeSolutionInstanceConfigurationRequestParameters) GoString() string {
	return s.String()
}

func (s *DescribeSolutionInstanceConfigurationRequestParameters) SetParameterKey(v string) *DescribeSolutionInstanceConfigurationRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *DescribeSolutionInstanceConfigurationRequestParameters) SetParameterValue(v string) *DescribeSolutionInstanceConfigurationRequestParameters {
	s.ParameterValue = &v
	return s
}

type DescribeSolutionInstanceConfigurationResponseBody struct {
	// The returned data.
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSolutionInstanceConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSolutionInstanceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSolutionInstanceConfigurationResponseBody) SetData(v []map[string]interface{}) *DescribeSolutionInstanceConfigurationResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSolutionInstanceConfigurationResponseBody) SetRequestId(v string) *DescribeSolutionInstanceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSolutionInstanceConfigurationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSolutionInstanceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSolutionInstanceConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSolutionInstanceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeSolutionInstanceConfigurationResponse) SetHeaders(v map[string]*string) *DescribeSolutionInstanceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeSolutionInstanceConfigurationResponse) SetStatusCode(v int32) *DescribeSolutionInstanceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSolutionInstanceConfigurationResponse) SetBody(v *DescribeSolutionInstanceConfigurationResponseBody) *DescribeSolutionInstanceConfigurationResponse {
	s.Body = v
	return s
}

type FailoverDiskReplicaGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the secondary site of the replication pair-consistent group.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group.
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s FailoverDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaGroupRequest) SetClientToken(v string) *FailoverDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *FailoverDiskReplicaGroupRequest) SetRegionId(v string) *FailoverDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *FailoverDiskReplicaGroupRequest) SetReplicaGroupId(v string) *FailoverDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

type FailoverDiskReplicaGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailoverDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaGroupResponseBody) SetRequestId(v string) *FailoverDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type FailoverDiskReplicaGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FailoverDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FailoverDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *FailoverDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *FailoverDiskReplicaGroupResponse) SetStatusCode(v int32) *FailoverDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *FailoverDiskReplicaGroupResponse) SetBody(v *FailoverDiskReplicaGroupResponseBody) *FailoverDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type FailoverDiskReplicaPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the secondary disk in the replication pair. You can call the [DescribeDiskReplicaPairs](~~354206~~) operation to query region IDs of secondary disks in replication pairs.
	//
	// >  The failover feature must be enabled for the region where the secondary disk is located.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair.
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s FailoverDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaPairRequest) SetClientToken(v string) *FailoverDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *FailoverDiskReplicaPairRequest) SetRegionId(v string) *FailoverDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *FailoverDiskReplicaPairRequest) SetReplicaPairId(v string) *FailoverDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type FailoverDiskReplicaPairResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailoverDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaPairResponseBody) SetRequestId(v string) *FailoverDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type FailoverDiskReplicaPairResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FailoverDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FailoverDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaPairResponse) SetHeaders(v map[string]*string) *FailoverDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *FailoverDiskReplicaPairResponse) SetStatusCode(v int32) *FailoverDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *FailoverDiskReplicaPairResponse) SetBody(v *FailoverDiskReplicaPairResponseBody) *FailoverDiskReplicaPairResponse {
	s.Body = v
	return s
}

type GetDiskRequest struct {
	// The ID of the disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of the disk.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDiskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiskRequest) GoString() string {
	return s.String()
}

func (s *GetDiskRequest) SetDiskId(v string) *GetDiskRequest {
	s.DiskId = &v
	return s
}

func (s *GetDiskRequest) SetRegionId(v string) *GetDiskRequest {
	s.RegionId = &v
	return s
}

type GetDiskResponseBody struct {
	// The information of the disk.
	Disk *GetDiskResponseBodyDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiskResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiskResponseBody) SetDisk(v *GetDiskResponseBodyDisk) *GetDiskResponseBody {
	s.Disk = v
	return s
}

func (s *GetDiskResponseBody) SetRequestId(v string) *GetDiskResponseBody {
	s.RequestId = &v
	return s
}

type GetDiskResponseBodyDisk struct {
	// The access authorization information of the disk.
	AccessAuthorization []*GetDiskResponseBodyDiskAccessAuthorization `json:"AccessAuthorization,omitempty" xml:"AccessAuthorization,omitempty" type:"Repeated"`
	// The time when the disk was attached.
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	// The attachment information of the disk.
	Attachments []*GetDiskResponseBodyDiskAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	// Specifies whether to enable the performance burst feature for the disk. Valid values:
	//
	// * true: enable the performance burst feature.
	// * false: disable the performance burst feature.
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of the disk.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the disk was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC. For more information, see [ISO 8601](~~25696~~).
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the disk was detached.
	DetachedTime *string `json:"DetachedTime,omitempty" xml:"DetachedTime,omitempty"`
	// The maximum number of read and write operations per second. Unit: operations/s.
	DeviceNguid *string `json:"DeviceNguid,omitempty" xml:"DeviceNguid,omitempty"`
	// The ID of the disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The name of the disk.
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the disk. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false.
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// IOPS.
	Iops *int64 `json:"Iops,omitempty" xml:"Iops,omitempty"`
	// The maximum number of read operations per second. Unit: operations/s.
	IopsRead *int64 `json:"IopsRead,omitempty" xml:"IopsRead,omitempty"`
	// The maximum number of write operations per second. Unit: operations/s.
	IopsWrite *int64 `json:"IopsWrite,omitempty" xml:"IopsWrite,omitempty"`
	// The ID of the KMS key.
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The time when the disk was modfied.
	ModifyAt *string `json:"ModifyAt,omitempty" xml:"ModifyAt,omitempty"`
	// The performance level of the ESSD. Valid values:
	//
	// * PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	// * PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// * PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// * PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}
	//
	// Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}
	//
	// >  This parameter is available only if the DiskCategory parameter is set to cloud_auto. For more information, see [ESSD AutoPL disks](~~368372~~) and [Modify the performance configurations of an ESSD AutoPL disk](~~413275~~).
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The region ID of the disk. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the disk belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The size of the disk. Unit: GB.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The states of disk. Valid values:
	//
	// * In_use
	// * Available
	// * Attaching
	// * Detaching
	// * Creating
	// * ReIniting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The details of the tags.
	Tags []*GetDiskResponseBodyDiskTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The throughput of the disk.
	//
	// Unit: MB/s.
	Throughput *int64 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	// The ID of the zone where the disk resides.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetDiskResponseBodyDisk) String() string {
	return tea.Prettify(s)
}

func (s GetDiskResponseBodyDisk) GoString() string {
	return s.String()
}

func (s *GetDiskResponseBodyDisk) SetAccessAuthorization(v []*GetDiskResponseBodyDiskAccessAuthorization) *GetDiskResponseBodyDisk {
	s.AccessAuthorization = v
	return s
}

func (s *GetDiskResponseBodyDisk) SetAttachedTime(v string) *GetDiskResponseBodyDisk {
	s.AttachedTime = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetAttachments(v []*GetDiskResponseBodyDiskAttachments) *GetDiskResponseBodyDisk {
	s.Attachments = v
	return s
}

func (s *GetDiskResponseBodyDisk) SetBurstingEnabled(v bool) *GetDiskResponseBodyDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetCategory(v string) *GetDiskResponseBodyDisk {
	s.Category = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetCreationTime(v string) *GetDiskResponseBodyDisk {
	s.CreationTime = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetDescription(v string) *GetDiskResponseBodyDisk {
	s.Description = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetDetachedTime(v string) *GetDiskResponseBodyDisk {
	s.DetachedTime = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetDeviceNguid(v string) *GetDiskResponseBodyDisk {
	s.DeviceNguid = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetDiskId(v string) *GetDiskResponseBodyDisk {
	s.DiskId = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetDiskName(v string) *GetDiskResponseBodyDisk {
	s.DiskName = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetEncrypted(v bool) *GetDiskResponseBodyDisk {
	s.Encrypted = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetIops(v int64) *GetDiskResponseBodyDisk {
	s.Iops = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetIopsRead(v int64) *GetDiskResponseBodyDisk {
	s.IopsRead = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetIopsWrite(v int64) *GetDiskResponseBodyDisk {
	s.IopsWrite = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetKmsKeyId(v string) *GetDiskResponseBodyDisk {
	s.KmsKeyId = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetModifyAt(v string) *GetDiskResponseBodyDisk {
	s.ModifyAt = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetPerformanceLevel(v string) *GetDiskResponseBodyDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetProvisionedIops(v int64) *GetDiskResponseBodyDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetRegionId(v string) *GetDiskResponseBodyDisk {
	s.RegionId = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetResourceGroupId(v string) *GetDiskResponseBodyDisk {
	s.ResourceGroupId = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetSize(v int64) *GetDiskResponseBodyDisk {
	s.Size = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetStatus(v string) *GetDiskResponseBodyDisk {
	s.Status = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetTags(v []*GetDiskResponseBodyDiskTags) *GetDiskResponseBodyDisk {
	s.Tags = v
	return s
}

func (s *GetDiskResponseBodyDisk) SetThroughput(v int64) *GetDiskResponseBodyDisk {
	s.Throughput = &v
	return s
}

func (s *GetDiskResponseBodyDisk) SetZoneId(v string) *GetDiskResponseBodyDisk {
	s.ZoneId = &v
	return s
}

type GetDiskResponseBodyDiskAccessAuthorization struct {
	// The token that is used to access the disk.
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The information of ecs.
	Ecs *GetDiskResponseBodyDiskAccessAuthorizationEcs `json:"Ecs,omitempty" xml:"Ecs,omitempty" type:"Struct"`
	// The information of pod.
	Pod *GetDiskResponseBodyDiskAccessAuthorizationPod `json:"Pod,omitempty" xml:"Pod,omitempty" type:"Struct"`
}

func (s GetDiskResponseBodyDiskAccessAuthorization) String() string {
	return tea.Prettify(s)
}

func (s GetDiskResponseBodyDiskAccessAuthorization) GoString() string {
	return s.String()
}

func (s *GetDiskResponseBodyDiskAccessAuthorization) SetAccessToken(v string) *GetDiskResponseBodyDiskAccessAuthorization {
	s.AccessToken = &v
	return s
}

func (s *GetDiskResponseBodyDiskAccessAuthorization) SetEcs(v *GetDiskResponseBodyDiskAccessAuthorizationEcs) *GetDiskResponseBodyDiskAccessAuthorization {
	s.Ecs = v
	return s
}

func (s *GetDiskResponseBodyDiskAccessAuthorization) SetPod(v *GetDiskResponseBodyDiskAccessAuthorizationPod) *GetDiskResponseBodyDiskAccessAuthorization {
	s.Pod = v
	return s
}

type GetDiskResponseBodyDiskAccessAuthorizationEcs struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDiskResponseBodyDiskAccessAuthorizationEcs) String() string {
	return tea.Prettify(s)
}

func (s GetDiskResponseBodyDiskAccessAuthorizationEcs) GoString() string {
	return s.String()
}

func (s *GetDiskResponseBodyDiskAccessAuthorizationEcs) SetInstanceId(v string) *GetDiskResponseBodyDiskAccessAuthorizationEcs {
	s.InstanceId = &v
	return s
}

type GetDiskResponseBodyDiskAccessAuthorizationPod struct {
	// The cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The pod ID.
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
}

func (s GetDiskResponseBodyDiskAccessAuthorizationPod) String() string {
	return tea.Prettify(s)
}

func (s GetDiskResponseBodyDiskAccessAuthorizationPod) GoString() string {
	return s.String()
}

func (s *GetDiskResponseBodyDiskAccessAuthorizationPod) SetClusterId(v string) *GetDiskResponseBodyDiskAccessAuthorizationPod {
	s.ClusterId = &v
	return s
}

func (s *GetDiskResponseBodyDiskAccessAuthorizationPod) SetPodId(v string) *GetDiskResponseBodyDiskAccessAuthorizationPod {
	s.PodId = &v
	return s
}

type GetDiskResponseBodyDiskAttachments struct {
	// The time when the disk was attached.
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	NamespaceId *int32 `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s GetDiskResponseBodyDiskAttachments) String() string {
	return tea.Prettify(s)
}

func (s GetDiskResponseBodyDiskAttachments) GoString() string {
	return s.String()
}

func (s *GetDiskResponseBodyDiskAttachments) SetAttachedTime(v string) *GetDiskResponseBodyDiskAttachments {
	s.AttachedTime = &v
	return s
}

func (s *GetDiskResponseBodyDiskAttachments) SetInstanceId(v string) *GetDiskResponseBodyDiskAttachments {
	s.InstanceId = &v
	return s
}

func (s *GetDiskResponseBodyDiskAttachments) SetNamespaceId(v int32) *GetDiskResponseBodyDiskAttachments {
	s.NamespaceId = &v
	return s
}

type GetDiskResponseBodyDiskTags struct {
	// The tag key.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetDiskResponseBodyDiskTags) String() string {
	return tea.Prettify(s)
}

func (s GetDiskResponseBodyDiskTags) GoString() string {
	return s.String()
}

func (s *GetDiskResponseBodyDiskTags) SetTagKey(v string) *GetDiskResponseBodyDiskTags {
	s.TagKey = &v
	return s
}

func (s *GetDiskResponseBodyDiskTags) SetTagValue(v string) *GetDiskResponseBodyDiskTags {
	s.TagValue = &v
	return s
}

type GetDiskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDiskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiskResponse) GoString() string {
	return s.String()
}

func (s *GetDiskResponse) SetHeaders(v map[string]*string) *GetDiskResponse {
	s.Headers = v
	return s
}

func (s *GetDiskResponse) SetStatusCode(v int32) *GetDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiskResponse) SetBody(v *GetDiskResponseBody) *GetDiskResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The **ClientToken** value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the resource. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID list of the resource. You can specify up to 50 resource IDs in each call.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Valid values:
	//
	// *   dedicatedblockstoragecluster: dedicated block storage cluster
	// *   diskreplicapair: replication pair
	// *   diskreplicagroup: replication pair-consistent group
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The information about the tags.
	//
	// You can specify at most 20 tags in each call.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetClientToken(v string) *ListTagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of tag N used for exact search of EBS resources. The tag key must be 1 to 128 characters in length. Valid values of N: 1 to 20.
	//
	// The `Tag.N` parameter pair (Tag.N.Key and Tag.N.Value) is used for exact search of EBS resources that have specified tags added. Each tag is a key-value pair.
	//
	// *   If you specify only `Tag.N.Key`, all EBS resources whose tags contain the specified tag key are returned.
	// *   If you specify only `Tag.N.Value`, the `InvalidParameter.TagValue` error is returned.
	// *   If you specify multiple tag key-value pairs at the same time, only EBS resources that match all tag key-value pairs are returned.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N used for exact search of EBS resources. The tag value must be 1 to 128 characters in length. Valid values of N: 1 to 20.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// The token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request. The request ID is returned regardless of whether the call is successful.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the resources and tags, including resource IDs, resource types, and tag key-value pairs.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// *   dedicatedblockstoragecluster: dedicated block storage cluster
	// *   diskreplicapair: replication pair
	// *   diskreplicagroup: replication pair-consistent group
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key of the resource.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the resource.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyDedicatedBlockStorageClusterAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure idempotence ](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the dedicated block storage cluster.
	DbscId *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	// The new name of the dedicated block storage cluster.
	DbscName *string `json:"DbscName,omitempty" xml:"DbscName,omitempty"`
	// The new description of dedicated block storage cluster.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID of the dedicated block storage cluster. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDedicatedBlockStorageClusterAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedBlockStorageClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetClientToken(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetDbscId(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.DbscId = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetDbscName(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.DbscName = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetDescription(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetRegionId(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.RegionId = &v
	return s
}

type ModifyDedicatedBlockStorageClusterAttributeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDedicatedBlockStorageClusterAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedBlockStorageClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponseBody) SetRequestId(v string) *ModifyDedicatedBlockStorageClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDedicatedBlockStorageClusterAttributeResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDedicatedBlockStorageClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDedicatedBlockStorageClusterAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedBlockStorageClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) SetHeaders(v map[string]*string) *ModifyDedicatedBlockStorageClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) SetStatusCode(v int32) *ModifyDedicatedBlockStorageClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) SetBody(v *ModifyDedicatedBlockStorageClusterAttributeResponseBody) *ModifyDedicatedBlockStorageClusterAttributeResponse {
	s.Body = v
	return s
}

type ModifyDiskReplicaGroupRequest struct {
	// The bandwidth value. Unit: Kbit/s.
	//
	// >  This parameter is not publicly available.
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the replication pair-consistent group. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the replication pair-consistent group. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (\_), and hyphens (-).
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The RPO of the replication pair-consistent group. Unit: seconds. Valid value: 900.
	RPO *int64 `json:"RPO,omitempty" xml:"RPO,omitempty"`
	// The region ID of the replication pair-consistent group.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](~~426614~~) operation to query the IDs of replication pair-consistent groups.
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s ModifyDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaGroupRequest) SetBandwidth(v int64) *ModifyDiskReplicaGroupRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetClientToken(v string) *ModifyDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetDescription(v string) *ModifyDiskReplicaGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetGroupName(v string) *ModifyDiskReplicaGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetRPO(v int64) *ModifyDiskReplicaGroupRequest {
	s.RPO = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetRegionId(v string) *ModifyDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetReplicaGroupId(v string) *ModifyDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

type ModifyDiskReplicaGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaGroupResponseBody) SetRequestId(v string) *ModifyDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDiskReplicaGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *ModifyDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskReplicaGroupResponse) SetStatusCode(v int32) *ModifyDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskReplicaGroupResponse) SetBody(v *ModifyDiskReplicaGroupResponseBody) *ModifyDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type ModifyDiskReplicaPairRequest struct {
	// The bandwidth value. Unit: Kbit/s.
	//
	// >  This parameter is not publicly available.
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the replication pair.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the replication pair.
	PairName *string `json:"PairName,omitempty" xml:"PairName,omitempty"`
	// The recovery point objective (RPO) of the replication pair-consistent group. Unit: seconds. Valid value: 900.
	RPO *int64 `json:"RPO,omitempty" xml:"RPO,omitempty"`
	// The region ID of the primary or secondary disk in the replication pair. You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which async replication is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair.
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s ModifyDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaPairRequest) SetBandwidth(v int64) *ModifyDiskReplicaPairRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetClientToken(v string) *ModifyDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetDescription(v string) *ModifyDiskReplicaPairRequest {
	s.Description = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetPairName(v string) *ModifyDiskReplicaPairRequest {
	s.PairName = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetRPO(v int64) *ModifyDiskReplicaPairRequest {
	s.RPO = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetRegionId(v string) *ModifyDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetReplicaPairId(v string) *ModifyDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type ModifyDiskReplicaPairResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaPairResponseBody) SetRequestId(v string) *ModifyDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDiskReplicaPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaPairResponse) SetHeaders(v map[string]*string) *ModifyDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskReplicaPairResponse) SetStatusCode(v int32) *ModifyDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskReplicaPairResponse) SetBody(v *ModifyDiskReplicaPairResponseBody) *ModifyDiskReplicaPairResponse {
	s.Body = v
	return s
}

type QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the request of SetDedicatedBlockStorageClusterDiskThroughput api.
	QosRequestId *string `json:"QosRequestId,omitempty" xml:"QosRequestId,omitempty"`
	// The region ID of the dedicated block storage cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) SetClientToken(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) SetQosRequestId(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest {
	s.QosRequestId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) SetRegionId(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest {
	s.RegionId = &v
	return s
}

type QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the throughput after setting the throughput by SetDedicatedBlockStorageClusterDiskThroughput api.
	//
	// - SUCCESS: The throughput has been successfully set.
	// - RUNNING: The throughput is currently being set.
	// - WAIT(): The throughput is waiting to be set.
	// - FAIL(): The throughput setting has failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) SetRequestId(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) SetStatus(v string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody {
	s.Status = &v
	return s
}

type QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse struct {
	Headers    map[string]*string                                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) SetHeaders(v map[string]*string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) SetStatusCode(v int32) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) SetBody(v *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse {
	s.Body = v
	return s
}

type QueryDedicatedBlockStorageClusterInventoryDataRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure idempotence ](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the dedicated block storage cluster.
	DbscId *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	// End timestamp of trend data.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval (seconds) between data retrieval points.
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the dedicated block storage cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Start timestamp of trend data.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterInventoryDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterInventoryDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) SetClientToken(v string) *QueryDedicatedBlockStorageClusterInventoryDataRequest {
	s.ClientToken = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) SetDbscId(v string) *QueryDedicatedBlockStorageClusterInventoryDataRequest {
	s.DbscId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) SetEndTime(v int64) *QueryDedicatedBlockStorageClusterInventoryDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) SetPeriod(v int32) *QueryDedicatedBlockStorageClusterInventoryDataRequest {
	s.Period = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) SetRegionId(v string) *QueryDedicatedBlockStorageClusterInventoryDataRequest {
	s.RegionId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataRequest) SetStartTime(v int64) *QueryDedicatedBlockStorageClusterInventoryDataRequest {
	s.StartTime = &v
	return s
}

type QueryDedicatedBlockStorageClusterInventoryDataResponseBody struct {
	// The returned data.
	Data []*QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the dedicated block storage cluster.
	DbscId *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	// The name of the dedicated block storage cluster.
	DbscName *string `json:"DbscName,omitempty" xml:"DbscName,omitempty"`
	// The type of the disk. Valid values:
	//
	// *   cloud_essd: enhanced SSD (ESSD).
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) SetData(v []*QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	s.Data = v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) SetDbscId(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	s.DbscId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) SetDbscName(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	s.DbscName = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) SetDiskCategory(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	s.DiskCategory = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) SetRequestId(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) SetTotalCount(v int32) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	s.TotalCount = &v
	return s
}

type QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData struct {
	// The returned metrics.
	MonitorItems *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems `json:"MonitorItems,omitempty" xml:"MonitorItems,omitempty" type:"Struct"`
	// The ID list of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The timestamp when the data is collected.
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) SetMonitorItems(v *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData {
	s.MonitorItems = v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) SetResourceId(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) SetTimestamp(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData {
	s.Timestamp = &v
	return s
}

type QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems struct {
	// Available capacity size of the dedicated block storage cluster.
	AvailableSize *int64 `json:"AvailableSize,omitempty" xml:"AvailableSize,omitempty"`
	// Total capacity size of the dedicated block storage cluster.
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) String() string {
	return tea.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) SetAvailableSize(v int64) *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems {
	s.AvailableSize = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) SetTotalSize(v int64) *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems {
	s.TotalSize = &v
	return s
}

type QueryDedicatedBlockStorageClusterInventoryDataResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDedicatedBlockStorageClusterInventoryDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponse) SetHeaders(v map[string]*string) *QueryDedicatedBlockStorageClusterInventoryDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponse) SetStatusCode(v int32) *QueryDedicatedBlockStorageClusterInventoryDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponse) SetBody(v *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) *QueryDedicatedBlockStorageClusterInventoryDataResponse {
	s.Body = v
	return s
}

type RemoveDiskReplicaPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the replication pair-consistent group.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group.
	//
	// You can call the [DescribeDiskReplicaGroups](~~426614~~) operation to query the IDs of replication pair-consistent groups.
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// The ID of the replication pair.
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s RemoveDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *RemoveDiskReplicaPairRequest) SetClientToken(v string) *RemoveDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveDiskReplicaPairRequest) SetRegionId(v string) *RemoveDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveDiskReplicaPairRequest) SetReplicaGroupId(v string) *RemoveDiskReplicaPairRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *RemoveDiskReplicaPairRequest) SetReplicaPairId(v string) *RemoveDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type RemoveDiskReplicaPairResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDiskReplicaPairResponseBody) SetRequestId(v string) *RemoveDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type RemoveDiskReplicaPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *RemoveDiskReplicaPairResponse) SetHeaders(v map[string]*string) *RemoveDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *RemoveDiskReplicaPairResponse) SetStatusCode(v int32) *RemoveDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDiskReplicaPairResponse) SetBody(v *RemoveDiskReplicaPairResponseBody) *RemoveDiskReplicaPairResponse {
	s.Body = v
	return s
}

type ReprotectDiskReplicaGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](~~426614~~) operation to query the IDs of replication pair-consistent groups.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](~~426614~~) operation to query the IDs of replication pair-consistent groups.
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// Specifies whether to enable the reverse replication sub-feature. Valid values: true and false. Default value: true.
	ReverseReplicate *bool `json:"ReverseReplicate,omitempty" xml:"ReverseReplicate,omitempty"`
}

func (s ReprotectDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaGroupRequest) SetClientToken(v string) *ReprotectDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ReprotectDiskReplicaGroupRequest) SetRegionId(v string) *ReprotectDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ReprotectDiskReplicaGroupRequest) SetReplicaGroupId(v string) *ReprotectDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *ReprotectDiskReplicaGroupRequest) SetReverseReplicate(v bool) *ReprotectDiskReplicaGroupRequest {
	s.ReverseReplicate = &v
	return s
}

type ReprotectDiskReplicaGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReprotectDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaGroupResponseBody) SetRequestId(v string) *ReprotectDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type ReprotectDiskReplicaGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReprotectDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReprotectDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *ReprotectDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *ReprotectDiskReplicaGroupResponse) SetStatusCode(v int32) *ReprotectDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ReprotectDiskReplicaGroupResponse) SetBody(v *ReprotectDiskReplicaGroupResponseBody) *ReprotectDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type ReprotectDiskReplicaPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the secondary disk in the replication pair. You can call the [DescribeDiskReplicaPairs](~~354206~~) operation to query region IDs of secondary disks in replication pairs.
	//
	// >  The reverse replication feature must be enabled from the region where the secondary disk is located.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair.
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
	// Specifies whether to enable the reverse replication sub-feature. Valid values: true and false. Default value: true.
	ReverseReplicate *bool `json:"ReverseReplicate,omitempty" xml:"ReverseReplicate,omitempty"`
}

func (s ReprotectDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaPairRequest) SetClientToken(v string) *ReprotectDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *ReprotectDiskReplicaPairRequest) SetRegionId(v string) *ReprotectDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *ReprotectDiskReplicaPairRequest) SetReplicaPairId(v string) *ReprotectDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

func (s *ReprotectDiskReplicaPairRequest) SetReverseReplicate(v bool) *ReprotectDiskReplicaPairRequest {
	s.ReverseReplicate = &v
	return s
}

type ReprotectDiskReplicaPairResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReprotectDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaPairResponseBody) SetRequestId(v string) *ReprotectDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type ReprotectDiskReplicaPairResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReprotectDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReprotectDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaPairResponse) SetHeaders(v map[string]*string) *ReprotectDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *ReprotectDiskReplicaPairResponse) SetStatusCode(v int32) *ReprotectDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *ReprotectDiskReplicaPairResponse) SetBody(v *ReprotectDiskReplicaPairResponseBody) *ReprotectDiskReplicaPairResponse {
	s.Body = v
	return s
}

type SetDedicatedBlockStorageClusterDiskThroughputRequest struct {
	// Target throughput.
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of disk.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetDedicatedBlockStorageClusterDiskThroughputRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDedicatedBlockStorageClusterDiskThroughputRequest) GoString() string {
	return s.String()
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) SetBps(v int32) *SetDedicatedBlockStorageClusterDiskThroughputRequest {
	s.Bps = &v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) SetClientToken(v string) *SetDedicatedBlockStorageClusterDiskThroughputRequest {
	s.ClientToken = &v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) SetDiskId(v string) *SetDedicatedBlockStorageClusterDiskThroughputRequest {
	s.DiskId = &v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) SetRegionId(v string) *SetDedicatedBlockStorageClusterDiskThroughputRequest {
	s.RegionId = &v
	return s
}

type SetDedicatedBlockStorageClusterDiskThroughputResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDedicatedBlockStorageClusterDiskThroughputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDedicatedBlockStorageClusterDiskThroughputResponseBody) GoString() string {
	return s.String()
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponseBody) SetRequestId(v string) *SetDedicatedBlockStorageClusterDiskThroughputResponseBody {
	s.RequestId = &v
	return s
}

type SetDedicatedBlockStorageClusterDiskThroughputResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDedicatedBlockStorageClusterDiskThroughputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDedicatedBlockStorageClusterDiskThroughputResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDedicatedBlockStorageClusterDiskThroughputResponse) GoString() string {
	return s.String()
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponse) SetHeaders(v map[string]*string) *SetDedicatedBlockStorageClusterDiskThroughputResponse {
	s.Headers = v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponse) SetStatusCode(v int32) *SetDedicatedBlockStorageClusterDiskThroughputResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponse) SetBody(v *SetDedicatedBlockStorageClusterDiskThroughputResponseBody) *SetDedicatedBlockStorageClusterDiskThroughputResponse {
	s.Body = v
	return s
}

type StartDiskReplicaGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to immediately synchronize data once. Valid values:
	//
	// *   true: immediately synchronizes data once.
	// *   false: synchronizes data based on the RPO of the replication pair-consistent group.
	//
	// Default value: false.
	OneShot *bool `json:"OneShot,omitempty" xml:"OneShot,omitempty"`
	// The ID of the replication pair-consistent group.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](~~426614~~) operation to query the IDs of replication pair-consistent groups.
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s StartDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaGroupRequest) SetClientToken(v string) *StartDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDiskReplicaGroupRequest) SetOneShot(v bool) *StartDiskReplicaGroupRequest {
	s.OneShot = &v
	return s
}

func (s *StartDiskReplicaGroupRequest) SetRegionId(v string) *StartDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *StartDiskReplicaGroupRequest) SetReplicaGroupId(v string) *StartDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

type StartDiskReplicaGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaGroupResponseBody) SetRequestId(v string) *StartDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type StartDiskReplicaGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *StartDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *StartDiskReplicaGroupResponse) SetStatusCode(v int32) *StartDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDiskReplicaGroupResponse) SetBody(v *StartDiskReplicaGroupResponseBody) *StartDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type StartDiskReplicaPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to immediately synchronize data. Valid values:
	//
	// *   true: immediately synchronizes data.
	// *   false: synchronizes data based on the recovery point objective (RPO).
	//
	// Default value: false.
	OneShot *bool `json:"OneShot,omitempty" xml:"OneShot,omitempty"`
	// The region ID of the primary or secondary disk in the replication pair. You can call the [DescribeDiskReplicaPairs](~~354206~~) operation to query the region information of replication pairs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair.
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s StartDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaPairRequest) SetClientToken(v string) *StartDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDiskReplicaPairRequest) SetOneShot(v bool) *StartDiskReplicaPairRequest {
	s.OneShot = &v
	return s
}

func (s *StartDiskReplicaPairRequest) SetRegionId(v string) *StartDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *StartDiskReplicaPairRequest) SetReplicaPairId(v string) *StartDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type StartDiskReplicaPairResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaPairResponseBody) SetRequestId(v string) *StartDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type StartDiskReplicaPairResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaPairResponse) SetHeaders(v map[string]*string) *StartDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *StartDiskReplicaPairResponse) SetStatusCode(v int32) *StartDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDiskReplicaPairResponse) SetBody(v *StartDiskReplicaPairResponseBody) *StartDiskReplicaPairResponse {
	s.Body = v
	return s
}

type StartPairDrillRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the replication pair. You can call the [DescribeDiskReplicaPairs](~~354206~~) operation to query a list of replication pairs, including replication pair IDs.
	PairId *string `json:"PairId,omitempty" xml:"PairId,omitempty"`
	// The region ID of the secondary disk in the replication pair. You can call the [DescribeDiskReplicaPairs](~~354206~~) operation to query the region in which the secondary disk of the replication pair resides.
	//
	// >  You must enable the disaster recovery drill feature in the region in which the secondary site resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartPairDrillRequest) String() string {
	return tea.Prettify(s)
}

func (s StartPairDrillRequest) GoString() string {
	return s.String()
}

func (s *StartPairDrillRequest) SetClientToken(v string) *StartPairDrillRequest {
	s.ClientToken = &v
	return s
}

func (s *StartPairDrillRequest) SetPairId(v string) *StartPairDrillRequest {
	s.PairId = &v
	return s
}

func (s *StartPairDrillRequest) SetRegionId(v string) *StartPairDrillRequest {
	s.RegionId = &v
	return s
}

type StartPairDrillResponseBody struct {
	// The drill ID.
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartPairDrillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartPairDrillResponseBody) GoString() string {
	return s.String()
}

func (s *StartPairDrillResponseBody) SetDrillId(v string) *StartPairDrillResponseBody {
	s.DrillId = &v
	return s
}

func (s *StartPairDrillResponseBody) SetRequestId(v string) *StartPairDrillResponseBody {
	s.RequestId = &v
	return s
}

type StartPairDrillResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartPairDrillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPairDrillResponse) String() string {
	return tea.Prettify(s)
}

func (s StartPairDrillResponse) GoString() string {
	return s.String()
}

func (s *StartPairDrillResponse) SetHeaders(v map[string]*string) *StartPairDrillResponse {
	s.Headers = v
	return s
}

func (s *StartPairDrillResponse) SetStatusCode(v int32) *StartPairDrillResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPairDrillResponse) SetBody(v *StartPairDrillResponseBody) *StartPairDrillResponse {
	s.Body = v
	return s
}

type StartReplicaGroupDrillRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the replication pair-consistent group ID. You can call the [DescribeDiskReplicaGroups](~~426614~~) operation the most recent list of async replication pair-consistent groups, including group IDs.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the region where the secondary site in the replication pair-consistent group is located. You can call the [DescribeDiskReplicaGroups](~~426614~~) operation to query the region where the secondary site in the replication pair-consistent group is located.
	//
	// >  You must enable the disaster recovery drill feature in the region in which the secondary site resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartReplicaGroupDrillRequest) String() string {
	return tea.Prettify(s)
}

func (s StartReplicaGroupDrillRequest) GoString() string {
	return s.String()
}

func (s *StartReplicaGroupDrillRequest) SetClientToken(v string) *StartReplicaGroupDrillRequest {
	s.ClientToken = &v
	return s
}

func (s *StartReplicaGroupDrillRequest) SetGroupId(v string) *StartReplicaGroupDrillRequest {
	s.GroupId = &v
	return s
}

func (s *StartReplicaGroupDrillRequest) SetRegionId(v string) *StartReplicaGroupDrillRequest {
	s.RegionId = &v
	return s
}

type StartReplicaGroupDrillResponseBody struct {
	// The drill ID.
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartReplicaGroupDrillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartReplicaGroupDrillResponseBody) GoString() string {
	return s.String()
}

func (s *StartReplicaGroupDrillResponseBody) SetDrillId(v string) *StartReplicaGroupDrillResponseBody {
	s.DrillId = &v
	return s
}

func (s *StartReplicaGroupDrillResponseBody) SetRequestId(v string) *StartReplicaGroupDrillResponseBody {
	s.RequestId = &v
	return s
}

type StartReplicaGroupDrillResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartReplicaGroupDrillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartReplicaGroupDrillResponse) String() string {
	return tea.Prettify(s)
}

func (s StartReplicaGroupDrillResponse) GoString() string {
	return s.String()
}

func (s *StartReplicaGroupDrillResponse) SetHeaders(v map[string]*string) *StartReplicaGroupDrillResponse {
	s.Headers = v
	return s
}

func (s *StartReplicaGroupDrillResponse) SetStatusCode(v int32) *StartReplicaGroupDrillResponse {
	s.StatusCode = &v
	return s
}

func (s *StartReplicaGroupDrillResponse) SetBody(v *StartReplicaGroupDrillResponseBody) *StartReplicaGroupDrillResponse {
	s.Body = v
	return s
}

type StopDiskReplicaGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the replication pair-consistent group.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](~~426614~~) operation to query the IDs of replication pair-consistent groups.
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s StopDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaGroupRequest) SetClientToken(v string) *StopDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDiskReplicaGroupRequest) SetRegionId(v string) *StopDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *StopDiskReplicaGroupRequest) SetReplicaGroupId(v string) *StopDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

type StopDiskReplicaGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaGroupResponseBody) SetRequestId(v string) *StopDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type StopDiskReplicaGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *StopDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *StopDiskReplicaGroupResponse) SetStatusCode(v int32) *StopDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDiskReplicaGroupResponse) SetBody(v *StopDiskReplicaGroupResponseBody) *StopDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type StopDiskReplicaPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the primary or secondary disk in the replication pair. You can call the [DescribeDiskReplicaPairs](~~354206~~) operation to query the region information of replication pairs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair.
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s StopDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaPairRequest) SetClientToken(v string) *StopDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDiskReplicaPairRequest) SetRegionId(v string) *StopDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *StopDiskReplicaPairRequest) SetReplicaPairId(v string) *StopDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type StopDiskReplicaPairResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaPairResponseBody) SetRequestId(v string) *StopDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type StopDiskReplicaPairResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaPairResponse) SetHeaders(v map[string]*string) *StopDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *StopDiskReplicaPairResponse) SetStatusCode(v int32) *StopDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDiskReplicaPairResponse) SetBody(v *StopDiskReplicaPairResponseBody) *StopDiskReplicaPairResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The **ClientToken** value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the resource. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID list of the resources. You can specify up to 50 IDs in each request.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Valid values:
	//
	// *   dedicatedblockstoragecluster: dedicated block storage cluster
	// *   diskreplicapair: replication pair
	// *   diskreplicagroup: replication pair-consistent group
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource tags. You can specify up to 20 tags.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetClientToken(v string) *TagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of tag N to add to the resource. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:` or contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The ID of the request. The request ID is returned regardless of whether the call is successful.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnbindEnterpriseSnapshotPolicyRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of disks.
	DiskTargets []*string `json:"DiskTargets,omitempty" xml:"DiskTargets,omitempty" type:"Repeated"`
	// The id of the policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID . You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which snapshot policy is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UnbindEnterpriseSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindEnterpriseSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) SetClientToken(v string) *UnbindEnterpriseSnapshotPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) SetDiskTargets(v []*string) *UnbindEnterpriseSnapshotPolicyRequest {
	s.DiskTargets = v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) SetPolicyId(v string) *UnbindEnterpriseSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) SetRegionId(v string) *UnbindEnterpriseSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

type UnbindEnterpriseSnapshotPolicyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindEnterpriseSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindEnterpriseSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindEnterpriseSnapshotPolicyResponseBody) SetRequestId(v string) *UnbindEnterpriseSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type UnbindEnterpriseSnapshotPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindEnterpriseSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindEnterpriseSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindEnterpriseSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *UnbindEnterpriseSnapshotPolicyResponse) SetHeaders(v map[string]*string) *UnbindEnterpriseSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyResponse) SetStatusCode(v int32) *UnbindEnterpriseSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyResponse) SetBody(v *UnbindEnterpriseSnapshotPolicyResponseBody) *UnbindEnterpriseSnapshotPolicyResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the resource. This parameter is valid only when the TagKey.N parameter is not specified. Valid values:
	//
	// *   true: removes all tags from the resource.
	// *   false: does not remove all tags from the resource.
	//
	// Default value: false.
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The **ClientToken** value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the resource. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID list of the resource. You can specify up to 50 resource IDs in each call.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Valid values:
	//
	// *   dedicatedblockstoragecluster: dedicated block storage cluster
	// *   diskreplicapair: the replication pair.
	// *   diskreplicagroup: replication pair-consistent group
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tag keys. You can specify up to 20 tag keys in the list.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetClientToken(v string) *UntagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The ID of the request. The request ID is returned regardless of whether the call is successful.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateEnterpriseSnapshotPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Snapshot replication destination information.
	CrossRegionCopyInfo *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo `json:"CrossRegionCopyInfo,omitempty" xml:"CrossRegionCopyInfo,omitempty" type:"Struct"`
	// The description of the policy.
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The name of the policy.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The id of the policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID . You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which snapshot policy is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Snapshot retention rule.
	RetainRule *UpdateEnterpriseSnapshotPolicyRequestRetainRule `json:"RetainRule,omitempty" xml:"RetainRule,omitempty" type:"Struct"`
	// The rule for scheduling.
	Schedule *UpdateEnterpriseSnapshotPolicyRequestSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Struct"`
	// The special snapshot retention rules.
	SpecialRetainRules *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules `json:"SpecialRetainRules,omitempty" xml:"SpecialRetainRules,omitempty" type:"Struct"`
	// The status of the policy. Valid values:
	//
	// *   **ENABLED**: Enable snapshot policy execution.
	// *   **DISABLED**: Disable snapshot policy execution.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Advanced snapshot features.
	StorageRule *UpdateEnterpriseSnapshotPolicyRequestStorageRule `json:"StorageRule,omitempty" xml:"StorageRule,omitempty" type:"Struct"`
}

func (s UpdateEnterpriseSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetClientToken(v string) *UpdateEnterpriseSnapshotPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetCrossRegionCopyInfo(v *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) *UpdateEnterpriseSnapshotPolicyRequest {
	s.CrossRegionCopyInfo = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetDesc(v string) *UpdateEnterpriseSnapshotPolicyRequest {
	s.Desc = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetName(v string) *UpdateEnterpriseSnapshotPolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetPolicyId(v string) *UpdateEnterpriseSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetRegionId(v string) *UpdateEnterpriseSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetRetainRule(v *UpdateEnterpriseSnapshotPolicyRequestRetainRule) *UpdateEnterpriseSnapshotPolicyRequest {
	s.RetainRule = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetSchedule(v *UpdateEnterpriseSnapshotPolicyRequestSchedule) *UpdateEnterpriseSnapshotPolicyRequest {
	s.Schedule = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetSpecialRetainRules(v *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) *UpdateEnterpriseSnapshotPolicyRequest {
	s.SpecialRetainRules = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetState(v string) *UpdateEnterpriseSnapshotPolicyRequest {
	s.State = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetStorageRule(v *UpdateEnterpriseSnapshotPolicyRequestStorageRule) *UpdateEnterpriseSnapshotPolicyRequest {
	s.StorageRule = v
	return s
}

type UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo struct {
	// Whether cross-region replication is enabled. The range of values:
	//
	// - true
	//
	// - false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Destination region information.
	Regions []*UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) SetEnabled(v bool) *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo {
	s.Enabled = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) SetRegions(v []*UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo {
	s.Regions = v
	return s
}

type UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions struct {
	// The region ID. You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which async replication is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Number of days to retain the destination snapshot. The range of values is greater than 1.
	RetainDays *int32 `json:"RetainDays,omitempty" xml:"RetainDays,omitempty"`
}

func (s UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) SetRegionId(v string) *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions {
	s.RegionId = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) SetRetainDays(v int32) *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions {
	s.RetainDays = &v
	return s
}

type UpdateEnterpriseSnapshotPolicyRequestRetainRule struct {
	// Maximum number of retained snapshots.
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// The time interval , valid value greater than 1.
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// The unit of time, valid values:
	//
	// - DAYS
	// - WEEKS
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s UpdateEnterpriseSnapshotPolicyRequestRetainRule) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestRetainRule) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestRetainRule) SetNumber(v int32) *UpdateEnterpriseSnapshotPolicyRequestRetainRule {
	s.Number = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestRetainRule) SetTimeInterval(v int32) *UpdateEnterpriseSnapshotPolicyRequestRetainRule {
	s.TimeInterval = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestRetainRule) SetTimeUnit(v string) *UpdateEnterpriseSnapshotPolicyRequestRetainRule {
	s.TimeUnit = &v
	return s
}

type UpdateEnterpriseSnapshotPolicyRequestSchedule struct {
	// The time when the policy will to be scheduled. Valid values: Set the parameter in a cron expression.
	//
	// For example, you can use `0 0 4 1/1 * ?` to specify 04:00:00 (UTC+8) on the first day of each month.
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
}

func (s UpdateEnterpriseSnapshotPolicyRequestSchedule) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestSchedule) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSchedule) SetCronExpression(v string) *UpdateEnterpriseSnapshotPolicyRequestSchedule {
	s.CronExpression = &v
	return s
}

type UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules struct {
	// Indicates whether the special retention is enabled.
	//
	// *   true: enable
	// *   false: disable
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The special retention rules.
	Rules []*UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) SetEnabled(v bool) *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules {
	s.Enabled = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) SetRules(v []*UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules {
	s.Rules = v
	return s
}

type UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules struct {
	// The periodic unit for specially retained snapshots. If configured to WEEKS, it provides special retention for the first snapshot of each week. The retention period is determined by TimeUnit and TimeInterval. The range of values are:
	// - WEEKS
	// - MONTHS
	// - YEARS"
	SpecialPeriodUnit *string `json:"SpecialPeriodUnit,omitempty" xml:"SpecialPeriodUnit,omitempty"`
	// Retention Time Value. The range of values is greater than 1.
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// Retention time unit for special snapshots. The range of values:
	//
	// - DAYS
	//
	// - WEEKS
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) SetSpecialPeriodUnit(v string) *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	s.SpecialPeriodUnit = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) SetTimeInterval(v int32) *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	s.TimeInterval = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) SetTimeUnit(v string) *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	s.TimeUnit = &v
	return s
}

type UpdateEnterpriseSnapshotPolicyRequestStorageRule struct {
	// Whether to enable the rapid availability of snapshots. The range of values:
	//
	// - true
	//
	// - false
	EnableImmediateAccess *bool `json:"EnableImmediateAccess,omitempty" xml:"EnableImmediateAccess,omitempty"`
}

func (s UpdateEnterpriseSnapshotPolicyRequestStorageRule) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestStorageRule) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestStorageRule) SetEnableImmediateAccess(v bool) *UpdateEnterpriseSnapshotPolicyRequestStorageRule {
	s.EnableImmediateAccess = &v
	return s
}

type UpdateEnterpriseSnapshotPolicyShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Snapshot replication destination information.
	CrossRegionCopyInfoShrink *string `json:"CrossRegionCopyInfo,omitempty" xml:"CrossRegionCopyInfo,omitempty"`
	// The description of the policy.
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The name of the policy.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The id of the policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID . You can call the [DescribeRegions](~~354276~~) operation to query the most recent list of regions in which snapshot policy is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Snapshot retention rule.
	RetainRuleShrink *string `json:"RetainRule,omitempty" xml:"RetainRule,omitempty"`
	// The rule for scheduling.
	ScheduleShrink *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The special snapshot retention rules.
	SpecialRetainRulesShrink *string `json:"SpecialRetainRules,omitempty" xml:"SpecialRetainRules,omitempty"`
	// The status of the policy. Valid values:
	//
	// *   **ENABLED**: Enable snapshot policy execution.
	// *   **DISABLED**: Disable snapshot policy execution.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Advanced snapshot features.
	StorageRuleShrink *string `json:"StorageRule,omitempty" xml:"StorageRule,omitempty"`
}

func (s UpdateEnterpriseSnapshotPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetClientToken(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetCrossRegionCopyInfoShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.CrossRegionCopyInfoShrink = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetDesc(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.Desc = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetName(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetPolicyId(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetRegionId(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetRetainRuleShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.RetainRuleShrink = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetScheduleShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.ScheduleShrink = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetSpecialRetainRulesShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.SpecialRetainRulesShrink = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetState(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.State = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetStorageRuleShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.StorageRuleShrink = &v
	return s
}

type UpdateEnterpriseSnapshotPolicyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEnterpriseSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyResponseBody) SetRequestId(v string) *UpdateEnterpriseSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateEnterpriseSnapshotPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnterpriseSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnterpriseSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyResponse) SetHeaders(v map[string]*string) *UpdateEnterpriseSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyResponse) SetStatusCode(v int32) *UpdateEnterpriseSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyResponse) SetBody(v *UpdateEnterpriseSnapshotPolicyResponseBody) *UpdateEnterpriseSnapshotPolicyResponse {
	s.Body = v
	return s
}

type UpdateSolutionInstanceAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The **ClientToken** value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the dedicated block storage cluster resides. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance.
	SolutionInstanceId *string `json:"SolutionInstanceId,omitempty" xml:"SolutionInstanceId,omitempty"`
}

func (s UpdateSolutionInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSolutionInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSolutionInstanceAttributeRequest) SetClientToken(v string) *UpdateSolutionInstanceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSolutionInstanceAttributeRequest) SetDescription(v string) *UpdateSolutionInstanceAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateSolutionInstanceAttributeRequest) SetName(v string) *UpdateSolutionInstanceAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateSolutionInstanceAttributeRequest) SetRegionId(v string) *UpdateSolutionInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSolutionInstanceAttributeRequest) SetSolutionInstanceId(v string) *UpdateSolutionInstanceAttributeRequest {
	s.SolutionInstanceId = &v
	return s
}

type UpdateSolutionInstanceAttributeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSolutionInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSolutionInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSolutionInstanceAttributeResponseBody) SetRequestId(v string) *UpdateSolutionInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSolutionInstanceAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSolutionInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSolutionInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSolutionInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSolutionInstanceAttributeResponse) SetHeaders(v map[string]*string) *UpdateSolutionInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSolutionInstanceAttributeResponse) SetStatusCode(v int32) *UpdateSolutionInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSolutionInstanceAttributeResponse) SetBody(v *UpdateSolutionInstanceAttributeResponseBody) *UpdateSolutionInstanceAttributeResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ebs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   A replication pair and a replication pair-consistent group replicate in the same direction if they have the same primary region (production region), primary zone (production zone), secondary region (disaster recovery region), and secondary zone (disaster recovery zone). A replication pair can be added only to a replication pair-consistent group that replicates in the same direction as the replication pair.
 * *   Before you can add a replication pair to a replication pair-consistent group, make sure that the pair and the group are in the **Created** (`created`) or **Stopped** (`stopped`) state.
 * *   Up to 17 replication pairs can be added to a single replication pair-consistent group.
 * *   After replication pairs are added to a replication pair-consistent group, the recovery point objective (RPO) of the group takes effect on the pairs in place of their original RPOs.
 *
 * @param request AddDiskReplicaPairRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddDiskReplicaPairResponse
 */
func (client *Client) AddDiskReplicaPairWithOptions(request *AddDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *AddDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   A replication pair and a replication pair-consistent group replicate in the same direction if they have the same primary region (production region), primary zone (production zone), secondary region (disaster recovery region), and secondary zone (disaster recovery zone). A replication pair can be added only to a replication pair-consistent group that replicates in the same direction as the replication pair.
 * *   Before you can add a replication pair to a replication pair-consistent group, make sure that the pair and the group are in the **Created** (`created`) or **Stopped** (`stopped`) state.
 * *   Up to 17 replication pairs can be added to a single replication pair-consistent group.
 * *   After replication pairs are added to a replication pair-consistent group, the recovery point objective (RPO) of the group takes effect on the pairs in place of their original RPOs.
 *
 * @param request AddDiskReplicaPairRequest
 * @return AddDiskReplicaPairResponse
 */
func (client *Client) AddDiskReplicaPair(request *AddDiskReplicaPairRequest) (_result *AddDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDiskReplicaPairResponse{}
	_body, _err := client.AddDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Usage notes
 * CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
 *
 * @param request ApplyLensServiceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ApplyLensServiceResponse
 */
func (client *Client) ApplyLensServiceWithOptions(runtime *util.RuntimeOptions) (_result *ApplyLensServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ApplyLensService"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyLensServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Usage notes
 * CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
 *
 * @return ApplyLensServiceResponse
 */
func (client *Client) ApplyLensService() (_result *ApplyLensServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyLensServiceResponse{}
	_body, _err := client.ApplyLensServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindEnterpriseSnapshotPolicyWithOptions(request *BindEnterpriseSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *BindEnterpriseSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DiskTargets)) {
		query["DiskTargets"] = request.DiskTargets
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindEnterpriseSnapshotPolicy"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindEnterpriseSnapshotPolicy(request *BindEnterpriseSnapshotPolicyRequest) (_result *BindEnterpriseSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.BindEnterpriseSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Usage notes
 * CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
 *
 * @param request CancelLensServiceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelLensServiceResponse
 */
func (client *Client) CancelLensServiceWithOptions(runtime *util.RuntimeOptions) (_result *CancelLensServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("CancelLensService"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelLensServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Usage notes
 * CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
 *
 * @return CancelLensServiceResponse
 */
func (client *Client) CancelLensService() (_result *CancelLensServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelLensServiceResponse{}
	_body, _err := client.CancelLensServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClearPairDrillWithOptions(request *ClearPairDrillRequest, runtime *util.RuntimeOptions) (_result *ClearPairDrillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrillId)) {
		query["DrillId"] = request.DrillId
	}

	if !tea.BoolValue(util.IsUnset(request.PairId)) {
		query["PairId"] = request.PairId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClearPairDrill"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClearPairDrillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearPairDrill(request *ClearPairDrillRequest) (_result *ClearPairDrillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearPairDrillResponse{}
	_body, _err := client.ClearPairDrillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClearReplicaGroupDrillWithOptions(request *ClearReplicaGroupDrillRequest, runtime *util.RuntimeOptions) (_result *ClearReplicaGroupDrillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrillId)) {
		query["DrillId"] = request.DrillId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClearReplicaGroupDrill"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClearReplicaGroupDrillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearReplicaGroupDrill(request *ClearReplicaGroupDrillRequest) (_result *ClearReplicaGroupDrillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearReplicaGroupDrillResponse{}
	_body, _err := client.ClearReplicaGroupDrillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Dedicated block storage clusters are physically isolated from public block storage clusters. The owner of each dedicated block storage cluster has exclusive access to all resources in the cluster.
 * Disks created in a dedicated block storage cluster can be attached only to Elastic Compute Service (ECS) instances that reside in the same zone as the cluster. Before you create a dedicated block storage cluster, decide the regions and zones in which to deploy your cloud resources.
 * Dedicated block storage clusters are classified into basic and performance types. When you create a dedicated block storage cluster, select a cluster type based on your business requirements.
 * You are charged for creating dedicated block storage clusters.
 *
 * @param request CreateDedicatedBlockStorageClusterRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDedicatedBlockStorageClusterResponse
 */
func (client *Client) CreateDedicatedBlockStorageClusterWithOptions(request *CreateDedicatedBlockStorageClusterRequest, runtime *util.RuntimeOptions) (_result *CreateDedicatedBlockStorageClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Azone)) {
		query["Azone"] = request.Azone
	}

	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		query["Capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.DbscId)) {
		query["DbscId"] = request.DbscId
	}

	if !tea.BoolValue(util.IsUnset(request.DbscName)) {
		query["DbscName"] = request.DbscName
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDedicatedBlockStorageCluster"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDedicatedBlockStorageClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Dedicated block storage clusters are physically isolated from public block storage clusters. The owner of each dedicated block storage cluster has exclusive access to all resources in the cluster.
 * Disks created in a dedicated block storage cluster can be attached only to Elastic Compute Service (ECS) instances that reside in the same zone as the cluster. Before you create a dedicated block storage cluster, decide the regions and zones in which to deploy your cloud resources.
 * Dedicated block storage clusters are classified into basic and performance types. When you create a dedicated block storage cluster, select a cluster type based on your business requirements.
 * You are charged for creating dedicated block storage clusters.
 *
 * @param request CreateDedicatedBlockStorageClusterRequest
 * @return CreateDedicatedBlockStorageClusterResponse
 */
func (client *Client) CreateDedicatedBlockStorageCluster(request *CreateDedicatedBlockStorageClusterRequest) (_result *CreateDedicatedBlockStorageClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDedicatedBlockStorageClusterResponse{}
	_body, _err := client.CreateDedicatedBlockStorageClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * The replication pair-consistent group feature allows you to batch manage multiple disks in disaster recovery scenarios. You can restore the data of all disks in the same replication pair-consistent group to the same point in time to allow for disaster recovery of instances.
 * Take note of the following items:
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   Replication pair-consistent groups can be used to implement disaster recovery across zones within the same region and disaster recovery across regions.
 * *   A replication pair and a replication pair-consistent group can replicate in the same direction if they have the same primary region (production region), primary zone (production zone), secondary region (disaster recovery region), and secondary zone (disaster recovery zone). A replication pair can be added to only a replication pair-consistent group that replicates in the same direction as the replication pair.
 * *   After replication pairs are added to a replication pair-consistent group, the recovery point objective (RPO) of the group takes effect on the pairs instead of their original RPOs.
 *
 * @param request CreateDiskReplicaGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDiskReplicaGroupResponse
 */
func (client *Client) CreateDiskReplicaGroupWithOptions(request *CreateDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *CreateDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationRegionId)) {
		query["DestinationRegionId"] = request.DestinationRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationZoneId)) {
		query["DestinationZoneId"] = request.DestinationZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RPO)) {
		query["RPO"] = request.RPO
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceZoneId)) {
		query["SourceZoneId"] = request.SourceZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * The replication pair-consistent group feature allows you to batch manage multiple disks in disaster recovery scenarios. You can restore the data of all disks in the same replication pair-consistent group to the same point in time to allow for disaster recovery of instances.
 * Take note of the following items:
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   Replication pair-consistent groups can be used to implement disaster recovery across zones within the same region and disaster recovery across regions.
 * *   A replication pair and a replication pair-consistent group can replicate in the same direction if they have the same primary region (production region), primary zone (production zone), secondary region (disaster recovery region), and secondary zone (disaster recovery zone). A replication pair can be added to only a replication pair-consistent group that replicates in the same direction as the replication pair.
 * *   After replication pairs are added to a replication pair-consistent group, the recovery point objective (RPO) of the group takes effect on the pairs instead of their original RPOs.
 *
 * @param request CreateDiskReplicaGroupRequest
 * @return CreateDiskReplicaGroupResponse
 */
func (client *Client) CreateDiskReplicaGroup(request *CreateDiskReplicaGroupRequest) (_result *CreateDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDiskReplicaGroupResponse{}
	_body, _err := client.CreateDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Async replication is a feature that protects data across regions by using the data replication capability of Elastic Block Storage (EBS). This feature can be used to asynchronously replicate data from a disk in one region to a disk in another region for disaster recovery purposes. You can use this feature to implement disaster recovery for critical business to protect data in your databases and improve business continuity.
 * Currently, the async replication feature can asynchronously replicate data only between enhanced SSDs (ESSDs). The functionality of disks in replication pairs is limited. You are charged on a subscription basis for the bandwidth that is used by the async replication feature.
 * Before you call this operation, take note of the following items:
 * *   Make sure that the source disk (primary disk) from which to replicate data and the destination disk (secondary disk) to which to replicate data are created. You can call the [CreateDisk](~~25513~~) operation to create disks.
 * *   The secondary disk cannot reside the same region as the primary disk. The async replication feature is supported in the China (Hangzhou), China (Shanghai), China (Beijing), China (Shenzhen), China (Heyuan), China (Chengdu), China (Hong Kong), Singapore, US (Silicon Valley), and US (Virginia) regions.
 * *   After you call this operation to create a replication pair, you must call the [StartDiskReplicaPair](~~354205~~) operation to enable async replication to periodically replicate data from the primary disk to the secondary disk across regions.
 *
 * @param request CreateDiskReplicaPairRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDiskReplicaPairResponse
 */
func (client *Client) CreateDiskReplicaPairWithOptions(request *CreateDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *CreateDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationDiskId)) {
		query["DestinationDiskId"] = request.DestinationDiskId
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationRegionId)) {
		query["DestinationRegionId"] = request.DestinationRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationZoneId)) {
		query["DestinationZoneId"] = request.DestinationZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.PairName)) {
		query["PairName"] = request.PairName
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RPO)) {
		query["RPO"] = request.RPO
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceZoneId)) {
		query["SourceZoneId"] = request.SourceZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Async replication is a feature that protects data across regions by using the data replication capability of Elastic Block Storage (EBS). This feature can be used to asynchronously replicate data from a disk in one region to a disk in another region for disaster recovery purposes. You can use this feature to implement disaster recovery for critical business to protect data in your databases and improve business continuity.
 * Currently, the async replication feature can asynchronously replicate data only between enhanced SSDs (ESSDs). The functionality of disks in replication pairs is limited. You are charged on a subscription basis for the bandwidth that is used by the async replication feature.
 * Before you call this operation, take note of the following items:
 * *   Make sure that the source disk (primary disk) from which to replicate data and the destination disk (secondary disk) to which to replicate data are created. You can call the [CreateDisk](~~25513~~) operation to create disks.
 * *   The secondary disk cannot reside the same region as the primary disk. The async replication feature is supported in the China (Hangzhou), China (Shanghai), China (Beijing), China (Shenzhen), China (Heyuan), China (Chengdu), China (Hong Kong), Singapore, US (Silicon Valley), and US (Virginia) regions.
 * *   After you call this operation to create a replication pair, you must call the [StartDiskReplicaPair](~~354205~~) operation to enable async replication to periodically replicate data from the primary disk to the secondary disk across regions.
 *
 * @param request CreateDiskReplicaPairRequest
 * @return CreateDiskReplicaPairResponse
 */
func (client *Client) CreateDiskReplicaPair(request *CreateDiskReplicaPairRequest) (_result *CreateDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDiskReplicaPairResponse{}
	_body, _err := client.CreateDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEnterpriseSnapshotPolicyWithOptions(tmpReq *CreateEnterpriseSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateEnterpriseSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateEnterpriseSnapshotPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CrossRegionCopyInfo)) {
		request.CrossRegionCopyInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CrossRegionCopyInfo, tea.String("CrossRegionCopyInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RetainRule)) {
		request.RetainRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RetainRule, tea.String("RetainRule"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Schedule)) {
		request.ScheduleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Schedule, tea.String("Schedule"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SpecialRetainRules)) {
		request.SpecialRetainRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SpecialRetainRules, tea.String("SpecialRetainRules"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StorageRule)) {
		request.StorageRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StorageRule, tea.String("StorageRule"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CrossRegionCopyInfoShrink)) {
		query["CrossRegionCopyInfo"] = request.CrossRegionCopyInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		query["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RetainRuleShrink)) {
		query["RetainRule"] = request.RetainRuleShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleShrink)) {
		query["Schedule"] = request.ScheduleShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SpecialRetainRulesShrink)) {
		query["SpecialRetainRules"] = request.SpecialRetainRulesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.StorageRuleShrink)) {
		query["StorageRule"] = request.StorageRuleShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEnterpriseSnapshotPolicy"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEnterpriseSnapshotPolicy(request *CreateEnterpriseSnapshotPolicyRequest) (_result *CreateEnterpriseSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CreateEnterpriseSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Currently, this API is only available for use with ACS resource hosting and is not yet open for direct invocation.
 *
 * @param request DeleteDiskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDiskResponse
 */
func (client *Client) DeleteDiskWithOptions(request *DeleteDiskRequest, runtime *util.RuntimeOptions) (_result *DeleteDiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDisk"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Currently, this API is only available for use with ACS resource hosting and is not yet open for direct invocation.
 *
 * @param request DeleteDiskRequest
 * @return DeleteDiskResponse
 */
func (client *Client) DeleteDisk(request *DeleteDiskRequest) (_result *DeleteDiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDiskResponse{}
	_body, _err := client.DeleteDiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   Before you can delete a replication pair-consistent group, make sure that no replication pairs exist in the group.
 * *   The replication pair-consistent group that you want to delete must be in the **Created** (`created`), **Creation Failed** (`create_failed`), **Stopped** (`stopped`), **Failovered** (`failovered`), **Deleting** (`deleting`), **Deletion Failed** (`delete_failed`), or **Invalid** (`invalid`) state.
 *
 * @param request DeleteDiskReplicaGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDiskReplicaGroupResponse
 */
func (client *Client) DeleteDiskReplicaGroupWithOptions(request *DeleteDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   Before you can delete a replication pair-consistent group, make sure that no replication pairs exist in the group.
 * *   The replication pair-consistent group that you want to delete must be in the **Created** (`created`), **Creation Failed** (`create_failed`), **Stopped** (`stopped`), **Failovered** (`failovered`), **Deleting** (`deleting`), **Deletion Failed** (`delete_failed`), or **Invalid** (`invalid`) state.
 *
 * @param request DeleteDiskReplicaGroupRequest
 * @return DeleteDiskReplicaGroupResponse
 */
func (client *Client) DeleteDiskReplicaGroup(request *DeleteDiskReplicaGroupRequest) (_result *DeleteDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDiskReplicaGroupResponse{}
	_body, _err := client.DeleteDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   Only replication pairs that are in the **Stopped** (`stopped`), **Invalid** (`invalid`), or **Failovered** (`failovered`) state can be deleted. This operation deletes only replication pairs. The primary and secondary disks in the deleted replication pairs are retained.
 * *   To delete a replication pair, you must call this operation in the region where the primary disk is located. After the replication pair is deleted, the functionality limits are lifted from the primary and secondary disks. For example, you can attach the secondary disk, resize the disk, or read data from or write data to the disk.
 *
 * @param request DeleteDiskReplicaPairRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDiskReplicaPairResponse
 */
func (client *Client) DeleteDiskReplicaPairWithOptions(request *DeleteDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *DeleteDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   Only replication pairs that are in the **Stopped** (`stopped`), **Invalid** (`invalid`), or **Failovered** (`failovered`) state can be deleted. This operation deletes only replication pairs. The primary and secondary disks in the deleted replication pairs are retained.
 * *   To delete a replication pair, you must call this operation in the region where the primary disk is located. After the replication pair is deleted, the functionality limits are lifted from the primary and secondary disks. For example, you can attach the secondary disk, resize the disk, or read data from or write data to the disk.
 *
 * @param request DeleteDiskReplicaPairRequest
 * @return DeleteDiskReplicaPairResponse
 */
func (client *Client) DeleteDiskReplicaPair(request *DeleteDiskReplicaPairRequest) (_result *DeleteDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDiskReplicaPairResponse{}
	_body, _err := client.DeleteDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEnterpriseSnapshotPolicyWithOptions(request *DeleteEnterpriseSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteEnterpriseSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEnterpriseSnapshotPolicy"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEnterpriseSnapshotPolicy(request *DeleteEnterpriseSnapshotPolicyRequest) (_result *DeleteEnterpriseSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.DeleteEnterpriseSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   You can use one of the following methods to check the responses:
 *     *   Method 1: Use `NextToken` to configure the query token. Set the value to the `NextToken` value that is returned in the last call to the DescribeDisks operation. Then, use `MaxResults` to specify the maximum number of entries to return on each page.
 *     *   Method 2: Use `PageSize` to specify the number of entries to return on each page and then use `PageNumber` to specify the number of the page to return.
 *         You can use only one of the preceding methods. If a large number of entries are to be returned, we recommend that you use method 1. When `NextToken` is specified, `PageSize` and `PageNumber` do not take effect and `TotalCount` in the response is invalid.
 * *   A disk that has the multi-attach feature enabled can be attached to multiple instances. You can query the attachment information of the disk based on the `Attachment` values in the response.
 * When you call an API operation by using Alibaba Cloud CLI, you must specify request parameter values of different data types in the required formats. For more information, see [Parameter format overview](~~110340~~).
 *
 * @param request DescribeDedicatedBlockStorageClusterDisksRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDedicatedBlockStorageClusterDisksResponse
 */
func (client *Client) DescribeDedicatedBlockStorageClusterDisksWithOptions(request *DescribeDedicatedBlockStorageClusterDisksRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedBlockStorageClusterDisksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbscId)) {
		query["DbscId"] = request.DbscId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDedicatedBlockStorageClusterDisks"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDedicatedBlockStorageClusterDisksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   You can use one of the following methods to check the responses:
 *     *   Method 1: Use `NextToken` to configure the query token. Set the value to the `NextToken` value that is returned in the last call to the DescribeDisks operation. Then, use `MaxResults` to specify the maximum number of entries to return on each page.
 *     *   Method 2: Use `PageSize` to specify the number of entries to return on each page and then use `PageNumber` to specify the number of the page to return.
 *         You can use only one of the preceding methods. If a large number of entries are to be returned, we recommend that you use method 1. When `NextToken` is specified, `PageSize` and `PageNumber` do not take effect and `TotalCount` in the response is invalid.
 * *   A disk that has the multi-attach feature enabled can be attached to multiple instances. You can query the attachment information of the disk based on the `Attachment` values in the response.
 * When you call an API operation by using Alibaba Cloud CLI, you must specify request parameter values of different data types in the required formats. For more information, see [Parameter format overview](~~110340~~).
 *
 * @param request DescribeDedicatedBlockStorageClusterDisksRequest
 * @return DescribeDedicatedBlockStorageClusterDisksResponse
 */
func (client *Client) DescribeDedicatedBlockStorageClusterDisks(request *DescribeDedicatedBlockStorageClusterDisksRequest) (_result *DescribeDedicatedBlockStorageClusterDisksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedBlockStorageClusterDisksResponse{}
	_body, _err := client.DescribeDedicatedBlockStorageClusterDisksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * >  The Dedicated Block Storage Cluster feature is available only in the China (Heyuan), Indonesia (Jakarta), and China (Shenzhen) regions.
 * *   You can specify multiple request parameters to be queried. Specified parameters are evaluated by using the AND operator. Only the specified parameters are included in the filter conditions.
 * *   We recommend that you use NextToken and MaxResults to perform paged queries. We recommend that you use MaxResults to specify the maximum number of entries to return in each request. The return value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. When you call the DescribeDedicatedBlockStorageClusters operation to retrieve a new page of results, set NextToken to the NextToken value that is returned in the previous call and specify MaxResults to limit the number of entries returned.
 *
 * @param request DescribeDedicatedBlockStorageClustersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDedicatedBlockStorageClustersResponse
 */
func (client *Client) DescribeDedicatedBlockStorageClustersWithOptions(request *DescribeDedicatedBlockStorageClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedBlockStorageClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedBlockStorageClusterId)) {
		query["DedicatedBlockStorageClusterId"] = request.DedicatedBlockStorageClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AzoneId)) {
		body["AzoneId"] = request.AzoneId
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDedicatedBlockStorageClusters"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDedicatedBlockStorageClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * >  The Dedicated Block Storage Cluster feature is available only in the China (Heyuan), Indonesia (Jakarta), and China (Shenzhen) regions.
 * *   You can specify multiple request parameters to be queried. Specified parameters are evaluated by using the AND operator. Only the specified parameters are included in the filter conditions.
 * *   We recommend that you use NextToken and MaxResults to perform paged queries. We recommend that you use MaxResults to specify the maximum number of entries to return in each request. The return value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. When you call the DescribeDedicatedBlockStorageClusters operation to retrieve a new page of results, set NextToken to the NextToken value that is returned in the previous call and specify MaxResults to limit the number of entries returned.
 *
 * @param request DescribeDedicatedBlockStorageClustersRequest
 * @return DescribeDedicatedBlockStorageClustersResponse
 */
func (client *Client) DescribeDedicatedBlockStorageClusters(request *DescribeDedicatedBlockStorageClustersRequest) (_result *DescribeDedicatedBlockStorageClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedBlockStorageClustersResponse{}
	_body, _err := client.DescribeDedicatedBlockStorageClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Usage notes
 * CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
 *
 * @param request DescribeDiskEventsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDiskEventsResponse
 */
func (client *Client) DescribeDiskEventsWithOptions(request *DescribeDiskEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiskEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiskCategory)) {
		query["DiskCategory"] = request.DiskCategory
	}

	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiskEvents"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiskEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Usage notes
 * CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
 *
 * @param request DescribeDiskEventsRequest
 * @return DescribeDiskEventsResponse
 */
func (client *Client) DescribeDiskEvents(request *DescribeDiskEventsRequest) (_result *DescribeDiskEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiskEventsResponse{}
	_body, _err := client.DescribeDiskEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Usage notes
 * *   CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
 * *   Up to 400 monitoring data entries can be returned at a time. An error is returned if the value calculated based on the following formula is greater than 400: `(EndTime - StartTime)/Period`.
 * *   You can query the monitoring data collected in the last three days. An error is returned if the time specified by `StartTime` is more than three days prior to the current time.
 *
 * @param request DescribeDiskMonitorDataRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDiskMonitorDataResponse
 */
func (client *Client) DescribeDiskMonitorDataWithOptions(request *DescribeDiskMonitorDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDiskMonitorDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiskMonitorData"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiskMonitorDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Usage notes
 * *   CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
 * *   Up to 400 monitoring data entries can be returned at a time. An error is returned if the value calculated based on the following formula is greater than 400: `(EndTime - StartTime)/Period`.
 * *   You can query the monitoring data collected in the last three days. An error is returned if the time specified by `StartTime` is more than three days prior to the current time.
 *
 * @param request DescribeDiskMonitorDataRequest
 * @return DescribeDiskMonitorDataResponse
 */
func (client *Client) DescribeDiskMonitorData(request *DescribeDiskMonitorDataRequest) (_result *DescribeDiskMonitorDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiskMonitorDataResponse{}
	_body, _err := client.DescribeDiskMonitorDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Usage notes
 * CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
 *
 * @param request DescribeDiskMonitorDataListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDiskMonitorDataListResponse
 */
func (client *Client) DescribeDiskMonitorDataListWithOptions(request *DescribeDiskMonitorDataListRequest, runtime *util.RuntimeOptions) (_result *DescribeDiskMonitorDataListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiskIds)) {
		query["DiskIds"] = request.DiskIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiskMonitorDataList"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiskMonitorDataListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Usage notes
 * CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
 *
 * @param request DescribeDiskMonitorDataListRequest
 * @return DescribeDiskMonitorDataListResponse
 */
func (client *Client) DescribeDiskMonitorDataList(request *DescribeDiskMonitorDataListRequest) (_result *DescribeDiskMonitorDataListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiskMonitorDataListResponse{}
	_body, _err := client.DescribeDiskMonitorDataListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * To perform a paged query, set the MaxResults and NextToken parameters.
 * During a paged query, when you call the DescribeDiskReplicaGroups operation to retrieve the first page of results, set `MaxResults` to specify the maximum number of entries to return in the call. The return value of `NextToken` is a pagination token, which can be used in the next call to retrieve a new page of results. When you call the DescribeDiskReplicaGroups operation to retrieve a new page of results, set `NextToken` to the `NextToken` value returned in the previous call and set MaxResults to specify the maximum number of entries to return in this call.
 *
 * @param request DescribeDiskReplicaGroupsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDiskReplicaGroupsResponse
 */
func (client *Client) DescribeDiskReplicaGroupsWithOptions(request *DescribeDiskReplicaGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiskReplicaGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		query["GroupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Site)) {
		query["Site"] = request.Site
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiskReplicaGroups"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiskReplicaGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * To perform a paged query, set the MaxResults and NextToken parameters.
 * During a paged query, when you call the DescribeDiskReplicaGroups operation to retrieve the first page of results, set `MaxResults` to specify the maximum number of entries to return in the call. The return value of `NextToken` is a pagination token, which can be used in the next call to retrieve a new page of results. When you call the DescribeDiskReplicaGroups operation to retrieve a new page of results, set `NextToken` to the `NextToken` value returned in the previous call and set MaxResults to specify the maximum number of entries to return in this call.
 *
 * @param request DescribeDiskReplicaGroupsRequest
 * @return DescribeDiskReplicaGroupsResponse
 */
func (client *Client) DescribeDiskReplicaGroups(request *DescribeDiskReplicaGroupsRequest) (_result *DescribeDiskReplicaGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiskReplicaGroupsResponse{}
	_body, _err := client.DescribeDiskReplicaGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiskReplicaPairProgressWithOptions(request *DescribeDiskReplicaPairProgressRequest, runtime *util.RuntimeOptions) (_result *DescribeDiskReplicaPairProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiskReplicaPairProgress"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiskReplicaPairProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiskReplicaPairProgress(request *DescribeDiskReplicaPairProgressRequest) (_result *DescribeDiskReplicaPairProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiskReplicaPairProgressResponse{}
	_body, _err := client.DescribeDiskReplicaPairProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   When you call this operation for a specific region, if the primary disk (source disk) or secondary disk (destination disk) of a replication pair resides in the region, information about the replication pair is displayed in the response.
 * *   If you want to perform a paged query, configure the `NextToken` and `MaxResults` parameters. During a paged query, when you call the DescribeDiskReplicaPairs operation to retrieve the first page of results, set `MaxResults` to limit the maximum number of entries to return in the call. The return value of NextToken is a pagination token, which can be used in the next call to retrieve a new page of results. When you call the DescribeDiskReplicaPairs operation to retrieve a new page of results, set NextToken to the NextToken value returned in the previous call and set MaxResults to specify the maximum number of entries to return in this call.
 *
 * @param request DescribeDiskReplicaPairsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDiskReplicaPairsResponse
 */
func (client *Client) DescribeDiskReplicaPairsWithOptions(request *DescribeDiskReplicaPairsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiskReplicaPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PairIds)) {
		query["PairIds"] = request.PairIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Site)) {
		query["Site"] = request.Site
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiskReplicaPairs"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiskReplicaPairsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   When you call this operation for a specific region, if the primary disk (source disk) or secondary disk (destination disk) of a replication pair resides in the region, information about the replication pair is displayed in the response.
 * *   If you want to perform a paged query, configure the `NextToken` and `MaxResults` parameters. During a paged query, when you call the DescribeDiskReplicaPairs operation to retrieve the first page of results, set `MaxResults` to limit the maximum number of entries to return in the call. The return value of NextToken is a pagination token, which can be used in the next call to retrieve a new page of results. When you call the DescribeDiskReplicaPairs operation to retrieve a new page of results, set NextToken to the NextToken value returned in the previous call and set MaxResults to specify the maximum number of entries to return in this call.
 *
 * @param request DescribeDiskReplicaPairsRequest
 * @return DescribeDiskReplicaPairsResponse
 */
func (client *Client) DescribeDiskReplicaPairs(request *DescribeDiskReplicaPairsRequest) (_result *DescribeDiskReplicaPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiskReplicaPairsResponse{}
	_body, _err := client.DescribeDiskReplicaPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDisksWithOptions(request *DescribeDisksRequest, runtime *util.RuntimeOptions) (_result *DescribeDisksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.DiskIds)) {
		query["DiskIds"] = request.DiskIds
	}

	if !tea.BoolValue(util.IsUnset(request.DiskName)) {
		query["DiskName"] = request.DiskName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDisks"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDisksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDisks(request *DescribeDisksRequest) (_result *DescribeDisksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDisksResponse{}
	_body, _err := client.DescribeDisksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnterpriseSnapshotPolicyWithOptions(request *DescribeEnterpriseSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeEnterpriseSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DiskIds)) {
		query["DiskIds"] = request.DiskIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyIds)) {
		query["PolicyIds"] = request.PolicyIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEnterpriseSnapshotPolicy"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnterpriseSnapshotPolicy(request *DescribeEnterpriseSnapshotPolicyRequest) (_result *DescribeEnterpriseSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.DescribeEnterpriseSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventsWithOptions(request *DescribeEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		query["EventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEvents"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEvents(request *DescribeEventsRequest) (_result *DescribeEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventsResponse{}
	_body, _err := client.DescribeEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Usage notes
 * CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
 *
 * @param request DescribeLensServiceStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeLensServiceStatusResponse
 */
func (client *Client) DescribeLensServiceStatusWithOptions(runtime *util.RuntimeOptions) (_result *DescribeLensServiceStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeLensServiceStatus"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLensServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Usage notes
 * CloudLens for EBS is in invitational preview in the China (Hangzhou), China (Shanghai), China (Zhangjiakou), China (Shenzhen), and China (Hong Kong) regions. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
 *
 * @return DescribeLensServiceStatusResponse
 */
func (client *Client) DescribeLensServiceStatus() (_result *DescribeLensServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLensServiceStatusResponse{}
	_body, _err := client.DescribeLensServiceStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetricDataWithOptions(request *DescribeMetricDataRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMetricData"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetricData(request *DescribeMetricDataRequest) (_result *DescribeMetricDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricDataResponse{}
	_body, _err := client.DescribeMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePairDrillsWithOptions(request *DescribePairDrillsRequest, runtime *util.RuntimeOptions) (_result *DescribePairDrillsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrillId)) {
		query["DrillId"] = request.DrillId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PairId)) {
		query["PairId"] = request.PairId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePairDrills"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePairDrillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePairDrills(request *DescribePairDrillsRequest) (_result *DescribePairDrillsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePairDrillsResponse{}
	_body, _err := client.DescribePairDrillsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeReplicaGroupDrillsWithOptions(request *DescribeReplicaGroupDrillsRequest, runtime *util.RuntimeOptions) (_result *DescribeReplicaGroupDrillsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrillId)) {
		query["DrillId"] = request.DrillId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeReplicaGroupDrills"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeReplicaGroupDrillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeReplicaGroupDrills(request *DescribeReplicaGroupDrillsRequest) (_result *DescribeReplicaGroupDrillsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeReplicaGroupDrillsResponse{}
	_body, _err := client.DescribeReplicaGroupDrillsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSolutionInstanceConfigurationWithOptions(request *DescribeSolutionInstanceConfigurationRequest, runtime *util.RuntimeOptions) (_result *DescribeSolutionInstanceConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionId)) {
		query["SolutionId"] = request.SolutionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSolutionInstanceConfiguration"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSolutionInstanceConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSolutionInstanceConfiguration(request *DescribeSolutionInstanceConfigurationRequest) (_result *DescribeSolutionInstanceConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSolutionInstanceConfigurationResponse{}
	_body, _err := client.DescribeSolutionInstanceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   The replication pair-consistent group must be in the **One-time Syncing** (`manual_syncing`), **Syncing** (`syncing`), **Normal** (`normal`), **Stopping** (`stopping`), **Stop Failed** (`stop_failed`), **Stopped** (`stopped`), **In Failover** (`failovering`), **Failover Failed** (`failover_failed`), or **Failovered** (`failovered`) state.
 * *   After a failover is performed, the replication pair-consistent group enters the **Failovered** (`failovered`) state.
 * *   Before you perform a failover, make sure that the first full data synchronization is completed between the primary site and secondary site.
 *
 * @param request FailoverDiskReplicaGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return FailoverDiskReplicaGroupResponse
 */
func (client *Client) FailoverDiskReplicaGroupWithOptions(request *FailoverDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *FailoverDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FailoverDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FailoverDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   The replication pair-consistent group must be in the **One-time Syncing** (`manual_syncing`), **Syncing** (`syncing`), **Normal** (`normal`), **Stopping** (`stopping`), **Stop Failed** (`stop_failed`), **Stopped** (`stopped`), **In Failover** (`failovering`), **Failover Failed** (`failover_failed`), or **Failovered** (`failovered`) state.
 * *   After a failover is performed, the replication pair-consistent group enters the **Failovered** (`failovered`) state.
 * *   Before you perform a failover, make sure that the first full data synchronization is completed between the primary site and secondary site.
 *
 * @param request FailoverDiskReplicaGroupRequest
 * @return FailoverDiskReplicaGroupResponse
 */
func (client *Client) FailoverDiskReplicaGroup(request *FailoverDiskReplicaGroupRequest) (_result *FailoverDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FailoverDiskReplicaGroupResponse{}
	_body, _err := client.FailoverDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   The replication pair for which you want to enable failover cannot be in the **Invalid** (`invalid`) or **Deleted** (`deleted`) state.
 * *   After a failover is performed, the replication pair enters the **Failovered** (`failovered`) state.
 *
 * @param request FailoverDiskReplicaPairRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return FailoverDiskReplicaPairResponse
 */
func (client *Client) FailoverDiskReplicaPairWithOptions(request *FailoverDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *FailoverDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FailoverDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FailoverDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   The replication pair for which you want to enable failover cannot be in the **Invalid** (`invalid`) or **Deleted** (`deleted`) state.
 * *   After a failover is performed, the replication pair enters the **Failovered** (`failovered`) state.
 *
 * @param request FailoverDiskReplicaPairRequest
 * @return FailoverDiskReplicaPairResponse
 */
func (client *Client) FailoverDiskReplicaPair(request *FailoverDiskReplicaPairRequest) (_result *FailoverDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FailoverDiskReplicaPairResponse{}
	_body, _err := client.FailoverDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiskWithOptions(request *GetDiskRequest, runtime *util.RuntimeOptions) (_result *GetDiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDisk"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDisk(request *GetDiskRequest) (_result *GetDiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiskResponse{}
	_body, _err := client.GetDiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Specify at least one of the following parameters or parameter pairs in a request to determine a query object:
 * *   `ResourceId.N`
 * *   `Tag.N` parameter pair (`Tag.N.Key` and `Tag.N.Value`)
 * If you set `Tag.N` and `ResourceId.N` at the same time, the EBS resources that match both the parameters are returned.
 *
 * @param request ListTagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagResourcesResponse
 */
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Specify at least one of the following parameters or parameter pairs in a request to determine a query object:
 * *   `ResourceId.N`
 * *   `Tag.N` parameter pair (`Tag.N.Key` and `Tag.N.Value`)
 * If you set `Tag.N` and `ResourceId.N` at the same time, the EBS resources that match both the parameters are returned.
 *
 * @param request ListTagResourcesRequest
 * @return ListTagResourcesResponse
 */
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to modify the information of a dedicated block storage cluster. The information includes the name and description of the cluster.
 *
 * @param request ModifyDedicatedBlockStorageClusterAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDedicatedBlockStorageClusterAttributeResponse
 */
func (client *Client) ModifyDedicatedBlockStorageClusterAttributeWithOptions(request *ModifyDedicatedBlockStorageClusterAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyDedicatedBlockStorageClusterAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DbscId)) {
		query["DbscId"] = request.DbscId
	}

	if !tea.BoolValue(util.IsUnset(request.DbscName)) {
		query["DbscName"] = request.DbscName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDedicatedBlockStorageClusterAttribute"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDedicatedBlockStorageClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to modify the information of a dedicated block storage cluster. The information includes the name and description of the cluster.
 *
 * @param request ModifyDedicatedBlockStorageClusterAttributeRequest
 * @return ModifyDedicatedBlockStorageClusterAttributeResponse
 */
func (client *Client) ModifyDedicatedBlockStorageClusterAttribute(request *ModifyDedicatedBlockStorageClusterAttributeRequest) (_result *ModifyDedicatedBlockStorageClusterAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDedicatedBlockStorageClusterAttributeResponse{}
	_body, _err := client.ModifyDedicatedBlockStorageClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   The replication pair-consistent group must be in the **Created** (`created`) or **Stopped** (`stopped`) state.
 *
 * @param request ModifyDiskReplicaGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDiskReplicaGroupResponse
 */
func (client *Client) ModifyDiskReplicaGroupWithOptions(request *ModifyDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RPO)) {
		query["RPO"] = request.RPO
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   The replication pair-consistent group must be in the **Created** (`created`) or **Stopped** (`stopped`) state.
 *
 * @param request ModifyDiskReplicaGroupRequest
 * @return ModifyDiskReplicaGroupResponse
 */
func (client *Client) ModifyDiskReplicaGroup(request *ModifyDiskReplicaGroupRequest) (_result *ModifyDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDiskReplicaGroupResponse{}
	_body, _err := client.ModifyDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   Only replication pairs that are in the **Created** (`created`) or **Stopped** (`stopped`) state can have their names or descriptions modified.
 *
 * @param request ModifyDiskReplicaPairRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDiskReplicaPairResponse
 */
func (client *Client) ModifyDiskReplicaPairWithOptions(request *ModifyDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *ModifyDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PairName)) {
		query["PairName"] = request.PairName
	}

	if !tea.BoolValue(util.IsUnset(request.RPO)) {
		query["RPO"] = request.RPO
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   Only replication pairs that are in the **Created** (`created`) or **Stopped** (`stopped`) state can have their names or descriptions modified.
 *
 * @param request ModifyDiskReplicaPairRequest
 * @return ModifyDiskReplicaPairResponse
 */
func (client *Client) ModifyDiskReplicaPair(request *ModifyDiskReplicaPairRequest) (_result *ModifyDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDiskReplicaPairResponse{}
	_body, _err := client.ModifyDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDedicatedBlockStorageClusterDiskThroughputStatusWithOptions(request *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest, runtime *util.RuntimeOptions) (_result *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QosRequestId)) {
		body["QosRequestId"] = request.QosRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDedicatedBlockStorageClusterDiskThroughputStatus"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDedicatedBlockStorageClusterDiskThroughputStatus(request *QueryDedicatedBlockStorageClusterDiskThroughputStatusRequest) (_result *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse{}
	_body, _err := client.QueryDedicatedBlockStorageClusterDiskThroughputStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Period is the time interval between data retrieval points. When set to 60 (minute interval), a maximum of 1440 data points can be returned; when set to 3600 (hour interval), a maximum of 744 data points can be returned; and when set to 86400 (day interval), a maximum of 366 data points can be returned.
 *
 * @param request QueryDedicatedBlockStorageClusterInventoryDataRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryDedicatedBlockStorageClusterInventoryDataResponse
 */
func (client *Client) QueryDedicatedBlockStorageClusterInventoryDataWithOptions(request *QueryDedicatedBlockStorageClusterInventoryDataRequest, runtime *util.RuntimeOptions) (_result *QueryDedicatedBlockStorageClusterInventoryDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbscId)) {
		body["DbscId"] = request.DbscId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDedicatedBlockStorageClusterInventoryData"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDedicatedBlockStorageClusterInventoryDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Period is the time interval between data retrieval points. When set to 60 (minute interval), a maximum of 1440 data points can be returned; when set to 3600 (hour interval), a maximum of 744 data points can be returned; and when set to 86400 (day interval), a maximum of 366 data points can be returned.
 *
 * @param request QueryDedicatedBlockStorageClusterInventoryDataRequest
 * @return QueryDedicatedBlockStorageClusterInventoryDataResponse
 */
func (client *Client) QueryDedicatedBlockStorageClusterInventoryData(request *QueryDedicatedBlockStorageClusterInventoryDataRequest) (_result *QueryDedicatedBlockStorageClusterInventoryDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDedicatedBlockStorageClusterInventoryDataResponse{}
	_body, _err := client.QueryDedicatedBlockStorageClusterInventoryDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   The replication pair-consistent group from which you want to remove a replication pair must be in the **Created** (`created`), **Stopped** (`stopped`), or **Invalid** (`invalid`) state.
 *
 * @param request RemoveDiskReplicaPairRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveDiskReplicaPairResponse
 */
func (client *Client) RemoveDiskReplicaPairWithOptions(request *RemoveDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *RemoveDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   The replication pair-consistent group from which you want to remove a replication pair must be in the **Created** (`created`), **Stopped** (`stopped`), or **Invalid** (`invalid`) state.
 *
 * @param request RemoveDiskReplicaPairRequest
 * @return RemoveDiskReplicaPairResponse
 */
func (client *Client) RemoveDiskReplicaPair(request *RemoveDiskReplicaPairRequest) (_result *RemoveDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveDiskReplicaPairResponse{}
	_body, _err := client.RemoveDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   The replication pair-consistent group for which you want to enable reverse replication must be in the **Failovered** (`failovered`) state. You can call the `FailoverDiskReplicaPair` operation to enable failover.
 * *   Before a reverse replication is performed, the primary disks must be detached from its associated Elastic Compute Service (ECS) instance and must be in the Unattached state. You can call the [DetachDisk](~~25516~~) operation to detach the disks.
 * *   After you enable reverse replication, you must call the `StartDiskReplicaPair` operation again to enable the async replication feature before data can be replicated from the original secondary disks to the original primary disks.
 * *   You can set the ReverseReplicate parameter to false to cancel the **Failovered** (`failovered`) state and restore the original replication direction.
 *
 * @param request ReprotectDiskReplicaGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ReprotectDiskReplicaGroupResponse
 */
func (client *Client) ReprotectDiskReplicaGroupWithOptions(request *ReprotectDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *ReprotectDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ReverseReplicate)) {
		query["ReverseReplicate"] = request.ReverseReplicate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReprotectDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReprotectDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   The replication pair-consistent group for which you want to enable reverse replication must be in the **Failovered** (`failovered`) state. You can call the `FailoverDiskReplicaPair` operation to enable failover.
 * *   Before a reverse replication is performed, the primary disks must be detached from its associated Elastic Compute Service (ECS) instance and must be in the Unattached state. You can call the [DetachDisk](~~25516~~) operation to detach the disks.
 * *   After you enable reverse replication, you must call the `StartDiskReplicaPair` operation again to enable the async replication feature before data can be replicated from the original secondary disks to the original primary disks.
 * *   You can set the ReverseReplicate parameter to false to cancel the **Failovered** (`failovered`) state and restore the original replication direction.
 *
 * @param request ReprotectDiskReplicaGroupRequest
 * @return ReprotectDiskReplicaGroupResponse
 */
func (client *Client) ReprotectDiskReplicaGroup(request *ReprotectDiskReplicaGroupRequest) (_result *ReprotectDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReprotectDiskReplicaGroupResponse{}
	_body, _err := client.ReprotectDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   The replication pair for which you want to enable reverse replication must be in the **Failovered** (`failovered`) state. You can call the [FailoverDiskReplicaPair](~~354358~~) operation to enable failover.
 * *   The primary disk must be detached from its associated Elastic Compute Service (ECS) instance and is in the Unattached state. You can call the [DetachDisk](~~25516~~) operation to detach the disk.
 * *   After you enable reverse replication, you must call the [StartDiskReplicaPair](~~354205~~) operation again to activate the replication pair before data can be replicated from the original secondary disk to the original primary disk.
 * *   You can set the ReverseReplicate parameter to false to cancel the **Failovered** (`failovered`) state and restore the original replication direction.
 *
 * @param request ReprotectDiskReplicaPairRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ReprotectDiskReplicaPairResponse
 */
func (client *Client) ReprotectDiskReplicaPairWithOptions(request *ReprotectDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *ReprotectDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	if !tea.BoolValue(util.IsUnset(request.ReverseReplicate)) {
		query["ReverseReplicate"] = request.ReverseReplicate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReprotectDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReprotectDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   The replication pair for which you want to enable reverse replication must be in the **Failovered** (`failovered`) state. You can call the [FailoverDiskReplicaPair](~~354358~~) operation to enable failover.
 * *   The primary disk must be detached from its associated Elastic Compute Service (ECS) instance and is in the Unattached state. You can call the [DetachDisk](~~25516~~) operation to detach the disk.
 * *   After you enable reverse replication, you must call the [StartDiskReplicaPair](~~354205~~) operation again to activate the replication pair before data can be replicated from the original secondary disk to the original primary disk.
 * *   You can set the ReverseReplicate parameter to false to cancel the **Failovered** (`failovered`) state and restore the original replication direction.
 *
 * @param request ReprotectDiskReplicaPairRequest
 * @return ReprotectDiskReplicaPairResponse
 */
func (client *Client) ReprotectDiskReplicaPair(request *ReprotectDiskReplicaPairRequest) (_result *ReprotectDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReprotectDiskReplicaPairResponse{}
	_body, _err := client.ReprotectDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDedicatedBlockStorageClusterDiskThroughputWithOptions(request *SetDedicatedBlockStorageClusterDiskThroughputRequest, runtime *util.RuntimeOptions) (_result *SetDedicatedBlockStorageClusterDiskThroughputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bps)) {
		body["Bps"] = request.Bps
	}

	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		body["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDedicatedBlockStorageClusterDiskThroughput"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDedicatedBlockStorageClusterDiskThroughputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDedicatedBlockStorageClusterDiskThroughput(request *SetDedicatedBlockStorageClusterDiskThroughputRequest) (_result *SetDedicatedBlockStorageClusterDiskThroughputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDedicatedBlockStorageClusterDiskThroughputResponse{}
	_body, _err := client.SetDedicatedBlockStorageClusterDiskThroughputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   If you set the `OneShot` to `false`, the replication pair-consistent group must be in the **Created** (`created` ), **Synchronizing** (`syncing` ), **Normal** (`normal` ), or **Stopped** (`stopped`) state.
 * *   If you set `OneShot` to `true`, the replication pair-consistent group must be in the **Created** (`created` ), **One-time Syncing** (`manual_syncing` ), or **Stopped** (`stopped`) state. The time interval between two consecutive one-time synchronizations must be longer than one half of the recovery point objective (RPO).
 * *   After a replication pair-consistent group is activated, the group enters the **Initial Syncing** (`initial_syncing`) state and the system performs the first async replication to replicate all data from the primary disks to secondary disks.
 *
 * @param request StartDiskReplicaGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartDiskReplicaGroupResponse
 */
func (client *Client) StartDiskReplicaGroupWithOptions(request *StartDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *StartDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OneShot)) {
		query["OneShot"] = request.OneShot
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   If you set the `OneShot` to `false`, the replication pair-consistent group must be in the **Created** (`created` ), **Synchronizing** (`syncing` ), **Normal** (`normal` ), or **Stopped** (`stopped`) state.
 * *   If you set `OneShot` to `true`, the replication pair-consistent group must be in the **Created** (`created` ), **One-time Syncing** (`manual_syncing` ), or **Stopped** (`stopped`) state. The time interval between two consecutive one-time synchronizations must be longer than one half of the recovery point objective (RPO).
 * *   After a replication pair-consistent group is activated, the group enters the **Initial Syncing** (`initial_syncing`) state and the system performs the first async replication to replicate all data from the primary disks to secondary disks.
 *
 * @param request StartDiskReplicaGroupRequest
 * @return StartDiskReplicaGroupResponse
 */
func (client *Client) StartDiskReplicaGroup(request *StartDiskReplicaGroupRequest) (_result *StartDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDiskReplicaGroupResponse{}
	_body, _err := client.StartDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   Only replication pairs that are in the **Created** (`created`) or **Stopped** (`stopped`) state can be activated.
 * *   After a replication pair is activated, it enters the **Initial Syncing** (`initial_syncing`) state and the system performs the first asynchronous replication to replicate all data from the primary disk to the secondary disk.
 *
 * @param request StartDiskReplicaPairRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartDiskReplicaPairResponse
 */
func (client *Client) StartDiskReplicaPairWithOptions(request *StartDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *StartDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OneShot)) {
		query["OneShot"] = request.OneShot
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   Only replication pairs that are in the **Created** (`created`) or **Stopped** (`stopped`) state can be activated.
 * *   After a replication pair is activated, it enters the **Initial Syncing** (`initial_syncing`) state and the system performs the first asynchronous replication to replicate all data from the primary disk to the secondary disk.
 *
 * @param request StartDiskReplicaPairRequest
 * @return StartDiskReplicaPairResponse
 */
func (client *Client) StartDiskReplicaPair(request *StartDiskReplicaPairRequest) (_result *StartDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDiskReplicaPairResponse{}
	_body, _err := client.StartDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After the disaster recovery drill is complete on the secondary disk, a pay-as-you-go drill disk that has the same capacity and category as the secondary disk is created in the zone where the secondary disk resides. The drill disk contains last-recovery-point data that can be used to test the completeness and correctness of applications.
 *
 * @param request StartPairDrillRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartPairDrillResponse
 */
func (client *Client) StartPairDrillWithOptions(request *StartPairDrillRequest, runtime *util.RuntimeOptions) (_result *StartPairDrillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.PairId)) {
		query["PairId"] = request.PairId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartPairDrill"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartPairDrillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After the disaster recovery drill is complete on the secondary disk, a pay-as-you-go drill disk that has the same capacity and category as the secondary disk is created in the zone where the secondary disk resides. The drill disk contains last-recovery-point data that can be used to test the completeness and correctness of applications.
 *
 * @param request StartPairDrillRequest
 * @return StartPairDrillResponse
 */
func (client *Client) StartPairDrill(request *StartPairDrillRequest) (_result *StartPairDrillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartPairDrillResponse{}
	_body, _err := client.StartPairDrillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After the disaster recovery drill is complete on secondary disks, a pay-as-you-go drill disk is created in the zone where the secondary disk of each replication pair resides. The latest-recovery-point data is restored to the drill disks to test the completeness and correctness of applications.
 *
 * @param request StartReplicaGroupDrillRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartReplicaGroupDrillResponse
 */
func (client *Client) StartReplicaGroupDrillWithOptions(request *StartReplicaGroupDrillRequest, runtime *util.RuntimeOptions) (_result *StartReplicaGroupDrillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartReplicaGroupDrill"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartReplicaGroupDrillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After the disaster recovery drill is complete on secondary disks, a pay-as-you-go drill disk is created in the zone where the secondary disk of each replication pair resides. The latest-recovery-point data is restored to the drill disks to test the completeness and correctness of applications.
 *
 * @param request StartReplicaGroupDrillRequest
 * @return StartReplicaGroupDrillResponse
 */
func (client *Client) StartReplicaGroupDrill(request *StartReplicaGroupDrillRequest) (_result *StartReplicaGroupDrillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartReplicaGroupDrillResponse{}
	_body, _err := client.StartReplicaGroupDrillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   The replication pair-consistent group that you want to stop must be in the **One-time Syncing** (`manual_syncing`), **Syncing** (`syncing`), **Normal** (`normal`), **Stopping** (`stopping`), **Stop Failed** (`stop_failed`), or **Stopped** (`stopped`) state.
 * *   When a replication pair-consistent group is stopped, it enters the **Stopped** (`stopped`) state. If a replication pair-consistent group cannot be stopped, the state of the group remains unchanged or changes to **Stop Failed** (`stop_failed`). In this case, try again later.
 *
 * @param request StopDiskReplicaGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopDiskReplicaGroupResponse
 */
func (client *Client) StopDiskReplicaGroupWithOptions(request *StopDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *StopDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which the replication pair-consistent group feature is available, see [Overview](~~314563~~).
 * *   The replication pair-consistent group that you want to stop must be in the **One-time Syncing** (`manual_syncing`), **Syncing** (`syncing`), **Normal** (`normal`), **Stopping** (`stopping`), **Stop Failed** (`stop_failed`), or **Stopped** (`stopped`) state.
 * *   When a replication pair-consistent group is stopped, it enters the **Stopped** (`stopped`) state. If a replication pair-consistent group cannot be stopped, the state of the group remains unchanged or changes to **Stop Failed** (`stop_failed`). In this case, try again later.
 *
 * @param request StopDiskReplicaGroupRequest
 * @return StopDiskReplicaGroupResponse
 */
func (client *Client) StopDiskReplicaGroup(request *StopDiskReplicaGroupRequest) (_result *StopDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDiskReplicaGroupResponse{}
	_body, _err := client.StopDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   Only replication pairs that are in the **Initial Syncing** (`initial_syncing`), **Syncing** (`syncing`), **One-time Syncing** (`manual_syncing`), or **Normal** (`normal`) state can be stopped. When a replication pair is stopped, it enters the Stopped (`stopped`) state. The secondary disk rolls back to the point in time when the last async replication was complete and drops all the data that is being replicated from the primary disk.
 *
 * @param request StopDiskReplicaPairRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopDiskReplicaPairResponse
 */
func (client *Client) StopDiskReplicaPairWithOptions(request *StopDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *StopDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## [](#)Usage notes
 * *   For information about the regions in which async replication is available, see [Overview](~~314563~~).
 * *   Only replication pairs that are in the **Initial Syncing** (`initial_syncing`), **Syncing** (`syncing`), **One-time Syncing** (`manual_syncing`), or **Normal** (`normal`) state can be stopped. When a replication pair is stopped, it enters the Stopped (`stopped`) state. The secondary disk rolls back to the point in time when the last async replication was complete and drops all the data that is being replicated from the primary disk.
 *
 * @param request StopDiskReplicaPairRequest
 * @return StopDiskReplicaPairResponse
 */
func (client *Client) StopDiskReplicaPair(request *StopDiskReplicaPairRequest) (_result *StopDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDiskReplicaPairResponse{}
	_body, _err := client.StopDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you add tags to a resource, Alibaba Cloud checks the number of existing tags of the resource. If the maximum number of tags is reached, an error message is returned. For more information, see the "Tag limits" section in [Limits](~~25412~~).
 *
 * @param request TagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return TagResourcesResponse
 */
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you add tags to a resource, Alibaba Cloud checks the number of existing tags of the resource. If the maximum number of tags is reached, an error message is returned. For more information, see the "Tag limits" section in [Limits](~~25412~~).
 *
 * @param request TagResourcesRequest
 * @return TagResourcesResponse
 */
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindEnterpriseSnapshotPolicyWithOptions(request *UnbindEnterpriseSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *UnbindEnterpriseSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DiskTargets)) {
		query["DiskTargets"] = request.DiskTargets
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindEnterpriseSnapshotPolicy"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindEnterpriseSnapshotPolicy(request *UnbindEnterpriseSnapshotPolicyRequest) (_result *UnbindEnterpriseSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.UnbindEnterpriseSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   You can remove up to 20 tags at a time.
 * *   After a tag is removed from an EBS resource, the tag is automatically deleted if the tag is not added to any instance.
 *
 * @param request UntagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UntagResourcesResponse
 */
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   You can remove up to 20 tags at a time.
 * *   After a tag is removed from an EBS resource, the tag is automatically deleted if the tag is not added to any instance.
 *
 * @param request UntagResourcesRequest
 * @return UntagResourcesResponse
 */
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEnterpriseSnapshotPolicyWithOptions(tmpReq *UpdateEnterpriseSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateEnterpriseSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateEnterpriseSnapshotPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CrossRegionCopyInfo)) {
		request.CrossRegionCopyInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CrossRegionCopyInfo, tea.String("CrossRegionCopyInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RetainRule)) {
		request.RetainRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RetainRule, tea.String("RetainRule"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Schedule)) {
		request.ScheduleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Schedule, tea.String("Schedule"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SpecialRetainRules)) {
		request.SpecialRetainRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SpecialRetainRules, tea.String("SpecialRetainRules"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StorageRule)) {
		request.StorageRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StorageRule, tea.String("StorageRule"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CrossRegionCopyInfoShrink)) {
		query["CrossRegionCopyInfo"] = request.CrossRegionCopyInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		query["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RetainRuleShrink)) {
		query["RetainRule"] = request.RetainRuleShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleShrink)) {
		query["Schedule"] = request.ScheduleShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SpecialRetainRulesShrink)) {
		query["SpecialRetainRules"] = request.SpecialRetainRulesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.StorageRuleShrink)) {
		query["StorageRule"] = request.StorageRuleShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEnterpriseSnapshotPolicy"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEnterpriseSnapshotPolicy(request *UpdateEnterpriseSnapshotPolicyRequest) (_result *UpdateEnterpriseSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEnterpriseSnapshotPolicyResponse{}
	_body, _err := client.UpdateEnterpriseSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSolutionInstanceAttributeWithOptions(request *UpdateSolutionInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateSolutionInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionInstanceId)) {
		query["SolutionInstanceId"] = request.SolutionInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSolutionInstanceAttribute"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSolutionInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSolutionInstanceAttribute(request *UpdateSolutionInstanceAttributeRequest) (_result *UpdateSolutionInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSolutionInstanceAttributeResponse{}
	_body, _err := client.UpdateSolutionInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
