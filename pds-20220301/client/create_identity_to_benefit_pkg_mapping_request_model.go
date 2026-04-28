// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityToBenefitPkgMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *CreateIdentityToBenefitPkgMappingRequest
	GetAmount() *int64
	SetBenefitPkgId(v string) *CreateIdentityToBenefitPkgMappingRequest
	GetBenefitPkgId() *string
	SetExpireTime(v int64) *CreateIdentityToBenefitPkgMappingRequest
	GetExpireTime() *int64
	SetIdentityId(v string) *CreateIdentityToBenefitPkgMappingRequest
	GetIdentityId() *string
	SetIdentityType(v string) *CreateIdentityToBenefitPkgMappingRequest
	GetIdentityType() *string
}

type CreateIdentityToBenefitPkgMappingRequest struct {
	// The number of benefit packages.
	//
	// This parameter takes effect only for the benefit packages of the resource type. Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The unique identifier of the benefit package.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40cb7794c9294
	BenefitPkgId *string `json:"benefit_pkg_id,omitempty" xml:"benefit_pkg_id,omitempty"`
	// The time when the benefit package expires. Set the value to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// By default, the benefit package never expires.
	//
	// example:
	//
	// 1633167071000
	ExpireTime *int64 `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// The unique identifier of the entity.
	//
	// If you want to manage the benefits of a user, set this parameter to a user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user123
	IdentityId *string `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	// The type of the entity.
	//
	// If you want to manage the benefits of a user, set this parameter to user.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	IdentityType *string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
}

func (s CreateIdentityToBenefitPkgMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityToBenefitPkgMappingRequest) GoString() string {
	return s.String()
}

func (s *CreateIdentityToBenefitPkgMappingRequest) GetAmount() *int64 {
	return s.Amount
}

func (s *CreateIdentityToBenefitPkgMappingRequest) GetBenefitPkgId() *string {
	return s.BenefitPkgId
}

func (s *CreateIdentityToBenefitPkgMappingRequest) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *CreateIdentityToBenefitPkgMappingRequest) GetIdentityId() *string {
	return s.IdentityId
}

func (s *CreateIdentityToBenefitPkgMappingRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *CreateIdentityToBenefitPkgMappingRequest) SetAmount(v int64) *CreateIdentityToBenefitPkgMappingRequest {
	s.Amount = &v
	return s
}

func (s *CreateIdentityToBenefitPkgMappingRequest) SetBenefitPkgId(v string) *CreateIdentityToBenefitPkgMappingRequest {
	s.BenefitPkgId = &v
	return s
}

func (s *CreateIdentityToBenefitPkgMappingRequest) SetExpireTime(v int64) *CreateIdentityToBenefitPkgMappingRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateIdentityToBenefitPkgMappingRequest) SetIdentityId(v string) *CreateIdentityToBenefitPkgMappingRequest {
	s.IdentityId = &v
	return s
}

func (s *CreateIdentityToBenefitPkgMappingRequest) SetIdentityType(v string) *CreateIdentityToBenefitPkgMappingRequest {
	s.IdentityType = &v
	return s
}

func (s *CreateIdentityToBenefitPkgMappingRequest) Validate() error {
	return dara.Validate(s)
}
