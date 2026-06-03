// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOutboundCallNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyOutboundCallNumberRequest
	GetInstanceId() *string
	SetNumber(v string) *ModifyOutboundCallNumberRequest
	GetNumber() *string
	SetOutboundCallNumberId(v string) *ModifyOutboundCallNumberRequest
	GetOutboundCallNumberId() *string
	SetRateLimitCount(v int32) *ModifyOutboundCallNumberRequest
	GetRateLimitCount() *int32
	SetRateLimitPeriod(v int32) *ModifyOutboundCallNumberRequest
	GetRateLimitPeriod() *int32
}

type ModifyOutboundCallNumberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 52e80b02-0126-4556-a1e6-ef5b3747ed53/a9a3ddc7-d7d7-48cd-82b5-b31bb5510e71_2a66f8ad-dfbb-4980-9b84-439171295a11.xlsx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10088
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fa0e21e9-caab-4629-9121-1e341243d599
	OutboundCallNumberId *string `json:"OutboundCallNumberId,omitempty" xml:"OutboundCallNumberId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	RateLimitCount *int32 `json:"RateLimitCount,omitempty" xml:"RateLimitCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	RateLimitPeriod *int32 `json:"RateLimitPeriod,omitempty" xml:"RateLimitPeriod,omitempty"`
}

func (s ModifyOutboundCallNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOutboundCallNumberRequest) GoString() string {
	return s.String()
}

func (s *ModifyOutboundCallNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyOutboundCallNumberRequest) GetNumber() *string {
	return s.Number
}

func (s *ModifyOutboundCallNumberRequest) GetOutboundCallNumberId() *string {
	return s.OutboundCallNumberId
}

func (s *ModifyOutboundCallNumberRequest) GetRateLimitCount() *int32 {
	return s.RateLimitCount
}

func (s *ModifyOutboundCallNumberRequest) GetRateLimitPeriod() *int32 {
	return s.RateLimitPeriod
}

func (s *ModifyOutboundCallNumberRequest) SetInstanceId(v string) *ModifyOutboundCallNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyOutboundCallNumberRequest) SetNumber(v string) *ModifyOutboundCallNumberRequest {
	s.Number = &v
	return s
}

func (s *ModifyOutboundCallNumberRequest) SetOutboundCallNumberId(v string) *ModifyOutboundCallNumberRequest {
	s.OutboundCallNumberId = &v
	return s
}

func (s *ModifyOutboundCallNumberRequest) SetRateLimitCount(v int32) *ModifyOutboundCallNumberRequest {
	s.RateLimitCount = &v
	return s
}

func (s *ModifyOutboundCallNumberRequest) SetRateLimitPeriod(v int32) *ModifyOutboundCallNumberRequest {
	s.RateLimitPeriod = &v
	return s
}

func (s *ModifyOutboundCallNumberRequest) Validate() error {
	return dara.Validate(s)
}
