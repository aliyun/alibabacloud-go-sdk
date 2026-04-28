// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityToBenefitPkgMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityId(v string) *ListIdentityToBenefitPkgMappingRequest
	GetIdentityId() *string
	SetIdentityType(v string) *ListIdentityToBenefitPkgMappingRequest
	GetIdentityType() *string
	SetIncludeExpired(v bool) *ListIdentityToBenefitPkgMappingRequest
	GetIncludeExpired() *bool
}

type ListIdentityToBenefitPkgMappingRequest struct {
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
	// Specifies whether to return the benefit packages that expire. Default value: false.
	//
	// example:
	//
	// false
	IncludeExpired *bool `json:"include_expired,omitempty" xml:"include_expired,omitempty"`
}

func (s ListIdentityToBenefitPkgMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityToBenefitPkgMappingRequest) GoString() string {
	return s.String()
}

func (s *ListIdentityToBenefitPkgMappingRequest) GetIdentityId() *string {
	return s.IdentityId
}

func (s *ListIdentityToBenefitPkgMappingRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *ListIdentityToBenefitPkgMappingRequest) GetIncludeExpired() *bool {
	return s.IncludeExpired
}

func (s *ListIdentityToBenefitPkgMappingRequest) SetIdentityId(v string) *ListIdentityToBenefitPkgMappingRequest {
	s.IdentityId = &v
	return s
}

func (s *ListIdentityToBenefitPkgMappingRequest) SetIdentityType(v string) *ListIdentityToBenefitPkgMappingRequest {
	s.IdentityType = &v
	return s
}

func (s *ListIdentityToBenefitPkgMappingRequest) SetIncludeExpired(v bool) *ListIdentityToBenefitPkgMappingRequest {
	s.IncludeExpired = &v
	return s
}

func (s *ListIdentityToBenefitPkgMappingRequest) Validate() error {
	return dara.Validate(s)
}
