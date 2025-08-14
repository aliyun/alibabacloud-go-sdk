// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCostCenterShareRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateShareRuleListShrink(v string) *SaveCostCenterShareRuleShrinkRequest
	GetCreateShareRuleListShrink() *string
	SetModifyShareRuleListShrink(v string) *SaveCostCenterShareRuleShrinkRequest
	GetModifyShareRuleListShrink() *string
	SetNbid(v string) *SaveCostCenterShareRuleShrinkRequest
	GetNbid() *string
	SetOwnerAccountId(v int64) *SaveCostCenterShareRuleShrinkRequest
	GetOwnerAccountId() *int64
	SetRemoveShareRuleListShrink(v string) *SaveCostCenterShareRuleShrinkRequest
	GetRemoveShareRuleListShrink() *string
}

type SaveCostCenterShareRuleShrinkRequest struct {
	CreateShareRuleListShrink *string `json:"CreateShareRuleList,omitempty" xml:"CreateShareRuleList,omitempty"`
	ModifyShareRuleListShrink *string `json:"ModifyShareRuleList,omitempty" xml:"ModifyShareRuleList,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 1977800748053695
	OwnerAccountId            *int64  `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	RemoveShareRuleListShrink *string `json:"RemoveShareRuleList,omitempty" xml:"RemoveShareRuleList,omitempty"`
}

func (s SaveCostCenterShareRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveCostCenterShareRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveCostCenterShareRuleShrinkRequest) GetCreateShareRuleListShrink() *string {
	return s.CreateShareRuleListShrink
}

func (s *SaveCostCenterShareRuleShrinkRequest) GetModifyShareRuleListShrink() *string {
	return s.ModifyShareRuleListShrink
}

func (s *SaveCostCenterShareRuleShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *SaveCostCenterShareRuleShrinkRequest) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *SaveCostCenterShareRuleShrinkRequest) GetRemoveShareRuleListShrink() *string {
	return s.RemoveShareRuleListShrink
}

func (s *SaveCostCenterShareRuleShrinkRequest) SetCreateShareRuleListShrink(v string) *SaveCostCenterShareRuleShrinkRequest {
	s.CreateShareRuleListShrink = &v
	return s
}

func (s *SaveCostCenterShareRuleShrinkRequest) SetModifyShareRuleListShrink(v string) *SaveCostCenterShareRuleShrinkRequest {
	s.ModifyShareRuleListShrink = &v
	return s
}

func (s *SaveCostCenterShareRuleShrinkRequest) SetNbid(v string) *SaveCostCenterShareRuleShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *SaveCostCenterShareRuleShrinkRequest) SetOwnerAccountId(v int64) *SaveCostCenterShareRuleShrinkRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *SaveCostCenterShareRuleShrinkRequest) SetRemoveShareRuleListShrink(v string) *SaveCostCenterShareRuleShrinkRequest {
	s.RemoveShareRuleListShrink = &v
	return s
}

func (s *SaveCostCenterShareRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
