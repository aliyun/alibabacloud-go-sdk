// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateCacheAnalysisTaskRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateCacheAnalysisTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisTaskRequest) SetSecurityToken(v string) *CreateCacheAnalysisTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetOwnerId(v int64) *CreateCacheAnalysisTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetResourceOwnerAccount(v string) *CreateCacheAnalysisTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetResourceOwnerId(v int64) *CreateCacheAnalysisTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetOwnerAccount(v string) *CreateCacheAnalysisTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetInstanceId(v string) *CreateCacheAnalysisTaskRequest {
	s.InstanceId = &v
	return s
}

type CreateCacheAnalysisTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCacheAnalysisTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisTaskResponseBody) SetRequestId(v string) *CreateCacheAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateCacheAnalysisTaskResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCacheAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCacheAnalysisTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisTaskResponse) SetHeaders(v map[string]*string) *CreateCacheAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCacheAnalysisTaskResponse) SetBody(v *CreateCacheAnalysisTaskResponseBody) *CreateCacheAnalysisTaskResponse {
	s.Body = v
	return s
}

type DescribeDedicatedClusterInstanceListRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus       *int32  `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceNetType      *string `json:"InstanceNetType,omitempty" xml:"InstanceNetType,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DedicatedHostName    *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetSecurityToken(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetOwnerId(v int64) *DescribeDedicatedClusterInstanceListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetResourceOwnerId(v int64) *DescribeDedicatedClusterInstanceListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetOwnerAccount(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetRegionId(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetInstanceId(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetInstanceStatus(v int32) *DescribeDedicatedClusterInstanceListRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetInstanceNetType(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.InstanceNetType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetEngine(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetEngineVersion(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetClusterId(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetDedicatedHostName(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetPageNumber(v int32) *DescribeDedicatedClusterInstanceListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetPageSize(v int32) *DescribeDedicatedClusterInstanceListRequest {
	s.PageSize = &v
	return s
}

