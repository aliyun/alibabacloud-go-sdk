// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigBackupRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *ConfigBackupRemarkRequest
	GetBackupId() *string
	SetName(v string) *ConfigBackupRemarkRequest
	GetName() *string
	SetRemark(v string) *ConfigBackupRemarkRequest
	GetRemark() *string
}

type ConfigBackupRemarkRequest struct {
	// The ID of the backup.
	//
	// This parameter is required.
	//
	// example:
	//
	// backup-fdb897sdfg****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The name of the backup.
	//
	// example:
	//
	// backup-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the backup.
	//
	// example:
	//
	// Test backup.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ConfigBackupRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigBackupRemarkRequest) GoString() string {
	return s.String()
}

func (s *ConfigBackupRemarkRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *ConfigBackupRemarkRequest) GetName() *string {
	return s.Name
}

func (s *ConfigBackupRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *ConfigBackupRemarkRequest) SetBackupId(v string) *ConfigBackupRemarkRequest {
	s.BackupId = &v
	return s
}

func (s *ConfigBackupRemarkRequest) SetName(v string) *ConfigBackupRemarkRequest {
	s.Name = &v
	return s
}

func (s *ConfigBackupRemarkRequest) SetRemark(v string) *ConfigBackupRemarkRequest {
	s.Remark = &v
	return s
}

func (s *ConfigBackupRemarkRequest) Validate() error {
	return dara.Validate(s)
}
