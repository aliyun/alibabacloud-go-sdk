// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBClustersResponseBodyItems) *DescribeDBClustersResponseBody
	GetItems() *DescribeDBClustersResponseBodyItems
	SetPageNumber(v int32) *DescribeDBClustersResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDBClustersResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDBClustersResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDBClustersResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBClustersResponseBody struct {
	Items *DescribeDBClustersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The number of the page to return.
	//
	// example:
	//
	// 12
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of clusters returned per page.
	//
	// example:
	//
	// 5
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 16
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBody) GetItems() *DescribeDBClustersResponseBodyItems {
	return s.Items
}

func (s *DescribeDBClustersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClustersResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDBClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClustersResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBClustersResponseBody) SetItems(v *DescribeDBClustersResponseBodyItems) *DescribeDBClustersResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageNumber(v int32) *DescribeDBClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageRecordCount(v int32) *DescribeDBClustersResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetRequestId(v string) *DescribeDBClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetTotalRecordCount(v int32) *DescribeDBClustersResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBClustersResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClustersResponseBodyItems struct {
	DBCluster []*DescribeDBClustersResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItems) GetDBCluster() []*DescribeDBClustersResponseBodyItemsDBCluster {
	return s.DBCluster
}

func (s *DescribeDBClustersResponseBodyItems) SetDBCluster(v []*DescribeDBClustersResponseBodyItemsDBCluster) *DescribeDBClustersResponseBodyItems {
	s.DBCluster = v
	return s
}

