// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityToBenefitPkgMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBenefitPkgId(v string) *GetIdentityToBenefitPkgMappingRequest
	GetBenefitPkgId() *string
	SetIdentityId(v string) *GetIdentityToBenefitPkgMappingRequest
	GetIdentityId() *string
	SetIdentityType(v string) *GetIdentityToBenefitPkgMappingRequest
	GetIdentityType() *string
}

type GetIdentityToBenefitPkgMappingRequest struct {
	// The unique identifier of the benefit package.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40cb7794c9294
	BenefitPkgId *string `json:"benefit_pkg_id,omitempty" xml:"benefit_pkg_id,omitempty"`
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
	// The type of the entity. If you want to manage the benefits of a user, set this parameter to user.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	IdentityType *string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
}

func (s GetIdentityToBenefitPkgMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityToBenefitPkgMappingRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityToBenefitPkgMappingRequest) GetBenefitPkgId() *string {
	return s.BenefitPkgId
}

func (s *GetIdentityToBenefitPkgMappingRequest) GetIdentityId() *string {
	return s.IdentityId
}

func (s *GetIdentityToBenefitPkgMappingRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *GetIdentityToBenefitPkgMappingRequest) SetBenefitPkgId(v string) *GetIdentityToBenefitPkgMappingRequest {
	s.BenefitPkgId = &v
	return s
}

func (s *GetIdentityToBenefitPkgMappingRequest) SetIdentityId(v string) *GetIdentityToBenefitPkgMappingRequest {
	s.IdentityId = &v
	return s
}

func (s *GetIdentityToBenefitPkgMappingRequest) SetIdentityType(v string) *GetIdentityToBenefitPkgMappingRequest {
	s.IdentityType = &v
	return s
}

func (s *GetIdentityToBenefitPkgMappingRequest) Validate() error {
	return dara.Validate(s)
}
