// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClientNodeConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *ClientNodeConfiguration
	GetAmount() *int64
	SetDisk(v int64) *ClientNodeConfiguration
	GetDisk() *int64
	SetDiskType(v string) *ClientNodeConfiguration
	GetDiskType() *string
	SetSpec(v string) *ClientNodeConfiguration
	GetSpec() *string
}

type ClientNodeConfiguration struct {
	// This parameter is required.
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	Disk *int64 `json:"disk,omitempty" xml:"disk,omitempty"`
	// This parameter is required.
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// This parameter is required.
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ClientNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ClientNodeConfiguration) GoString() string {
	return s.String()
}

func (s *ClientNodeConfiguration) GetAmount() *int64 {
	return s.Amount
}

func (s *ClientNodeConfiguration) GetDisk() *int64 {
	return s.Disk
}

func (s *ClientNodeConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *ClientNodeConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *ClientNodeConfiguration) SetAmount(v int64) *ClientNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *ClientNodeConfiguration) SetDisk(v int64) *ClientNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *ClientNodeConfiguration) SetDiskType(v string) *ClientNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *ClientNodeConfiguration) SetSpec(v string) *ClientNodeConfiguration {
	s.Spec = &v
	return s
}

func (s *ClientNodeConfiguration) Validate() error {
	return dara.Validate(s)
}
