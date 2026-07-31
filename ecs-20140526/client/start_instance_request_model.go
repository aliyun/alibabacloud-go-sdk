// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *StartInstanceRequest
	GetDryRun() *bool
	SetInitLocalDisk(v bool) *StartInstanceRequest
	GetInitLocalDisk() *bool
	SetInstanceId(v string) *StartInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *StartInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StartInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StartInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartInstanceRequest
	GetResourceOwnerId() *int64
}

type StartInstanceRequest struct {
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - true: performs only a dry run. The instance is not started. The system checks whether the AccessKey pair is valid, whether Resource Access Management (RAM) users are granted required permissions, and whether the required parameters are specified. If the check fails, the corresponding error is returned. If the check succeeds, the DryRunOperation error code is returned.
	//
	// - false: performs a dry run and sends the request. If the check succeeds, a 2XX HTTP status code is returned and the instance is started.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to recover the instance to its initial health state when a local disk fails. This parameter is applicable to instances that use local disks, such as instances in the d1, i1, or i2 instance families. Valid values:
	//
	// - true: Recovers the instance to its initial health state.
	//
	// 	Warning: All data stored on the local disks of the instance is lost.
	//
	// - false: Does not perform any action and maintains the current state.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	InitLocalDisk *bool `json:"InitLocalDisk,omitempty" xml:"InitLocalDisk,omitempty"`
	// The instance ID of the instance that you want to start.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StartInstanceRequest) GetInitLocalDisk() *bool {
	return s.InitLocalDisk
}

func (s *StartInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StartInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartInstanceRequest) SetDryRun(v bool) *StartInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *StartInstanceRequest) SetInitLocalDisk(v bool) *StartInstanceRequest {
	s.InitLocalDisk = &v
	return s
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceRequest) SetOwnerAccount(v string) *StartInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartInstanceRequest) SetOwnerId(v int64) *StartInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *StartInstanceRequest) SetResourceOwnerAccount(v string) *StartInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartInstanceRequest) SetResourceOwnerId(v int64) *StartInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
