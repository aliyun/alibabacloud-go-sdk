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
	// The instance ID. You can invoke DescribeDBInstances to query the instance ID.
	//
	// > 	- The storage type of the ApsaraDB RDS for PostgreSQL instance must be **cloud disks**. For an instance with Premium Local SSDs, you can invoke the [RestartDBInstance](https://help.aliyun.com/document_detail/26230.html) operation to restart the instance, which automatically upgrades the instance to the latest minor engine version.
	//
	// > 	- Only the 2019 version of ApsaraDB RDS for SQL Server supports minor engine version upgrades.
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
	// The specified time. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// > This parameter takes effect only when **UpgradeTime*	- is set to **SpecifyTime**.
	//
	// example:
	//
	// 2020-01-15T00:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The minor database engine version to which you want to upgrade. Format:
	//
	// 	- **PostgreSQL**: `rds_postgres_<Major version number>00_<Minor version number>`. Example for version 12 with minor version 20200830: `rds_postgres_1200_20200830`.
	//
	// 	- **MySQL**: `<Instance version>_<Minor version number>`. Examples: `rds_20200229`, `xcluster_20200229`, or `xcluster80_20200229`. The instance version can be one of the following:
	//
	//     	- **rds**: high-availability series or Basic Edition.
	//
	//     	- **xcluster**: MySQL 5.7 RDS Enterprise Edition.
	//
	//     	- **xcluster80**: MySQL 8.0 RDS Enterprise Edition.
	//
	// 	- **SQLServer**: `<Minor version number>`. Example: `15.0.4073.23`.
	//
	// If you do not specify this parameter, the instance is upgraded to the latest minor engine version by default.
	//
	// > For minor engine version numbers, see [Release notes of ApsaraDB RDS for PostgreSQL minor engine versions](https://help.aliyun.com/document_detail/126002.html), [Release notes of ApsaraDB RDS for MySQL minor engine versions](https://help.aliyun.com/document_detail/96060.html), and [Release notes of ApsaraDB RDS for SQL Server minor engine versions](https://help.aliyun.com/document_detail/213577.html).
	//
	// example:
	//
	// xcluster80_20210305
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The upgrade time. Valid values:
	//
	// 	- **Immediate*	- (default): The upgrade takes effect immediately.
	//
	// 	- **MaintainTime**: The upgrade takes effect during the maintenance window. To modify the maintenance window, call ModifyDBInstanceMaintainTime.
	//
	// 	- **SpecifyTime**: The upgrade takes effect at a specified time.
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
