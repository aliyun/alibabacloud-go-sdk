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
	SetPageSize(v int32) *DescribeDBClustersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDBClustersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDBClustersResponseBody
	GetTotalCount() *int32
}

type DescribeDBClustersResponseBody struct {
	Items *DescribeDBClustersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D65A809F-34CE-4550-9BC1-0ED21ETG380
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeDBClustersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClustersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDBClustersResponseBody) SetItems(v *DescribeDBClustersResponseBodyItems) *DescribeDBClustersResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageNumber(v int32) *DescribeDBClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageSize(v int32) *DescribeDBClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetRequestId(v string) *DescribeDBClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetTotalCount(v int32) *DescribeDBClustersResponseBody {
	s.TotalCount = &v
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
	Category             *string                                               `json:"Category,omitempty" xml:"Category,omitempty"`
	CommodityCode        *string                                               `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ComputeResource      *string                                               `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	ConnectionString     *string                                               `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	CreateTime           *string                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBClusterDescription *string                                               `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterId          *string                                               `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterNetworkType *string                                               `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterStatus      *string                                               `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	DBClusterType        *string                                               `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	DBNodeClass          *string                                               `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeCount          *int64                                                `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodeStorage        *int64                                                `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	DBVersion            *string                                               `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DiskType             *string                                               `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DtsJobId             *string                                               `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ElasticIOResource    *int32                                                `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	Engine               *string                                               `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ExecutorCount        *string                                               `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	ExpireTime           *string                                               `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired              *string                                               `json:"Expired,omitempty" xml:"Expired,omitempty"`
	InnerIp              *string                                               `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	InnerPort            *string                                               `json:"InnerPort,omitempty" xml:"InnerPort,omitempty"`
	LockMode             *string                                               `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason           *string                                               `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	Mode                 *string                                               `json:"Mode,omitempty" xml:"Mode,omitempty"`
	PayType              *string                                               `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                 *string                                               `json:"Port,omitempty" xml:"Port,omitempty"`
	ProductVersion       *string                                               `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	RdsInstanceId        *string                                               `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	RegionId             *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StorageResource      *string                                               `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	Tags                 *DescribeDBClustersResponseBodyItemsDBClusterTags     `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TaskInfo             *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
	VPCCloudInstanceId   *string                                               `json:"VPCCloudInstanceId,omitempty" xml:"VPCCloudInstanceId,omitempty"`
	VPCId                *string                                               `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string                                               `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId               *string                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetConnectionString() *string {
	return s.ConnectionString
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

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBClusterType() *string {
	return s.DBClusterType
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBNodeCount() *int64 {
	return s.DBNodeCount
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBNodeStorage() *int64 {
	return s.DBNodeStorage
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetElasticIOResource() *int32 {
	return s.ElasticIOResource
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetExecutorCount() *string {
	return s.ExecutorCount
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetInnerIp() *string {
	return s.InnerIp
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetInnerPort() *string {
	return s.InnerPort
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetMode() *string {
	return s.Mode
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetPort() *string {
	return s.Port
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetRdsInstanceId() *string {
	return s.RdsInstanceId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetStorageResource() *string {
	return s.StorageResource
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetTags() *DescribeDBClustersResponseBodyItemsDBClusterTags {
	return s.Tags
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetTaskInfo() *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo {
	return s.TaskInfo
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetVPCCloudInstanceId() *string {
	return s.VPCCloudInstanceId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCategory(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCommodityCode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetComputeResource(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ComputeResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetConnectionString(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ConnectionString = &v
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

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeCount(v int64) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeStorage(v int64) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDiskType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DiskType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDtsJobId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetElasticIOResource(v int32) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ElasticIOResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExecutorCount(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ExecutorCount = &v
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

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetInnerIp(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.InnerIp = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetInnerPort(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.InnerPort = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetLockReason(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetMode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Mode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetPort(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetProductVersion(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetRdsInstanceId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetStorageResource(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.StorageResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetTags(v *DescribeDBClustersResponseBodyItemsDBClusterTags) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetTaskInfo(v *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.TaskInfo = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVPCCloudInstanceId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VPCCloudInstanceId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVPCId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVSwitchId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.TaskInfo != nil {
		if err := s.TaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type DescribeDBClustersResponseBodyItemsDBClusterTaskInfo struct {
	Name     *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Progress *string                                                       `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status   *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	StepList *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepList `json:"StepList,omitempty" xml:"StepList,omitempty" type:"Struct"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo) GetName() *string {
	return s.Name
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo) GetStepList() *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepList {
	return s.StepList
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo) SetName(v string) *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo {
	s.Name = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo) SetProgress(v string) *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo {
	s.Progress = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo) SetStatus(v string) *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo {
	s.Status = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo) SetStepList(v *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepList) *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo {
	s.StepList = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo) Validate() error {
	if s.StepList != nil {
		if err := s.StepList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepList struct {
	StepList []*DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList `json:"StepList,omitempty" xml:"StepList,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepList) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepList) GetStepList() []*DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList {
	return s.StepList
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepList) SetStepList(v []*DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepList {
	s.StepList = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepList) Validate() error {
	if s.StepList != nil {
		for _, item := range s.StepList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StepDesc     *string `json:"StepDesc,omitempty" xml:"StepDesc,omitempty"`
	StepName     *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	StepProgress *string `json:"StepProgress,omitempty" xml:"StepProgress,omitempty"`
	StepStatus   *string `json:"StepStatus,omitempty" xml:"StepStatus,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) GetStepDesc() *string {
	return s.StepDesc
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) GetStepName() *string {
	return s.StepName
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) GetStepProgress() *string {
	return s.StepProgress
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) GetStepStatus() *string {
	return s.StepStatus
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) SetEndTime(v string) *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) SetStartTime(v string) *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) SetStepDesc(v string) *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList {
	s.StepDesc = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) SetStepName(v string) *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList {
	s.StepName = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) SetStepProgress(v string) *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList {
	s.StepProgress = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) SetStepStatus(v string) *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList {
	s.StepStatus = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList) Validate() error {
	return dara.Validate(s)
}
