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
	SetReverseImage(v *ReverseImageConfig) *DatasetConfig
	GetReverseImage() *ReverseImageConfig
	SetSmartCluster(v *SmartClusterConfig) *DatasetConfig
	GetSmartCluster() *SmartClusterConfig
}

type DatasetConfig struct {
	// The content awareness configuration.
	Insights     *InsightsConfig     `json:"Insights,omitempty" xml:"Insights,omitempty"`
	ReverseImage *ReverseImageConfig `json:"ReverseImage,omitempty" xml:"ReverseImage,omitempty"`
	// The intelligent clustering configuration.
	SmartCluster *SmartClusterConfig `json:"SmartCluster,omitempty" xml:"SmartCluster,omitempty"`
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

func (s *DatasetConfig) GetReverseImage() *ReverseImageConfig {
	return s.ReverseImage
}

func (s *DatasetConfig) GetSmartCluster() *SmartClusterConfig {
	return s.SmartCluster
}

func (s *DatasetConfig) SetInsights(v *InsightsConfig) *DatasetConfig {
	s.Insights = v
	return s
}

func (s *DatasetConfig) SetReverseImage(v *ReverseImageConfig) *DatasetConfig {
	s.ReverseImage = v
	return s
}

func (s *DatasetConfig) SetSmartCluster(v *SmartClusterConfig) *DatasetConfig {
	s.SmartCluster = v
	return s
}

func (s *DatasetConfig) Validate() error {
	if s.Insights != nil {
		if err := s.Insights.Validate(); err != nil {
			return err
		}
	}
	if s.ReverseImage != nil {
		if err := s.ReverseImage.Validate(); err != nil {
			return err
		}
	}
	if s.SmartCluster != nil {
		if err := s.SmartCluster.Validate(); err != nil {
			return err
		}
	}
	return nil
}
