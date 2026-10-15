// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppAiStaff interface {
	dara.Model
	String() string
	GoString() string
	SetStaffId(v string) *AppAiStaff
	GetStaffId() *string
	SetStaffName(v string) *AppAiStaff
	GetStaffName() *string
	SetStaffType(v string) *AppAiStaff
	GetStaffType() *string
	SetStatus(v string) *AppAiStaff
	GetStatus() *string
}

type AppAiStaff struct {
	// The employee ID.
	//
	// example:
	//
	// 370196
	StaffId *string `json:"StaffId,omitempty" xml:"StaffId,omitempty"`
	// The employee name.
	StaffName *string `json:"StaffName,omitempty" xml:"StaffName,omitempty"`
	// The employee type.
	StaffType *string `json:"StaffType,omitempty" xml:"StaffType,omitempty"`
	// Valid values: unknown, init, testing, and online.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AppAiStaff) String() string {
	return dara.Prettify(s)
}

func (s AppAiStaff) GoString() string {
	return s.String()
}

func (s *AppAiStaff) GetStaffId() *string {
	return s.StaffId
}

func (s *AppAiStaff) GetStaffName() *string {
	return s.StaffName
}

func (s *AppAiStaff) GetStaffType() *string {
	return s.StaffType
}

func (s *AppAiStaff) GetStatus() *string {
	return s.Status
}

func (s *AppAiStaff) SetStaffId(v string) *AppAiStaff {
	s.StaffId = &v
	return s
}

func (s *AppAiStaff) SetStaffName(v string) *AppAiStaff {
	s.StaffName = &v
	return s
}

func (s *AppAiStaff) SetStaffType(v string) *AppAiStaff {
	s.StaffType = &v
	return s
}

func (s *AppAiStaff) SetStatus(v string) *AppAiStaff {
	s.Status = &v
	return s
}

func (s *AppAiStaff) Validate() error {
	return dara.Validate(s)
}
