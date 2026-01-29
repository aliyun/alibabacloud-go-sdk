// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSavingPlanUserDeductRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcIdAccountIds(v []*SetSavingPlanUserDeductRuleRequestEcIdAccountIds) *SetSavingPlanUserDeductRuleRequest
	GetEcIdAccountIds() []*SetSavingPlanUserDeductRuleRequestEcIdAccountIds
	SetNbid(v string) *SetSavingPlanUserDeductRuleRequest
	GetNbid() *string
	SetSpnInstanceCode(v string) *SetSavingPlanUserDeductRuleRequest
	GetSpnInstanceCode() *string
	SetUserDeductRules(v []*SetSavingPlanUserDeductRuleRequestUserDeductRules) *SetSavingPlanUserDeductRuleRequest
	GetUserDeductRules() []*SetSavingPlanUserDeductRuleRequestUserDeductRules
}

type SetSavingPlanUserDeductRuleRequest struct {
	EcIdAccountIds  []*SetSavingPlanUserDeductRuleRequestEcIdAccountIds  `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid            *string                                              `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	SpnInstanceCode *string                                              `json:"SpnInstanceCode,omitempty" xml:"SpnInstanceCode,omitempty"`
	UserDeductRules []*SetSavingPlanUserDeductRuleRequestUserDeductRules `json:"UserDeductRules,omitempty" xml:"UserDeductRules,omitempty" type:"Repeated"`
}

func (s SetSavingPlanUserDeductRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSavingPlanUserDeductRuleRequest) GoString() string {
	return s.String()
}

func (s *SetSavingPlanUserDeductRuleRequest) GetEcIdAccountIds() []*SetSavingPlanUserDeductRuleRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *SetSavingPlanUserDeductRuleRequest) GetNbid() *string {
	return s.Nbid
}

func (s *SetSavingPlanUserDeductRuleRequest) GetSpnInstanceCode() *string {
	return s.SpnInstanceCode
}

func (s *SetSavingPlanUserDeductRuleRequest) GetUserDeductRules() []*SetSavingPlanUserDeductRuleRequestUserDeductRules {
	return s.UserDeductRules
}

func (s *SetSavingPlanUserDeductRuleRequest) SetEcIdAccountIds(v []*SetSavingPlanUserDeductRuleRequestEcIdAccountIds) *SetSavingPlanUserDeductRuleRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequest) SetNbid(v string) *SetSavingPlanUserDeductRuleRequest {
	s.Nbid = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequest) SetSpnInstanceCode(v string) *SetSavingPlanUserDeductRuleRequest {
	s.SpnInstanceCode = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequest) SetUserDeductRules(v []*SetSavingPlanUserDeductRuleRequestUserDeductRules) *SetSavingPlanUserDeductRuleRequest {
	s.UserDeductRules = v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequest) Validate() error {
	if s.EcIdAccountIds != nil {
		for _, item := range s.EcIdAccountIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserDeductRules != nil {
		for _, item := range s.UserDeductRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetSavingPlanUserDeductRuleRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	EcId       *string  `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s SetSavingPlanUserDeductRuleRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s SetSavingPlanUserDeductRuleRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *SetSavingPlanUserDeductRuleRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *SetSavingPlanUserDeductRuleRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *SetSavingPlanUserDeductRuleRequestEcIdAccountIds) SetAccountIds(v []*int64) *SetSavingPlanUserDeductRuleRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequestEcIdAccountIds) SetEcId(v string) *SetSavingPlanUserDeductRuleRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}

type SetSavingPlanUserDeductRuleRequestUserDeductRules struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ModuleCode    *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	SkipDeduct    *bool   `json:"SkipDeduct,omitempty" xml:"SkipDeduct,omitempty"`
}

func (s SetSavingPlanUserDeductRuleRequestUserDeductRules) String() string {
	return dara.Prettify(s)
}

func (s SetSavingPlanUserDeductRuleRequestUserDeductRules) GoString() string {
	return s.String()
}

func (s *SetSavingPlanUserDeductRuleRequestUserDeductRules) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *SetSavingPlanUserDeductRuleRequestUserDeductRules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *SetSavingPlanUserDeductRuleRequestUserDeductRules) GetSkipDeduct() *bool {
	return s.SkipDeduct
}

func (s *SetSavingPlanUserDeductRuleRequestUserDeductRules) SetCommodityCode(v string) *SetSavingPlanUserDeductRuleRequestUserDeductRules {
	s.CommodityCode = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequestUserDeductRules) SetModuleCode(v string) *SetSavingPlanUserDeductRuleRequestUserDeductRules {
	s.ModuleCode = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequestUserDeductRules) SetSkipDeduct(v bool) *SetSavingPlanUserDeductRuleRequestUserDeductRules {
	s.SkipDeduct = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequestUserDeductRules) Validate() error {
	return dara.Validate(s)
}
