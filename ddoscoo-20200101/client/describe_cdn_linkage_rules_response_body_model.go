// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnLinkageRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCdnLinkageRulesResponseBody
	GetRequestId() *string
	SetSchedulerRules(v []*DescribeCdnLinkageRulesResponseBodySchedulerRules) *DescribeCdnLinkageRulesResponseBody
	GetSchedulerRules() []*DescribeCdnLinkageRulesResponseBodySchedulerRules
	SetTotalCount(v string) *DescribeCdnLinkageRulesResponseBody
	GetTotalCount() *string
}

type DescribeCdnLinkageRulesResponseBody struct {
	// example:
	//
	// 02FE96D9-C77B-5735-B36D-329E052C8047
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SchedulerRules []*DescribeCdnLinkageRulesResponseBodySchedulerRules `json:"SchedulerRules,omitempty" xml:"SchedulerRules,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCdnLinkageRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnLinkageRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnLinkageRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnLinkageRulesResponseBody) GetSchedulerRules() []*DescribeCdnLinkageRulesResponseBodySchedulerRules {
	return s.SchedulerRules
}

func (s *DescribeCdnLinkageRulesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeCdnLinkageRulesResponseBody) SetRequestId(v string) *DescribeCdnLinkageRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBody) SetSchedulerRules(v []*DescribeCdnLinkageRulesResponseBodySchedulerRules) *DescribeCdnLinkageRulesResponseBody {
	s.SchedulerRules = v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBody) SetTotalCount(v string) *DescribeCdnLinkageRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBody) Validate() error {
	if s.SchedulerRules != nil {
		for _, item := range s.SchedulerRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnLinkageRulesResponseBodySchedulerRules struct {
	// example:
	//
	// 0
	CdnLinkageEnable *int32                                                           `json:"CdnLinkageEnable,omitempty" xml:"CdnLinkageEnable,omitempty"`
	CdnLinkageRule   *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule `json:"CdnLinkageRule,omitempty" xml:"CdnLinkageRule,omitempty" type:"Struct"`
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeCdnLinkageRulesResponseBodySchedulerRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnLinkageRulesResponseBodySchedulerRules) GoString() string {
	return s.String()
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRules) GetCdnLinkageEnable() *int32 {
	return s.CdnLinkageEnable
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRules) GetCdnLinkageRule() *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule {
	return s.CdnLinkageRule
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRules) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRules) SetCdnLinkageEnable(v int32) *DescribeCdnLinkageRulesResponseBodySchedulerRules {
	s.CdnLinkageEnable = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRules) SetCdnLinkageRule(v *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule) *DescribeCdnLinkageRulesResponseBodySchedulerRules {
	s.CdnLinkageRule = v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRules) SetDomain(v string) *DescribeCdnLinkageRulesResponseBodySchedulerRules {
	s.Domain = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRules) Validate() error {
	if s.CdnLinkageRule != nil {
		if err := s.CdnLinkageRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule struct {
	// example:
	//
	// example.aliyundoc.com
	Cname *string                                                               `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Param *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// example:
	//
	// testDDos
	RuleName *string                                                                 `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Rules    []*DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule) GoString() string {
	return s.String()
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule) GetCname() *string {
	return s.Cname
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule) GetParam() *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam {
	return s.Param
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule) GetRules() []*DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules {
	return s.Rules
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule) SetCname(v string) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule {
	s.Cname = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule) SetParam(v *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule {
	s.Param = v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule) SetRuleName(v string) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule {
	s.RuleName = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule) SetRules(v []*DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule {
	s.Rules = v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRule) Validate() error {
	if s.Param != nil {
		if err := s.Param.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam struct {
	ParamData *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData `json:"ParamData,omitempty" xml:"ParamData,omitempty" type:"Struct"`
	// example:
	//
	// cdn
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam) GoString() string {
	return s.String()
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam) GetParamData() *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData {
	return s.ParamData
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam) GetParamType() *string {
	return s.ParamType
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam) SetParamData(v *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam {
	s.ParamData = v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam) SetParamType(v string) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam {
	s.ParamType = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParam) Validate() error {
	if s.ParamData != nil {
		if err := s.ParamData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData struct {
	// example:
	//
	// 100
	AccessQps *int64 `json:"AccessQps,omitempty" xml:"AccessQps,omitempty"`
	// example:
	//
	// 0
	UpstreamQps *int64 `json:"UpstreamQps,omitempty" xml:"UpstreamQps,omitempty"`
}

func (s DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData) GoString() string {
	return s.String()
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData) GetAccessQps() *int64 {
	return s.AccessQps
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData) GetUpstreamQps() *int64 {
	return s.UpstreamQps
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData) SetAccessQps(v int64) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData {
	s.AccessQps = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData) SetUpstreamQps(v int64) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData {
	s.UpstreamQps = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleParamParamData) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules struct {
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// ""
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 203.107.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 1
	ValueType *int32 `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) GoString() string {
	return s.String()
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) GetType() *string {
	return s.Type
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) GetValue() *string {
	return s.Value
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) GetValueType() *int32 {
	return s.ValueType
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) SetPriority(v int32) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules {
	s.Priority = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) SetRegionId(v string) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules {
	s.RegionId = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) SetStatus(v int32) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules {
	s.Status = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) SetType(v string) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules {
	s.Type = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) SetValue(v string) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules {
	s.Value = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) SetValueType(v int32) *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules {
	s.ValueType = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponseBodySchedulerRulesCdnLinkageRuleRules) Validate() error {
	return dara.Validate(s)
}
