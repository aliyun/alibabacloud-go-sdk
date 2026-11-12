// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseInstances(v []*DescribeDatabaseInstancesResponseBodyDatabaseInstances) *DescribeDatabaseInstancesResponseBody
	GetDatabaseInstances() []*DescribeDatabaseInstancesResponseBodyDatabaseInstances
	SetPageNumber(v int32) *DescribeDatabaseInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDatabaseInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDatabaseInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDatabaseInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeDatabaseInstancesResponseBody struct {
	// The details of the lightweight database instances.
	DatabaseInstances []*DescribeDatabaseInstancesResponseBodyDatabaseInstances `json:"DatabaseInstances,omitempty" xml:"DatabaseInstances,omitempty" type:"Repeated"`
	// The page number for a paged query.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in the lightweight database instance list.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatabaseInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstancesResponseBody) GetDatabaseInstances() []*DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	return s.DatabaseInstances
}

func (s *DescribeDatabaseInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatabaseInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatabaseInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatabaseInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDatabaseInstancesResponseBody) SetDatabaseInstances(v []*DescribeDatabaseInstancesResponseBodyDatabaseInstances) *DescribeDatabaseInstancesResponseBody {
	s.DatabaseInstances = v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) SetPageNumber(v int32) *DescribeDatabaseInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) SetPageSize(v int32) *DescribeDatabaseInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) SetRequestId(v string) *DescribeDatabaseInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) SetTotalCount(v int32) *DescribeDatabaseInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) Validate() error {
	if s.DatabaseInstances != nil {
		for _, item := range s.DatabaseInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDatabaseInstancesResponseBodyDatabaseInstances struct {
	// The business status of the instance. Valid values:
	//
	// - normal: Normal.
	//
	// - expired: Expired.
	//
	// - overdue: Overdue payment.
	//
	// example:
	//
	// normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The billing method of the lightweight database instance. Only the subscription billing method is supported. Set the value to PrePaid.
	//
	// Default value: PrePaid.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 1
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the instance was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format and is displayed in UTC+0.
	//
	// example:
	//
	// 2022-09-01T02:39:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The plan ID of the lightweight database instance. Valid values:
	//
	// - swas.db.c1m1s25: CNY 35 per month.
	//
	// - swas.db.c1m2s80: CNY 100 per month.
	//
	// - swas.db.c2m4s120: CNY 200 per month.
	//
	// - swas.db.c2m8s240: CNY 300 per month.
	//
	// example:
	//
	// swas.db.c1m1s25
	DatabaseInstanceEdition *string `json:"DatabaseInstanceEdition,omitempty" xml:"DatabaseInstanceEdition,omitempty"`
	// The instance ID of the lightweight database.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The name of the lightweight database instance.
	//
	// example:
	//
	// My database
	DatabaseInstanceName *string `json:"DatabaseInstanceName,omitempty" xml:"DatabaseInstanceName,omitempty"`
	// The status of the lightweight database instance. Valid values:
	//
	// - Pending: The instance is being created.
	//
	// - Restarting: The instance is being restarted.
	//
	// - Running: The instance is running.
	//
	// - Stopping: The instance is being stopped.
	//
	// - Stopped: The instance is stopped.
	//
	// - UPGRADING: The instance is being upgraded.
	//
	// - DISABLED: The instance is disabled.
	//
	// example:
	//
	// Running
	DatabaseInstanceStatus *string `json:"DatabaseInstanceStatus,omitempty" xml:"DatabaseInstanceStatus,omitempty"`
	// The database version. Valid values:
	//
	// - 5.7: MySQL 5.7.
	//
	// - 8.0: MySQL 8.0.
	//
	// example:
	//
	// 5.7
	DatabaseVersion *string `json:"DatabaseVersion,omitempty" xml:"DatabaseVersion,omitempty"`
	// The time when the instance expires. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format and is displayed in UTC+0.
	//
	// > The expiration time displayed in the console is 8 hours ahead of the UTC time.
	//
	// example:
	//
	// 2022-10-01T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 1
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// rm-bp1d39opj7906****.mysql.rds.aliyuncs.com
	PrivateConnection *string `json:"PrivateConnection,omitempty" xml:"PrivateConnection,omitempty"`
	// The public endpoint.
	//
	// > This parameter is returned only if a public endpoint is applied for. You can call the [AllocatePublicConnection](https://help.aliyun.com/document_detail/451413.html) operation to apply for a public endpoint.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****.mysql.rds.aliyuncs.com
	PublicConnection *string `json:"PublicConnection,omitempty" xml:"PublicConnection,omitempty"`
	// The region ID of the lightweight database instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The capacity of the ESSD cloud disk. Unit: GB.
	//
	// example:
	//
	// 25
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The name of the privileged account of the lightweight database instance.
	//
	// example:
	//
	// administrator
	SuperAccountName *string `json:"SuperAccountName,omitempty" xml:"SuperAccountName,omitempty"`
}

func (s DescribeDatabaseInstancesResponseBodyDatabaseInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseInstancesResponseBodyDatabaseInstances) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetCpu() *string {
	return s.Cpu
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetDatabaseInstanceEdition() *string {
	return s.DatabaseInstanceEdition
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetDatabaseInstanceName() *string {
	return s.DatabaseInstanceName
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetDatabaseInstanceStatus() *string {
	return s.DatabaseInstanceStatus
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetDatabaseVersion() *string {
	return s.DatabaseVersion
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetMemory() *string {
	return s.Memory
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetPrivateConnection() *string {
	return s.PrivateConnection
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetPublicConnection() *string {
	return s.PublicConnection
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetStorage() *int32 {
	return s.Storage
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) GetSuperAccountName() *string {
	return s.SuperAccountName
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetBusinessStatus(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetChargeType(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetCpu(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.Cpu = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetCreationTime(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.CreationTime = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseInstanceEdition(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseInstanceEdition = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseInstanceId(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseInstanceName(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseInstanceName = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseInstanceStatus(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseInstanceStatus = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseVersion(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseVersion = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetExpiredTime(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetMemory(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.Memory = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetPrivateConnection(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.PrivateConnection = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetPublicConnection(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.PublicConnection = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetRegionId(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetStorage(v int32) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.Storage = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetSuperAccountName(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.SuperAccountName = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) Validate() error {
	return dara.Validate(s)
}
