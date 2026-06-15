// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckModelFeatureFGFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFGCheckResults(v []*CheckModelFeatureFGFeatureResponseBodyFGCheckResults) *CheckModelFeatureFGFeatureResponseBody
	GetFGCheckResults() []*CheckModelFeatureFGFeatureResponseBodyFGCheckResults
	SetRequestId(v string) *CheckModelFeatureFGFeatureResponseBody
	GetRequestId() *string
}

type CheckModelFeatureFGFeatureResponseBody struct {
	// The FG configuration check results.
	FGCheckResults []*CheckModelFeatureFGFeatureResponseBodyFGCheckResults `json:"FGCheckResults,omitempty" xml:"FGCheckResults,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ED4DEA2F-F216-57F0-AE28-08D791233280
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CheckModelFeatureFGFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckModelFeatureFGFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *CheckModelFeatureFGFeatureResponseBody) GetFGCheckResults() []*CheckModelFeatureFGFeatureResponseBodyFGCheckResults {
	return s.FGCheckResults
}

func (s *CheckModelFeatureFGFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckModelFeatureFGFeatureResponseBody) SetFGCheckResults(v []*CheckModelFeatureFGFeatureResponseBodyFGCheckResults) *CheckModelFeatureFGFeatureResponseBody {
	s.FGCheckResults = v
	return s
}

func (s *CheckModelFeatureFGFeatureResponseBody) SetRequestId(v string) *CheckModelFeatureFGFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckModelFeatureFGFeatureResponseBody) Validate() error {
	if s.FGCheckResults != nil {
		for _, item := range s.FGCheckResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckModelFeatureFGFeatureResponseBodyFGCheckResults struct {
	// The error message.
	//
	// example:
	//
	// f1[1]: these lookup_features\\"s LookupValueFeature(key) not exist in model features
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The rule code for the check.
	//
	// - SeqSubEx—Are the sub-features of sequence features in the FG present in the model feature?
	//
	// - LookupKeyEx—Does the lookup key feature of query features exist in the model feature?
	//
	// - FGSetEx—Are all FG features present in the model feature?
	//
	// - ModelSetEx—Are all model feature features present in the FG?
	//
	// example:
	//
	// SeqSubEx
	RuleCode *string `json:"RuleCode,omitempty" xml:"RuleCode,omitempty"`
	// Whether the rule passes.
	//
	// example:
	//
	// True
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckModelFeatureFGFeatureResponseBodyFGCheckResults) String() string {
	return dara.Prettify(s)
}

func (s CheckModelFeatureFGFeatureResponseBodyFGCheckResults) GoString() string {
	return s.String()
}

func (s *CheckModelFeatureFGFeatureResponseBodyFGCheckResults) GetMessage() *string {
	return s.Message
}

func (s *CheckModelFeatureFGFeatureResponseBodyFGCheckResults) GetRuleCode() *string {
	return s.RuleCode
}

func (s *CheckModelFeatureFGFeatureResponseBodyFGCheckResults) GetStatus() *bool {
	return s.Status
}

func (s *CheckModelFeatureFGFeatureResponseBodyFGCheckResults) SetMessage(v string) *CheckModelFeatureFGFeatureResponseBodyFGCheckResults {
	s.Message = &v
	return s
}

func (s *CheckModelFeatureFGFeatureResponseBodyFGCheckResults) SetRuleCode(v string) *CheckModelFeatureFGFeatureResponseBodyFGCheckResults {
	s.RuleCode = &v
	return s
}

func (s *CheckModelFeatureFGFeatureResponseBodyFGCheckResults) SetStatus(v bool) *CheckModelFeatureFGFeatureResponseBodyFGCheckResults {
	s.Status = &v
	return s
}

func (s *CheckModelFeatureFGFeatureResponseBodyFGCheckResults) Validate() error {
	return dara.Validate(s)
}
