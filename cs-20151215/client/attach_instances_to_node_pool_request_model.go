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
	// Specifies whether to store container data and images on a data cloud disk. Valid values:
	//
	// - `true`: Stores container data and images on a data cloud disk.
	//
	// - `false`: Does not store container data and images on a data cloud disk.
	//
	// Default value: `false`.
	//
	//
	// Data cloud disk mounting rules:
	//
	// - If the ECS instance has data cloud disks attached and the file system of the last data cloud disk is not initialized, the system automatically formats the data cloud disk as EXT4 to store /var/lib/docker and /var/lib/kubelet.
	//
	// - If the ECS instance has no data cloud disks attached, no new data cloud disk is mounted.
	//
	// > If you choose to store data on a data cloud disk and the ECS instance already has data cloud disks attached, existing data on the data cloud disk is lost. Back up your data in advance.
	//
	// example:
	//
	// false
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// The list of ECS instances to be added.
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// Specifies whether to retain the original instance name. Valid values:
	//
	// - `true`: Retains the instance name.
	//
	// - `false`: Does not retain the instance name.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	KeepInstanceName *bool `json:"keep_instance_name,omitempty" xml:"keep_instance_name,omitempty"`
	// The SSH logon password of the instances to be added.
	//
	// example:
	//
	// ******
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
