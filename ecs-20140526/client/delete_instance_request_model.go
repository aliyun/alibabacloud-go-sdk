// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DeleteInstanceRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteInstanceRequest
	GetForce() *bool
	SetForceStop(v bool) *DeleteInstanceRequest
	GetForceStop() *bool
	SetInstanceId(v string) *DeleteInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DeleteInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteInstanceRequest
	GetResourceOwnerId() *int64
	SetTerminateSubscription(v bool) *DeleteInstanceRequest
	GetTerminateSubscription() *bool
}

type DeleteInstanceRequest struct {
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - true: Sends a check request without releasing the instance. The system checks whether the required parameters are specified, the request format is valid, business requirements are met, and ECS resources are sufficient. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - false (default): Sends a normal request. After the request passes the check, the instance is directly deleted.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully release a **running*	- (`Running`) instance.
	//
	// - true: Forcefully releases a **running*	- (`Running`) instance.
	//
	// - false: Releases the instance in the normal way. The instance must be in the **Stopped*	- (`Stopped`) state.
	//
	// Default value: false.
	//
	// 	Warning: A forceful release is equivalent to a power-off. Temporary data in the instance memory and storage is erased and cannot be recovered..
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// Specifies whether to use the forced shutdown policy when releasing a **running*	- (`Running`) instance. This parameter takes effect only when `Force=true`. Valid values:
	//
	// - true: Forcefully shuts down and releases the instance. This is equivalent to a typical power-off operation. The instance directly enters the resource release process.
	//
	// 	Warning: A forceful release is equivalent to a power-off. Temporary data in the instance memory and storage is erased and cannot be recovered.
	//
	// - false: Before the instance is released, the system preferentially performs a standard shutdown process. This mode causes the instance release to take several minutes. You can configure service draining actions during the operating system shutdown to reduce noise in your business systems.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1g6zv0ce8oghu7****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to release an expired subscription instance.
	//
	// - true: Releases the instance.
	//
	// - false: Does not release the instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	TerminateSubscription *bool `json:"TerminateSubscription,omitempty" xml:"TerminateSubscription,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteInstanceRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteInstanceRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *DeleteInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteInstanceRequest) GetTerminateSubscription() *bool {
	return s.TerminateSubscription
}

func (s *DeleteInstanceRequest) SetDryRun(v bool) *DeleteInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteInstanceRequest) SetForce(v bool) *DeleteInstanceRequest {
	s.Force = &v
	return s
}

func (s *DeleteInstanceRequest) SetForceStop(v bool) *DeleteInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetOwnerAccount(v string) *DeleteInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteInstanceRequest) SetOwnerId(v int64) *DeleteInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceOwnerAccount(v string) *DeleteInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceOwnerId(v int64) *DeleteInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteInstanceRequest) SetTerminateSubscription(v bool) *DeleteInstanceRequest {
	s.TerminateSubscription = &v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
