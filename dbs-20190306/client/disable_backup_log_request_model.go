// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBackupLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DisableBackupLogRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *DisableBackupLogRequest
	GetClientToken() *string
	SetDisableMysqlPhysicalBackupBinlogOnly(v string) *DisableBackupLogRequest
	GetDisableMysqlPhysicalBackupBinlogOnly() *string
	SetOwnerId(v string) *DisableBackupLogRequest
	GetOwnerId() *string
}

type DisableBackupLogRequest struct {
	// The ID of the backup plan. Call [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) to query it.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi01****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// Ensures idempotence and prevents duplicate requests.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCziJZNwH****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DisableMysqlPhysicalBackupBinlogOnly *string `json:"DisableMysqlPhysicalBackupBinlogOnly,omitempty" xml:"DisableMysqlPhysicalBackupBinlogOnly,omitempty"`
	OwnerId                              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DisableBackupLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableBackupLogRequest) GoString() string {
	return s.String()
}

func (s *DisableBackupLogRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DisableBackupLogRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableBackupLogRequest) GetDisableMysqlPhysicalBackupBinlogOnly() *string {
	return s.DisableMysqlPhysicalBackupBinlogOnly
}

func (s *DisableBackupLogRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DisableBackupLogRequest) SetBackupPlanId(v string) *DisableBackupLogRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DisableBackupLogRequest) SetClientToken(v string) *DisableBackupLogRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableBackupLogRequest) SetDisableMysqlPhysicalBackupBinlogOnly(v string) *DisableBackupLogRequest {
	s.DisableMysqlPhysicalBackupBinlogOnly = &v
	return s
}

func (s *DisableBackupLogRequest) SetOwnerId(v string) *DisableBackupLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableBackupLogRequest) Validate() error {
	return dara.Validate(s)
}
