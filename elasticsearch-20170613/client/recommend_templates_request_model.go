// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecommendTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUsageScenario(v string) *RecommendTemplatesRequest
	GetUsageScenario() *string
}

type RecommendTemplatesRequest struct {
	// The scenario-specific templatetype used by the cluster. Valid values:
	//
	// - general: general-purpose scenario
	//
	// - analysisVisualization: data analytics scenario
	//
	// - dbAcceleration: database acceleration scenario
	//
	// - search: search scenario
	//
	// - log: log scenario
	//
	// > Commercial model instances support the general-purpose, data analytics, database acceleration, and search scenarios. Advanced Edition instances support only the log scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// general
	UsageScenario *string `json:"usageScenario,omitempty" xml:"usageScenario,omitempty"`
}

func (s RecommendTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s RecommendTemplatesRequest) GoString() string {
	return s.String()
}

func (s *RecommendTemplatesRequest) GetUsageScenario() *string {
	return s.UsageScenario
}

func (s *RecommendTemplatesRequest) SetUsageScenario(v string) *RecommendTemplatesRequest {
	s.UsageScenario = &v
	return s
}

func (s *RecommendTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
