// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachRCDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteWithInstance(v bool) *DetachRCDiskRequest
	GetDeleteWithInstance() *bool
	SetDiskId(v string) *DetachRCDiskRequest
	GetDiskId() *string
	SetInstanceId(v string) *DetachRCDiskRequest
	GetInstanceId() *string
	SetRegionId(v string) *DetachRCDiskRequest
	GetRegionId() *string
}

type DetachRCDiskRequest struct {
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The ID of the disk that you want to detach.
	//
	// This parameter is required.
	//
	// example:
	//
	// rcd-f8zh55g5gbk1byjr****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-dh2jf9n6j4s14926****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachRCDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachRCDiskRequest) GoString() string {
	return s.String()
}

func (s *DetachRCDiskRequest) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DetachRCDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DetachRCDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachRCDiskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachRCDiskRequest) SetDeleteWithInstance(v bool) *DetachRCDiskRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *DetachRCDiskRequest) SetDiskId(v string) *DetachRCDiskRequest {
	s.DiskId = &v
	return s
}

func (s *DetachRCDiskRequest) SetInstanceId(v string) *DetachRCDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachRCDiskRequest) SetRegionId(v string) *DetachRCDiskRequest {
	s.RegionId = &v
	return s
}

func (s *DetachRCDiskRequest) Validate() error {
	return dara.Validate(s)
}
