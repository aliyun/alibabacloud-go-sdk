// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunHotTopicSummaryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotTopicVersion(v string) *RunHotTopicSummaryShrinkRequest
	GetHotTopicVersion() *string
	SetStepForCustomSummaryStyleConfigShrink(v string) *RunHotTopicSummaryShrinkRequest
	GetStepForCustomSummaryStyleConfigShrink() *string
	SetTopicIdsShrink(v string) *RunHotTopicSummaryShrinkRequest
	GetTopicIdsShrink() *string
}

type RunHotTopicSummaryShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2024-10-16_8
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	// This parameter is required.
	StepForCustomSummaryStyleConfigShrink *string `json:"stepForCustomSummaryStyleConfig,omitempty" xml:"stepForCustomSummaryStyleConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxx
	TopicIdsShrink *string `json:"topicIds,omitempty" xml:"topicIds,omitempty"`
}

func (s RunHotTopicSummaryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicSummaryShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryShrinkRequest) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *RunHotTopicSummaryShrinkRequest) GetStepForCustomSummaryStyleConfigShrink() *string {
	return s.StepForCustomSummaryStyleConfigShrink
}

func (s *RunHotTopicSummaryShrinkRequest) GetTopicIdsShrink() *string {
	return s.TopicIdsShrink
}

func (s *RunHotTopicSummaryShrinkRequest) SetHotTopicVersion(v string) *RunHotTopicSummaryShrinkRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *RunHotTopicSummaryShrinkRequest) SetStepForCustomSummaryStyleConfigShrink(v string) *RunHotTopicSummaryShrinkRequest {
	s.StepForCustomSummaryStyleConfigShrink = &v
	return s
}

func (s *RunHotTopicSummaryShrinkRequest) SetTopicIdsShrink(v string) *RunHotTopicSummaryShrinkRequest {
	s.TopicIdsShrink = &v
	return s
}

func (s *RunHotTopicSummaryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
