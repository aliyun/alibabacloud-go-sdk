// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleOutClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudMonitorFlags(v bool) *ScaleOutClusterRequest
	GetCloudMonitorFlags() *bool
	SetCount(v int64) *ScaleOutClusterRequest
	GetCount() *int64
	SetCpuPolicy(v string) *ScaleOutClusterRequest
	GetCpuPolicy() *string
	SetImageId(v string) *ScaleOutClusterRequest
	GetImageId() *string
	SetKeyPair(v string) *ScaleOutClusterRequest
	GetKeyPair() *string
	SetLoginPassword(v string) *ScaleOutClusterRequest
	GetLoginPassword() *string
	SetRdsInstances(v []*string) *ScaleOutClusterRequest
	GetRdsInstances() []*string
	SetRuntime(v *Runtime) *ScaleOutClusterRequest
	GetRuntime() *Runtime
	SetTags(v []*Tag) *ScaleOutClusterRequest
	GetTags() []*Tag
	SetTaints(v []*Taint) *ScaleOutClusterRequest
	GetTaints() []*Taint
	SetUserData(v string) *ScaleOutClusterRequest
	GetUserData() *string
	SetVswitchIds(v []*string) *ScaleOutClusterRequest
	GetVswitchIds() []*string
	SetWorkerAutoRenew(v bool) *ScaleOutClusterRequest
	GetWorkerAutoRenew() *bool
	SetWorkerAutoRenewPeriod(v int64) *ScaleOutClusterRequest
	GetWorkerAutoRenewPeriod() *int64
	SetWorkerDataDisks(v []*ScaleOutClusterRequestWorkerDataDisks) *ScaleOutClusterRequest
	GetWorkerDataDisks() []*ScaleOutClusterRequestWorkerDataDisks
	SetWorkerInstanceChargeType(v string) *ScaleOutClusterRequest
	GetWorkerInstanceChargeType() *string
	SetWorkerInstanceTypes(v []*string) *ScaleOutClusterRequest
	GetWorkerInstanceTypes() []*string
	SetWorkerPeriod(v int64) *ScaleOutClusterRequest
	GetWorkerPeriod() *int64
	SetWorkerPeriodUnit(v string) *ScaleOutClusterRequest
	GetWorkerPeriodUnit() *string
	SetWorkerSystemDiskCategory(v string) *ScaleOutClusterRequest
	GetWorkerSystemDiskCategory() *string
	SetWorkerSystemDiskSize(v int64) *ScaleOutClusterRequest
	GetWorkerSystemDiskSize() *int64
}

