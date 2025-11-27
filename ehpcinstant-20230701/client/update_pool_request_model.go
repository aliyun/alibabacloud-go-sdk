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
}

type UpdatePoolRequest struct {
	// The name of the resource pool.
	//
	// 	- The value can be up to 15 characters in length.
	//
	// 	- It can contain digits, uppercase letters, lowercase letters, underscores (_), and dots (.).
	//
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The priority of the resource pool.
	//
	// 	- You can set a priority in the range of 1 to 99. The default value is 1, which is the lowest priority.
	//
	// 	- Jobs submitted to a resource pool with a higher priority level value will be scheduled before pending jobs in a resource pool with a lower priority level value, and the priority level of the resource pool takes precedence over the priority of the job.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The quota of resources that users are allowed to concurrently use in a resource pool.
	ResourceLimits *UpdatePoolRequestResourceLimits `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty" type:"Struct"`
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

func (s *UpdatePoolRequest) Validate() error {
	if s.ResourceLimits != nil {
		if err := s.ResourceLimits.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePoolRequestResourceLimits struct {
	// The maximum number of concurrent execution nodes in a resource pool.
	//
	// example:
	//
	// 2000
	MaxExectorNum *int32 `json:"MaxExectorNum,omitempty" xml:"MaxExectorNum,omitempty"`
}

func (s UpdatePoolRequestResourceLimits) String() string {
	return dara.Prettify(s)
}

func (s UpdatePoolRequestResourceLimits) GoString() string {
	return s.String()
}

func (s *UpdatePoolRequestResourceLimits) GetMaxExectorNum() *int32 {
	return s.MaxExectorNum
}

func (s *UpdatePoolRequestResourceLimits) SetMaxExectorNum(v int32) *UpdatePoolRequestResourceLimits {
	s.MaxExectorNum = &v
	return s
}

func (s *UpdatePoolRequestResourceLimits) Validate() error {
	return dara.Validate(s)
}
