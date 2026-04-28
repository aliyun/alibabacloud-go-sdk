// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDriveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDriveRequest
	GetDescription() *string
	SetDriveId(v string) *UpdateDriveRequest
	GetDriveId() *string
	SetDriveName(v string) *UpdateDriveRequest
	GetDriveName() *string
	SetOwner(v string) *UpdateDriveRequest
	GetOwner() *string
	SetStatus(v string) *UpdateDriveRequest
	GetStatus() *string
	SetTotalSize(v int64) *UpdateDriveRequest
	GetTotalSize() *int64
}

type UpdateDriveRequest struct {
	// The description of the drive. The description can be up to 1,024 characters in length.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The name of the drive. The name can be up to 128 characters in length.
	//
	// example:
	//
	// my_drive
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	// The owner of the drive. Note: You can modify the owner of a personal drive only by using an AccessKey pair.
	//
	// example:
	//
	// user1
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The state of the drive. Valid values:
	//
	// enabled and disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total size of the drive. Unit: bytes. A value of -1 specifies that the size is unlimited.
	//
	// example:
	//
	// 10240
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s UpdateDriveRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDriveRequest) GoString() string {
	return s.String()
}

func (s *UpdateDriveRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDriveRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *UpdateDriveRequest) GetDriveName() *string {
	return s.DriveName
}

func (s *UpdateDriveRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateDriveRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateDriveRequest) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *UpdateDriveRequest) SetDescription(v string) *UpdateDriveRequest {
	s.Description = &v
	return s
}

func (s *UpdateDriveRequest) SetDriveId(v string) *UpdateDriveRequest {
	s.DriveId = &v
	return s
}

func (s *UpdateDriveRequest) SetDriveName(v string) *UpdateDriveRequest {
	s.DriveName = &v
	return s
}

func (s *UpdateDriveRequest) SetOwner(v string) *UpdateDriveRequest {
	s.Owner = &v
	return s
}

func (s *UpdateDriveRequest) SetStatus(v string) *UpdateDriveRequest {
	s.Status = &v
	return s
}

func (s *UpdateDriveRequest) SetTotalSize(v int64) *UpdateDriveRequest {
	s.TotalSize = &v
	return s
}

func (s *UpdateDriveRequest) Validate() error {
	return dara.Validate(s)
}
