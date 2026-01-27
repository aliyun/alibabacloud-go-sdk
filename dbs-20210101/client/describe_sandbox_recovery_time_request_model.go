// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSandboxRecoveryTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DescribeSandboxRecoveryTimeRequest
	GetBackupPlanId() *string
}

type DescribeSandboxRecoveryTimeRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/437215.html) operation to obtain the ID of the backup schedule. If you set this parameter to the backup schedule ID obtained by calling the DescribeBackupPlanList operation, the dbs prefix must be removed. Otherwise, the call fails.
	//
	// > If your instance is an ApsaraDB RDS for MySQL instance, you can [configure automatic access to a data source](https://help.aliyun.com/document_detail/193091.html) to automatically add the instance to DBS and obtain the ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1jyjal15l****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
}

func (s DescribeSandboxRecoveryTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxRecoveryTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSandboxRecoveryTimeRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeSandboxRecoveryTimeRequest) SetBackupPlanId(v string) *DescribeSandboxRecoveryTimeRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeRequest) Validate() error {
	return dara.Validate(s)
}
