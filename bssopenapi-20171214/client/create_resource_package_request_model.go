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
	// The ID of the resource plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 2020-03-03T12:00:00Z
	EffectiveDate *string `json:"EffectiveDate,omitempty" xml:"EffectiveDate,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the order.
	//
	// This parameter is required.
	//
	// example:
	//
	// FPT_ossbag_absolute_Storage_sh
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// Indicates whether the request is successful.
	//
	// This parameter is required.
	//
	// example:
	//
	// ossbag
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40
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
