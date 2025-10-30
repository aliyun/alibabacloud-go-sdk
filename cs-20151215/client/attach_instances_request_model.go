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
	// The CPU management policy of the node. The following policies are supported if the Kubernetes version of the cluster is 1.12.6 or later:
	//
	// 	- `static`: allows pods with specific resource characteristics on the node to be granted enhanced CPU affinity and exclusivity.
	//
	// 	- `none`: uses default CPU affinity.
	//
	// Default value: `none`
	//
	// >  This parameter is not supported if you specify `nodepool_id`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// Specifies whether to store container data and images on data disks. Valid value:
	//
	// 	- `true`: stores container data and images on data disks.
	//
	// 	- `false`: does not store container data or images on data disks.
	//
	// Default value: `false`.
	//
	// How data disks are attached:
	//
	// 	- If the ECS instance is already attached with data disks and the file system of the last data disk is not initialized, the system automatically formats this data disk to ext4. Then, the system uses the disk to store the data in the /var/lib/docker and /var/lib/kubelet directories.
	//
	// 	- If no data disk is attached to the ECS instance, the system does not purchase a new data disk.
	//
	// >  If you choose to store container data and images on data disks and a data disk is already attached to the ECS instance, the original data on this data disk is cleared. You can back up the disk to prevent data loss.
	//
	// example:
	//
	// false
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// The custom image ID. If you do not specify this parameter, the default system image is used.
	//
	// >
	//
	// 	- If you specify a custom image, the custom image is used to deploy the operating system on the system disk of the node.
	//
	// 	- This parameter is not supported if you specify `nodepool_id`.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200529.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The ECS instances that you want to add.
	//
	// This parameter is required.
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// Specifies whether the node that you want to add is an Edge Node Service (ENS) node. Valid value:
	//
	// 	- `true`: the node that you want to add is an ENS node.
	//
	// 	- `false`: the node that you want to add is not an ENS node.
	//
	// Default value: `false`.
	//
	// >  If the node that you want to add is an ENS node, you must set the value to `true`. This allows you to identify the node.
	//
	// example:
	//
	// false
	IsEdgeWorker *bool `json:"is_edge_worker,omitempty" xml:"is_edge_worker,omitempty"`
	// Specifies whether to retain the instance name. Valid value:
	//
	// 	- `true`: retains the instance name.
	//
	// 	- `false`: does not retain the instance name.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	KeepInstanceName *bool `json:"keep_instance_name,omitempty" xml:"keep_instance_name,omitempty"`
	// The name of the key pair used to log on to the ECS instances. You must specify this parameter or `login_password`.
	//
	// >  This parameter is not supported if you specify `nodepool_id`.
	//
	// example:
	//
	// secrity-key
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The ID of the node pool to which the node is added. If you do not specify this parameter, the node is added to the default node pool.
	//
	// example:
	//
	// np615c0e0966124216a0412e10afe0****
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// The SSH logon password used to log on to the ECS instances. You must specify this parameter or `key_pair`. The password must be 8 to 30 characters in length, and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The password cannot contain backslashes (\\\\) or double quotation marks (").
	//
	// The password is encrypted during data transfer to ensure security.
	//
	// example:
	//
	// Hello1234
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// A list of ApsaraDB RDS instances.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The container runtime.
	//
	// >  This parameter is not supported if you specify `nodepool_id`.
	Runtime *Runtime `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The labels that you want to add to the node. When you add labels to a node, the following rules apply:
	//
	// 	- A label is a case-sensitive key-value pair. You can add up to 20 labels.
	//
	// 	- The key must be unique and cannot exceed 64 characters in length. The value can be empty and cannot exceed 128 characters in length. Keys and values cannot start with `aliyun`, `acs:`, `https://`, or `http://`. For more information, see [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set).
	//
	// >  This parameter is not supported if you specify `nodepool_id`.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The user-defined data on the node. For more information, see [Use instance user data to automatically run commands or scripts on instance startup](https://help.aliyun.com/document_detail/49121.html).
	//
	// >  This parameter is not supported if you specify `nodepool_id`.
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
