// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSavingPlanUserDeductRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcIdAccountIdsShrink(v string) *SetSavingPlanUserDeductRuleShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetNbid(v string) *SetSavingPlanUserDeductRuleShrinkRequest
	GetNbid() *string
	SetSpnInstanceCode(v string) *SetSavingPlanUserDeductRuleShrinkRequest
	GetSpnInstanceCode() *string
	SetUserDeductRulesShrink(v string) *SetSavingPlanUserDeductRuleShrinkRequest
	GetUserDeductRulesShrink() *string
}

type SetSavingPlanUserDeductRuleShrinkRequest struct {
	EcIdAccountIdsShrink  *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                  *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	SpnInstanceCode       *string `json:"SpnInstanceCode,omitempty" xml:"SpnInstanceCode,omitempty"`
	UserDeductRulesShrink *string `json:"UserDeductRules,omitempty" xml:"UserDeductRules,omitempty"`
}

func (s SetSavingPlanUserDeductRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSavingPlanUserDeductRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) GetSpnInstanceCode() *string {
	return s.SpnInstanceCode
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) GetUserDeductRulesShrink() *string {
	return s.UserDeductRulesShrink
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) SetEcIdAccountIdsShrink(v string) *SetSavingPlanUserDeductRuleShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) SetNbid(v string) *SetSavingPlanUserDeductRuleShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) SetSpnInstanceCode(v string) *SetSavingPlanUserDeductRuleShrinkRequest {
	s.SpnInstanceCode = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) SetUserDeductRulesShrink(v string) *SetSavingPlanUserDeductRuleShrinkRequest {
	s.UserDeductRulesShrink = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
