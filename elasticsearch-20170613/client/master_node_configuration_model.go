// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMasterNodeConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *MasterNodeConfiguration
	GetAmount() *int64
	SetDisk(v int64) *MasterNodeConfiguration
	GetDisk() *int64
	SetDiskType(v string) *MasterNodeConfiguration
	GetDiskType() *string
	SetSpec(v string) *MasterNodeConfiguration
	GetSpec() *string
}

type MasterNodeConfiguration struct {
	// This parameter is required.
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	Disk *int64 `json:"disk,omitempty" xml:"disk,omitempty"`
	// This parameter is required.
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// This parameter is required.
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s MasterNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s MasterNodeConfiguration) GoString() string {
	return s.String()
}

func (s *MasterNodeConfiguration) GetAmount() *int64 {
	return s.Amount
}

func (s *MasterNodeConfiguration) GetDisk() *int64 {
	return s.Disk
}

func (s *MasterNodeConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *MasterNodeConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *MasterNodeConfiguration) SetAmount(v int64) *MasterNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *MasterNodeConfiguration) SetDisk(v int64) *MasterNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *MasterNodeConfiguration) SetDiskType(v string) *MasterNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *MasterNodeConfiguration) SetSpec(v string) *MasterNodeConfiguration {
	s.Spec = &v
	return s
}

func (s *MasterNodeConfiguration) Validate() error {
	return dara.Validate(s)
}
