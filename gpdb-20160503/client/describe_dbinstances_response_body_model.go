// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBInstancesResponseBodyItems) *DescribeDBInstancesResponseBody
	GetItems() *DescribeDBInstancesResponseBodyItems
	SetPageNumber(v int32) *DescribeDBInstancesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDBInstancesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDBInstancesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDBInstancesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBInstancesResponseBody struct {
	// The queried instances.
	Items *DescribeDBInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
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
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 53EA07B7-FC2A-521B-AB7C-27**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) GetItems() *DescribeDBInstancesResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDBInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBInstancesResponseBody) SetItems(v *DescribeDBInstancesResponseBodyItems) *DescribeDBInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageNumber(v int32) *DescribeDBInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageRecordCount(v int32) *DescribeDBInstancesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetRequestId(v string) *DescribeDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyItems struct {
	DBInstance []*DescribeDBInstancesResponseBodyItemsDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItems) GetDBInstance() []*DescribeDBInstancesResponseBodyItemsDBInstance {
	return s.DBInstance
}

func (s *DescribeDBInstancesResponseBodyItems) SetDBInstance(v []*DescribeDBInstancesResponseBodyItemsDBInstance) *DescribeDBInstancesResponseBodyItems {
	s.DBInstance = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) Validate() error {
	if s.DBInstance != nil {
		for _, item := range s.DBInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyItemsDBInstance struct {
	// An invalid parameter. It is no longer returned when you call this operation.
	//
	// You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/86910.html) operation to query the access mode of an instance.
	//
	// example:
	//
	// null
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The time when the instance was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-10-09T04:54:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The edition of the instance. Valid values:
	//
	// 	- **Basic**: Basic Edition.
	//
	// 	- **HighAvailability**: High-availability Edition.
	//
	// 	- **Finance**: Enterprise Edition.
	//
	// example:
	//
	// HighAvailability
	DBInstanceCategory *string `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The resource type of the instance. Valid values:
	//
	// 	- **Serverless**: Serverless mode.
	//
	// 	- **StorageElastic**: elastic storage mode.
	//
	// 	- **Classic**: reserved storage mode.
	//
	// example:
	//
	// StorageElastic
	DBInstanceMode *string `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	// The type of the network interface card (NIC) that is used by the instance. Valid values:
	//
	// 	- **0**: Internet.
	//
	// 	- **1**: internal network.
	//
	// 	- **2**: VPC.
	//
	// example:
	//
	// 2
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The status of the instance. For more information, see [Instance statuses](https://help.aliyun.com/document_detail/86944.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// gpdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance.
	//
	// example:
	//
	// 6.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The expiration time of the instance. The time is displayed in UTC.
	//
	// >  The expiration time of a pay-as-you-go instance is `2999-09-08T16:00:00Z`.
	//
	// example:
	//
	// 2999-09-08T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The resource type of the instance. Valid values:
	//
	// 	- **cluster**: Serverless mode or elastic storage mode.
	//
	// 	- **replicaSet**: reserved storage mode.
	//
	// example:
	//
	// cluster
	InstanceDeployType *string `json:"InstanceDeployType,omitempty" xml:"InstanceDeployType,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **Classic**: classic network.
	//
	// 	- **VPC**: VPC.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// 	- **Unlock**: The instance is not locked.
	//
	// 	- **ManualLock**: The instance is manually locked.
	//
	// 	- **LockByExpiration**: The instance is automatically locked due to instance expiration.
	//
	// 	- **LockByRestoration**: The instance is automatically locked due to instance restoration.
	//
	// 	- **LockByDiskQuota**: The instance is automatically locked due to exhausted storage.
	//
	// 	- **LockReadInstanceByDiskQuota**: The instance is a read-only instance and is automatically locked when the disk space is full.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the instance is locked. Valid values:
	//
	// 	- **0**: The instance is not locked.
	//
	// 	- **1**: The instance is manually locked.
	//
	// 	- **2**: The instance is automatically locked due to instance expiration.
	//
	// 	- **3**: The instance is automatically locked due to instance restoration.
	//
	// 	- **4**: The instance is automatically locked due to exhausted storage.
	//
	// >  If the instance is in reserved storage mode and is not locked, null is returned.
	//
	// example:
	//
	// 0
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The number of coordinator nodes.
	//
	// example:
	//
	// 1
	MasterNodeNum *int32 `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// product type
	//
	// example:
	//
	// standard
	ProdType *string `json:"ProdType,omitempty" xml:"ProdType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of compute nodes.
	//
	// example:
	//
	// 4
	SegNodeNum *string `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	// The type of the Serverless mode. Valid values:
	//
	// 	- **Manual**: manual scheduling.
	//
	// 	- **Auto**: automatic scheduling.
	//
	// >  This parameter is returned only for instances in Serverless mode.
	//
	// example:
	//
	// Manual
	ServerlessMode *string `json:"ServerlessMode,omitempty" xml:"ServerlessMode,omitempty"`
	// The storage capacity of the instance. Unit: GB.
	//
	// example:
	//
	// 50
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **cloud_essd**: enhanced SSD (ESSD).
	//
	// 	- **cloud_efficiency**: ultra disk.
	//
	// example:
	//
	// cloud_essd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tags that are added to the instance.
	Tags *DescribeDBInstancesResponseBodyItemsDBInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID of the instance.
	//
	// example:
	//
	// vpc-bp19ame5m1r3oejns****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceCategory() *string {
	return s.DBInstanceCategory
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceMode() *string {
	return s.DBInstanceMode
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetInstanceDeployType() *string {
	return s.InstanceDeployType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetMasterNodeNum() *int32 {
	return s.MasterNodeNum
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetProdType() *string {
	return s.ProdType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetSegNodeNum() *string {
	return s.SegNodeNum
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetServerlessMode() *string {
	return s.ServerlessMode
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetStorageSize() *string {
	return s.StorageSize
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetTags() *DescribeDBInstancesResponseBodyItemsDBInstanceTags {
	return s.Tags
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetConnectionMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetCreateTime(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceCategory(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceCategory = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceDescription(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceNetType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceStatus(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetEngine(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetEngineVersion(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetExpireTime(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetInstanceDeployType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.InstanceDeployType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetInstanceNetworkType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetLockMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetLockReason(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetMasterNodeNum(v int32) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.MasterNodeNum = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetPayType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetProdType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ProdType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetResourceGroupId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetSegNodeNum(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.SegNodeNum = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetServerlessMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ServerlessMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetStorageSize(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetStorageType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetTags(v *DescribeDBInstancesResponseBodyItemsDBInstanceTags) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.Tags = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetVSwitchId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetVpcId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyItemsDBInstanceTags struct {
	Tag []*DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTags) GetTag() []*DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag {
	return s.Tag
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTags) SetTag(v []*DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) *DescribeDBInstancesResponseBodyItemsDBInstanceTags {
	s.Tag = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTags) Validate() error {
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

type DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag struct {
	// The key of tag N.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) SetKey(v string) *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) SetValue(v string) *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) Validate() error {
	return dara.Validate(s)
}
