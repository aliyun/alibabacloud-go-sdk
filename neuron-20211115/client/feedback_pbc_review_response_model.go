// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedbackPbcReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FeedbackPbcReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FeedbackPbcReviewResponse
	GetStatusCode() *int32
	SetBody(v *PbcReview) *FeedbackPbcReviewResponse
	GetBody() *PbcReview
}

type FeedbackPbcReviewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcReview         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FeedbackPbcReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s FeedbackPbcReviewResponse) GoString() string {
	return s.String()
}

func (s *FeedbackPbcReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FeedbackPbcReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FeedbackPbcReviewResponse) GetBody() *PbcReview {
	return s.Body
}

func (s *FeedbackPbcReviewResponse) SetHeaders(v map[string]*string) *FeedbackPbcReviewResponse {
	s.Headers = v
	return s
}

func (s *FeedbackPbcReviewResponse) SetStatusCode(v int32) *FeedbackPbcReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *FeedbackPbcReviewResponse) SetBody(v *PbcReview) *FeedbackPbcReviewResponse {
	s.Body = v
	return s
}

func (s *FeedbackPbcReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
