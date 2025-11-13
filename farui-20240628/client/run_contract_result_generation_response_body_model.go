// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContractResultGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunContractResultGenerationResponseBody
	GetCode() *string
	SetMessage(v string) *RunContractResultGenerationResponseBody
	GetMessage() *string
	SetOutput(v *RunContractResultGenerationResponseBodyOutput) *RunContractResultGenerationResponseBody
	GetOutput() *RunContractResultGenerationResponseBodyOutput
	SetRequestId(v string) *RunContractResultGenerationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunContractResultGenerationResponseBody
	GetSuccess() *bool
	SetUsage(v *RunContractResultGenerationResponseBodyUsage) *RunContractResultGenerationResponseBody
	GetUsage() *RunContractResultGenerationResponseBodyUsage
	SetHttpStatusCode(v string) *RunContractResultGenerationResponseBody
	GetHttpStatusCode() *string
}

type RunContractResultGenerationResponseBody struct {
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Message *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Output  *RunContractResultGenerationResponseBodyOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// example:
	//
	// 744419D0-671A-5997-9840-E8AE48356194
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Usage   *RunContractResultGenerationResponseBodyUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
}

func (s RunContractResultGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunContractResultGenerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunContractResultGenerationResponseBody) GetOutput() *RunContractResultGenerationResponseBodyOutput {
	return s.Output
}

func (s *RunContractResultGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunContractResultGenerationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunContractResultGenerationResponseBody) GetUsage() *RunContractResultGenerationResponseBodyUsage {
	return s.Usage
}

func (s *RunContractResultGenerationResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *RunContractResultGenerationResponseBody) SetCode(v string) *RunContractResultGenerationResponseBody {
	s.Code = &v
	return s
}

func (s *RunContractResultGenerationResponseBody) SetMessage(v string) *RunContractResultGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *RunContractResultGenerationResponseBody) SetOutput(v *RunContractResultGenerationResponseBodyOutput) *RunContractResultGenerationResponseBody {
	s.Output = v
	return s
}

func (s *RunContractResultGenerationResponseBody) SetRequestId(v string) *RunContractResultGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunContractResultGenerationResponseBody) SetSuccess(v bool) *RunContractResultGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *RunContractResultGenerationResponseBody) SetUsage(v *RunContractResultGenerationResponseBodyUsage) *RunContractResultGenerationResponseBody {
	s.Usage = v
	return s
}

func (s *RunContractResultGenerationResponseBody) SetHttpStatusCode(v string) *RunContractResultGenerationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunContractResultGenerationResponseBody) Validate() error {
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

type RunContractResultGenerationResponseBodyOutput struct {
	Result *RunContractResultGenerationResponseBodyOutputResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// eaa56e1e-e205-4f5e-926e-5e2269ae7f68
	ResultTaskId *string `json:"resultTaskId,omitempty" xml:"resultTaskId,omitempty"`
}

func (s RunContractResultGenerationResponseBodyOutput) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationResponseBodyOutput) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationResponseBodyOutput) GetResult() *RunContractResultGenerationResponseBodyOutputResult {
	return s.Result
}

func (s *RunContractResultGenerationResponseBodyOutput) GetResultTaskId() *string {
	return s.ResultTaskId
}

func (s *RunContractResultGenerationResponseBodyOutput) SetResult(v *RunContractResultGenerationResponseBodyOutputResult) *RunContractResultGenerationResponseBodyOutput {
	s.Result = v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutput) SetResultTaskId(v string) *RunContractResultGenerationResponseBodyOutput {
	s.ResultTaskId = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutput) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunContractResultGenerationResponseBodyOutputResult struct {
	ExamineBrief  *string `json:"examineBrief,omitempty" xml:"examineBrief,omitempty"`
	ExamineResult *string `json:"examineResult,omitempty" xml:"examineResult,omitempty"`
	// example:
	//
	// high
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// example:
	//
	// 1.1
	RuleSequence *string                                                        `json:"ruleSequence,omitempty" xml:"ruleSequence,omitempty"`
	RuleTag      *string                                                        `json:"ruleTag,omitempty" xml:"ruleTag,omitempty"`
	RuleTitle    *string                                                        `json:"ruleTitle,omitempty" xml:"ruleTitle,omitempty"`
	SubRisks     []*RunContractResultGenerationResponseBodyOutputResultSubRisks `json:"subRisks,omitempty" xml:"subRisks,omitempty" type:"Repeated"`
}

