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
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 8fa1953f-4a84-46d8-b80c-8ce9cf684fb3
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number to be modified
	//
	// This parameter is required.
	//
	// example:
	//
	// 10088xxx
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// Number ID
	//
	// This parameter is required.
	//
	// example:
	//
	// fa0e21e9-caab-4629-9121-1e341243d599
	OutboundCallNumberId *string `json:"OutboundCallNumberId,omitempty" xml:"OutboundCallNumberId,omitempty"`
	// Count of Rate Limiting rules
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	RateLimitCount *int32 `json:"RateLimitCount,omitempty" xml:"RateLimitCount,omitempty"`
	// Time Range for Rate Limiting, in seconds
	//
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
