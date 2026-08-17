// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPoolName(v string) *CreatePoolRequest
	GetPoolName() *string
	SetPriority(v int32) *CreatePoolRequest
	GetPriority() *int32
	SetResourceLimits(v *CreatePoolRequestResourceLimits) *CreatePoolRequest
	GetResourceLimits() *CreatePoolRequestResourceLimits
	SetSchedulingPolicyId(v string) *CreatePoolRequest
	GetSchedulingPolicyId() *string
}

type CreatePoolRequest struct {
	// The resource pool name.
	//
	// - The name can be up to 15 characters in length.
	//
	// - The name can contain digits, uppercase letters, lowercase letters, underscores (_), and periods (.).
	//
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The priority of the resource pool.
	//
	// - Valid values: 1 to 99. Default value: 1, which indicates the lowest priority.
	//
	// - Jobs submitted to a resource pool with a higher priority value are scheduled before pending jobs in a resource pool with a lower priority value. The priority of a resource pool takes precedence over the priority of a job.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resource quota limits for concurrent usage allowed for a user within a resource pool.
	ResourceLimits *CreatePoolRequestResourceLimits `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty" type:"Struct"`
	// The scheduling policy.
	//
	// example:
	//
	// policy-xxx
	SchedulingPolicyId *string `json:"SchedulingPolicyId,omitempty" xml:"SchedulingPolicyId,omitempty"`
}

func (s CreatePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePoolRequest) GoString() string {
	return s.String()
}

func (s *CreatePoolRequest) GetPoolName() *string {
	return s.PoolName
}

func (s *CreatePoolRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreatePoolRequest) GetResourceLimits() *CreatePoolRequestResourceLimits {
	return s.ResourceLimits
}

func (s *CreatePoolRequest) GetSchedulingPolicyId() *string {
	return s.SchedulingPolicyId
}

func (s *CreatePoolRequest) SetPoolName(v string) *CreatePoolRequest {
	s.PoolName = &v
	return s
}

func (s *CreatePoolRequest) SetPriority(v int32) *CreatePoolRequest {
	s.Priority = &v
	return s
}

func (s *CreatePoolRequest) SetResourceLimits(v *CreatePoolRequestResourceLimits) *CreatePoolRequest {
	s.ResourceLimits = v
	return s
}

func (s *CreatePoolRequest) SetSchedulingPolicyId(v string) *CreatePoolRequest {
	s.SchedulingPolicyId = &v
	return s
}

func (s *CreatePoolRequest) Validate() error {
	if s.ResourceLimits != nil {
		if err := s.ResourceLimits.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePoolRequestResourceLimits struct {
	// The maximum number of executor nodes that a user can concurrently run within a resource pool.
	//
	// example:
	//
	// 100
	MaxExecutorNum *int32 `json:"MaxExecutorNum,omitempty" xml:"MaxExecutorNum,omitempty"`
}

func (s CreatePoolRequestResourceLimits) String() string {
	return dara.Prettify(s)
}

func (s CreatePoolRequestResourceLimits) GoString() string {
	return s.String()
}

func (s *CreatePoolRequestResourceLimits) GetMaxExecutorNum() *int32 {
	return s.MaxExecutorNum
}

func (s *CreatePoolRequestResourceLimits) SetMaxExecutorNum(v int32) *CreatePoolRequestResourceLimits {
	s.MaxExecutorNum = &v
	return s
}

func (s *CreatePoolRequestResourceLimits) Validate() error {
	return dara.Validate(s)
}
