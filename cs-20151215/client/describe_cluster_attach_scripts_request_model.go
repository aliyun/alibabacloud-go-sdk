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
	// > This parameter is required when the cluster type is managed edge cluster.
	//
	// example:
	//
	// amd64
	Arch *string `json:"arch,omitempty" xml:"arch,omitempty"`
	// The expiration time of the generated token. The value is a UNIX timestamp. For example, 1739980800 indicates 2025-02-20 00:00:00.
	//
	// example:
	//
	// 1740037333
	Expired *int64 `json:"expired,omitempty" xml:"expired,omitempty"`
	// Specifies whether to mount data disks to the instance when you manually add the existing instance to the cluster. Container and image data is stored on the data disks. Valid values:
	//
	// - `true`: Mounts data disks to the instance. Existing data on the data disks will be lost. Back up your data before you proceed.
	//
	// - `false`: Does not mount data disks to the instance.
	//
	// Default value: `false`.
	//
	// Data disk mounting rules:
	//
	// - If data disks are already mounted to the ECS instance and the file system of the last data disk is not initialized, the system automatically formats the data disk as ext4 to store /var/lib/docker and /var/lib/kubelet.
	//
	// - If no data disks are mounted to the ECS instance, no new data disks are mounted.
	//
	// example:
	//
	// false
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// Specifies whether to retain the instance name when adding an existing instance to the cluster. If the instance name is not retained, the instance name is in the format of `worker-k8s-for-cs-<clusterid>`. Valid values:
	//
	// - `true`: Retains the instance name.
	//
	// - `false`: Does not retain the instance name. The instance name is replaced based on system rules.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	KeepInstanceName *bool `json:"keep_instance_name,omitempty" xml:"keep_instance_name,omitempty"`
	// The node pool ID. You can add the node to a specified node pool.
	//
	// > If you do not specify a node pool ID, the node is added to the default node pool.
	//
	// example:
	//
	// np1c9229d9be2d432c93f77a88fca0****
	NodepoolId   *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	OneTimeToken *bool   `json:"one_time_token,omitempty" xml:"one_time_token,omitempty"`
	// The configuration parameters for node registration.
	//
	// > This parameter is required when the cluster type is managed edge cluster.
	//
	// example:
	//
	// {"enableIptables": true,"manageRuntime": true,"quiet": true,"allowedClusterAddons": ["kube-proxy","flannel","coredns"]}
	Options *string `json:"options,omitempty" xml:"options,omitempty"`
	// If you specify a list of ApsaraDB RDS instances, the ECS instances in the cluster are automatically added to the whitelists of the specified ApsaraDB RDS instances.
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
