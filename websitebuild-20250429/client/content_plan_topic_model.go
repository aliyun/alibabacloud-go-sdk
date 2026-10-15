// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContentPlanTopic interface {
	dara.Model
	String() string
	GoString() string
	SetArticleConversationId(v string) *ContentPlanTopic
	GetArticleConversationId() *string
	SetArticleId(v string) *ContentPlanTopic
	GetArticleId() *string
	SetArticleStatus(v string) *ContentPlanTopic
	GetArticleStatus() *string
	SetDiagnosisReferences(v []*string) *ContentPlanTopic
	GetDiagnosisReferences() []*string
	SetDirection(v string) *ContentPlanTopic
	GetDirection() *string
	SetFailReason(v string) *ContentPlanTopic
	GetFailReason() *string
	SetPlanId(v string) *ContentPlanTopic
	GetPlanId() *string
	SetProjectId(v string) *ContentPlanTopic
	GetProjectId() *string
	SetRationale(v string) *ContentPlanTopic
	GetRationale() *string
	SetRecommendedChannels(v []*string) *ContentPlanTopic
	GetRecommendedChannels() []*string
	SetScheduledAt(v string) *ContentPlanTopic
	GetScheduledAt() *string
	SetSkipReason(v string) *ContentPlanTopic
	GetSkipReason() *string
	SetStatus(v string) *ContentPlanTopic
	GetStatus() *string
	SetTitle(v string) *ContentPlanTopic
	GetTitle() *string
	SetTopicId(v string) *ContentPlanTopic
	GetTopicId() *string
	SetWritingDirection(v string) *ContentPlanTopic
	GetWritingDirection() *string
}

type ContentPlanTopic struct {
	// The article generation session ID.
	//
	// example:
	//
	// conv_8a2b3c
	ArticleConversationId *string `json:"ArticleConversationId,omitempty" xml:"ArticleConversationId,omitempty"`
	// The generated article ID.
	//
	// example:
	//
	// art_9f2k3d
	ArticleId *string `json:"ArticleId,omitempty" xml:"ArticleId,omitempty"`
	// The status of the associated article. Valid values: DRAFT, PUBLISHED, and PUBLISH_FAILED. This parameter is empty if the article is not generated.
	//
	// example:
	//
	// PUBLISHED
	ArticleStatus *string `json:"ArticleStatus,omitempty" xml:"ArticleStatus,omitempty"`
	// The GEO diagnostic item references.
	DiagnosisReferences []*string `json:"DiagnosisReferences,omitempty" xml:"DiagnosisReferences,omitempty" type:"Repeated"`
	// The content direction.
	//
	// example:
	//
	// Introduction to product capabilities
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The reason for the execution failure.
	//
	// example:
	//
	// Article generation failed. Try again later
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// The content plan ID.
	//
	// example:
	//
	// plan_3f7a1c2e9b7d4a5f8c0e1d2b3a4c5d6e
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The workspace ID. This refers to the GEO workspace for brand marketing, which is the appWorkspaceId.
	//
	// example:
	//
	// d4e5f60718293a4b5c6d7e8f90a1b2c3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The rationale for topic selection.
	//
	// example:
	//
	// Users show high interest in AI website builder capabilities, making it a suitable entry point for brand awareness
	Rationale *string `json:"Rationale,omitempty" xml:"Rationale,omitempty"`
	// The recommended publishing channels.
	RecommendedChannels []*string `json:"RecommendedChannels,omitempty" xml:"RecommendedChannels,omitempty" type:"Repeated"`
	// The scheduled publish date in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2026-10-12
	ScheduledAt *string `json:"ScheduledAt,omitempty" xml:"ScheduledAt,omitempty"`
	// The reason for skipping.
	//
	// example:
	//
	// This topic duplicates a published article and has been skipped
	SkipReason *string `json:"SkipReason,omitempty" xml:"SkipReason,omitempty"`
	// The topic status.
	//
	// example:
	//
	// PENDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The content title.
	//
	// example:
	//
	// Analysis of the core capabilities of Wanxiaozhi AI website builder
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The topic ID.
	//
	// example:
	//
	// topic_3f7a1c2e9b7d4a5f8c0e1d2b3a4c5d6e
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// The writing direction, which only describes the structure and perspective.
	//
	// example:
	//
	// First summarize the product positioning, then introduce the core capabilities point by point, and finally provide the applicable scenarios
	WritingDirection *string `json:"WritingDirection,omitempty" xml:"WritingDirection,omitempty"`
}

