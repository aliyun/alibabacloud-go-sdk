// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *ListBackupsRequest
	GetBackupId() *string
	SetCurrentPage(v int64) *ListBackupsRequest
	GetCurrentPage() *int64
	SetInstanceId(v string) *ListBackupsRequest
	GetInstanceId() *string
	SetName(v string) *ListBackupsRequest
	GetName() *string
	SetPageSize(v int64) *ListBackupsRequest
	GetPageSize() *int64
	SetRegionId(v string) *ListBackupsRequest
	GetRegionId() *string
}

type ListBackupsRequest struct {
	// The ID of the backup.
	//
	// example:
	//
	// backup-1648438****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the hardware security module (HSM).
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the backup.
	//
	// example:
	//
	// hsm-te****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBackupsRequest) GoString() string {
	return s.String()
}

func (s *ListBackupsRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *ListBackupsRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListBackupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBackupsRequest) GetName() *string {
	return s.Name
}

func (s *ListBackupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListBackupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBackupsRequest) SetBackupId(v string) *ListBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *ListBackupsRequest) SetCurrentPage(v int64) *ListBackupsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListBackupsRequest) SetInstanceId(v string) *ListBackupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBackupsRequest) SetName(v string) *ListBackupsRequest {
	s.Name = &v
	return s
}

func (s *ListBackupsRequest) SetPageSize(v int64) *ListBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListBackupsRequest) SetRegionId(v string) *ListBackupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListBackupsRequest) Validate() error {
	return dara.Validate(s)
}
