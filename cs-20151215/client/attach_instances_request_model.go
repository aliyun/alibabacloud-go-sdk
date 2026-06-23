// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpuPolicy(v string) *AttachInstancesRequest
	GetCpuPolicy() *string
	SetFormatDisk(v bool) *AttachInstancesRequest
	GetFormatDisk() *bool
	SetImageId(v string) *AttachInstancesRequest
	GetImageId() *string
	SetInstances(v []*string) *AttachInstancesRequest
	GetInstances() []*string
	SetIsEdgeWorker(v bool) *AttachInstancesRequest
	GetIsEdgeWorker() *bool
	SetKeepInstanceName(v bool) *AttachInstancesRequest
	GetKeepInstanceName() *bool
	SetKeyPair(v string) *AttachInstancesRequest
	GetKeyPair() *string
	SetNodepoolId(v string) *AttachInstancesRequest
	GetNodepoolId() *string
	SetPassword(v string) *AttachInstancesRequest
	GetPassword() *string
	SetRdsInstances(v []*string) *AttachInstancesRequest
	GetRdsInstances() []*string
	SetRuntime(v *Runtime) *AttachInstancesRequest
	GetRuntime() *Runtime
	SetTags(v []*Tag) *AttachInstancesRequest
	GetTags() []*Tag
	SetUserData(v string) *AttachInstancesRequest
	GetUserData() *string
}

type AttachInstancesRequest struct {
	// The CPU management policy of the node. The following policies are supported for clusters of version 1.12.6 or later:
	//
	// - `static`: Allows pods with certain resource characteristics on the node to be granted enhanced CPU affinity and exclusivity.
	//
	// - `none`: Uses the existing default CPU affinity scheme.
	//
	// Default value: `none`.
	//
	// > After you specify `nodepool_id`, this parameter is not supported.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
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
	// - If the ECS instance has data cloud disks mounted and the file system of the last data cloud disk is not initialized, the system automatically formats the data cloud disk to EXT4 to store the content of /var/lib/docker and /var/lib/kubelet (the default data directories for the Docker container runtime and the kubelet component, respectively).
	//
	// - If the ECS instance has no data cloud disks mounted, no new data cloud disk is mounted.
	//
	// > If you choose to store data on a data cloud disk and the ECS instance already has data cloud disks mounted, existing data on the data cloud disk is lost. Back up your data in advance.
	//
	// example:
	//
	// false
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// The custom image ID.
	//
	// - If you specify a custom image ID, the system cloud disk image of the instance is replaced with the custom image.
	//
	// - If you do not specify this parameter, the default system image is used.
	//
	// > After you specify `nodepool_id`, this parameter is not supported.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200529.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The list of ECS instances to be added.
	//
	// This parameter is required.
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// Specifies whether the node to be added is an edge node, that is, an Edge Node Service (ENS) node. Valid values:
	//
	// - `true`: The node to be added is an edge node.
	//
	// - `false`: The node to be added is not an edge node.
	//
	// Default value: `false`.
	//
	// > If the node is an edge node, set this parameter to `true` to identify the node type as an ENS node.
	//
	// example:
	//
	// false
	IsEdgeWorker *bool `json:"is_edge_worker,omitempty" xml:"is_edge_worker,omitempty"`
	// Specifies whether to retain the original instance name. Valid values:
	//
	// - `true`: Retains the instance name.
	//
	// - `false`: Does not retain the instance name.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	KeepInstanceName *bool `json:"keep_instance_name,omitempty" xml:"keep_instance_name,omitempty"`
	// The name of the key pair for the instances to be added. Specify either key_pair or password. You can also leave both parameters empty.
	//
	// > After you specify `nodepool_id`, this parameter is not supported.
	//
	// example:
	//
	// security-key
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The node pool ID. If you do not specify this parameter, the node is added to the default node pool.
	//
	// example:
	//
	// np615c0e0966124216a0412e10afe0****
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// The SSH logon password for the instances to be added. Specify either key_pair or password. You can also leave both parameters empty.
	//
	// The password must meet the following requirements:
	//
	// - The password must be 8 to 30 characters in length.
	//
	// - The password must contain uppercase letters, lowercase letters, digits, and special characters at the same time.
	//
	// - The password cannot contain backslashes (\\) or double quotation marks (").
	//
	// The password is encrypted during transmission for security purposes.
	//
	// example:
	//
	// Hello1234
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The list of ApsaraDB RDS instances.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The container runtime.
	//
	// > After you specify `nodepool_id`, this parameter is not supported.
	//
	// name: The name of the container runtime. ACK supports the following three container runtimes:
	//
	// - containerd: Recommended. Supported by all cluster versions.
	//
	// - Sandboxed-Container.runv: Sandboxed container that provides higher isolation. Supported by clusters of version 1.24 or earlier.
	//
	// - docker: Supported by clusters of version 1.22 or earlier.
	//
	// Default value: containerd.
	//
	// containerd: The container runtime version. Default value: the latest version.
	//
	// For more information about changes to the sandboxed container runtime, see [Release notes for the sandboxed container runtime](https://help.aliyun.com/document_detail/160312.html).
	Runtime *Runtime `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The node labels. Label definition rules:
	//
	// - Labels are case-sensitive key-value pairs. You can set up to 20 labels.
	//
	// - Label keys cannot be duplicate and can be up to 64 characters in length.
	//
	// - Label values can be empty and can be up to 128 characters in length.
	//
	// - Label keys and values cannot start with `aliyun`, `acs:`, `https://`, or `http://`.
	//
	// For more information, see [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set).
	//
	// > After you specify `nodepool_id`, this parameter is not supported.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The instance user data of the node. For more information, see [Generate instance user data](https://help.aliyun.com/document_detail/49121.html).
	//
	// > After you specify `nodepool_id`, this parameter is not supported.
	//
	// example:
	//
	// IyEvdXNyL2Jpbi9iYXNoCmVjaG8gIkhlbGxvIEFDSyEi
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
}

