// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteInstancesRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteInstancesRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteInstancesRequest
	GetForce() *bool
	SetForceStop(v bool) *DeleteInstancesRequest
	GetForceStop() *bool
	SetInstanceId(v []*string) *DeleteInstancesRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *DeleteInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteInstancesRequest
	GetResourceOwnerId() *int64
	SetTerminateSubscription(v bool) *DeleteInstancesRequest
	GetTerminateSubscription() *bool
}

type DeleteInstancesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request.
	//
	// - true: sends a check request without querying resource status. The check items include whether your AccessKey pair is valid, whether Resource Access Management (RAM) user authorization is granted, and whether the required parameters are specified. If the check fails, the corresponding error is returned. If the check succeeds, the DRYRUN.SUCCESS error code is returned.
	//
	// - false: sends a Normal request. After the check succeeds, a 2xx HTTP status code is returned and the resource status is queried directly.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully release an instance that is in the **Running*	- (`Running`) state.
	//
	// - true: forcefully releases ECS instance that is in the **Running*	- (`Running`) state.
	//
	// - false: releases ECS instance only when it is in the **Stopped*	- (`Stopped`) state.
	//
	// Default value: false.
	//
	// 	Warning: Forceful release is equivalent to powering off ECS instance. All in-memory data and temporary data in the storage are erased and cannot be recovered..
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// Specifies whether to forcefully shut down the instance before release when the instance is in the **Running*	- (`Running`) state. This parameter takes effect only when `Force=true`. Valid values:
	//
	// - true: forcefully shuts down and releases the instance. This is equivalent to a power-off operation. The instance directly enters the resource release process.
	//
	// 	Warning: Forceful release is equivalent to powering off the instance. All in-memory data and temporary data in the storage are erased and cannot be recovered.
	//
	// - false: performs a standard shutdown before releasing the instance. This mode causes the release process to take several minutes. You can configure service draining actions during the operating system shutdown to reduce noise in your business systems.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The instance ID array. Array length: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1g6zv0ce8oghu7****
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
	// Specifies whether to release an expired subscription instance.
	//
	// - true: releases the instance.
	//
	// - false: does not release the instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	TerminateSubscription *bool `json:"TerminateSubscription,omitempty" xml:"TerminateSubscription,omitempty"`
}

func (s DeleteInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteInstancesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteInstancesRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteInstancesRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *DeleteInstancesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DeleteInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteInstancesRequest) GetTerminateSubscription() *bool {
	return s.TerminateSubscription
}

func (s *DeleteInstancesRequest) SetClientToken(v string) *DeleteInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteInstancesRequest) SetDryRun(v bool) *DeleteInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteInstancesRequest) SetForce(v bool) *DeleteInstancesRequest {
	s.Force = &v
	return s
}

func (s *DeleteInstancesRequest) SetForceStop(v bool) *DeleteInstancesRequest {
	s.ForceStop = &v
	return s
}

func (s *DeleteInstancesRequest) SetInstanceId(v []*string) *DeleteInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DeleteInstancesRequest) SetOwnerAccount(v string) *DeleteInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteInstancesRequest) SetOwnerId(v int64) *DeleteInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteInstancesRequest) SetRegionId(v string) *DeleteInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteInstancesRequest) SetResourceOwnerAccount(v string) *DeleteInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteInstancesRequest) SetResourceOwnerId(v int64) *DeleteInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteInstancesRequest) SetTerminateSubscription(v bool) *DeleteInstancesRequest {
	s.TerminateSubscription = &v
	return s
}

func (s *DeleteInstancesRequest) Validate() error {
	return dara.Validate(s)
}
