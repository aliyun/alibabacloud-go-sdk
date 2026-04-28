// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityToBenefitPkgMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*IdentityToBenefitPkgMapping) *ListIdentityToBenefitPkgMappingResponseBody
	GetItems() []*IdentityToBenefitPkgMapping
}

type ListIdentityToBenefitPkgMappingResponseBody struct {
	// The information about the benefit packages that are associated with an entity.
	Items []*IdentityToBenefitPkgMapping `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListIdentityToBenefitPkgMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityToBenefitPkgMappingResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdentityToBenefitPkgMappingResponseBody) GetItems() []*IdentityToBenefitPkgMapping {
	return s.Items
}

func (s *ListIdentityToBenefitPkgMappingResponseBody) SetItems(v []*IdentityToBenefitPkgMapping) *ListIdentityToBenefitPkgMappingResponseBody {
	s.Items = v
	return s
}

func (s *ListIdentityToBenefitPkgMappingResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
