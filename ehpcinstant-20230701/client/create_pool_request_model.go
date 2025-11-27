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
}

type CreatePoolRequest struct {
	// The name of the resource pool.
	//
	// 	- The name can be up to 15 characters in length.
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
	ResourceLimits *CreatePoolRequestResourceLimits `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty" type:"Struct"`
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

func (s *CreatePoolRequest) Validate() error {
	if s.ResourceLimits != nil {
		if err := s.ResourceLimits.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePoolRequestResourceLimits struct {
	// The maximum number of concurrent execution nodes in a resource pool.
	//
	// example:
	//
	// 2000
	MaxExectorNum *int32 `json:"MaxExectorNum,omitempty" xml:"MaxExectorNum,omitempty"`
}

func (s CreatePoolRequestResourceLimits) String() string {
	return dara.Prettify(s)
}

func (s CreatePoolRequestResourceLimits) GoString() string {
	return s.String()
}

func (s *CreatePoolRequestResourceLimits) GetMaxExectorNum() *int32 {
	return s.MaxExectorNum
}

func (s *CreatePoolRequestResourceLimits) SetMaxExectorNum(v int32) *CreatePoolRequestResourceLimits {
	s.MaxExectorNum = &v
	return s
}

func (s *CreatePoolRequestResourceLimits) Validate() error {
	return dara.Validate(s)
}