func (s ContentPlanTopic) String() string {
	return dara.Prettify(s)
}

func (s ContentPlanTopic) GoString() string {
	return s.String()
}

func (s *ContentPlanTopic) GetArticleConversationId() *string {
	return s.ArticleConversationId
}

func (s *ContentPlanTopic) GetArticleId() *string {
	return s.ArticleId
}

func (s *ContentPlanTopic) GetArticleStatus() *string {
	return s.ArticleStatus
}

func (s *ContentPlanTopic) GetDiagnosisReferences() []*string {
	return s.DiagnosisReferences
}

func (s *ContentPlanTopic) GetDirection() *string {
	return s.Direction
}

func (s *ContentPlanTopic) GetFailReason() *string {
	return s.FailReason
}

func (s *ContentPlanTopic) GetPlanId() *string {
	return s.PlanId
}

func (s *ContentPlanTopic) GetProjectId() *string {
	return s.ProjectId
}

func (s *ContentPlanTopic) GetRationale() *string {
	return s.Rationale
}

func (s *ContentPlanTopic) GetRecommendedChannels() []*string {
	return s.RecommendedChannels
}

func (s *ContentPlanTopic) GetScheduledAt() *string {
	return s.ScheduledAt
}

func (s *ContentPlanTopic) GetSkipReason() *string {
	return s.SkipReason
}

func (s *ContentPlanTopic) GetStatus() *string {
	return s.Status
}

func (s *ContentPlanTopic) GetTitle() *string {
	return s.Title
}

func (s *ContentPlanTopic) GetTopicId() *string {
	return s.TopicId
}

func (s *ContentPlanTopic) GetWritingDirection() *string {
	return s.WritingDirection
}

func (s *ContentPlanTopic) SetArticleConversationId(v string) *ContentPlanTopic {
	s.ArticleConversationId = &v
	return s
}

func (s *ContentPlanTopic) SetArticleId(v string) *ContentPlanTopic {
	s.ArticleId = &v
	return s
}

func (s *ContentPlanTopic) SetArticleStatus(v string) *ContentPlanTopic {
	s.ArticleStatus = &v
	return s
}

func (s *ContentPlanTopic) SetDiagnosisReferences(v []*string) *ContentPlanTopic {
	s.DiagnosisReferences = v
	return s
}

func (s *ContentPlanTopic) SetDirection(v string) *ContentPlanTopic {
	s.Direction = &v
	return s
}

func (s *ContentPlanTopic) SetFailReason(v string) *ContentPlanTopic {
	s.FailReason = &v
	return s
}

func (s *ContentPlanTopic) SetPlanId(v string) *ContentPlanTopic {
	s.PlanId = &v
	return s
}

func (s *ContentPlanTopic) SetProjectId(v string) *ContentPlanTopic {
	s.ProjectId = &v
	return s
}

func (s *ContentPlanTopic) SetRationale(v string) *ContentPlanTopic {
	s.Rationale = &v
	return s
}

func (s *ContentPlanTopic) SetRecommendedChannels(v []*string) *ContentPlanTopic {
	s.RecommendedChannels = v
	return s
}

func (s *ContentPlanTopic) SetScheduledAt(v string) *ContentPlanTopic {
	s.ScheduledAt = &v
	return s
}

func (s *ContentPlanTopic) SetSkipReason(v string) *ContentPlanTopic {
	s.SkipReason = &v
	return s
}

func (s *ContentPlanTopic) SetStatus(v string) *ContentPlanTopic {
	s.Status = &v
	return s
}

func (s *ContentPlanTopic) SetTitle(v string) *ContentPlanTopic {
	s.Title = &v
	return s
}

func (s *ContentPlanTopic) SetTopicId(v string) *ContentPlanTopic {
	s.TopicId = &v
	return s
}

func (s *ContentPlanTopic) SetWritingDirection(v string) *ContentPlanTopic {
	s.WritingDirection = &v
	return s
}

func (s *ContentPlanTopic) Validate() error {
	return dara.Validate(s)
}