func (s RunContractResultGenerationResponseBodyOutputResult) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationResponseBodyOutputResult) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationResponseBodyOutputResult) GetExamineBrief() *string {
	return s.ExamineBrief
}

func (s *RunContractResultGenerationResponseBodyOutputResult) GetExamineResult() *string {
	return s.ExamineResult
}

func (s *RunContractResultGenerationResponseBodyOutputResult) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *RunContractResultGenerationResponseBodyOutputResult) GetRuleSequence() *string {
	return s.RuleSequence
}

func (s *RunContractResultGenerationResponseBodyOutputResult) GetRuleTag() *string {
	return s.RuleTag
}

func (s *RunContractResultGenerationResponseBodyOutputResult) GetRuleTitle() *string {
	return s.RuleTitle
}

func (s *RunContractResultGenerationResponseBodyOutputResult) GetSubRisks() []*RunContractResultGenerationResponseBodyOutputResultSubRisks {
	return s.SubRisks
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetExamineBrief(v string) *RunContractResultGenerationResponseBodyOutputResult {
	s.ExamineBrief = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetExamineResult(v string) *RunContractResultGenerationResponseBodyOutputResult {
	s.ExamineResult = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetRiskLevel(v string) *RunContractResultGenerationResponseBodyOutputResult {
	s.RiskLevel = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetRuleSequence(v string) *RunContractResultGenerationResponseBodyOutputResult {
	s.RuleSequence = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetRuleTag(v string) *RunContractResultGenerationResponseBodyOutputResult {
	s.RuleTag = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetRuleTitle(v string) *RunContractResultGenerationResponseBodyOutputResult {
	s.RuleTitle = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetSubRisks(v []*RunContractResultGenerationResponseBodyOutputResultSubRisks) *RunContractResultGenerationResponseBodyOutputResult {
	s.SubRisks = v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) Validate() error {
	if s.SubRisks != nil {
		for _, item := range s.SubRisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunContractResultGenerationResponseBodyOutputResultSubRisks struct {
	OriginalContent         *string `json:"originalContent,omitempty" xml:"originalContent,omitempty"`
	ResultContent           *string `json:"resultContent,omitempty" xml:"resultContent,omitempty"`
	ResultType              *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
	RiskBrief               *string `json:"riskBrief,omitempty" xml:"riskBrief,omitempty"`
	RiskClause              *string `json:"riskClause,omitempty" xml:"riskClause,omitempty"`
	RiskExplain             *string `json:"riskExplain,omitempty" xml:"riskExplain,omitempty"`
	StandardOriginalContent *string `json:"standardOriginalContent,omitempty" xml:"standardOriginalContent,omitempty"`
}

func (s RunContractResultGenerationResponseBodyOutputResultSubRisks) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationResponseBodyOutputResultSubRisks) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) GetOriginalContent() *string {
	return s.OriginalContent
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) GetResultContent() *string {
	return s.ResultContent
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) GetResultType() *string {
	return s.ResultType
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) GetRiskBrief() *string {
	return s.RiskBrief
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) GetRiskClause() *string {
	return s.RiskClause
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) GetRiskExplain() *string {
	return s.RiskExplain
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) GetStandardOriginalContent() *string {
	return s.StandardOriginalContent
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetOriginalContent(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.OriginalContent = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetResultContent(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.ResultContent = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetResultType(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.ResultType = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetRiskBrief(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.RiskBrief = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetRiskClause(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.RiskClause = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetRiskExplain(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.RiskExplain = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetStandardOriginalContent(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.StandardOriginalContent = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) Validate() error {
	return dara.Validate(s)
}

type RunContractResultGenerationResponseBodyUsage struct {
	// example:
	//
	// 5
	Input *int64 `json:"input,omitempty" xml:"input,omitempty"`
	// example:
	//
	// page
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s RunContractResultGenerationResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationResponseBodyUsage) GetInput() *int64 {
	return s.Input
}

func (s *RunContractResultGenerationResponseBodyUsage) GetUnit() *string {
	return s.Unit
}

func (s *RunContractResultGenerationResponseBodyUsage) SetInput(v int64) *RunContractResultGenerationResponseBodyUsage {
	s.Input = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyUsage) SetUnit(v string) *RunContractResultGenerationResponseBodyUsage {
	s.Unit = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
