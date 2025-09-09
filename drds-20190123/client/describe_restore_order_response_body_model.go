// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRestoreOrderResponseBody
	GetRequestId() *string
	SetRestoreOrderDO(v *DescribeRestoreOrderResponseBodyRestoreOrderDO) *DescribeRestoreOrderResponseBody
	GetRestoreOrderDO() *DescribeRestoreOrderResponseBodyRestoreOrderDO
	SetSuccess(v bool) *DescribeRestoreOrderResponseBody
	GetSuccess() *bool
}

type DescribeRestoreOrderResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0AD2DE5D-B86B-40B5-9678-487D37******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned data object.
	RestoreOrderDO *DescribeRestoreOrderResponseBodyRestoreOrderDO `json:"RestoreOrderDO,omitempty" xml:"RestoreOrderDO,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRestoreOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestoreOrderResponseBody) GetRestoreOrderDO() *DescribeRestoreOrderResponseBodyRestoreOrderDO {
	return s.RestoreOrderDO
}

func (s *DescribeRestoreOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRestoreOrderResponseBody) SetRequestId(v string) *DescribeRestoreOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBody) SetRestoreOrderDO(v *DescribeRestoreOrderResponseBodyRestoreOrderDO) *DescribeRestoreOrderResponseBody {
	s.RestoreOrderDO = v
	return s
}

func (s *DescribeRestoreOrderResponseBody) SetSuccess(v bool) *DescribeRestoreOrderResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRestoreOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreOrderResponseBodyRestoreOrderDO struct {
	// The information of the restored order.
	DrdsOrderDOList *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList `json:"DrdsOrderDOList,omitempty" xml:"DrdsOrderDOList,omitempty" type:"Struct"`
	// The ID of the restored apsaradb for PolarDB cluster.
	PolarOrderDOList *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList `json:"PolarOrderDOList,omitempty" xml:"PolarOrderDOList,omitempty" type:"Struct"`
	// The information of the restored RDS order.
	RdsOrderDOList *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList `json:"RdsOrderDOList,omitempty" xml:"RdsOrderDOList,omitempty" type:"Struct"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDO) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDO) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) GetDrdsOrderDOList() *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList {
	return s.DrdsOrderDOList
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) GetPolarOrderDOList() *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList {
	return s.PolarOrderDOList
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) GetRdsOrderDOList() *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList {
	return s.RdsOrderDOList
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) SetDrdsOrderDOList(v *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDO {
	s.DrdsOrderDOList = v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) SetPolarOrderDOList(v *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDO {
	s.PolarOrderDOList = v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) SetRdsOrderDOList(v *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDO {
	s.RdsOrderDOList = v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList struct {
	DrdsOrderDOList []*DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList `json:"DrdsOrderDOList,omitempty" xml:"DrdsOrderDOList,omitempty" type:"Repeated"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) GetDrdsOrderDOList() []*DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	return s.DrdsOrderDOList
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) SetDrdsOrderDOList(v []*DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList {
	s.DrdsOrderDOList = v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList struct {
	// The ID of the zone for which to query resources.
	//
	// example:
	//
	// cn-hangzhou-e
	AzoneId *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	// The instance type of the instance.
	//
	// example:
	//
	// 4C8G 	- 2
	InstSpec *string `json:"InstSpec,omitempty" xml:"InstSpec,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Classic **: Classic Network
	//
	// 	- **vpc**: VPC
	//
	// example:
	//
	// vpc
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the vSwitch in the VPC.
	//
	// example:
	//
	// vsw-*******************
	VSwtichId *string `json:"VSwtichId,omitempty" xml:"VSwtichId,omitempty"`
	// The ID of the VPC network.
	//
	// example:
	//
	// vpc-*******************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) GetAzoneId() *string {
	return s.AzoneId
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) GetInstSpec() *string {
	return s.InstSpec
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) GetNetwork() *string {
	return s.Network
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) GetVSwtichId() *string {
	return s.VSwtichId
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetAzoneId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.AzoneId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetInstSpec(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.InstSpec = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetNetwork(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.Network = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetRegionId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.RegionId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetVSwtichId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.VSwtichId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetVpcId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.VpcId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList struct {
	PolarOrderDOList []*DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList `json:"PolarOrderDOList,omitempty" xml:"PolarOrderDOList,omitempty" type:"Repeated"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) GetPolarOrderDOList() []*DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	return s.PolarOrderDOList
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) SetPolarOrderDOList(v []*DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList {
	s.PolarOrderDOList = v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList struct {
	// The zone ID of the node.
	//
	// example:
	//
	// cn-hangzhou-g
	AzoneId *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	// The capacity of disk.
	//
	// example:
	//
	// 10240
	DbInstanceStorage *string `json:"DbInstanceStorage,omitempty" xml:"DbInstanceStorage,omitempty"`
	// The storage engine of PolarDB.
	//
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// polar.mysql.x4.large
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Classic**: Classic Network
	//
	// 	- **vpc**: VPC
	//
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The number of streams that were returned.
	//
	// example:
	//
	// 1
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The version of the operating system.
	//
	// example:
	//
	// 5.6
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) GetAzoneId() *string {
	return s.AzoneId
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) GetDbInstanceStorage() *string {
	return s.DbInstanceStorage
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) GetEngine() *string {
	return s.Engine
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) GetNetwork() *string {
	return s.Network
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) GetNum() *int64 {
	return s.Num
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) GetVersion() *string {
	return s.Version
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetAzoneId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.AzoneId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetDbInstanceStorage(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.DbInstanceStorage = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetEngine(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Engine = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetInstanceClass(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.InstanceClass = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetNetwork(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Network = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetNum(v int64) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Num = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetRegionId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.RegionId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetVersion(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Version = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList struct {
	RdsOrderDOList []*DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList `json:"RdsOrderDOList,omitempty" xml:"RdsOrderDOList,omitempty" type:"Repeated"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) GetRdsOrderDOList() []*DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	return s.RdsOrderDOList
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) SetRdsOrderDOList(v []*DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList {
	s.RdsOrderDOList = v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList struct {
	// The zone ID of the node.
	//
	// example:
	//
	// cn-hangzhou-g
	AzoneId *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	// The capacity of disk.
	//
	// example:
	//
	// 10240
	DbInstanceStorage *string `json:"DbInstanceStorage,omitempty" xml:"DbInstanceStorage,omitempty"`
	// The storage engine of the instance.
	//
	// example:
	//
	// MYSQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The instance type of the instance.
	//
	// example:
	//
	// rds.mysql.s2.large
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The network type. Valid values: - **Classic **: Classic Network
	//
	// - **vpc**: VPC
	//
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The number of streams that were returned.
	//
	// example:
	//
	// 1
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The version of the operating system.
	//
	// example:
	//
	// 5.6
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) GetAzoneId() *string {
	return s.AzoneId
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) GetDbInstanceStorage() *string {
	return s.DbInstanceStorage
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) GetEngine() *string {
	return s.Engine
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) GetNetwork() *string {
	return s.Network
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) GetNum() *int64 {
	return s.Num
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) GetVersion() *string {
	return s.Version
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetAzoneId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.AzoneId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetDbInstanceStorage(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.DbInstanceStorage = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetEngine(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Engine = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetInstanceClass(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.InstanceClass = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetNetwork(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Network = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetNum(v int64) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Num = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetRegionId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.RegionId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetVersion(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Version = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) Validate() error {
	return dara.Validate(s)
}
