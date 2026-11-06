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
	Items *DescribeDBInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 2
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 03C46A7A-E680-5B9F-B890-38789B52EB50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
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
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBClusterId           *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBInstanceClass       *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStatus      *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceType        *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	DBType                *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion             *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ExpireTime            *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceNetworkType   *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	LockMode              *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason            *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	PayType               *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                 *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBClusterId() *string {
	return s.DBClusterId
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

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetExpireTime() *string {
	return s.ExpireTime
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

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) GetRegionId() *string {
	return s.RegionId
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

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetCreateTime(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBClusterId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBClusterId = &v
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

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceStatus(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBVersion(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetEngine(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetExpireTime(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ExpireTime = &v
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

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetPayType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.RegionId = &v
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
	return dara.Validate(s)
}
