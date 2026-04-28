// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityToBenefitPkgMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *UpdateIdentityToBenefitPkgMappingRequest
	GetAmount() *int64
	SetBenefitPkgId(v string) *UpdateIdentityToBenefitPkgMappingRequest
	GetBenefitPkgId() *string
	SetExpireTime(v int64) *UpdateIdentityToBenefitPkgMappingRequest
	GetExpireTime() *int64
	SetIdentityId(v string) *UpdateIdentityToBenefitPkgMappingRequest
	GetIdentityId() *string
	SetIdentityType(v string) *UpdateIdentityToBenefitPkgMappingRequest
	GetIdentityType() *string
}

type UpdateIdentityToBenefitPkgMappingRequest struct {
	// The number of benefit packages.
	//
	// This parameter specifies the number of benefit packages of the resource type. Default value: 1.
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
	// The expiration time of the benefit package. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// By default, the benefit package never expires.
	//
	// example:
	//
	// 1633167071000
	ExpireTime *int64 `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// The unique identifier of the entity.
	//
	// If you call this operation to manage the benefits of a user, set this parameter to the ID of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// user123
	IdentityId *string `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	// The type of the entity. If you call this operation to manage the benefits of a user, set this parameter to user.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	IdentityType *string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
}

func (s UpdateIdentityToBenefitPkgMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityToBenefitPkgMappingRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) GetAmount() *int64 {
	return s.Amount
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) GetBenefitPkgId() *string {
	return s.BenefitPkgId
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) GetIdentityId() *string {
	return s.IdentityId
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) SetAmount(v int64) *UpdateIdentityToBenefitPkgMappingRequest {
	s.Amount = &v
	return s
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) SetBenefitPkgId(v string) *UpdateIdentityToBenefitPkgMappingRequest {
	s.BenefitPkgId = &v
	return s
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) SetExpireTime(v int64) *UpdateIdentityToBenefitPkgMappingRequest {
	s.ExpireTime = &v
	return s
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) SetIdentityId(v string) *UpdateIdentityToBenefitPkgMappingRequest {
	s.IdentityId = &v
	return s
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) SetIdentityType(v string) *UpdateIdentityToBenefitPkgMappingRequest {
	s.IdentityType = &v
	return s
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) Validate() error {
	return dara.Validate(s)
}
