// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscription interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*SubscriptionConditions) *Subscription
	GetConditions() []*SubscriptionConditions
	SetCreateTime(v string) *Subscription
	GetCreateTime() *string
	SetDescription(v string) *Subscription
	GetDescription() *string
	SetEnabled(v bool) *Subscription
	GetEnabled() *bool
	SetName(v string) *Subscription
	GetName() *string
	SetProduct(v string) *Subscription
	GetProduct() *string
	SetRelation(v string) *Subscription
	GetRelation() *string
	SetStrategyUuid(v string) *Subscription
	GetStrategyUuid() *string
	SetUpdateTime(v string) *Subscription
	GetUpdateTime() *string
	SetUuid(v string) *Subscription
	GetUuid() *string
}

type Subscription struct {
	Conditions  []*SubscriptionConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	CreateTime  *string                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                   `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled     *bool                     `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Relation     *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	StrategyUuid *string `json:"StrategyUuid,omitempty" xml:"StrategyUuid,omitempty"`
	UpdateTime   *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s Subscription) String() string {
	return dara.Prettify(s)
}

func (s Subscription) GoString() string {
	return s.String()
}

func (s *Subscription) GetConditions() []*SubscriptionConditions {
	return s.Conditions
}

func (s *Subscription) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Subscription) GetDescription() *string {
	return s.Description
}

func (s *Subscription) GetEnabled() *bool {
	return s.Enabled
}

func (s *Subscription) GetName() *string {
	return s.Name
}

func (s *Subscription) GetProduct() *string {
	return s.Product
}

func (s *Subscription) GetRelation() *string {
	return s.Relation
}

func (s *Subscription) GetStrategyUuid() *string {
	return s.StrategyUuid
}

func (s *Subscription) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *Subscription) GetUuid() *string {
	return s.Uuid
}

func (s *Subscription) SetConditions(v []*SubscriptionConditions) *Subscription {
	s.Conditions = v
	return s
}

func (s *Subscription) SetCreateTime(v string) *Subscription {
	s.CreateTime = &v
	return s
}

func (s *Subscription) SetDescription(v string) *Subscription {
	s.Description = &v
	return s
}

func (s *Subscription) SetEnabled(v bool) *Subscription {
	s.Enabled = &v
	return s
}

func (s *Subscription) SetName(v string) *Subscription {
	s.Name = &v
	return s
}

func (s *Subscription) SetProduct(v string) *Subscription {
	s.Product = &v
	return s
}

func (s *Subscription) SetRelation(v string) *Subscription {
	s.Relation = &v
	return s
}

func (s *Subscription) SetStrategyUuid(v string) *Subscription {
	s.StrategyUuid = &v
	return s
}

func (s *Subscription) SetUpdateTime(v string) *Subscription {
	s.UpdateTime = &v
	return s
}

func (s *Subscription) SetUuid(v string) *Subscription {
	s.Uuid = &v
	return s
}

func (s *Subscription) Validate() error {
	return dara.Validate(s)
}

type SubscriptionConditions struct {
	Field    *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SubscriptionConditions) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionConditions) GoString() string {
	return s.String()
}

func (s *SubscriptionConditions) GetField() *string {
	return s.Field
}

func (s *SubscriptionConditions) GetOperator() *string {
	return s.Operator
}

func (s *SubscriptionConditions) GetValue() *string {
	return s.Value
}

func (s *SubscriptionConditions) SetField(v string) *SubscriptionConditions {
	s.Field = &v
	return s
}

func (s *SubscriptionConditions) SetOperator(v string) *SubscriptionConditions {
	s.Operator = &v
	return s
}

func (s *SubscriptionConditions) SetValue(v string) *SubscriptionConditions {
	s.Value = &v
	return s
}

func (s *SubscriptionConditions) Validate() error {
	return dara.Validate(s)
}
