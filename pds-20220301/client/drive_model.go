// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDrive interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *Drive
	GetCreatedAt() *string
	SetCreator(v string) *Drive
	GetCreator() *string
	SetDescription(v string) *Drive
	GetDescription() *string
	SetDomainId(v string) *Drive
	GetDomainId() *string
	SetDriveId(v string) *Drive
	GetDriveId() *string
	SetDriveName(v string) *Drive
	GetDriveName() *string
	SetDriveType(v string) *Drive
	GetDriveType() *string
	SetOwner(v string) *Drive
	GetOwner() *string
	SetOwnerType(v string) *Drive
	GetOwnerType() *string
	SetStatus(v string) *Drive
	GetStatus() *string
	SetTotalSize(v int64) *Drive
	GetTotalSize() *int64
	SetUsedSize(v int64) *Drive
	GetUsedSize() *int64
}

type Drive struct {
	// The time when the drive was created.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The user who created the drive.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The description of the drive.
	//
	// example:
	//
	// vipdrive
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The drive name.
	//
	// example:
	//
	// drv1
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	// The type of the drive.
	//
	// example:
	//
	// normal
	DriveType *string `json:"drive_type,omitempty" xml:"drive_type,omitempty"`
	// The owner of the drive.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The type of the owner.
	//
	// example:
	//
	// user
	OwnerType *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	// The status of the driver.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total storage space of the drive.
	//
	// example:
	//
	// 204800
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
	// The occupied storage space of the drive.
	//
	// example:
	//
	// 20480
	UsedSize *int64 `json:"used_size,omitempty" xml:"used_size,omitempty"`
}

func (s Drive) String() string {
	return dara.Prettify(s)
}

func (s Drive) GoString() string {
	return s.String()
}

func (s *Drive) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Drive) GetCreator() *string {
	return s.Creator
}

func (s *Drive) GetDescription() *string {
	return s.Description
}

func (s *Drive) GetDomainId() *string {
	return s.DomainId
}

func (s *Drive) GetDriveId() *string {
	return s.DriveId
}

func (s *Drive) GetDriveName() *string {
	return s.DriveName
}

func (s *Drive) GetDriveType() *string {
	return s.DriveType
}

func (s *Drive) GetOwner() *string {
	return s.Owner
}

func (s *Drive) GetOwnerType() *string {
	return s.OwnerType
}

func (s *Drive) GetStatus() *string {
	return s.Status
}

func (s *Drive) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *Drive) GetUsedSize() *int64 {
	return s.UsedSize
}

func (s *Drive) SetCreatedAt(v string) *Drive {
	s.CreatedAt = &v
	return s
}

func (s *Drive) SetCreator(v string) *Drive {
	s.Creator = &v
	return s
}

func (s *Drive) SetDescription(v string) *Drive {
	s.Description = &v
	return s
}

func (s *Drive) SetDomainId(v string) *Drive {
	s.DomainId = &v
	return s
}

func (s *Drive) SetDriveId(v string) *Drive {
	s.DriveId = &v
	return s
}

func (s *Drive) SetDriveName(v string) *Drive {
	s.DriveName = &v
	return s
}

func (s *Drive) SetDriveType(v string) *Drive {
	s.DriveType = &v
	return s
}

func (s *Drive) SetOwner(v string) *Drive {
	s.Owner = &v
	return s
}

func (s *Drive) SetOwnerType(v string) *Drive {
	s.OwnerType = &v
	return s
}

func (s *Drive) SetStatus(v string) *Drive {
	s.Status = &v
	return s
}

func (s *Drive) SetTotalSize(v int64) *Drive {
	s.TotalSize = &v
	return s
}

func (s *Drive) SetUsedSize(v int64) *Drive {
	s.UsedSize = &v
	return s
}

func (s *Drive) Validate() error {
	return dara.Validate(s)
}
