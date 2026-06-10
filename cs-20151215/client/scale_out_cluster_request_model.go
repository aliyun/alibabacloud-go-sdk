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
	// example:
	//
	// true
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// example:
	//
	// centos_7_7_x64_20G_alibase_20200426.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// key_name
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hello1234
	LoginPassword *string   `json:"login_password,omitempty" xml:"login_password,omitempty"`
	RdsInstances  []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	Runtime       *Runtime  `json:"runtime,omitempty" xml:"runtime,omitempty"`
	Tags          []*Tag    `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	Taints        []*Taint  `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// example:
	//
	// IyEvdXNyL2Jpbi9iYXNoCmVjaG8gIkhlbGxvIEFDSyEi
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// This parameter is required.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// example:
	//
	// true
	WorkerAutoRenew *bool `json:"worker_auto_renew,omitempty" xml:"worker_auto_renew,omitempty"`
	// example:
	//
	// 1
	WorkerAutoRenewPeriod *int64                                   `json:"worker_auto_renew_period,omitempty" xml:"worker_auto_renew_period,omitempty"`
	WorkerDataDisks       []*ScaleOutClusterRequestWorkerDataDisks `json:"worker_data_disks,omitempty" xml:"worker_data_disks,omitempty" type:"Repeated"`
	// example:
	//
	// PrePaid
	WorkerInstanceChargeType *string `json:"worker_instance_charge_type,omitempty" xml:"worker_instance_charge_type,omitempty"`
	// This parameter is required.
	WorkerInstanceTypes []*string `json:"worker_instance_types,omitempty" xml:"worker_instance_types,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	WorkerPeriod *int64 `json:"worker_period,omitempty" xml:"worker_period,omitempty"`
	// example:
	//
	// Month
	WorkerPeriodUnit *string `json:"worker_period_unit,omitempty" xml:"worker_period_unit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cloud_efficiency
	WorkerSystemDiskCategory *string `json:"worker_system_disk_category,omitempty" xml:"worker_system_disk_category,omitempty"`
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
	// example:
	//
	// sp-2zej1nogjvovnz4z****
	AutoSnapshotPolicyId *string `json:"auto_snapshot_policy_id,omitempty" xml:"auto_snapshot_policy_id,omitempty"`
	// example:
	//
	// cloud_efficiency
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// true
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
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
