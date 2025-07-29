// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudMonitorFlags(v bool) *ScaleClusterRequest
	GetCloudMonitorFlags() *bool
	SetCount(v int64) *ScaleClusterRequest
	GetCount() *int64
	SetCpuPolicy(v string) *ScaleClusterRequest
	GetCpuPolicy() *string
	SetDisableRollback(v bool) *ScaleClusterRequest
	GetDisableRollback() *bool
	SetKeyPair(v string) *ScaleClusterRequest
	GetKeyPair() *string
	SetLoginPassword(v string) *ScaleClusterRequest
	GetLoginPassword() *string
	SetTags(v []*ScaleClusterRequestTags) *ScaleClusterRequest
	GetTags() []*ScaleClusterRequestTags
	SetTaints(v []*ScaleClusterRequestTaints) *ScaleClusterRequest
	GetTaints() []*ScaleClusterRequestTaints
	SetVswitchIds(v []*string) *ScaleClusterRequest
	GetVswitchIds() []*string
	SetWorkerAutoRenew(v bool) *ScaleClusterRequest
	GetWorkerAutoRenew() *bool
	SetWorkerAutoRenewPeriod(v int64) *ScaleClusterRequest
	GetWorkerAutoRenewPeriod() *int64
	SetWorkerDataDisk(v bool) *ScaleClusterRequest
	GetWorkerDataDisk() *bool
	SetWorkerDataDisks(v []*ScaleClusterRequestWorkerDataDisks) *ScaleClusterRequest
	GetWorkerDataDisks() []*ScaleClusterRequestWorkerDataDisks
	SetWorkerInstanceChargeType(v string) *ScaleClusterRequest
	GetWorkerInstanceChargeType() *string
	SetWorkerInstanceTypes(v []*string) *ScaleClusterRequest
	GetWorkerInstanceTypes() []*string
	SetWorkerPeriod(v int64) *ScaleClusterRequest
	GetWorkerPeriod() *int64
	SetWorkerPeriodUnit(v string) *ScaleClusterRequest
	GetWorkerPeriodUnit() *string
	SetWorkerSystemDiskCategory(v string) *ScaleClusterRequest
	GetWorkerSystemDiskCategory() *string
	SetWorkerSystemDiskSize(v int64) *ScaleClusterRequest
	GetWorkerSystemDiskSize() *int64
}

type ScaleClusterRequest struct {
	CloudMonitorFlags        *bool                                 `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	Count                    *int64                                `json:"count,omitempty" xml:"count,omitempty"`
	CpuPolicy                *string                               `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	DisableRollback          *bool                                 `json:"disable_rollback,omitempty" xml:"disable_rollback,omitempty"`
	KeyPair                  *string                               `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	LoginPassword            *string                               `json:"login_password,omitempty" xml:"login_password,omitempty"`
	Tags                     []*ScaleClusterRequestTags            `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	Taints                   []*ScaleClusterRequestTaints          `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	VswitchIds               []*string                             `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	WorkerAutoRenew          *bool                                 `json:"worker_auto_renew,omitempty" xml:"worker_auto_renew,omitempty"`
	WorkerAutoRenewPeriod    *int64                                `json:"worker_auto_renew_period,omitempty" xml:"worker_auto_renew_period,omitempty"`
	WorkerDataDisk           *bool                                 `json:"worker_data_disk,omitempty" xml:"worker_data_disk,omitempty"`
	WorkerDataDisks          []*ScaleClusterRequestWorkerDataDisks `json:"worker_data_disks,omitempty" xml:"worker_data_disks,omitempty" type:"Repeated"`
	WorkerInstanceChargeType *string                               `json:"worker_instance_charge_type,omitempty" xml:"worker_instance_charge_type,omitempty"`
	WorkerInstanceTypes      []*string                             `json:"worker_instance_types,omitempty" xml:"worker_instance_types,omitempty" type:"Repeated"`
	WorkerPeriod             *int64                                `json:"worker_period,omitempty" xml:"worker_period,omitempty"`
	WorkerPeriodUnit         *string                               `json:"worker_period_unit,omitempty" xml:"worker_period_unit,omitempty"`
	WorkerSystemDiskCategory *string                               `json:"worker_system_disk_category,omitempty" xml:"worker_system_disk_category,omitempty"`
	WorkerSystemDiskSize     *int64                                `json:"worker_system_disk_size,omitempty" xml:"worker_system_disk_size,omitempty"`
}

