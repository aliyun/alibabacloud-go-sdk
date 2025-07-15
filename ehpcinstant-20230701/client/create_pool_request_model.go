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
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// example:
	//
	// 1
	Priority       *int32                           `json:"Priority,omitempty" xml:"Priority,omitempty"`
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
	return dara.Validate(s)
}

type CreatePoolRequestResourceLimits struct {
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
