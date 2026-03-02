// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedbackPbcReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *FeedbackReviewCmd) *FeedbackPbcReviewRequest
	GetBody() *FeedbackReviewCmd
}

type FeedbackPbcReviewRequest struct {
	Body *FeedbackReviewCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FeedbackPbcReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s FeedbackPbcReviewRequest) GoString() string {
	return s.String()
}

func (s *FeedbackPbcReviewRequest) GetBody() *FeedbackReviewCmd {
	return s.Body
}

func (s *FeedbackPbcReviewRequest) SetBody(v *FeedbackReviewCmd) *FeedbackPbcReviewRequest {
	s.Body = v
	return s
}

func (s *FeedbackPbcReviewRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
