// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachRCDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteWithInstance(v bool) *AttachRCDiskRequest
	GetDeleteWithInstance() *bool
	SetDiskId(v string) *AttachRCDiskRequest
	GetDiskId() *string
	SetInstanceId(v string) *AttachRCDiskRequest
	GetInstanceId() *string
	SetRegionId(v string) *AttachRCDiskRequest
	GetRegionId() *string
}

type AttachRCDiskRequest struct {
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The disk ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rcd-wz98hnpj2sjo85zc7t2w
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-dh2jf9n6j4s14926****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachRCDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachRCDiskRequest) GoString() string {
	return s.String()
}

func (s *AttachRCDiskRequest) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *AttachRCDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *AttachRCDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachRCDiskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachRCDiskRequest) SetDeleteWithInstance(v bool) *AttachRCDiskRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *AttachRCDiskRequest) SetDiskId(v string) *AttachRCDiskRequest {
	s.DiskId = &v
	return s
}

func (s *AttachRCDiskRequest) SetInstanceId(v string) *AttachRCDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachRCDiskRequest) SetRegionId(v string) *AttachRCDiskRequest {
	s.RegionId = &v
	return s
}

func (s *AttachRCDiskRequest) Validate() error {
	return dara.Validate(s)
}
