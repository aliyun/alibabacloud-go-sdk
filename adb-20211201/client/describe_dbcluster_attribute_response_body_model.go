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
	// The queried cluster.
	Items *DescribeDBClusterAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A5C433C2-001F-58E3-99F5-3274C14DF8BD
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeDBClusterAttributeResponseBodyItemsDBCluster struct {
	AINodeNumber *int32  `json:"AINodeNumber,omitempty" xml:"AINodeNumber,omitempty"`
	AINodeSpec   *string `json:"AINodeSpec,omitempty" xml:"AINodeSpec,omitempty"`
	// The cache size of the ClickHouse wide table engine. Unit: GB. If a value of -1 is returned, the ClickHouse wide table engine is disabled. If a value other than -1 is returned, this parameter indicates the disk cache size.
	//
	// example:
	//
	// 100
	ClickhouseEngineCacheSize *int32 `json:"ClickhouseEngineCacheSize,omitempty" xml:"ClickhouseEngineCacheSize,omitempty"`
	// Indicates whether the ClickHouse wide table engine is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ClickhouseEngineEnabled *bool `json:"ClickhouseEngineEnabled,omitempty" xml:"ClickhouseEngineEnabled,omitempty"`
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
	// The total amount of computing resources in the cluster. Each ACU is approximately equal to 1 core and 4 GB memory.
	//
	// example:
	//
	// 48ACU
	ComputeResourceTotal *string `json:"ComputeResourceTotal,omitempty" xml:"ComputeResourceTotal,omitempty"`
	// The public endpoint that is used to connect to the cluster.
	//
	// example:
	//
	// amv-wz9509beptiz****.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-07-01T09:50:18Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// adb_test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// example:
	//
	// amv-wz9509beptiz****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster. **VPC*	- is returned.
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
	// The engine version of the AnalyticDB for MySQL Data Lakehouse Edition cluster. **5.0*	- is returned.
	//
	// example:
	//
	// 5.0
	DBVersion      *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DiskEncryption *bool   `json:"DiskEncryption,omitempty" xml:"DiskEncryption,omitempty"`
	// The engine of the cluster. **AnalyticDB*	- is returned.
	//
	// example:
	//
	// AnalyticDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The minor version of the cluster.
	//
	// example:
	//
	// 3.1.16
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the cluster expires.
	//
	// 	- If the billing method of the cluster is subscription, the actual expiration time is returned.
	//
	// 	- If the billing method of the cluster is pay-as-you-go, null is returned.
	//
	// example:
	//
	// 2022-10-01T09:50:18Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the subscription cluster has expired. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >
	//
	// 	- If the cluster has expired, the system locks or releases the cluster within a period of time. We recommend that you renew the expired cluster. For more information, see [Renewal policy](https://help.aliyun.com/document_detail/135248.html).
	//
	// 	- This parameter is not returned for pay-as-you-go clusters.
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The ID of the key that is used to encrypt disk data.
	//
	// >  This parameter is returned only when disk encryption is enabled.
	//
	// example:
	//
	// e1935511-cf88-1123-a0f8-1be8d251****
	KmsId *string `json:"KmsId,omitempty" xml:"KmsId,omitempty"`
	// The lock mode of the cluster. Valid values:
	//
	// 	- **Unlock**: The cluster is not locked.
	//
	// 	- **ManualLock**: The cluster is manually locked.
	//
	// 	- **LockByExpiration**: The cluster is automatically locked due to cluster expiration.
	//
	// example:
	//
	// ManualLock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the cluster is locked.
	//
	// >  This parameter is returned only when the cluster was locked. **instance_expire*	- is returned.
	//
	// example:
	//
	// instance_expire
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The maintenance window of the cluster. The time is displayed in the `HH:mmZ-HH:mmZ` format in UTC.
	//
	// >  For more information about maintenance windows, see [Configure a maintenance window](https://help.aliyun.com/document_detail/122569.html).
	//
	// example:
	//
	// 04:00Z-05:00Z
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
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
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	ProductForm *string `json:"ProductForm,omitempty" xml:"ProductForm,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// 	- **BasicVersion**: Basic Edition.
	//
	// 	- **EnterpriseVersion**: Enterprise Edition.
	//
	// example:
	//
	// BasicVersion
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The amount of remaining reserved computing resources that are available in the cluster. Each ACU is approximately equal to 1 core and 4 GB memory.
	//
	// example:
	//
	// 24ACU
	ReservedACU *string `json:"ReservedACU,omitempty" xml:"ReservedACU,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	ReservedNodeCount *int32 `json:"ReservedNodeCount,omitempty" xml:"ReservedNodeCount,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	ReservedNodeSize *string `json:"ReservedNodeSize,omitempty" xml:"ReservedNodeSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId    *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecondaryVSwitchId *string `json:"SecondaryVSwitchId,omitempty" xml:"SecondaryVSwitchId,omitempty"`
	SecondaryZoneId    *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The specifications of reserved storage resources. Each AnalyticDB compute unit (ACU) is approximately equal to 1 core and 4 GB memory. Storage resources are used to read and write data. The increase in the storage resources can improve the read and write performance of the cluster.
	//
	// example:
	//
	// 24ACU
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	// The total amount of storage resources in the cluster. Each ACU is approximately equal to 1 core and 4 GB memory.
	//
	// example:
	//
	// 24ACU
	StorageResourceTotal *string `json:"StorageResourceTotal,omitempty" xml:"StorageResourceTotal,omitempty"`
	// Reserved parameters.
	SupportedFeatures map[string]*string `json:"SupportedFeatures,omitempty" xml:"SupportedFeatures,omitempty"`
	// The tags that are added to the cluster.
	Tags *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The job information.
	TaskInfo *DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
	// Indicates whether Elastic Network Interface (ENI) is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	UserENIStatus *bool `json:"UserENIStatus,omitempty" xml:"UserENIStatus,omitempty"`
	// The vSwitch connected to the ENI. Separate multiple vSwitches with commas (,).
	//
	// example:
	//
	// vsw-rj9ixufmywqq98z******,vsw-rj95ij6wcz656v7******
	UserENIVSwitchOptions *string `json:"UserENIVSwitchOptions,omitempty" xml:"UserENIVSwitchOptions,omitempty"`
	// The VPC information of the ENI.
	//
	// example:
	//
	// vpc-rj9hnedlfm645uj******
	UserENIVpcId *string `json:"UserENIVpcId,omitempty" xml:"UserENIVpcId,omitempty"`
	// The zone associated with the ENI. Separate multiple zones with commas (,).
	//
	// example:
	//
	// cn-hangzhou-k,cn-hangzhou-h
	UserENIZoneOptions *string `json:"UserENIZoneOptions,omitempty" xml:"UserENIZoneOptions,omitempty"`
	// The virtual private cloud (VPC) ID of the cluster.
	//
	// example:
	//
	// vpc-bp13h7uzhulpu****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	//
	// example:
	//
	// vsw-uf629gydd54ld****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetAINodeNumber() *int32 {
	return s.AINodeNumber
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetAINodeSpec() *string {
	return s.AINodeSpec
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetClickhouseEngineCacheSize() *int32 {
	return s.ClickhouseEngineCacheSize
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetClickhouseEngineEnabled() *bool {
	return s.ClickhouseEngineEnabled
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetComputeResourceTotal() *string {
	return s.ComputeResourceTotal
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

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetExpired() *string {
	return s.Expired
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

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetReservedACU() *string {
	return s.ReservedACU
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

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetStorageResourceTotal() *string {
	return s.StorageResourceTotal
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetSupportedFeatures() map[string]*string {
	return s.SupportedFeatures
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

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetUserENIVSwitchOptions() *string {
	return s.UserENIVSwitchOptions
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetUserENIVpcId() *string {
	return s.UserENIVpcId
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) GetUserENIZoneOptions() *string {
	return s.UserENIZoneOptions
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

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetAINodeNumber(v int32) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.AINodeNumber = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetAINodeSpec(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.AINodeSpec = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetClickhouseEngineCacheSize(v int32) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ClickhouseEngineCacheSize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetClickhouseEngineEnabled(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ClickhouseEngineEnabled = &v
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

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetComputeResourceTotal(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ComputeResourceTotal = &v
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

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDiskEncryption(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DiskEncryption = &v
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

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Expired = &v
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

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetReservedACU(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ReservedACU = &v
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

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetStorageResourceTotal(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.StorageResourceTotal = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetSupportedFeatures(v map[string]*string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.SupportedFeatures = v
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

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetUserENIVSwitchOptions(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.UserENIVSwitchOptions = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetUserENIVpcId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.UserENIVpcId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetUserENIZoneOptions(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.UserENIZoneOptions = &v
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag struct {
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
	// The name of the job.
	//
	// example:
	//
	// ScaleUpDBCluster
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The progress of the job. Unit: %.
	//
	// example:
	//
	// 10
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- NOT_RUN
	//
	// 	- RUNNING
	//
	// 	- SUCCEED
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The job steps.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTaskInfoStepListStepList struct {
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
	// Prepare resources
	StepDesc *string `json:"StepDesc,omitempty" xml:"StepDesc,omitempty"`
	// The name of the job step.
	//
	// example:
	//
	// PrepareResources
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The progress of the job step. Unit: %.
	//
	// example:
	//
	// 50
	StepProgress *string `json:"StepProgress,omitempty" xml:"StepProgress,omitempty"`
	// The status of the job step. Valid values:
	//
	// 	- NOT_RUN
	//
	// 	- RUNNING
	//
	// 	- SUCCEED
	//
	// example:
	//
	// RUNNING
	StepStatus *string `json:"StepStatus,omitempty" xml:"StepStatus,omitempty"`
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
