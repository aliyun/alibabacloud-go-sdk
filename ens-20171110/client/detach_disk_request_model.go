// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DetachDiskRequest
	GetDiskId() *string
	SetInstanceId(v string) *DetachDiskRequest
	GetInstanceId() *string
}

type DetachDiskRequest struct {
	// The ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-5r7v69e0bejrnzger09w7****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5inkeimcipxk26yqtzm4q****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DetachDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachDiskRequest) GoString() string {
	return s.String()
}

func (s *DetachDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DetachDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachDiskRequest) SetDiskId(v string) *DetachDiskRequest {
	s.DiskId = &v
	return s
}

func (s *DetachDiskRequest) SetInstanceId(v string) *DetachDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachDiskRequest) Validate() error {
	return dara.Validate(s)
}
