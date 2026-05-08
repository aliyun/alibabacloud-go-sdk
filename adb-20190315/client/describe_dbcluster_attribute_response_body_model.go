// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBClusterAttributeResponseBodyItems) *DescribeDBClusterAttributeResponseBody
	GetItems() *DescribeDBClusterAttributeResponseBodyItems
	SetRequestId(v string) *DescribeDBClusterAttributeResponseBody
	GetRequestId() *string
}

type DescribeDBClusterAttributeResponseBody struct {
	Items *DescribeDBClusterAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2895BB82-B2C1-408E-AA73-DB8D59******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBody) GetItems() *DescribeDBClusterAttributeResponseBodyItems {
	return s.Items
}

func (s *DescribeDBClusterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterAttributeResponseBody) SetItems(v *DescribeDBClusterAttributeResponseBodyItems) *DescribeDBClusterAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRequestId(v string) *DescribeDBClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterAttributeResponseBodyItems struct {
	DBCluster []*DescribeDBClusterAttributeResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAttributeResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItems) GetDBCluster() []*DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	return s.DBCluster
}

func (s *DescribeDBClusterAttributeResponseBodyItems) SetDBCluster(v []*DescribeDBClusterAttributeResponseBodyItemsDBCluster) *DescribeDBClusterAttributeResponseBodyItems {
	s.DBCluster = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItems) Validate() error {
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

type DescribeDBClusterAttributeResponseBodyItemsDBCluster struct {
	Category              *string                                                       `json:"Category,omitempty" xml:"Category,omitempty"`
	CommodityCode         *string                                                       `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ComputeResource       *string                                                       `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	ConnectionString      *string                                                       `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	CreationTime          *string                                                       `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DBClusterDescription  *string                                                       `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterId           *string                                                       `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterNetworkType  *string                                                       `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterStatus       *string                                                       `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	DBClusterType         *string                                                       `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	DBNodeClass           *string                                                       `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeCount           *int64                                                        `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodeStorage         *int64                                                        `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	DBVersion             *string                                                       `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DiskEncryption        *bool                                                         `json:"DiskEncryption,omitempty" xml:"DiskEncryption,omitempty"`
	DiskPerformanceLevel  *string                                                       `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	DiskType              *string                                                       `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DtsJobId              *string                                                       `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ElasticIOResource     *int32                                                        `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	ElasticIOResourceSize *string                                                       `json:"ElasticIOResourceSize,omitempty" xml:"ElasticIOResourceSize,omitempty"`
	EnableAirflow         *bool                                                         `json:"EnableAirflow,omitempty" xml:"EnableAirflow,omitempty"`
	EnableSpark           *bool                                                         `json:"EnableSpark,omitempty" xml:"EnableSpark,omitempty"`
	Engine                *string                                                       `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string                                                       `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExecutorCount         *string                                                       `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	ExpireTime            *string                                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired               *string                                                       `json:"Expired,omitempty" xml:"Expired,omitempty"`
	InnerIp               *string                                                       `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	InnerPort             *string                                                       `json:"InnerPort,omitempty" xml:"InnerPort,omitempty"`
	KmsId                 *string                                                       `json:"KmsId,omitempty" xml:"KmsId,omitempty"`
	LockMode              *string                                                       `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason            *string                                                       `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainTime          *string                                                       `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	Mode                  *string                                                       `json:"Mode,omitempty" xml:"Mode,omitempty"`
	PayType               *string                                                       `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                  *int32                                                        `json:"Port,omitempty" xml:"Port,omitempty"`
	ProductForm           *string                                                       `json:"ProductForm,omitempty" xml:"ProductForm,omitempty"`
	ProductVersion        *string                                                       `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	RdsInstanceId         *string                                                       `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	RegionId              *string                                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReservedNodeCount     *int32                                                        `json:"ReservedNodeCount,omitempty" xml:"ReservedNodeCount,omitempty"`
	ReservedNodeSize      *string                                                       `json:"ReservedNodeSize,omitempty" xml:"ReservedNodeSize,omitempty"`
	ResourceGroupId       *string                                                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecondaryVSwitchId    *string                                                       `json:"SecondaryVSwitchId,omitempty" xml:"SecondaryVSwitchId,omitempty"`
	SecondaryZoneId       *string                                                       `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	StorageResource       *string                                                       `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	Tags                  *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags     `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TaskInfo              *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
	UserENIStatus         *bool                                                         `json:"UserENIStatus,omitempty" xml:"UserENIStatus,omitempty"`
	VPCCloudInstanceId    *string                                                       `json:"VPCCloudInstanceId,omitempty" xml:"VPCCloudInstanceId,omitempty"`
	VPCId                 *string                                                       `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId             *string                                                       `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                *string                                                       `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDBClusterNetworkType() *string {
	return s.DBClusterNetworkType
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDBClusterType() *string {
	return s.DBClusterType
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDBNodeCount() *int64 {
	return s.DBNodeCount
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDBNodeStorage() *int64 {
	return s.DBNodeStorage
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDiskPerformanceLevel() *string {
	return s.DiskPerformanceLevel
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetElasticIOResource() *int32 {
	return s.ElasticIOResource
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetElasticIOResourceSize() *string {
	return s.ElasticIOResourceSize
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetEnableAirflow() *bool {
	return s.EnableAirflow
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetEnableSpark() *bool {
	return s.EnableSpark
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetExecutorCount() *string {
	return s.ExecutorCount
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetInnerIp() *string {
	return s.InnerIp
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetInnerPort() *string {
	return s.InnerPort
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetKmsId() *string {
	return s.KmsId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetMaintainTime() *string {
	return s.MaintainTime
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetMode() *string {
	return s.Mode
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetProductForm() *string {
	return s.ProductForm
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetRdsInstanceId() *string {
	return s.RdsInstanceId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetReservedNodeCount() *int32 {
	return s.ReservedNodeCount
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetReservedNodeSize() *string {
	return s.ReservedNodeSize
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetSecondaryVSwitchId() *string {
	return s.SecondaryVSwitchId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetStorageResource() *string {
	return s.StorageResource
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetTags() *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags {
	return s.Tags
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetTaskInfo() *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo {
	return s.TaskInfo
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetUserENIStatus() *bool {
	return s.UserENIStatus
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetVPCCloudInstanceId() *string {
	return s.VPCCloudInstanceId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCategory(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCommodityCode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetComputeResource(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ComputeResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetConnectionString(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCreationTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBNodeCount(v int64) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBNodeStorage(v int64) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDiskEncryption(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DiskEncryption = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDiskPerformanceLevel(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DiskPerformanceLevel = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDiskType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DiskType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDtsJobId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetElasticIOResource(v int32) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ElasticIOResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetElasticIOResourceSize(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ElasticIOResourceSize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEnableAirflow(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EnableAirflow = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEnableSpark(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EnableSpark = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEngineVersion(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExecutorCount(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ExecutorCount = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetInnerIp(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.InnerIp = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetInnerPort(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.InnerPort = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetKmsId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.KmsId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetLockReason(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetMaintainTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetMode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Mode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetPort(v int32) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetProductForm(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ProductForm = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetProductVersion(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetRdsInstanceId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetReservedNodeCount(v int32) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ReservedNodeCount = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetReservedNodeSize(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ReservedNodeSize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetSecondaryVSwitchId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.SecondaryVSwitchId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetSecondaryZoneId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.SecondaryZoneId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetStorageResource(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.StorageResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetTags(v *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetTaskInfo(v *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.TaskInfo = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetUserENIStatus(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.UserENIStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVPCCloudInstanceId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VPCCloudInstanceId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVPCId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVSwitchId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) Validate() error {
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

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTags struct {
	Tag []*DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) GetTag() []*DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag {
	return s.Tag
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) SetTag(v []*DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags {
	s.Tag = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) Validate() error {
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

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) SetKey(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) SetValue(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo struct {
	Name     *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Progress *string                                                               `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status   *string                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	StepList *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepList `json:"StepList,omitempty" xml:"StepList,omitempty" type:"Struct"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo) GetName() *string {
	return s.Name
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo) GetStepList() *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepList {
	return s.StepList
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo) SetName(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo) SetProgress(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo {
	s.Progress = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo) SetStatus(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo {
	s.Status = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo) SetStepList(v *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepList) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo {
	s.StepList = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo) Validate() error {
	if s.StepList != nil {
		if err := s.StepList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepList struct {
	StepList []*DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList `json:"StepList,omitempty" xml:"StepList,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepList) GetStepList() []*DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList {
	return s.StepList
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepList) SetStepList(v []*DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepList {
	s.StepList = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepList) Validate() error {
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

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StepDesc     *string `json:"StepDesc,omitempty" xml:"StepDesc,omitempty"`
	StepName     *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	StepProgress *string `json:"StepProgress,omitempty" xml:"StepProgress,omitempty"`
	StepStatus   *string `json:"StepStatus,omitempty" xml:"StepStatus,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) GetStepDesc() *string {
	return s.StepDesc
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) GetStepName() *string {
	return s.StepName
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) GetStepProgress() *string {
	return s.StepProgress
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) GetStepStatus() *string {
	return s.StepStatus
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) SetEndTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) SetStartTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) SetStepDesc(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList {
	s.StepDesc = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) SetStepName(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList {
	s.StepName = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) SetStepProgress(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList {
	s.StepProgress = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) SetStepStatus(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList {
	s.StepStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList) Validate() error {
	return dara.Validate(s)
}
