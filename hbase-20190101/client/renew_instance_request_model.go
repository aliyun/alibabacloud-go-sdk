// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RenewInstanceRequest
	GetClusterId() *string
	SetDuration(v int32) *RenewInstanceRequest
	GetDuration() *int32
	SetPricingCycle(v string) *RenewInstanceRequest
	GetPricingCycle() *string
}

type RenewInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1u0639js2h7****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RenewInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *RenewInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *RenewInstanceRequest) SetClusterId(v string) *RenewInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *RenewInstanceRequest) SetDuration(v int32) *RenewInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewInstanceRequest) SetPricingCycle(v string) *RenewInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
