// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionDTO interface {
	dara.Model
	String() string
	GoString() string
	SetBalanceType(v string) *SubscriptionDTO
	GetBalanceType() *string
	SetClientId(v int64) *SubscriptionDTO
	GetClientId() *int64
	SetCreateTime(v string) *SubscriptionDTO
	GetCreateTime() *string
	SetId(v int64) *SubscriptionDTO
	GetId() *int64
	SetStatus(v string) *SubscriptionDTO
	GetStatus() *string
	SetStopTime(v string) *SubscriptionDTO
	GetStopTime() *string
	SetSubscriptionAmount(v float64) *SubscriptionDTO
	GetSubscriptionAmount() *float64
	SetUpdateTime(v string) *SubscriptionDTO
	GetUpdateTime() *string
	SetValidFrom(v string) *SubscriptionDTO
	GetValidFrom() *string
}

type SubscriptionDTO struct {
	// The balance type (permanent/monthly).
	//
	// example:
	//
	// permanent
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// The department ID.
	//
	// example:
	//
	// 100
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-06-15T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The subscription ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The subscription status. Valid values:
	//
	// - active: The subscription is active.
	//
	// - stopped: The subscription is stopped.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The stop time. This value is empty if the subscription has not been stopped.
	//
	// example:
	//
	// 2026-07-01T00:00:00Z
	StopTime *string `json:"stopTime,omitempty" xml:"stopTime,omitempty"`
	// The subscription recharge amount.
	//
	// example:
	//
	// 100.00
	SubscriptionAmount *float64 `json:"subscriptionAmount,omitempty" xml:"subscriptionAmount,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2024-06-15T10:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The effective period.
	//
	// example:
	//
	// 2024-07-01T00:00:00Z
	ValidFrom *string `json:"validFrom,omitempty" xml:"validFrom,omitempty"`
}

func (s SubscriptionDTO) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionDTO) GoString() string {
	return s.String()
}

func (s *SubscriptionDTO) GetBalanceType() *string {
	return s.BalanceType
}

func (s *SubscriptionDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *SubscriptionDTO) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SubscriptionDTO) GetId() *int64 {
	return s.Id
}

func (s *SubscriptionDTO) GetStatus() *string {
	return s.Status
}

func (s *SubscriptionDTO) GetStopTime() *string {
	return s.StopTime
}

func (s *SubscriptionDTO) GetSubscriptionAmount() *float64 {
	return s.SubscriptionAmount
}

func (s *SubscriptionDTO) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SubscriptionDTO) GetValidFrom() *string {
	return s.ValidFrom
}

func (s *SubscriptionDTO) SetBalanceType(v string) *SubscriptionDTO {
	s.BalanceType = &v
	return s
}

func (s *SubscriptionDTO) SetClientId(v int64) *SubscriptionDTO {
	s.ClientId = &v
	return s
}

func (s *SubscriptionDTO) SetCreateTime(v string) *SubscriptionDTO {
	s.CreateTime = &v
	return s
}

func (s *SubscriptionDTO) SetId(v int64) *SubscriptionDTO {
	s.Id = &v
	return s
}

func (s *SubscriptionDTO) SetStatus(v string) *SubscriptionDTO {
	s.Status = &v
	return s
}

func (s *SubscriptionDTO) SetStopTime(v string) *SubscriptionDTO {
	s.StopTime = &v
	return s
}

func (s *SubscriptionDTO) SetSubscriptionAmount(v float64) *SubscriptionDTO {
	s.SubscriptionAmount = &v
	return s
}

func (s *SubscriptionDTO) SetUpdateTime(v string) *SubscriptionDTO {
	s.UpdateTime = &v
	return s
}

func (s *SubscriptionDTO) SetValidFrom(v string) *SubscriptionDTO {
	s.ValidFrom = &v
	return s
}

func (s *SubscriptionDTO) Validate() error {
	return dara.Validate(s)
}
