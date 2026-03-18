// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdateBackupRequest
	GetRegionId() *string
	SetBackupTaskId(v string) *UpdateBackupRequest
	GetBackupTaskId() *string
	SetDescription(v string) *UpdateBackupRequest
	GetDescription() *string
}

type UpdateBackupRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// bk-t1232839
	BackupTaskId *string `json:"backupTaskId,omitempty" xml:"backupTaskId,omitempty"`
	// example:
	//
	// backupTask-desc1
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupRequest) GoString() string {
	return s.String()
}

func (s *UpdateBackupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateBackupRequest) GetBackupTaskId() *string {
	return s.BackupTaskId
}

func (s *UpdateBackupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateBackupRequest) SetRegionId(v string) *UpdateBackupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBackupRequest) SetBackupTaskId(v string) *UpdateBackupRequest {
	s.BackupTaskId = &v
	return s
}

func (s *UpdateBackupRequest) SetDescription(v string) *UpdateBackupRequest {
	s.Description = &v
	return s
}

func (s *UpdateBackupRequest) Validate() error {
	return dara.Validate(s)
}