func (s ScaleClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterRequest) GoString() string {
	return s.String()
}

func (s *ScaleClusterRequest) GetCloudMonitorFlags() *bool {
	return s.CloudMonitorFlags
}

func (s *ScaleClusterRequest) GetCount() *int64 {
	return s.Count
}

func (s *ScaleClusterRequest) GetCpuPolicy() *string {
	return s.CpuPolicy
}

func (s *ScaleClusterRequest) GetDisableRollback() *bool {
	return s.DisableRollback
}

func (s *ScaleClusterRequest) GetKeyPair() *string {
	return s.KeyPair
}

func (s *ScaleClusterRequest) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *ScaleClusterRequest) GetTags() []*ScaleClusterRequestTags {
	return s.Tags
}

func (s *ScaleClusterRequest) GetTaints() []*ScaleClusterRequestTaints {
	return s.Taints
}

func (s *ScaleClusterRequest) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *ScaleClusterRequest) GetWorkerAutoRenew() *bool {
	return s.WorkerAutoRenew
}

func (s *ScaleClusterRequest) GetWorkerAutoRenewPeriod() *int64 {
	return s.WorkerAutoRenewPeriod
}

func (s *ScaleClusterRequest) GetWorkerDataDisk() *bool {
	return s.WorkerDataDisk
}

func (s *ScaleClusterRequest) GetWorkerDataDisks() []*ScaleClusterRequestWorkerDataDisks {
	return s.WorkerDataDisks
}

func (s *ScaleClusterRequest) GetWorkerInstanceChargeType() *string {
	return s.WorkerInstanceChargeType
}

func (s *ScaleClusterRequest) GetWorkerInstanceTypes() []*string {
	return s.WorkerInstanceTypes
}

func (s *ScaleClusterRequest) GetWorkerPeriod() *int64 {
	return s.WorkerPeriod
}

func (s *ScaleClusterRequest) GetWorkerPeriodUnit() *string {
	return s.WorkerPeriodUnit
}

func (s *ScaleClusterRequest) GetWorkerSystemDiskCategory() *string {
	return s.WorkerSystemDiskCategory
}

func (s *ScaleClusterRequest) GetWorkerSystemDiskSize() *int64 {
	return s.WorkerSystemDiskSize
}

func (s *ScaleClusterRequest) SetCloudMonitorFlags(v bool) *ScaleClusterRequest {
	s.CloudMonitorFlags = &v
	return s
}

func (s *ScaleClusterRequest) SetCount(v int64) *ScaleClusterRequest {
	s.Count = &v
	return s
}

func (s *ScaleClusterRequest) SetCpuPolicy(v string) *ScaleClusterRequest {
	s.CpuPolicy = &v
	return s
}

func (s *ScaleClusterRequest) SetDisableRollback(v bool) *ScaleClusterRequest {
	s.DisableRollback = &v
	return s
}

func (s *ScaleClusterRequest) SetKeyPair(v string) *ScaleClusterRequest {
	s.KeyPair = &v
	return s
}

func (s *ScaleClusterRequest) SetLoginPassword(v string) *ScaleClusterRequest {
	s.LoginPassword = &v
	return s
}

func (s *ScaleClusterRequest) SetTags(v []*ScaleClusterRequestTags) *ScaleClusterRequest {
	s.Tags = v
	return s
}

func (s *ScaleClusterRequest) SetTaints(v []*ScaleClusterRequestTaints) *ScaleClusterRequest {
	s.Taints = v
	return s
}

func (s *ScaleClusterRequest) SetVswitchIds(v []*string) *ScaleClusterRequest {
	s.VswitchIds = v
	return s
}

