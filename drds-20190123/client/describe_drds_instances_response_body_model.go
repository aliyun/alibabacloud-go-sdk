// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeDrdsInstancesResponseBodyInstances) *DescribeDrdsInstancesResponseBody
	GetInstances() *DescribeDrdsInstancesResponseBodyInstances
	SetPageNumber(v int32) *DescribeDrdsInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDrdsInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDrdsInstancesResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeDrdsInstancesResponseBody
	GetTotal() *int32
}

type DescribeDrdsInstancesResponseBody struct {
	// The list of returned instances.
	Instances *DescribeDrdsInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8734773E-7B21-4A22-9106-CBD245F8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances returned.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBody) GetInstances() *DescribeDrdsInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeDrdsInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDrdsInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDrdsInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsInstancesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeDrdsInstancesResponseBody) SetInstances(v *DescribeDrdsInstancesResponseBodyInstances) *DescribeDrdsInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetPageNumber(v int32) *DescribeDrdsInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetPageSize(v int32) *DescribeDrdsInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetRequestId(v string) *DescribeDrdsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetTotal(v int32) *DescribeDrdsInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstancesResponseBodyInstances struct {
	Instance []*DescribeDrdsInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstances) GetInstance() []*DescribeDrdsInstancesResponseBodyInstancesInstance {
	return s.Instance
}

func (s *DescribeDrdsInstancesResponseBodyInstances) SetInstance(v []*DescribeDrdsInstancesResponseBodyInstancesInstance) *DescribeDrdsInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstancesResponseBodyInstancesInstance struct {
	// The commodity code of the service.
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
	// drdssen12****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The timestamp that indicates when the instance expires.
	//
	// example:
	//
	// 4724323200000
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The role of the instance. Valid values:
	//
	// 	- MASTER: The instance is a primary instance.
	//
	// 	- SLAVE: The instance is a read-only instance to analyze complex queries.
	//
	// 	- SLAVE_FLOW: The instance is a read-only instance for high-concurrency scenarios.
	//
	// example:
	//
	// MASTER
	InstRole *string `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
	// The instance series.
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
	// The machine type of the instance. Valid value: ecs.
	//
	// example:
	//
	// ecs
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The ID of the primary instance.
	//
	// example:
	//
	// drdssen12****
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **CLASSIC**
	//
	// 	- **VPC**
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
	// The version of the service.
	//
	// example:
	//
	// V1
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// The IDs of read-only instances that are associated with the instance.
	ReadOnlyDBInstanceIds *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds `json:"ReadOnlyDBInstanceIds,omitempty" xml:"ReadOnlyDBInstanceIds,omitempty" type:"Struct"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou-e
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aek2ljh3ye4****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// RUN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the instance. Valid values:
	//
	// 	- **PUBLIC**: The returned instance is a shared instance.
	//
	// 	- **PRIVATE**: The returned instance is a dedicated instance.
	//
	// example:
	//
	// PRIVATE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the instance.
	//
	// example:
	//
	// 0
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
	// Indicates whether the version of the instance can be upgraded.
	//
	// example:
	//
	// Upgradable
	VersionAction *string `json:"VersionAction,omitempty" xml:"VersionAction,omitempty"`
	// The list of returned virtual IP addresses (VIPs).
	Vips *DescribeDrdsInstancesResponseBodyInstancesInstanceVips `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
	// The ID of the instance that is deployed in the VPC.
	//
	// example:
	//
	// drdssen12****
	VpcCloudInstanceId *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	// The ID of the VPC to which the instance belongs.
	//
	// example:
	//
	// vpc-bp**********
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone in which the resource is located.
	//
	// example:
	//
	// vsw-bpxxxxxxxxxxxxx96
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The edition of the instance. Valid values:
	//
	// 	- **starter**: Starter Edition
	//
	// 	- **enterprise**: Enterprise Edition
	//
	// 	- **standard**: Standard Edition
	//
	// example:
	//
	// enterprise
	Series *string `json:"series,omitempty" xml:"series,omitempty"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetDescription() *string {
	return s.Description
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetExpireDate() *int64 {
	return s.ExpireDate
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetInstRole() *string {
	return s.InstRole
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetInstanceSeries() *string {
	return s.InstanceSeries
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetLabel() *string {
	return s.Label
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetMachineType() *string {
	return s.MachineType
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetMasterInstanceId() *string {
	return s.MasterInstanceId
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetOrderInstanceId() *string {
	return s.OrderInstanceId
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetReadOnlyDBInstanceIds() *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds {
	return s.ReadOnlyDBInstanceIds
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetType() *string {
	return s.Type
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetVersion() *int64 {
	return s.Version
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetVersionAction() *string {
	return s.VersionAction
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetVips() *DescribeDrdsInstancesResponseBodyInstancesInstanceVips {
	return s.Vips
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetVpcCloudInstanceId() *string {
	return s.VpcCloudInstanceId
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) GetSeries() *string {
	return s.Series
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetCommodityCode(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetCreateTime(v int64) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetDescription(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Description = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetDrdsInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetExpireDate(v int64) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetInstRole(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.InstRole = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetInstanceSeries(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.InstanceSeries = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetInstanceSpec(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetLabel(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Label = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetMachineType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.MachineType = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetMasterInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetNetworkType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetOrderInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.OrderInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetProductVersion(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetReadOnlyDBInstanceIds(v *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ReadOnlyDBInstanceIds = v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetResourceGroupId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVersion(v int64) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Version = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVersionAction(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.VersionAction = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVips(v *DescribeDrdsInstancesResponseBodyInstancesInstanceVips) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Vips = v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVpcCloudInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVpcId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetZoneId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetSeries(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Series = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds struct {
	ReadOnlyDBInstanceId []*string `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) GetReadOnlyDBInstanceId() []*string {
	return s.ReadOnlyDBInstanceId
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) SetReadOnlyDBInstanceId(v []*string) *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds {
	s.ReadOnlyDBInstanceId = v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstancesResponseBodyInstancesInstanceVips struct {
	Vip []*DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVips) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVips) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVips) GetVip() []*DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	return s.Vip
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVips) SetVip(v []*DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) *DescribeDrdsInstancesResponseBodyInstancesInstanceVips {
	s.Vip = v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVips) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip struct {
	// The virtual IP address.
	//
	// example:
	//
	// 10.23.***.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ports that are opened on the VIP.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the VIP. Valid values:
	//
	// 	- intranet: a private IP address
	//
	// 	- internet: a public IP address
	//
	// example:
	//
	// intranet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bpxxxxxxxxxxxy
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bpxxxxxxxxxxxxx96
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The domain name that is mapped to the VIP.
	//
	// example:
	//
	// drdssen1243as.drds.aliyuncs.com
	Dns *string `json:"dns,omitempty" xml:"dns,omitempty"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) GetIP() *string {
	return s.IP
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) GetPort() *string {
	return s.Port
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) GetType() *string {
	return s.Type
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) GetDns() *string {
	return s.Dns
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetIP(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.IP = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetPort(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.Port = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetVpcId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.VpcId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetVswitchId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.VswitchId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetDns(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.Dns = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) Validate() error {
	return dara.Validate(s)
}
