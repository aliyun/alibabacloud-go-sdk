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
	SetNextToken(v string) *DescribeDBInstancesResponseBody
	GetNextToken() *string
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
	Items *DescribeDBInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The token that is used to display the next page. If the returned entries are displayed on multiple pages, the next page can be displayed when you call this operation again with **NextToken*	- specified.
	//
	// example:
	//
	// o7PORW5o2TJg**********
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number of the returned page.
	//
	// > If you specify **MaxResults*	- or **NextToken**, only the value **1*	- is returned. You can ignore the value 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// > If you specify **MaxResults*	- or **NextToken**, only the number of entries on the current page is returned. You can ignore the number.
	//
	// example:
	//
	// 100
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

func (s *DescribeDBInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
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

func (s *DescribeDBInstancesResponseBody) SetNextToken(v string) *DescribeDBInstancesResponseBody {
	s.NextToken = &v
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
	AutoRenewal                  *bool                                                                `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	BlueGreenDeploymentName      *string                                                              `json:"BlueGreenDeploymentName,omitempty" xml:"BlueGreenDeploymentName,omitempty"`
	BlueInstanceName             *string                                                              `json:"BlueInstanceName,omitempty" xml:"BlueInstanceName,omitempty"`
	BpeEnabled                   *string                                                              `json:"BpeEnabled,omitempty" xml:"BpeEnabled,omitempty"`
	BurstingEnabled              *bool                                                                `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Category                     *string                                                              `json:"Category,omitempty" xml:"Category,omitempty"`
	ColdDataEnabled              *bool                                                                `json:"ColdDataEnabled,omitempty" xml:"ColdDataEnabled,omitempty"`
	ConnectionMode               *string                                                              `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	ConnectionString             *string                                                              `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	CreateTime                   *string                                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBInstanceCPU                *string                                                              `json:"DBInstanceCPU,omitempty" xml:"DBInstanceCPU,omitempty"`
	DBInstanceClass              *string                                                              `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceDescription        *string                                                              `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId                 *string                                                              `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceMemory             *int32                                                               `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty"`
	DBInstanceNetType            *string                                                              `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	DBInstanceStatus             *string                                                              `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceStorageType        *string                                                              `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	DBInstanceType               *string                                                              `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	DedicatedHostGroupId         *string                                                              `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DedicatedHostGroupName       *string                                                              `json:"DedicatedHostGroupName,omitempty" xml:"DedicatedHostGroupName,omitempty"`
	DedicatedHostIdForLog        *string                                                              `json:"DedicatedHostIdForLog,omitempty" xml:"DedicatedHostIdForLog,omitempty"`
	DedicatedHostIdForMaster     *string                                                              `json:"DedicatedHostIdForMaster,omitempty" xml:"DedicatedHostIdForMaster,omitempty"`
	DedicatedHostIdForSlave      *string                                                              `json:"DedicatedHostIdForSlave,omitempty" xml:"DedicatedHostIdForSlave,omitempty"`
	DedicatedHostNameForLog      *string                                                              `json:"DedicatedHostNameForLog,omitempty" xml:"DedicatedHostNameForLog,omitempty"`
	DedicatedHostNameForMaster   *string                                                              `json:"DedicatedHostNameForMaster,omitempty" xml:"DedicatedHostNameForMaster,omitempty"`
	DedicatedHostNameForSlave    *string                                                              `json:"DedicatedHostNameForSlave,omitempty" xml:"DedicatedHostNameForSlave,omitempty"`
	DedicatedHostZoneIdForLog    *string                                                              `json:"DedicatedHostZoneIdForLog,omitempty" xml:"DedicatedHostZoneIdForLog,omitempty"`
	DedicatedHostZoneIdForMaster *string                                                              `json:"DedicatedHostZoneIdForMaster,omitempty" xml:"DedicatedHostZoneIdForMaster,omitempty"`
	DedicatedHostZoneIdForSlave  *string                                                              `json:"DedicatedHostZoneIdForSlave,omitempty" xml:"DedicatedHostZoneIdForSlave,omitempty"`
	DeletionProtection           *bool                                                                `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	DestroyTime                  *string                                                              `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	Engine                       *string                                                              `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion                *string                                                              `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime                   *string                                                              `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GeneralGroupName             *string                                                              `json:"GeneralGroupName,omitempty" xml:"GeneralGroupName,omitempty"`
	GreenInstanceName            *string                                                              `json:"GreenInstanceName,omitempty" xml:"GreenInstanceName,omitempty"`
	GuardDBInstanceId            *string                                                              `json:"GuardDBInstanceId,omitempty" xml:"GuardDBInstanceId,omitempty"`
	InstanceNetworkType          *string                                                              `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	IoAccelerationEnabled        *string                                                              `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	IsAnalyticIns                *string                                                              `json:"IsAnalyticIns,omitempty" xml:"IsAnalyticIns,omitempty"`
	IsAnalyticReadOnlyIns        *bool                                                                `json:"IsAnalyticReadOnlyIns,omitempty" xml:"IsAnalyticReadOnlyIns,omitempty"`
	LockMode                     *string                                                              `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason                   *string                                                              `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MasterInstanceId             *string                                                              `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	MutriORsignle                *bool                                                                `json:"MutriORsignle,omitempty" xml:"MutriORsignle,omitempty"`
	PayType                      *string                                                              `json:"PayType,omitempty" xml:"PayType,omitempty"`
	ReadOnlyDBInstanceIds        *DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIds `json:"ReadOnlyDBInstanceIds,omitempty" xml:"ReadOnlyDBInstanceIds,omitempty" type:"Struct"`
	RegionId                     *string                                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId              *string                                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SwitchWeight                 *int32                                                               `json:"SwitchWeight,omitempty" xml:"SwitchWeight,omitempty"`
	TempDBInstanceId             *string                                                              `json:"TempDBInstanceId,omitempty" xml:"TempDBInstanceId,omitempty"`
	Tips                         *string                                                              `json:"Tips,omitempty" xml:"Tips,omitempty"`
	TipsLevel                    *int32                                                               `json:"TipsLevel,omitempty" xml:"TipsLevel,omitempty"`
	VSwitchId                    *string                                                              `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcCloudInstanceId           *string                                                              `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	VpcId                        *string                                                              `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName                      *string                                                              `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	ZoneId                       *string                                                              `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetBlueGreenDeploymentName() *string {
	return s.BlueGreenDeploymentName
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetBlueInstanceName() *string {
	return s.BlueInstanceName
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetBpeEnabled() *string {
	return s.BpeEnabled
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceCPU() *string {
	return s.DBInstanceCPU
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceMemory() *int32 {
	return s.DBInstanceMemory
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDedicatedHostGroupName() *string {
	return s.DedicatedHostGroupName
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDedicatedHostIdForLog() *string {
	return s.DedicatedHostIdForLog
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDedicatedHostIdForMaster() *string {
	return s.DedicatedHostIdForMaster
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDedicatedHostIdForSlave() *string {
	return s.DedicatedHostIdForSlave
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDedicatedHostNameForLog() *string {
	return s.DedicatedHostNameForLog
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDedicatedHostNameForMaster() *string {
	return s.DedicatedHostNameForMaster
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDedicatedHostNameForSlave() *string {
	return s.DedicatedHostNameForSlave
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDedicatedHostZoneIdForLog() *string {
	return s.DedicatedHostZoneIdForLog
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDedicatedHostZoneIdForMaster() *string {
	return s.DedicatedHostZoneIdForMaster
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDedicatedHostZoneIdForSlave() *string {
	return s.DedicatedHostZoneIdForSlave
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDestroyTime() *string {
	return s.DestroyTime
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

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetGeneralGroupName() *string {
	return s.GeneralGroupName
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetGreenInstanceName() *string {
	return s.GreenInstanceName
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetGuardDBInstanceId() *string {
	return s.GuardDBInstanceId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetIsAnalyticIns() *string {
	return s.IsAnalyticIns
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetIsAnalyticReadOnlyIns() *bool {
	return s.IsAnalyticReadOnlyIns
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetMasterInstanceId() *string {
	return s.MasterInstanceId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetMutriORsignle() *bool {
	return s.MutriORsignle
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetReadOnlyDBInstanceIds() *DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIds {
	return s.ReadOnlyDBInstanceIds
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetSwitchWeight() *int32 {
	return s.SwitchWeight
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetTempDBInstanceId() *string {
	return s.TempDBInstanceId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetTips() *string {
	return s.Tips
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetTipsLevel() *int32 {
	return s.TipsLevel
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetVpcCloudInstanceId() *string {
	return s.VpcCloudInstanceId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetAutoRenewal(v bool) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetBlueGreenDeploymentName(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.BlueGreenDeploymentName = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetBlueInstanceName(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.BlueInstanceName = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetBpeEnabled(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.BpeEnabled = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetBurstingEnabled(v bool) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetCategory(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.Category = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetColdDataEnabled(v bool) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ColdDataEnabled = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetConnectionMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetConnectionString(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetCreateTime(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceCPU(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceCPU = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceClass(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceClass = &v
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

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceMemory(v int32) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceMemory = &v
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

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceStorageType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceStorageType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDedicatedHostGroupId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDedicatedHostGroupName(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DedicatedHostGroupName = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDedicatedHostIdForLog(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DedicatedHostIdForLog = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDedicatedHostIdForMaster(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DedicatedHostIdForMaster = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDedicatedHostIdForSlave(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DedicatedHostIdForSlave = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDedicatedHostNameForLog(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DedicatedHostNameForLog = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDedicatedHostNameForMaster(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DedicatedHostNameForMaster = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDedicatedHostNameForSlave(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DedicatedHostNameForSlave = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDedicatedHostZoneIdForLog(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DedicatedHostZoneIdForLog = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDedicatedHostZoneIdForMaster(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DedicatedHostZoneIdForMaster = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDedicatedHostZoneIdForSlave(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DedicatedHostZoneIdForSlave = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDeletionProtection(v bool) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDestroyTime(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DestroyTime = &v
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

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetGeneralGroupName(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.GeneralGroupName = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetGreenInstanceName(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.GreenInstanceName = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetGuardDBInstanceId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.GuardDBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetInstanceNetworkType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetIoAccelerationEnabled(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetIsAnalyticIns(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.IsAnalyticIns = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetIsAnalyticReadOnlyIns(v bool) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.IsAnalyticReadOnlyIns = &v
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

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetMasterInstanceId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetMutriORsignle(v bool) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.MutriORsignle = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetPayType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetReadOnlyDBInstanceIds(v *DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIds) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ReadOnlyDBInstanceIds = v
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

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetSwitchWeight(v int32) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.SwitchWeight = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetTempDBInstanceId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.TempDBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetTips(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.Tips = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetTipsLevel(v int32) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.TipsLevel = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetVSwitchId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetVpcCloudInstanceId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetVpcId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetVpcName(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.VpcName = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) Validate() error {
	if s.ReadOnlyDBInstanceIds != nil {
		if err := s.ReadOnlyDBInstanceIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIds struct {
	ReadOnlyDBInstanceId []*DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIds) GetReadOnlyDBInstanceId() []*DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId {
	return s.ReadOnlyDBInstanceId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIds) SetReadOnlyDBInstanceId(v []*DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId) *DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIds {
	s.ReadOnlyDBInstanceId = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIds) Validate() error {
	if s.ReadOnlyDBInstanceId != nil {
		for _, item := range s.ReadOnlyDBInstanceId {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId) SetDBInstanceId(v string) *DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId) Validate() error {
	return dara.Validate(s)
}