type ScaleOutClusterRequest struct {
	// Specifies whether to install the CloudMonitor agent. Valid values:
	//
	// 	- `true`: installs the CloudMonitor agent.
	//
	// 	- `false`: does not install the CloudMonitor agent.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// The number of worker nodes that you want to add.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The CPU management policy of nodes in the node pool. The following policies are supported if the Kubernetes version of the cluster is 1.12.6 or later:
	//
	// 	- `static`: allows pods with specific resource characteristics on the node to be granted with enhanced CPU affinity and exclusivity.
	//
	// 	- `none`: specifies that the default CPU affinity is used.
	//
	// Default value: `none`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// Specifies a custom image for nodes. By default, the image provided by ACK is used. You can select a custom image to replace the default image. For more information, see [Custom images](https://help.aliyun.com/document_detail/146647.html).
	//
	// example:
	//
	// m-bp16z7xko3vvv8gt****
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The name of the key pair. You must configure this parameter or the `login_password` parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// secrity-key
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The password for SSH logon. You must configure this parameter or the `key_pair` parameter. The password must be 8 to 30 characters in length, and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello@1234
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The ApsaraDB RDS instances. If you specify a list of ApsaraDB RDS instances, ECS instances in the cluster are automatically added to the whitelist of the ApsaraDB RDS instances.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The container runtime.
	Runtime *Runtime `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The labels that you want to add to the node. When you add labels to a node, the following rules apply:
	//
	// 	- A label is a case-sensitive key-value pair. You can add up to 20 labels.
	//
	// 	- The key must be unique and cannot exceed 64 characters in length. The value can be empty and cannot exceed 128 characters in length. Keys and values cannot start with aliyun, acs:, https://, or http://. For more information, see [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set).
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The taints that you want to add to nodes. Taints can be used together with tolerations to avoid scheduling pods to specified nodes. For more information, see [taint-and-toleration](https://kubernetes.io/zh/docs/concepts/scheduling-eviction/taint-and-toleration/).
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// The user-defined data of the node pool. For more information, see [Generate user-defined data](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// IyEvdXNyL2Jpbi9iYXNoCmVjaG8gIkhlbGxvIEFD****
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// The vSwitch IDs. You can select one to three vSwitches when you create a cluster. To ensure the high availability of the cluster, we recommend that you select vSwitches in different zones.
	//
	// This parameter is required.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// Specifies whether to enable auto-renewal for worker nodes. This parameter takes effect and is required only if `worker_instance_charge_type` is set to `PrePaid`. Valid values:
	//
	// 	- `true`: enables auto-renewal.
	//
	// 	- `false`: does not enable auto-renewal.
	//
	// Default value: `true`
	//
	// example:
	//
	// true
	WorkerAutoRenew *bool `json:"worker_auto_renew,omitempty" xml:"worker_auto_renew,omitempty"`
	// The auto-renewal duration of worker nodes. This parameter takes effect and is required only if the subscription billing method is selected for worker nodes.
	//
	// Valid values: 1, 2, 3, 6, and 12.
	//
	// Default value: `1`.
	//
	// example:
	//
	// 6
	WorkerAutoRenewPeriod *int64 `json:"worker_auto_renew_period,omitempty" xml:"worker_auto_renew_period,omitempty"`
	// The configurations of the data disks that you want to mount to worker nodes. The configurations include the disk type and disk size.
	WorkerDataDisks []*ScaleOutClusterRequestWorkerDataDisks `json:"worker_data_disks,omitempty" xml:"worker_data_disks,omitempty" type:"Repeated"`
	// The billing method of worker nodes. Valid values:
	//
	// 	- `PrePaid`: subscription.
	//
	// 	- `PostPaid`: pay-as-you-go.
	//
	// Default value: `PostPaid`.
	//
	// example:
	//
	// PrePaid
	WorkerInstanceChargeType *string `json:"worker_instance_charge_type,omitempty" xml:"worker_instance_charge_type,omitempty"`
	// The instance configurations of worker nodes.
	//
	// This parameter is required.
	WorkerInstanceTypes []*string `json:"worker_instance_types,omitempty" xml:"worker_instance_types,omitempty" type:"Repeated"`
	// The subscription duration of worker nodes. This parameter takes effect and is required only if `worker_instance_charge_type` is set to `PrePaid`.
	//
	// Valid values: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	WorkerPeriod *int64 `json:"worker_period,omitempty" xml:"worker_period,omitempty"`
	// The billing cycle of worker nodes. This parameter is required only if worker_instance_charge_type is set to `PrePaid`.
	//
	// Set the value to `Month`.
	//
	// example:
	//
	// Month
	WorkerPeriodUnit *string `json:"worker_period_unit,omitempty" xml:"worker_period_unit,omitempty"`
	// The system disk category of worker nodes. Valid values:
	//
	// 	- `cloud_efficiency`: ultra disk.
	//
	// 	- `cloud_ssd`: standard SSD.
	//
	// 	- `cloud_essd`: Enterprise SSD (ESSD).
	//
	// Default value: `cloud_ssd`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_efficiency
	WorkerSystemDiskCategory *string `json:"worker_system_disk_category,omitempty" xml:"worker_system_disk_category,omitempty"`
	// The system disk size of worker nodes. Unit: GiB.
	//
	// Valid values: 40 to 500.
	//
	// Default value: `120`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	WorkerSystemDiskSize *int64 `json:"worker_system_disk_size,omitempty" xml:"worker_system_disk_size,omitempty"`
}

func (s ScaleOutClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleOutClusterRequest) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterRequest) GetCloudMonitorFlags() *bool {
	return s.CloudMonitorFlags
}

func (s *ScaleOutClusterRequest) GetCount() *int64 {
	return s.Count
}

func (s *ScaleOutClusterRequest) GetCpuPolicy() *string {
	return s.CpuPolicy
}

func (s *ScaleOutClusterRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ScaleOutClusterRequest) GetKeyPair() *string {
	return s.KeyPair
}

func (s *ScaleOutClusterRequest) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *ScaleOutClusterRequest) GetRdsInstances() []*string {
	return s.RdsInstances
}

func (s *ScaleOutClusterRequest) GetRuntime() *Runtime {
	return s.Runtime
}

func (s *ScaleOutClusterRequest) GetTags() []*Tag {
	return s.Tags
}

func (s *ScaleOutClusterRequest) GetTaints() []*Taint {
	return s.Taints
}

func (s *ScaleOutClusterRequest) GetUserData() *string {
	return s.UserData
}

func (s *ScaleOutClusterRequest) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *ScaleOutClusterRequest) GetWorkerAutoRenew() *bool {
	return s.WorkerAutoRenew
}

func (s *ScaleOutClusterRequest) GetWorkerAutoRenewPeriod() *int64 {
	return s.WorkerAutoRenewPeriod
}

func (s *ScaleOutClusterRequest) GetWorkerDataDisks() []*ScaleOutClusterRequestWorkerDataDisks {
	return s.WorkerDataDisks
}

func (s *ScaleOutClusterRequest) GetWorkerInstanceChargeType() *string {
	return s.WorkerInstanceChargeType
}

func (s *ScaleOutClusterRequest) GetWorkerInstanceTypes() []*string {
	return s.WorkerInstanceTypes
}

func (s *ScaleOutClusterRequest) GetWorkerPeriod() *int64 {
	return s.WorkerPeriod
}

func (s *ScaleOutClusterRequest) GetWorkerPeriodUnit() *string {
	return s.WorkerPeriodUnit
}

func (s *ScaleOutClusterRequest) GetWorkerSystemDiskCategory() *string {
	return s.WorkerSystemDiskCategory
}

func (s *ScaleOutClusterRequest) GetWorkerSystemDiskSize() *int64 {
	return s.WorkerSystemDiskSize
}

func (s *ScaleOutClusterRequest) SetCloudMonitorFlags(v bool) *ScaleOutClusterRequest {
	s.CloudMonitorFlags = &v
	return s
}

func (s *ScaleOutClusterRequest) SetCount(v int64) *ScaleOutClusterRequest {
	s.Count = &v
	return s
}

func (s *ScaleOutClusterRequest) SetCpuPolicy(v string) *ScaleOutClusterRequest {
	s.CpuPolicy = &v
	return s
}

func (s *ScaleOutClusterRequest) SetImageId(v string) *ScaleOutClusterRequest {
	s.ImageId = &v
	return s
}

func (s *ScaleOutClusterRequest) SetKeyPair(v string) *ScaleOutClusterRequest {
	s.KeyPair = &v
	return s
}

func (s *ScaleOutClusterRequest) SetLoginPassword(v string) *ScaleOutClusterRequest {
	s.LoginPassword = &v
	return s
}

func (s *ScaleOutClusterRequest) SetRdsInstances(v []*string) *ScaleOutClusterRequest {
	s.RdsInstances = v
	return s
}

func (s *ScaleOutClusterRequest) SetRuntime(v *Runtime) *ScaleOutClusterRequest {
	s.Runtime = v
	return s
}

func (s *ScaleOutClusterRequest) SetTags(v []*Tag) *ScaleOutClusterRequest {
	s.Tags = v
	return s
}

func (s *ScaleOutClusterRequest) SetTaints(v []*Taint) *ScaleOutClusterRequest {
	s.Taints = v
	return s
}

func (s *ScaleOutClusterRequest) SetUserData(v string) *ScaleOutClusterRequest {
	s.UserData = &v
	return s
}

func (s *ScaleOutClusterRequest) SetVswitchIds(v []*string) *ScaleOutClusterRequest {
	s.VswitchIds = v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerAutoRenew(v bool) *ScaleOutClusterRequest {
	s.WorkerAutoRenew = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerAutoRenewPeriod(v int64) *ScaleOutClusterRequest {
	s.WorkerAutoRenewPeriod = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerDataDisks(v []*ScaleOutClusterRequestWorkerDataDisks) *ScaleOutClusterRequest {
	s.WorkerDataDisks = v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerInstanceChargeType(v string) *ScaleOutClusterRequest {
	s.WorkerInstanceChargeType = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerInstanceTypes(v []*string) *ScaleOutClusterRequest {
	s.WorkerInstanceTypes = v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerPeriod(v int64) *ScaleOutClusterRequest {
	s.WorkerPeriod = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerPeriodUnit(v string) *ScaleOutClusterRequest {
	s.WorkerPeriodUnit = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerSystemDiskCategory(v string) *ScaleOutClusterRequest {
	s.WorkerSystemDiskCategory = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerSystemDiskSize(v int64) *ScaleOutClusterRequest {
	s.WorkerSystemDiskSize = &v
	return s
}

func (s *ScaleOutClusterRequest) Validate() error {
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
	if s.Taints != nil {
		for _, item := range s.Taints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WorkerDataDisks != nil {
		for _, item := range s.WorkerDataDisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ScaleOutClusterRequestWorkerDataDisks struct {
	// The ID of the automatic snapshot policy. The system performs automatic backup for a cloud disk based on the specified automatic snapshot policy.
	//
	// By default, this parameter is left empty, which indicates that automatic backup is disabled.
	//
	// example:
	//
	// sp-bp14yziiuvu3s6jn****
	AutoSnapshotPolicyId *string `json:"auto_snapshot_policy_id,omitempty" xml:"auto_snapshot_policy_id,omitempty"`
	// The type of the data disk.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Specifies whether to encrypt the data disks. Valid values:
	//
	// 	- `true`: encrypts the data disk.
	//
	// 	- `false`: does not encrypt the data disk.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// The size of the data disk. Valid values: 40 to 32767.
	//
	// example:
	//
	// 120
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ScaleOutClusterRequestWorkerDataDisks) String() string {
	return dara.Prettify(s)
}

func (s ScaleOutClusterRequestWorkerDataDisks) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterRequestWorkerDataDisks) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *ScaleOutClusterRequestWorkerDataDisks) GetCategory() *string {
	return s.Category
}

func (s *ScaleOutClusterRequestWorkerDataDisks) GetEncrypted() *string {
	return s.Encrypted
}

func (s *ScaleOutClusterRequestWorkerDataDisks) GetSize() *string {
	return s.Size
}

func (s *ScaleOutClusterRequestWorkerDataDisks) SetAutoSnapshotPolicyId(v string) *ScaleOutClusterRequestWorkerDataDisks {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ScaleOutClusterRequestWorkerDataDisks) SetCategory(v string) *ScaleOutClusterRequestWorkerDataDisks {
	s.Category = &v
	return s
}

func (s *ScaleOutClusterRequestWorkerDataDisks) SetEncrypted(v string) *ScaleOutClusterRequestWorkerDataDisks {
	s.Encrypted = &v
	return s
}

func (s *ScaleOutClusterRequestWorkerDataDisks) SetSize(v string) *ScaleOutClusterRequestWorkerDataDisks {
	s.Size = &v
	return s
}

func (s *ScaleOutClusterRequestWorkerDataDisks) Validate() error {
	return dara.Validate(s)
}
