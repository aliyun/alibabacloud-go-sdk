// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetConfig interface {
	dara.Model
	String() string
	GoString() string
	SetInsights(v *InsightsConfig) *DatasetConfig
	GetInsights() *InsightsConfig
}

type DatasetConfig struct {
	Insights *InsightsConfig `json:"Insights,omitempty" xml:"Insights,omitempty"`
}

func (s DatasetConfig) String() string {
	return dara.Prettify(s)
}

func (s DatasetConfig) GoString() string {
	return s.String()
}

func (s *DatasetConfig) GetInsights() *InsightsConfig {
	return s.Insights
}

func (s *DatasetConfig) SetInsights(v *InsightsConfig) *DatasetConfig {
	s.Insights = v
	return s
}

func (s *DatasetConfig) Validate() error {
	if s.Insights != nil {
		if err := s.Insights.Validate(); err != nil {
			return err
		}
	}
	return nil
}
