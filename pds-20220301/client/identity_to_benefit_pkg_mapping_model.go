// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdentityToBenefitPkgMapping interface {
	dara.Model
	String() string
	GoString() string
	SetBenefitPkgComputationRule(v string) *IdentityToBenefitPkgMapping
	GetBenefitPkgComputationRule() *string
	SetBenefitPkgId(v string) *IdentityToBenefitPkgMapping
	GetBenefitPkgId() *string
	SetBenefitPkgName(v string) *IdentityToBenefitPkgMapping
	GetBenefitPkgName() *string
	SetBenefitPkgOwnerId(v string) *IdentityToBenefitPkgMapping
	GetBenefitPkgOwnerId() *string
	SetBenefitPkgPriority(v int64) *IdentityToBenefitPkgMapping
	GetBenefitPkgPriority() *int64
	SetBenefitPkgType(v string) *IdentityToBenefitPkgMapping
	GetBenefitPkgType() *string
	SetCreatedAt(v string) *IdentityToBenefitPkgMapping
	GetCreatedAt() *string
	SetDeliveryInfoList(v []*BenefitPkgDeliveryInfo) *IdentityToBenefitPkgMapping
	GetDeliveryInfoList() []*BenefitPkgDeliveryInfo
	SetIdentityId(v string) *IdentityToBenefitPkgMapping
	GetIdentityId() *string
	SetIdentityType(v string) *IdentityToBenefitPkgMapping
	GetIdentityType() *string
	SetUpdatedAt(v string) *IdentityToBenefitPkgMapping
	GetUpdatedAt() *string
}

type IdentityToBenefitPkgMapping struct {
	// Calculation rules of equity in the benefit package.
	//
	// The user identity benefit package. The return value is empty. Only the quota of the effective equity is calculated based on the priority.
	//
	// The user resource benefit package, which can return null or non-null values. If the return value is not empty, this benefit package is used to append the quota of existing benefits in other benefit packages, which is limited to quota-type benefits. For example, if a user identity benefit package already contains 10 GB of user storage capacity, you can define one or more user resource benefit packages to add additional storage capacity for some users.
	//
	// The following append calculation rules are supported:
	//
	// sum: Multiple benefit packages have the same equity and are accumulated.
	//
	// max: If multiple benefit packages have the same equity, the max value is used.
	//
	// min: If multiple benefit packages have the same equity, the value of min is used.
	//
	// example:
	//
	// sum
	BenefitPkgComputationRule *string `json:"benefit_pkg_computation_rule,omitempty" xml:"benefit_pkg_computation_rule,omitempty"`
	// The ID of the benefit package.
	//
	// example:
	//
	// 40cb7794c9294
	BenefitPkgId *string `json:"benefit_pkg_id,omitempty" xml:"benefit_pkg_id,omitempty"`
	// The name of the benefit package.
	BenefitPkgName *string `json:"benefit_pkg_name,omitempty" xml:"benefit_pkg_name,omitempty"`
	// The ID of the owner of the benefit package.
	//
	// example:
	//
	// bj1
	BenefitPkgOwnerId *string `json:"benefit_pkg_owner_id,omitempty" xml:"benefit_pkg_owner_id,omitempty"`
	// Priority of the benefit package.
	//
	// The priority returned for the user identity benefit package. A smaller number indicates a higher priority.
	//
	// example:
	//
	// 1
	BenefitPkgPriority *int64 `json:"benefit_pkg_priority,omitempty" xml:"benefit_pkg_priority,omitempty"`
	// The type of benefit package.
	//
	// Valid values:
	//
	// user_identity : user identity benefit package
	//
	// user_resource: user resource benefit package
	//
	// example:
	//
	// user_identity
	BenefitPkgType *string `json:"benefit_pkg_type,omitempty" xml:"benefit_pkg_type,omitempty"`
	// Creation time of the entity and benefit package association.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The information about the benefit packages that are associated with an entity.
	DeliveryInfoList []*BenefitPkgDeliveryInfo `json:"delivery_info_list,omitempty" xml:"delivery_info_list,omitempty" type:"Repeated"`
	// The ID of the entity.
	//
	// example:
	//
	// user123
	IdentityId *string `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	// The type of the entity.
	//
	// example:
	//
	// user
	IdentityType *string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
	// Update time associated with the entity and benefit package.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s IdentityToBenefitPkgMapping) String() string {
	return dara.Prettify(s)
}

func (s IdentityToBenefitPkgMapping) GoString() string {
	return s.String()
}

func (s *IdentityToBenefitPkgMapping) GetBenefitPkgComputationRule() *string {
	return s.BenefitPkgComputationRule
}

func (s *IdentityToBenefitPkgMapping) GetBenefitPkgId() *string {
	return s.BenefitPkgId
}

func (s *IdentityToBenefitPkgMapping) GetBenefitPkgName() *string {
	return s.BenefitPkgName
}

func (s *IdentityToBenefitPkgMapping) GetBenefitPkgOwnerId() *string {
	return s.BenefitPkgOwnerId
}

func (s *IdentityToBenefitPkgMapping) GetBenefitPkgPriority() *int64 {
	return s.BenefitPkgPriority
}

func (s *IdentityToBenefitPkgMapping) GetBenefitPkgType() *string {
	return s.BenefitPkgType
}

func (s *IdentityToBenefitPkgMapping) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *IdentityToBenefitPkgMapping) GetDeliveryInfoList() []*BenefitPkgDeliveryInfo {
	return s.DeliveryInfoList
}

func (s *IdentityToBenefitPkgMapping) GetIdentityId() *string {
	return s.IdentityId
}

func (s *IdentityToBenefitPkgMapping) GetIdentityType() *string {
	return s.IdentityType
}

func (s *IdentityToBenefitPkgMapping) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *IdentityToBenefitPkgMapping) SetBenefitPkgComputationRule(v string) *IdentityToBenefitPkgMapping {
	s.BenefitPkgComputationRule = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetBenefitPkgId(v string) *IdentityToBenefitPkgMapping {
	s.BenefitPkgId = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetBenefitPkgName(v string) *IdentityToBenefitPkgMapping {
	s.BenefitPkgName = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetBenefitPkgOwnerId(v string) *IdentityToBenefitPkgMapping {
	s.BenefitPkgOwnerId = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetBenefitPkgPriority(v int64) *IdentityToBenefitPkgMapping {
	s.BenefitPkgPriority = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetBenefitPkgType(v string) *IdentityToBenefitPkgMapping {
	s.BenefitPkgType = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetCreatedAt(v string) *IdentityToBenefitPkgMapping {
	s.CreatedAt = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetDeliveryInfoList(v []*BenefitPkgDeliveryInfo) *IdentityToBenefitPkgMapping {
	s.DeliveryInfoList = v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetIdentityId(v string) *IdentityToBenefitPkgMapping {
	s.IdentityId = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetIdentityType(v string) *IdentityToBenefitPkgMapping {
	s.IdentityType = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetUpdatedAt(v string) *IdentityToBenefitPkgMapping {
	s.UpdatedAt = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) Validate() error {
	if s.DeliveryInfoList != nil {
		for _, item := range s.DeliveryInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
