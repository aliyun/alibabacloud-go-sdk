// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupTaskId(v string) *DeleteBackupRequest
	GetBackupTaskId() *string
	SetInstanceId(v string) *DeleteBackupRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteBackupRequest
	GetRegionId() *string
}

type DeleteBackupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bk-yi2378sdhf2
	BackupTaskId *string `json:"BackupTaskId,omitempty" xml:"BackupTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupRequest) GetBackupTaskId() *string {
	return s.BackupTaskId
}

func (s *DeleteBackupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteBackupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBackupRequest) SetBackupTaskId(v string) *DeleteBackupRequest {
	s.BackupTaskId = &v
	return s
}

func (s *DeleteBackupRequest) SetInstanceId(v string) *DeleteBackupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBackupRequest) SetRegionId(v string) *DeleteBackupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBackupRequest) Validate() error {
	return dara.Validate(s)
}
