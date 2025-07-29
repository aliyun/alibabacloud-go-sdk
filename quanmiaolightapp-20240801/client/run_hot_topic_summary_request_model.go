// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunHotTopicSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotTopicVersion(v string) *RunHotTopicSummaryRequest
	GetHotTopicVersion() *string
	SetStepForCustomSummaryStyleConfig(v *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) *RunHotTopicSummaryRequest
	GetStepForCustomSummaryStyleConfig() *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig
	SetTopicIds(v []*string) *RunHotTopicSummaryRequest
	GetTopicIds() []*string
}

type RunHotTopicSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2024-10-16_8
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	// This parameter is required.
	StepForCustomSummaryStyleConfig *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig `json:"stepForCustomSummaryStyleConfig,omitempty" xml:"stepForCustomSummaryStyleConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxx
	TopicIds []*string `json:"topicIds,omitempty" xml:"topicIds,omitempty" type:"Repeated"`
}

func (s RunHotTopicSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicSummaryRequest) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryRequest) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *RunHotTopicSummaryRequest) GetStepForCustomSummaryStyleConfig() *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig {
	return s.StepForCustomSummaryStyleConfig
}

func (s *RunHotTopicSummaryRequest) GetTopicIds() []*string {
	return s.TopicIds
}

func (s *RunHotTopicSummaryRequest) SetHotTopicVersion(v string) *RunHotTopicSummaryRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *RunHotTopicSummaryRequest) SetStepForCustomSummaryStyleConfig(v *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) *RunHotTopicSummaryRequest {
	s.StepForCustomSummaryStyleConfig = v
	return s
}

func (s *RunHotTopicSummaryRequest) SetTopicIds(v []*string) *RunHotTopicSummaryRequest {
	s.TopicIds = v
	return s
}

func (s *RunHotTopicSummaryRequest) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig struct {
	// example:
	//
	// 2
	SummaryImageCount *int32 `json:"summaryImageCount,omitempty" xml:"summaryImageCount,omitempty"`
	// example:
	//
	// qwen-max
	SummaryModel *string `json:"summaryModel,omitempty" xml:"summaryModel,omitempty"`
	// example:
	//
	// xxxx
	SummaryPrompt *string `json:"summaryPrompt,omitempty" xml:"summaryPrompt,omitempty"`
}

func (s RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) GetSummaryImageCount() *int32 {
	return s.SummaryImageCount
}

func (s *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) GetSummaryModel() *string {
	return s.SummaryModel
}

func (s *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) GetSummaryPrompt() *string {
	return s.SummaryPrompt
}

func (s *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) SetSummaryImageCount(v int32) *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig {
	s.SummaryImageCount = &v
	return s
}

func (s *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) SetSummaryModel(v string) *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig {
	s.SummaryModel = &v
	return s
}

func (s *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) SetSummaryPrompt(v string) *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig {
	s.SummaryPrompt = &v
	return s
}

func (s *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) Validate() error {
	return dara.Validate(s)
}
