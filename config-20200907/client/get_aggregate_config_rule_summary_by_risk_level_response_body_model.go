// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRuleSummaryByRiskLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleSummaries(v []*GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) *GetAggregateConfigRuleSummaryByRiskLevelResponseBody
	GetConfigRuleSummaries() []*GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries
	SetRequestId(v string) *GetAggregateConfigRuleSummaryByRiskLevelResponseBody
	GetRequestId() *string
}

type GetAggregateConfigRuleSummaryByRiskLevelResponseBody struct {
	ConfigRuleSummaries []*GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries `json:"ConfigRuleSummaries,omitempty" xml:"ConfigRuleSummaries,omitempty" type:"Repeated"`
	RequestId           *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateConfigRuleSummaryByRiskLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleSummaryByRiskLevelResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBody) GetConfigRuleSummaries() []*GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	return s.ConfigRuleSummaries
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBody) SetConfigRuleSummaries(v []*GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) *GetAggregateConfigRuleSummaryByRiskLevelResponseBody {
	s.ConfigRuleSummaries = v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBody) SetRequestId(v string) *GetAggregateConfigRuleSummaryByRiskLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBody) Validate() error {
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

type GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries struct {
	CompliantCount    *int32 `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	NonCompliantCount *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	RiskLevel         *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) GetCompliantCount() *int32 {
	return s.CompliantCount
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) GetNonCompliantCount() *int32 {
	return s.NonCompliantCount
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) SetCompliantCount(v int32) *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	s.CompliantCount = &v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) SetNonCompliantCount(v int32) *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	s.NonCompliantCount = &v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) SetRiskLevel(v int32) *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	s.RiskLevel = &v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) Validate() error {
	return dara.Validate(s)
}
