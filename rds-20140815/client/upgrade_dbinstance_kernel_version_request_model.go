// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceKernelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBInstanceKernelVersionRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpgradeDBInstanceKernelVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *UpgradeDBInstanceKernelVersionRequest
	GetSwitchTime() *string
	SetTargetMinorVersion(v string) *UpgradeDBInstanceKernelVersionRequest
	GetTargetMinorVersion() *string
	SetUpgradeTime(v string) *UpgradeDBInstanceKernelVersionRequest
	GetUpgradeTime() *string
}

type UpgradeDBInstanceKernelVersionRequest struct {
	// The ID of the instance. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// > 	- The storage type of the ApsaraDB RDS for PostgreSQL instance must be **cloud disk**. You can call the [restart the instance](https://help.aliyun.com/document_detail/26230.html) operation to restart an instance with Premium Local SSDs to automatically upgrade the instance to the latest minor engine version.
	//
	// > 	- Only ApsaraDB RDS for SQL Server 2019 instances support minor engine version upgrades.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The specified time when the upgrade is performed. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// >This parameter takes effect only when **UpgradeTime*	- is set to **SpecifyTime**.
	//
	// example:
	//
	// 2020-01-15T00:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The target minor database version. Format:
	//
	// 	- **PostgreSQL**: `rds_postgres_<Major version number>00_<Minor version number>`. Example: `rds_postgres_1200_20200830` for PostgreSQL 12 version 20200830.
	//
	// 	- **MySQL**: `<Instance edition>_<Minor version number>`. Examples: `rds_20200229`, `xcluster_20200229`, and `xcluster80_20200229`. The instance editions are as follows:
	//
	//     	- **rds**: high availability series or Basic Edition.
	//
	//     	- **xcluster**: MySQL 5.7 RDS Enterprise Edition.
	//
	//     	- **xcluster80**: MySQL 8.0 RDS Enterprise Edition.
	//
	// 	- **SQL Server**: `<Minor version number>`. Example: `15.0.4073.23`.
	//
	// If you do not specify this parameter, the instance is upgraded to the latest minor database version by default.
	//
	// > For more information about minor database versions, see [Release notes for minor engine versions of ApsaraDB RDS for PostgreSQL](https://help.aliyun.com/document_detail/126002.html), [Release notes for minor engine versions of ApsaraDB RDS for MySQL](https://help.aliyun.com/document_detail/96060.html), and [Release notes for minor engine versions of ApsaraDB RDS for SQL Server](https://help.aliyun.com/document_detail/213577.html).
	//
	// example:
	//
	// xcluster80_20210305
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The time when the upgrade is performed. Valid values:
	//
	// 	- **Immediate**: The upgrade is immediately performed. This is the default value.
	//
	// 	- **MaintainTime**: The upgrade is performed within the specified maintenance window. For more information about how to modify the maintenance window, see ModifyDBInstanceMaintainTime.
	//
	// 	- **SpecifyTime**: The upgrade is performed at the specified time.
	//
	// example:
	//
	// Immediate
	UpgradeTime *string `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty"`
}

func (s UpgradeDBInstanceKernelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetUpgradeTime() *string {
	return s.UpgradeTime
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetDBInstanceId(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetResourceOwnerAccount(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetSwitchTime(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.SwitchTime = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetTargetMinorVersion(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetUpgradeTime(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.UpgradeTime = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) Validate() error {
	return dara.Validate(s)
}
