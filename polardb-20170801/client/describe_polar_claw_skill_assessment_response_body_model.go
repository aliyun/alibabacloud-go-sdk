// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillAssessmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssessmentResult(v string) *DescribePolarClawSkillAssessmentResponseBody
	GetAssessmentResult() *string
	SetRequestId(v string) *DescribePolarClawSkillAssessmentResponseBody
	GetRequestId() *string
}

type DescribePolarClawSkillAssessmentResponseBody struct {
	// The assessment result.
	//
	// example:
	//
	// {"zhReportContent":"# Skill Security Assessment Report\\n\\n> **Risk Level:*	- ...}
	AssessmentResult *string `json:"AssessmentResult,omitempty" xml:"AssessmentResult,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolarClawSkillAssessmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillAssessmentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillAssessmentResponseBody) GetAssessmentResult() *string {
	return s.AssessmentResult
}

func (s *DescribePolarClawSkillAssessmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawSkillAssessmentResponseBody) SetAssessmentResult(v string) *DescribePolarClawSkillAssessmentResponseBody {
	s.AssessmentResult = &v
	return s
}

func (s *DescribePolarClawSkillAssessmentResponseBody) SetRequestId(v string) *DescribePolarClawSkillAssessmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawSkillAssessmentResponseBody) Validate() error {
	return dara.Validate(s)
}
