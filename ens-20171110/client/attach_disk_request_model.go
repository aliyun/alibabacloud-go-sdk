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
	// Specifies whether the cloud disk to be mounted is released when the instance is released. Valid values:
	//
	// - true: The cloud disk is released when the instance is released.
	//
	// - false: The cloud disk is retained when the instance is released.
	//
	// - Empty: The default value is false.
	//
	// example:
	//
	// false
	DeleteWithInstance *string `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The ID of the cloud disk to mount. The cloud disk and the instance must be in the same edge zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-5saf13yy6sopmmg88mzsg****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The instance ID.
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
