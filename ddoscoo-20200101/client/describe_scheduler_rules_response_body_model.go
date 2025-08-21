// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSchedulerRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSchedulerRulesResponseBody
	GetRequestId() *string
	SetSchedulerRules(v []*DescribeSchedulerRulesResponseBodySchedulerRules) *DescribeSchedulerRulesResponseBody
	GetSchedulerRules() []*DescribeSchedulerRulesResponseBodySchedulerRules
	SetTotalCount(v string) *DescribeSchedulerRulesResponseBody
	GetTotalCount() *string
}

type DescribeSchedulerRulesResponseBody struct {
	// example:
	//
	// 11C55595-1757-4B17-9ACE-4ACB68C2D989
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SchedulerRules []*DescribeSchedulerRulesResponseBodySchedulerRules `json:"SchedulerRules,omitempty" xml:"SchedulerRules,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSchedulerRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchedulerRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSchedulerRulesResponseBody) GetSchedulerRules() []*DescribeSchedulerRulesResponseBodySchedulerRules {
	return s.SchedulerRules
}

func (s *DescribeSchedulerRulesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeSchedulerRulesResponseBody) SetRequestId(v string) *DescribeSchedulerRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBody) SetSchedulerRules(v []*DescribeSchedulerRulesResponseBodySchedulerRules) *DescribeSchedulerRulesResponseBody {
	s.SchedulerRules = v
	return s
}

func (s *DescribeSchedulerRulesResponseBody) SetTotalCount(v string) *DescribeSchedulerRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSchedulerRulesResponseBodySchedulerRules struct {
	// example:
	//
	// 4eru5229a843****.aliyunddos0001.com
	Cname *string                                                `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Param *DescribeSchedulerRulesResponseBodySchedulerRulesParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// example:
	//
	// doctest
	RuleName *string                                                  `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType *string                                                  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Rules    []*DescribeSchedulerRulesResponseBodySchedulerRulesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeSchedulerRulesResponseBodySchedulerRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchedulerRulesResponseBodySchedulerRules) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) GetCname() *string {
	return s.Cname
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) GetParam() *DescribeSchedulerRulesResponseBodySchedulerRulesParam {
	return s.Param
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) GetRules() []*DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	return s.Rules
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) SetCname(v string) *DescribeSchedulerRulesResponseBodySchedulerRules {
	s.Cname = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) SetParam(v *DescribeSchedulerRulesResponseBodySchedulerRulesParam) *DescribeSchedulerRulesResponseBodySchedulerRules {
	s.Param = v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) SetRuleName(v string) *DescribeSchedulerRulesResponseBodySchedulerRules {
	s.RuleName = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) SetRuleType(v string) *DescribeSchedulerRulesResponseBodySchedulerRules {
	s.RuleType = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) SetRules(v []*DescribeSchedulerRulesResponseBodySchedulerRulesRules) *DescribeSchedulerRulesResponseBodySchedulerRules {
	s.Rules = v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) Validate() error {
	return dara.Validate(s)
}

type DescribeSchedulerRulesResponseBodySchedulerRulesParam struct {
	ParamData *DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData `json:"ParamData,omitempty" xml:"ParamData,omitempty" type:"Struct"`
	// example:
	//
	// GA
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s DescribeSchedulerRulesResponseBodySchedulerRulesParam) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchedulerRulesResponseBodySchedulerRulesParam) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesParam) GetParamData() *DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData {
	return s.ParamData
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesParam) GetParamType() *string {
	return s.ParamType
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesParam) SetParamData(v *DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData) *DescribeSchedulerRulesResponseBodySchedulerRulesParam {
	s.ParamData = v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesParam) SetParamType(v string) *DescribeSchedulerRulesResponseBodySchedulerRulesParam {
	s.ParamType = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesParam) Validate() error {
	return dara.Validate(s)
}

type DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData struct {
	// example:
	//
	// ga-bp1htlajy5509rc99****
	CloudInstanceId *string `json:"CloudInstanceId,omitempty" xml:"CloudInstanceId,omitempty"`
}

func (s DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData) GetCloudInstanceId() *string {
	return s.CloudInstanceId
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData) SetCloudInstanceId(v string) *DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData {
	s.CloudInstanceId = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData) Validate() error {
	return dara.Validate(s)
}

type DescribeSchedulerRulesResponseBodySchedulerRulesRules struct {
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// example:
	//
	// 100
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 60
	RestoreDelay *int32 `json:"RestoreDelay,omitempty" xml:"RestoreDelay,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 203.***.***.39
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 1
	ValueType *int32 `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s DescribeSchedulerRulesResponseBodySchedulerRulesRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchedulerRulesResponseBodySchedulerRulesRules) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) GetLine() *string {
	return s.Line
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) GetRestoreDelay() *int32 {
	return s.RestoreDelay
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) GetType() *string {
	return s.Type
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) GetValue() *string {
	return s.Value
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) GetValueType() *int32 {
	return s.ValueType
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetLine(v string) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.Line = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetPriority(v int32) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.Priority = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetRegionId(v string) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.RegionId = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetRestoreDelay(v int32) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.RestoreDelay = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetStatus(v int32) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.Status = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetType(v string) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.Type = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetValue(v string) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.Value = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetValueType(v int32) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.ValueType = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) Validate() error {
	return dara.Validate(s)
}
