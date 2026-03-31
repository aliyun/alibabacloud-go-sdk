// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRuleSummaryByRiskLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleSummaries(v []*GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) *GetConfigRuleSummaryByRiskLevelResponseBody
	GetConfigRuleSummaries() []*GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries
	SetRequestId(v string) *GetConfigRuleSummaryByRiskLevelResponseBody
	GetRequestId() *string
}

type GetConfigRuleSummaryByRiskLevelResponseBody struct {
	// The summary of compliance evaluation results by rule risk level.
	ConfigRuleSummaries []*GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries `json:"ConfigRuleSummaries,omitempty" xml:"ConfigRuleSummaries,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// A3CED98C-DE65-46AC-B2D2-04A4A9AB5B36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConfigRuleSummaryByRiskLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleSummaryByRiskLevelResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBody) GetConfigRuleSummaries() []*GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	return s.ConfigRuleSummaries
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBody) SetConfigRuleSummaries(v []*GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) *GetConfigRuleSummaryByRiskLevelResponseBody {
	s.ConfigRuleSummaries = v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBody) SetRequestId(v string) *GetConfigRuleSummaryByRiskLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBody) Validate() error {
	if s.ConfigRuleSummaries != nil {
		for _, item := range s.ConfigRuleSummaries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries struct {
	// The number of rules against which specific resources are evaluated as compliant.
	//
	// example:
	//
	// 3
	CompliantCount *int32 `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	// The number of rules against which specific resources are evaluated as non-compliant.
	//
	// example:
	//
	// 1
	NonCompliantCount *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	// The risk level of the resources that are not compliant with the rules. Valid values:
	//
	// 	- 1: high risk level.
	//
	// 	- 2: medium risk level.
	//
	// 	- 3: low risk level.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) GoString() string {
	return s.String()
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) GetCompliantCount() *int32 {
	return s.CompliantCount
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) GetNonCompliantCount() *int32 {
	return s.NonCompliantCount
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) SetCompliantCount(v int32) *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	s.CompliantCount = &v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) SetNonCompliantCount(v int32) *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	s.NonCompliantCount = &v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) SetRiskLevel(v int32) *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	s.RiskLevel = &v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) Validate() error {
	return dara.Validate(s)
}
