// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomHotTopicBroadcastJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotTopicBroadcastConfig(v *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig) *SubmitCustomHotTopicBroadcastJobRequest
	GetHotTopicBroadcastConfig() *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig
	SetHotTopicVersion(v string) *SubmitCustomHotTopicBroadcastJobRequest
	GetHotTopicVersion() *string
	SetTopics(v []*string) *SubmitCustomHotTopicBroadcastJobRequest
	GetTopics() []*string
	SetWorkspaceId(v string) *SubmitCustomHotTopicBroadcastJobRequest
	GetWorkspaceId() *string
}

type SubmitCustomHotTopicBroadcastJobRequest struct {
	// The configuration for the news broadcast job.
	//
	// This parameter is required.
	HotTopicBroadcastConfig *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig `json:"HotTopicBroadcastConfig,omitempty" xml:"HotTopicBroadcastConfig,omitempty" type:"Struct"`
	// The version of the hot topic.
	//
	// example:
	//
	// 热点版本
	HotTopicVersion *string `json:"HotTopicVersion,omitempty" xml:"HotTopicVersion,omitempty"`
	// The topic filter.
	Topics []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
	// The unique identifier of the Alibaba Cloud Model Studio workspace. For more information, see [Get a workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitCustomHotTopicBroadcastJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomHotTopicBroadcastJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitCustomHotTopicBroadcastJobRequest) GetHotTopicBroadcastConfig() *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig {
	return s.HotTopicBroadcastConfig
}

func (s *SubmitCustomHotTopicBroadcastJobRequest) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *SubmitCustomHotTopicBroadcastJobRequest) GetTopics() []*string {
	return s.Topics
}

func (s *SubmitCustomHotTopicBroadcastJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitCustomHotTopicBroadcastJobRequest) SetHotTopicBroadcastConfig(v *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig) *SubmitCustomHotTopicBroadcastJobRequest {
	s.HotTopicBroadcastConfig = v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequest) SetHotTopicVersion(v string) *SubmitCustomHotTopicBroadcastJobRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequest) SetTopics(v []*string) *SubmitCustomHotTopicBroadcastJobRequest {
	s.Topics = v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequest) SetWorkspaceId(v string) *SubmitCustomHotTopicBroadcastJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequest) Validate() error {
	if s.HotTopicBroadcastConfig != nil {
		if err := s.HotTopicBroadcastConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig struct {
	// The configuration for the custom output style.
	//
	// This parameter is required.
	StepForCustomSummaryStyleConfig *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig `json:"StepForCustomSummaryStyleConfig,omitempty" xml:"StepForCustomSummaryStyleConfig,omitempty" type:"Struct"`
	// The configuration for the broadcast content.
	//
	// This parameter is required.
	StepForNewsBroadcastContentConfig *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig `json:"StepForNewsBroadcastContentConfig,omitempty" xml:"StepForNewsBroadcastContentConfig,omitempty" type:"Struct"`
}

func (s SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig) GoString() string {
	return s.String()
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig) GetStepForCustomSummaryStyleConfig() *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig {
	return s.StepForCustomSummaryStyleConfig
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig) GetStepForNewsBroadcastContentConfig() *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig {
	return s.StepForNewsBroadcastContentConfig
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig) SetStepForCustomSummaryStyleConfig(v *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig) *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig {
	s.StepForCustomSummaryStyleConfig = v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig) SetStepForNewsBroadcastContentConfig(v *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig) *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig {
	s.StepForNewsBroadcastContentConfig = v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfig) Validate() error {
	if s.StepForCustomSummaryStyleConfig != nil {
		if err := s.StepForCustomSummaryStyleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.StepForNewsBroadcastContentConfig != nil {
		if err := s.StepForNewsBroadcastContentConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig struct {
	// The number of images in the summary.
	//
	// example:
	//
	// 3
	SummaryImageCount *int32 `json:"SummaryImageCount,omitempty" xml:"SummaryImageCount,omitempty"`
	// The summary model.
	//
	// example:
	//
	// qwen-max
	SummaryModel *string `json:"SummaryModel,omitempty" xml:"SummaryModel,omitempty"`
	// The custom prompt for the summary.
	//
	// example:
	//
	// xxxx
	SummaryPrompt *string `json:"SummaryPrompt,omitempty" xml:"SummaryPrompt,omitempty"`
}

func (s SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig) GoString() string {
	return s.String()
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig) GetSummaryImageCount() *int32 {
	return s.SummaryImageCount
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig) GetSummaryModel() *string {
	return s.SummaryModel
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig) GetSummaryPrompt() *string {
	return s.SummaryPrompt
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig) SetSummaryImageCount(v int32) *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig {
	s.SummaryImageCount = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig) SetSummaryModel(v string) *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig {
	s.SummaryModel = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig) SetSummaryPrompt(v string) *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig {
	s.SummaryPrompt = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForCustomSummaryStyleConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig struct {
	// The list of selected channels.
	//
	// example:
	//
	// ["科技","经济","时政","娱乐"]
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The custom weights for hot topics.
	CustomHotValueWeights []*SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights `json:"CustomHotValueWeights,omitempty" xml:"CustomHotValueWeights,omitempty" type:"Repeated"`
	// The number of topics.
	//
	// example:
	//
	// 10
	TopicCount *int32 `json:"TopicCount,omitempty" xml:"TopicCount,omitempty"`
}

func (s SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig) GoString() string {
	return s.String()
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig) GetCategories() []*string {
	return s.Categories
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig) GetCustomHotValueWeights() []*SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights {
	return s.CustomHotValueWeights
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig) GetTopicCount() *int32 {
	return s.TopicCount
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig) SetCategories(v []*string) *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig {
	s.Categories = v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig) SetCustomHotValueWeights(v []*SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights) *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig {
	s.CustomHotValueWeights = v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig) SetTopicCount(v int32) *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig {
	s.TopicCount = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfig) Validate() error {
	if s.CustomHotValueWeights != nil {
		for _, item := range s.CustomHotValueWeights {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights struct {
	// The key of the dimension.
	//
	// example:
	//
	// views
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The name of the dimension.
	//
	// example:
	//
	// 维度名称
	DimensionName *string `json:"DimensionName,omitempty" xml:"DimensionName,omitempty"`
	// The weight.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights) GoString() string {
	return s.String()
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights) GetDimension() *string {
	return s.Dimension
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights) GetDimensionName() *string {
	return s.DimensionName
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights) GetWeight() *int32 {
	return s.Weight
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights) SetDimension(v string) *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights {
	s.Dimension = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights) SetDimensionName(v string) *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights {
	s.DimensionName = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights) SetWeight(v int32) *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights {
	s.Weight = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobRequestHotTopicBroadcastConfigStepForNewsBroadcastContentConfigCustomHotValueWeights) Validate() error {
	return dara.Validate(s)
}
