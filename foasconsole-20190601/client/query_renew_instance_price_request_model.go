// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRenewInstancePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenewInstanceRequest(v *QueryRenewInstancePriceRequestRenewInstanceRequest) *QueryRenewInstancePriceRequest
	GetRenewInstanceRequest() *QueryRenewInstancePriceRequestRenewInstanceRequest
}

type QueryRenewInstancePriceRequest struct {
	// This parameter is required.
	RenewInstanceRequest *QueryRenewInstancePriceRequestRenewInstanceRequest `json:"RenewInstanceRequest,omitempty" xml:"RenewInstanceRequest,omitempty" type:"Struct"`
}

func (s QueryRenewInstancePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceRequest) GetRenewInstanceRequest() *QueryRenewInstancePriceRequestRenewInstanceRequest {
	return s.RenewInstanceRequest
}

func (s *QueryRenewInstancePriceRequest) SetRenewInstanceRequest(v *QueryRenewInstancePriceRequestRenewInstanceRequest) *QueryRenewInstancePriceRequest {
	s.RenewInstanceRequest = v
	return s
}

func (s *QueryRenewInstancePriceRequest) Validate() error {
	if s.RenewInstanceRequest != nil {
		if err := s.RenewInstanceRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRenewInstancePriceRequestRenewInstanceRequest struct {
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
	// sc_flinkserverless_public_cn-******
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

func (s QueryRenewInstancePriceRequestRenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewInstancePriceRequestRenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) SetDuration(v int32) *QueryRenewInstancePriceRequestRenewInstanceRequest {
	s.Duration = &v
	return s
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) SetInstanceId(v string) *QueryRenewInstancePriceRequestRenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) SetPricingCycle(v string) *QueryRenewInstancePriceRequestRenewInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) SetRegion(v string) *QueryRenewInstancePriceRequestRenewInstanceRequest {
	s.Region = &v
	return s
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
