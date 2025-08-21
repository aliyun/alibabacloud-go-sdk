// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteWithInstance(v string) *AttachDiskRequest
	GetDeleteWithInstance() *string
	SetDiskId(v string) *AttachDiskRequest
	GetDiskId() *string
	SetInstanceId(v string) *AttachDiskRequest
	GetInstanceId() *string
}

type AttachDiskRequest struct {
	// Specifies whether the disk to be attached is released with the instance. Valid values:
	//
	// 	- true: The disk will be released when the ECS instance is released.
	//
	// 	- false: The disk will be retained when the ECS instance is released.
	//
	// 	- If you leave this parameter empty, the default value is used.
	//
	// example:
	//
	// false
	DeleteWithInstance *string `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The ID of the disk to be attached. The cloud disk and the instance must belong to the same node.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-5saf13yy6sopmmg88mzsg****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5rr19av7tkpgi9os52ag1****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AttachDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachDiskRequest) GoString() string {
	return s.String()
}

func (s *AttachDiskRequest) GetDeleteWithInstance() *string {
	return s.DeleteWithInstance
}

func (s *AttachDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *AttachDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachDiskRequest) SetDeleteWithInstance(v string) *AttachDiskRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *AttachDiskRequest) SetDiskId(v string) *AttachDiskRequest {
	s.DiskId = &v
	return s
}

func (s *AttachDiskRequest) SetInstanceId(v string) *AttachDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachDiskRequest) Validate() error {
	return dara.Validate(s)
}
