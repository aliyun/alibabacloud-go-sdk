// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstancesToNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFormatDisk(v bool) *AttachInstancesToNodePoolRequest
	GetFormatDisk() *bool
	SetInstances(v []*string) *AttachInstancesToNodePoolRequest
	GetInstances() []*string
	SetKeepInstanceName(v bool) *AttachInstancesToNodePoolRequest
	GetKeepInstanceName() *bool
	SetPassword(v string) *AttachInstancesToNodePoolRequest
	GetPassword() *string
}

type AttachInstancesToNodePoolRequest struct {
	// Specifies whether to store container data and images on data disks. Valid values:
	//
	// 	- `true`: stores container data and images on data disks.
	//
	// 	- `false`: does not store container data or images on data disks.
	//
	// Default value: `false`.
	//
	// How to mount a data disk:
	//
	// 	- If the ECS instances are already mounted with data disks and the file system of the last data disk is not initialized, the system automatically formats this data disk to ext4 and mounts it to /var/lib/docker and /var/lib/kubelet.
	//
	// 	- If no data disk is attached to the ECS instances, the system does not purchase a new data disk.
	//
	// > If you choose to store container data and images on a data disk and the data disk is already mounted to the ECS instance, the existing data on the data disk will be cleared. You can back up the disk to avoid data loss.
	//
	// example:
	//
	// false
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// The IDs of the instances to be added.
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// Specifies whether to retain the instance name. Valid values:
	//
	// 	- `true`: retains the instance name.
	//
	// 	- `false`: does not retain the instance name.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	KeepInstanceName *bool `json:"keep_instance_name,omitempty" xml:"keep_instance_name,omitempty"`
	// The SSH password that is used to log on to the instance.
	//
	// example:
	//
	// Hello1234
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s AttachInstancesToNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachInstancesToNodePoolRequest) GoString() string {
	return s.String()
}

func (s *AttachInstancesToNodePoolRequest) GetFormatDisk() *bool {
	return s.FormatDisk
}

func (s *AttachInstancesToNodePoolRequest) GetInstances() []*string {
	return s.Instances
}

func (s *AttachInstancesToNodePoolRequest) GetKeepInstanceName() *bool {
	return s.KeepInstanceName
}

func (s *AttachInstancesToNodePoolRequest) GetPassword() *string {
	return s.Password
}

func (s *AttachInstancesToNodePoolRequest) SetFormatDisk(v bool) *AttachInstancesToNodePoolRequest {
	s.FormatDisk = &v
	return s
}

func (s *AttachInstancesToNodePoolRequest) SetInstances(v []*string) *AttachInstancesToNodePoolRequest {
	s.Instances = v
	return s
}

func (s *AttachInstancesToNodePoolRequest) SetKeepInstanceName(v bool) *AttachInstancesToNodePoolRequest {
	s.KeepInstanceName = &v
	return s
}

func (s *AttachInstancesToNodePoolRequest) SetPassword(v string) *AttachInstancesToNodePoolRequest {
	s.Password = &v
	return s
}

func (s *AttachInstancesToNodePoolRequest) Validate() error {
	return dara.Validate(s)
}
