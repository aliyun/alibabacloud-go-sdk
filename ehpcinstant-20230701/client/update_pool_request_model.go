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
	return dara.Validate(s)
}

type UpdatePoolRequestResourceLimits struct {
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
