// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportInstancesStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ReportInstancesStatusRequest
	GetDescription() *string
	SetDevice(v []*string) *ReportInstancesStatusRequest
	GetDevice() []*string
	SetDiskId(v []*string) *ReportInstancesStatusRequest
	GetDiskId() []*string
	SetEndTime(v string) *ReportInstancesStatusRequest
	GetEndTime() *string
	SetInstanceId(v []*string) *ReportInstancesStatusRequest
	GetInstanceId() []*string
	SetIssueCategory(v string) *ReportInstancesStatusRequest
	GetIssueCategory() *string
	SetOwnerAccount(v string) *ReportInstancesStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReportInstancesStatusRequest
	GetOwnerId() *int64
	SetReason(v string) *ReportInstancesStatusRequest
	GetReason() *string
	SetRegionId(v string) *ReportInstancesStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ReportInstancesStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReportInstancesStatusRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *ReportInstancesStatusRequest
	GetStartTime() *string
}

type ReportInstancesStatusRequest struct {
	// The detailed description of the anomalous issue.
	//
	// This parameter is required.
	//
	// example:
	//
	// 本地盘不可用，挂载点拒绝访问，无法加载文件。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of device names of the disks that have the same anomalous issue and are attached to the instance. You can specify up to 100 device names.
	//
	// If you are using an ECS Bare Metal server instance, specify the SLOT information list of the disk devices.
	//
	// > For ECS bare metal instances, this parameter is required when the `Reason` parameter is set to `abnormal-local-disk` or `abnormal-cloud-disk`, or when the `IssueCategory` parameter is set to `hardware-disk-error`.
	//
	// example:
	//
	// /dev/xvdb
	Device []*string `json:"Device,omitempty" xml:"Device,omitempty" type:"Repeated"`
	// The list of IDs of the disks that have the same anomalous issue. You can specify up to 100 disk IDs. If you are using an ECS Bare Metal server instance, specify the SN list of the disk devices.
	//
	// > This parameter is required when the `Reason` parameter is set to `abnormal-local-disk` or `abnormal-cloud-disk`, or when the `IssueCategory` parameter is set to `hardware-disk-error`.
	//
	// example:
	//
	// d-bp1aeljlfad7x6u1****
	DiskId []*string `json:"DiskId,omitempty" xml:"DiskId,omitempty" type:"Repeated"`
	// The time when the instance failures ended. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-11-31T06:32:31Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of ECS instance IDs. You can specify up to 100 instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp165p6xk2tmdhj0****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The category of the anomalous issue. This parameter is applicable only to Elastic Compute Service Bare Metal Instance instances. Valid values:
	//
	// - hardware-cpu-error: CPU failure.
	//
	// - hardware-motherboard-error: Motherboard failure.
	//
	// - hardware-mem-error: Memory failure.
	//
	// - hardware-power-error: Power failure.
	//
	// - hardware-disk-error: Disk failure.
	//
	// - hardware-networkcard-error: Network interface controller (NIC) failure.
	//
	// - hardware-raidcard-error: SAS/RAID card failure.
	//
	// - hardware-fan-error: Fan failure.
	//
	// - others: Other failures.
	//
	// example:
	//
	// hardware-cpu-error
	IssueCategory *string `json:"IssueCategory,omitempty" xml:"IssueCategory,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The impact of the anomalous issue on the ECS instance. Valid values:
	//
	// - instance-hang: The ECS instance is unavailable or cannot be connected to.
	//
	// - instance-stuck-in-status: The ECS instance is stuck in a specific state, such as Starting or Stopping, for an extended period of time.
	//
	// - abnormal-network: A network exception occurred on the ECS instance.
	//
	// - abnormal-local-disk: A local disk attached to the ECS instance is abnormal.
	//
	// - abnormal-cloud-disk: A cloud disk or Shared Block Storage device attached to the ECS instance is abnormal.
	//
	// - others: Other exception types. If none of the preceding values apply, set `Reason=others` and provide more information in `Description`.
	//
	// example:
	//
	// abnormal-local-disk
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The time when the instance failures started. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-11-30T06:32:31Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ReportInstancesStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportInstancesStatusRequest) GoString() string {
	return s.String()
}

func (s *ReportInstancesStatusRequest) GetDescription() *string {
	return s.Description
}

func (s *ReportInstancesStatusRequest) GetDevice() []*string {
	return s.Device
}

func (s *ReportInstancesStatusRequest) GetDiskId() []*string {
	return s.DiskId
}

func (s *ReportInstancesStatusRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ReportInstancesStatusRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *ReportInstancesStatusRequest) GetIssueCategory() *string {
	return s.IssueCategory
}

func (s *ReportInstancesStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReportInstancesStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReportInstancesStatusRequest) GetReason() *string {
	return s.Reason
}

func (s *ReportInstancesStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReportInstancesStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReportInstancesStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReportInstancesStatusRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ReportInstancesStatusRequest) SetDescription(v string) *ReportInstancesStatusRequest {
	s.Description = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetDevice(v []*string) *ReportInstancesStatusRequest {
	s.Device = v
	return s
}

func (s *ReportInstancesStatusRequest) SetDiskId(v []*string) *ReportInstancesStatusRequest {
	s.DiskId = v
	return s
}

func (s *ReportInstancesStatusRequest) SetEndTime(v string) *ReportInstancesStatusRequest {
	s.EndTime = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetInstanceId(v []*string) *ReportInstancesStatusRequest {
	s.InstanceId = v
	return s
}

func (s *ReportInstancesStatusRequest) SetIssueCategory(v string) *ReportInstancesStatusRequest {
	s.IssueCategory = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetOwnerAccount(v string) *ReportInstancesStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetOwnerId(v int64) *ReportInstancesStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetReason(v string) *ReportInstancesStatusRequest {
	s.Reason = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetRegionId(v string) *ReportInstancesStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetResourceOwnerAccount(v string) *ReportInstancesStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetResourceOwnerId(v int64) *ReportInstancesStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetStartTime(v string) *ReportInstancesStatusRequest {
	s.StartTime = &v
	return s
}

func (s *ReportInstancesStatusRequest) Validate() error {
	return dara.Validate(s)
}
