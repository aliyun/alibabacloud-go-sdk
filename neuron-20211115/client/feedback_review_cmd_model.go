// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedbackReviewCmd interface {
	dara.Model
	String() string
	GoString() string
	SetFeedbackPbcAPI(v string) *FeedbackReviewCmd
	GetFeedbackPbcAPI() *string
	SetFeedbackPbcInstruction(v string) *FeedbackReviewCmd
	GetFeedbackPbcInstruction() *string
	SetFeedbackPbcSchema(v string) *FeedbackReviewCmd
	GetFeedbackPbcSchema() *string
	SetFeedbackSecurity(v string) *FeedbackReviewCmd
	GetFeedbackSecurity() *string
	SetFeedbackServiceAvailable(v string) *FeedbackReviewCmd
	GetFeedbackServiceAvailable() *string
	SetReviewId(v int64) *FeedbackReviewCmd
	GetReviewId() *int64
	SetStatus(v string) *FeedbackReviewCmd
	GetStatus() *string
}

type FeedbackReviewCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 符合规范
	FeedbackPbcAPI *string `json:"feedbackPbcAPI,omitempty" xml:"feedbackPbcAPI,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 符合规范
	FeedbackPbcInstruction *string `json:"feedbackPbcInstruction,omitempty" xml:"feedbackPbcInstruction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 符合规范
	FeedbackPbcSchema *string `json:"feedbackPbcSchema,omitempty" xml:"feedbackPbcSchema,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 服务安全
	FeedbackSecurity *string `json:"feedbackSecurity,omitempty" xml:"feedbackSecurity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 服务可运行
	FeedbackServiceAvailable *string `json:"feedbackServiceAvailable,omitempty" xml:"feedbackServiceAvailable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReviewId *int64 `json:"reviewId,omitempty" xml:"reviewId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// REVIEW_FAILED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s FeedbackReviewCmd) String() string {
	return dara.Prettify(s)
}

func (s FeedbackReviewCmd) GoString() string {
	return s.String()
}

func (s *FeedbackReviewCmd) GetFeedbackPbcAPI() *string {
	return s.FeedbackPbcAPI
}

func (s *FeedbackReviewCmd) GetFeedbackPbcInstruction() *string {
	return s.FeedbackPbcInstruction
}

func (s *FeedbackReviewCmd) GetFeedbackPbcSchema() *string {
	return s.FeedbackPbcSchema
}

func (s *FeedbackReviewCmd) GetFeedbackSecurity() *string {
	return s.FeedbackSecurity
}

func (s *FeedbackReviewCmd) GetFeedbackServiceAvailable() *string {
	return s.FeedbackServiceAvailable
}

func (s *FeedbackReviewCmd) GetReviewId() *int64 {
	return s.ReviewId
}

func (s *FeedbackReviewCmd) GetStatus() *string {
	return s.Status
}

func (s *FeedbackReviewCmd) SetFeedbackPbcAPI(v string) *FeedbackReviewCmd {
	s.FeedbackPbcAPI = &v
	return s
}

func (s *FeedbackReviewCmd) SetFeedbackPbcInstruction(v string) *FeedbackReviewCmd {
	s.FeedbackPbcInstruction = &v
	return s
}

func (s *FeedbackReviewCmd) SetFeedbackPbcSchema(v string) *FeedbackReviewCmd {
	s.FeedbackPbcSchema = &v
	return s
}

func (s *FeedbackReviewCmd) SetFeedbackSecurity(v string) *FeedbackReviewCmd {
	s.FeedbackSecurity = &v
	return s
}

func (s *FeedbackReviewCmd) SetFeedbackServiceAvailable(v string) *FeedbackReviewCmd {
	s.FeedbackServiceAvailable = &v
	return s
}

func (s *FeedbackReviewCmd) SetReviewId(v int64) *FeedbackReviewCmd {
	s.ReviewId = &v
	return s
}

func (s *FeedbackReviewCmd) SetStatus(v string) *FeedbackReviewCmd {
	s.Status = &v
	return s
}

func (s *FeedbackReviewCmd) Validate() error {
	return dara.Validate(s)
}
