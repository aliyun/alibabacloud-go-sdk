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
	SetFormatDiskOption(v string) *DescribeClusterAttachScriptsRequest
	GetFormatDiskOption() *string
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
	// > Required when the cluster type is ACK Edge.
	//
	// example:
	//
	// amd64
	Arch *string `json:"arch,omitempty" xml:"arch,omitempty"`
	// The expiration time of the generated token. The time is in the UNIX timestamp format. For example, 1739980800 indicates 2025-02-20 00:00:00.
	//
	// example:
	//
	// 1740037333
	Expired *int64 `json:"expired,omitempty" xml:"expired,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use `format_disk_option` instead.
	//
	// - `true`: Equivalent to setting `format_disk_option` to `DataDisk`.
	//
	// - `false`: Equivalent to setting `format_disk_option` to `SystemDisk`.
	//
	// example:
	//
	// false
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// The disk formatting method. This parameter is mutually exclusive with `format_disk` when `format_disk` is set to `true`. Valid values:
	//
	// - `SystemDisk`: No additional initialization is performed. Container runtime data is stored on the system cloud disk or any disk that is already mounted to the relevant directory during import.
	//
	// - `DataDisk`: If a data cloud disk or local disk is mounted, the last disk is formatted and mounted to store container runtime data.
	//
	// - `NodePoolDiskConfig`: Initializes data cloud disks based on the configurations of the imported node pool. Cloud disks and elastic ephemeral disks are matched based on the `disk_name` field of the node pool and the `DiskName` field of the block storage. Local disks are matched sequentially based on the PCI address order. Node pool configurations that do not match any disk are ignored.
	//
	// 	Warning: The `DataDisk` option is retained only for historical compatibility and is strongly discouraged.
	//
	// This option uses the last disk in lexicographical order of device names under `/dev` (NVMe is preferred). Note that nvme9 takes precedence over nvme10. Operating systems with newer kernels such as Alinux 4 cannot guarantee a stable device order.
	//
	//
	// Disks with a recognizable file system during import skip the formatting step and are directly mounted.
	//
	// example:
	//
	// NodePoolDiskConfig
	FormatDiskOption *string `json:"format_disk_option,omitempty" xml:"format_disk_option,omitempty"`
	// Specifies whether to retain the instance name when adding an existing instance to the cluster. If you do not retain the instance name, the instance name format is `worker-k8s-for-cs-<clusterid>`. Valid values:
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
	// The ID of the node pool. Specifies the node pool to which the node is added.
	//
	// > If you do not specify a node pool ID, the node is added to the default node pool by default.
	//
	// example:
	//
	// np1c9229d9be2d432c93f77a88fca0****
	NodepoolId   *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	OneTimeToken *bool   `json:"one_time_token,omitempty" xml:"one_time_token,omitempty"`
	// The access configuration parameters for the node.
	//
	// > Required when the cluster type is ACK Edge.
	//
	// example:
	//
	// {"enableIptables": true,"manageRuntime": true,"quiet": true,"allowedClusterAddons": ["kube-proxy","flannel","coredns"]}
	Options *string `json:"options,omitempty" xml:"options,omitempty"`
	// If an ApsaraDB RDS instance list is specified, the ECS instances of the cluster nodes are automatically added to the RDS access whitelist.
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

func (s *DescribeClusterAttachScriptsRequest) GetFormatDiskOption() *string {
	return s.FormatDiskOption
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

func (s *DescribeClusterAttachScriptsRequest) SetFormatDiskOption(v string) *DescribeClusterAttachScriptsRequest {
	s.FormatDiskOption = &v
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
