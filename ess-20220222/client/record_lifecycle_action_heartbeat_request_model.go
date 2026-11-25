// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordLifecycleActionHeartbeatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *RecordLifecycleActionHeartbeatRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RecordLifecycleActionHeartbeatRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RecordLifecycleActionHeartbeatRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RecordLifecycleActionHeartbeatRequest
	GetResourceOwnerAccount() *string
	SetHeartbeatTimeout(v int32) *RecordLifecycleActionHeartbeatRequest
	GetHeartbeatTimeout() *int32
	SetLifecycleActionToken(v string) *RecordLifecycleActionHeartbeatRequest
	GetLifecycleActionToken() *string
	SetLifecycleHookId(v string) *RecordLifecycleActionHeartbeatRequest
	GetLifecycleHookId() *string
}

type RecordLifecycleActionHeartbeatRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The wait time configured in the lifecycle hook for the scaling activity. If the wait state times out, the default execution policy is executed. Valid values: 30 to 21600. Unit: seconds.
	//
	// After you create a lifecycle hook, you can call this operation to extend the wait time of the ECS instance, or call the [CompleteLifecycleAction](https://help.aliyun.com/document_detail/459335.html) operation to end the wait state of the scaling activity in advance.
	//
	// Default value: 600.
	//
	// example:
	//
	// 600
	HeartbeatTimeout *int32 `json:"heartbeatTimeout,omitempty" xml:"heartbeatTimeout,omitempty"`
	// The wait state token of the scaling activity. When the instance is suspended, you can obtain this token from the details page of the message received from the MNS queue specified by the lifecycle hook.
	//
	// You can also call the [DescribeLifecycleActions](https://help.aliyun.com/document_detail/459333.html) operation to obtain this token.
	//
	// Alternatively, you can obtain this value from the MNS topic specified by the lifecycle hook.
	//
	// This parameter is required.
	//
	// example:
	//
	// F324B880-900E-4968-85DD-81691113****
	LifecycleActionToken *string `json:"lifecycleActionToken,omitempty" xml:"lifecycleActionToken,omitempty"`
	// The ID of the lifecycle hook.
	//
	// This parameter is required.
	//
	// example:
	//
	// ash-bp1fxuqyi98w0aib****
	LifecycleHookId *string `json:"lifecycleHookId,omitempty" xml:"lifecycleHookId,omitempty"`
}

func (s RecordLifecycleActionHeartbeatRequest) String() string {
	return dara.Prettify(s)
}

func (s RecordLifecycleActionHeartbeatRequest) GoString() string {
	return s.String()
}

func (s *RecordLifecycleActionHeartbeatRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RecordLifecycleActionHeartbeatRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RecordLifecycleActionHeartbeatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RecordLifecycleActionHeartbeatRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RecordLifecycleActionHeartbeatRequest) GetHeartbeatTimeout() *int32 {
	return s.HeartbeatTimeout
}

func (s *RecordLifecycleActionHeartbeatRequest) GetLifecycleActionToken() *string {
	return s.LifecycleActionToken
}

func (s *RecordLifecycleActionHeartbeatRequest) GetLifecycleHookId() *string {
	return s.LifecycleHookId
}

func (s *RecordLifecycleActionHeartbeatRequest) SetOwnerAccount(v string) *RecordLifecycleActionHeartbeatRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetOwnerId(v int64) *RecordLifecycleActionHeartbeatRequest {
	s.OwnerId = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetRegionId(v string) *RecordLifecycleActionHeartbeatRequest {
	s.RegionId = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetResourceOwnerAccount(v string) *RecordLifecycleActionHeartbeatRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetHeartbeatTimeout(v int32) *RecordLifecycleActionHeartbeatRequest {
	s.HeartbeatTimeout = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetLifecycleActionToken(v string) *RecordLifecycleActionHeartbeatRequest {
	s.LifecycleActionToken = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetLifecycleHookId(v string) *RecordLifecycleActionHeartbeatRequest {
	s.LifecycleHookId = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) Validate() error {
	return dara.Validate(s)
}
