// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDbInstancesResponseBodyItems) *DescribeDbInstancesResponseBody
	GetItems() *DescribeDbInstancesResponseBodyItems
	SetRequestId(v string) *DescribeDbInstancesResponseBody
	GetRequestId() *string
}

type DescribeDbInstancesResponseBody struct {
	// The details of the instance.
	Items *DescribeDbInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 293275B3-8FC0-4619-A26E-6F062FASD56R
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDbInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBody) GetItems() *DescribeDbInstancesResponseBodyItems {
	return s.Items
}

func (s *DescribeDbInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDbInstancesResponseBody) SetItems(v *DescribeDbInstancesResponseBodyItems) *DescribeDbInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDbInstancesResponseBody) SetRequestId(v string) *DescribeDbInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDbInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDbInstancesResponseBodyItems struct {
	DBInstance []*DescribeDbInstancesResponseBodyItemsDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDbInstancesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBodyItems) GetDBInstance() []*DescribeDbInstancesResponseBodyItemsDBInstance {
	return s.DBInstance
}

func (s *DescribeDbInstancesResponseBodyItems) SetDBInstance(v []*DescribeDbInstancesResponseBodyItemsDBInstance) *DescribeDbInstancesResponseBodyItems {
	s.DBInstance = v
	return s
}

func (s *DescribeDbInstancesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDbInstancesResponseBodyItemsDBInstance struct {
	AllowAllCategory *bool `json:"AllowAllCategory,omitempty" xml:"AllowAllCategory,omitempty"`
	// The description of the storage instance.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The ID of the storage instance.
	//
	// example:
	//
	// rm-****************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Storage layer instance status. Valid values:
	//
	// 	- **0**: creating
	//
	// 	- **1**: In use
	//
	// 	- **3**: Deleting
	//
	// 	- **5**: restarting
	//
	// 	- **6**: upgrading /Downgrading
	//
	// 	- **7**: Recovering
	//
	// 	- **8**: switching the Internet and intranet
	//
	// example:
	//
	// 0
	DBInstanceStatus *int32 `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The storage layer instance type.
	//
	// example:
	//
	// Primary
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The engine of the storage instance.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the engine for the storage instance.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The network type of the storage layer. Valid values:
	//
	// 	- **VPC**: VPC
	//
	// 	- **CLASSIC **: Classic Network
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The details about a read-only storage instance.
	ReadOnlyDBInstanceId *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Struct"`
	// The ID of the region where the storage instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the zone where the storage instance resides.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDbInstancesResponseBodyItemsDBInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) GetAllowAllCategory() *bool {
	return s.AllowAllCategory
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) GetDBInstanceStatus() *int32 {
	return s.DBInstanceStatus
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) GetReadOnlyDBInstanceId() *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId {
	return s.ReadOnlyDBInstanceId
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetAllowAllCategory(v bool) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.AllowAllCategory = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceDescription(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceId(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceStatus(v int32) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceType(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetEngine(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetEngineVersion(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetInstanceNetworkType(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetReadOnlyDBInstanceId(v *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.ReadOnlyDBInstanceId = v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) Validate() error {
	return dara.Validate(s)
}

type DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId struct {
	ReadOnlyDBInstanceId []*string `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) GetReadOnlyDBInstanceId() []*string {
	return s.ReadOnlyDBInstanceId
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) SetReadOnlyDBInstanceId(v []*string) *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId {
	s.ReadOnlyDBInstanceId = v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) Validate() error {
	return dara.Validate(s)
}
