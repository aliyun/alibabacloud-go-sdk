// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDDoSBillingMode(v string) *DescribeDDoSPriceRequest
	GetDDoSBillingMode() *string
	SetDDoSBurstableDomesticProtection(v string) *DescribeDDoSPriceRequest
	GetDDoSBurstableDomesticProtection() *string
	SetDDoSBurstableOverseasProtection(v string) *DescribeDDoSPriceRequest
	GetDDoSBurstableOverseasProtection() *string
}

type DescribeDDoSPriceRequest struct {
	// The billing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// CleanTraffic
	DDoSBillingMode *string `json:"DDoSBillingMode,omitempty" xml:"DDoSBillingMode,omitempty"`
	// The instance specifications for the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn_300
	DDoSBurstableDomesticProtection *string `json:"DDoSBurstableDomesticProtection,omitempty" xml:"DDoSBurstableDomesticProtection,omitempty"`
	// The instance specifications for regions outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// overseas_300
	DDoSBurstableOverseasProtection *string `json:"DDoSBurstableOverseasProtection,omitempty" xml:"DDoSBurstableOverseasProtection,omitempty"`
}

func (s DescribeDDoSPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSPriceRequest) GetDDoSBillingMode() *string {
	return s.DDoSBillingMode
}

func (s *DescribeDDoSPriceRequest) GetDDoSBurstableDomesticProtection() *string {
	return s.DDoSBurstableDomesticProtection
}

func (s *DescribeDDoSPriceRequest) GetDDoSBurstableOverseasProtection() *string {
	return s.DDoSBurstableOverseasProtection
}

func (s *DescribeDDoSPriceRequest) SetDDoSBillingMode(v string) *DescribeDDoSPriceRequest {
	s.DDoSBillingMode = &v
	return s
}

func (s *DescribeDDoSPriceRequest) SetDDoSBurstableDomesticProtection(v string) *DescribeDDoSPriceRequest {
	s.DDoSBurstableDomesticProtection = &v
	return s
}

func (s *DescribeDDoSPriceRequest) SetDDoSBurstableOverseasProtection(v string) *DescribeDDoSPriceRequest {
	s.DDoSBurstableOverseasProtection = &v
	return s
}

func (s *DescribeDDoSPriceRequest) Validate() error {
	return dara.Validate(s)
}
