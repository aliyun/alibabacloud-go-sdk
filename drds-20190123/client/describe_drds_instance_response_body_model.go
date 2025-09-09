// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDrdsInstanceResponseBodyData) *DescribeDrdsInstanceResponseBody
	GetData() *DescribeDrdsInstanceResponseBodyData
	SetRequestId(v string) *DescribeDrdsInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsInstanceResponseBody
	GetSuccess() *bool
}

type DescribeDrdsInstanceResponseBody struct {
	// The details of the instance.
	Data *DescribeDrdsInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B4F76641-BA45-4320-BE7C-9C62CFDAC9B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBody) GetData() *DescribeDrdsInstanceResponseBodyData {
	return s.Data
}

func (s *DescribeDrdsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsInstanceResponseBody) SetData(v *DescribeDrdsInstanceResponseBodyData) *DescribeDrdsInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceResponseBody) SetRequestId(v string) *DescribeDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstanceResponseBodyData struct {
	// The commodity code of the instance.
	//
	// example:
	//
	// drdsPost
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The timestamp that indicates when the instance is created.
	//
	// example:
	//
	// 1568620311000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// drds_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// drdssen1243as
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The timestamp that indicates when the instance expires.
	//
	// example:
	//
	// 4724323200000
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The role of the instance. Valid values:
	//
	// 	- **MASTER**: The instance is a primary instance.
	//
	// 	- **SLAVE**: The instance is a read-only instance to analyze complex queries
	//
	// 	- **SLAVE_FLOW**: The instance is a read-only instance for high-concurrency scenarios
	//
	// example:
	//
	// MASTER
	InstRole *string `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
	// The instance series of the instance.
	//
	// example:
	//
	// drds.sn2.4c16g
	InstanceSeries *string `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	// The specification of the instance.
	//
	// example:
	//
	// drds.sn2.4c16g.8C32G
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The tag of the instance. Valid values:
	//
	// 	- **NORMAL**: The instance is a standard instance.
	//
	// 	- **HA**: The instance is a high-availability (HA) instance.
	//
	// 	- **VPC**: The instance is a VPC-based instance.
	//
	// example:
	//
	// NORMAL
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The machine type of the instance. The value of this parameter is **ecs**.
	//
	// example:
	//
	// ecs
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The ID of the primary instance.
	//
	// >  This parameter is returned only when the instance is a primary instance.
	//
	// example:
	//
	// drdssen1243as
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	// The MySQL version that is supported by the instance.
	//
	// example:
	//
	// 5
	MysqlVersion *int32 `json:"MysqlVersion,omitempty" xml:"MysqlVersion,omitempty"`
	// The network type of the instance. Valid values: CLASSIC and VPC.
	//
	// example:
	//
	// CLASSIC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the purchased instance.
	//
	// example:
	//
	// drdssen12****
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// The version of .
	//
	// example:
	//
	// 5.3.*
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// The details about each read-only instance that is associated with the instance.
	ReadOnlyDBInstanceIds *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds `json:"ReadOnlyDBInstanceIds,omitempty" xml:"ReadOnlyDBInstanceIds,omitempty" type:"Struct"`
	// The ID of the region in which the instance is created.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. The value of this parameter can be null.
	//
	// example:
	//
	// NULL
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The state of the instance.
	//
	// example:
	//
	// RUN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the instance used for storage.
	//
	// example:
	//
	// RDS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The type of the instance. Valid values: PRIVATE and PUBLIC. The value of PRIVATE indicates that the instance is a dedicated instance. The value of PUBLIC indicates that the instance is a shared instance.
	//
	// example:
	//
	// PRIVATE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the instance. The value of this parameter is 0.
	//
	// example:
	//
	// 0
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
	// Indicates whether the version of the instance can be upgraded.
	//
	// example:
	//
	// Upgradeable
	VersionAction *string `json:"VersionAction,omitempty" xml:"VersionAction,omitempty"`
	// The list of returned virtual IP addresses (VIPs).
	Vips *DescribeDrdsInstanceResponseBodyDataVips `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
	// The ID of the instance that is deployed in the VPC.
	//
	// example:
	//
	// drdssen12****
	VpcCloudInstanceId *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	// The ID of the zone in which the instance is located.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDrdsInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDrdsInstanceResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeDrdsInstanceResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeDrdsInstanceResponseBodyData) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsInstanceResponseBodyData) GetExpireDate() *int64 {
	return s.ExpireDate
}

func (s *DescribeDrdsInstanceResponseBodyData) GetInstRole() *string {
	return s.InstRole
}

func (s *DescribeDrdsInstanceResponseBodyData) GetInstanceSeries() *string {
	return s.InstanceSeries
}

func (s *DescribeDrdsInstanceResponseBodyData) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeDrdsInstanceResponseBodyData) GetLabel() *string {
	return s.Label
}

func (s *DescribeDrdsInstanceResponseBodyData) GetMachineType() *string {
	return s.MachineType
}

func (s *DescribeDrdsInstanceResponseBodyData) GetMasterInstanceId() *string {
	return s.MasterInstanceId
}

func (s *DescribeDrdsInstanceResponseBodyData) GetMysqlVersion() *int32 {
	return s.MysqlVersion
}

func (s *DescribeDrdsInstanceResponseBodyData) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDrdsInstanceResponseBodyData) GetOrderInstanceId() *string {
	return s.OrderInstanceId
}

func (s *DescribeDrdsInstanceResponseBodyData) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *DescribeDrdsInstanceResponseBodyData) GetReadOnlyDBInstanceIds() *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds {
	return s.ReadOnlyDBInstanceIds
}

func (s *DescribeDrdsInstanceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrdsInstanceResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDrdsInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeDrdsInstanceResponseBodyData) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDrdsInstanceResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DescribeDrdsInstanceResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *DescribeDrdsInstanceResponseBodyData) GetVersionAction() *string {
	return s.VersionAction
}

func (s *DescribeDrdsInstanceResponseBodyData) GetVips() *DescribeDrdsInstanceResponseBodyDataVips {
	return s.Vips
}

func (s *DescribeDrdsInstanceResponseBodyData) GetVpcCloudInstanceId() *string {
	return s.VpcCloudInstanceId
}

func (s *DescribeDrdsInstanceResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDrdsInstanceResponseBodyData) SetCommodityCode(v string) *DescribeDrdsInstanceResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetCreateTime(v int64) *DescribeDrdsInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetDescription(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetDrdsInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetExpireDate(v int64) *DescribeDrdsInstanceResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetInstRole(v string) *DescribeDrdsInstanceResponseBodyData {
	s.InstRole = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetInstanceSeries(v string) *DescribeDrdsInstanceResponseBodyData {
	s.InstanceSeries = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetInstanceSpec(v string) *DescribeDrdsInstanceResponseBodyData {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetLabel(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Label = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetMachineType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.MachineType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetMasterInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetMysqlVersion(v int32) *DescribeDrdsInstanceResponseBodyData {
	s.MysqlVersion = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetNetworkType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetOrderInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.OrderInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetProductVersion(v string) *DescribeDrdsInstanceResponseBodyData {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetReadOnlyDBInstanceIds(v *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) *DescribeDrdsInstanceResponseBodyData {
	s.ReadOnlyDBInstanceIds = v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetRegionId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetResourceGroupId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetStatus(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetStorageType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.StorageType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVersion(v int64) *DescribeDrdsInstanceResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVersionAction(v string) *DescribeDrdsInstanceResponseBodyData {
	s.VersionAction = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVips(v *DescribeDrdsInstanceResponseBodyDataVips) *DescribeDrdsInstanceResponseBodyData {
	s.Vips = v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVpcCloudInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetZoneId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds struct {
	ReadOnlyDBInstanceId []*string `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) GetReadOnlyDBInstanceId() []*string {
	return s.ReadOnlyDBInstanceId
}

