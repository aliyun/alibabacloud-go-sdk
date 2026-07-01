// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchOptimization(v string) *StopInstancesRequest
	GetBatchOptimization() *string
	SetDryRun(v bool) *StopInstancesRequest
	GetDryRun() *bool
	SetForceStop(v bool) *StopInstancesRequest
	GetForceStop() *bool
	SetInstanceId(v []*string) *StopInstancesRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *StopInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StopInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StopInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StopInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StopInstancesRequest
	GetResourceOwnerId() *int64
	SetStoppedMode(v string) *StopInstancesRequest
	GetStoppedMode() *string
}

type StopInstancesRequest struct {
	// The batch operation mode. Valid values:
	//
	// - AllTogether: All operations must succeed for the entire batch operation to be considered successful. If any operation fails, the entire batch operation fails and all completed operations are rolled back to the previous state.
	//
	// - SuccessFirst: Each operation in the batch is executed independently. If an operation fails, other operations can still be executed and confirmed as successful. Successful operations are committed, and failed operations are marked as failed without affecting the results of other operations.
	//
	// Default value: AllTogether.
	//
	// example:
	//
	// AllTogether
	BatchOptimization *string `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	// Specifies whether to send a dry run request. Valid values:
	//
	// - true: sends a dry run request without stopping the instances. The system checks the required parameters, request format, and instance status. If the check fails, the corresponding error is returned. If the check succeeds, `DRYRUN.SUCCESS` is returned.
	//
	// > If the BatchOptimization parameter is set to `SuccessFirst`, the dry run result for `DryRun=true` returns only `DRYRUN.SUCCESS`.
	//
	// - false: sends a normal request. After the request passes the check, the instances are stopped.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully stop the instances. Valid values:
	//
	// - true: forcefully stops the instances.
	//
	//   	Warning: A forced stop is equivalent to a power-off. Data that is not written to disks in the instance operating system may be lost. Proceed with caution.
	//
	// - false: normally stops the instances.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The instance IDs. Array length: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId   []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instances. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The stop mode. Valid values:
	//
	//   - StopCharging: economical mode. After economical mode is enabled:
	//
	//     - Billing for compute resources (vCPUs, memory, and GPUs), image license fees, and fixed bandwidth of static public IP addresses is suspended.
	//
	//     - Billing for system disks, data disks, and fixed bandwidth of Elastic IP Addresses (EIPs) continues.
	//
	//     - Because compute resources are released, the instance may fail to start due to insufficient resources. Try again later or change the instance type.
	//
	//     - If an EIP is associated with the instance before it is stopped, the IP address remains unchanged after the instance is restarted. Otherwise, the static public IP address may change, but the private IP address remains unchanged.
	//
	//     For more information, see [Economical mode](https://help.aliyun.com/document_detail/63353.html).
	//
	//     	Notice:
	//
	// If the instance does not support economical mode, the API does not return an error. Stopping the instance takes priority. Instance types that do not support economical mode include instances with local SSDs and subscription instances.
	//
	//
	//
	//   - KeepCharging: standard stop mode. After the instance is stopped, resources are retained and billing continues. The instance type inventory and public IP address are also retained. If you stop the instance to replace the operating system, reinitialize a disk, change the instance type, or modify the private IP address, select this mode to avoid startup failures.
	//
	// Default value: If you [enable economical mode for VPC-connected instances](~~63353#default~~) and the conditions are met, the default value is `StopCharging`. Otherwise, the default value is `KeepCharging`.
	//
	// example:
	//
	// KeepCharging
	StoppedMode *string `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
}

func (s StopInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstancesRequest) GoString() string {
	return s.String()
}

func (s *StopInstancesRequest) GetBatchOptimization() *string {
	return s.BatchOptimization
}

func (s *StopInstancesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StopInstancesRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *StopInstancesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *StopInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StopInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopInstancesRequest) GetStoppedMode() *string {
	return s.StoppedMode
}

func (s *StopInstancesRequest) SetBatchOptimization(v string) *StopInstancesRequest {
	s.BatchOptimization = &v
	return s
}

func (s *StopInstancesRequest) SetDryRun(v bool) *StopInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *StopInstancesRequest) SetForceStop(v bool) *StopInstancesRequest {
	s.ForceStop = &v
	return s
}

func (s *StopInstancesRequest) SetInstanceId(v []*string) *StopInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *StopInstancesRequest) SetOwnerAccount(v string) *StopInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StopInstancesRequest) SetOwnerId(v int64) *StopInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *StopInstancesRequest) SetRegionId(v string) *StopInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *StopInstancesRequest) SetResourceOwnerAccount(v string) *StopInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopInstancesRequest) SetResourceOwnerId(v int64) *StopInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopInstancesRequest) SetStoppedMode(v string) *StopInstancesRequest {
	s.StoppedMode = &v
	return s
}

func (s *StopInstancesRequest) Validate() error {
	return dara.Validate(s)
}
