// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectionStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetectionStatistic(v *GetDetectionStatisticResponseBodyDetectionStatistic) *GetDetectionStatisticResponseBody
	GetDetectionStatistic() *GetDetectionStatisticResponseBodyDetectionStatistic
	SetRequestId(v string) *GetDetectionStatisticResponseBody
	GetRequestId() *string
}

type GetDetectionStatisticResponseBody struct {
	// The detection rule count result.
	DetectionStatistic *GetDetectionStatisticResponseBodyDetectionStatistic `json:"DetectionStatistic,omitempty" xml:"DetectionStatistic,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6FB890AC-90B2-5EEA-845B-F7C86FB2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDetectionStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDetectionStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectionStatisticResponseBody) GetDetectionStatistic() *GetDetectionStatisticResponseBodyDetectionStatistic {
	return s.DetectionStatistic
}

func (s *GetDetectionStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDetectionStatisticResponseBody) SetDetectionStatistic(v *GetDetectionStatisticResponseBodyDetectionStatistic) *GetDetectionStatisticResponseBody {
	s.DetectionStatistic = v
	return s
}

func (s *GetDetectionStatisticResponseBody) SetRequestId(v string) *GetDetectionStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetectionStatisticResponseBody) Validate() error {
	if s.DetectionStatistic != nil {
		if err := s.DetectionStatistic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDetectionStatisticResponseBodyDetectionStatistic struct {
	AiPoweredAggregationRuleCount *int32 `json:"AiPoweredAggregationRuleCount,omitempty" xml:"AiPoweredAggregationRuleCount,omitempty"`
	// The number of online rules.
	//
	// example:
	//
	// 10
	DetectionRuleOnlineCount *int32 `json:"DetectionRuleOnlineCount,omitempty" xml:"DetectionRuleOnlineCount,omitempty"`
	// The number of rule templates.
	//
	// example:
	//
	// 20
	DetectionRuleTemplateCount *int32 `json:"DetectionRuleTemplateCount,omitempty" xml:"DetectionRuleTemplateCount,omitempty"`
	// The number of test rules.
	//
	// example:
	//
	// 5
	DetectionRuleTestCount *int32 `json:"DetectionRuleTestCount,omitempty" xml:"DetectionRuleTestCount,omitempty"`
	// The number of graph computing rules.
	//
	// example:
	//
	// 12
	GraphComputeRuleCount *int32 `json:"GraphComputeRuleCount,omitempty" xml:"GraphComputeRuleCount,omitempty"`
	// The number of alert pass-through rules.
	//
	// example:
	//
	// 2
	PassthroughRuleCount *int32 `json:"PassthroughRuleCount,omitempty" xml:"PassthroughRuleCount,omitempty"`
	// The number of similar aggregation rules.
	//
	// example:
	//
	// 6
	WindowRuleCount *int32 `json:"WindowRuleCount,omitempty" xml:"WindowRuleCount,omitempty"`
}

func (s GetDetectionStatisticResponseBodyDetectionStatistic) String() string {
	return dara.Prettify(s)
}

func (s GetDetectionStatisticResponseBodyDetectionStatistic) GoString() string {
	return s.String()
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) GetAiPoweredAggregationRuleCount() *int32 {
	return s.AiPoweredAggregationRuleCount
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) GetDetectionRuleOnlineCount() *int32 {
	return s.DetectionRuleOnlineCount
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) GetDetectionRuleTemplateCount() *int32 {
	return s.DetectionRuleTemplateCount
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) GetDetectionRuleTestCount() *int32 {
	return s.DetectionRuleTestCount
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) GetGraphComputeRuleCount() *int32 {
	return s.GraphComputeRuleCount
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) GetPassthroughRuleCount() *int32 {
	return s.PassthroughRuleCount
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) GetWindowRuleCount() *int32 {
	return s.WindowRuleCount
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) SetAiPoweredAggregationRuleCount(v int32) *GetDetectionStatisticResponseBodyDetectionStatistic {
	s.AiPoweredAggregationRuleCount = &v
	return s
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) SetDetectionRuleOnlineCount(v int32) *GetDetectionStatisticResponseBodyDetectionStatistic {
	s.DetectionRuleOnlineCount = &v
	return s
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) SetDetectionRuleTemplateCount(v int32) *GetDetectionStatisticResponseBodyDetectionStatistic {
	s.DetectionRuleTemplateCount = &v
	return s
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) SetDetectionRuleTestCount(v int32) *GetDetectionStatisticResponseBodyDetectionStatistic {
	s.DetectionRuleTestCount = &v
	return s
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) SetGraphComputeRuleCount(v int32) *GetDetectionStatisticResponseBodyDetectionStatistic {
	s.GraphComputeRuleCount = &v
	return s
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) SetPassthroughRuleCount(v int32) *GetDetectionStatisticResponseBodyDetectionStatistic {
	s.PassthroughRuleCount = &v
	return s
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) SetWindowRuleCount(v int32) *GetDetectionStatisticResponseBodyDetectionStatistic {
	s.WindowRuleCount = &v
	return s
}

func (s *GetDetectionStatisticResponseBodyDetectionStatistic) Validate() error {
	return dara.Validate(s)
}
