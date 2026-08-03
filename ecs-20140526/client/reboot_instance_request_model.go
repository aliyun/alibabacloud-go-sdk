// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *RebootInstanceRequest
	GetDryRun() *bool
	SetForceStop(v bool) *RebootInstanceRequest
	GetForceStop() *bool
	SetInstanceId(v string) *RebootInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *RebootInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RebootInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RebootInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RebootInstanceRequest
	GetResourceOwnerId() *int64
}

type RebootInstanceRequest struct {
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - true: Performs only a dry run without restarting the instance. The system checks the required parameters, request syntax, business restrictions, and ECS inventory. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - false: Performs a dry run and sends the request. If the check succeeds, the instance is restarted.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully stop ECS instance before restarting it. Valid values:
	//
	// -   true: Forcefully stops ECS instance. This is equivalent to a power-off operation. Cached data that has not been written to storage devices is lost.
	//
	// -   false: Normally stops ECS instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The ID of the instance.
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
}

func (s RebootInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RebootInstanceRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *RebootInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RebootInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RebootInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RebootInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RebootInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RebootInstanceRequest) SetDryRun(v bool) *RebootInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *RebootInstanceRequest) SetForceStop(v bool) *RebootInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *RebootInstanceRequest) SetInstanceId(v string) *RebootInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootInstanceRequest) SetOwnerAccount(v string) *RebootInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RebootInstanceRequest) SetOwnerId(v int64) *RebootInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RebootInstanceRequest) SetResourceOwnerAccount(v string) *RebootInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RebootInstanceRequest) SetResourceOwnerId(v int64) *RebootInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RebootInstanceRequest) Validate() error {
	return dara.Validate(s)
}
