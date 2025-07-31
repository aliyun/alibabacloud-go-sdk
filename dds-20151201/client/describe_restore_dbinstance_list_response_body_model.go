// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreDBInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstances(v *DescribeRestoreDBInstanceListResponseBodyDBInstances) *DescribeRestoreDBInstanceListResponseBody
	GetDBInstances() *DescribeRestoreDBInstanceListResponseBodyDBInstances
	SetPageNumber(v int32) *DescribeRestoreDBInstanceListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRestoreDBInstanceListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRestoreDBInstanceListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRestoreDBInstanceListResponseBody
	GetTotalCount() *int32
}

type DescribeRestoreDBInstanceListResponseBody struct {
	// DB instances list.
	DBInstances *DescribeRestoreDBInstanceListResponseBodyDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AF0AD89-ED4F-44AD-B65F-BFC1D5Cxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of instances in the query results.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRestoreDBInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreDBInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreDBInstanceListResponseBody) GetDBInstances() *DescribeRestoreDBInstanceListResponseBodyDBInstances {
	return s.DBInstances
}

func (s *DescribeRestoreDBInstanceListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRestoreDBInstanceListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreDBInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestoreDBInstanceListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRestoreDBInstanceListResponseBody) SetDBInstances(v *DescribeRestoreDBInstanceListResponseBodyDBInstances) *DescribeRestoreDBInstanceListResponseBody {
	s.DBInstances = v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBody) SetPageNumber(v int32) *DescribeRestoreDBInstanceListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBody) SetPageSize(v int32) *DescribeRestoreDBInstanceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBody) SetRequestId(v string) *DescribeRestoreDBInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBody) SetTotalCount(v int32) *DescribeRestoreDBInstanceListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreDBInstanceListResponseBodyDBInstances struct {
	DBInstance []*DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeRestoreDBInstanceListResponseBodyDBInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreDBInstanceListResponseBodyDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstances) GetDBInstance() []*DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	return s.DBInstance
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstances) SetDBInstance(v []*DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) *DescribeRestoreDBInstanceListResponseBodyDBInstances {
	s.DBInstance = v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance struct {
	// The time of instance creation, formatted as <i>yyyy-MM-dd</i>T<i>HH:00:00</i>Z (UTC time).
	//
	// example:
	//
	// 2022-01-02T07:43:59Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// test-database
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dds-bp12c5b040dc****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The status of the instance. For more information, see [Instance states](https://help.aliyun.com/document_detail/63870.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The architecture of the instance. Valid values:
	//
	// 	- **sharding**: sharded cluster instance
	//
	// 	- **replicate**: replica set or standalone instance
	//
	// example:
	//
	// replicate
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The database engine version of the instance. Valid values:
	//
	// 	- **7.0**
	//
	// 	- **6.0**
	//
	// 	- **5.0**
	//
	// 	- **4.4**
	//
	// 	- **4.2**
	//
	// 	- **4.0**
	//
	// 	- **3.4**
	//
	// example:
	//
	// 4.2
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The secondary availability zone 2 for the instance when implementing multi-AZ deployment.
	//
	// example:
	//
	// cn-hangzhou-h
	HiddenZoneId *string `json:"HiddenZoneId,omitempty" xml:"HiddenZoneId,omitempty"`
	// Specifies whether the instance is deleted. Valid values:
	//
	// 	- **0**: not deleted
	//
	// 	- **1**: deleted
	//
	// example:
	//
	// 0
	IsDeleted *int32 `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The locked state of the instance, value description:
	//
	// - Unlock: Normal.
	//
	// - ManualLock: Manually triggered lock.
	//
	// - LockByExpiration: Automatically locked due to expiration.
	//
	// - LockByRestoration: Automatically locked before restoration.
	//
	// - LockByDiskQuota: Automatically locked due to disk quota exceeded.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The secondary availability zone 1 for the instance when implementing multi-AZ deployment.
	//
	// example:
	//
	// cn-hangzhou-i
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GetHiddenZoneId() *string {
	return s.HiddenZoneId
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GetIsDeleted() *int32 {
	return s.IsDeleted
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) SetCreationTime(v string) *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) SetDBInstanceDescription(v string) *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) SetDBInstanceId(v string) *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) SetDBInstanceStatus(v string) *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) SetDBInstanceType(v string) *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) SetEngineVersion(v string) *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) SetHiddenZoneId(v string) *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	s.HiddenZoneId = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) SetIsDeleted(v int32) *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	s.IsDeleted = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) SetLockMode(v string) *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) SetRegionId(v string) *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) SetSecondaryZoneId(v string) *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	s.SecondaryZoneId = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) SetZoneId(v string) *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponseBodyDBInstancesDBInstance) Validate() error {
	return dara.Validate(s)
}
