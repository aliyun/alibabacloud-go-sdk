// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPoolName(v string) *UpdatePoolRequest
	GetPoolName() *string
	SetPriority(v int32) *UpdatePoolRequest
	GetPriority() *int32
	SetResourceLimits(v *UpdatePoolRequestResourceLimits) *UpdatePoolRequest
	GetResourceLimits() *UpdatePoolRequestResourceLimits
	SetSchedulingPolicyId(v string) *UpdatePoolRequest
	GetSchedulingPolicyId() *string
}

type UpdatePoolRequest struct {
	// The name of the resource pool.
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
	// - Jobs submitted to a resource pool with a higher priority value are scheduled before pending jobs in resource pools with lower priority values. The priority of a resource pool takes precedence over the priority of a job.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resource quota limit for concurrent usage by a user within a resource pool.
	ResourceLimits *UpdatePoolRequestResourceLimits `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty" type:"Struct"`
	// The scheduling policy ID.
	//
	// example:
	//
	// policy-xxxx
	SchedulingPolicyId *string `json:"SchedulingPolicyId,omitempty" xml:"SchedulingPolicyId,omitempty"`
}

func (s UpdatePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePoolRequest) GoString() string {
	return s.String()
}

func (s *UpdatePoolRequest) GetPoolName() *string {
	return s.PoolName
}

func (s *UpdatePoolRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdatePoolRequest) GetResourceLimits() *UpdatePoolRequestResourceLimits {
	return s.ResourceLimits
}

func (s *UpdatePoolRequest) GetSchedulingPolicyId() *string {
	return s.SchedulingPolicyId
}

func (s *UpdatePoolRequest) SetPoolName(v string) *UpdatePoolRequest {
	s.PoolName = &v
	return s
}

func (s *UpdatePoolRequest) SetPriority(v int32) *UpdatePoolRequest {
	s.Priority = &v
	return s
}

func (s *UpdatePoolRequest) SetResourceLimits(v *UpdatePoolRequestResourceLimits) *UpdatePoolRequest {
	s.ResourceLimits = v
	return s
}

func (s *UpdatePoolRequest) SetSchedulingPolicyId(v string) *UpdatePoolRequest {
	s.SchedulingPolicyId = &v
	return s
}

func (s *UpdatePoolRequest) Validate() error {
	if s.ResourceLimits != nil {
		if err := s.ResourceLimits.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePoolRequestResourceLimits struct {
	// The maximum number of executor nodes that a user can concurrently run within a resource pool.
	//
	// example:
	//
	// 100
	MaxExecutorNum *int32 `json:"MaxExecutorNum,omitempty" xml:"MaxExecutorNum,omitempty"`
}

func (s UpdatePoolRequestResourceLimits) String() string {
	return dara.Prettify(s)
}

func (s UpdatePoolRequestResourceLimits) GoString() string {
	return s.String()
}

func (s *UpdatePoolRequestResourceLimits) GetMaxExecutorNum() *int32 {
	return s.MaxExecutorNum
}

func (s *UpdatePoolRequestResourceLimits) SetMaxExecutorNum(v int32) *UpdatePoolRequestResourceLimits {
	s.MaxExecutorNum = &v
	return s
}

func (s *UpdatePoolRequestResourceLimits) Validate() error {
	return dara.Validate(s)
}
