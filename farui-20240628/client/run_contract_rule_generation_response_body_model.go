// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContractRuleGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunContractRuleGenerationResponseBody
	GetCode() *string
	SetMessage(v string) *RunContractRuleGenerationResponseBody
	GetMessage() *string
	SetOutput(v *RunContractRuleGenerationResponseBodyOutput) *RunContractRuleGenerationResponseBody
	GetOutput() *RunContractRuleGenerationResponseBodyOutput
	SetRequestId(v string) *RunContractRuleGenerationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunContractRuleGenerationResponseBody
	GetSuccess() *bool
	SetUsage(v *RunContractRuleGenerationResponseBodyUsage) *RunContractRuleGenerationResponseBody
	GetUsage() *RunContractRuleGenerationResponseBodyUsage
	SetHttpStatusCode(v int32) *RunContractRuleGenerationResponseBody
	GetHttpStatusCode() *int32
}

type RunContractRuleGenerationResponseBody struct {
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Message *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Output  *RunContractRuleGenerationResponseBodyOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// example:
	//
	// 744419D0-671A-5997-9840-E8AE48356194
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Usage   *RunContractRuleGenerationResponseBodyUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
}

func (s RunContractRuleGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunContractRuleGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunContractRuleGenerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunContractRuleGenerationResponseBody) GetOutput() *RunContractRuleGenerationResponseBodyOutput {
	return s.Output
}

func (s *RunContractRuleGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunContractRuleGenerationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunContractRuleGenerationResponseBody) GetUsage() *RunContractRuleGenerationResponseBodyUsage {
	return s.Usage
}

func (s *RunContractRuleGenerationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RunContractRuleGenerationResponseBody) SetCode(v string) *RunContractRuleGenerationResponseBody {
	s.Code = &v
	return s
}

func (s *RunContractRuleGenerationResponseBody) SetMessage(v string) *RunContractRuleGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *RunContractRuleGenerationResponseBody) SetOutput(v *RunContractRuleGenerationResponseBodyOutput) *RunContractRuleGenerationResponseBody {
	s.Output = v
	return s
}

func (s *RunContractRuleGenerationResponseBody) SetRequestId(v string) *RunContractRuleGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunContractRuleGenerationResponseBody) SetSuccess(v bool) *RunContractRuleGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *RunContractRuleGenerationResponseBody) SetUsage(v *RunContractRuleGenerationResponseBodyUsage) *RunContractRuleGenerationResponseBody {
	s.Usage = v
	return s
}

func (s *RunContractRuleGenerationResponseBody) SetHttpStatusCode(v int32) *RunContractRuleGenerationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunContractRuleGenerationResponseBody) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunContractRuleGenerationResponseBodyOutput struct {
	// example:
	//
	// b265b416-ca1f-425d-9340-c968f39624e9
	RuleTaskId *string                                             `json:"ruleTaskId,omitempty" xml:"ruleTaskId,omitempty"`
	Rules      []*RunContractRuleGenerationResponseBodyOutputRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s RunContractRuleGenerationResponseBodyOutput) String() string {
	return dara.Prettify(s)
}

func (s RunContractRuleGenerationResponseBodyOutput) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationResponseBodyOutput) GetRuleTaskId() *string {
	return s.RuleTaskId
}

func (s *RunContractRuleGenerationResponseBodyOutput) GetRules() []*RunContractRuleGenerationResponseBodyOutputRules {
	return s.Rules
}

func (s *RunContractRuleGenerationResponseBodyOutput) SetRuleTaskId(v string) *RunContractRuleGenerationResponseBodyOutput {
	s.RuleTaskId = &v
	return s
}

func (s *RunContractRuleGenerationResponseBodyOutput) SetRules(v []*RunContractRuleGenerationResponseBodyOutputRules) *RunContractRuleGenerationResponseBodyOutput {
	s.Rules = v
	return s
}

func (s *RunContractRuleGenerationResponseBodyOutput) Validate() error {
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

type RunContractRuleGenerationResponseBodyOutputRules struct {
	// example:
	//
	// medium
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// example:
	//
	// 1.1
	RuleSequence *string `json:"ruleSequence,omitempty" xml:"ruleSequence,omitempty"`
	RuleTag      *string `json:"ruleTag,omitempty" xml:"ruleTag,omitempty"`
	RuleTitle    *string `json:"ruleTitle,omitempty" xml:"ruleTitle,omitempty"`
}

func (s RunContractRuleGenerationResponseBodyOutputRules) String() string {
	return dara.Prettify(s)
}

func (s RunContractRuleGenerationResponseBodyOutputRules) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) GetRuleSequence() *string {
	return s.RuleSequence
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) GetRuleTag() *string {
	return s.RuleTag
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) GetRuleTitle() *string {
	return s.RuleTitle
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) SetRiskLevel(v string) *RunContractRuleGenerationResponseBodyOutputRules {
	s.RiskLevel = &v
	return s
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) SetRuleSequence(v string) *RunContractRuleGenerationResponseBodyOutputRules {
	s.RuleSequence = &v
	return s
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) SetRuleTag(v string) *RunContractRuleGenerationResponseBodyOutputRules {
	s.RuleTag = &v
	return s
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) SetRuleTitle(v string) *RunContractRuleGenerationResponseBodyOutputRules {
	s.RuleTitle = &v
	return s
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) Validate() error {
	return dara.Validate(s)
}

type RunContractRuleGenerationResponseBodyUsage struct {
	// example:
	//
	// 5
	Input *int64 `json:"input,omitempty" xml:"input,omitempty"`
	// example:
	//
	// page
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s RunContractRuleGenerationResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s RunContractRuleGenerationResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationResponseBodyUsage) GetInput() *int64 {
	return s.Input
}

func (s *RunContractRuleGenerationResponseBodyUsage) GetUnit() *string {
	return s.Unit
}

func (s *RunContractRuleGenerationResponseBodyUsage) SetInput(v int64) *RunContractRuleGenerationResponseBodyUsage {
	s.Input = &v
	return s
}

func (s *RunContractRuleGenerationResponseBodyUsage) SetUnit(v string) *RunContractRuleGenerationResponseBodyUsage {
	s.Unit = &v
	return s
}

func (s *RunContractRuleGenerationResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
