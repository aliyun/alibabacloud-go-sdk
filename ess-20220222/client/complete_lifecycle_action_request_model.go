// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteLifecycleActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CompleteLifecycleActionRequest
	GetClientToken() *string
	SetLifecycleActionResult(v string) *CompleteLifecycleActionRequest
	GetLifecycleActionResult() *string
	SetLifecycleActionToken(v string) *CompleteLifecycleActionRequest
	GetLifecycleActionToken() *string
	SetLifecycleHookId(v string) *CompleteLifecycleActionRequest
	GetLifecycleHookId() *string
	SetOwnerAccount(v string) *CompleteLifecycleActionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CompleteLifecycleActionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CompleteLifecycleActionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CompleteLifecycleActionRequest
	GetResourceOwnerAccount() *string
}

type CompleteLifecycleActionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The action that you want Auto Scaling to perform after the lifecycle hook times out. Valid values:
	//
	// 	- CONTINUE: Auto Scaling continues to respond to a scale-in or scale-out request.
	//
	// 	- ABANDON: Auto Scaling releases Elastic Compute Service (ECS) instances that are created during scale-out activities or removes ECS instances from the scaling group during scale-in activities.
	//
	// 	- ROLLBACK: For scale-in activities, Auto Scaling rejects the requests to release ECS instances but rolls back ECS instances. For scale-out activities, the ROLLBACK setting has the same effect as the ABANDON setting.
	//
	// If you do not specify this parameter, Auto Scaling performs the action that is specified by the `DefaultResult` parameter after the lifecycle hook times out.
	//
	// If multiple lifecycle hooks exist in a scaling group and the lifecycle hooks are triggered at the same time, the following rules apply:
	//
	// 	- For scale-in activities, when lifecycle hooks whose LifecycleActionResult parameter is set to ABANDON or ROLLBACK time out, other lifecycle hooks time out ahead of schedule.
	//
	// 	- For scale-in and scale-out activities, if you set the LifecycleActionResult parameter for all lifecycle hooks to CONTINUE, Auto Scaling performs the next action only after the last lifecycle hook times out. The action that Auto Scaling performs varies based on the value that you specify for the LifecycleActionResult parameter of the lifecycle hook that last times out.
	//
	// example:
	//
	// CONTINUE
	LifecycleActionResult *string `json:"LifecycleActionResult,omitempty" xml:"LifecycleActionResult,omitempty"`
	// The token of the lifecycle action. You can obtain the token from the Simple Message Queue (SMQ, formerly MNS) queue or topic that is specified for the lifecycle hook.
	//
	// This parameter is required.
	//
	// example:
	//
	// aaaa-bbbbb-cccc-ddddd
	LifecycleActionToken *string `json:"LifecycleActionToken,omitempty" xml:"LifecycleActionToken,omitempty"`
	// The ID of the lifecycle hook.
	//
	// This parameter is required.
	//
	// example:
	//
	// ash-bp14g3ee6bt3sc98****
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s CompleteLifecycleActionRequest) String() string {
	return dara.Prettify(s)
}

func (s CompleteLifecycleActionRequest) GoString() string {
	return s.String()
}

func (s *CompleteLifecycleActionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CompleteLifecycleActionRequest) GetLifecycleActionResult() *string {
	return s.LifecycleActionResult
}

func (s *CompleteLifecycleActionRequest) GetLifecycleActionToken() *string {
	return s.LifecycleActionToken
}

func (s *CompleteLifecycleActionRequest) GetLifecycleHookId() *string {
	return s.LifecycleHookId
}

func (s *CompleteLifecycleActionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CompleteLifecycleActionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CompleteLifecycleActionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CompleteLifecycleActionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CompleteLifecycleActionRequest) SetClientToken(v string) *CompleteLifecycleActionRequest {
	s.ClientToken = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetLifecycleActionResult(v string) *CompleteLifecycleActionRequest {
	s.LifecycleActionResult = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetLifecycleActionToken(v string) *CompleteLifecycleActionRequest {
	s.LifecycleActionToken = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetLifecycleHookId(v string) *CompleteLifecycleActionRequest {
	s.LifecycleHookId = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetOwnerAccount(v string) *CompleteLifecycleActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetOwnerId(v int64) *CompleteLifecycleActionRequest {
	s.OwnerId = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetRegionId(v string) *CompleteLifecycleActionRequest {
	s.RegionId = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetResourceOwnerAccount(v string) *CompleteLifecycleActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CompleteLifecycleActionRequest) Validate() error {
	return dara.Validate(s)
}
