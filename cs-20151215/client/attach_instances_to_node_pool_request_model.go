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
	SetFormatDiskOption(v string) *AttachInstancesToNodePoolRequest
	GetFormatDiskOption() *string
	SetInstances(v []*string) *AttachInstancesToNodePoolRequest
	GetInstances() []*string
	SetKeepInstanceName(v bool) *AttachInstancesToNodePoolRequest
	GetKeepInstanceName() *bool
	SetPassword(v string) *AttachInstancesToNodePoolRequest
	GetPassword() *string
}

type AttachInstancesToNodePoolRequest struct {
	// Deprecated
	//
	// **[Deprecated]*	- Use `format_disk_option` instead. Valid values:
	//
	// - `true`: equivalent to setting `format_disk_option` to `DataDisk`.
	//
	// - `false`: equivalent to setting `format_disk_option` to `SystemDisk`.
	//
	// example:
	//
	// false
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// The disk formatting method. This parameter is mutually exclusive with `format_disk` set to `true`. Valid values:
	//
	// - `SystemDisk`: No additional initialization is performed. Container runtime data is stored on the system cloud disk or any disk that is already mounted to the relevant directory during import.
	//
	// - `DataDisk`: If a data cloud disk or local disk is mounted, the last disk is formatted and mounted to store container runtime data.
	//
	// - `NodePoolDiskConfig`: Initializes data cloud disks based on the configuration of the imported node pool. Cloud disks and elastic ephemeral disks are matched based on the `disk_name` field of the node pool and the `DiskName` field of the block storage. Local disks are matched sequentially based on the PCI address order. Node pool configurations that do not match any disk are ignored.
	//
	// 	Warning: The `DataDisk` option is retained only for historical compatibility and is strongly discouraged.
	//
	// This option uses the last disk in lexicographical order of device names under `/dev` (NVMe is preferred). Note that nvme9 takes precedence over nvme10. Operating systems with newer kernels such as Alinux 4 cannot guarantee a stable device order.
	//
	//
	// Disks with recognizable file systems during import skip the formatting step and are mounted directly.
	//
	// example:
	//
	// NodePoolDiskConfig
	FormatDiskOption *string `json:"format_disk_option,omitempty" xml:"format_disk_option,omitempty"`
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
	// The SSH logon password for the instances to be added.
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

func (s *AttachInstancesToNodePoolRequest) GetFormatDiskOption() *string {
	return s.FormatDiskOption
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

func (s *AttachInstancesToNodePoolRequest) SetFormatDiskOption(v string) *AttachInstancesToNodePoolRequest {
	s.FormatDiskOption = &v
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