type DescribeDedicatedClusterInstanceListResponseBody struct {
	Instances  []*DescribeDedicatedClusterInstanceListResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	TotalCount *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetInstances(v []*DescribeDedicatedClusterInstanceListResponseBodyInstances) *DescribeDedicatedClusterInstanceListResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetTotalCount(v int32) *DescribeDedicatedClusterInstanceListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetPageSize(v int32) *DescribeDedicatedClusterInstanceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetRequestId(v string) *DescribeDedicatedClusterInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetPageNumber(v int32) *DescribeDedicatedClusterInstanceListResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeDedicatedClusterInstanceListResponseBodyInstances struct {
	VpcId             *string                                                                      `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	CharacterType     *int32                                                                       `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	VswitchId         *string                                                                      `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	MaintainStartTime *string                                                                      `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	InstanceClass     *string                                                                      `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	ConnectionDomain  *string                                                                      `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	CreateTime        *string                                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MaintainEndTime   *string                                                                      `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	StorageType       *string                                                                      `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	InstanceNodeList  []*DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList `json:"InstanceNodeList,omitempty" xml:"InstanceNodeList,omitempty" type:"Repeated"`
	InstanceId        *string                                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	EngineVersion     *string                                                                      `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RegionId          *string                                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceName      *string                                                                      `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ZoneId            *string                                                                      `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClusterName       *string                                                                      `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	InstanceStatus    *string                                                                      `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	Engine            *string                                                                      `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ShardCount        *int32                                                                       `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	CustomId          *string                                                                      `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	ClusterId         *string                                                                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetVpcId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetCharacterType(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.CharacterType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetVswitchId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.VswitchId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetMaintainStartTime(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceClass(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceClass = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetConnectionDomain(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetCreateTime(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetMaintainEndTime(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetStorageType(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.StorageType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceNodeList(v []*DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceNodeList = v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetEngineVersion(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetRegionId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetZoneId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetClusterName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ClusterName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceStatus(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetEngine(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetShardCount(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ShardCount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetCustomId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.CustomId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetClusterId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ClusterId = &v
	return s
}

type DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList struct {
	NodeIp            *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	NodeType          *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	ZoneId            *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Port              *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Role              *string `json:"Role,omitempty" xml:"Role,omitempty"`
	NodeId            *int32  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetNodeIp(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.NodeIp = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetDedicatedHostName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetNodeType(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.NodeType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetZoneId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetInstanceId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetPort(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.Port = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetRole(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.Role = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetNodeId(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.NodeId = &v
	return s
}

type DescribeDedicatedClusterInstanceListResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedClusterInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedClusterInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponse) SetHeaders(v map[string]*string) *DescribeDedicatedClusterInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponse) SetBody(v *DescribeDedicatedClusterInstanceListResponseBody) *DescribeDedicatedClusterInstanceListResponse {
	s.Body = v
	return s
}

type DescribeRoleZoneInfoRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Product              *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	QueryType            *int32  `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRoleZoneInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoRequest) SetSecurityToken(v string) *DescribeRoleZoneInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetOwnerId(v int64) *DescribeRoleZoneInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetResourceOwnerAccount(v string) *DescribeRoleZoneInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetResourceOwnerId(v int64) *DescribeRoleZoneInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetOwnerAccount(v string) *DescribeRoleZoneInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetInstanceId(v string) *DescribeRoleZoneInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetProduct(v string) *DescribeRoleZoneInfoRequest {
	s.Product = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetCategory(v string) *DescribeRoleZoneInfoRequest {
	s.Category = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetQueryType(v int32) *DescribeRoleZoneInfoRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetPageNumber(v int32) *DescribeRoleZoneInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetPageSize(v int32) *DescribeRoleZoneInfoRequest {
	s.PageSize = &v
	return s
}

type DescribeRoleZoneInfoResponseBody struct {
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Node       *DescribeRoleZoneInfoResponseBodyNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
}

func (s DescribeRoleZoneInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBody) SetTotalCount(v int32) *DescribeRoleZoneInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetRequestId(v string) *DescribeRoleZoneInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetPageSize(v int32) *DescribeRoleZoneInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetPageNumber(v int32) *DescribeRoleZoneInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetNode(v *DescribeRoleZoneInfoResponseBodyNode) *DescribeRoleZoneInfoResponseBody {
	s.Node = v
	return s
}

type DescribeRoleZoneInfoResponseBodyNode struct {
	NodeInfo []*DescribeRoleZoneInfoResponseBodyNodeNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s DescribeRoleZoneInfoResponseBodyNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBodyNode) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBodyNode) SetNodeInfo(v []*DescribeRoleZoneInfoResponseBodyNodeNodeInfo) *DescribeRoleZoneInfoResponseBodyNode {
	s.NodeInfo = v
	return s
}

type DescribeRoleZoneInfoResponseBodyNodeNodeInfo struct {
	CurrentMinorVersion *string `json:"CurrentMinorVersion,omitempty" xml:"CurrentMinorVersion,omitempty"`
	InsType             *int32  `json:"InsType,omitempty" xml:"InsType,omitempty"`
	IsLatestVersion     *int32  `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	InsName             *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	NodeType            *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	ZoneId              *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Role                *string `json:"Role,omitempty" xml:"Role,omitempty"`
	CustinsId           *string `json:"CustinsId,omitempty" xml:"CustinsId,omitempty"`
	NodeId              *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeRoleZoneInfoResponseBodyNodeNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetCurrentMinorVersion(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.CurrentMinorVersion = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetInsType(v int32) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.InsType = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetIsLatestVersion(v int32) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetInsName(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.InsName = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetNodeType(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.NodeType = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetZoneId(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetRole(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.Role = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetCustinsId(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.CustinsId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetNodeId(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.NodeId = &v
	return s
}

type DescribeRoleZoneInfoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRoleZoneInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRoleZoneInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponse) SetHeaders(v map[string]*string) *DescribeRoleZoneInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoleZoneInfoResponse) SetBody(v *DescribeRoleZoneInfoResponseBody) *DescribeRoleZoneInfoResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":                  tea.String("r-kvstore.aliyuncs.com"),
		"cn-beijing":                  tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai":                 tea.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("r-kvstore.aliyuncs.com"),
		"cn-heyuan":                   tea.String("r-kvstore.aliyuncs.com"),
		"ap-southeast-1":              tea.String("r-kvstore.aliyuncs.com"),
		"us-west-1":                   tea.String("r-kvstore.aliyuncs.com"),
		"us-east-1":                   tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("r-kvstore.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("r-kvstore.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("r-kvstore.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("r-kvstore.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("r-kvstore.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("r-kvstore.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("r-kvstore.aliyuncs.com"),
		"cn-edge-1":                   tea.String("r-kvstore.aliyuncs.com"),
		"cn-fujian":                   tea.String("r-kvstore.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("r-kvstore.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("r-kvstore.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("r-kvstore.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("r-kvstore.aliyuncs.com"),
		"cn-wuhan":                    tea.String("r-kvstore.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("r-kvstore.aliyuncs.com"),
		"cn-yushanfang":               tea.String("r-kvstore.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("r-kvstore.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("r-kvstore.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("r-kvstore.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("r-kvstore.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("r-kvstore.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("r-kvstore"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateCacheAnalysisTaskWithOptions(request *CreateCacheAnalysisTaskRequest, runtime *util.RuntimeOptions) (_result *CreateCacheAnalysisTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCacheAnalysisTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCacheAnalysisTask"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCacheAnalysisTask(request *CreateCacheAnalysisTaskRequest) (_result *CreateCacheAnalysisTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCacheAnalysisTaskResponse{}
	_body, _err := client.CreateCacheAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedClusterInstanceListWithOptions(request *DescribeDedicatedClusterInstanceListRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedClusterInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedClusterInstanceListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedClusterInstanceList"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedClusterInstanceList(request *DescribeDedicatedClusterInstanceListRequest) (_result *DescribeDedicatedClusterInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedClusterInstanceListResponse{}
	_body, _err := client.DescribeDedicatedClusterInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRoleZoneInfoWithOptions(request *DescribeRoleZoneInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeRoleZoneInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRoleZoneInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRoleZoneInfo"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRoleZoneInfo(request *DescribeRoleZoneInfoRequest) (_result *DescribeRoleZoneInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRoleZoneInfoResponse{}
	_body, _err := client.DescribeRoleZoneInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
