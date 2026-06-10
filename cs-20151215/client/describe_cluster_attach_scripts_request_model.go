// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAttachScriptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArch(v string) *DescribeClusterAttachScriptsRequest
	GetArch() *string
	SetExpired(v int64) *DescribeClusterAttachScriptsRequest
	GetExpired() *int64
	SetFormatDisk(v bool) *DescribeClusterAttachScriptsRequest
	GetFormatDisk() *bool
	SetKeepInstanceName(v bool) *DescribeClusterAttachScriptsRequest
	GetKeepInstanceName() *bool
	SetNodepoolId(v string) *DescribeClusterAttachScriptsRequest
	GetNodepoolId() *string
	SetOneTimeToken(v bool) *DescribeClusterAttachScriptsRequest
	GetOneTimeToken() *bool
	SetOptions(v string) *DescribeClusterAttachScriptsRequest
	GetOptions() *string
	SetRdsInstances(v []*string) *DescribeClusterAttachScriptsRequest
	GetRdsInstances() []*string
}

type DescribeClusterAttachScriptsRequest struct {
	// The CPU architecture of the node. Supported CPU architectures include `amd64`, `arm`, and `arm64`.
	//
	// Default value: `amd64`.
	//
	// > This parameter is required if the cluster is a managed edge cluster.
	//
	// example:
	//
	// amd64
	Arch *string `json:"arch,omitempty" xml:"arch,omitempty"`
	// The Unix timestamp that indicates when the generated token expires. For example, the timestamp 1739980800 corresponds to 00:00:00 on February 20, 2025 (UTC).
	//
	// example:
	//
	// 1740037333
	Expired *int64 `json:"expired,omitempty" xml:"expired,omitempty"`
	// Specifies whether to mount a data disk to the instance and store containers and images on the data disk when you manually add an existing instance to the cluster. Valid values:
	//
	// - `true`: Mounts the data disk to the instance. The original data on the data disk will be erased. Back up your data before you proceed.
	//
	// - `false`: Does not mount the data disk to the instance.
	//
	// Default value: `false`.
	//
	// Data disk mounting rules:
	//
	// - If an ECS instance has data disks attached and the last data disk is uninitialized, the system automatically formats that disk to ext4 and uses it to store content for `/var/lib/docker` and `/var/lib/kubelet`.
	//
	// - If no data disk is attached to the ECS instance, the system does not mount a new data disk.
	//
	// example:
	//
	// false
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// Specifies whether to retain the instance name when the instance is added to the cluster. If you do not retain the instance name, the system renames the instance to use the `worker-k8s-for-cs-<clusterid>` format. Valid values:
	//
	// - `true`: Retains the instance name.
	//
	// - `false`: Does not retain the instance name. The system renames the instance based on a system rule.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	KeepInstanceName *bool `json:"keep_instance_name,omitempty" xml:"keep_instance_name,omitempty"`
	// The node pool ID. You can add the node to a specific node pool.
	//
	// > If you do not specify a node pool ID, the node is added to the default node pool.
	//
	// example:
	//
	// np1c9229d9be2d432c93f77a88fca0****
	NodepoolId   *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	OneTimeToken *bool   `json:"one_time_token,omitempty" xml:"one_time_token,omitempty"`
	// The configuration parameters for node attachment.
	//
	// > This parameter is required if the cluster is a managed edge cluster.
	//
	// example:
	//
	// {"enableIptables": true,"manageRuntime": true,"quiet": true,"allowedClusterAddons": ["kube-proxy","flannel","coredns"]}
	Options *string `json:"options,omitempty" xml:"options,omitempty"`
	// If you specify a list of RDS instances, the system automatically adds the ECS instances of the cluster nodes to the access whitelists of the specified RDS instances.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttachScriptsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttachScriptsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttachScriptsRequest) GetArch() *string {
	return s.Arch
}

func (s *DescribeClusterAttachScriptsRequest) GetExpired() *int64 {
	return s.Expired
}

func (s *DescribeClusterAttachScriptsRequest) GetFormatDisk() *bool {
	return s.FormatDisk
}

func (s *DescribeClusterAttachScriptsRequest) GetKeepInstanceName() *bool {
	return s.KeepInstanceName
}

func (s *DescribeClusterAttachScriptsRequest) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *DescribeClusterAttachScriptsRequest) GetOneTimeToken() *bool {
	return s.OneTimeToken
}

func (s *DescribeClusterAttachScriptsRequest) GetOptions() *string {
	return s.Options
}

func (s *DescribeClusterAttachScriptsRequest) GetRdsInstances() []*string {
	return s.RdsInstances
}

func (s *DescribeClusterAttachScriptsRequest) SetArch(v string) *DescribeClusterAttachScriptsRequest {
	s.Arch = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) SetExpired(v int64) *DescribeClusterAttachScriptsRequest {
	s.Expired = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) SetFormatDisk(v bool) *DescribeClusterAttachScriptsRequest {
	s.FormatDisk = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) SetKeepInstanceName(v bool) *DescribeClusterAttachScriptsRequest {
	s.KeepInstanceName = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) SetNodepoolId(v string) *DescribeClusterAttachScriptsRequest {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) SetOneTimeToken(v bool) *DescribeClusterAttachScriptsRequest {
	s.OneTimeToken = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) SetOptions(v string) *DescribeClusterAttachScriptsRequest {
	s.Options = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) SetRdsInstances(v []*string) *DescribeClusterAttachScriptsRequest {
	s.RdsInstances = v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) Validate() error {
	return dara.Validate(s)
}
