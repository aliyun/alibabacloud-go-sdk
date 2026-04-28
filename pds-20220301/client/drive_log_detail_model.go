// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDriveLogDetail interface {
	dara.Model
	String() string
	GoString() string
	SetForceDelete(v bool) *DriveLogDetail
	GetForceDelete() *bool
	SetHandoverOwnerName(v string) *DriveLogDetail
	GetHandoverOwnerName() *string
	SetName(v string) *DriveLogDetail
	GetName() *string
	SetOwnerId(v string) *DriveLogDetail
	GetOwnerId() *string
	SetOwnerName(v string) *DriveLogDetail
	GetOwnerName() *string
	SetOwnerType(v string) *DriveLogDetail
	GetOwnerType() *string
	SetTotalSize(v int64) *DriveLogDetail
	GetTotalSize() *int64
	SetUpdateTo(v *DriveLogDetailUpdateTo) *DriveLogDetail
	GetUpdateTo() *DriveLogDetailUpdateTo
}

type DriveLogDetail struct {
	ForceDelete       *bool                   `json:"force_delete,omitempty" xml:"force_delete,omitempty"`
	HandoverOwnerName *string                 `json:"handover_owner_name,omitempty" xml:"handover_owner_name,omitempty"`
	Name              *string                 `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId           *string                 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
	OwnerName         *string                 `json:"owner_name,omitempty" xml:"owner_name,omitempty"`
	OwnerType         *string                 `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	TotalSize         *int64                  `json:"total_size,omitempty" xml:"total_size,omitempty"`
	UpdateTo          *DriveLogDetailUpdateTo `json:"update_to,omitempty" xml:"update_to,omitempty" type:"Struct"`
}

func (s DriveLogDetail) String() string {
	return dara.Prettify(s)
}

func (s DriveLogDetail) GoString() string {
	return s.String()
}

func (s *DriveLogDetail) GetForceDelete() *bool {
	return s.ForceDelete
}

func (s *DriveLogDetail) GetHandoverOwnerName() *string {
	return s.HandoverOwnerName
}

func (s *DriveLogDetail) GetName() *string {
	return s.Name
}

func (s *DriveLogDetail) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DriveLogDetail) GetOwnerName() *string {
	return s.OwnerName
}

func (s *DriveLogDetail) GetOwnerType() *string {
	return s.OwnerType
}

func (s *DriveLogDetail) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DriveLogDetail) GetUpdateTo() *DriveLogDetailUpdateTo {
	return s.UpdateTo
}

func (s *DriveLogDetail) SetForceDelete(v bool) *DriveLogDetail {
	s.ForceDelete = &v
	return s
}

func (s *DriveLogDetail) SetHandoverOwnerName(v string) *DriveLogDetail {
	s.HandoverOwnerName = &v
	return s
}

func (s *DriveLogDetail) SetName(v string) *DriveLogDetail {
	s.Name = &v
	return s
}

func (s *DriveLogDetail) SetOwnerId(v string) *DriveLogDetail {
	s.OwnerId = &v
	return s
}

func (s *DriveLogDetail) SetOwnerName(v string) *DriveLogDetail {
	s.OwnerName = &v
	return s
}

func (s *DriveLogDetail) SetOwnerType(v string) *DriveLogDetail {
	s.OwnerType = &v
	return s
}

func (s *DriveLogDetail) SetTotalSize(v int64) *DriveLogDetail {
	s.TotalSize = &v
	return s
}

func (s *DriveLogDetail) SetUpdateTo(v *DriveLogDetailUpdateTo) *DriveLogDetail {
	s.UpdateTo = v
	return s
}

func (s *DriveLogDetail) Validate() error {
	if s.UpdateTo != nil {
		if err := s.UpdateTo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DriveLogDetailUpdateTo struct {
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId   *string `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
	OwnerName *string `json:"owner_name,omitempty" xml:"owner_name,omitempty"`
	OwnerType *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	TotalSize *int64  `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s DriveLogDetailUpdateTo) String() string {
	return dara.Prettify(s)
}

func (s DriveLogDetailUpdateTo) GoString() string {
	return s.String()
}

func (s *DriveLogDetailUpdateTo) GetName() *string {
	return s.Name
}

func (s *DriveLogDetailUpdateTo) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DriveLogDetailUpdateTo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *DriveLogDetailUpdateTo) GetOwnerType() *string {
	return s.OwnerType
}

func (s *DriveLogDetailUpdateTo) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DriveLogDetailUpdateTo) SetName(v string) *DriveLogDetailUpdateTo {
	s.Name = &v
	return s
}

func (s *DriveLogDetailUpdateTo) SetOwnerId(v string) *DriveLogDetailUpdateTo {
	s.OwnerId = &v
	return s
}

func (s *DriveLogDetailUpdateTo) SetOwnerName(v string) *DriveLogDetailUpdateTo {
	s.OwnerName = &v
	return s
}

func (s *DriveLogDetailUpdateTo) SetOwnerType(v string) *DriveLogDetailUpdateTo {
	s.OwnerType = &v
	return s
}

func (s *DriveLogDetailUpdateTo) SetTotalSize(v int64) *DriveLogDetailUpdateTo {
	s.TotalSize = &v
	return s
}

func (s *DriveLogDetailUpdateTo) Validate() error {
	return dara.Validate(s)
}