func (s *ScaleClusterRequest) SetWorkerAutoRenew(v bool) *ScaleClusterRequest {
	s.WorkerAutoRenew = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerAutoRenewPeriod(v int64) *ScaleClusterRequest {
	s.WorkerAutoRenewPeriod = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerDataDisk(v bool) *ScaleClusterRequest {
	s.WorkerDataDisk = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerDataDisks(v []*ScaleClusterRequestWorkerDataDisks) *ScaleClusterRequest {
	s.WorkerDataDisks = v
	return s
}

func (s *ScaleClusterRequest) SetWorkerInstanceChargeType(v string) *ScaleClusterRequest {
	s.WorkerInstanceChargeType = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerInstanceTypes(v []*string) *ScaleClusterRequest {
	s.WorkerInstanceTypes = v
	return s
}

func (s *ScaleClusterRequest) SetWorkerPeriod(v int64) *ScaleClusterRequest {
	s.WorkerPeriod = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerPeriodUnit(v string) *ScaleClusterRequest {
	s.WorkerPeriodUnit = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerSystemDiskCategory(v string) *ScaleClusterRequest {
	s.WorkerSystemDiskCategory = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerSystemDiskSize(v int64) *ScaleClusterRequest {
	s.WorkerSystemDiskSize = &v
	return s
}

func (s *ScaleClusterRequest) Validate() error {
	return dara.Validate(s)
}

type ScaleClusterRequestTags struct {
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ScaleClusterRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterRequestTags) GoString() string {
	return s.String()
}

func (s *ScaleClusterRequestTags) GetKey() *string {
	return s.Key
}

func (s *ScaleClusterRequestTags) SetKey(v string) *ScaleClusterRequestTags {
	s.Key = &v
	return s
}

func (s *ScaleClusterRequestTags) Validate() error {
	return dara.Validate(s)
}

type ScaleClusterRequestTaints struct {
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	Key    *string `json:"key,omitempty" xml:"key,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ScaleClusterRequestTaints) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterRequestTaints) GoString() string {
	return s.String()
}

func (s *ScaleClusterRequestTaints) GetEffect() *string {
	return s.Effect
}

func (s *ScaleClusterRequestTaints) GetKey() *string {
	return s.Key
}

func (s *ScaleClusterRequestTaints) GetValue() *string {
	return s.Value
}

func (s *ScaleClusterRequestTaints) SetEffect(v string) *ScaleClusterRequestTaints {
	s.Effect = &v
	return s
}

func (s *ScaleClusterRequestTaints) SetKey(v string) *ScaleClusterRequestTaints {
	s.Key = &v
	return s
}

func (s *ScaleClusterRequestTaints) SetValue(v string) *ScaleClusterRequestTaints {
	s.Value = &v
	return s
}

func (s *ScaleClusterRequestTaints) Validate() error {
	return dara.Validate(s)
}

type ScaleClusterRequestWorkerDataDisks struct {
	Category  *string `json:"category,omitempty" xml:"category,omitempty"`
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	Size      *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ScaleClusterRequestWorkerDataDisks) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterRequestWorkerDataDisks) GoString() string {
	return s.String()
}

func (s *ScaleClusterRequestWorkerDataDisks) GetCategory() *string {
	return s.Category
}

func (s *ScaleClusterRequestWorkerDataDisks) GetEncrypted() *string {
	return s.Encrypted
}

func (s *ScaleClusterRequestWorkerDataDisks) GetSize() *string {
	return s.Size
}

func (s *ScaleClusterRequestWorkerDataDisks) SetCategory(v string) *ScaleClusterRequestWorkerDataDisks {
	s.Category = &v
	return s
}

func (s *ScaleClusterRequestWorkerDataDisks) SetEncrypted(v string) *ScaleClusterRequestWorkerDataDisks {
	s.Encrypted = &v
	return s
}

func (s *ScaleClusterRequestWorkerDataDisks) SetSize(v string) *ScaleClusterRequestWorkerDataDisks {
	s.Size = &v
	return s
}

func (s *ScaleClusterRequestWorkerDataDisks) Validate() error {
	return dara.Validate(s)
}
