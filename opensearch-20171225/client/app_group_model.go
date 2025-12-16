// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppGroup interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *AppGroup
	GetChargeType() *string
	SetDescription(v string) *AppGroup
	GetDescription() *string
	SetDomain(v string) *AppGroup
	GetDomain() *string
	SetName(v string) *AppGroup
	GetName() *string
	SetOrder(v *AppGroupOrder) *AppGroup
	GetOrder() *AppGroupOrder
	SetQuota(v *Quota) *AppGroup
	GetQuota() *Quota
	SetResourceGroupId(v string) *AppGroup
	GetResourceGroupId() *string
	SetType(v string) *AppGroup
	GetType() *string
}

type AppGroup struct {
	ChargeType      *string        `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	Description     *string        `json:"description,omitempty" xml:"description,omitempty"`
	Domain          *string        `json:"domain,omitempty" xml:"domain,omitempty"`
	Name            *string        `json:"name,omitempty" xml:"name,omitempty"`
	Order           *AppGroupOrder `json:"order,omitempty" xml:"order,omitempty" type:"Struct"`
	Quota           *Quota         `json:"quota,omitempty" xml:"quota,omitempty"`
	ResourceGroupId *string        `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Type            *string        `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AppGroup) String() string {
	return dara.Prettify(s)
}

func (s AppGroup) GoString() string {
	return s.String()
}

func (s *AppGroup) GetChargeType() *string {
	return s.ChargeType
}

func (s *AppGroup) GetDescription() *string {
	return s.Description
}

func (s *AppGroup) GetDomain() *string {
	return s.Domain
}

func (s *AppGroup) GetName() *string {
	return s.Name
}

func (s *AppGroup) GetOrder() *AppGroupOrder {
	return s.Order
}

func (s *AppGroup) GetQuota() *Quota {
	return s.Quota
}

func (s *AppGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AppGroup) GetType() *string {
	return s.Type
}

func (s *AppGroup) SetChargeType(v string) *AppGroup {
	s.ChargeType = &v
	return s
}

func (s *AppGroup) SetDescription(v string) *AppGroup {
	s.Description = &v
	return s
}

func (s *AppGroup) SetDomain(v string) *AppGroup {
	s.Domain = &v
	return s
}

func (s *AppGroup) SetName(v string) *AppGroup {
	s.Name = &v
	return s
}

func (s *AppGroup) SetOrder(v *AppGroupOrder) *AppGroup {
	s.Order = v
	return s
}

func (s *AppGroup) SetQuota(v *Quota) *AppGroup {
	s.Quota = v
	return s
}

func (s *AppGroup) SetResourceGroupId(v string) *AppGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *AppGroup) SetType(v string) *AppGroup {
	s.Type = &v
	return s
}

func (s *AppGroup) Validate() error {
	if s.Order != nil {
		if err := s.Order.Validate(); err != nil {
			return err
		}
	}
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AppGroupOrder struct {
	// example:
	//
	// false
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// example:
	//
	// 1
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
}

func (s AppGroupOrder) String() string {
	return dara.Prettify(s)
}

func (s AppGroupOrder) GoString() string {
	return s.String()
}

func (s *AppGroupOrder) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *AppGroupOrder) GetDuration() *int64 {
	return s.Duration
}

func (s *AppGroupOrder) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *AppGroupOrder) SetAutoRenew(v bool) *AppGroupOrder {
	s.AutoRenew = &v
	return s
}

func (s *AppGroupOrder) SetDuration(v int64) *AppGroupOrder {
	s.Duration = &v
	return s
}

func (s *AppGroupOrder) SetPricingCycle(v string) *AppGroupOrder {
	s.PricingCycle = &v
	return s
}

func (s *AppGroupOrder) Validate() error {
	return dara.Validate(s)
}
