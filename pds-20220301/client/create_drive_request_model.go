// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDriveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefault(v bool) *CreateDriveRequest
	GetDefault() *bool
	SetDescription(v string) *CreateDriveRequest
	GetDescription() *string
	SetDriveName(v string) *CreateDriveRequest
	GetDriveName() *string
	SetDriveType(v string) *CreateDriveRequest
	GetDriveType() *string
	SetOwner(v string) *CreateDriveRequest
	GetOwner() *string
	SetOwnerType(v string) *CreateDriveRequest
	GetOwnerType() *string
	SetStatus(v string) *CreateDriveRequest
	GetStatus() *string
	SetTotalSize(v int64) *CreateDriveRequest
	GetTotalSize() *int64
}

type CreateDriveRequest struct {
	// Specifies whether the drive is the default drive. Default value: false.
	//
	// example:
	//
	// true
	Default *bool `json:"default,omitempty" xml:"default,omitempty"`
	// The description of the drive. The description can be up to 1,024 characters in length.
	//
	// example:
	//
	// drive for test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the drive. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_drive
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	// The type of the drive. Set the value to normal.
	//
	// example:
	//
	// normal
	DriveType *string `json:"drive_type,omitempty" xml:"drive_type,omitempty"`
	// The owner of the drive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3b3d7245c159488da17d081ad6c64687
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The type of the owner. Valid values:
	//
	// user and group.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	OwnerType *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	// The state of the drive. Valid values:
	//
	// enabled and disabled.
	//
	// Default value: enabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total size of the drive. Unit: bytes. By default, the size is unlimited.
	//
	// example:
	//
	// 1024
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s CreateDriveRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDriveRequest) GoString() string {
	return s.String()
}

func (s *CreateDriveRequest) GetDefault() *bool {
	return s.Default
}

func (s *CreateDriveRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDriveRequest) GetDriveName() *string {
	return s.DriveName
}

func (s *CreateDriveRequest) GetDriveType() *string {
	return s.DriveType
}

func (s *CreateDriveRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateDriveRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *CreateDriveRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateDriveRequest) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *CreateDriveRequest) SetDefault(v bool) *CreateDriveRequest {
	s.Default = &v
	return s
}

func (s *CreateDriveRequest) SetDescription(v string) *CreateDriveRequest {
	s.Description = &v
	return s
}

func (s *CreateDriveRequest) SetDriveName(v string) *CreateDriveRequest {
	s.DriveName = &v
	return s
}

func (s *CreateDriveRequest) SetDriveType(v string) *CreateDriveRequest {
	s.DriveType = &v
	return s
}

func (s *CreateDriveRequest) SetOwner(v string) *CreateDriveRequest {
	s.Owner = &v
	return s
}

func (s *CreateDriveRequest) SetOwnerType(v string) *CreateDriveRequest {
	s.OwnerType = &v
	return s
}

func (s *CreateDriveRequest) SetStatus(v string) *CreateDriveRequest {
	s.Status = &v
	return s
}

func (s *CreateDriveRequest) SetTotalSize(v int64) *CreateDriveRequest {
	s.TotalSize = &v
	return s
}

func (s *CreateDriveRequest) Validate() error {
	return dara.Validate(s)
}
