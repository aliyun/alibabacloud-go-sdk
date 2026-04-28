// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDriveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *CreateDriveResponseBody
	GetCreatedAt() *string
	SetCreator(v string) *CreateDriveResponseBody
	GetCreator() *string
	SetDescription(v string) *CreateDriveResponseBody
	GetDescription() *string
	SetDomainId(v string) *CreateDriveResponseBody
	GetDomainId() *string
	SetDriveId(v string) *CreateDriveResponseBody
	GetDriveId() *string
	SetDriveName(v string) *CreateDriveResponseBody
	GetDriveName() *string
	SetDriveType(v string) *CreateDriveResponseBody
	GetDriveType() *string
	SetOwner(v string) *CreateDriveResponseBody
	GetOwner() *string
	SetOwnerType(v string) *CreateDriveResponseBody
	GetOwnerType() *string
	SetStatus(v string) *CreateDriveResponseBody
	GetStatus() *string
	SetTotalSize(v int64) *CreateDriveResponseBody
	GetTotalSize() *int64
	SetUsedSize(v int64) *CreateDriveResponseBody
	GetUsedSize() *int64
}

type CreateDriveResponseBody struct {
	CreatedAt   *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator     *string `json:"creator,omitempty" xml:"creator,omitempty"`
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
	DriveId   *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	DriveType *string `json:"drive_type,omitempty" xml:"drive_type,omitempty"`
	Owner     *string `json:"owner,omitempty" xml:"owner,omitempty"`
	OwnerType *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	TotalSize *int64  `json:"total_size,omitempty" xml:"total_size,omitempty"`
	UsedSize  *int64  `json:"used_size,omitempty" xml:"used_size,omitempty"`
}

func (s CreateDriveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDriveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDriveResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateDriveResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *CreateDriveResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateDriveResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *CreateDriveResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *CreateDriveResponseBody) GetDriveName() *string {
	return s.DriveName
}

func (s *CreateDriveResponseBody) GetDriveType() *string {
	return s.DriveType
}

func (s *CreateDriveResponseBody) GetOwner() *string {
	return s.Owner
}

func (s *CreateDriveResponseBody) GetOwnerType() *string {
	return s.OwnerType
}

func (s *CreateDriveResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateDriveResponseBody) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *CreateDriveResponseBody) GetUsedSize() *int64 {
	return s.UsedSize
}

func (s *CreateDriveResponseBody) SetCreatedAt(v string) *CreateDriveResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateDriveResponseBody) SetCreator(v string) *CreateDriveResponseBody {
	s.Creator = &v
	return s
}

func (s *CreateDriveResponseBody) SetDescription(v string) *CreateDriveResponseBody {
	s.Description = &v
	return s
}

func (s *CreateDriveResponseBody) SetDomainId(v string) *CreateDriveResponseBody {
	s.DomainId = &v
	return s
}

func (s *CreateDriveResponseBody) SetDriveId(v string) *CreateDriveResponseBody {
	s.DriveId = &v
	return s
}

func (s *CreateDriveResponseBody) SetDriveName(v string) *CreateDriveResponseBody {
	s.DriveName = &v
	return s
}

func (s *CreateDriveResponseBody) SetDriveType(v string) *CreateDriveResponseBody {
	s.DriveType = &v
	return s
}

func (s *CreateDriveResponseBody) SetOwner(v string) *CreateDriveResponseBody {
	s.Owner = &v
	return s
}

func (s *CreateDriveResponseBody) SetOwnerType(v string) *CreateDriveResponseBody {
	s.OwnerType = &v
	return s
}

func (s *CreateDriveResponseBody) SetStatus(v string) *CreateDriveResponseBody {
	s.Status = &v
	return s
}

func (s *CreateDriveResponseBody) SetTotalSize(v int64) *CreateDriveResponseBody {
	s.TotalSize = &v
	return s
}

func (s *CreateDriveResponseBody) SetUsedSize(v int64) *CreateDriveResponseBody {
	s.UsedSize = &v
	return s
}

func (s *CreateDriveResponseBody) Validate() error {
	return dara.Validate(s)
}
