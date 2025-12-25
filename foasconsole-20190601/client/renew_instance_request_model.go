// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenewInstanceRequest(v *RenewInstanceRequestRenewInstanceRequest) *RenewInstanceRequest
	GetRenewInstanceRequest() *RenewInstanceRequestRenewInstanceRequest
}

type RenewInstanceRequest struct {
	// This parameter is required.
	RenewInstanceRequest *RenewInstanceRequestRenewInstanceRequest `json:"RenewInstanceRequest,omitempty" xml:"RenewInstanceRequest,omitempty" type:"Struct"`
}

func (s RenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) GetRenewInstanceRequest() *RenewInstanceRequestRenewInstanceRequest {
	return s.RenewInstanceRequest
}

func (s *RenewInstanceRequest) SetRenewInstanceRequest(v *RenewInstanceRequestRenewInstanceRequest) *RenewInstanceRequest {
	s.RenewInstanceRequest = v
	return s
}

func (s *RenewInstanceRequest) Validate() error {
	if s.RenewInstanceRequest != nil {
		if err := s.RenewInstanceRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RenewInstanceRequestRenewInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sc_flinkserverless_public_cn-7e22ae5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s RenewInstanceRequestRenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequestRenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequestRenewInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *RenewInstanceRequestRenewInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewInstanceRequestRenewInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *RenewInstanceRequestRenewInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *RenewInstanceRequestRenewInstanceRequest) SetDuration(v int32) *RenewInstanceRequestRenewInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewInstanceRequestRenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequestRenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequestRenewInstanceRequest) SetPricingCycle(v string) *RenewInstanceRequestRenewInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewInstanceRequestRenewInstanceRequest) SetRegion(v string) *RenewInstanceRequestRenewInstanceRequest {
	s.Region = &v
	return s
}

func (s *RenewInstanceRequestRenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
