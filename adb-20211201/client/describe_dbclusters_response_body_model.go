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
	// The queried clusters.
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
	// A5EDBA27-AF3E-5966-9503-FD1557E19167
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeDBClustersResponseBodyItemsDBCluster struct {
	// The mode of the cluster. This parameter is returned only for Data Warehouse Edition clusters. Valid values:
	//
	// 	- **BASIC**: reserved mode for Basic Edition.
	//
	// 	- **CLUSTER**: reserved mode for Cluster Edition.
	//
	// 	- **MIXED_STORAGE**: elastic mode for Cluster Edition.
	//
	// >  For more information about cluster editions, see [Editions](https://help.aliyun.com/document_detail/205001.html).
	//
	// example:
	//
	// MIXED_STORAGE
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **ads**: pay-as-you-go.
	//
	// 	- **ads_pre**: subscription.
	//
	// example:
	//
	// ads_pre
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The specifications of reserved computing resources. Each ACU is approximately equal to 1 core and 4 GB memory. Computing resources are used to compute data. The increase in the computing resources can accelerate queries. You can scale computing resources based on your business requirements.
	//
	// example:
	//
	// 16ACU
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The public endpoint that is used to connect to the cluster.
	//
	// example:
	//
	// amv-bp163885f8q21****.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the *yyyy-mm-ddThh:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-01T09:50:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// adb_test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// amv-bp163885f8q21****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster. Only **VPC*	- is supported.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- **Preparing**
	//
	// 	- **Creating**
	//
	// 	- **Running**
	//
	// 	- **Deleting**
	//
	// 	- **Restoring**
	//
	// 	- **ClassChanging**
	//
	// 	- **NetAddressCreating**
	//
	// 	- **NetAddressDeleting**
	//
	// 	- **NetAddressModifying**
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The type of the cluster. By default, **Common*	- is returned, which indicates a common cluster.
	//
	// example:
	//
	// Common
	DBClusterType *string `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	// The node specifications of the cluster. This parameter is returned only for Data Warehouse Edition clusters.
	//
	// example:
	//
	// E8
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of node groups.
	//
	// example:
	//
	// 2
	DBNodeCount *int64 `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	// The storage capacity of the cluster. Unit: GB.
	//
	// example:
	//
	// 300
	DBNodeStorage *int64 `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The version number corresponding to the edition of the cluster. Only **5.0*	- is supported.
	//
	// example:
	//
	// 5.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The disk type of the cluster. Valid values:
	//
	// 	- **local_ssd**: local disk.
	//
	// 	- **cloud**: basic disk.
	//
	// 	- **cloud_ssd**: standard SSD.
	//
	// 	- **cloud_efficiency**: ultra disk.
	//
	// 	- **cloud_essd**: PL1 Enterprise SSD (ESSD).
	//
	// 	- **cloud_essd2**: PL2 ESSD.
	//
	// 	- **cloud_essd3**: PL3 ESSD.
	//
	// >  For more information about ESSDs, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// cloud_essd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the Data Transmission Service (DTS) synchronization job This parameter is returned only for MySQL analytic instances.
	//
	// example:
	//
	// dtsb1578j90XXXX
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The number of elastic I/O units (EIUs). For more information, see the "[EIUs](https://help.aliyun.com/document_detail/189505.html)" section of the Scale out elastic I/O resources topic.
	//
	// >  This parameter is returned only for clusters in elastic mode.
	//
	// example:
	//
	// 2
	ElasticIOResource *int32 `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	// The engine of the cluster. **AnalyticDB*	- is returned.
	//
	// example:
	//
	// AnalyticDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The number of compute nodes that are used by the cluster in elastic mode.
	//
	// example:
	//
	// 1
	ExecutorCount *string `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	// The time when the cluster expires. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// >
	//
	// 	- If the billing method of the cluster is subscription, the actual expiration time is returned.
	//
	// 	- If the billing method of the cluster is pay-as-you-go, null is returned.
	//
	// example:
	//
	// 2022-07-01T09:50:18Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the subscription cluster has expired. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >
	//
	// 	- If the cluster has expired, the system locks or releases the cluster within a period of time. We recommend that you renew the expired cluster. For more information, see [Renewal policy](https://help.aliyun.com/document_detail/135246.html).
	//
	// 	- This parameter is not returned for pay-as-you-go clusters.
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The internal IP address of the cluster.
	//
	// example:
	//
	// 10.1.xx.xx
	InnerIp *string `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	// The internal port of the cluster.
	//
	// example:
	//
	// 3306
	InnerPort *string `json:"InnerPort,omitempty" xml:"InnerPort,omitempty"`
	// The lock status of the cluster. Valid values:
	//
	// 	- **Unlock**: The cluster is not locked.
	//
	// 	- **ManualLock**: The cluster is manually locked.
	//
	// 	- **LockByExpiration**: The cluster is automatically locked due to cluster expiration.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the cluster is locked.
	//
	// >  This parameter is returned only when the cluster was locked. **instance_expire*	- is returned.
	//
	// example:
	//
	// instance_expire
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The mode of the cluster. By default, **flexible*	- is returned, which indicates that the cluster is in elastic mode.
	//
	// example:
	//
	// flexible
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port number that is used to connect to the cluster.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The service type of the cluster. Valid values:
	//
	// 	- **LegacyForm**
	//
	// 	- **IntegrationForm**
	//
	// example:
	//
	// IntegrationForm
	ProductForm *string `json:"ProductForm,omitempty" xml:"ProductForm,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// 	- **BasicVersion**: Basic Edition.
	//
	// 	- **EnterpriseVersion**: Enterprise Edition.
	//
	// example:
	//
	// EnterpriseVersion
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// The ID of the ApsaraDB RDS instance from which data is synchronized to the cluster. This parameter is returned only for MySQL analytic instances.
	//
	// example:
	//
	// rm-bp11q28kvl688****
	RdsInstanceId *string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remaining reserved computing resources that are available in the cluster. Each ACU is approximately equal to 1 core and 4 GB memory.
	//
	// example:
	//
	// 32ACU
	ReservedACU *string `json:"ReservedACU,omitempty" xml:"ReservedACU,omitempty"`
	// The number of reserved resource nodes.
	//
	// example:
	//
	// 1
	ReservedNodeCount *int32 `json:"ReservedNodeCount,omitempty" xml:"ReservedNodeCount,omitempty"`
	// The single-node specifications of reserved resources.
	//
	// example:
	//
	// 8ACU
	ReservedNodeSize *string `json:"ReservedNodeSize,omitempty" xml:"ReservedNodeSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The specifications of reserved storage resources. Each AnalyticDB compute unit (ACU) is approximately equal to 1 core and 4 GB memory. Storage resources are used to read and write data. The increase in the storage resources can improve the read and write performance of the cluster.
	//
	// example:
	//
	// 24ACU
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	// The tags that are added to the cluster.
	Tags *DescribeDBClustersResponseBodyItemsDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The information about the job.
	TaskInfo *DescribeDBClustersResponseBodyItemsDBClusterTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
	// The VPC endpoint.
	//
	// example:
	//
	// am-bp163885f8q21****-controller
	VPCCloudInstanceId *string `json:"VPCCloudInstanceId,omitempty" xml:"VPCCloudInstanceId,omitempty"`
	// The virtual private cloud (VPC) ID of the cluster.
	//
	// example:
	//
	// vpc-bp13h7uzhulpuxvnp****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	//
	// example:
	//
	// vsw-bp1syh8vvw8yech7n****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetProductForm() *string {
	return s.ProductForm
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

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetReservedACU() *string {
	return s.ReservedACU
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetReservedNodeCount() *int32 {
	return s.ReservedNodeCount
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) GetReservedNodeSize() *string {
	return s.ReservedNodeSize
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

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetProductForm(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ProductForm = &v
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

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetReservedACU(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ReservedACU = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetReservedNodeCount(v int32) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ReservedNodeCount = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetReservedNodeSize(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ReservedNodeSize = &v
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
	return dara.Validate(s)
}

type DescribeDBClustersResponseBodyItemsDBClusterTagsTag struct {
	// The tag key.
	//
	// >  You can call the [TagResources](https://help.aliyun.com/document_detail/179253.html) operation to add tags to a cluster.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test1
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
	// The name of the job.
	//
	// example:
	//
	// analyticDBFlexibleScaleOut
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The progress of the job. Unit: %.
	//
	// example:
	//
	// 10
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- **NOT_RUN**
	//
	// 	- **RUNNING**
	//
	// 	- **SUCCEED**
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The job steps.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeDBClustersResponseBodyItemsDBClusterTaskInfoStepListStepList struct {
	// The end time of the job step. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-03-10T10:28:34Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the job step. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-03-10T09:28:34Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The description of the job step.
	//
	// example:
	//
	// Apply resource
	StepDesc *string `json:"StepDesc,omitempty" xml:"StepDesc,omitempty"`
	// The name of the job step.
	//
	// example:
	//
	// ApplyResource
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The progress of the job step. Unit: %.
	//
	// example:
	//
	// 50
	StepProgress *string `json:"StepProgress,omitempty" xml:"StepProgress,omitempty"`
	// The status of the job step. Valid values:
	//
	// 	- **NOT_RUN**
	//
	// 	- **RUNNING**
	//
	// 	- **SUCCEED**
	//
	// example:
	//
	// SUCCEED
	StepStatus *string `json:"StepStatus,omitempty" xml:"StepStatus,omitempty"`
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