func (s *DescribeDBClustersResponseBodyItems) Validate() error {
	if s.DBCluster != nil {
		for _, item := range s.DBCluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClustersResponseBodyItemsDBCluster struct {
	AiType               *string                                              `json:"AiType,omitempty" xml:"AiType,omitempty"`
	Category             *string                                              `json:"Category,omitempty" xml:"Category,omitempty"`
	CpuCores             *string                                              `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	CreateTime           *string                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBClusterDescription *string                                              `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterId          *string                                              `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterNetworkType *string                                              `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterStatus      *string                                              `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	DBNodeClass          *string                                              `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeNumber         *int32                                               `json:"DBNodeNumber,omitempty" xml:"DBNodeNumber,omitempty"`
	DBNodes              *DescribeDBClustersResponseBodyItemsDBClusterDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Struct"`
	DBType               *string                                              `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion            *string                                              `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DeletionLock         *int32                                               `json:"DeletionLock,omitempty" xml:"DeletionLock,omitempty"`
	Engine               *string                                              `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ExpireTime           *string                                              `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired              *string                                              `json:"Expired,omitempty" xml:"Expired,omitempty"`
	HotStandbyCluster    *string                                              `json:"HotStandbyCluster,omitempty" xml:"HotStandbyCluster,omitempty"`
	LockMode             *string                                              `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MemorySize           *string                                              `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	PayType              *string                                              `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId             *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RemoteMemorySize     *string                                              `json:"RemoteMemorySize,omitempty" xml:"RemoteMemorySize,omitempty"`
	ResourceGroupId      *string                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SearchStorageUsed    *int64                                               `json:"SearchStorageUsed,omitempty" xml:"SearchStorageUsed,omitempty"`
	ServerlessType       *string                                              `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	StoragePayType       *string                                              `json:"StoragePayType,omitempty" xml:"StoragePayType,omitempty"`
	StorageSpace         *int64                                               `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	StorageType          *string                                              `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	StorageUsed          *int64                                               `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	StrictConsistency    *string                                              `json:"StrictConsistency,omitempty" xml:"StrictConsistency,omitempty"`
	SubCategory          *string                                              `json:"SubCategory,omitempty" xml:"SubCategory,omitempty"`
	Tags                 *DescribeDBClustersResponseBodyItemsDBClusterTags    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcId                *string                                              `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId            *string                                              `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ZoneId               *string                                              `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetAiType() *string {
	return s.AiType
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetCpuCores() *string {
	return s.CpuCores
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBClusterNetworkType() *string {
	return s.DBClusterNetworkType
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBNodeNumber() *int32 {
	return s.DBNodeNumber
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBNodes() *DescribeDBClustersResponseBodyItemsDBClusterDBNodes {
	return s.DBNodes
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDeletionLock() *int32 {
	return s.DeletionLock
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetHotStandbyCluster() *string {
	return s.HotStandbyCluster
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetMemorySize() *string {
	return s.MemorySize
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetRemoteMemorySize() *string {
	return s.RemoteMemorySize
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetSearchStorageUsed() *int64 {
	return s.SearchStorageUsed
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetStoragePayType() *string {
	return s.StoragePayType
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetStorageSpace() *int64 {
	return s.StorageSpace
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetStorageUsed() *int64 {
	return s.StorageUsed
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetStrictConsistency() *string {
	return s.StrictConsistency
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetSubCategory() *string {
	return s.SubCategory
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetTags() *DescribeDBClustersResponseBodyItemsDBClusterTags {
	return s.Tags
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetAiType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.AiType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCategory(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCpuCores(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCreateTime(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeNumber(v int32) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeNumber = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodes(v *DescribeDBClustersResponseBodyItemsDBClusterDBNodes) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodes = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDeletionLock(v int32) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DeletionLock = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetHotStandbyCluster(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.HotStandbyCluster = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetMemorySize(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.MemorySize = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetRemoteMemorySize(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.RemoteMemorySize = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetSearchStorageUsed(v int64) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.SearchStorageUsed = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetServerlessType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ServerlessType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetStoragePayType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.StoragePayType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetStorageSpace(v int64) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.StorageSpace = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetStorageType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.StorageType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetStorageUsed(v int64) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetStrictConsistency(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.StrictConsistency = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetSubCategory(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.SubCategory = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetTags(v *DescribeDBClustersResponseBodyItemsDBClusterTags) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVpcId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVswitchId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VswitchId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) Validate() error {
	if s.DBNodes != nil {
		if err := s.DBNodes.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClustersResponseBodyItemsDBClusterDBNodes struct {
	DBNode []*DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterDBNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodes) GetDBNode() []*DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	return s.DBNode
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodes) SetDBNode(v []*DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) *DescribeDBClustersResponseBodyItemsDBClusterDBNodes {
	s.DBNode = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodes) Validate() error {
	if s.DBNode != nil {
		for _, item := range s.DBNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode struct {
	DBNodeClass    *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeId       *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	DBNodeRole     *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	HotReplicaMode *string `json:"HotReplicaMode,omitempty" xml:"HotReplicaMode,omitempty"`
	ImciSwitch     *string `json:"ImciSwitch,omitempty" xml:"ImciSwitch,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Serverless     *string `json:"Serverless,omitempty" xml:"Serverless,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) GetDBNodeRole() *string {
	return s.DBNodeRole
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) GetHotReplicaMode() *string {
	return s.HotReplicaMode
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) GetImciSwitch() *string {
	return s.ImciSwitch
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) GetServerless() *string {
	return s.Serverless
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetDBNodeClass(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetDBNodeId(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetDBNodeRole(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetHotReplicaMode(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.HotReplicaMode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetImciSwitch(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.ImciSwitch = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetRegionId(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetServerless(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.Serverless = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetZoneId(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClustersResponseBodyItemsDBClusterTags struct {
	Tag []*DescribeDBClustersResponseBodyItemsDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTags) GetTag() []*DescribeDBClustersResponseBodyItemsDBClusterTagsTag {
	return s.Tag
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTags) SetTag(v []*DescribeDBClustersResponseBodyItemsDBClusterTagsTag) *DescribeDBClustersResponseBodyItemsDBClusterTags {
	s.Tag = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClustersResponseBodyItemsDBClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) SetKey(v string) *DescribeDBClustersResponseBodyItemsDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) SetValue(v string) *DescribeDBClustersResponseBodyItemsDBClusterTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) Validate() error {
	return dara.Validate(s)
}