func (s AttachInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachInstancesRequest) GetCpuPolicy() *string {
	return s.CpuPolicy
}

func (s *AttachInstancesRequest) GetFormatDisk() *bool {
	return s.FormatDisk
}

func (s *AttachInstancesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *AttachInstancesRequest) GetInstances() []*string {
	return s.Instances
}

func (s *AttachInstancesRequest) GetIsEdgeWorker() *bool {
	return s.IsEdgeWorker
}

func (s *AttachInstancesRequest) GetKeepInstanceName() *bool {
	return s.KeepInstanceName
}

func (s *AttachInstancesRequest) GetKeyPair() *string {
	return s.KeyPair
}

func (s *AttachInstancesRequest) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *AttachInstancesRequest) GetPassword() *string {
	return s.Password
}

func (s *AttachInstancesRequest) GetRdsInstances() []*string {
	return s.RdsInstances
}

func (s *AttachInstancesRequest) GetRuntime() *Runtime {
	return s.Runtime
}

func (s *AttachInstancesRequest) GetTags() []*Tag {
	return s.Tags
}

func (s *AttachInstancesRequest) GetUserData() *string {
	return s.UserData
}

func (s *AttachInstancesRequest) SetCpuPolicy(v string) *AttachInstancesRequest {
	s.CpuPolicy = &v
	return s
}

func (s *AttachInstancesRequest) SetFormatDisk(v bool) *AttachInstancesRequest {
	s.FormatDisk = &v
	return s
}

func (s *AttachInstancesRequest) SetImageId(v string) *AttachInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *AttachInstancesRequest) SetInstances(v []*string) *AttachInstancesRequest {
	s.Instances = v
	return s
}

func (s *AttachInstancesRequest) SetIsEdgeWorker(v bool) *AttachInstancesRequest {
	s.IsEdgeWorker = &v
	return s
}

func (s *AttachInstancesRequest) SetKeepInstanceName(v bool) *AttachInstancesRequest {
	s.KeepInstanceName = &v
	return s
}

func (s *AttachInstancesRequest) SetKeyPair(v string) *AttachInstancesRequest {
	s.KeyPair = &v
	return s
}

func (s *AttachInstancesRequest) SetNodepoolId(v string) *AttachInstancesRequest {
	s.NodepoolId = &v
	return s
}

func (s *AttachInstancesRequest) SetPassword(v string) *AttachInstancesRequest {
	s.Password = &v
	return s
}

func (s *AttachInstancesRequest) SetRdsInstances(v []*string) *AttachInstancesRequest {
	s.RdsInstances = v
	return s
}

func (s *AttachInstancesRequest) SetRuntime(v *Runtime) *AttachInstancesRequest {
	s.Runtime = v
	return s
}

func (s *AttachInstancesRequest) SetTags(v []*Tag) *AttachInstancesRequest {
	s.Tags = v
	return s
}

func (s *AttachInstancesRequest) SetUserData(v string) *AttachInstancesRequest {
	s.UserData = &v
	return s
}

func (s *AttachInstancesRequest) Validate() error {
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
