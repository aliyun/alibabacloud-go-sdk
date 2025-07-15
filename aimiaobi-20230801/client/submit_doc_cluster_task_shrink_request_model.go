// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocClusterTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SubmitDocClusterTaskShrinkRequest
	GetAgentKey() *string
	SetDocumentsShrink(v string) *SubmitDocClusterTaskShrinkRequest
	GetDocumentsShrink() *string
	SetSummaryLength(v int32) *SubmitDocClusterTaskShrinkRequest
	GetSummaryLength() *int32
	SetTitleLength(v int32) *SubmitDocClusterTaskShrinkRequest
	GetTitleLength() *int32
	SetTopicCount(v int32) *SubmitDocClusterTaskShrinkRequest
	GetTopicCount() *int32
}

type SubmitDocClusterTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	DocumentsShrink *string `json:"Documents,omitempty" xml:"Documents,omitempty"`
	// example:
	//
	// 49
	SummaryLength *int32 `json:"SummaryLength,omitempty" xml:"SummaryLength,omitempty"`
	// example:
	//
	// 69
	TitleLength *int32 `json:"TitleLength,omitempty" xml:"TitleLength,omitempty"`
	// example:
	//
	// 15
	TopicCount *int32 `json:"TopicCount,omitempty" xml:"TopicCount,omitempty"`
}

func (s SubmitDocClusterTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocClusterTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocClusterTaskShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SubmitDocClusterTaskShrinkRequest) GetDocumentsShrink() *string {
	return s.DocumentsShrink
}

func (s *SubmitDocClusterTaskShrinkRequest) GetSummaryLength() *int32 {
	return s.SummaryLength
}

func (s *SubmitDocClusterTaskShrinkRequest) GetTitleLength() *int32 {
	return s.TitleLength
}

func (s *SubmitDocClusterTaskShrinkRequest) GetTopicCount() *int32 {
	return s.TopicCount
}

func (s *SubmitDocClusterTaskShrinkRequest) SetAgentKey(v string) *SubmitDocClusterTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitDocClusterTaskShrinkRequest) SetDocumentsShrink(v string) *SubmitDocClusterTaskShrinkRequest {
	s.DocumentsShrink = &v
	return s
}

func (s *SubmitDocClusterTaskShrinkRequest) SetSummaryLength(v int32) *SubmitDocClusterTaskShrinkRequest {
	s.SummaryLength = &v
	return s
}

func (s *SubmitDocClusterTaskShrinkRequest) SetTitleLength(v int32) *SubmitDocClusterTaskShrinkRequest {
	s.TitleLength = &v
	return s
}

func (s *SubmitDocClusterTaskShrinkRequest) SetTopicCount(v int32) *SubmitDocClusterTaskShrinkRequest {
	s.TopicCount = &v
	return s
}

func (s *SubmitDocClusterTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
