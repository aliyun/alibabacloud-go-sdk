// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourcePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *CreateResourcePackageRequest
	GetDuration() *int32
	SetEffectiveDate(v string) *CreateResourcePackageRequest
	GetEffectiveDate() *string
	SetOwnerId(v int64) *CreateResourcePackageRequest
	GetOwnerId() *int64
	SetPackageType(v string) *CreateResourcePackageRequest
	GetPackageType() *string
	SetPricingCycle(v string) *CreateResourcePackageRequest
	GetPricingCycle() *string
	SetProductCode(v string) *CreateResourcePackageRequest
	GetProductCode() *string
	SetSpecification(v string) *CreateResourcePackageRequest
	GetSpecification() *string
}

type CreateResourcePackageRequest struct {
	// This parameter is required.
	Duration      *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EffectiveDate *string `json:"EffectiveDate,omitempty" xml:"EffectiveDate,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PackageType  *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// This parameter is required.
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s CreateResourcePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *CreateResourcePackageRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateResourcePackageRequest) GetEffectiveDate() *string {
	return s.EffectiveDate
}

func (s *CreateResourcePackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateResourcePackageRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *CreateResourcePackageRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateResourcePackageRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateResourcePackageRequest) GetSpecification() *string {
	return s.Specification
}

func (s *CreateResourcePackageRequest) SetDuration(v int32) *CreateResourcePackageRequest {
	s.Duration = &v
	return s
}

func (s *CreateResourcePackageRequest) SetEffectiveDate(v string) *CreateResourcePackageRequest {
	s.EffectiveDate = &v
	return s
}

func (s *CreateResourcePackageRequest) SetOwnerId(v int64) *CreateResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateResourcePackageRequest) SetPackageType(v string) *CreateResourcePackageRequest {
	s.PackageType = &v
	return s
}

func (s *CreateResourcePackageRequest) SetPricingCycle(v string) *CreateResourcePackageRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateResourcePackageRequest) SetProductCode(v string) *CreateResourcePackageRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateResourcePackageRequest) SetSpecification(v string) *CreateResourcePackageRequest {
	s.Specification = &v
	return s
}

func (s *CreateResourcePackageRequest) Validate() error {
	return dara.Validate(s)
}
