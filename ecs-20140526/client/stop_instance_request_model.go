// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfirmStop(v bool) *StopInstanceRequest
	GetConfirmStop() *bool
	SetDryRun(v bool) *StopInstanceRequest
	GetDryRun() *bool
	SetForceStop(v bool) *StopInstanceRequest
	GetForceStop() *bool
	SetHibernate(v bool) *StopInstanceRequest
	GetHibernate() *bool
	SetInstanceId(v string) *StopInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *StopInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StopInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StopInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StopInstanceRequest
	GetResourceOwnerId() *int64
	SetStoppedMode(v string) *StopInstanceRequest
	GetStoppedMode() *string
}

type StopInstanceRequest struct {
	// This parameter is being deprecated and is retained only for compatibility purposes. Ignore this parameter when you call this operation.
	//
	// example:
	//
	// true
	ConfirmStop *bool `json:"ConfirmStop,omitempty" xml:"ConfirmStop,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - true: Performs a dry run without stopping the instance. The system checks whether the required parameters are specified, the request format is valid, service limits are met, and ECS inventory is sufficient. If the check fails, the corresponding error is returned. If the check passes, the error code `DryRunOperation` is returned.
	//
	// - false: Performs a normal request. After the check passes, the instance is stopped.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully stop the instance. Valid values:
	//
	// - true: Forcefully stops the instance. This is equivalent to a typical power-off operation. All cached data that is not written to the storage device is lost.
	//
	// - false: Normally stops the instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// >This parameter is in invitational preview and is not available for use.
	//
	// example:
	//
	// hide
	Hibernate *bool `json:"Hibernate,omitempty" xml:"Hibernate,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The stop mode for a pay-as-you-go ECS instance. Valid values:
	//
	//   - StopCharging: Economical mode. After economical mode is enabled:
	//
	//     - Billing is suspended for compute resources (vCPUs, memory, and GPUs), image license fees, and the pay-by-bandwidth mode for static public IP addresses.
	//
	//     - Billing continues for system disks, data disks, and the pay-by-bandwidth mode for elastic IP addresses (EIPs).
	//
	//     - Because compute resources are reclaimed, the instance may fail to start due to insufficient inventory. In this case, try again later or change the instance type.
	//
	//     - If an EIP is associated with the instance before the instance is stopped, the IP address remains unchanged after the instance is restarted. Otherwise, the static public IP address may change, but the private IP address remains unchanged.
	//
	//     For more information, see [Economical mode](https://help.aliyun.com/document_detail/63353.html).
	//
	//     	Notice:
	//
	// If the instance does not support economical mode, the API does not return an error. The instance is stopped as a priority. Instance types that do not support economical mode include instances with local disks and subscription instances.
	//
	//
	//
	//   - KeepCharging: Standard stop mode. Billing continues after the instance is stopped.
	//
	// Default value: If you enable the economical mode for instances in a VPC in the ECS console (for more information, see [Enable economical mode by default](~~63353#default~~)) and the conditions are met, the default value is `StopCharging`. Otherwise, the default value is `KeepCharging`.
	//
	// example:
	//
	// KeepCharging
	StoppedMode *string `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) GetConfirmStop() *bool {
	return s.ConfirmStop
}

func (s *StopInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StopInstanceRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *StopInstanceRequest) GetHibernate() *bool {
	return s.Hibernate
}

func (s *StopInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StopInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopInstanceRequest) GetStoppedMode() *string {
	return s.StoppedMode
}

func (s *StopInstanceRequest) SetConfirmStop(v bool) *StopInstanceRequest {
	s.ConfirmStop = &v
	return s
}

func (s *StopInstanceRequest) SetDryRun(v bool) *StopInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *StopInstanceRequest) SetForceStop(v bool) *StopInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *StopInstanceRequest) SetHibernate(v bool) *StopInstanceRequest {
	s.Hibernate = &v
	return s
}

func (s *StopInstanceRequest) SetInstanceId(v string) *StopInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceRequest) SetOwnerAccount(v string) *StopInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StopInstanceRequest) SetOwnerId(v int64) *StopInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *StopInstanceRequest) SetResourceOwnerAccount(v string) *StopInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopInstanceRequest) SetResourceOwnerId(v int64) *StopInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopInstanceRequest) SetStoppedMode(v string) *StopInstanceRequest {
	s.StoppedMode = &v
	return s
}

func (s *StopInstanceRequest) Validate() error {
	return dara.Validate(s)
}
