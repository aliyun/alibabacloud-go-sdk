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
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// > 	- If your instance runs PostgreSQL, you must make sure that the instance uses **cloud disks**. If the instance uses local disks, you must call the [RestartDBInstance](https://help.aliyun.com/document_detail/26230.html) operation to restart the instance. The system automatically updates the minor engine version of the instance to the latest version during the restart.
	//
	// > 	- If your instance runs SQL Server, you must make sure that the instance runs SQL Server 2019.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bpxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The update time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// > This parameter takes effect only when you set **UpgradeTime*	- to **SpecifyTime**.
	//
	// example:
	//
	// 2020-01-15T00:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The minor engine version to which you want to update. Format:
	//
	// 	- **PostgreSQL**: `rds_postgres_<Major engine version>00_<Minor engine version>`. Example: `rds_postgres_1200_20200830`.
	//
	// 	- **MySQL**: `<RDS edition and MySQL version>_<Minor engine version>`. Examples: `rds_20200229`, `xcluster_20200229`, and `xcluster80_20200229`. The following RDS editions and MySQL versions are supported:
	//
	//     	- **rds**: RDS Basic Edition or RDS High-availability Edition
	//
	//     	- **xcluster**: MySQL 5.7 on RDS Enterprise Edition
	//
	//     	- **xcluster80**: MySQL 8.0 on RDS Enterprise Edition
	//
	// 	- **SQLServer**: `<Minor engine version>`. Example: `15.0.4073.23`.
	//
	// If you do not specify this parameter, the instance is updated to the latest minor engine version.
	//
	// >  For more information about minor engine versions, see [Release notes of AliPG](https://help.aliyun.com/document_detail/126002.html), [Release notes of AliSQL](https://help.aliyun.com/document_detail/96060.html), and [Release notes of minor engine versions of ApsaraDB RDS for SQL Server](https://help.aliyun.com/document_detail/213577.html).
	//
	// example:
	//
	// xcluster80_20210305
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The time when the update takes effect. Valid values:
	//
	// 	- **Immediate*	- (default): The update takes effect immediately.
	//
	// 	- **MaintainTime**: The update takes effect during the maintenance window that you specify. For more information about how to change the maintenance window, see ModifyDBInstanceMaintainTime.
	//
	// 	- **SpecifyTime**: The update takes effect at the point in time you specify.
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
