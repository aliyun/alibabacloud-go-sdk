// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataAgentBillingInstance interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSeats(v int32) *DataAgentBillingInstance
	GetAgentSeats() *int32
	SetBillingInstanceId(v string) *DataAgentBillingInstance
	GetBillingInstanceId() *string
	SetBoundWorkspaceIds(v []*string) *DataAgentBillingInstance
	GetBoundWorkspaceIds() []*string
	SetChargeType(v string) *DataAgentBillingInstance
	GetChargeType() *string
	SetCommodityCode(v string) *DataAgentBillingInstance
	GetCommodityCode() *string
	SetExpireTime(v int64) *DataAgentBillingInstance
	GetExpireTime() *int64
	SetFreeAgentSeats(v int32) *DataAgentBillingInstance
	GetFreeAgentSeats() *int32
	SetGmtCreated(v int64) *DataAgentBillingInstance
	GetGmtCreated() *int64
	SetGmtModified(v int64) *DataAgentBillingInstance
	GetGmtModified() *int64
	SetIsDefault(v bool) *DataAgentBillingInstance
	GetIsDefault() *bool
	SetLLM(v int32) *DataAgentBillingInstance
	GetLLM() *int32
	SetPayLevel(v string) *DataAgentBillingInstance
	GetPayLevel() *string
	SetSessionAvailableDurationQuota(v int32) *DataAgentBillingInstance
	GetSessionAvailableDurationQuota() *int32
	SetSessionSeats(v int32) *DataAgentBillingInstance
	GetSessionSeats() *int32
	SetTokenLimit(v int32) *DataAgentBillingInstance
	GetTokenLimit() *int32
}

type DataAgentBillingInstance struct {
	AgentSeats                    *int32    `json:"AgentSeats,omitempty" xml:"AgentSeats,omitempty"`
	BillingInstanceId             *string   `json:"BillingInstanceId,omitempty" xml:"BillingInstanceId,omitempty"`
	BoundWorkspaceIds             []*string `json:"BoundWorkspaceIds,omitempty" xml:"BoundWorkspaceIds,omitempty" type:"Repeated"`
	ChargeType                    *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CommodityCode                 *string   `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ExpireTime                    *int64    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FreeAgentSeats                *int32    `json:"FreeAgentSeats,omitempty" xml:"FreeAgentSeats,omitempty"`
	GmtCreated                    *int64    `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified                   *int64    `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IsDefault                     *bool     `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	LLM                           *int32    `json:"LLM,omitempty" xml:"LLM,omitempty"`
	PayLevel                      *string   `json:"PayLevel,omitempty" xml:"PayLevel,omitempty"`
	SessionAvailableDurationQuota *int32    `json:"SessionAvailableDurationQuota,omitempty" xml:"SessionAvailableDurationQuota,omitempty"`
	SessionSeats                  *int32    `json:"SessionSeats,omitempty" xml:"SessionSeats,omitempty"`
	TokenLimit                    *int32    `json:"TokenLimit,omitempty" xml:"TokenLimit,omitempty"`
}

func (s DataAgentBillingInstance) String() string {
	return dara.Prettify(s)
}

func (s DataAgentBillingInstance) GoString() string {
	return s.String()
}

func (s *DataAgentBillingInstance) GetAgentSeats() *int32 {
	return s.AgentSeats
}

func (s *DataAgentBillingInstance) GetBillingInstanceId() *string {
	return s.BillingInstanceId
}

func (s *DataAgentBillingInstance) GetBoundWorkspaceIds() []*string {
	return s.BoundWorkspaceIds
}

func (s *DataAgentBillingInstance) GetChargeType() *string {
	return s.ChargeType
}

func (s *DataAgentBillingInstance) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DataAgentBillingInstance) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DataAgentBillingInstance) GetFreeAgentSeats() *int32 {
	return s.FreeAgentSeats
}

func (s *DataAgentBillingInstance) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *DataAgentBillingInstance) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DataAgentBillingInstance) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DataAgentBillingInstance) GetLLM() *int32 {
	return s.LLM
}

func (s *DataAgentBillingInstance) GetPayLevel() *string {
	return s.PayLevel
}

func (s *DataAgentBillingInstance) GetSessionAvailableDurationQuota() *int32 {
	return s.SessionAvailableDurationQuota
}

func (s *DataAgentBillingInstance) GetSessionSeats() *int32 {
	return s.SessionSeats
}

func (s *DataAgentBillingInstance) GetTokenLimit() *int32 {
	return s.TokenLimit
}

func (s *DataAgentBillingInstance) SetAgentSeats(v int32) *DataAgentBillingInstance {
	s.AgentSeats = &v
	return s
}

func (s *DataAgentBillingInstance) SetBillingInstanceId(v string) *DataAgentBillingInstance {
	s.BillingInstanceId = &v
	return s
}

func (s *DataAgentBillingInstance) SetBoundWorkspaceIds(v []*string) *DataAgentBillingInstance {
	s.BoundWorkspaceIds = v
	return s
}

func (s *DataAgentBillingInstance) SetChargeType(v string) *DataAgentBillingInstance {
	s.ChargeType = &v
	return s
}

func (s *DataAgentBillingInstance) SetCommodityCode(v string) *DataAgentBillingInstance {
	s.CommodityCode = &v
	return s
}

func (s *DataAgentBillingInstance) SetExpireTime(v int64) *DataAgentBillingInstance {
	s.ExpireTime = &v
	return s
}

func (s *DataAgentBillingInstance) SetFreeAgentSeats(v int32) *DataAgentBillingInstance {
	s.FreeAgentSeats = &v
	return s
}

func (s *DataAgentBillingInstance) SetGmtCreated(v int64) *DataAgentBillingInstance {
	s.GmtCreated = &v
	return s
}

func (s *DataAgentBillingInstance) SetGmtModified(v int64) *DataAgentBillingInstance {
	s.GmtModified = &v
	return s
}

func (s *DataAgentBillingInstance) SetIsDefault(v bool) *DataAgentBillingInstance {
	s.IsDefault = &v
	return s
}

func (s *DataAgentBillingInstance) SetLLM(v int32) *DataAgentBillingInstance {
	s.LLM = &v
	return s
}

func (s *DataAgentBillingInstance) SetPayLevel(v string) *DataAgentBillingInstance {
	s.PayLevel = &v
	return s
}

func (s *DataAgentBillingInstance) SetSessionAvailableDurationQuota(v int32) *DataAgentBillingInstance {
	s.SessionAvailableDurationQuota = &v
	return s
}

func (s *DataAgentBillingInstance) SetSessionSeats(v int32) *DataAgentBillingInstance {
	s.SessionSeats = &v
	return s
}

func (s *DataAgentBillingInstance) SetTokenLimit(v int32) *DataAgentBillingInstance {
	s.TokenLimit = &v
	return s
}

func (s *DataAgentBillingInstance) Validate() error {
	return dara.Validate(s)
}
