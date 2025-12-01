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
	// The CPU architecture of the node. Valid values: `amd64`, `arm`, and `arm64`.
	//
	// Default value: `amd64`.
	//
	// >  This parameter is required if you want to add a node to an ACK Edge cluster.
	//
	// example:
	//
	// amd64
	Arch *string `json:"arch,omitempty" xml:"arch,omitempty"`
	// The expiration time of the token that is generated. The value is a UNIX timestamp. For example, a value of 1739980800 indicates 00:00:00 (UTC+8) on February 20, 2025.
	//
	// example:
	//
	// 1740037333
	Expired *int64 `json:"expired,omitempty" xml:"expired,omitempty"`
	// Specifies whether to mount data disks to an existing instance when you manually add this instance to the cluster. You can use data disks to store container data and images. Valid values:
	//
	// 	- `true`: mounts data disks to the instance that you want to add. After a data disk is mounted, the original data on the disk is erased. Back up data before you mount a data disk.
	//
	// 	- `false`: does not mount data disks to the instance.
	//
	// Default value: `false`.
	//
	// How a data disk is mounted:
	//
	// 	- If the Elastic Compute Service (ECS) instances are already mounted with data disks and the file system of the last data disk is uninitialized, the system automatically formats this data disk to ext4 and uses the disk to store the data in the /var/lib/docker and /var/lib/kubelet directories.
	//
	// 	- If no data disk is mounted to the ECS instance, the system does not purchase a new data disk.
	//
	// example:
	//
	// false
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// Specifies whether to retain the name of an existing instance when it is added to the cluster. If you do not retain the instance name, the instance is renamed in the `worker-k8s-for-cs-<clusterid>` format. Valid values:
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
	// The ID of the node pool to which you want to add an existing node.
	//
	// >  If you do not specify a node pool ID, the node is added to the default node pool.
	//
	// example:
	//
	// np1c9229d9be2d432c93f77a88fca0****
	NodepoolId   *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	OneTimeToken *bool   `json:"one_time_token,omitempty" xml:"one_time_token,omitempty"`
	// The node configurations for the node that you want to add.
	//
	// >  This parameter is required if you want to add a node to an ACK Edge cluster.
	//
	// example:
	//
	// {\\"enableIptables\\": true,\\"manageRuntime\\": true,\\"quiet\\": true,\\"allowedClusterAddons\\": [\\"kube-proxy\\",\\"flannel\\",\\"coredns\\"]}
	Options *string `json:"options,omitempty" xml:"options,omitempty"`
	// A list of ApsaraDB RDS instances. ECS instances in the cluster are automatically added to the whitelist of the ApsaraDB RDS instances.
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
