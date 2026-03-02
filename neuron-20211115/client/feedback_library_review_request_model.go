// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedbackLibraryReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprove(v string) *FeedbackLibraryReviewRequest
	GetApprove() *string
	SetFeedback(v string) *FeedbackLibraryReviewRequest
	GetFeedback() *string
	SetReviewId(v int64) *FeedbackLibraryReviewRequest
	GetReviewId() *int64
}

type FeedbackLibraryReviewRequest struct {
	Approve  *string `json:"approve,omitempty" xml:"approve,omitempty"`
	Feedback *string `json:"feedback,omitempty" xml:"feedback,omitempty"`
	ReviewId *int64  `json:"reviewId,omitempty" xml:"reviewId,omitempty"`
}

func (s FeedbackLibraryReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s FeedbackLibraryReviewRequest) GoString() string {
	return s.String()
}

func (s *FeedbackLibraryReviewRequest) GetApprove() *string {
	return s.Approve
}

func (s *FeedbackLibraryReviewRequest) GetFeedback() *string {
	return s.Feedback
}

func (s *FeedbackLibraryReviewRequest) GetReviewId() *int64 {
	return s.ReviewId
}

func (s *FeedbackLibraryReviewRequest) SetApprove(v string) *FeedbackLibraryReviewRequest {
	s.Approve = &v
	return s
}

func (s *FeedbackLibraryReviewRequest) SetFeedback(v string) *FeedbackLibraryReviewRequest {
	s.Feedback = &v
	return s
}

func (s *FeedbackLibraryReviewRequest) SetReviewId(v int64) *FeedbackLibraryReviewRequest {
	s.ReviewId = &v
	return s
}

func (s *FeedbackLibraryReviewRequest) Validate() error {
	return dara.Validate(s)
}