func (s *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) SetReadOnlyDBInstanceId(v []*string) *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds {
	s.ReadOnlyDBInstanceId = v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstanceResponseBodyDataVips struct {
	Vip []*DescribeDrdsInstanceResponseBodyDataVipsVip `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceResponseBodyDataVips) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyDataVips) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyDataVips) GetVip() []*DescribeDrdsInstanceResponseBodyDataVipsVip {
	return s.Vip
}

func (s *DescribeDrdsInstanceResponseBodyDataVips) SetVip(v []*DescribeDrdsInstanceResponseBodyDataVipsVip) *DescribeDrdsInstanceResponseBodyDataVips {
	s.Vip = v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVips) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstanceResponseBodyDataVipsVip struct {
	// The domain name that is mapped to the VIP.
	//
	// example:
	//
	// drdssen1243as.drds.aliyuncs.com
	Dns *string `json:"Dns,omitempty" xml:"Dns,omitempty"`
	// The number of remaining days before the VIP expires.
	//
	// example:
	//
	// 0
	ExpireDays *int64 `json:"ExpireDays,omitempty" xml:"ExpireDays,omitempty"`
	// The ports that are opened on the VIP.
	//
	// example:
	//
	// 3306
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	RemoveWeight *bool   `json:"RemoveWeight,omitempty" xml:"RemoveWeight,omitempty"`
	// The type of the VIP. Valid values: intranet and internet.
	//
	// example:
	//
	// intranet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp**********
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp***********
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeDrdsInstanceResponseBodyDataVipsVip) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyDataVipsVip) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) GetDns() *string {
	return s.Dns
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) GetExpireDays() *int64 {
	return s.ExpireDays
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) GetPort() *string {
	return s.Port
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) GetRemoveWeight() *bool {
	return s.RemoveWeight
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) GetType() *string {
	return s.Type
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetDns(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.Dns = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetExpireDays(v int64) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.ExpireDays = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetPort(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.Port = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetRemoveWeight(v bool) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.RemoveWeight = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetType(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetVpcId(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.VpcId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetVswitchId(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.VswitchId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) Validate() error {
	return dara.Validate(s)
}
